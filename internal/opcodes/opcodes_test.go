package opcodes_test

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"testing/quick"

	"github.com/google/go-cmp/cmp"
	"github.com/vsinha/vm/internal/memory"
	"github.com/vsinha/vm/internal/opcodes"
	"github.com/vsinha/vm/internal/registers"
	"github.com/vsinha/vm/internal/vm"
)

func TestParse(t *testing.T) {
	tests := []struct {
		in   []byte
		want string // Stringified version of the opcode
	}{
		{[]byte{
			0x10}, // STOP
			`STOP 0`},
		{[]byte{
			0x81}, // ADD A, C
			`ADD A C`},
		{[]byte{
			0xC0}, // SUB B
			`RET NZ`},
		{[]byte{
			0x81,        // ADD A, C
			0x06, 0xAA}, // LD B, d8
			`
ADD A C
LD B AA`},
		{[]byte{
			0x81,       // ADD A, C
			0x06, 0xAA, // LD B, d8
			0xCB, 0x00}, // CB-Prefix RLC B
			`
ADD A C
LD B AA
RLC B`},
		{[]byte{
			0xCB, 0x01, // RLC C
			0xE0, 0xAA}, // LDH 0xAA, A
			`
RLC C
LDH (AA) A`},
		{[]byte{
			0xF8, 0x12}, // LD HL SP+r8
			`LD HL SP+12`}, // SAD
		// Test that "negative" offsets for SP (look backward from the SP) are
		// represented by negative numbers in the output.
		{[]byte{
			0xF8, 0x81}, // LD HL SP+r8
			`LD HL SP+-7F`}, // SAD
	}

	for _, test := range tests {
		t.Run(test.want, func(t *testing.T) {
			var got strings.Builder
			r := bytes.NewBuffer(test.in)
			var w bytes.Buffer
			for {
				o, err := opcodes.ReadInstruction(r)
				if err == io.EOF {
					break
				}
				if err != nil {
					t.Fatalf("opcodes.ReadInstruction(%v) error: %v", test.in, err)
				}

				_, err = o.Write(&w)
				if err != nil {
					t.Fatalf("%q.Write() error: %v", o, err)
				}

				fmt.Fprintln(&got, o.String())
				t.Log(o.String())
			}
			if diff := cmp.Diff(strings.TrimSpace(test.want), strings.TrimSpace(got.String())); diff != "" {
				t.Errorf("Diff (-want, +got):\n%s", diff)
			}

			if diff := cmp.Diff(test.in, w.Bytes()); diff != "" {
				t.Errorf("Byte diff (-want, +got):\n%s", diff)
			}
		})
	}
}

type randomInstructions struct {
	Bytes []byte
}

// Generate implements quick.Generate.
func (_ randomInstructions) Generate(rand *rand.Rand, size int) reflect.Value {
	var instructions []opcodes.Instruction
	for j := 0; j < size; j++ {
		var b []byte

	generate:
		for {
			// Make a byte slice of 4 random bytes
			// 0 - 2^64
			b = []byte{
				byte(rand.Intn(1 << 8)),
				byte(rand.Intn(1 << 8)),
				byte(rand.Intn(1 << 8)),
				byte(rand.Intn(1 << 8)),
			}

			switch b[0] {
			case 0xD3, 0xDB, 0xDD, 0xE3, 0xE4, 0xEB, 0xEC, 0xED, 0xF4, 0xFC, 0xFD:
				// If the opcode is a hole in the z80 instruciton set, repeat
				// the generation process to see if we can get a valid instruction.
			default:
				// The opcode is not a hole in the z80 instruction set.
				break generate
			}
		}

		// Read 1 opcode off the random bytes
		i, err := opcodes.ReadInstruction(bytes.NewBuffer(b))
		if err != nil {
			panic(err)
		}

		// Append to the list of opcodes the generated opcode
		instructions = append(instructions, i)
	}

	// Write all the opcodes to a buffer and use them.
	var w bytes.Buffer
	for _, i := range instructions {
		_, err := i.Write(&w)
		if err != nil {
			panic(err)
		}
	}

	r := randomInstructions{
		Bytes: w.Bytes(),
	}
	return reflect.ValueOf(r)

}

// https://golang.org/pkg/testing/quick/
func TestRoundTrip(t *testing.T) {
	f := func(x randomInstructions) bool {
		t.Logf("x.Bytes = %x", x.Bytes)
		r := bytes.NewBuffer(x.Bytes)
		var w bytes.Buffer
		for {
			o, err := opcodes.ReadInstruction(r)
			if err == io.EOF {
				break
			}
			if err != nil {
				t.Fatalf("opcodes.ReadInstruction(%v) error: %v", x.Bytes, err)
				return false
			}

			_, err = o.Write(&w)
			if err != nil {
				t.Fatalf("%q.Write() error: %v", o, err)
				return false
			}
		}

		if diff := cmp.Diff(x.Bytes, w.Bytes()); diff != "" {
			t.Errorf("Byte diff (-want, +got):\n%s", diff)
			return false
		}
		return true
	}

	config := quick.Config{
		MaxCount:      10000,
		MaxCountScale: 0,
		Rand:          nil,
		Values:        nil,
	}

	if err := quick.Check(f, &config); err != nil {
		t.Error(err)
	}
}

func TestOpcodes(t *testing.T) {
	v := vm.New(memory.Memory{
		0xc6, // ADD A,d8
		0xff, // d8
		0x76, // Halt
	})

	err := v.Run()
	if err == opcodes.ErrHalt {
		t.Errorf("Got non HALTED error: %v", err)
	}

	want := &registers.Registers{
		A:  255,
		PC: 2,
	}
	if diff := cmp.Diff(v.Reg(), want); diff != "" {
		t.Errorf("Register difference (-got,+want): %v", diff)
	}
}
