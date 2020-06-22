package main

//go:generate go run main.go -name=String -type=string

// The https://www.pastraiser.com/cpu/gameboy/gameboy_opcodes.html will guide your way.

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

const jsonFile = "./opcodes.json"

// opcode represents the data loaded from the accompanying JSON file
type opcode struct {
	ExtendedMnemonic string
	CBPrefixed       bool
	Mnemonic         string   `json:"mnemonic"`
	Length           uint     `json:"length"`
	Cycles           []uint   `json:"cycles"`
	Flags            []string `json:"flags"`
	Addr             string   `json:"addr"`
	RawOperand1      string   `json:"operand1"`
	RawOperand2      string   `json:"operand2"`
}

func (o *opcode) Operand1() string {
	// Matches (a16) and (d16)
	if strings.HasSuffix(o.RawOperand1, "16)") {
		// ("+o.operand1+")
		return "(\"+o.operand1+\")"
	} else if strings.HasSuffix(o.RawOperand1, "16") {
		return "o.operand1"
	}

	return o.RawOperand1
}

func (o *opcode) Operand2() string {
	// Matches (a16) and (d16)
	if strings.HasSuffix(o.RawOperand2, "16)") {
		// ("+o.operand1+")
		return "(\"+o.operand1+\")"
	} else if strings.HasSuffix(o.RawOperand2, "16") {
		return "o.operand1"
	}

	return o.RawOperand2
}

type opcodes struct {
	Unprefixed map[string]*opcode `json:"unprefixed"`

	// An optional prefix byte may appear before the opcode, changing its
	// meaning and causing the Z80 to look up the opcode in a different
	// bank of instructions. The prefix byte, if present, may have the values
	// CB. Although there are opcodes which have these values too, there is
	// no ambiguity: the first byte in the instruction, if it has one of
	// these values, is always a prefix byte.
	CBPrefixed map[string]*opcode `json:"cbprefixed"`
}

// LD A,(C)    has alternative mnemonic LD A,($FF00+C)
// LD C,(A)    has alternative mnemonic LD ($FF00+C),A
// LDH A,(a8)  has alternative mnemonic LD A,($FF00+a8)
// LDH (a8),A  has alternative mnemonic LD ($FF00+a8),A
// LD A,(HL+)  has alternative mnemonic LD A,(HLI) or LDI A,(HL)
// LD (HL+),A  has alternative mnemonic LD (HLI),A or LDI (HL),A
// LD A,(HL-)  has alternative mnemonic LD A,(HLD) or LDD A,(HL)
// LD (HL-),A  has alternative mnemonic LD (HLD),A or LDD (HL),A
// LD HL,SP+r8 has alternative mnemonic LDHL SP,r8

//  LDI (HL),A ; Write A to (HL) and increment HL
//  LDD (HL),A ; Write A to (HL) and decrement HL
//  LDI A,(HL) ; Write (HL) to A and increment HL
//  LDD A,(HL) ; Write (HL) to A and decrement HL

// cleanMnemonic makes the mnemonic a legal golang type name.
func cleanMnemonic(s string) string {
	switch true {
	case s == "":
		return ""
	case s == "(HL)":
		// The H and L registers are special due to the fact that they are
		// extensively used for indirect addressing as register pair HL.
		// Indirect Addressing is when instead of specifying an specific address
		// for an operation, you could just use the 16-bit value in HL as an
		// address. It's pretty handy for things like address calculations when
		// you need to access an array of value for example. This is the ONLY
		// register pair that can be used indirectly in the instructions ADC,
		return "HLPtr"
	case s == "(HL+)":
		return "HLPtrInc"
	case s == "(HL-)":
		return "HLPtrDec"
	case s[0] == '(':
		if s[len(s)-1] != ')' {
			panic(fmt.Sprintf("%q is not a valid opcode argument, no matching paren.", s))
		}
		return s[1:len(s)-1] + "Deref"
	case s == "SP+r8":
		return "SP_plus_r8"
	default:
		return s
	}
}

