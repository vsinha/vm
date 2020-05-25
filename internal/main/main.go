package main

import (
	"fmt"
	"io"

	"github.com/vsinha/vm/internal/assembler"
	"github.com/vsinha/vm/internal/vm"
)

// Main runs a vm.
func Main(r io.Reader) error {

	instructions := []interface{}{
		vm.Nop,
		vm.Halt,
	}

	mem, err := assembler.Assemble(instructions)
	if err != nil {
		return fmt.Errorf("unable to assemble error: %v", err)
	}

	v := vm.New(mem)
	if err := v.Run(); err != nil {
		return fmt.Errorf("virtual machine error: %v", err)
	}

	return nil
}
