package vm

import (
	"fmt"
	"strings"
)

// Op is an enum of the opcodes in our language
type Op int

// Instructions used in the language
const (
	Nop Op = iota
	Add
	Addi
	Halt
	J
	JR
	Load
	Loadi
	Sub
)

// Reg is an enum alias of the named registers
type Reg int

// Registers
const (
	R0 Reg = iota
	R1
	R2
	R3
	R4
	R5
	R6
	R7
	R8
	R9
	R10
	R11
	R12
	R13
	R14
)

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

type gbRegisters struct {
	A  uint8 // Accumulator % & flags
	B  uint8
	C  uint8
	D  uint8
	E  uint8
	F  FlagRegister
	H  uint8
	L  uint8
	SP uint8
	PC uint8
}

func (reg gbRegisters) GetBC() uint16 {
	return uint16(reg.B) << 8 & uint16(reg.C)
}

func (reg *gbRegisters) SetBC(val uint16) {
	reg.B = uint8(val >> 8 & 0xFF)
	reg.C = uint8(val & 0xFF)
}

// OpKind stores informationa about the kind of the opcode
type OpKind struct {
	Size uint
}

var (
	unary  OpKind = OpKind{1} // NOP and HALT
	rType         = OpKind{4} // works on registers
	riType        = OpKind{4} // works on registers (with immediate values)
	lType         = OpKind{3} // works on addresses
	liType        = OpKind{3} // works on addresses (with immediate values)
	jType         = OpKind{2} // works for control flow instructions (eg j)
)

// OpExec is a map that points from an opcode to the operational parameters of the opcode.
var OpExec = map[Op]struct {
	Name string

	// f is the function which will be called upon encountering this opcode
	f func(*VM)

	// kind describes how many bytes in memory the function takes as arguments
	Kind OpKind
}{
	Nop:   {"nop", (*VM).nop, unary},      // nop           // nop
	Add:   {"add", (*VM).add, rType},      // add r1 r2 r3  // r1 = r1 + r2
	Addi:  {"addi", (*VM).addi, riType},   // add r1 r2 4   // r1 = r2 + 4
	Halt:  {"halt", (*VM).halt, unary},    // halt          // halt
	Loadi: {"loadi", (*VM).loadi, liType}, // loadi r1 5    // r1 = 5
	Load:  {"load", (*VM).load, lType},    // loadi r1 r2   // r1 = r2
	Sub:   {"sub", (*VM).sub, rType},      // sub r1 r2 r3  // r1 = r2 - r3
	J:     {"j", (*VM).j, jType},          // j 5  			// jumps to address 5
	JR:    {"jr", (*VM).jr, jType},        // jr r1 		// jumps to the address stored in r1
}

// Registers are the registers of our VM
type Registers [15]int

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

func (v *VM) nop() {}

func (v *VM) load() {
	reg1, reg2 := v.mem[v.pc+1], v.mem[v.pc+2]
	v.log("reg1: %v, reg2: %v\n", reg1, reg2)
	v.r[reg1] = v.r[reg2]
}

func (v *VM) loadi() {
	reg, constant := v.mem[v.pc+1], v.mem[v.pc+2]
	v.log("reg: %v, constant: %v\n", reg, constant)
	v.r[reg] = constant
}

func (v *VM) add() {
	dest, reg1, reg2 := v.mem[v.pc+1], v.mem[v.pc+2], v.mem[v.pc+3]
	v.log("dest: %v, r1: %v, r2: %v, r[%v]: %v, r[%v]: %v\n",
		dest, reg1, reg2, reg1, v.r[reg1], reg2, v.r[reg2])
	v.r[dest] = v.r[reg1] + v.r[reg2]
}

func (v *VM) addi() {
	dest, src, constant := v.mem[v.pc+1], v.mem[v.pc+2], v.mem[v.pc+3]
	v.log("dest: %v, src: %v, constant: %v\n", dest, src, constant)
	v.r[dest] = v.r[src] + constant
}

func (v *VM) sub() {
	dest, reg1, reg2 := v.mem[v.pc+1], v.mem[v.pc+2], v.mem[v.pc+3]
	v.log("dest: %v, r1: %v, r2: %v, r[%v]: %v, r[%v]: %v\n",
		dest, reg1, reg2, reg1, v.r[reg1], reg2, v.r[reg2])
	v.r[dest] = v.r[reg1] - v.r[reg2]
}

func (v *VM) halt() {}

func (v *VM) j() {
	location := v.mem[v.pc+1]
	v.pc = uint(location)
}

func (v *VM) jr() {
	reg := v.mem[v.pc+1]
	v.log("r[%v]: %v", reg, v.r[reg])
	v.pc = uint(v.r[reg])
}

// Registers retursn the VM's registers.
func (v *VM) Registers() Registers {
	return v.r
}

// Memory returns the VM's memory
func (v *VM) Memory() Memory {
	return v.mem
}
