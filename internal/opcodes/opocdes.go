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
