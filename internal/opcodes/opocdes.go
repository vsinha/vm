package opcodes

import (
	"fmt"
	"io"
)

func readBytesAsString(r io.Reader, n int) (string, error) {
	d := make([]byte, n)

	_, err := r.Read(d)
	if err != nil {
		return "", err
	}

	s := fmt.Sprintf("%x", d)

	return s, nil
}

func printLiteral(rawOperand string) bool {
	if rawOperand == "a8" ||
		rawOperand == "d8" ||
		rawOperand == "a16" ||
		rawOperand == "d16" {
		return true
	} else {
		return false
	}
}
