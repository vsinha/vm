package registers

// Reg is an enum alias of the named registers
type Reg uint8

// Registers are the gameboy's registers.
type Registers struct {
	A  Reg // Accumulator % & flags
	B  Reg
	C  Reg
	D  Reg
	E  Reg
	F  FlagRegister
	H  Reg
	L  Reg
	SP uint16
	PC uint16
}

// BC returns the B register cat'd to the C register as an int16.
func (reg Registers) BC() uint16 {
	return uint16(reg.B) << 8 & uint16(reg.C)
}

// SetBC sets the B register and the C register individually as int8.
func (reg *Registers) SetBC(val uint16) {
	reg.B = Reg(val >> 8 & 0xFF)
	reg.C = Reg(val & 0xFF)
}

// FlagRegister provides easier encapsulation for the Accumulator and Flag registers
type FlagRegister uint8

// FlagRegisterFlag does things
type FlagRegisterFlag uint8

// These are the accessors for the flags in the F flag register
const (
	FlagZf FlagRegisterFlag = 1 << 7 // Zero Flag
	FlagN                   = 1 << 6 // Add/Sub Flag (BCD)
	FlagH                   = 1 << 5 // Half Carry Flag
	FlagCy                  = 1 << 4 // Carry Flag
)

func hasBit(n uint8, pos uint8) bool {
	val := n & (1 << pos)
	return (val > 0)
}

func (af FlagRegister) isFlagSet(which FlagRegisterFlag) bool {
	return hasBit(uint8(af), uint8(which))
}
