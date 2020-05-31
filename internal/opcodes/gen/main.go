package main

//go:generate go run main.go -name=String -type=string > string.go

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/joncalhoun/pipe"
)

const jsonFile = "./opcodes.json"

// Opcode represents the data loaded from the accompanying JSON file
type Opcode struct {
	Mnemonic string   `json:"mnemonic"`
	Length   uint     `json:"length"`
	Cycles   uint     `json:"cycles"`
	Flags    []string `json:"flags"`
	Addr     string   `json:"addr"`
	Operand1 string   `json:"operand1"`
	Operand2 string   `json:"operand2"`
}

// Opcodes represents the data loaded from the accompanying JSON file
type Opcodes struct {
	Unprefixed map[string]Opcode `json:"unprefixed"`
	CBPrefixed map[string]Opcode `json:"cbprefixed"`
}

// Parse parses the opcodes.json file into nice struct
// to be used by the codgen tool
func parse() Opcodes {
	jsonFile, err := os.Open(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	fmt.Printf("Succesfully opened %v\n", jsonFile)

	bytes, _ := ioutil.ReadAll(jsonFile)

	var opcodes Opcodes
	json.Unmarshal(bytes, &opcodes)

	return opcodes
}

func generateFromTemplate(wr io.Writer, name string, templ string, opcodes Opcodes) {
	t := template.Must(template.New(name).Parse(templ))

	err := t.Execute(wr, opcodes)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Generated section: %s\n", name)
}

func main() {

	opcodes := parse()

	rc, wc, _ := pipe.Commands(
		exec.Command("gofmt"),
		exec.Command("goimports"),
	)

	generateFromTemplate(wc, "header", header, opcodes)
	generateFromTemplate(wc, "op", opTemplate, opcodes)
	generateFromTemplate(wc, "opExec", opExecTemplate, opcodes)

	wc.Close()

	outFile, err := os.Create("../opcodes_generated.go")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, rc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var header = `
package vm
`

// TODO this has problems because many of the opcodes have the same mnemonic,
// so we have to think about how to disambiguate them or if we even want to
var opTemplate = `
type Op int

const (
	{{range $key, $value := .Unprefixed }}
		{{.Mnemonic}} Op = {{ $key }}
	{{- end}}
)
`

var opExecTemplate = `
// put something here
`
