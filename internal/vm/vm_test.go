package vm_test

// import (
// 	"fmt"
// 	"strings"
// 	"testing"

// 	"github.com/google/go-cmp/cmp"

// 	"github.com/vsinha/vm/internal/assembler"
// 	"github.com/vsinha/vm/internal/vm"
// )

// func Example() {
// 	mem, err := assembler.Assemble([]interface{}{
// 		vm.Loadi, // r1 = 5
// 		vm.R1,
// 		5,
// 		vm.J,
// 		8,
// 		vm.Loadi, // r2 = 3
// 		vm.R2,
// 		1,
// 		vm.Halt,
// 	})
// 	if err != nil {
// 		fmt.Printf("error assembling: %v", err)
// 	}

// 	v := vm.New(mem)
// 	if err := v.Run(); err != nil {
// 		fmt.Printf("v.Run() error: %v", err)
// 	}

// 	fmt.Printf("Registers: %v", v.Registers())
// 	// Output: Registers: [0 5 0 0 0 0 0 0 0 0 0 0 0 0 0]
// }

// func TestNoZeroLength(t *testing.T) {
// 	for _, opex := range vm.OpExec {
// 		if opex.Kind.Size == 0 {
// 			t.Errorf("Instruction %q is length 0, this is forbidden", opex.Name)
// 		}
// 	}
// }

// func TestAssembleString(t *testing.T) {
// 	tests := []struct {
// 		code []interface{}
// 		want string
// 	}{
// 		{
// 			[]interface{}{
// 				vm.Addi, // r1 = r0 + 1
// 				vm.R1,
// 				vm.R0,
// 				1,
// 				vm.Halt,
// 			},
// 			"addi 1 0 1",
// 			//			`
// 			//addi r1 = r0 + 1 (2 1 0 1)
// 			//`,
// 		},
// 	}

// 	for _, test := range tests {
// 		got, err := assembler.Assemble(test.code)
// 		if err != nil {
// 			t.Fatalf("assemble(%v) error: %v", test.code, err)
// 		}

// 		if diff := cmp.Diff(test.want, strings.TrimSpace(got.String())); diff != "" {
// 			t.Errorf("assemble diff (-want,+got):\n%v", diff)
// 		}
// 	}
// }

// func TestProgram(t *testing.T) {
// 	tests := map[string]struct {
// 		code []interface{}
// 		reg  vm.Registers
// 	}{
// 		"addi": {
// 			[]interface{}{
// 				vm.Addi, // r1 = r0 + 1
// 				vm.R1,
// 				vm.R0,
// 				1,
// 				vm.Halt,
// 			},
// 			vm.Registers{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
// 		},
// 		"load": {
// 			[]interface{}{
// 				vm.Addi, // r1 = r0 + 1
// 				vm.R1,
// 				vm.R0,
// 				1,
// 				vm.Load, // r2 = r1
// 				vm.R2,
// 				vm.R1,
// 				vm.Halt,
// 			},
// 			vm.Registers{0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
// 		},
// 		"loadi": {
// 			[]interface{}{
// 				vm.Loadi, // r2 = 2
// 				vm.R2,
// 				2,
// 				vm.Halt,
// 			},
// 			vm.Registers{0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
// 		},
// 		"add": {
// 			[]interface{}{
// 				vm.Loadi, // r1 = 5
// 				vm.R1,
// 				5,
// 				vm.Loadi, // r2 = 3
// 				vm.R2,
// 				3,
// 				vm.Add, // r3 = r1 + r2
// 				vm.R3,
// 				vm.R1,
// 				vm.R2,
// 				vm.Halt,
// 			},
// 			vm.Registers{0, 5, 3, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
// 		},
// 		"sub": {
// 			[]interface{}{
// 				vm.Loadi, // r1 = 5
// 				vm.R1,
// 				5,
// 				vm.Loadi, // r2 = 3
// 				vm.R2,
// 				3,
// 				vm.Sub, // r3 = r1 - r2
// 				vm.R3,
// 				vm.R1,
// 				vm.R2,
// 				vm.Halt,
// 			},
// 			vm.Registers{0, 5, 3, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
// 		},
// 		"j": {
// 			[]interface{}{
// 				vm.Loadi, // r1 = 5
// 				vm.R1,
// 				5,
// 				vm.J,
// 				8,
// 				vm.Loadi, // r2 = 3
// 				vm.R2,
// 				1,
// 				vm.Halt,
// 			},
// 			vm.Registers{0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
// 		},
// 		"jr": {
// 			[]interface{}{
// 				vm.Loadi, // r1 = 5
// 				vm.R1,
// 				5,
// 				vm.Loadi, // r1 = 5
// 				vm.R2,
// 				11,
// 				vm.JR,
// 				vm.R2,
// 				vm.Loadi, // r2 = 3
// 				vm.R2,
// 				1,
// 				vm.Halt,
// 			},
// 			vm.Registers{0, 5, 11, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
// 		},
// 	}
// 	for name, test := range tests {
// 		t.Run(name, func(t *testing.T) {
// 			mem, err := assembler.Assemble(test.code)
// 			if err != nil {
// 				t.Errorf("assemble(%v) error: %v", test.code, err)
// 			}
// 			v := vm.New(mem)
// 			if err := v.Run(); err != nil {
// 				t.Errorf("v.Run() error: %v", err)
// 			}

// 			if diff := cmp.Diff(test.reg, v.Registers()); diff != "" {
// 				t.Errorf("vm registers diff (-test.reg,+got):\n%s", diff)
// 			}
// 		})
// 	}
// }
