package internal

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInit(t *testing.T) {
	err := Main()
	if err != nil {
		t.Errorf("doSomething() error: %v", err)
	}
}

func TestNoZeroLength(t *testing.T) {
	for _, opex := range opExec {
		if opex.kind.size == 0 {
			t.Errorf("Instruction %q is length 0, this is forbidden", opex.name)
		}
	}
}

func TestAssembleString(t *testing.T) {
	tests := []struct {
		code []interface{}
		want string
	}{
		{
			[]interface{}{
				addi, // r1 = r0 + 1
				r1,
				r0,
				1,
				halt,
			},
			"addi 1 0 1",
			//			`
			//addi r1 = r0 + 1 (2 1 0 1)
			//`,
		},
	}

	for _, test := range tests {
		got, err := assemble(test.code)
		if err != nil {
			t.Fatalf("assemble(%v) error: %v", test.code, err)
		}

		if diff := cmp.Diff(test.want, strings.TrimSpace(got.String())); diff != "" {
			t.Errorf("assemble diff (-want,+got):\n%v", diff)
		}
	}
}

func TestProgram(t *testing.T) {
	tests := map[string]struct {
		code []interface{}
		reg  registers
	}{
		"addi": {
			[]interface{}{
				addi, // r1 = r0 + 1
				r1,
				r0,
				1,
				halt,
			},
			registers{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		"load": {
			[]interface{}{
				addi, // r1 = r0 + 1
				r1,
				r0,
				1,
				load, // r2 = r1
				r2,
				r1,
				halt,
			},
			registers{0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		"loadi": {
			[]interface{}{
				loadi, // r2 = 2
				r2,
				2,
				halt,
			},
			registers{0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		"add": {
			[]interface{}{
				loadi, // r1 = 5
				r1,
				5,
				loadi, // r2 = 3
				r2,
				3,
				add, // r3 = r1 + r2
				r3,
				r1,
				r2,
				halt,
			},
			registers{0, 5, 3, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			mem, err := assemble(test.code)
			if err != nil {
				t.Errorf("assemble(%v) error: %v", test.code, err)
			}
			vm := &vm{
				mem: mem,
			}
			if err := vm.run(); err != nil {
				t.Errorf("vm.run() error: %v", err)
			}

			if diff := cmp.Diff(test.reg, vm.r); diff != "" {
				t.Errorf("vm registers diff (-test.reg,+got):\n%s", diff)
			}
		})
	}
}
