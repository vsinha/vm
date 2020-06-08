package opcodes

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
			`ADD A, C`},
		{[]byte{
			0xC0}, // SUB B
			`RET NZ`},
		{[]byte{
			0x81,        // ADD A, C
			0x06, 0xFF}, // LD B, d8
			`
ADD A, C
LD B, ff`},
		{[]byte{
			0x81,       // ADD A, C
			0x06, 0xFF, // LD B, d8
			0xCB, 0x00}, // CB-Prefix RLC B
			`
ADD A, C
LD B, ff
RLC B`},
		{[]byte{
			0xCB, 0x01, // RLC C
			0xE0, 0xFF}, // LDH 0xFF, A
			`
RLC C
LDH ff, A`},
		{[]byte{
			0xF8, 0xFF}, // LD HL SP+r8
			`LD HL, SP+ff`}, // SAD
	}

	for _, test := range tests {
		t.Run(test.want, func(t *testing.T) {
			var got strings.Builder
			r := bytes.NewBuffer(test.in)
			var w bytes.Buffer
			for {
				o, err := ReadInstruction(r)
				if err == io.EOF {
					break
				}
				if err != nil {
					t.Fatalf("ReadInstruction(%v) error: %v", test.in, err)
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

	var instructions []Instruction
	for j := 0; j < size; j++ {
		// Make a byte slice of 4 random bytes
		// 0 - 2^64
		b := []byte{
			0x11,
			byte(rand.Intn(1 << 8)),
		}

		// Read 1 opcode off the random bytes
		i, err := ReadInstruction(bytes.NewBuffer(b))
		if err != nil {
			panic(err)
		}
		panic(fmt.Sprintf("%v: %v", i.String(), b))

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
	t.Skipf("Reenable once we have adding working")
	f := func(x randomInstructions) bool {
		t.Logf("x.Bytes = %x", x.Bytes)
		r := bytes.NewBuffer(x.Bytes)
		var w bytes.Buffer
		for {
			o, err := ReadInstruction(r)
			if err == io.EOF {
				break
			}
			if err != nil {
				t.Fatalf("ReadInstruction(%v) error: %v", x.Bytes, err)
				return false
			}

			_, err = o.Write(&w)
			if err != nil {
				t.Fatalf("%q.Write() error: %v", o, err)
				return false
			}

			t.Logf("Instruction: %s", r)
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