func cleanMnemonics(opcodes map[string]*opcode) {
	for k, v := range opcodes {
		var args []string
		args = append(args, cleanMnemonic(v.Mnemonic))

		if v.RawOperand1 != "" {
			args = append(args, cleanMnemonic(v.RawOperand1))
		}

		if v.RawOperand2 != "" {
			args = append(args, cleanMnemonic(v.RawOperand2))
		}

		extended := strings.Join(args, "_")

		opcodes[k].ExtendedMnemonic = extended
	}
}

// Parse parses the opcodes.json file into nice struct
// to be used by the codgen tool
func parse() (*opcodes, error) {
	jsonFile, err := os.Open(jsonFile)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	fmt.Printf("Succesfully opened %s\n", jsonFile.Name())

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll(%q) error: %v", jsonFile.Name(), err)
	}

	var data opcodes
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, fmt.Errorf("Unmarshal error: %v", err)
	}

	return &data, nil
}

type validateOpcode func(*opcode) error

func main() {
	opcodes, err := parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cleanMnemonics(opcodes.Unprefixed)
	cleanMnemonics(opcodes.CBPrefixed)

	for _, val := range opcodes.CBPrefixed {
		val.CBPrefixed = true
	}

	validate := func(f validateOpcode, opcodes map[string]*opcode) {
		for _, o := range opcodes {
			if err := f(o); err != nil {
				panic(fmt.Sprintf("Error validating opcode %v error: %v", o.Addr, err))
			}
		}
	}

	for _, f := range []validateOpcode{
		addMnemonicHasOnlyOneCycle,
	} {
		validate(f, opcodes.Unprefixed)
		validate(f, opcodes.CBPrefixed)
	}

	outFile, err := os.Create("../opcodes_generated.go")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer outFile.Close()

	templates := generateTemplate()
	if err := templates.ExecuteTemplate(outFile, "opcodes.tmpl", opcodes); err != nil {
		log.Fatalf("Error rendering template %v: %v", templates, err)
	}
	if err := outFile.Close(); err != nil {
		log.Fatalf("Unable to close pipe", err)
	}

	cmd := exec.Command("gofmt", "-w", outFile.Name())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func handleStringOfOperand(opcode opcode, operandID int) string {
	var rawOperand string
	if operandID == 1 {
		rawOperand = opcode.RawOperand1
	} else if operandID == 2 {
		rawOperand = opcode.RawOperand2
	}
	switch rawOperand {
	case "(a8)":
		return fmt.Sprintf(`fmt.Sprintf("(%%X)", o.operand%d)`, operandID)
	case "d8":
		return fmt.Sprintf(`fmt.Sprintf("%%X", o.operand%d)`, operandID)
	case "r8", "SP+r8":
		return fmt.Sprintf(`fmt.Sprintf("SP+%%X", o.operand%d)`, operandID)
	case "a16", "(a16)":
		return fmt.Sprintf(`fmt.Sprintf("%%X", o.operand%d)`, operandID)
	case "d16":
		return fmt.Sprintf(`fmt.Sprintf("%%X", o.operand%d)`, operandID)
	case "HL", "(HL)", "(HL-)", "(HL+)",
		"A", "AF", "B", "C", "(C)", "BC", "(BC)", "D", "E", "DE", "(DE)",
		"H", "L", "F",
		"SP", "PC",
		"Z",
		"NZ",
		"NC",
		"",
		"0", "1", "2", "3", "4", "5", "6", "7",
		"00H", "08H", "10H", "18H", "20H", "28H", "30H", "38H",
		"CB" /* CB prefixed opcodes, not the register C+B */ :
		return fmt.Sprintf(`"%s" /* operand%d */`, rawOperand, operandID)
	default:
		panic(fmt.Sprintf("You need to implement an stringer for %v operand%d \"%v\"", opcode.Mnemonic, operandID, rawOperand))
	}
}

func handleReadOperand(opcode opcode, operandID int) string {
	var rawOperand string
	if operandID == 1 {
		rawOperand = opcode.RawOperand1
	} else if operandID == 2 {
		rawOperand = opcode.RawOperand2
	}
	switch rawOperand {
	case "(a8)":
		return "s, err := readImmediate8BitAddress(data)"
	case "d8":
		return "s, err := readImmediate8BitData(data)"
	case "r8", "SP+r8":
		return "s, err := readImmediateSigned8BitData(data)"
	case "a16", "(a16)":
		return "s, err := readImmediate16BitAddress(data)"
	case "d16":
		return "s, err := readImmediate16BitData(data)"
	case "HL", "(HL)", "(HL-)", "(HL+)",
		"A", "AF", "B", "C", "(C)", "BC", "(BC)", "D", "E", "DE", "(DE)",
		"H", "L", "F",
		"SP", "PC",
		"Z",
		"NZ",
		"NC",
		"",
		"0", "1", "2", "3", "4", "5", "6", "7",
		"00H", "08H", "10H", "18H", "20H", "28H", "30H", "38H",
		"CB" /* CB prefixed opcodes, not the register C+B */ :
		return fmt.Sprintf("s := \"%v\" // %v is implicit in this instruction.", rawOperand, rawOperand)
	default:
		panic(fmt.Sprintf("You need to implement an writer for %v operand%d \"%v\"", opcode.Mnemonic, operandID, rawOperand))
	}
}

func handleWriteByteOperand(opcode opcode, operandID int) string {
	var rawOperand string
	if operandID == 1 {
		rawOperand = opcode.RawOperand1
	} else if operandID == 2 {
		rawOperand = opcode.RawOperand2
	}
	switch rawOperand {
	case "(a8)":
		return fmt.Sprintf("err = writeImmediate8BitAddress   (w, o.operand%d); written += 1", operandID)
	case "d8":
		return fmt.Sprintf("err = writeImmediate8BitData      (w, o.operand%d); written += 1", operandID)
	case "r8", "SP+r8":
		return fmt.Sprintf("err = writeImmediateSigned8BitData(w, o.operand%d); written += 1", operandID)
	case "a16", "(a16)":
		return fmt.Sprintf("err = writeImmediate16BitAddress  (w, o.operand%d); written += 2", operandID)
	case "d16":
		return fmt.Sprintf("err = writeImmediate16BitData     (w, o.operand%d); written += 2", operandID)
	case "HL", "(HL)", "(HL-)", "(HL+)",
		"A", "AF", "B", "C", "(C)", "BC", "(BC)", "D", "E", "DE", "(DE)",
		"H", "L", "F",
		"SP", "PC",
		"Z",
		"NZ",
		"NC",
		"",
		"0", "1", "2", "3", "4", "5", "6", "7",
		"00H", "08H", "10H", "18H", "20H", "28H", "30H", "38H",
		"CB" /* CB prefixed opcodes, not the register C+B */ :
		return fmt.Sprintf("// %v is implicit in this instruction.", rawOperand)
	default:
		panic(fmt.Sprintf("You need to implement an writer for %v operand%d \"%v\"", opcode.Mnemonic, operandID, rawOperand))
	}
}

func generateTemplate() *template.Template {
	b, err := ioutil.ReadFile("opcodes.tmpl")
	if err != nil {
		panic(err)
	}

	return template.Must(template.New("opcodes.tmpl").
		Funcs(template.FuncMap{
			"handleWriteByteOperand": handleWriteByteOperand,
			"handleStringOfOperand":  handleStringOfOperand,
			"handleReadOperand":      handleReadOperand,
		}).
		Parse(string(b)))
}

func addMnemonicHasOnlyOneCycle(o *opcode) error {
	if o.Mnemonic == "ADD" && len(o.Cycles) != 1 {
		return fmt.Errorf("ADD mnemoniced opcode with hex value %v has multiple cycle counts to return. All add instructions should be constant time", o.Addr)
	}

	return nil
}
