package internal

import (
	"fmt"
	"strings"
)

type op int

const (
	nop op = iota
	add
	addi
	halt
	load
	loadi
	sub
)

type reg int

const (
	r0 reg = iota
	r1
	r2
	r3
	r4
	r5
	r6
	// TODO finish these
)

type opKind struct {
	size uint
}

var (
	unary  opKind = opKind{1} // NOP and HALT
	rType         = opKind{4} // works on registers
	riType        = opKind{4} // works on registers (with immediate values)
	lType         = opKind{3} // works on addresses
	liType        = opKind{3} // works on addresses (with immediate values)
)

// opExec is a map that points from an opcode to the operational parameters of the opcode.
var opExec = map[op]struct {
	name string

	// f is the function which will be called upon encountering this opcode
	f func(*vm)

	// kind describes how many bytes in memory the function takes as arguments
	kind opKind
}{
	nop:   {"nop", (*vm).nop, unary},      // nop           // nop
	add:   {"add", (*vm).add, rType},      // add r1 r2 r3  // r1 = r1 + r2
	addi:  {"addi", (*vm).addi, riType},   // add r1 r2 4   // r1 = r2 + 4
	halt:  {"halt", (*vm).halt, unary},    // halt          // halt
	loadi: {"loadi", (*vm).loadi, liType}, // loadi r1 5    // r1 = 5
	load:  {"load", (*vm).load, lType},    // loadi r1 r2   // r1 = r2
	sub:   {"sub", (*vm).sub, rType},      // sub r1 r2 r3  // r1 = r2 - r3
}

type registers [15]int

type memory []int

func (m memory) String() string {
	b := &strings.Builder{}

	nopCount := 0
	for i := 0; i < len(m); i++ {
		o := op(m[i])

		if o == nop {
			nopCount++
			continue
		} else if nopCount != 0 {
			fmt.Fprintf(b, "Found %d nops\n", nopCount)
		}

		if ex, ok := opExec[o]; ok {
			fmt.Fprintf(b, "%v ", ex.name)
			for j := 1; j < int(ex.kind.size); j++ {
				fmt.Fprintf(b, "%v ", m[i+j])
			}
			i += int(ex.kind.size)
			fmt.Fprintf(b, "\n")
		}
	}
	return b.String()
}

type vm struct {
	pc  uint
	r   registers
	mem memory

	trace bool
}

func (v *vm) run() error {
	o := op(v.mem[v.pc])
	for o != halt {
		ex := opExec[o]

		v.log("PC: %d Executing: %s\n", v.pc, ex.name)
		ex.f(v)

		v.pc += ex.kind.size
		if v.pc == uint(len(v.mem)) {
			return fmt.Errorf("finished VM without halting")
		}
		o = op(v.mem[v.pc])
	}

	v.log("Found halt instruction. Done.\n")
	return nil
}

func (v *vm) log(msg string, args ...interface{}) {
	if v.trace {
		fmt.Printf(msg, args...)
	}
}

func (v *vm) nop() {}

func (v *vm) load() {
	reg1, reg2 := v.mem[v.pc+1], v.mem[v.pc+2]
	v.log("reg1: %v, reg2: %v\n", reg1, reg2)
	v.r[reg1] = v.r[reg2]
}

func (v *vm) loadi() {
	reg, constant := v.mem[v.pc+1], v.mem[v.pc+2]
	v.log("reg: %v, constant: %v\n", reg, constant)
	v.r[reg] = constant
}

func (v *vm) add() {
	dest, reg1, reg2 := v.mem[v.pc+1], v.mem[v.pc+2], v.mem[v.pc+3]
	v.log("dest: %v, r1: %v, r2: %v, r[%v]: %v, r[%v]: %v\n",
		dest, reg1, reg2, reg1, v.r[reg1], reg2, v.r[reg2])
	v.r[dest] = v.r[reg1] + v.r[reg2]
}

func (v *vm) addi() {
	dest, src, constant := v.mem[v.pc+1], v.mem[v.pc+2], v.mem[v.pc+3]
	v.log("dest: %v, src: %v, constant: %v\n", dest, src, constant)
	v.r[dest] = v.r[src] + constant
}

func (v *vm) sub()  {}
func (v *vm) halt() {}

func assemble(instructions []interface{}) (memory, error) {
	empty := make(memory, 256)

	for i, inst := range instructions {
		switch v := inst.(type) {
		case op:
			empty[i] = int(v)
		case int:
			empty[i] = v
		case reg:
			empty[i] = int(v)
		default:
			return nil, fmt.Errorf("Unknown type %T", inst)
		}
	}

	return empty, nil
}

// Main runs a vm.
func Main() error {

	instructions := []interface{}{
		nop,
		halt,
	}

	mem, err := assemble(instructions)
	if err != nil {
		return fmt.Errorf("unable to assemble error: %v", err)
	}

	v := &vm{
		mem: mem,
	}

	if err := v.run(); err != nil {
		return fmt.Errorf("virtual machine error: %v", err)
	}

	return nil
}
