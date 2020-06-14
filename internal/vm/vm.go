package vm

import (
	"fmt"
	"io"

	"github.com/vsinha/vm/internal/memory"
	"github.com/vsinha/vm/internal/opcodes"
	"github.com/vsinha/vm/internal/registers"
)

// VM is the in-memory virtual machine!
type VM struct {
	r   registers.Registers
	mem memory.Memory

	trace bool
}

// Reg returns the registers of the vm.
func (v *VM) Reg() *registers.Registers {
	return &v.r
}

// Mem returns the memory of the vm.
func (v *VM) Mem() memory.Memory {
	return v.mem
}

// New creates anew Virtual Machine with the provided memory initialized. The PC
// will be set to 0 and the VM will be ready to run.
func New(mem memory.Memory) *VM {
	return &VM{
		mem: mem,
	}
}

// Run executes the virtual machine.
func (v *VM) Run() error {

	v.r.PC = 0
	instructionCount := 1
	cycles := uint(0)
	max := 100
	for {
		instructionCount++
		if instructionCount > max {
			return fmt.Errorf("Ran %d instructions, bailing", max)
		}

		// Instructions are parseable in a direct way where the CPU can run an
		// opcode and move on to the next. However, there are times when you
		// will jump not to the beginning of an opcode, but instead to the
		// argument of an opcode.
		// For example:
		// 0x30, 0x01, 0x00, 0x00
		// This can be interpreted in two ways based on where you start parsing.
		// If you start parsing at byte 0, this is:
		// JR NC, 01
		// NOP
		// NOP
		// If you start parsing at byte 1, this is:
		// LD BC, 0000
		// There is no guaruntee that we are going to jump to the correct byte
		// alignment and sometimes this is used as a trick in obfuscated code.
		i, err := opcodes.ReadInstruction(io.NewSectionReader(v.mem, int64(v.r.PC), 4))
		if err != nil {
			return err
		}

		fmt.Printf("Running PC = %d: %v\n", v.r.PC, i)

		// execute
		executionResult, err := i.Execute(v)
		if err != nil {
			return err
		}

		// Increment PC by the length of the instruction.
		if !executionResult.DidSetPC {
			v.Reg().PC += uint16(i.Length())
		}

		// increment cycles
		cycles += uint(executionResult.Cycles)

		// if cycles > cycles_between_interrupts
		// then handle interrupts
	}
}

func (v *VM) log(msg string, args ...interface{}) {
	if v.trace {
		fmt.Printf(msg, args...)
	}
}
