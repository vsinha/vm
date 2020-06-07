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
		{[]byte{
			/* STOP */
			0x10}, `
STOP 0`},
		{[]byte{
			/* ADD A, C*/
			0x81}, `
ADD A,C`},
		{[]byte{
			/* ADD A, C*/
			0x81,
			/* LD B, d8 */
			0x06, 0xFF}, `
ADD A,C
LD B,ff`},
		{[]byte{
			/* ADD A, C*/
			0x81,
			/* LD B, d8 */
			0x06, 0xFF,
			/* CB-Prefix RLC B */
			0xCB, 0x00}, `
ADD A,C
LD B,ff
RLC B`},
		{[]byte{
			/* SUB B*/
			0xC0}, `
RET NZ`},
		{[]byte{
			/* LD B, d8*/
			0x06, 0xFF}, `
LD B,ff`},
	}

	for _, test := range tests {
		t.Run(test.want, func(t *testing.T) {
			var got strings.Builder
			r := bytes.NewBuffer(test.in)
			var w bytes.Buffer
			for {
				o, err := ReadOpCode(r)
				if err == io.EOF {
					break
				}
				if err != nil {
					t.Fatalf("ReadOpCode(%v) error: %v", test.in, err)
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
