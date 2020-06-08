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
	//o := Op(v.mem[v.pc])

	v.r.PC = 0

	count := 1
	max := 100
	for {
		count++
		if count > max {
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
		if err := i.Execute(v); err != nil {
			return err
		}

		// increment cycles

		// if cycles > cycles_between_interrupts
		// then handle interrupts
	}

	/*
		for o != Halt {

			// while !timeForVblank
			// fetch (this depends on width)
			// Go to memory, read instruction, decode instruction (generated) returns a *opcodes.NOP

			// execute  instruction.Execute(vm)
			// exceute vm.Execute(instruction)
			// instruction.Execute(vm)
			// cycles += instrucntion.Cycles()
			//

			// increment something

			// vblank
			// sound things

			ex := OpExec[o]

			v.log("PC: %d Executing: %s\n", v.pc, ex.Name)
			ex.f(v)

			if ex.Kind != jType {
				// If we're jType, we assume the implementation handled moving the pc by itself
				v.pc += ex.Kind.Size

			}
			if v.pc == uint(len(v.mem)) {
				return fmt.Errorf("finished VM without halting")
			}
			o = Op(v.mem[v.pc])
		}
	*/

	v.log("Found halt instruction. Done.\n")
	return nil
}

func (v *VM) log(msg string, args ...interface{}) {
	if v.trace {
		fmt.Printf(msg, args...)
	}
}
