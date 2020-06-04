package opcodes

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	tests := []struct {
		in   []byte
		want string // Stringified version of the opcode
	}{
		{[]byte{0x81}, "ADD A,C"},
		{[]byte{0x06, 0x08}, "LD B,08"},
	}

	for _, test := range tests {
		r := bytes.NewBuffer(test.in)
		o, err := ReadOpCode(r)
		if err != nil {
			t.Fatalf("ReadOpCode(%v) error: %v", test.in, err)
		}
		if o.String() != test.want {
			t.Errorf("o.String() == %q != %q == test.want", o.String(), test.want)
		}
	}
}

func TestParseMultiple(t *testing.T) {
	tests := []struct {
		in   []byte
		want string // Stringified version of the opcode
	}{
		{[]byte{0x81, 0x06, 0x08}, `
ADD A,C
LD B,08`},
		{[]byte{0x81, 0x06, 0x08, 0x32, 0x16, 0x16}, `
ADD A,C
LD B,08`},
	}

	for _, test := range tests {
		var got strings.Builder
		r := bytes.NewBuffer(test.in)
		for {
			o, err := ReadOpCode(r)
			if err == io.EOF {
				break
			}
			if err != nil {
				t.Fatalf("ReadOpCode(%v) error: %v", test.in, err)
			}
			fmt.Fprintln(&got, o.String())
		}
		if diff := cmp.Diff(strings.TrimSpace(test.want), strings.TrimSpace(got.String())); diff != "" {
			t.Errorf("Diff (-want, + got):\n%s", diff)
		}
	}
}
