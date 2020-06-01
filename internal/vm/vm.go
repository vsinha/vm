package vm

import (
	"fmt"
	"strings"
)

// Reg is an enum alias of the named registers
type Reg uint8

// FlagRegister provides easier encapsulation for the Accumulator and Flag registers
type FlagRegister uint8

// FlagRegisterFlag does things
type FlagRegisterFlag uint8

// These are the accessors for the flags in the F flag register
const (
	FlagZf FlagRegisterFlag = 1 << 7 // Zero Flag
	FlagN                   = 1 << 6 // Add/Sub Flag (BCD)
	FlagH                   = 1 << 5 // Half Carry Flag
	FlagCy                  = 1 << 4 // Carry Flag
)

func hasBit(n uint8, pos uint8) bool {
	val := n & (1 << pos)
	return (val > 0)
}

func (af FlagRegister) isFlagSet(which FlagRegisterFlag) bool {
	return hasBit(uint8(af), uint8(which))
}

// Registers are the gameboy's registers.
type Registers struct {
	a  Reg // Accumulator % & flags
	b  Reg
	c  Reg
	d  Reg
	e  Reg
	f  FlagRegister
	h  Reg
	l  Reg
	sp uint16
	pc uint16
}

// BC returns the B register cat'd to the C register as an int16.
func (reg Registers) BC() uint16 {
	return uint16(reg.b) << 8 & uint16(reg.c)
}

// SetBC sets the B register and the C register individually as int8.
func (reg *Registers) SetBC(val uint16) {
	reg.b = Reg(val >> 8 & 0xFF)
	reg.c = Reg(val & 0xFF)
}

// Memory is the initialized starting memory to be passed into the VM
type Memory []int

func (m Memory) String() string {
	b := &strings.Builder{}

	nopCount := 0
	for i := 0; i < len(m); i++ {
		o := Op(m[i])

		if o == Nop {
			nopCount++
			continue
		} else if nopCount != 0 {
			fmt.Fprintf(b, "Found %d nops\n", nopCount)
		}

		if ex, ok := OpExec[o]; ok {
			fmt.Fprintf(b, "%v ", ex.Name)
			for j := 1; j < int(ex.Kind.Size); j++ {
				fmt.Fprintf(b, "%v ", m[i+j])
			}
			i += int(ex.Kind.Size)
			fmt.Fprintf(b, "\n")
		}
	}
	return b.String()
}

// VM is the in-memory virtual machine!
type VM struct {
	pc  uint
	r   Registers
	mem Memory

	trace bool
}

// New creates anew Virtual Machine with the provided memory initialized. The PC
// will be set to 0 and the VM will be ready to run.
func New(mem Memory) *VM {
	return &VM{
		mem: mem,
	}
}

// Run executes the virtual machine.
func (v *VM) Run() error {
	o := Op(v.mem[v.pc])

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

	v.log("Found halt instruction. Done.\n")
	return nil
}

func (v *VM) log(msg string, args ...interface{}) {
	if v.trace {
		fmt.Printf(msg, args...)
	}
}
