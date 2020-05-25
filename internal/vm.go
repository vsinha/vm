package internal

import "fmt"

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

// opExec is a map that points from an opcode to the operational parameters of the opcode.
var opExec = map[op]struct {
	name string

	// f is the function which will be called upon encountering this opcode
	f func(*vm)

	// size describes how many bytes in memory the function takes as arguments
	size uint
}{
	nop:   {"nop", (*vm).nop, 1},     // nop           // nop
	add:   {"add", (*vm).add, 4},     // add r1 r2 r3  // r1 = r1 + r2
	addi:  {"addi", (*vm).addi, 4},   // add r1 r2 4   // r1 = r2 + 4
	halt:  {"halt", (*vm).halt, 1},   // halt          // halt
	loadi: {"loadi", (*vm).loadi, 3}, // loadi r1 5    // r1 = 5
	load:  {"load", (*vm).load, 3},   // loadi r1 r2   // r1 = r2
	sub:   {"sub", (*vm).sub, 4},     // sub r1 r2 r3  // r1 = r2 - r3
}

type registers [15]int
type memory []int

type vm struct {
	pc  uint
	r   registers
	mem memory
}

func (v *vm) run() error {
	o := op(v.mem[v.pc])
	for o != halt {
		ex := opExec[o]

		fmt.Printf("PC: %d Executing: %s\n", v.pc, ex.name)
		ex.f(v)

		v.pc += ex.size
		if v.pc == uint(len(v.mem)) {
			return fmt.Errorf("finished VM without halting")
		}
		o = op(v.mem[v.pc])
	}

	fmt.Printf("Found halt instruction. Done.\n")
	return nil
}

func (v *vm) nop() {}

func (v *vm) load() {
	reg1, reg2 := v.mem[v.pc+1], v.mem[v.pc+2]
	fmt.Printf("reg1: %v, reg2: %v\n", reg1, reg2)
	v.r[reg1] = v.r[reg2]
}

func (v *vm) loadi() {
	reg, constant := v.mem[v.pc+1], v.mem[v.pc+2]
	fmt.Printf("reg: %v, constant: %v\n", reg, constant)
	v.r[reg] = constant
}

func (v *vm) add() {
	dest, reg1, reg2 := v.mem[v.pc+1], v.mem[v.pc+2], v.mem[v.pc+3]
	fmt.Printf("dest: %v, r1: %v, r2: %v, r[%v]: %v, r[%v]: %v\n",
		dest, reg1, reg2, reg1, v.r[reg1], reg2, v.r[reg2])
	v.r[dest] = v.r[reg1] + v.r[reg2]
}

func (v *vm) addi() {
	dest, src, constant := v.mem[v.pc+1], v.mem[v.pc+2], v.mem[v.pc+3]
	fmt.Printf("dest: %v, src: %v, constant: %v\n", dest, src, constant)
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
