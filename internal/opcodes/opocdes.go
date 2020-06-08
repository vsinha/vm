package opcodes

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/vsinha/vm/internal/memory"
	"github.com/vsinha/vm/internal/registers"
)

type Instruction interface {
	Execute(vm) error
	String() string
	Write(io.Writer) (int, error)
}

type vm interface {
	Mem() memory.Memory
	Reg() *registers.Registers
}

var endianness = binary.BigEndian

func readBytesAsString(r io.Reader, n int) (string, error) {
	d := make([]byte, n)

	readBytes, err := r.Read(d)
	if err != nil {
		return "", err
	}
	if readBytes != n {
		return "", fmt.Errorf("wanted to read %d bytes but only read %d", n, readBytes)
	}

	s := fmt.Sprintf("%x", d)

	return s, nil
}

func printLiteral(rawOperand string) bool {
	if rawOperand == "a8" ||
		rawOperand == "d8" ||
		rawOperand == "a16" ||
		rawOperand == "d16" {
		return true
	} else {
		return false
	}
}

// Execute DO NOT DO THIS
func (o *ADD_A_B) Execute(v vm) error {
	r := v.Reg()
	r.A = 5 + r.A // dont't hardcode 5
	r.PC += 1     // store instruction width in the opcode struct
	// set flags
	return nil
}

// Execute CB Prefixed badness
func (o *RES_0_B) Execute(v vm) error {
	return nil
}

// Execute Halt.
func (o *HALT) Execute(v vm) error {
	return fmt.Errorf("HALTED")
}

// Execute CB prefixed badness.
func (o *BIT_6_HLPtr) Execute(v vm) error {
	return nil
}
