package memory

// Memory is the initialized starting memory to be passed into the VM
type Memory []byte

func (m Memory) ReadAt(p []byte, off int64) (n int, err error) {
	i := 0
	// Read up to len(p) bytes (as long as you don't overflow).
	for ; i < len(p) && int(off)+i < len(m); i++ {
		p[i] = m[off+int64(i)]
	}
	// If you overflow, return the remainder as 0x00.
	for ; i < len(p); i++ {
		p[i] = 0x00
	}
	return len(p), nil
}

/*
func (m Memory) String() string {
	b := &strings.Builder{}

	nopCount := 0
	for i := 0; i < len(m); i++ {
		o := Op(m[i])

		if o == Nop {
			nopCount++
			continue
		} else if nopCount != 0 {
			fmt.Fprintf(b, "Found %d nops\n", nopCount)
		}

		if ex, ok := OpExec[o]; ok {
			fmt.Fprintf(b, "%v ", ex.Name)
			for j := 1; j < int(ex.Kind.Size); j++ {
				fmt.Fprintf(b, "%v ", m[i+j])
			}
			i += int(ex.Kind.Size)
			fmt.Fprintf(b, "\n")
		}
	}
	return b.String()
}
*/
