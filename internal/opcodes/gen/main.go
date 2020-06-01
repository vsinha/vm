package main

//go:generate go run main.go -name=String -type=string

import (
	"encoding/json"
	"fmt"
	"io"
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
	Operand1         string   `json:"operand1"`
	Operand2         string   `json:"operand2"`
}

var templates = generateTemplate()

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
	default:
		return s
	}
}

// Parse parses the opcodes.json file into nice struct
// to be used by the codgen tool
func parse() []opcode {
	jsonFile, err := os.Open(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	fmt.Printf("Succesfully opened %s\n", jsonFile.Name())

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(fmt.Sprintf("ioutil.ReadAll(%q) error: %v", jsonFile.Name(), err))
	}

	var data struct {
		Unprefixed map[string]opcode `json:"unprefixed"`
		CBPrefixed map[string]opcode `json:"cbprefixed"`
	}
	if err := json.Unmarshal(bytes, &data); err != nil {
		panic(fmt.Sprintf("Unmarshal error: %v", err))
	}

	var opcodes []opcode
	for _, v := range data.Unprefixed {
		opcodes = append(opcodes, v)
	}
	for _, v := range data.CBPrefixed {
		v.CBPrefixed = true
		opcodes = append(opcodes, v)
	}

	for i, v := range opcodes {
		var args []string
		args = append(args, cleanMnemonic(v.Mnemonic))
		args = append(args, cleanMnemonic(v.Operand1))

		if v.Operand2 != "" {
			args = append(args, cleanMnemonic(v.Operand2))
		}

		extended := strings.Join(args, "_")

		opcodes[i].ExtendedMnemonic = extended
	}

	return opcodes
}

func main() {
	opcodes := parse()

	outFile, err := os.Create("../opcodes_generated.go")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer outFile.Close()

	// Synchronus reader/writer
	r, _ := io.Pipe()

	cmd := exec.Command("goimports")
	cmd.Stdin = r
	cmd.Stdout = outFile
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, t := range templates.Templates() {
		if err := t.Execute(outFile, opcodes); err != nil {
			log.Fatalf("Error rendering template %v: %v", t, err)
		}
	}
}

func generateTemplate() *template.Template {
	return template.Must(template.ParseGlob("*.tmpl"))
}
