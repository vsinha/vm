package vm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
func Parse() Opcodes {
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
