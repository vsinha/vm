package main

import (
	"fmt"
	"os"

	"github.com/vsinha/vm/internal/assembler"
	"github.com/vsinha/vm/internal/vm"
)

func main() {
	instructions := []interface{}{
		vm.Nop,
		vm.Halt,
	}

	mem, err := assembler.Assemble(instructions)
	if err != nil {
		fmt.Printf("unable to assemble error: %v", err)
		os.Exit(1)
	}

	v := vm.New(mem)
	if err := v.Run(); err != nil {
		fmt.Printf("virtual machine error: %v", err)
		os.Exit(1)
	}
}
