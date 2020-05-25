package assembler

import (
	"fmt"

	"github.com/vsinha/vm/internal/vm"
)

// Assemble assembles
func Assemble(instructions []interface{}) (vm.Memory, error) {
	empty := make(vm.Memory, 256)

	for i, inst := range instructions {
		switch v := inst.(type) {
		case vm.Op:
			empty[i] = int(v)
		case int:
			empty[i] = v
		case vm.Reg:
			empty[i] = int(v)
		default:
			return nil, fmt.Errorf("Unknown type %T", inst)
		}
	}

	return empty, nil
}
