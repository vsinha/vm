package opcodes

import (
	"fmt"
	"io"
)

func readBytesAsString(reader io.Reader, n int) (string, error) {
	d := make([]byte, n)

	_, err := data.Read(d)
	if err != nil {
		return nil, err
	}

	s := fmt.Sprintf("%x", d)

	return s, nil
}
