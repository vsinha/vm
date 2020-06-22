package opcodes

import (
	"errors"
	"fmt"
	"io"
)

// ErrNoOpCode is a little alias for the error message returned below
var ErrNoOpCode = errors.New("no opcode with that address exists")

// TODO this has problems because many of the opcodes have the same mnemonic,
// so we have to think about how to disambiguate them or if we even want to
type Op int

type NOP struct {
	addr         string      // 0x0
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *NOP) Write(w io.Writer) (int, error) { // 0x0
	var b []byte

	b = append(b, 0x0)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x0,
	})
	if err != nil {
		return written, err
	}

	//  is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *NOP) Length() uint8 { // 0x0
	return 1
}

func (o *NOP) cycles() []uint8 { // 0x0
	return []uint8{4}
}

func (o *NOP) String() string { // 0x0
	return "NOP"
}

func (o *NOP) SymbolicString() string { // 0x0
	return "NOP"
}

type LD_BC_d16 struct {
	addr         string      // 0x1
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_BC_d16) Write(w io.Writer) (int, error) { // 0x1
	var b []byte

	b = append(b, 0x1)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x1,
	})
	if err != nil {
		return written, err
	}

	// BC is implicit in this instruction.
	err = writeImmediate16BitData(w, o.operand2)
	written += 2

	return written, err
}

func (o *LD_BC_d16) Length() uint8 { // 0x1
	return 3
}

func (o *LD_BC_d16) cycles() []uint8 { // 0x1
	return []uint8{12}
}

func (o *LD_BC_d16) String() string { // 0x1
	return "LD" + " " + "BC" /* operand1 */ + " " + fmt.Sprintf("%X", o.operand2)
}

func (o *LD_BC_d16) SymbolicString() string { // 0x1
	return "LD BC,d16"
}

type STOP_0 struct {
	addr         string      // 0x10
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *STOP_0) Write(w io.Writer) (int, error) { // 0x10
	var b []byte

	b = append(b, 0x10)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x10,
	})
	if err != nil {
		return written, err
	}

	// 0 is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *STOP_0) Length() uint8 { // 0x10
	return 1
}

func (o *STOP_0) cycles() []uint8 { // 0x10
	return []uint8{4}
}

func (o *STOP_0) String() string { // 0x10
	return "STOP" + " " + "0" /* operand1 */
}

func (o *STOP_0) SymbolicString() string { // 0x10
	return "STOP 0"
}

type LD_DE_d16 struct {
	addr         string      // 0x11
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_DE_d16) Write(w io.Writer) (int, error) { // 0x11
	var b []byte

	b = append(b, 0x11)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x11,
	})
	if err != nil {
		return written, err
	}

	// DE is implicit in this instruction.
	err = writeImmediate16BitData(w, o.operand2)
	written += 2

	return written, err
}

func (o *LD_DE_d16) Length() uint8 { // 0x11
	return 3
}

func (o *LD_DE_d16) cycles() []uint8 { // 0x11
	return []uint8{12}
}

func (o *LD_DE_d16) String() string { // 0x11
	return "LD" + " " + "DE" /* operand1 */ + " " + fmt.Sprintf("%X", o.operand2)
}

func (o *LD_DE_d16) SymbolicString() string { // 0x11
	return "LD DE,d16"
}

type LD_DEDeref_A struct {
	addr         string      // 0x12
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_DEDeref_A) Write(w io.Writer) (int, error) { // 0x12
	var b []byte

	b = append(b, 0x12)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x12,
	})
	if err != nil {
		return written, err
	}

	// (DE) is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *LD_DEDeref_A) Length() uint8 { // 0x12
	return 1
}

func (o *LD_DEDeref_A) cycles() []uint8 { // 0x12
	return []uint8{8}
}

func (o *LD_DEDeref_A) String() string { // 0x12
	return "LD" + " " + "(DE)" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *LD_DEDeref_A) SymbolicString() string { // 0x12
	return "LD (DE),A"
}

type INC_DE struct {
	addr         string      // 0x13
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *INC_DE) Write(w io.Writer) (int, error) { // 0x13
	var b []byte

	b = append(b, 0x13)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x13,
	})
	if err != nil {
		return written, err
	}

	// DE is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *INC_DE) Length() uint8 { // 0x13
	return 1
}

func (o *INC_DE) cycles() []uint8 { // 0x13
	return []uint8{8}
}

func (o *INC_DE) String() string { // 0x13
	return "INC" + " " + "DE" /* operand1 */
}

func (o *INC_DE) SymbolicString() string { // 0x13
	return "INC DE"
}

type INC_D struct {
	addr         string      // 0x14
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *INC_D) Write(w io.Writer) (int, error) { // 0x14
	var b []byte

	b = append(b, 0x14)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x14,
	})
	if err != nil {
		return written, err
	}

	// D is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *INC_D) Length() uint8 { // 0x14
	return 1
}

func (o *INC_D) cycles() []uint8 { // 0x14
	return []uint8{4}
}

func (o *INC_D) String() string { // 0x14
	return "INC" + " " + "D" /* operand1 */
}

func (o *INC_D) SymbolicString() string { // 0x14
	return "INC D"
}

type DEC_D struct {
	addr         string      // 0x15
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *DEC_D) Write(w io.Writer) (int, error) { // 0x15
	var b []byte

	b = append(b, 0x15)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x15,
	})
	if err != nil {
		return written, err
	}

	// D is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *DEC_D) Length() uint8 { // 0x15
	return 1
}

func (o *DEC_D) cycles() []uint8 { // 0x15
	return []uint8{4}
}

func (o *DEC_D) String() string { // 0x15
	return "DEC" + " " + "D" /* operand1 */
}

func (o *DEC_D) SymbolicString() string { // 0x15
	return "DEC D"
}

type LD_D_d8 struct {
	addr         string      // 0x16
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_D_d8) Write(w io.Writer) (int, error) { // 0x16
	var b []byte

	b = append(b, 0x16)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x16,
	})
	if err != nil {
		return written, err
	}

	// D is implicit in this instruction.
	err = writeImmediate8BitData(w, o.operand2)
	written += 1

	return written, err
}

func (o *LD_D_d8) Length() uint8 { // 0x16
	return 2
}

func (o *LD_D_d8) cycles() []uint8 { // 0x16
	return []uint8{8}
}

func (o *LD_D_d8) String() string { // 0x16
	return "LD" + " " + "D" /* operand1 */ + " " + fmt.Sprintf("%X", o.operand2)
}

func (o *LD_D_d8) SymbolicString() string { // 0x16
	return "LD D,d8"
}

type RLA struct {
	addr         string      // 0x17
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RLA) Write(w io.Writer) (int, error) { // 0x17
	var b []byte

	b = append(b, 0x17)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x17,
	})
	if err != nil {
		return written, err
	}

	//  is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RLA) Length() uint8 { // 0x17
	return 1
}

func (o *RLA) cycles() []uint8 { // 0x17
	return []uint8{4}
}

func (o *RLA) String() string { // 0x17
	return "RLA"
}

func (o *RLA) SymbolicString() string { // 0x17
	return "RLA"
}

type JR_r8 struct {
	addr         string      // 0x18
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *JR_r8) Write(w io.Writer) (int, error) { // 0x18
	var b []byte

	b = append(b, 0x18)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x18,
	})
	if err != nil {
		return written, err
	}

	err = writeImmediateSigned8BitData(w, o.operand1)
	written += 1
	//  is implicit in this instruction.

	return written, err
}

func (o *JR_r8) Length() uint8 { // 0x18
	return 2
}

func (o *JR_r8) cycles() []uint8 { // 0x18
	return []uint8{12}
}

func (o *JR_r8) String() string { // 0x18
	return "JR" + " " + fmt.Sprintf("SP+%X", o.operand1)
}

func (o *JR_r8) SymbolicString() string { // 0x18
	return "JR r8"
}

type ADD_HL_DE struct {
	addr         string      // 0x19
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *ADD_HL_DE) Write(w io.Writer) (int, error) { // 0x19
	var b []byte

	b = append(b, 0x19)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x19,
	})
	if err != nil {
		return written, err
	}

	// HL is implicit in this instruction.
	// DE is implicit in this instruction.

	return written, err
}

func (o *ADD_HL_DE) Length() uint8 { // 0x19
	return 1
}

func (o *ADD_HL_DE) cycles() []uint8 { // 0x19
	return []uint8{8}
}

func (o *ADD_HL_DE) String() string { // 0x19
	return "ADD" + " " + "HL" /* operand1 */ + " " + "DE" /* operand2 */
}

func (o *ADD_HL_DE) SymbolicString() string { // 0x19
	return "ADD HL,DE"
}

type LD_A_DEDeref struct {
	addr         string      // 0x1a
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_A_DEDeref) Write(w io.Writer) (int, error) { // 0x1a
	var b []byte

	b = append(b, 0x1a)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x1a,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// (DE) is implicit in this instruction.

	return written, err
}

func (o *LD_A_DEDeref) Length() uint8 { // 0x1a
	return 1
}

func (o *LD_A_DEDeref) cycles() []uint8 { // 0x1a
	return []uint8{8}
}

func (o *LD_A_DEDeref) String() string { // 0x1a
	return "LD" + " " + "A" /* operand1 */ + " " + "(DE)" /* operand2 */
}

func (o *LD_A_DEDeref) SymbolicString() string { // 0x1a
	return "LD A,(DE)"
}

type DEC_DE struct {
	addr         string      // 0x1b
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *DEC_DE) Write(w io.Writer) (int, error) { // 0x1b
	var b []byte

	b = append(b, 0x1b)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x1b,
	})
	if err != nil {
		return written, err
	}

	// DE is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *DEC_DE) Length() uint8 { // 0x1b
	return 1
}

func (o *DEC_DE) cycles() []uint8 { // 0x1b
	return []uint8{8}
}

func (o *DEC_DE) String() string { // 0x1b
	return "DEC" + " " + "DE" /* operand1 */
}

func (o *DEC_DE) SymbolicString() string { // 0x1b
	return "DEC DE"
}

type INC_E struct {
	addr         string      // 0x1c
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *INC_E) Write(w io.Writer) (int, error) { // 0x1c
	var b []byte

	b = append(b, 0x1c)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x1c,
	})
	if err != nil {
		return written, err
	}

	// E is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *INC_E) Length() uint8 { // 0x1c
	return 1
}

func (o *INC_E) cycles() []uint8 { // 0x1c
	return []uint8{4}
}

func (o *INC_E) String() string { // 0x1c
	return "INC" + " " + "E" /* operand1 */
}

func (o *INC_E) SymbolicString() string { // 0x1c
	return "INC E"
}

type DEC_E struct {
	addr         string      // 0x1d
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *DEC_E) Write(w io.Writer) (int, error) { // 0x1d
	var b []byte

	b = append(b, 0x1d)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x1d,
	})
	if err != nil {
		return written, err
	}

	// E is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *DEC_E) Length() uint8 { // 0x1d
	return 1
}

func (o *DEC_E) cycles() []uint8 { // 0x1d
	return []uint8{4}
}

func (o *DEC_E) String() string { // 0x1d
	return "DEC" + " " + "E" /* operand1 */
}

func (o *DEC_E) SymbolicString() string { // 0x1d
	return "DEC E"
}

type LD_E_d8 struct {
	addr         string      // 0x1e
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_E_d8) Write(w io.Writer) (int, error) { // 0x1e
	var b []byte

	b = append(b, 0x1e)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x1e,
	})
	if err != nil {
		return written, err
	}

	// E is implicit in this instruction.
	err = writeImmediate8BitData(w, o.operand2)
	written += 1

	return written, err
}

func (o *LD_E_d8) Length() uint8 { // 0x1e
	return 2
}

func (o *LD_E_d8) cycles() []uint8 { // 0x1e
	return []uint8{8}
}

func (o *LD_E_d8) String() string { // 0x1e
	return "LD" + " " + "E" /* operand1 */ + " " + fmt.Sprintf("%X", o.operand2)
}

func (o *LD_E_d8) SymbolicString() string { // 0x1e
	return "LD E,d8"
}

type RRA struct {
	addr         string      // 0x1f
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RRA) Write(w io.Writer) (int, error) { // 0x1f
	var b []byte

	b = append(b, 0x1f)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x1f,
	})
	if err != nil {
		return written, err
	}

	//  is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RRA) Length() uint8 { // 0x1f
	return 1
}

func (o *RRA) cycles() []uint8 { // 0x1f
	return []uint8{4}
}

func (o *RRA) String() string { // 0x1f
	return "RRA"
}

func (o *RRA) SymbolicString() string { // 0x1f
	return "RRA"
}

type LD_BCDeref_A struct {
	addr         string      // 0x2
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_BCDeref_A) Write(w io.Writer) (int, error) { // 0x2
	var b []byte

	b = append(b, 0x2)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x2,
	})
	if err != nil {
		return written, err
	}

	// (BC) is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *LD_BCDeref_A) Length() uint8 { // 0x2
	return 1
}

func (o *LD_BCDeref_A) cycles() []uint8 { // 0x2
	return []uint8{8}
}

func (o *LD_BCDeref_A) String() string { // 0x2
	return "LD" + " " + "(BC)" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *LD_BCDeref_A) SymbolicString() string { // 0x2
	return "LD (BC),A"
}

type JR_NZ_r8 struct {
	addr         string      // 0x20
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *JR_NZ_r8) Write(w io.Writer) (int, error) { // 0x20
	var b []byte

	b = append(b, 0x20)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x20,
	})
	if err != nil {
		return written, err
	}

	// NZ is implicit in this instruction.
	err = writeImmediateSigned8BitData(w, o.operand2)
	written += 1

	return written, err
}

func (o *JR_NZ_r8) Length() uint8 { // 0x20
	return 2
}

func (o *JR_NZ_r8) cycles() []uint8 { // 0x20
	return []uint8{12, 8}
}

func (o *JR_NZ_r8) String() string { // 0x20
	return "JR" + " " + "NZ" /* operand1 */ + " " + fmt.Sprintf("SP+%X", o.operand2)
}

func (o *JR_NZ_r8) SymbolicString() string { // 0x20
	return "JR NZ,r8"
}

type LD_HL_d16 struct {
	addr         string      // 0x21
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_HL_d16) Write(w io.Writer) (int, error) { // 0x21
	var b []byte

	b = append(b, 0x21)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x21,
	})
	if err != nil {
		return written, err
	}

	// HL is implicit in this instruction.
	err = writeImmediate16BitData(w, o.operand2)
	written += 2

	return written, err
}

func (o *LD_HL_d16) Length() uint8 { // 0x21
	return 3
}

func (o *LD_HL_d16) cycles() []uint8 { // 0x21
	return []uint8{12}
}

func (o *LD_HL_d16) String() string { // 0x21
	return "LD" + " " + "HL" /* operand1 */ + " " + fmt.Sprintf("%X", o.operand2)
}

func (o *LD_HL_d16) SymbolicString() string { // 0x21
	return "LD HL,d16"
}

type LD_HLPtrInc_A struct {
	addr         string      // 0x22
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_HLPtrInc_A) Write(w io.Writer) (int, error) { // 0x22
	var b []byte

	b = append(b, 0x22)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x22,
	})
	if err != nil {
		return written, err
	}

	// (HL+) is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *LD_HLPtrInc_A) Length() uint8 { // 0x22
	return 1
}

func (o *LD_HLPtrInc_A) cycles() []uint8 { // 0x22
	return []uint8{8}
}

func (o *LD_HLPtrInc_A) String() string { // 0x22
	return "LD" + " " + "(HL+)" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *LD_HLPtrInc_A) SymbolicString() string { // 0x22
	return "LD (HL+),A"
}

type INC_HL struct {
	addr         string      // 0x23
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *INC_HL) Write(w io.Writer) (int, error) { // 0x23
	var b []byte

	b = append(b, 0x23)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x23,
	})
	if err != nil {
		return written, err
	}

	// HL is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *INC_HL) Length() uint8 { // 0x23
	return 1
}

func (o *INC_HL) cycles() []uint8 { // 0x23
	return []uint8{8}
}

func (o *INC_HL) String() string { // 0x23
	return "INC" + " " + "HL" /* operand1 */
}

func (o *INC_HL) SymbolicString() string { // 0x23
	return "INC HL"
}

type INC_H struct {
	addr         string      // 0x24
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *INC_H) Write(w io.Writer) (int, error) { // 0x24
	var b []byte

	b = append(b, 0x24)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x24,
	})
	if err != nil {
		return written, err
	}

	// H is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *INC_H) Length() uint8 { // 0x24
	return 1
}

func (o *INC_H) cycles() []uint8 { // 0x24
	return []uint8{4}
}

func (o *INC_H) String() string { // 0x24
	return "INC" + " " + "H" /* operand1 */
}

func (o *INC_H) SymbolicString() string { // 0x24
	return "INC H"
}

type DEC_H struct {
	addr         string      // 0x25
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *DEC_H) Write(w io.Writer) (int, error) { // 0x25
	var b []byte

	b = append(b, 0x25)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x25,
	})
	if err != nil {
		return written, err
	}

	// H is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *DEC_H) Length() uint8 { // 0x25
	return 1
}

func (o *DEC_H) cycles() []uint8 { // 0x25
	return []uint8{4}
}

func (o *DEC_H) String() string { // 0x25
	return "DEC" + " " + "H" /* operand1 */
}

func (o *DEC_H) SymbolicString() string { // 0x25
	return "DEC H"
}

type LD_H_d8 struct {
	addr         string      // 0x26
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_H_d8) Write(w io.Writer) (int, error) { // 0x26
	var b []byte

	b = append(b, 0x26)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x26,
	})
	if err != nil {
		return written, err
	}

	// H is implicit in this instruction.
	err = writeImmediate8BitData(w, o.operand2)
	written += 1

	return written, err
}

func (o *LD_H_d8) Length() uint8 { // 0x26
	return 2
}

func (o *LD_H_d8) cycles() []uint8 { // 0x26
	return []uint8{8}
}

func (o *LD_H_d8) String() string { // 0x26
	return "LD" + " " + "H" /* operand1 */ + " " + fmt.Sprintf("%X", o.operand2)
}

func (o *LD_H_d8) SymbolicString() string { // 0x26
	return "LD H,d8"
}

type DAA struct {
	addr         string      // 0x27
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *DAA) Write(w io.Writer) (int, error) { // 0x27
	var b []byte

	b = append(b, 0x27)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x27,
	})
	if err != nil {
		return written, err
	}

	//  is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *DAA) Length() uint8 { // 0x27
	return 1
}

func (o *DAA) cycles() []uint8 { // 0x27
	return []uint8{4}
}

func (o *DAA) String() string { // 0x27
	return "DAA"
}

func (o *DAA) SymbolicString() string { // 0x27
	return "DAA"
}

type JR_Z_r8 struct {
	addr         string      // 0x28
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *JR_Z_r8) Write(w io.Writer) (int, error) { // 0x28
	var b []byte

	b = append(b, 0x28)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x28,
	})
	if err != nil {
		return written, err
	}

	// Z is implicit in this instruction.
	err = writeImmediateSigned8BitData(w, o.operand2)
	written += 1

	return written, err
}

func (o *JR_Z_r8) Length() uint8 { // 0x28
	return 2
}

func (o *JR_Z_r8) cycles() []uint8 { // 0x28
	return []uint8{12, 8}
}

func (o *JR_Z_r8) String() string { // 0x28
	return "JR" + " " + "Z" /* operand1 */ + " " + fmt.Sprintf("SP+%X", o.operand2)
}

func (o *JR_Z_r8) SymbolicString() string { // 0x28
	return "JR Z,r8"
}

type ADD_HL_HL struct {
	addr         string      // 0x29
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *ADD_HL_HL) Write(w io.Writer) (int, error) { // 0x29
	var b []byte

	b = append(b, 0x29)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x29,
	})
	if err != nil {
		return written, err
	}

	// HL is implicit in this instruction.
	// HL is implicit in this instruction.

	return written, err
}

func (o *ADD_HL_HL) Length() uint8 { // 0x29
	return 1
}

func (o *ADD_HL_HL) cycles() []uint8 { // 0x29
	return []uint8{8}
}

func (o *ADD_HL_HL) String() string { // 0x29
	return "ADD" + " " + "HL" /* operand1 */ + " " + "HL" /* operand2 */
}

func (o *ADD_HL_HL) SymbolicString() string { // 0x29
	return "ADD HL,HL"
}

type LD_A_HLPtrInc struct {
	addr         string      // 0x2a
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_A_HLPtrInc) Write(w io.Writer) (int, error) { // 0x2a
	var b []byte

	b = append(b, 0x2a)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x2a,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// (HL+) is implicit in this instruction.

	return written, err
}

func (o *LD_A_HLPtrInc) Length() uint8 { // 0x2a
	return 1
}

func (o *LD_A_HLPtrInc) cycles() []uint8 { // 0x2a
	return []uint8{8}
}

func (o *LD_A_HLPtrInc) String() string { // 0x2a
	return "LD" + " " + "A" /* operand1 */ + " " + "(HL+)" /* operand2 */
}

func (o *LD_A_HLPtrInc) SymbolicString() string { // 0x2a
	return "LD A,(HL+)"
}

type DEC_HL struct {
	addr         string      // 0x2b
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *DEC_HL) Write(w io.Writer) (int, error) { // 0x2b
	var b []byte

	b = append(b, 0x2b)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x2b,
	})
	if err != nil {
		return written, err
	}

	// HL is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *DEC_HL) Length() uint8 { // 0x2b
	return 1
}

func (o *DEC_HL) cycles() []uint8 { // 0x2b
	return []uint8{8}
}

func (o *DEC_HL) String() string { // 0x2b
	return "DEC" + " " + "HL" /* operand1 */
}

func (o *DEC_HL) SymbolicString() string { // 0x2b
	return "DEC HL"
}

type INC_L struct {
	addr         string      // 0x2c
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *INC_L) Write(w io.Writer) (int, error) { // 0x2c
	var b []byte

	b = append(b, 0x2c)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x2c,
	})
	if err != nil {
		return written, err
	}

	// L is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *INC_L) Length() uint8 { // 0x2c
	return 1
}

func (o *INC_L) cycles() []uint8 { // 0x2c
	return []uint8{4}
}

func (o *INC_L) String() string { // 0x2c
	return "INC" + " " + "L" /* operand1 */
}

func (o *INC_L) SymbolicString() string { // 0x2c
	return "INC L"
}

type DEC_L struct {
	addr         string      // 0x2d
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *DEC_L) Write(w io.Writer) (int, error) { // 0x2d
	var b []byte

	b = append(b, 0x2d)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x2d,
	})
	if err != nil {
		return written, err
	}

	// L is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *DEC_L) Length() uint8 { // 0x2d
	return 1
}

func (o *DEC_L) cycles() []uint8 { // 0x2d
	return []uint8{4}
}

func (o *DEC_L) String() string { // 0x2d
	return "DEC" + " " + "L" /* operand1 */
}

func (o *DEC_L) SymbolicString() string { // 0x2d
	return "DEC L"
}

type LD_L_d8 struct {
	addr         string      // 0x2e
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_L_d8) Write(w io.Writer) (int, error) { // 0x2e
	var b []byte

	b = append(b, 0x2e)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x2e,
	})
	if err != nil {
		return written, err
	}

	// L is implicit in this instruction.
	err = writeImmediate8BitData(w, o.operand2)
	written += 1

	return written, err
}

func (o *LD_L_d8) Length() uint8 { // 0x2e
	return 2
}

func (o *LD_L_d8) cycles() []uint8 { // 0x2e
	return []uint8{8}
}

func (o *LD_L_d8) String() string { // 0x2e
	return "LD" + " " + "L" /* operand1 */ + " " + fmt.Sprintf("%X", o.operand2)
}

func (o *LD_L_d8) SymbolicString() string { // 0x2e
	return "LD L,d8"
}

type CPL struct {
	addr         string      // 0x2f
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *CPL) Write(w io.Writer) (int, error) { // 0x2f
	var b []byte

	b = append(b, 0x2f)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x2f,
	})
	if err != nil {
		return written, err
	}

	//  is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *CPL) Length() uint8 { // 0x2f
	return 1
}

func (o *CPL) cycles() []uint8 { // 0x2f
	return []uint8{4}
}

func (o *CPL) String() string { // 0x2f
	return "CPL"
}

func (o *CPL) SymbolicString() string { // 0x2f
	return "CPL"
}

type INC_BC struct {
	addr         string      // 0x3
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *INC_BC) Write(w io.Writer) (int, error) { // 0x3
	var b []byte

	b = append(b, 0x3)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x3,
	})
	if err != nil {
		return written, err
	}

	// BC is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *INC_BC) Length() uint8 { // 0x3
	return 1
}

func (o *INC_BC) cycles() []uint8 { // 0x3
	return []uint8{8}
}

func (o *INC_BC) String() string { // 0x3
	return "INC" + " " + "BC" /* operand1 */
}

func (o *INC_BC) SymbolicString() string { // 0x3
	return "INC BC"
}

type JR_NC_r8 struct {
	addr         string      // 0x30
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *JR_NC_r8) Write(w io.Writer) (int, error) { // 0x30
	var b []byte

	b = append(b, 0x30)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x30,
	})
	if err != nil {
		return written, err
	}

	// NC is implicit in this instruction.
	err = writeImmediateSigned8BitData(w, o.operand2)
	written += 1

	return written, err
}

func (o *JR_NC_r8) Length() uint8 { // 0x30
	return 2
}

func (o *JR_NC_r8) cycles() []uint8 { // 0x30
	return []uint8{12, 8}
}

func (o *JR_NC_r8) String() string { // 0x30
	return "JR" + " " + "NC" /* operand1 */ + " " + fmt.Sprintf("SP+%X", o.operand2)
}

func (o *JR_NC_r8) SymbolicString() string { // 0x30
	return "JR NC,r8"
}

type LD_SP_d16 struct {
	addr         string      // 0x31
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_SP_d16) Write(w io.Writer) (int, error) { // 0x31
	var b []byte

	b = append(b, 0x31)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x31,
	})
	if err != nil {
		return written, err
	}

	// SP is implicit in this instruction.
	err = writeImmediate16BitData(w, o.operand2)
	written += 2

	return written, err
}

func (o *LD_SP_d16) Length() uint8 { // 0x31
	return 3
}

func (o *LD_SP_d16) cycles() []uint8 { // 0x31
	return []uint8{12}
}

func (o *LD_SP_d16) String() string { // 0x31
	return "LD" + " " + "SP" /* operand1 */ + " " + fmt.Sprintf("%X", o.operand2)
}

func (o *LD_SP_d16) SymbolicString() string { // 0x31
	return "LD SP,d16"
}

type LD_HLPtrDec_A struct {
	addr         string      // 0x32
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_HLPtrDec_A) Write(w io.Writer) (int, error) { // 0x32
	var b []byte

	b = append(b, 0x32)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x32,
	})
	if err != nil {
		return written, err
	}

	// (HL-) is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *LD_HLPtrDec_A) Length() uint8 { // 0x32
	return 1
}

func (o *LD_HLPtrDec_A) cycles() []uint8 { // 0x32
	return []uint8{8}
}

func (o *LD_HLPtrDec_A) String() string { // 0x32
	return "LD" + " " + "(HL-)" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *LD_HLPtrDec_A) SymbolicString() string { // 0x32
	return "LD (HL-),A"
}

type INC_SP struct {
	addr         string      // 0x33
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *INC_SP) Write(w io.Writer) (int, error) { // 0x33
	var b []byte

	b = append(b, 0x33)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x33,
	})
	if err != nil {
		return written, err
	}

	// SP is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *INC_SP) Length() uint8 { // 0x33
	return 1
}

func (o *INC_SP) cycles() []uint8 { // 0x33
	return []uint8{8}
}

func (o *INC_SP) String() string { // 0x33
	return "INC" + " " + "SP" /* operand1 */
}

func (o *INC_SP) SymbolicString() string { // 0x33
	return "INC SP"
}

type INC_HLPtr struct {
	addr         string      // 0x34
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *INC_HLPtr) Write(w io.Writer) (int, error) { // 0x34
	var b []byte

	b = append(b, 0x34)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x34,
	})
	if err != nil {
		return written, err
	}

	// (HL) is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *INC_HLPtr) Length() uint8 { // 0x34
	return 1
}

func (o *INC_HLPtr) cycles() []uint8 { // 0x34
	return []uint8{12}
}

func (o *INC_HLPtr) String() string { // 0x34
	return "INC" + " " + "(HL)" /* operand1 */
}

func (o *INC_HLPtr) SymbolicString() string { // 0x34
	return "INC (HL)"
}

type DEC_HLPtr struct {
	addr         string      // 0x35
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *DEC_HLPtr) Write(w io.Writer) (int, error) { // 0x35
	var b []byte

	b = append(b, 0x35)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x35,
	})
	if err != nil {
		return written, err
	}

	// (HL) is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *DEC_HLPtr) Length() uint8 { // 0x35
	return 1
}

func (o *DEC_HLPtr) cycles() []uint8 { // 0x35
	return []uint8{12}
}

func (o *DEC_HLPtr) String() string { // 0x35
	return "DEC" + " " + "(HL)" /* operand1 */
}

func (o *DEC_HLPtr) SymbolicString() string { // 0x35
	return "DEC (HL)"
}

type LD_HLPtr_d8 struct {
	addr         string      // 0x36
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_HLPtr_d8) Write(w io.Writer) (int, error) { // 0x36
	var b []byte

	b = append(b, 0x36)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x36,
	})
	if err != nil {
		return written, err
	}

	// (HL) is implicit in this instruction.
	err = writeImmediate8BitData(w, o.operand2)
	written += 1

	return written, err
}

func (o *LD_HLPtr_d8) Length() uint8 { // 0x36
	return 2
}

func (o *LD_HLPtr_d8) cycles() []uint8 { // 0x36
	return []uint8{12}
}

func (o *LD_HLPtr_d8) String() string { // 0x36
	return "LD" + " " + "(HL)" /* operand1 */ + " " + fmt.Sprintf("%X", o.operand2)
}

func (o *LD_HLPtr_d8) SymbolicString() string { // 0x36
	return "LD (HL),d8"
}

type SCF struct {
	addr         string      // 0x37
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SCF) Write(w io.Writer) (int, error) { // 0x37
	var b []byte

	b = append(b, 0x37)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x37,
	})
	if err != nil {
		return written, err
	}

	//  is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SCF) Length() uint8 { // 0x37
	return 1
}

func (o *SCF) cycles() []uint8 { // 0x37
	return []uint8{4}
}

func (o *SCF) String() string { // 0x37
	return "SCF"
}

func (o *SCF) SymbolicString() string { // 0x37
	return "SCF"
}

type JR_C_r8 struct {
	addr         string      // 0x38
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *JR_C_r8) Write(w io.Writer) (int, error) { // 0x38
	var b []byte

	b = append(b, 0x38)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x38,
	})
	if err != nil {
		return written, err
	}

	// C is implicit in this instruction.
	err = writeImmediateSigned8BitData(w, o.operand2)
	written += 1

	return written, err
}

func (o *JR_C_r8) Length() uint8 { // 0x38
	return 2
}

func (o *JR_C_r8) cycles() []uint8 { // 0x38
	return []uint8{12, 8}
}

func (o *JR_C_r8) String() string { // 0x38
	return "JR" + " " + "C" /* operand1 */ + " " + fmt.Sprintf("SP+%X", o.operand2)
}

func (o *JR_C_r8) SymbolicString() string { // 0x38
	return "JR C,r8"
}

type ADD_HL_SP struct {
	addr         string      // 0x39
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *ADD_HL_SP) Write(w io.Writer) (int, error) { // 0x39
	var b []byte

	b = append(b, 0x39)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x39,
	})
	if err != nil {
		return written, err
	}

	// HL is implicit in this instruction.
	// SP is implicit in this instruction.

	return written, err
}

func (o *ADD_HL_SP) Length() uint8 { // 0x39
	return 1
}

func (o *ADD_HL_SP) cycles() []uint8 { // 0x39
	return []uint8{8}
}

func (o *ADD_HL_SP) String() string { // 0x39
	return "ADD" + " " + "HL" /* operand1 */ + " " + "SP" /* operand2 */
}

func (o *ADD_HL_SP) SymbolicString() string { // 0x39
	return "ADD HL,SP"
}

type LD_A_HLPtrDec struct {
	addr         string      // 0x3a
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_A_HLPtrDec) Write(w io.Writer) (int, error) { // 0x3a
	var b []byte

	b = append(b, 0x3a)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x3a,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// (HL-) is implicit in this instruction.

	return written, err
}

func (o *LD_A_HLPtrDec) Length() uint8 { // 0x3a
	return 1
}

func (o *LD_A_HLPtrDec) cycles() []uint8 { // 0x3a
	return []uint8{8}
}

func (o *LD_A_HLPtrDec) String() string { // 0x3a
	return "LD" + " " + "A" /* operand1 */ + " " + "(HL-)" /* operand2 */
}

func (o *LD_A_HLPtrDec) SymbolicString() string { // 0x3a
	return "LD A,(HL-)"
}

type DEC_SP struct {
	addr         string      // 0x3b
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *DEC_SP) Write(w io.Writer) (int, error) { // 0x3b
	var b []byte

	b = append(b, 0x3b)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x3b,
	})
	if err != nil {
		return written, err
	}

	// SP is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *DEC_SP) Length() uint8 { // 0x3b
	return 1
}

func (o *DEC_SP) cycles() []uint8 { // 0x3b
	return []uint8{8}
}

func (o *DEC_SP) String() string { // 0x3b
	return "DEC" + " " + "SP" /* operand1 */
}

func (o *DEC_SP) SymbolicString() string { // 0x3b
	return "DEC SP"
}

type INC_A struct {
	addr         string      // 0x3c
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *INC_A) Write(w io.Writer) (int, error) { // 0x3c
	var b []byte

	b = append(b, 0x3c)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x3c,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *INC_A) Length() uint8 { // 0x3c
	return 1
}

func (o *INC_A) cycles() []uint8 { // 0x3c
	return []uint8{4}
}

func (o *INC_A) String() string { // 0x3c
	return "INC" + " " + "A" /* operand1 */
}

func (o *INC_A) SymbolicString() string { // 0x3c
	return "INC A"
}

type DEC_A struct {
	addr         string      // 0x3d
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *DEC_A) Write(w io.Writer) (int, error) { // 0x3d
	var b []byte

	b = append(b, 0x3d)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x3d,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *DEC_A) Length() uint8 { // 0x3d
	return 1
}

func (o *DEC_A) cycles() []uint8 { // 0x3d
	return []uint8{4}
}

func (o *DEC_A) String() string { // 0x3d
	return "DEC" + " " + "A" /* operand1 */
}

func (o *DEC_A) SymbolicString() string { // 0x3d
	return "DEC A"
}

type LD_A_d8 struct {
	addr         string      // 0x3e
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_A_d8) Write(w io.Writer) (int, error) { // 0x3e
	var b []byte

	b = append(b, 0x3e)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x3e,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	err = writeImmediate8BitData(w, o.operand2)
	written += 1

	return written, err
}

func (o *LD_A_d8) Length() uint8 { // 0x3e
	return 2
}

func (o *LD_A_d8) cycles() []uint8 { // 0x3e
	return []uint8{8}
}

func (o *LD_A_d8) String() string { // 0x3e
	return "LD" + " " + "A" /* operand1 */ + " " + fmt.Sprintf("%X", o.operand2)
}

func (o *LD_A_d8) SymbolicString() string { // 0x3e
	return "LD A,d8"
}

type CCF struct {
	addr         string      // 0x3f
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *CCF) Write(w io.Writer) (int, error) { // 0x3f
	var b []byte

	b = append(b, 0x3f)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x3f,
	})
	if err != nil {
		return written, err
	}

	//  is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *CCF) Length() uint8 { // 0x3f
	return 1
}

func (o *CCF) cycles() []uint8 { // 0x3f
	return []uint8{4}
}

func (o *CCF) String() string { // 0x3f
	return "CCF"
}

func (o *CCF) SymbolicString() string { // 0x3f
	return "CCF"
}

type INC_B struct {
	addr         string      // 0x4
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *INC_B) Write(w io.Writer) (int, error) { // 0x4
	var b []byte

	b = append(b, 0x4)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x4,
	})
	if err != nil {
		return written, err
	}

	// B is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *INC_B) Length() uint8 { // 0x4
	return 1
}

func (o *INC_B) cycles() []uint8 { // 0x4
	return []uint8{4}
}

func (o *INC_B) String() string { // 0x4
	return "INC" + " " + "B" /* operand1 */
}

func (o *INC_B) SymbolicString() string { // 0x4
	return "INC B"
}

type LD_B_B struct {
	addr         string      // 0x40
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_B_B) Write(w io.Writer) (int, error) { // 0x40
	var b []byte

	b = append(b, 0x40)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x40,
	})
	if err != nil {
		return written, err
	}

	// B is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *LD_B_B) Length() uint8 { // 0x40
	return 1
}

func (o *LD_B_B) cycles() []uint8 { // 0x40
	return []uint8{4}
}

func (o *LD_B_B) String() string { // 0x40
	return "LD" + " " + "B" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *LD_B_B) SymbolicString() string { // 0x40
	return "LD B,B"
}

type LD_B_C struct {
	addr         string      // 0x41
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_B_C) Write(w io.Writer) (int, error) { // 0x41
	var b []byte

	b = append(b, 0x41)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x41,
	})
	if err != nil {
		return written, err
	}

	// B is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *LD_B_C) Length() uint8 { // 0x41
	return 1
}

func (o *LD_B_C) cycles() []uint8 { // 0x41
	return []uint8{4}
}

func (o *LD_B_C) String() string { // 0x41
	return "LD" + " " + "B" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *LD_B_C) SymbolicString() string { // 0x41
	return "LD B,C"
}

type LD_B_D struct {
	addr         string      // 0x42
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_B_D) Write(w io.Writer) (int, error) { // 0x42
	var b []byte

	b = append(b, 0x42)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x42,
	})
	if err != nil {
		return written, err
	}

	// B is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *LD_B_D) Length() uint8 { // 0x42
	return 1
}

func (o *LD_B_D) cycles() []uint8 { // 0x42
	return []uint8{4}
}

func (o *LD_B_D) String() string { // 0x42
	return "LD" + " " + "B" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *LD_B_D) SymbolicString() string { // 0x42
	return "LD B,D"
}

type LD_B_E struct {
	addr         string      // 0x43
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_B_E) Write(w io.Writer) (int, error) { // 0x43
	var b []byte

	b = append(b, 0x43)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x43,
	})
	if err != nil {
		return written, err
	}

	// B is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *LD_B_E) Length() uint8 { // 0x43
	return 1
}

func (o *LD_B_E) cycles() []uint8 { // 0x43
	return []uint8{4}
}

func (o *LD_B_E) String() string { // 0x43
	return "LD" + " " + "B" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *LD_B_E) SymbolicString() string { // 0x43
	return "LD B,E"
}

type LD_B_H struct {
	addr         string      // 0x44
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_B_H) Write(w io.Writer) (int, error) { // 0x44
	var b []byte

	b = append(b, 0x44)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x44,
	})
	if err != nil {
		return written, err
	}

	// B is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *LD_B_H) Length() uint8 { // 0x44
	return 1
}

func (o *LD_B_H) cycles() []uint8 { // 0x44
	return []uint8{4}
}

func (o *LD_B_H) String() string { // 0x44
	return "LD" + " " + "B" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *LD_B_H) SymbolicString() string { // 0x44
	return "LD B,H"
}

type LD_B_L struct {
	addr         string      // 0x45
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_B_L) Write(w io.Writer) (int, error) { // 0x45
	var b []byte

	b = append(b, 0x45)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x45,
	})
	if err != nil {
		return written, err
	}

	// B is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *LD_B_L) Length() uint8 { // 0x45
	return 1
}

func (o *LD_B_L) cycles() []uint8 { // 0x45
	return []uint8{4}
}

func (o *LD_B_L) String() string { // 0x45
	return "LD" + " " + "B" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *LD_B_L) SymbolicString() string { // 0x45
	return "LD B,L"
}

type LD_B_HLPtr struct {
	addr         string      // 0x46
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_B_HLPtr) Write(w io.Writer) (int, error) { // 0x46
	var b []byte

	b = append(b, 0x46)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x46,
	})
	if err != nil {
		return written, err
	}

	// B is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *LD_B_HLPtr) Length() uint8 { // 0x46
	return 1
}

func (o *LD_B_HLPtr) cycles() []uint8 { // 0x46
	return []uint8{8}
}

func (o *LD_B_HLPtr) String() string { // 0x46
	return "LD" + " " + "B" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *LD_B_HLPtr) SymbolicString() string { // 0x46
	return "LD B,(HL)"
}

type LD_B_A struct {
	addr         string      // 0x47
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_B_A) Write(w io.Writer) (int, error) { // 0x47
	var b []byte

	b = append(b, 0x47)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x47,
	})
	if err != nil {
		return written, err
	}

	// B is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *LD_B_A) Length() uint8 { // 0x47
	return 1
}

func (o *LD_B_A) cycles() []uint8 { // 0x47
	return []uint8{4}
}

func (o *LD_B_A) String() string { // 0x47
	return "LD" + " " + "B" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *LD_B_A) SymbolicString() string { // 0x47
	return "LD B,A"
}

type LD_C_B struct {
	addr         string      // 0x48
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_C_B) Write(w io.Writer) (int, error) { // 0x48
	var b []byte

	b = append(b, 0x48)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x48,
	})
	if err != nil {
		return written, err
	}

	// C is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *LD_C_B) Length() uint8 { // 0x48
	return 1
}

func (o *LD_C_B) cycles() []uint8 { // 0x48
	return []uint8{4}
}

func (o *LD_C_B) String() string { // 0x48
	return "LD" + " " + "C" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *LD_C_B) SymbolicString() string { // 0x48
	return "LD C,B"
}

type LD_C_C struct {
	addr         string      // 0x49
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_C_C) Write(w io.Writer) (int, error) { // 0x49
	var b []byte

	b = append(b, 0x49)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x49,
	})
	if err != nil {
		return written, err
	}

	// C is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *LD_C_C) Length() uint8 { // 0x49
	return 1
}

func (o *LD_C_C) cycles() []uint8 { // 0x49
	return []uint8{4}
}

func (o *LD_C_C) String() string { // 0x49
	return "LD" + " " + "C" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *LD_C_C) SymbolicString() string { // 0x49
	return "LD C,C"
}

type LD_C_D struct {
	addr         string      // 0x4a
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_C_D) Write(w io.Writer) (int, error) { // 0x4a
	var b []byte

	b = append(b, 0x4a)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x4a,
	})
	if err != nil {
		return written, err
	}

	// C is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *LD_C_D) Length() uint8 { // 0x4a
	return 1
}

func (o *LD_C_D) cycles() []uint8 { // 0x4a
	return []uint8{4}
}

func (o *LD_C_D) String() string { // 0x4a
	return "LD" + " " + "C" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *LD_C_D) SymbolicString() string { // 0x4a
	return "LD C,D"
}

type LD_C_E struct {
	addr         string      // 0x4b
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_C_E) Write(w io.Writer) (int, error) { // 0x4b
	var b []byte

	b = append(b, 0x4b)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x4b,
	})
	if err != nil {
		return written, err
	}

	// C is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *LD_C_E) Length() uint8 { // 0x4b
	return 1
}

func (o *LD_C_E) cycles() []uint8 { // 0x4b
	return []uint8{4}
}

func (o *LD_C_E) String() string { // 0x4b
	return "LD" + " " + "C" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *LD_C_E) SymbolicString() string { // 0x4b
	return "LD C,E"
}

type LD_C_H struct {
	addr         string      // 0x4c
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_C_H) Write(w io.Writer) (int, error) { // 0x4c
	var b []byte

	b = append(b, 0x4c)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x4c,
	})
	if err != nil {
		return written, err
	}

	// C is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *LD_C_H) Length() uint8 { // 0x4c
	return 1
}

func (o *LD_C_H) cycles() []uint8 { // 0x4c
	return []uint8{4}
}

func (o *LD_C_H) String() string { // 0x4c
	return "LD" + " " + "C" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *LD_C_H) SymbolicString() string { // 0x4c
	return "LD C,H"
}

type LD_C_L struct {
	addr         string      // 0x4d
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_C_L) Write(w io.Writer) (int, error) { // 0x4d
	var b []byte

	b = append(b, 0x4d)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x4d,
	})
	if err != nil {
		return written, err
	}

	// C is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *LD_C_L) Length() uint8 { // 0x4d
	return 1
}

func (o *LD_C_L) cycles() []uint8 { // 0x4d
	return []uint8{4}
}

func (o *LD_C_L) String() string { // 0x4d
	return "LD" + " " + "C" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *LD_C_L) SymbolicString() string { // 0x4d
	return "LD C,L"
}

type LD_C_HLPtr struct {
	addr         string      // 0x4e
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_C_HLPtr) Write(w io.Writer) (int, error) { // 0x4e
	var b []byte

	b = append(b, 0x4e)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x4e,
	})
	if err != nil {
		return written, err
	}

	// C is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *LD_C_HLPtr) Length() uint8 { // 0x4e
	return 1
}

func (o *LD_C_HLPtr) cycles() []uint8 { // 0x4e
	return []uint8{8}
}

func (o *LD_C_HLPtr) String() string { // 0x4e
	return "LD" + " " + "C" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *LD_C_HLPtr) SymbolicString() string { // 0x4e
	return "LD C,(HL)"
}

type LD_C_A struct {
	addr         string      // 0x4f
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_C_A) Write(w io.Writer) (int, error) { // 0x4f
	var b []byte

	b = append(b, 0x4f)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x4f,
	})
	if err != nil {
		return written, err
	}

	// C is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *LD_C_A) Length() uint8 { // 0x4f
	return 1
}

func (o *LD_C_A) cycles() []uint8 { // 0x4f
	return []uint8{4}
}

func (o *LD_C_A) String() string { // 0x4f
	return "LD" + " " + "C" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *LD_C_A) SymbolicString() string { // 0x4f
	return "LD C,A"
}

type DEC_B struct {
	addr         string      // 0x5
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *DEC_B) Write(w io.Writer) (int, error) { // 0x5
	var b []byte

	b = append(b, 0x5)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x5,
	})
	if err != nil {
		return written, err
	}

	// B is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *DEC_B) Length() uint8 { // 0x5
	return 1
}

func (o *DEC_B) cycles() []uint8 { // 0x5
	return []uint8{4}
}

func (o *DEC_B) String() string { // 0x5
	return "DEC" + " " + "B" /* operand1 */
}

func (o *DEC_B) SymbolicString() string { // 0x5
	return "DEC B"
}

type LD_D_B struct {
	addr         string      // 0x50
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_D_B) Write(w io.Writer) (int, error) { // 0x50
	var b []byte

	b = append(b, 0x50)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x50,
	})
	if err != nil {
		return written, err
	}

	// D is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *LD_D_B) Length() uint8 { // 0x50
	return 1
}

func (o *LD_D_B) cycles() []uint8 { // 0x50
	return []uint8{4}
}

func (o *LD_D_B) String() string { // 0x50
	return "LD" + " " + "D" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *LD_D_B) SymbolicString() string { // 0x50
	return "LD D,B"
}

type LD_D_C struct {
	addr         string      // 0x51
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_D_C) Write(w io.Writer) (int, error) { // 0x51
	var b []byte

	b = append(b, 0x51)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x51,
	})
	if err != nil {
		return written, err
	}

	// D is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *LD_D_C) Length() uint8 { // 0x51
	return 1
}

func (o *LD_D_C) cycles() []uint8 { // 0x51
	return []uint8{4}
}

func (o *LD_D_C) String() string { // 0x51
	return "LD" + " " + "D" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *LD_D_C) SymbolicString() string { // 0x51
	return "LD D,C"
}

type LD_D_D struct {
	addr         string      // 0x52
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_D_D) Write(w io.Writer) (int, error) { // 0x52
	var b []byte

	b = append(b, 0x52)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x52,
	})
	if err != nil {
		return written, err
	}

	// D is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *LD_D_D) Length() uint8 { // 0x52
	return 1
}

func (o *LD_D_D) cycles() []uint8 { // 0x52
	return []uint8{4}
}

func (o *LD_D_D) String() string { // 0x52
	return "LD" + " " + "D" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *LD_D_D) SymbolicString() string { // 0x52
	return "LD D,D"
}

type LD_D_E struct {
	addr         string      // 0x53
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_D_E) Write(w io.Writer) (int, error) { // 0x53
	var b []byte

	b = append(b, 0x53)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x53,
	})
	if err != nil {
		return written, err
	}

	// D is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *LD_D_E) Length() uint8 { // 0x53
	return 1
}

func (o *LD_D_E) cycles() []uint8 { // 0x53
	return []uint8{4}
}

func (o *LD_D_E) String() string { // 0x53
	return "LD" + " " + "D" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *LD_D_E) SymbolicString() string { // 0x53
	return "LD D,E"
}

type LD_D_H struct {
	addr         string      // 0x54
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_D_H) Write(w io.Writer) (int, error) { // 0x54
	var b []byte

	b = append(b, 0x54)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x54,
	})
	if err != nil {
		return written, err
	}

	// D is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *LD_D_H) Length() uint8 { // 0x54
	return 1
}

func (o *LD_D_H) cycles() []uint8 { // 0x54
	return []uint8{4}
}

func (o *LD_D_H) String() string { // 0x54
	return "LD" + " " + "D" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *LD_D_H) SymbolicString() string { // 0x54
	return "LD D,H"
}

type LD_D_L struct {
	addr         string      // 0x55
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_D_L) Write(w io.Writer) (int, error) { // 0x55
	var b []byte

	b = append(b, 0x55)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x55,
	})
	if err != nil {
		return written, err
	}

	// D is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *LD_D_L) Length() uint8 { // 0x55
	return 1
}

func (o *LD_D_L) cycles() []uint8 { // 0x55
	return []uint8{4}
}

func (o *LD_D_L) String() string { // 0x55
	return "LD" + " " + "D" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *LD_D_L) SymbolicString() string { // 0x55
	return "LD D,L"
}

type LD_D_HLPtr struct {
	addr         string      // 0x56
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_D_HLPtr) Write(w io.Writer) (int, error) { // 0x56
	var b []byte

	b = append(b, 0x56)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x56,
	})
	if err != nil {
		return written, err
	}

	// D is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *LD_D_HLPtr) Length() uint8 { // 0x56
	return 1
}

func (o *LD_D_HLPtr) cycles() []uint8 { // 0x56
	return []uint8{8}
}

func (o *LD_D_HLPtr) String() string { // 0x56
	return "LD" + " " + "D" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *LD_D_HLPtr) SymbolicString() string { // 0x56
	return "LD D,(HL)"
}

type LD_D_A struct {
	addr         string      // 0x57
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_D_A) Write(w io.Writer) (int, error) { // 0x57
	var b []byte

	b = append(b, 0x57)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x57,
	})
	if err != nil {
		return written, err
	}

	// D is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *LD_D_A) Length() uint8 { // 0x57
	return 1
}

func (o *LD_D_A) cycles() []uint8 { // 0x57
	return []uint8{4}
}

func (o *LD_D_A) String() string { // 0x57
	return "LD" + " " + "D" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *LD_D_A) SymbolicString() string { // 0x57
	return "LD D,A"
}

type LD_E_B struct {
	addr         string      // 0x58
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_E_B) Write(w io.Writer) (int, error) { // 0x58
	var b []byte

	b = append(b, 0x58)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x58,
	})
	if err != nil {
		return written, err
	}

	// E is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *LD_E_B) Length() uint8 { // 0x58
	return 1
}

func (o *LD_E_B) cycles() []uint8 { // 0x58
	return []uint8{4}
}

func (o *LD_E_B) String() string { // 0x58
	return "LD" + " " + "E" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *LD_E_B) SymbolicString() string { // 0x58
	return "LD E,B"
}

type LD_E_C struct {
	addr         string      // 0x59
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_E_C) Write(w io.Writer) (int, error) { // 0x59
	var b []byte

	b = append(b, 0x59)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x59,
	})
	if err != nil {
		return written, err
	}

	// E is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *LD_E_C) Length() uint8 { // 0x59
	return 1
}

func (o *LD_E_C) cycles() []uint8 { // 0x59
	return []uint8{4}
}

func (o *LD_E_C) String() string { // 0x59
	return "LD" + " " + "E" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *LD_E_C) SymbolicString() string { // 0x59
	return "LD E,C"
}

type LD_E_D struct {
	addr         string      // 0x5a
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_E_D) Write(w io.Writer) (int, error) { // 0x5a
	var b []byte

	b = append(b, 0x5a)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x5a,
	})
	if err != nil {
		return written, err
	}

	// E is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *LD_E_D) Length() uint8 { // 0x5a
	return 1
}

func (o *LD_E_D) cycles() []uint8 { // 0x5a
	return []uint8{4}
}

func (o *LD_E_D) String() string { // 0x5a
	return "LD" + " " + "E" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *LD_E_D) SymbolicString() string { // 0x5a
	return "LD E,D"
}

type LD_E_E struct {
	addr         string      // 0x5b
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_E_E) Write(w io.Writer) (int, error) { // 0x5b
	var b []byte

	b = append(b, 0x5b)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x5b,
	})
	if err != nil {
		return written, err
	}

	// E is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *LD_E_E) Length() uint8 { // 0x5b
	return 1
}

func (o *LD_E_E) cycles() []uint8 { // 0x5b
	return []uint8{4}
}

func (o *LD_E_E) String() string { // 0x5b
	return "LD" + " " + "E" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *LD_E_E) SymbolicString() string { // 0x5b
	return "LD E,E"
}

type LD_E_H struct {
	addr         string      // 0x5c
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_E_H) Write(w io.Writer) (int, error) { // 0x5c
	var b []byte

	b = append(b, 0x5c)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x5c,
	})
	if err != nil {
		return written, err
	}

	// E is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *LD_E_H) Length() uint8 { // 0x5c
	return 1
}

func (o *LD_E_H) cycles() []uint8 { // 0x5c
	return []uint8{4}
}

func (o *LD_E_H) String() string { // 0x5c
	return "LD" + " " + "E" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *LD_E_H) SymbolicString() string { // 0x5c
	return "LD E,H"
}

type LD_E_L struct {
	addr         string      // 0x5d
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_E_L) Write(w io.Writer) (int, error) { // 0x5d
	var b []byte

	b = append(b, 0x5d)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x5d,
	})
	if err != nil {
		return written, err
	}

	// E is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *LD_E_L) Length() uint8 { // 0x5d
	return 1
}

func (o *LD_E_L) cycles() []uint8 { // 0x5d
	return []uint8{4}
}

func (o *LD_E_L) String() string { // 0x5d
	return "LD" + " " + "E" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *LD_E_L) SymbolicString() string { // 0x5d
	return "LD E,L"
}

type LD_E_HLPtr struct {
	addr         string      // 0x5e
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_E_HLPtr) Write(w io.Writer) (int, error) { // 0x5e
	var b []byte

	b = append(b, 0x5e)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x5e,
	})
	if err != nil {
		return written, err
	}

	// E is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *LD_E_HLPtr) Length() uint8 { // 0x5e
	return 1
}

func (o *LD_E_HLPtr) cycles() []uint8 { // 0x5e
	return []uint8{8}
}

func (o *LD_E_HLPtr) String() string { // 0x5e
	return "LD" + " " + "E" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *LD_E_HLPtr) SymbolicString() string { // 0x5e
	return "LD E,(HL)"
}

type LD_E_A struct {
	addr         string      // 0x5f
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_E_A) Write(w io.Writer) (int, error) { // 0x5f
	var b []byte

	b = append(b, 0x5f)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x5f,
	})
	if err != nil {
		return written, err
	}

	// E is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *LD_E_A) Length() uint8 { // 0x5f
	return 1
}

func (o *LD_E_A) cycles() []uint8 { // 0x5f
	return []uint8{4}
}

func (o *LD_E_A) String() string { // 0x5f
	return "LD" + " " + "E" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *LD_E_A) SymbolicString() string { // 0x5f
	return "LD E,A"
}

type LD_B_d8 struct {
	addr         string      // 0x6
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_B_d8) Write(w io.Writer) (int, error) { // 0x6
	var b []byte

	b = append(b, 0x6)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x6,
	})
	if err != nil {
		return written, err
	}

	// B is implicit in this instruction.
	err = writeImmediate8BitData(w, o.operand2)
	written += 1

	return written, err
}

func (o *LD_B_d8) Length() uint8 { // 0x6
	return 2
}

func (o *LD_B_d8) cycles() []uint8 { // 0x6
	return []uint8{8}
}

func (o *LD_B_d8) String() string { // 0x6
	return "LD" + " " + "B" /* operand1 */ + " " + fmt.Sprintf("%X", o.operand2)
}

func (o *LD_B_d8) SymbolicString() string { // 0x6
	return "LD B,d8"
}

type LD_H_B struct {
	addr         string      // 0x60
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_H_B) Write(w io.Writer) (int, error) { // 0x60
	var b []byte

	b = append(b, 0x60)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x60,
	})
	if err != nil {
		return written, err
	}

	// H is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *LD_H_B) Length() uint8 { // 0x60
	return 1
}

func (o *LD_H_B) cycles() []uint8 { // 0x60
	return []uint8{4}
}

func (o *LD_H_B) String() string { // 0x60
	return "LD" + " " + "H" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *LD_H_B) SymbolicString() string { // 0x60
	return "LD H,B"
}

type LD_H_C struct {
	addr         string      // 0x61
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_H_C) Write(w io.Writer) (int, error) { // 0x61
	var b []byte

	b = append(b, 0x61)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x61,
	})
	if err != nil {
		return written, err
	}

	// H is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *LD_H_C) Length() uint8 { // 0x61
	return 1
}

func (o *LD_H_C) cycles() []uint8 { // 0x61
	return []uint8{4}
}

func (o *LD_H_C) String() string { // 0x61
	return "LD" + " " + "H" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *LD_H_C) SymbolicString() string { // 0x61
	return "LD H,C"
}

type LD_H_D struct {
	addr         string      // 0x62
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_H_D) Write(w io.Writer) (int, error) { // 0x62
	var b []byte

	b = append(b, 0x62)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x62,
	})
	if err != nil {
		return written, err
	}

	// H is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *LD_H_D) Length() uint8 { // 0x62
	return 1
}

func (o *LD_H_D) cycles() []uint8 { // 0x62
	return []uint8{4}
}

func (o *LD_H_D) String() string { // 0x62
	return "LD" + " " + "H" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *LD_H_D) SymbolicString() string { // 0x62
	return "LD H,D"
}

type LD_H_E struct {
	addr         string      // 0x63
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_H_E) Write(w io.Writer) (int, error) { // 0x63
	var b []byte

	b = append(b, 0x63)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x63,
	})
	if err != nil {
		return written, err
	}

	// H is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *LD_H_E) Length() uint8 { // 0x63
	return 1
}

func (o *LD_H_E) cycles() []uint8 { // 0x63
	return []uint8{4}
}

func (o *LD_H_E) String() string { // 0x63
	return "LD" + " " + "H" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *LD_H_E) SymbolicString() string { // 0x63
	return "LD H,E"
}

type LD_H_H struct {
	addr         string      // 0x64
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_H_H) Write(w io.Writer) (int, error) { // 0x64
	var b []byte

	b = append(b, 0x64)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x64,
	})
	if err != nil {
		return written, err
	}

	// H is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *LD_H_H) Length() uint8 { // 0x64
	return 1
}

func (o *LD_H_H) cycles() []uint8 { // 0x64
	return []uint8{4}
}

func (o *LD_H_H) String() string { // 0x64
	return "LD" + " " + "H" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *LD_H_H) SymbolicString() string { // 0x64
	return "LD H,H"
}

type LD_H_L struct {
	addr         string      // 0x65
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_H_L) Write(w io.Writer) (int, error) { // 0x65
	var b []byte

	b = append(b, 0x65)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x65,
	})
	if err != nil {
		return written, err
	}

	// H is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *LD_H_L) Length() uint8 { // 0x65
	return 1
}

func (o *LD_H_L) cycles() []uint8 { // 0x65
	return []uint8{4}
}

func (o *LD_H_L) String() string { // 0x65
	return "LD" + " " + "H" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *LD_H_L) SymbolicString() string { // 0x65
	return "LD H,L"
}

type LD_H_HLPtr struct {
	addr         string      // 0x66
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_H_HLPtr) Write(w io.Writer) (int, error) { // 0x66
	var b []byte

	b = append(b, 0x66)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x66,
	})
	if err != nil {
		return written, err
	}

	// H is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *LD_H_HLPtr) Length() uint8 { // 0x66
	return 1
}

func (o *LD_H_HLPtr) cycles() []uint8 { // 0x66
	return []uint8{8}
}

func (o *LD_H_HLPtr) String() string { // 0x66
	return "LD" + " " + "H" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *LD_H_HLPtr) SymbolicString() string { // 0x66
	return "LD H,(HL)"
}

type LD_H_A struct {
	addr         string      // 0x67
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_H_A) Write(w io.Writer) (int, error) { // 0x67
	var b []byte

	b = append(b, 0x67)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x67,
	})
	if err != nil {
		return written, err
	}

	// H is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *LD_H_A) Length() uint8 { // 0x67
	return 1
}

func (o *LD_H_A) cycles() []uint8 { // 0x67
	return []uint8{4}
}

func (o *LD_H_A) String() string { // 0x67
	return "LD" + " " + "H" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *LD_H_A) SymbolicString() string { // 0x67
	return "LD H,A"
}

type LD_L_B struct {
	addr         string      // 0x68
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_L_B) Write(w io.Writer) (int, error) { // 0x68
	var b []byte

	b = append(b, 0x68)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x68,
	})
	if err != nil {
		return written, err
	}

	// L is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *LD_L_B) Length() uint8 { // 0x68
	return 1
}

func (o *LD_L_B) cycles() []uint8 { // 0x68
	return []uint8{4}
}

func (o *LD_L_B) String() string { // 0x68
	return "LD" + " " + "L" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *LD_L_B) SymbolicString() string { // 0x68
	return "LD L,B"
}

type LD_L_C struct {
	addr         string      // 0x69
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_L_C) Write(w io.Writer) (int, error) { // 0x69
	var b []byte

	b = append(b, 0x69)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x69,
	})
	if err != nil {
		return written, err
	}

	// L is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *LD_L_C) Length() uint8 { // 0x69
	return 1
}

func (o *LD_L_C) cycles() []uint8 { // 0x69
	return []uint8{4}
}

func (o *LD_L_C) String() string { // 0x69
	return "LD" + " " + "L" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *LD_L_C) SymbolicString() string { // 0x69
	return "LD L,C"
}

type LD_L_D struct {
	addr         string      // 0x6a
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_L_D) Write(w io.Writer) (int, error) { // 0x6a
	var b []byte

	b = append(b, 0x6a)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x6a,
	})
	if err != nil {
		return written, err
	}

	// L is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *LD_L_D) Length() uint8 { // 0x6a
	return 1
}

func (o *LD_L_D) cycles() []uint8 { // 0x6a
	return []uint8{4}
}

func (o *LD_L_D) String() string { // 0x6a
	return "LD" + " " + "L" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *LD_L_D) SymbolicString() string { // 0x6a
	return "LD L,D"
}

type LD_L_E struct {
	addr         string      // 0x6b
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_L_E) Write(w io.Writer) (int, error) { // 0x6b
	var b []byte

	b = append(b, 0x6b)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x6b,
	})
	if err != nil {
		return written, err
	}

	// L is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *LD_L_E) Length() uint8 { // 0x6b
	return 1
}

func (o *LD_L_E) cycles() []uint8 { // 0x6b
	return []uint8{4}
}

func (o *LD_L_E) String() string { // 0x6b
	return "LD" + " " + "L" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *LD_L_E) SymbolicString() string { // 0x6b
	return "LD L,E"
}

type LD_L_H struct {
	addr         string      // 0x6c
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_L_H) Write(w io.Writer) (int, error) { // 0x6c
	var b []byte

	b = append(b, 0x6c)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x6c,
	})
	if err != nil {
		return written, err
	}

	// L is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *LD_L_H) Length() uint8 { // 0x6c
	return 1
}

func (o *LD_L_H) cycles() []uint8 { // 0x6c
	return []uint8{4}
}

func (o *LD_L_H) String() string { // 0x6c
	return "LD" + " " + "L" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *LD_L_H) SymbolicString() string { // 0x6c
	return "LD L,H"
}

type LD_L_L struct {
	addr         string      // 0x6d
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_L_L) Write(w io.Writer) (int, error) { // 0x6d
	var b []byte

	b = append(b, 0x6d)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x6d,
	})
	if err != nil {
		return written, err
	}

	// L is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *LD_L_L) Length() uint8 { // 0x6d
	return 1
}

func (o *LD_L_L) cycles() []uint8 { // 0x6d
	return []uint8{4}
}

func (o *LD_L_L) String() string { // 0x6d
	return "LD" + " " + "L" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *LD_L_L) SymbolicString() string { // 0x6d
	return "LD L,L"
}

type LD_L_HLPtr struct {
	addr         string      // 0x6e
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_L_HLPtr) Write(w io.Writer) (int, error) { // 0x6e
	var b []byte

	b = append(b, 0x6e)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x6e,
	})
	if err != nil {
		return written, err
	}

	// L is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *LD_L_HLPtr) Length() uint8 { // 0x6e
	return 1
}

func (o *LD_L_HLPtr) cycles() []uint8 { // 0x6e
	return []uint8{8}
}

func (o *LD_L_HLPtr) String() string { // 0x6e
	return "LD" + " " + "L" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *LD_L_HLPtr) SymbolicString() string { // 0x6e
	return "LD L,(HL)"
}

type LD_L_A struct {
	addr         string      // 0x6f
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_L_A) Write(w io.Writer) (int, error) { // 0x6f
	var b []byte

	b = append(b, 0x6f)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x6f,
	})
	if err != nil {
		return written, err
	}

	// L is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *LD_L_A) Length() uint8 { // 0x6f
	return 1
}

func (o *LD_L_A) cycles() []uint8 { // 0x6f
	return []uint8{4}
}

func (o *LD_L_A) String() string { // 0x6f
	return "LD" + " " + "L" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *LD_L_A) SymbolicString() string { // 0x6f
	return "LD L,A"
}

type RLCA struct {
	addr         string      // 0x7
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RLCA) Write(w io.Writer) (int, error) { // 0x7
	var b []byte

	b = append(b, 0x7)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x7,
	})
	if err != nil {
		return written, err
	}

	//  is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RLCA) Length() uint8 { // 0x7
	return 1
}

func (o *RLCA) cycles() []uint8 { // 0x7
	return []uint8{4}
}

func (o *RLCA) String() string { // 0x7
	return "RLCA"
}

func (o *RLCA) SymbolicString() string { // 0x7
	return "RLCA"
}

type LD_HLPtr_B struct {
	addr         string      // 0x70
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_HLPtr_B) Write(w io.Writer) (int, error) { // 0x70
	var b []byte

	b = append(b, 0x70)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x70,
	})
	if err != nil {
		return written, err
	}

	// (HL) is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *LD_HLPtr_B) Length() uint8 { // 0x70
	return 1
}

func (o *LD_HLPtr_B) cycles() []uint8 { // 0x70
	return []uint8{8}
}

func (o *LD_HLPtr_B) String() string { // 0x70
	return "LD" + " " + "(HL)" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *LD_HLPtr_B) SymbolicString() string { // 0x70
	return "LD (HL),B"
}

type LD_HLPtr_C struct {
	addr         string      // 0x71
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_HLPtr_C) Write(w io.Writer) (int, error) { // 0x71
	var b []byte

	b = append(b, 0x71)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x71,
	})
	if err != nil {
		return written, err
	}

	// (HL) is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *LD_HLPtr_C) Length() uint8 { // 0x71
	return 1
}

func (o *LD_HLPtr_C) cycles() []uint8 { // 0x71
	return []uint8{8}
}

func (o *LD_HLPtr_C) String() string { // 0x71
	return "LD" + " " + "(HL)" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *LD_HLPtr_C) SymbolicString() string { // 0x71
	return "LD (HL),C"
}

type LD_HLPtr_D struct {
	addr         string      // 0x72
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_HLPtr_D) Write(w io.Writer) (int, error) { // 0x72
	var b []byte

	b = append(b, 0x72)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x72,
	})
	if err != nil {
		return written, err
	}

	// (HL) is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *LD_HLPtr_D) Length() uint8 { // 0x72
	return 1
}

func (o *LD_HLPtr_D) cycles() []uint8 { // 0x72
	return []uint8{8}
}

func (o *LD_HLPtr_D) String() string { // 0x72
	return "LD" + " " + "(HL)" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *LD_HLPtr_D) SymbolicString() string { // 0x72
	return "LD (HL),D"
}

type LD_HLPtr_E struct {
	addr         string      // 0x73
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_HLPtr_E) Write(w io.Writer) (int, error) { // 0x73
	var b []byte

	b = append(b, 0x73)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x73,
	})
	if err != nil {
		return written, err
	}

	// (HL) is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *LD_HLPtr_E) Length() uint8 { // 0x73
	return 1
}

func (o *LD_HLPtr_E) cycles() []uint8 { // 0x73
	return []uint8{8}
}

func (o *LD_HLPtr_E) String() string { // 0x73
	return "LD" + " " + "(HL)" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *LD_HLPtr_E) SymbolicString() string { // 0x73
	return "LD (HL),E"
}

type LD_HLPtr_H struct {
	addr         string      // 0x74
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_HLPtr_H) Write(w io.Writer) (int, error) { // 0x74
	var b []byte

	b = append(b, 0x74)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x74,
	})
	if err != nil {
		return written, err
	}

	// (HL) is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *LD_HLPtr_H) Length() uint8 { // 0x74
	return 1
}

func (o *LD_HLPtr_H) cycles() []uint8 { // 0x74
	return []uint8{8}
}

func (o *LD_HLPtr_H) String() string { // 0x74
	return "LD" + " " + "(HL)" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *LD_HLPtr_H) SymbolicString() string { // 0x74
	return "LD (HL),H"
}

type LD_HLPtr_L struct {
	addr         string      // 0x75
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_HLPtr_L) Write(w io.Writer) (int, error) { // 0x75
	var b []byte

	b = append(b, 0x75)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x75,
	})
	if err != nil {
		return written, err
	}

	// (HL) is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *LD_HLPtr_L) Length() uint8 { // 0x75
	return 1
}

func (o *LD_HLPtr_L) cycles() []uint8 { // 0x75
	return []uint8{8}
}

func (o *LD_HLPtr_L) String() string { // 0x75
	return "LD" + " " + "(HL)" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *LD_HLPtr_L) SymbolicString() string { // 0x75
	return "LD (HL),L"
}

type HALT struct {
	addr         string      // 0x76
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *HALT) Write(w io.Writer) (int, error) { // 0x76
	var b []byte

	b = append(b, 0x76)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x76,
	})
	if err != nil {
		return written, err
	}

	//  is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *HALT) Length() uint8 { // 0x76
	return 1
}

func (o *HALT) cycles() []uint8 { // 0x76
	return []uint8{4}
}

func (o *HALT) String() string { // 0x76
	return "HALT"
}

func (o *HALT) SymbolicString() string { // 0x76
	return "HALT"
}

type LD_HLPtr_A struct {
	addr         string      // 0x77
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_HLPtr_A) Write(w io.Writer) (int, error) { // 0x77
	var b []byte

	b = append(b, 0x77)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x77,
	})
	if err != nil {
		return written, err
	}

	// (HL) is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *LD_HLPtr_A) Length() uint8 { // 0x77
	return 1
}

func (o *LD_HLPtr_A) cycles() []uint8 { // 0x77
	return []uint8{8}
}

func (o *LD_HLPtr_A) String() string { // 0x77
	return "LD" + " " + "(HL)" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *LD_HLPtr_A) SymbolicString() string { // 0x77
	return "LD (HL),A"
}

type LD_A_B struct {
	addr         string      // 0x78
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_A_B) Write(w io.Writer) (int, error) { // 0x78
	var b []byte

	b = append(b, 0x78)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x78,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *LD_A_B) Length() uint8 { // 0x78
	return 1
}

func (o *LD_A_B) cycles() []uint8 { // 0x78
	return []uint8{4}
}

func (o *LD_A_B) String() string { // 0x78
	return "LD" + " " + "A" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *LD_A_B) SymbolicString() string { // 0x78
	return "LD A,B"
}

type LD_A_C struct {
	addr         string      // 0x79
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_A_C) Write(w io.Writer) (int, error) { // 0x79
	var b []byte

	b = append(b, 0x79)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x79,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *LD_A_C) Length() uint8 { // 0x79
	return 1
}

func (o *LD_A_C) cycles() []uint8 { // 0x79
	return []uint8{4}
}

func (o *LD_A_C) String() string { // 0x79
	return "LD" + " " + "A" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *LD_A_C) SymbolicString() string { // 0x79
	return "LD A,C"
}

type LD_A_D struct {
	addr         string      // 0x7a
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_A_D) Write(w io.Writer) (int, error) { // 0x7a
	var b []byte

	b = append(b, 0x7a)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x7a,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *LD_A_D) Length() uint8 { // 0x7a
	return 1
}

func (o *LD_A_D) cycles() []uint8 { // 0x7a
	return []uint8{4}
}

func (o *LD_A_D) String() string { // 0x7a
	return "LD" + " " + "A" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *LD_A_D) SymbolicString() string { // 0x7a
	return "LD A,D"
}

type LD_A_E struct {
	addr         string      // 0x7b
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_A_E) Write(w io.Writer) (int, error) { // 0x7b
	var b []byte

	b = append(b, 0x7b)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x7b,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *LD_A_E) Length() uint8 { // 0x7b
	return 1
}

func (o *LD_A_E) cycles() []uint8 { // 0x7b
	return []uint8{4}
}

func (o *LD_A_E) String() string { // 0x7b
	return "LD" + " " + "A" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *LD_A_E) SymbolicString() string { // 0x7b
	return "LD A,E"
}

type LD_A_H struct {
	addr         string      // 0x7c
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_A_H) Write(w io.Writer) (int, error) { // 0x7c
	var b []byte

	b = append(b, 0x7c)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x7c,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *LD_A_H) Length() uint8 { // 0x7c
	return 1
}

func (o *LD_A_H) cycles() []uint8 { // 0x7c
	return []uint8{4}
}

func (o *LD_A_H) String() string { // 0x7c
	return "LD" + " " + "A" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *LD_A_H) SymbolicString() string { // 0x7c
	return "LD A,H"
}

type LD_A_L struct {
	addr         string      // 0x7d
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_A_L) Write(w io.Writer) (int, error) { // 0x7d
	var b []byte

	b = append(b, 0x7d)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x7d,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *LD_A_L) Length() uint8 { // 0x7d
	return 1
}

func (o *LD_A_L) cycles() []uint8 { // 0x7d
	return []uint8{4}
}

func (o *LD_A_L) String() string { // 0x7d
	return "LD" + " " + "A" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *LD_A_L) SymbolicString() string { // 0x7d
	return "LD A,L"
}

type LD_A_HLPtr struct {
	addr         string      // 0x7e
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_A_HLPtr) Write(w io.Writer) (int, error) { // 0x7e
	var b []byte

	b = append(b, 0x7e)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x7e,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *LD_A_HLPtr) Length() uint8 { // 0x7e
	return 1
}

func (o *LD_A_HLPtr) cycles() []uint8 { // 0x7e
	return []uint8{8}
}

func (o *LD_A_HLPtr) String() string { // 0x7e
	return "LD" + " " + "A" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *LD_A_HLPtr) SymbolicString() string { // 0x7e
	return "LD A,(HL)"
}

type LD_A_A struct {
	addr         string      // 0x7f
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_A_A) Write(w io.Writer) (int, error) { // 0x7f
	var b []byte

	b = append(b, 0x7f)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x7f,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *LD_A_A) Length() uint8 { // 0x7f
	return 1
}

func (o *LD_A_A) cycles() []uint8 { // 0x7f
	return []uint8{4}
}

func (o *LD_A_A) String() string { // 0x7f
	return "LD" + " " + "A" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *LD_A_A) SymbolicString() string { // 0x7f
	return "LD A,A"
}

type LD_a16Deref_SP struct {
	addr         string      // 0x8
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_a16Deref_SP) Write(w io.Writer) (int, error) { // 0x8
	var b []byte

	b = append(b, 0x8)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x8,
	})
	if err != nil {
		return written, err
	}

	err = writeImmediate16BitAddress(w, o.operand1)
	written += 2
	// SP is implicit in this instruction.

	return written, err
}

func (o *LD_a16Deref_SP) Length() uint8 { // 0x8
	return 3
}

func (o *LD_a16Deref_SP) cycles() []uint8 { // 0x8
	return []uint8{20}
}

func (o *LD_a16Deref_SP) String() string { // 0x8
	return "LD" + " " + fmt.Sprintf("%X", o.operand1) + " " + "SP" /* operand2 */
}

func (o *LD_a16Deref_SP) SymbolicString() string { // 0x8
	return "LD (a16),SP"
}

type ADD_A_B struct {
	addr         string      // 0x80
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *ADD_A_B) Write(w io.Writer) (int, error) { // 0x80
	var b []byte

	b = append(b, 0x80)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x80,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *ADD_A_B) Length() uint8 { // 0x80
	return 1
}

func (o *ADD_A_B) cycles() []uint8 { // 0x80
	return []uint8{4}
}

func (o *ADD_A_B) String() string { // 0x80
	return "ADD" + " " + "A" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *ADD_A_B) SymbolicString() string { // 0x80
	return "ADD A,B"
}

type ADD_A_C struct {
	addr         string      // 0x81
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *ADD_A_C) Write(w io.Writer) (int, error) { // 0x81
	var b []byte

	b = append(b, 0x81)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x81,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *ADD_A_C) Length() uint8 { // 0x81
	return 1
}

func (o *ADD_A_C) cycles() []uint8 { // 0x81
	return []uint8{4}
}

func (o *ADD_A_C) String() string { // 0x81
	return "ADD" + " " + "A" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *ADD_A_C) SymbolicString() string { // 0x81
	return "ADD A,C"
}

type ADD_A_D struct {
	addr         string      // 0x82
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *ADD_A_D) Write(w io.Writer) (int, error) { // 0x82
	var b []byte

	b = append(b, 0x82)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x82,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *ADD_A_D) Length() uint8 { // 0x82
	return 1
}

func (o *ADD_A_D) cycles() []uint8 { // 0x82
	return []uint8{4}
}

func (o *ADD_A_D) String() string { // 0x82
	return "ADD" + " " + "A" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *ADD_A_D) SymbolicString() string { // 0x82
	return "ADD A,D"
}

type ADD_A_E struct {
	addr         string      // 0x83
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *ADD_A_E) Write(w io.Writer) (int, error) { // 0x83
	var b []byte

	b = append(b, 0x83)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x83,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *ADD_A_E) Length() uint8 { // 0x83
	return 1
}

func (o *ADD_A_E) cycles() []uint8 { // 0x83
	return []uint8{4}
}

func (o *ADD_A_E) String() string { // 0x83
	return "ADD" + " " + "A" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *ADD_A_E) SymbolicString() string { // 0x83
	return "ADD A,E"
}

type ADD_A_H struct {
	addr         string      // 0x84
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *ADD_A_H) Write(w io.Writer) (int, error) { // 0x84
	var b []byte

	b = append(b, 0x84)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x84,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *ADD_A_H) Length() uint8 { // 0x84
	return 1
}

func (o *ADD_A_H) cycles() []uint8 { // 0x84
	return []uint8{4}
}

func (o *ADD_A_H) String() string { // 0x84
	return "ADD" + " " + "A" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *ADD_A_H) SymbolicString() string { // 0x84
	return "ADD A,H"
}

type ADD_A_L struct {
	addr         string      // 0x85
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *ADD_A_L) Write(w io.Writer) (int, error) { // 0x85
	var b []byte

	b = append(b, 0x85)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x85,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *ADD_A_L) Length() uint8 { // 0x85
	return 1
}

func (o *ADD_A_L) cycles() []uint8 { // 0x85
	return []uint8{4}
}

func (o *ADD_A_L) String() string { // 0x85
	return "ADD" + " " + "A" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *ADD_A_L) SymbolicString() string { // 0x85
	return "ADD A,L"
}

type ADD_A_HLPtr struct {
	addr         string      // 0x86
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *ADD_A_HLPtr) Write(w io.Writer) (int, error) { // 0x86
	var b []byte

	b = append(b, 0x86)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x86,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *ADD_A_HLPtr) Length() uint8 { // 0x86
	return 1
}

func (o *ADD_A_HLPtr) cycles() []uint8 { // 0x86
	return []uint8{8}
}

func (o *ADD_A_HLPtr) String() string { // 0x86
	return "ADD" + " " + "A" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *ADD_A_HLPtr) SymbolicString() string { // 0x86
	return "ADD A,(HL)"
}

type ADD_A_A struct {
	addr         string      // 0x87
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *ADD_A_A) Write(w io.Writer) (int, error) { // 0x87
	var b []byte

	b = append(b, 0x87)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x87,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *ADD_A_A) Length() uint8 { // 0x87
	return 1
}

func (o *ADD_A_A) cycles() []uint8 { // 0x87
	return []uint8{4}
}

func (o *ADD_A_A) String() string { // 0x87
	return "ADD" + " " + "A" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *ADD_A_A) SymbolicString() string { // 0x87
	return "ADD A,A"
}

type ADC_A_B struct {
	addr         string      // 0x88
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *ADC_A_B) Write(w io.Writer) (int, error) { // 0x88
	var b []byte

	b = append(b, 0x88)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x88,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *ADC_A_B) Length() uint8 { // 0x88
	return 1
}

func (o *ADC_A_B) cycles() []uint8 { // 0x88
	return []uint8{4}
}

func (o *ADC_A_B) String() string { // 0x88
	return "ADC" + " " + "A" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *ADC_A_B) SymbolicString() string { // 0x88
	return "ADC A,B"
}

type ADC_A_C struct {
	addr         string      // 0x89
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *ADC_A_C) Write(w io.Writer) (int, error) { // 0x89
	var b []byte

	b = append(b, 0x89)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x89,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *ADC_A_C) Length() uint8 { // 0x89
	return 1
}

func (o *ADC_A_C) cycles() []uint8 { // 0x89
	return []uint8{4}
}

func (o *ADC_A_C) String() string { // 0x89
	return "ADC" + " " + "A" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *ADC_A_C) SymbolicString() string { // 0x89
	return "ADC A,C"
}

type ADC_A_D struct {
	addr         string      // 0x8a
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *ADC_A_D) Write(w io.Writer) (int, error) { // 0x8a
	var b []byte

	b = append(b, 0x8a)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x8a,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *ADC_A_D) Length() uint8 { // 0x8a
	return 1
}

func (o *ADC_A_D) cycles() []uint8 { // 0x8a
	return []uint8{4}
}

func (o *ADC_A_D) String() string { // 0x8a
	return "ADC" + " " + "A" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *ADC_A_D) SymbolicString() string { // 0x8a
	return "ADC A,D"
}

type ADC_A_E struct {
	addr         string      // 0x8b
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *ADC_A_E) Write(w io.Writer) (int, error) { // 0x8b
	var b []byte

	b = append(b, 0x8b)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x8b,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *ADC_A_E) Length() uint8 { // 0x8b
	return 1
}

func (o *ADC_A_E) cycles() []uint8 { // 0x8b
	return []uint8{4}
}

func (o *ADC_A_E) String() string { // 0x8b
	return "ADC" + " " + "A" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *ADC_A_E) SymbolicString() string { // 0x8b
	return "ADC A,E"
}

type ADC_A_H struct {
	addr         string      // 0x8c
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *ADC_A_H) Write(w io.Writer) (int, error) { // 0x8c
	var b []byte

	b = append(b, 0x8c)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x8c,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *ADC_A_H) Length() uint8 { // 0x8c
	return 1
}

func (o *ADC_A_H) cycles() []uint8 { // 0x8c
	return []uint8{4}
}

func (o *ADC_A_H) String() string { // 0x8c
	return "ADC" + " " + "A" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *ADC_A_H) SymbolicString() string { // 0x8c
	return "ADC A,H"
}

type ADC_A_L struct {
	addr         string      // 0x8d
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *ADC_A_L) Write(w io.Writer) (int, error) { // 0x8d
	var b []byte

	b = append(b, 0x8d)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x8d,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *ADC_A_L) Length() uint8 { // 0x8d
	return 1
}

func (o *ADC_A_L) cycles() []uint8 { // 0x8d
	return []uint8{4}
}

func (o *ADC_A_L) String() string { // 0x8d
	return "ADC" + " " + "A" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *ADC_A_L) SymbolicString() string { // 0x8d
	return "ADC A,L"
}

type ADC_A_HLPtr struct {
	addr         string      // 0x8e
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *ADC_A_HLPtr) Write(w io.Writer) (int, error) { // 0x8e
	var b []byte

	b = append(b, 0x8e)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x8e,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *ADC_A_HLPtr) Length() uint8 { // 0x8e
	return 1
}

func (o *ADC_A_HLPtr) cycles() []uint8 { // 0x8e
	return []uint8{8}
}

func (o *ADC_A_HLPtr) String() string { // 0x8e
	return "ADC" + " " + "A" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *ADC_A_HLPtr) SymbolicString() string { // 0x8e
	return "ADC A,(HL)"
}

type ADC_A_A struct {
	addr         string      // 0x8f
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *ADC_A_A) Write(w io.Writer) (int, error) { // 0x8f
	var b []byte

	b = append(b, 0x8f)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x8f,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *ADC_A_A) Length() uint8 { // 0x8f
	return 1
}

func (o *ADC_A_A) cycles() []uint8 { // 0x8f
	return []uint8{4}
}

func (o *ADC_A_A) String() string { // 0x8f
	return "ADC" + " " + "A" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *ADC_A_A) SymbolicString() string { // 0x8f
	return "ADC A,A"
}

type ADD_HL_BC struct {
	addr         string      // 0x9
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *ADD_HL_BC) Write(w io.Writer) (int, error) { // 0x9
	var b []byte

	b = append(b, 0x9)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x9,
	})
	if err != nil {
		return written, err
	}

	// HL is implicit in this instruction.
	// BC is implicit in this instruction.

	return written, err
}

func (o *ADD_HL_BC) Length() uint8 { // 0x9
	return 1
}

func (o *ADD_HL_BC) cycles() []uint8 { // 0x9
	return []uint8{8}
}

func (o *ADD_HL_BC) String() string { // 0x9
	return "ADD" + " " + "HL" /* operand1 */ + " " + "BC" /* operand2 */
}

func (o *ADD_HL_BC) SymbolicString() string { // 0x9
	return "ADD HL,BC"
}

type SUB_B struct {
	addr         string      // 0x90
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SUB_B) Write(w io.Writer) (int, error) { // 0x90
	var b []byte

	b = append(b, 0x90)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x90,
	})
	if err != nil {
		return written, err
	}

	// B is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SUB_B) Length() uint8 { // 0x90
	return 1
}

func (o *SUB_B) cycles() []uint8 { // 0x90
	return []uint8{4}
}

func (o *SUB_B) String() string { // 0x90
	return "SUB" + " " + "B" /* operand1 */
}

func (o *SUB_B) SymbolicString() string { // 0x90
	return "SUB B"
}

type SUB_C struct {
	addr         string      // 0x91
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SUB_C) Write(w io.Writer) (int, error) { // 0x91
	var b []byte

	b = append(b, 0x91)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x91,
	})
	if err != nil {
		return written, err
	}

	// C is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SUB_C) Length() uint8 { // 0x91
	return 1
}

func (o *SUB_C) cycles() []uint8 { // 0x91
	return []uint8{4}
}

func (o *SUB_C) String() string { // 0x91
	return "SUB" + " " + "C" /* operand1 */
}

func (o *SUB_C) SymbolicString() string { // 0x91
	return "SUB C"
}

type SUB_D struct {
	addr         string      // 0x92
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SUB_D) Write(w io.Writer) (int, error) { // 0x92
	var b []byte

	b = append(b, 0x92)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x92,
	})
	if err != nil {
		return written, err
	}

	// D is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SUB_D) Length() uint8 { // 0x92
	return 1
}

func (o *SUB_D) cycles() []uint8 { // 0x92
	return []uint8{4}
}

func (o *SUB_D) String() string { // 0x92
	return "SUB" + " " + "D" /* operand1 */
}

func (o *SUB_D) SymbolicString() string { // 0x92
	return "SUB D"
}

type SUB_E struct {
	addr         string      // 0x93
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SUB_E) Write(w io.Writer) (int, error) { // 0x93
	var b []byte

	b = append(b, 0x93)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x93,
	})
	if err != nil {
		return written, err
	}

	// E is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SUB_E) Length() uint8 { // 0x93
	return 1
}

func (o *SUB_E) cycles() []uint8 { // 0x93
	return []uint8{4}
}

func (o *SUB_E) String() string { // 0x93
	return "SUB" + " " + "E" /* operand1 */
}

func (o *SUB_E) SymbolicString() string { // 0x93
	return "SUB E"
}

type SUB_H struct {
	addr         string      // 0x94
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SUB_H) Write(w io.Writer) (int, error) { // 0x94
	var b []byte

	b = append(b, 0x94)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x94,
	})
	if err != nil {
		return written, err
	}

	// H is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SUB_H) Length() uint8 { // 0x94
	return 1
}

func (o *SUB_H) cycles() []uint8 { // 0x94
	return []uint8{4}
}

func (o *SUB_H) String() string { // 0x94
	return "SUB" + " " + "H" /* operand1 */
}

func (o *SUB_H) SymbolicString() string { // 0x94
	return "SUB H"
}

type SUB_L struct {
	addr         string      // 0x95
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SUB_L) Write(w io.Writer) (int, error) { // 0x95
	var b []byte

	b = append(b, 0x95)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x95,
	})
	if err != nil {
		return written, err
	}

	// L is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SUB_L) Length() uint8 { // 0x95
	return 1
}

func (o *SUB_L) cycles() []uint8 { // 0x95
	return []uint8{4}
}

func (o *SUB_L) String() string { // 0x95
	return "SUB" + " " + "L" /* operand1 */
}

func (o *SUB_L) SymbolicString() string { // 0x95
	return "SUB L"
}

type SUB_HLPtr struct {
	addr         string      // 0x96
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SUB_HLPtr) Write(w io.Writer) (int, error) { // 0x96
	var b []byte

	b = append(b, 0x96)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x96,
	})
	if err != nil {
		return written, err
	}

	// (HL) is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SUB_HLPtr) Length() uint8 { // 0x96
	return 1
}

func (o *SUB_HLPtr) cycles() []uint8 { // 0x96
	return []uint8{8}
}

func (o *SUB_HLPtr) String() string { // 0x96
	return "SUB" + " " + "(HL)" /* operand1 */
}

func (o *SUB_HLPtr) SymbolicString() string { // 0x96
	return "SUB (HL)"
}

type SUB_A struct {
	addr         string      // 0x97
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SUB_A) Write(w io.Writer) (int, error) { // 0x97
	var b []byte

	b = append(b, 0x97)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x97,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SUB_A) Length() uint8 { // 0x97
	return 1
}

func (o *SUB_A) cycles() []uint8 { // 0x97
	return []uint8{4}
}

func (o *SUB_A) String() string { // 0x97
	return "SUB" + " " + "A" /* operand1 */
}

func (o *SUB_A) SymbolicString() string { // 0x97
	return "SUB A"
}

type SBC_A_B struct {
	addr         string      // 0x98
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SBC_A_B) Write(w io.Writer) (int, error) { // 0x98
	var b []byte

	b = append(b, 0x98)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x98,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *SBC_A_B) Length() uint8 { // 0x98
	return 1
}

func (o *SBC_A_B) cycles() []uint8 { // 0x98
	return []uint8{4}
}

func (o *SBC_A_B) String() string { // 0x98
	return "SBC" + " " + "A" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *SBC_A_B) SymbolicString() string { // 0x98
	return "SBC A,B"
}

type SBC_A_C struct {
	addr         string      // 0x99
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SBC_A_C) Write(w io.Writer) (int, error) { // 0x99
	var b []byte

	b = append(b, 0x99)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x99,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *SBC_A_C) Length() uint8 { // 0x99
	return 1
}

func (o *SBC_A_C) cycles() []uint8 { // 0x99
	return []uint8{4}
}

func (o *SBC_A_C) String() string { // 0x99
	return "SBC" + " " + "A" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *SBC_A_C) SymbolicString() string { // 0x99
	return "SBC A,C"
}

type SBC_A_D struct {
	addr         string      // 0x9a
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SBC_A_D) Write(w io.Writer) (int, error) { // 0x9a
	var b []byte

	b = append(b, 0x9a)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x9a,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *SBC_A_D) Length() uint8 { // 0x9a
	return 1
}

func (o *SBC_A_D) cycles() []uint8 { // 0x9a
	return []uint8{4}
}

func (o *SBC_A_D) String() string { // 0x9a
	return "SBC" + " " + "A" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *SBC_A_D) SymbolicString() string { // 0x9a
	return "SBC A,D"
}

type SBC_A_E struct {
	addr         string      // 0x9b
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SBC_A_E) Write(w io.Writer) (int, error) { // 0x9b
	var b []byte

	b = append(b, 0x9b)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x9b,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *SBC_A_E) Length() uint8 { // 0x9b
	return 1
}

func (o *SBC_A_E) cycles() []uint8 { // 0x9b
	return []uint8{4}
}

func (o *SBC_A_E) String() string { // 0x9b
	return "SBC" + " " + "A" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *SBC_A_E) SymbolicString() string { // 0x9b
	return "SBC A,E"
}

type SBC_A_H struct {
	addr         string      // 0x9c
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SBC_A_H) Write(w io.Writer) (int, error) { // 0x9c
	var b []byte

	b = append(b, 0x9c)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x9c,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *SBC_A_H) Length() uint8 { // 0x9c
	return 1
}

func (o *SBC_A_H) cycles() []uint8 { // 0x9c
	return []uint8{4}
}

func (o *SBC_A_H) String() string { // 0x9c
	return "SBC" + " " + "A" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *SBC_A_H) SymbolicString() string { // 0x9c
	return "SBC A,H"
}

type SBC_A_L struct {
	addr         string      // 0x9d
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SBC_A_L) Write(w io.Writer) (int, error) { // 0x9d
	var b []byte

	b = append(b, 0x9d)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x9d,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *SBC_A_L) Length() uint8 { // 0x9d
	return 1
}

func (o *SBC_A_L) cycles() []uint8 { // 0x9d
	return []uint8{4}
}

func (o *SBC_A_L) String() string { // 0x9d
	return "SBC" + " " + "A" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *SBC_A_L) SymbolicString() string { // 0x9d
	return "SBC A,L"
}

type SBC_A_HLPtr struct {
	addr         string      // 0x9e
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SBC_A_HLPtr) Write(w io.Writer) (int, error) { // 0x9e
	var b []byte

	b = append(b, 0x9e)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x9e,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *SBC_A_HLPtr) Length() uint8 { // 0x9e
	return 1
}

func (o *SBC_A_HLPtr) cycles() []uint8 { // 0x9e
	return []uint8{8}
}

func (o *SBC_A_HLPtr) String() string { // 0x9e
	return "SBC" + " " + "A" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *SBC_A_HLPtr) SymbolicString() string { // 0x9e
	return "SBC A,(HL)"
}

type SBC_A_A struct {
	addr         string      // 0x9f
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SBC_A_A) Write(w io.Writer) (int, error) { // 0x9f
	var b []byte

	b = append(b, 0x9f)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0x9f,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *SBC_A_A) Length() uint8 { // 0x9f
	return 1
}

func (o *SBC_A_A) cycles() []uint8 { // 0x9f
	return []uint8{4}
}

func (o *SBC_A_A) String() string { // 0x9f
	return "SBC" + " " + "A" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *SBC_A_A) SymbolicString() string { // 0x9f
	return "SBC A,A"
}

type LD_A_BCDeref struct {
	addr         string      // 0xa
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_A_BCDeref) Write(w io.Writer) (int, error) { // 0xa
	var b []byte

	b = append(b, 0xa)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xa,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// (BC) is implicit in this instruction.

	return written, err
}

func (o *LD_A_BCDeref) Length() uint8 { // 0xa
	return 1
}

func (o *LD_A_BCDeref) cycles() []uint8 { // 0xa
	return []uint8{8}
}

func (o *LD_A_BCDeref) String() string { // 0xa
	return "LD" + " " + "A" /* operand1 */ + " " + "(BC)" /* operand2 */
}

func (o *LD_A_BCDeref) SymbolicString() string { // 0xa
	return "LD A,(BC)"
}

type AND_B struct {
	addr         string      // 0xa0
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *AND_B) Write(w io.Writer) (int, error) { // 0xa0
	var b []byte

	b = append(b, 0xa0)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xa0,
	})
	if err != nil {
		return written, err
	}

	// B is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *AND_B) Length() uint8 { // 0xa0
	return 1
}

func (o *AND_B) cycles() []uint8 { // 0xa0
	return []uint8{4}
}

func (o *AND_B) String() string { // 0xa0
	return "AND" + " " + "B" /* operand1 */
}

func (o *AND_B) SymbolicString() string { // 0xa0
	return "AND B"
}

type AND_C struct {
	addr         string      // 0xa1
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *AND_C) Write(w io.Writer) (int, error) { // 0xa1
	var b []byte

	b = append(b, 0xa1)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xa1,
	})
	if err != nil {
		return written, err
	}

	// C is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *AND_C) Length() uint8 { // 0xa1
	return 1
}

func (o *AND_C) cycles() []uint8 { // 0xa1
	return []uint8{4}
}

func (o *AND_C) String() string { // 0xa1
	return "AND" + " " + "C" /* operand1 */
}

func (o *AND_C) SymbolicString() string { // 0xa1
	return "AND C"
}

type AND_D struct {
	addr         string      // 0xa2
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *AND_D) Write(w io.Writer) (int, error) { // 0xa2
	var b []byte

	b = append(b, 0xa2)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xa2,
	})
	if err != nil {
		return written, err
	}

	// D is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *AND_D) Length() uint8 { // 0xa2
	return 1
}

func (o *AND_D) cycles() []uint8 { // 0xa2
	return []uint8{4}
}

func (o *AND_D) String() string { // 0xa2
	return "AND" + " " + "D" /* operand1 */
}

func (o *AND_D) SymbolicString() string { // 0xa2
	return "AND D"
}

type AND_E struct {
	addr         string      // 0xa3
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *AND_E) Write(w io.Writer) (int, error) { // 0xa3
	var b []byte

	b = append(b, 0xa3)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xa3,
	})
	if err != nil {
		return written, err
	}

	// E is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *AND_E) Length() uint8 { // 0xa3
	return 1
}

func (o *AND_E) cycles() []uint8 { // 0xa3
	return []uint8{4}
}

func (o *AND_E) String() string { // 0xa3
	return "AND" + " " + "E" /* operand1 */
}

func (o *AND_E) SymbolicString() string { // 0xa3
	return "AND E"
}

type AND_H struct {
	addr         string      // 0xa4
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *AND_H) Write(w io.Writer) (int, error) { // 0xa4
	var b []byte

	b = append(b, 0xa4)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xa4,
	})
	if err != nil {
		return written, err
	}

	// H is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *AND_H) Length() uint8 { // 0xa4
	return 1
}

func (o *AND_H) cycles() []uint8 { // 0xa4
	return []uint8{4}
}

func (o *AND_H) String() string { // 0xa4
	return "AND" + " " + "H" /* operand1 */
}

func (o *AND_H) SymbolicString() string { // 0xa4
	return "AND H"
}

type AND_L struct {
	addr         string      // 0xa5
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *AND_L) Write(w io.Writer) (int, error) { // 0xa5
	var b []byte

	b = append(b, 0xa5)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xa5,
	})
	if err != nil {
		return written, err
	}

	// L is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *AND_L) Length() uint8 { // 0xa5
	return 1
}

func (o *AND_L) cycles() []uint8 { // 0xa5
	return []uint8{4}
}

func (o *AND_L) String() string { // 0xa5
	return "AND" + " " + "L" /* operand1 */
}

func (o *AND_L) SymbolicString() string { // 0xa5
	return "AND L"
}

type AND_HLPtr struct {
	addr         string      // 0xa6
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *AND_HLPtr) Write(w io.Writer) (int, error) { // 0xa6
	var b []byte

	b = append(b, 0xa6)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xa6,
	})
	if err != nil {
		return written, err
	}

	// (HL) is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *AND_HLPtr) Length() uint8 { // 0xa6
	return 1
}

func (o *AND_HLPtr) cycles() []uint8 { // 0xa6
	return []uint8{8}
}

func (o *AND_HLPtr) String() string { // 0xa6
	return "AND" + " " + "(HL)" /* operand1 */
}

func (o *AND_HLPtr) SymbolicString() string { // 0xa6
	return "AND (HL)"
}

type AND_A struct {
	addr         string      // 0xa7
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *AND_A) Write(w io.Writer) (int, error) { // 0xa7
	var b []byte

	b = append(b, 0xa7)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xa7,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *AND_A) Length() uint8 { // 0xa7
	return 1
}

func (o *AND_A) cycles() []uint8 { // 0xa7
	return []uint8{4}
}

func (o *AND_A) String() string { // 0xa7
	return "AND" + " " + "A" /* operand1 */
}

func (o *AND_A) SymbolicString() string { // 0xa7
	return "AND A"
}

type XOR_B struct {
	addr         string      // 0xa8
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *XOR_B) Write(w io.Writer) (int, error) { // 0xa8
	var b []byte

	b = append(b, 0xa8)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xa8,
	})
	if err != nil {
		return written, err
	}

	// B is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *XOR_B) Length() uint8 { // 0xa8
	return 1
}

func (o *XOR_B) cycles() []uint8 { // 0xa8
	return []uint8{4}
}

func (o *XOR_B) String() string { // 0xa8
	return "XOR" + " " + "B" /* operand1 */
}

func (o *XOR_B) SymbolicString() string { // 0xa8
	return "XOR B"
}

type XOR_C struct {
	addr         string      // 0xa9
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *XOR_C) Write(w io.Writer) (int, error) { // 0xa9
	var b []byte

	b = append(b, 0xa9)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xa9,
	})
	if err != nil {
		return written, err
	}

	// C is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *XOR_C) Length() uint8 { // 0xa9
	return 1
}

func (o *XOR_C) cycles() []uint8 { // 0xa9
	return []uint8{4}
}

func (o *XOR_C) String() string { // 0xa9
	return "XOR" + " " + "C" /* operand1 */
}

func (o *XOR_C) SymbolicString() string { // 0xa9
	return "XOR C"
}

type XOR_D struct {
	addr         string      // 0xaa
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *XOR_D) Write(w io.Writer) (int, error) { // 0xaa
	var b []byte

	b = append(b, 0xaa)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xaa,
	})
	if err != nil {
		return written, err
	}

	// D is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *XOR_D) Length() uint8 { // 0xaa
	return 1
}

func (o *XOR_D) cycles() []uint8 { // 0xaa
	return []uint8{4}
}

func (o *XOR_D) String() string { // 0xaa
	return "XOR" + " " + "D" /* operand1 */
}

func (o *XOR_D) SymbolicString() string { // 0xaa
	return "XOR D"
}

type XOR_E struct {
	addr         string      // 0xab
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *XOR_E) Write(w io.Writer) (int, error) { // 0xab
	var b []byte

	b = append(b, 0xab)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xab,
	})
	if err != nil {
		return written, err
	}

	// E is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *XOR_E) Length() uint8 { // 0xab
	return 1
}

func (o *XOR_E) cycles() []uint8 { // 0xab
	return []uint8{4}
}

func (o *XOR_E) String() string { // 0xab
	return "XOR" + " " + "E" /* operand1 */
}

func (o *XOR_E) SymbolicString() string { // 0xab
	return "XOR E"
}

type XOR_H struct {
	addr         string      // 0xac
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *XOR_H) Write(w io.Writer) (int, error) { // 0xac
	var b []byte

	b = append(b, 0xac)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xac,
	})
	if err != nil {
		return written, err
	}

	// H is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *XOR_H) Length() uint8 { // 0xac
	return 1
}

func (o *XOR_H) cycles() []uint8 { // 0xac
	return []uint8{4}
}

func (o *XOR_H) String() string { // 0xac
	return "XOR" + " " + "H" /* operand1 */
}

func (o *XOR_H) SymbolicString() string { // 0xac
	return "XOR H"
}

type XOR_L struct {
	addr         string      // 0xad
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *XOR_L) Write(w io.Writer) (int, error) { // 0xad
	var b []byte

	b = append(b, 0xad)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xad,
	})
	if err != nil {
		return written, err
	}

	// L is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *XOR_L) Length() uint8 { // 0xad
	return 1
}

func (o *XOR_L) cycles() []uint8 { // 0xad
	return []uint8{4}
}

func (o *XOR_L) String() string { // 0xad
	return "XOR" + " " + "L" /* operand1 */
}

func (o *XOR_L) SymbolicString() string { // 0xad
	return "XOR L"
}

type XOR_HLPtr struct {
	addr         string      // 0xae
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *XOR_HLPtr) Write(w io.Writer) (int, error) { // 0xae
	var b []byte

	b = append(b, 0xae)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xae,
	})
	if err != nil {
		return written, err
	}

	// (HL) is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *XOR_HLPtr) Length() uint8 { // 0xae
	return 1
}

func (o *XOR_HLPtr) cycles() []uint8 { // 0xae
	return []uint8{8}
}

func (o *XOR_HLPtr) String() string { // 0xae
	return "XOR" + " " + "(HL)" /* operand1 */
}

func (o *XOR_HLPtr) SymbolicString() string { // 0xae
	return "XOR (HL)"
}

type XOR_A struct {
	addr         string      // 0xaf
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *XOR_A) Write(w io.Writer) (int, error) { // 0xaf
	var b []byte

	b = append(b, 0xaf)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xaf,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *XOR_A) Length() uint8 { // 0xaf
	return 1
}

func (o *XOR_A) cycles() []uint8 { // 0xaf
	return []uint8{4}
}

func (o *XOR_A) String() string { // 0xaf
	return "XOR" + " " + "A" /* operand1 */
}

func (o *XOR_A) SymbolicString() string { // 0xaf
	return "XOR A"
}

type DEC_BC struct {
	addr         string      // 0xb
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *DEC_BC) Write(w io.Writer) (int, error) { // 0xb
	var b []byte

	b = append(b, 0xb)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xb,
	})
	if err != nil {
		return written, err
	}

	// BC is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *DEC_BC) Length() uint8 { // 0xb
	return 1
}

func (o *DEC_BC) cycles() []uint8 { // 0xb
	return []uint8{8}
}

func (o *DEC_BC) String() string { // 0xb
	return "DEC" + " " + "BC" /* operand1 */
}

func (o *DEC_BC) SymbolicString() string { // 0xb
	return "DEC BC"
}

type OR_B struct {
	addr         string      // 0xb0
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *OR_B) Write(w io.Writer) (int, error) { // 0xb0
	var b []byte

	b = append(b, 0xb0)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xb0,
	})
	if err != nil {
		return written, err
	}

	// B is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *OR_B) Length() uint8 { // 0xb0
	return 1
}

func (o *OR_B) cycles() []uint8 { // 0xb0
	return []uint8{4}
}

func (o *OR_B) String() string { // 0xb0
	return "OR" + " " + "B" /* operand1 */
}

func (o *OR_B) SymbolicString() string { // 0xb0
	return "OR B"
}

type OR_C struct {
	addr         string      // 0xb1
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *OR_C) Write(w io.Writer) (int, error) { // 0xb1
	var b []byte

	b = append(b, 0xb1)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xb1,
	})
	if err != nil {
		return written, err
	}

	// C is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *OR_C) Length() uint8 { // 0xb1
	return 1
}

func (o *OR_C) cycles() []uint8 { // 0xb1
	return []uint8{4}
}

func (o *OR_C) String() string { // 0xb1
	return "OR" + " " + "C" /* operand1 */
}

func (o *OR_C) SymbolicString() string { // 0xb1
	return "OR C"
}

type OR_D struct {
	addr         string      // 0xb2
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *OR_D) Write(w io.Writer) (int, error) { // 0xb2
	var b []byte

	b = append(b, 0xb2)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xb2,
	})
	if err != nil {
		return written, err
	}

	// D is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *OR_D) Length() uint8 { // 0xb2
	return 1
}

func (o *OR_D) cycles() []uint8 { // 0xb2
	return []uint8{4}
}

func (o *OR_D) String() string { // 0xb2
	return "OR" + " " + "D" /* operand1 */
}

func (o *OR_D) SymbolicString() string { // 0xb2
	return "OR D"
}

type OR_E struct {
	addr         string      // 0xb3
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *OR_E) Write(w io.Writer) (int, error) { // 0xb3
	var b []byte

	b = append(b, 0xb3)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xb3,
	})
	if err != nil {
		return written, err
	}

	// E is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *OR_E) Length() uint8 { // 0xb3
	return 1
}

func (o *OR_E) cycles() []uint8 { // 0xb3
	return []uint8{4}
}

func (o *OR_E) String() string { // 0xb3
	return "OR" + " " + "E" /* operand1 */
}

func (o *OR_E) SymbolicString() string { // 0xb3
	return "OR E"
}

type OR_H struct {
	addr         string      // 0xb4
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *OR_H) Write(w io.Writer) (int, error) { // 0xb4
	var b []byte

	b = append(b, 0xb4)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xb4,
	})
	if err != nil {
		return written, err
	}

	// H is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *OR_H) Length() uint8 { // 0xb4
	return 1
}

func (o *OR_H) cycles() []uint8 { // 0xb4
	return []uint8{4}
}

func (o *OR_H) String() string { // 0xb4
	return "OR" + " " + "H" /* operand1 */
}

func (o *OR_H) SymbolicString() string { // 0xb4
	return "OR H"
}

type OR_L struct {
	addr         string      // 0xb5
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *OR_L) Write(w io.Writer) (int, error) { // 0xb5
	var b []byte

	b = append(b, 0xb5)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xb5,
	})
	if err != nil {
		return written, err
	}

	// L is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *OR_L) Length() uint8 { // 0xb5
	return 1
}

func (o *OR_L) cycles() []uint8 { // 0xb5
	return []uint8{4}
}

func (o *OR_L) String() string { // 0xb5
	return "OR" + " " + "L" /* operand1 */
}

func (o *OR_L) SymbolicString() string { // 0xb5
	return "OR L"
}

type OR_HLPtr struct {
	addr         string      // 0xb6
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *OR_HLPtr) Write(w io.Writer) (int, error) { // 0xb6
	var b []byte

	b = append(b, 0xb6)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xb6,
	})
	if err != nil {
		return written, err
	}

	// (HL) is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *OR_HLPtr) Length() uint8 { // 0xb6
	return 1
}

func (o *OR_HLPtr) cycles() []uint8 { // 0xb6
	return []uint8{8}
}

func (o *OR_HLPtr) String() string { // 0xb6
	return "OR" + " " + "(HL)" /* operand1 */
}

func (o *OR_HLPtr) SymbolicString() string { // 0xb6
	return "OR (HL)"
}

type OR_A struct {
	addr         string      // 0xb7
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *OR_A) Write(w io.Writer) (int, error) { // 0xb7
	var b []byte

	b = append(b, 0xb7)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xb7,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *OR_A) Length() uint8 { // 0xb7
	return 1
}

func (o *OR_A) cycles() []uint8 { // 0xb7
	return []uint8{4}
}

func (o *OR_A) String() string { // 0xb7
	return "OR" + " " + "A" /* operand1 */
}

func (o *OR_A) SymbolicString() string { // 0xb7
	return "OR A"
}

type CP_B struct {
	addr         string      // 0xb8
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *CP_B) Write(w io.Writer) (int, error) { // 0xb8
	var b []byte

	b = append(b, 0xb8)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xb8,
	})
	if err != nil {
		return written, err
	}

	// B is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *CP_B) Length() uint8 { // 0xb8
	return 1
}

func (o *CP_B) cycles() []uint8 { // 0xb8
	return []uint8{4}
}

func (o *CP_B) String() string { // 0xb8
	return "CP" + " " + "B" /* operand1 */
}

func (o *CP_B) SymbolicString() string { // 0xb8
	return "CP B"
}

type CP_C struct {
	addr         string      // 0xb9
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *CP_C) Write(w io.Writer) (int, error) { // 0xb9
	var b []byte

	b = append(b, 0xb9)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xb9,
	})
	if err != nil {
		return written, err
	}

	// C is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *CP_C) Length() uint8 { // 0xb9
	return 1
}

func (o *CP_C) cycles() []uint8 { // 0xb9
	return []uint8{4}
}

func (o *CP_C) String() string { // 0xb9
	return "CP" + " " + "C" /* operand1 */
}

func (o *CP_C) SymbolicString() string { // 0xb9
	return "CP C"
}

type CP_D struct {
	addr         string      // 0xba
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *CP_D) Write(w io.Writer) (int, error) { // 0xba
	var b []byte

	b = append(b, 0xba)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xba,
	})
	if err != nil {
		return written, err
	}

	// D is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *CP_D) Length() uint8 { // 0xba
	return 1
}

func (o *CP_D) cycles() []uint8 { // 0xba
	return []uint8{4}
}

func (o *CP_D) String() string { // 0xba
	return "CP" + " " + "D" /* operand1 */
}

func (o *CP_D) SymbolicString() string { // 0xba
	return "CP D"
}

type CP_E struct {
	addr         string      // 0xbb
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *CP_E) Write(w io.Writer) (int, error) { // 0xbb
	var b []byte

	b = append(b, 0xbb)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xbb,
	})
	if err != nil {
		return written, err
	}

	// E is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *CP_E) Length() uint8 { // 0xbb
	return 1
}

func (o *CP_E) cycles() []uint8 { // 0xbb
	return []uint8{4}
}

func (o *CP_E) String() string { // 0xbb
	return "CP" + " " + "E" /* operand1 */
}

func (o *CP_E) SymbolicString() string { // 0xbb
	return "CP E"
}

type CP_H struct {
	addr         string      // 0xbc
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *CP_H) Write(w io.Writer) (int, error) { // 0xbc
	var b []byte

	b = append(b, 0xbc)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xbc,
	})
	if err != nil {
		return written, err
	}

	// H is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *CP_H) Length() uint8 { // 0xbc
	return 1
}

func (o *CP_H) cycles() []uint8 { // 0xbc
	return []uint8{4}
}

func (o *CP_H) String() string { // 0xbc
	return "CP" + " " + "H" /* operand1 */
}

func (o *CP_H) SymbolicString() string { // 0xbc
	return "CP H"
}

type CP_L struct {
	addr         string      // 0xbd
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *CP_L) Write(w io.Writer) (int, error) { // 0xbd
	var b []byte

	b = append(b, 0xbd)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xbd,
	})
	if err != nil {
		return written, err
	}

	// L is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *CP_L) Length() uint8 { // 0xbd
	return 1
}

func (o *CP_L) cycles() []uint8 { // 0xbd
	return []uint8{4}
}

func (o *CP_L) String() string { // 0xbd
	return "CP" + " " + "L" /* operand1 */
}

func (o *CP_L) SymbolicString() string { // 0xbd
	return "CP L"
}

type CP_HLPtr struct {
	addr         string      // 0xbe
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *CP_HLPtr) Write(w io.Writer) (int, error) { // 0xbe
	var b []byte

	b = append(b, 0xbe)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xbe,
	})
	if err != nil {
		return written, err
	}

	// (HL) is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *CP_HLPtr) Length() uint8 { // 0xbe
	return 1
}

func (o *CP_HLPtr) cycles() []uint8 { // 0xbe
	return []uint8{8}
}

func (o *CP_HLPtr) String() string { // 0xbe
	return "CP" + " " + "(HL)" /* operand1 */
}

func (o *CP_HLPtr) SymbolicString() string { // 0xbe
	return "CP (HL)"
}

type CP_A struct {
	addr         string      // 0xbf
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *CP_A) Write(w io.Writer) (int, error) { // 0xbf
	var b []byte

	b = append(b, 0xbf)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xbf,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *CP_A) Length() uint8 { // 0xbf
	return 1
}

func (o *CP_A) cycles() []uint8 { // 0xbf
	return []uint8{4}
}

func (o *CP_A) String() string { // 0xbf
	return "CP" + " " + "A" /* operand1 */
}

func (o *CP_A) SymbolicString() string { // 0xbf
	return "CP A"
}

type INC_C struct {
	addr         string      // 0xc
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *INC_C) Write(w io.Writer) (int, error) { // 0xc
	var b []byte

	b = append(b, 0xc)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xc,
	})
	if err != nil {
		return written, err
	}

	// C is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *INC_C) Length() uint8 { // 0xc
	return 1
}

func (o *INC_C) cycles() []uint8 { // 0xc
	return []uint8{4}
}

func (o *INC_C) String() string { // 0xc
	return "INC" + " " + "C" /* operand1 */
}

func (o *INC_C) SymbolicString() string { // 0xc
	return "INC C"
}

type RET_NZ struct {
	addr         string      // 0xc0
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RET_NZ) Write(w io.Writer) (int, error) { // 0xc0
	var b []byte

	b = append(b, 0xc0)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xc0,
	})
	if err != nil {
		return written, err
	}

	// NZ is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RET_NZ) Length() uint8 { // 0xc0
	return 1
}

func (o *RET_NZ) cycles() []uint8 { // 0xc0
	return []uint8{20, 8}
}

func (o *RET_NZ) String() string { // 0xc0
	return "RET" + " " + "NZ" /* operand1 */
}

func (o *RET_NZ) SymbolicString() string { // 0xc0
	return "RET NZ"
}

type POP_BC struct {
	addr         string      // 0xc1
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *POP_BC) Write(w io.Writer) (int, error) { // 0xc1
	var b []byte

	b = append(b, 0xc1)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xc1,
	})
	if err != nil {
		return written, err
	}

	// BC is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *POP_BC) Length() uint8 { // 0xc1
	return 1
}

func (o *POP_BC) cycles() []uint8 { // 0xc1
	return []uint8{12}
}

func (o *POP_BC) String() string { // 0xc1
	return "POP" + " " + "BC" /* operand1 */
}

func (o *POP_BC) SymbolicString() string { // 0xc1
	return "POP BC"
}

type JP_NZ_a16 struct {
	addr         string      // 0xc2
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *JP_NZ_a16) Write(w io.Writer) (int, error) { // 0xc2
	var b []byte

	b = append(b, 0xc2)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xc2,
	})
	if err != nil {
		return written, err
	}

	// NZ is implicit in this instruction.
	err = writeImmediate16BitAddress(w, o.operand2)
	written += 2

	return written, err
}

func (o *JP_NZ_a16) Length() uint8 { // 0xc2
	return 3
}

func (o *JP_NZ_a16) cycles() []uint8 { // 0xc2
	return []uint8{16, 12}
}

func (o *JP_NZ_a16) String() string { // 0xc2
	return "JP" + " " + "NZ" /* operand1 */ + " " + fmt.Sprintf("%X", o.operand2)
}

func (o *JP_NZ_a16) SymbolicString() string { // 0xc2
	return "JP NZ,a16"
}

type JP_a16 struct {
	addr         string      // 0xc3
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *JP_a16) Write(w io.Writer) (int, error) { // 0xc3
	var b []byte

	b = append(b, 0xc3)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xc3,
	})
	if err != nil {
		return written, err
	}

	err = writeImmediate16BitAddress(w, o.operand1)
	written += 2
	//  is implicit in this instruction.

	return written, err
}

func (o *JP_a16) Length() uint8 { // 0xc3
	return 3
}

func (o *JP_a16) cycles() []uint8 { // 0xc3
	return []uint8{16}
}

func (o *JP_a16) String() string { // 0xc3
	return "JP" + " " + fmt.Sprintf("%X", o.operand1)
}

func (o *JP_a16) SymbolicString() string { // 0xc3
	return "JP a16"
}

type CALL_NZ_a16 struct {
	addr         string      // 0xc4
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *CALL_NZ_a16) Write(w io.Writer) (int, error) { // 0xc4
	var b []byte

	b = append(b, 0xc4)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xc4,
	})
	if err != nil {
		return written, err
	}

	// NZ is implicit in this instruction.
	err = writeImmediate16BitAddress(w, o.operand2)
	written += 2

	return written, err
}

func (o *CALL_NZ_a16) Length() uint8 { // 0xc4
	return 3
}

func (o *CALL_NZ_a16) cycles() []uint8 { // 0xc4
	return []uint8{24, 12}
}

func (o *CALL_NZ_a16) String() string { // 0xc4
	return "CALL" + " " + "NZ" /* operand1 */ + " " + fmt.Sprintf("%X", o.operand2)
}

func (o *CALL_NZ_a16) SymbolicString() string { // 0xc4
	return "CALL NZ,a16"
}

type PUSH_BC struct {
	addr         string      // 0xc5
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *PUSH_BC) Write(w io.Writer) (int, error) { // 0xc5
	var b []byte

	b = append(b, 0xc5)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xc5,
	})
	if err != nil {
		return written, err
	}

	// BC is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *PUSH_BC) Length() uint8 { // 0xc5
	return 1
}

func (o *PUSH_BC) cycles() []uint8 { // 0xc5
	return []uint8{16}
}

func (o *PUSH_BC) String() string { // 0xc5
	return "PUSH" + " " + "BC" /* operand1 */
}

func (o *PUSH_BC) SymbolicString() string { // 0xc5
	return "PUSH BC"
}

type ADD_A_d8 struct {
	addr         string      // 0xc6
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *ADD_A_d8) Write(w io.Writer) (int, error) { // 0xc6
	var b []byte

	b = append(b, 0xc6)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xc6,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	err = writeImmediate8BitData(w, o.operand2)
	written += 1

	return written, err
}

func (o *ADD_A_d8) Length() uint8 { // 0xc6
	return 2
}

func (o *ADD_A_d8) cycles() []uint8 { // 0xc6
	return []uint8{8}
}

func (o *ADD_A_d8) String() string { // 0xc6
	return "ADD" + " " + "A" /* operand1 */ + " " + fmt.Sprintf("%X", o.operand2)
}

func (o *ADD_A_d8) SymbolicString() string { // 0xc6
	return "ADD A,d8"
}

type RST_00H struct {
	addr         string      // 0xc7
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RST_00H) Write(w io.Writer) (int, error) { // 0xc7
	var b []byte

	b = append(b, 0xc7)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xc7,
	})
	if err != nil {
		return written, err
	}

	// 00H is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RST_00H) Length() uint8 { // 0xc7
	return 1
}

func (o *RST_00H) cycles() []uint8 { // 0xc7
	return []uint8{16}
}

func (o *RST_00H) String() string { // 0xc7
	return "RST" + " " + "00H" /* operand1 */
}

func (o *RST_00H) SymbolicString() string { // 0xc7
	return "RST 00H"
}

type RET_Z struct {
	addr         string      // 0xc8
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RET_Z) Write(w io.Writer) (int, error) { // 0xc8
	var b []byte

	b = append(b, 0xc8)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xc8,
	})
	if err != nil {
		return written, err
	}

	// Z is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RET_Z) Length() uint8 { // 0xc8
	return 1
}

func (o *RET_Z) cycles() []uint8 { // 0xc8
	return []uint8{20, 8}
}

func (o *RET_Z) String() string { // 0xc8
	return "RET" + " " + "Z" /* operand1 */
}

func (o *RET_Z) SymbolicString() string { // 0xc8
	return "RET Z"
}

type RET struct {
	addr         string      // 0xc9
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RET) Write(w io.Writer) (int, error) { // 0xc9
	var b []byte

	b = append(b, 0xc9)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xc9,
	})
	if err != nil {
		return written, err
	}

	//  is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RET) Length() uint8 { // 0xc9
	return 1
}

func (o *RET) cycles() []uint8 { // 0xc9
	return []uint8{16}
}

func (o *RET) String() string { // 0xc9
	return "RET"
}

func (o *RET) SymbolicString() string { // 0xc9
	return "RET"
}

type JP_Z_a16 struct {
	addr         string      // 0xca
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *JP_Z_a16) Write(w io.Writer) (int, error) { // 0xca
	var b []byte

	b = append(b, 0xca)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xca,
	})
	if err != nil {
		return written, err
	}

	// Z is implicit in this instruction.
	err = writeImmediate16BitAddress(w, o.operand2)
	written += 2

	return written, err
}

func (o *JP_Z_a16) Length() uint8 { // 0xca
	return 3
}

func (o *JP_Z_a16) cycles() []uint8 { // 0xca
	return []uint8{16, 12}
}

func (o *JP_Z_a16) String() string { // 0xca
	return "JP" + " " + "Z" /* operand1 */ + " " + fmt.Sprintf("%X", o.operand2)
}

func (o *JP_Z_a16) SymbolicString() string { // 0xca
	return "JP Z,a16"
}

type PREFIX_CB struct {
	addr         string      // 0xcb
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *PREFIX_CB) Write(w io.Writer) (int, error) { // 0xcb
	var b []byte

	b = append(b, 0xcb)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xcb,
	})
	if err != nil {
		return written, err
	}

	// CB is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *PREFIX_CB) Length() uint8 { // 0xcb
	return 1
}

func (o *PREFIX_CB) cycles() []uint8 { // 0xcb
	return []uint8{4}
}

func (o *PREFIX_CB) String() string { // 0xcb
	return "PREFIX" + " " + "CB" /* operand1 */
}

func (o *PREFIX_CB) SymbolicString() string { // 0xcb
	return "PREFIX CB"
}

type CALL_Z_a16 struct {
	addr         string      // 0xcc
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *CALL_Z_a16) Write(w io.Writer) (int, error) { // 0xcc
	var b []byte

	b = append(b, 0xcc)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xcc,
	})
	if err != nil {
		return written, err
	}

	// Z is implicit in this instruction.
	err = writeImmediate16BitAddress(w, o.operand2)
	written += 2

	return written, err
}

func (o *CALL_Z_a16) Length() uint8 { // 0xcc
	return 3
}

func (o *CALL_Z_a16) cycles() []uint8 { // 0xcc
	return []uint8{24, 12}
}

func (o *CALL_Z_a16) String() string { // 0xcc
	return "CALL" + " " + "Z" /* operand1 */ + " " + fmt.Sprintf("%X", o.operand2)
}

func (o *CALL_Z_a16) SymbolicString() string { // 0xcc
	return "CALL Z,a16"
}

type CALL_a16 struct {
	addr         string      // 0xcd
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *CALL_a16) Write(w io.Writer) (int, error) { // 0xcd
	var b []byte

	b = append(b, 0xcd)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xcd,
	})
	if err != nil {
		return written, err
	}

	err = writeImmediate16BitAddress(w, o.operand1)
	written += 2
	//  is implicit in this instruction.

	return written, err
}

func (o *CALL_a16) Length() uint8 { // 0xcd
	return 3
}

func (o *CALL_a16) cycles() []uint8 { // 0xcd
	return []uint8{24}
}

func (o *CALL_a16) String() string { // 0xcd
	return "CALL" + " " + fmt.Sprintf("%X", o.operand1)
}

func (o *CALL_a16) SymbolicString() string { // 0xcd
	return "CALL a16"
}

type ADC_A_d8 struct {
	addr         string      // 0xce
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *ADC_A_d8) Write(w io.Writer) (int, error) { // 0xce
	var b []byte

	b = append(b, 0xce)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xce,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	err = writeImmediate8BitData(w, o.operand2)
	written += 1

	return written, err
}

func (o *ADC_A_d8) Length() uint8 { // 0xce
	return 2
}

func (o *ADC_A_d8) cycles() []uint8 { // 0xce
	return []uint8{8}
}

func (o *ADC_A_d8) String() string { // 0xce
	return "ADC" + " " + "A" /* operand1 */ + " " + fmt.Sprintf("%X", o.operand2)
}

func (o *ADC_A_d8) SymbolicString() string { // 0xce
	return "ADC A,d8"
}

type RST_08H struct {
	addr         string      // 0xcf
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RST_08H) Write(w io.Writer) (int, error) { // 0xcf
	var b []byte

	b = append(b, 0xcf)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xcf,
	})
	if err != nil {
		return written, err
	}

	// 08H is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RST_08H) Length() uint8 { // 0xcf
	return 1
}

func (o *RST_08H) cycles() []uint8 { // 0xcf
	return []uint8{16}
}

func (o *RST_08H) String() string { // 0xcf
	return "RST" + " " + "08H" /* operand1 */
}

func (o *RST_08H) SymbolicString() string { // 0xcf
	return "RST 08H"
}

type DEC_C struct {
	addr         string      // 0xd
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *DEC_C) Write(w io.Writer) (int, error) { // 0xd
	var b []byte

	b = append(b, 0xd)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xd,
	})
	if err != nil {
		return written, err
	}

	// C is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *DEC_C) Length() uint8 { // 0xd
	return 1
}

func (o *DEC_C) cycles() []uint8 { // 0xd
	return []uint8{4}
}

func (o *DEC_C) String() string { // 0xd
	return "DEC" + " " + "C" /* operand1 */
}

func (o *DEC_C) SymbolicString() string { // 0xd
	return "DEC C"
}

type RET_NC struct {
	addr         string      // 0xd0
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RET_NC) Write(w io.Writer) (int, error) { // 0xd0
	var b []byte

	b = append(b, 0xd0)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xd0,
	})
	if err != nil {
		return written, err
	}

	// NC is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RET_NC) Length() uint8 { // 0xd0
	return 1
}

func (o *RET_NC) cycles() []uint8 { // 0xd0
	return []uint8{20, 8}
}

func (o *RET_NC) String() string { // 0xd0
	return "RET" + " " + "NC" /* operand1 */
}

func (o *RET_NC) SymbolicString() string { // 0xd0
	return "RET NC"
}

type POP_DE struct {
	addr         string      // 0xd1
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *POP_DE) Write(w io.Writer) (int, error) { // 0xd1
	var b []byte

	b = append(b, 0xd1)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xd1,
	})
	if err != nil {
		return written, err
	}

	// DE is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *POP_DE) Length() uint8 { // 0xd1
	return 1
}

func (o *POP_DE) cycles() []uint8 { // 0xd1
	return []uint8{12}
}

func (o *POP_DE) String() string { // 0xd1
	return "POP" + " " + "DE" /* operand1 */
}

func (o *POP_DE) SymbolicString() string { // 0xd1
	return "POP DE"
}

type JP_NC_a16 struct {
	addr         string      // 0xd2
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *JP_NC_a16) Write(w io.Writer) (int, error) { // 0xd2
	var b []byte

	b = append(b, 0xd2)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xd2,
	})
	if err != nil {
		return written, err
	}

	// NC is implicit in this instruction.
	err = writeImmediate16BitAddress(w, o.operand2)
	written += 2

	return written, err
}

func (o *JP_NC_a16) Length() uint8 { // 0xd2
	return 3
}

func (o *JP_NC_a16) cycles() []uint8 { // 0xd2
	return []uint8{16, 12}
}

func (o *JP_NC_a16) String() string { // 0xd2
	return "JP" + " " + "NC" /* operand1 */ + " " + fmt.Sprintf("%X", o.operand2)
}

func (o *JP_NC_a16) SymbolicString() string { // 0xd2
	return "JP NC,a16"
}

type CALL_NC_a16 struct {
	addr         string      // 0xd4
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *CALL_NC_a16) Write(w io.Writer) (int, error) { // 0xd4
	var b []byte

	b = append(b, 0xd4)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xd4,
	})
	if err != nil {
		return written, err
	}

	// NC is implicit in this instruction.
	err = writeImmediate16BitAddress(w, o.operand2)
	written += 2

	return written, err
}

func (o *CALL_NC_a16) Length() uint8 { // 0xd4
	return 3
}

func (o *CALL_NC_a16) cycles() []uint8 { // 0xd4
	return []uint8{24, 12}
}

func (o *CALL_NC_a16) String() string { // 0xd4
	return "CALL" + " " + "NC" /* operand1 */ + " " + fmt.Sprintf("%X", o.operand2)
}

func (o *CALL_NC_a16) SymbolicString() string { // 0xd4
	return "CALL NC,a16"
}

type PUSH_DE struct {
	addr         string      // 0xd5
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *PUSH_DE) Write(w io.Writer) (int, error) { // 0xd5
	var b []byte

	b = append(b, 0xd5)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xd5,
	})
	if err != nil {
		return written, err
	}

	// DE is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *PUSH_DE) Length() uint8 { // 0xd5
	return 1
}

func (o *PUSH_DE) cycles() []uint8 { // 0xd5
	return []uint8{16}
}

func (o *PUSH_DE) String() string { // 0xd5
	return "PUSH" + " " + "DE" /* operand1 */
}

func (o *PUSH_DE) SymbolicString() string { // 0xd5
	return "PUSH DE"
}

type SUB_d8 struct {
	addr         string      // 0xd6
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SUB_d8) Write(w io.Writer) (int, error) { // 0xd6
	var b []byte

	b = append(b, 0xd6)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xd6,
	})
	if err != nil {
		return written, err
	}

	err = writeImmediate8BitData(w, o.operand1)
	written += 1
	//  is implicit in this instruction.

	return written, err
}

func (o *SUB_d8) Length() uint8 { // 0xd6
	return 2
}

func (o *SUB_d8) cycles() []uint8 { // 0xd6
	return []uint8{8}
}

func (o *SUB_d8) String() string { // 0xd6
	return "SUB" + " " + fmt.Sprintf("%X", o.operand1)
}

func (o *SUB_d8) SymbolicString() string { // 0xd6
	return "SUB d8"
}

type RST_10H struct {
	addr         string      // 0xd7
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RST_10H) Write(w io.Writer) (int, error) { // 0xd7
	var b []byte

	b = append(b, 0xd7)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xd7,
	})
	if err != nil {
		return written, err
	}

	// 10H is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RST_10H) Length() uint8 { // 0xd7
	return 1
}

func (o *RST_10H) cycles() []uint8 { // 0xd7
	return []uint8{16}
}

func (o *RST_10H) String() string { // 0xd7
	return "RST" + " " + "10H" /* operand1 */
}

func (o *RST_10H) SymbolicString() string { // 0xd7
	return "RST 10H"
}

type RET_C struct {
	addr         string      // 0xd8
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RET_C) Write(w io.Writer) (int, error) { // 0xd8
	var b []byte

	b = append(b, 0xd8)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xd8,
	})
	if err != nil {
		return written, err
	}

	// C is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RET_C) Length() uint8 { // 0xd8
	return 1
}

func (o *RET_C) cycles() []uint8 { // 0xd8
	return []uint8{20, 8}
}

func (o *RET_C) String() string { // 0xd8
	return "RET" + " " + "C" /* operand1 */
}

func (o *RET_C) SymbolicString() string { // 0xd8
	return "RET C"
}

type RETI struct {
	addr         string      // 0xd9
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RETI) Write(w io.Writer) (int, error) { // 0xd9
	var b []byte

	b = append(b, 0xd9)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xd9,
	})
	if err != nil {
		return written, err
	}

	//  is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RETI) Length() uint8 { // 0xd9
	return 1
}

func (o *RETI) cycles() []uint8 { // 0xd9
	return []uint8{16}
}

func (o *RETI) String() string { // 0xd9
	return "RETI"
}

func (o *RETI) SymbolicString() string { // 0xd9
	return "RETI"
}

type JP_C_a16 struct {
	addr         string      // 0xda
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *JP_C_a16) Write(w io.Writer) (int, error) { // 0xda
	var b []byte

	b = append(b, 0xda)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xda,
	})
	if err != nil {
		return written, err
	}

	// C is implicit in this instruction.
	err = writeImmediate16BitAddress(w, o.operand2)
	written += 2

	return written, err
}

func (o *JP_C_a16) Length() uint8 { // 0xda
	return 3
}

func (o *JP_C_a16) cycles() []uint8 { // 0xda
	return []uint8{16, 12}
}

func (o *JP_C_a16) String() string { // 0xda
	return "JP" + " " + "C" /* operand1 */ + " " + fmt.Sprintf("%X", o.operand2)
}

func (o *JP_C_a16) SymbolicString() string { // 0xda
	return "JP C,a16"
}

type CALL_C_a16 struct {
	addr         string      // 0xdc
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *CALL_C_a16) Write(w io.Writer) (int, error) { // 0xdc
	var b []byte

	b = append(b, 0xdc)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xdc,
	})
	if err != nil {
		return written, err
	}

	// C is implicit in this instruction.
	err = writeImmediate16BitAddress(w, o.operand2)
	written += 2

	return written, err
}

func (o *CALL_C_a16) Length() uint8 { // 0xdc
	return 3
}

func (o *CALL_C_a16) cycles() []uint8 { // 0xdc
	return []uint8{24, 12}
}

func (o *CALL_C_a16) String() string { // 0xdc
	return "CALL" + " " + "C" /* operand1 */ + " " + fmt.Sprintf("%X", o.operand2)
}

func (o *CALL_C_a16) SymbolicString() string { // 0xdc
	return "CALL C,a16"
}

type SBC_A_d8 struct {
	addr         string      // 0xde
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SBC_A_d8) Write(w io.Writer) (int, error) { // 0xde
	var b []byte

	b = append(b, 0xde)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xde,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	err = writeImmediate8BitData(w, o.operand2)
	written += 1

	return written, err
}

func (o *SBC_A_d8) Length() uint8 { // 0xde
	return 2
}

func (o *SBC_A_d8) cycles() []uint8 { // 0xde
	return []uint8{8}
}

func (o *SBC_A_d8) String() string { // 0xde
	return "SBC" + " " + "A" /* operand1 */ + " " + fmt.Sprintf("%X", o.operand2)
}

func (o *SBC_A_d8) SymbolicString() string { // 0xde
	return "SBC A,d8"
}

type RST_18H struct {
	addr         string      // 0xdf
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RST_18H) Write(w io.Writer) (int, error) { // 0xdf
	var b []byte

	b = append(b, 0xdf)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xdf,
	})
	if err != nil {
		return written, err
	}

	// 18H is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RST_18H) Length() uint8 { // 0xdf
	return 1
}

func (o *RST_18H) cycles() []uint8 { // 0xdf
	return []uint8{16}
}

func (o *RST_18H) String() string { // 0xdf
	return "RST" + " " + "18H" /* operand1 */
}

func (o *RST_18H) SymbolicString() string { // 0xdf
	return "RST 18H"
}

type LD_C_d8 struct {
	addr         string      // 0xe
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_C_d8) Write(w io.Writer) (int, error) { // 0xe
	var b []byte

	b = append(b, 0xe)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xe,
	})
	if err != nil {
		return written, err
	}

	// C is implicit in this instruction.
	err = writeImmediate8BitData(w, o.operand2)
	written += 1

	return written, err
}

func (o *LD_C_d8) Length() uint8 { // 0xe
	return 2
}

func (o *LD_C_d8) cycles() []uint8 { // 0xe
	return []uint8{8}
}

func (o *LD_C_d8) String() string { // 0xe
	return "LD" + " " + "C" /* operand1 */ + " " + fmt.Sprintf("%X", o.operand2)
}

func (o *LD_C_d8) SymbolicString() string { // 0xe
	return "LD C,d8"
}

type LDH_a8Deref_A struct {
	addr         string      // 0xe0
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LDH_a8Deref_A) Write(w io.Writer) (int, error) { // 0xe0
	var b []byte

	b = append(b, 0xe0)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xe0,
	})
	if err != nil {
		return written, err
	}

	err = writeImmediate8BitAddress(w, o.operand1)
	written += 1
	// A is implicit in this instruction.

	return written, err
}

func (o *LDH_a8Deref_A) Length() uint8 { // 0xe0
	return 2
}

func (o *LDH_a8Deref_A) cycles() []uint8 { // 0xe0
	return []uint8{12}
}

func (o *LDH_a8Deref_A) String() string { // 0xe0
	return "LDH" + " " + fmt.Sprintf("(%X)", o.operand1) + " " + "A" /* operand2 */
}

func (o *LDH_a8Deref_A) SymbolicString() string { // 0xe0
	return "LDH (a8),A"
}

type POP_HL struct {
	addr         string      // 0xe1
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *POP_HL) Write(w io.Writer) (int, error) { // 0xe1
	var b []byte

	b = append(b, 0xe1)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xe1,
	})
	if err != nil {
		return written, err
	}

	// HL is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *POP_HL) Length() uint8 { // 0xe1
	return 1
}

func (o *POP_HL) cycles() []uint8 { // 0xe1
	return []uint8{12}
}

func (o *POP_HL) String() string { // 0xe1
	return "POP" + " " + "HL" /* operand1 */
}

func (o *POP_HL) SymbolicString() string { // 0xe1
	return "POP HL"
}

type LD_CDeref_A struct {
	addr         string      // 0xe2
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_CDeref_A) Write(w io.Writer) (int, error) { // 0xe2
	var b []byte

	b = append(b, 0xe2)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xe2,
	})
	if err != nil {
		return written, err
	}

	// (C) is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *LD_CDeref_A) Length() uint8 { // 0xe2
	return 1
}

func (o *LD_CDeref_A) cycles() []uint8 { // 0xe2
	return []uint8{8}
}

func (o *LD_CDeref_A) String() string { // 0xe2
	return "LD" + " " + "(C)" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *LD_CDeref_A) SymbolicString() string { // 0xe2
	return "LD (C),A"
}

type PUSH_HL struct {
	addr         string      // 0xe5
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *PUSH_HL) Write(w io.Writer) (int, error) { // 0xe5
	var b []byte

	b = append(b, 0xe5)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xe5,
	})
	if err != nil {
		return written, err
	}

	// HL is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *PUSH_HL) Length() uint8 { // 0xe5
	return 1
}

func (o *PUSH_HL) cycles() []uint8 { // 0xe5
	return []uint8{16}
}

func (o *PUSH_HL) String() string { // 0xe5
	return "PUSH" + " " + "HL" /* operand1 */
}

func (o *PUSH_HL) SymbolicString() string { // 0xe5
	return "PUSH HL"
}

type AND_d8 struct {
	addr         string      // 0xe6
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *AND_d8) Write(w io.Writer) (int, error) { // 0xe6
	var b []byte

	b = append(b, 0xe6)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xe6,
	})
	if err != nil {
		return written, err
	}

	err = writeImmediate8BitData(w, o.operand1)
	written += 1
	//  is implicit in this instruction.

	return written, err
}

func (o *AND_d8) Length() uint8 { // 0xe6
	return 2
}

func (o *AND_d8) cycles() []uint8 { // 0xe6
	return []uint8{8}
}

func (o *AND_d8) String() string { // 0xe6
	return "AND" + " " + fmt.Sprintf("%X", o.operand1)
}

func (o *AND_d8) SymbolicString() string { // 0xe6
	return "AND d8"
}

type RST_20H struct {
	addr         string      // 0xe7
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RST_20H) Write(w io.Writer) (int, error) { // 0xe7
	var b []byte

	b = append(b, 0xe7)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xe7,
	})
	if err != nil {
		return written, err
	}

	// 20H is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RST_20H) Length() uint8 { // 0xe7
	return 1
}

func (o *RST_20H) cycles() []uint8 { // 0xe7
	return []uint8{16}
}

func (o *RST_20H) String() string { // 0xe7
	return "RST" + " " + "20H" /* operand1 */
}

func (o *RST_20H) SymbolicString() string { // 0xe7
	return "RST 20H"
}

type ADD_SP_r8 struct {
	addr         string      // 0xe8
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *ADD_SP_r8) Write(w io.Writer) (int, error) { // 0xe8
	var b []byte

	b = append(b, 0xe8)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xe8,
	})
	if err != nil {
		return written, err
	}

	// SP is implicit in this instruction.
	err = writeImmediateSigned8BitData(w, o.operand2)
	written += 1

	return written, err
}

func (o *ADD_SP_r8) Length() uint8 { // 0xe8
	return 2
}

func (o *ADD_SP_r8) cycles() []uint8 { // 0xe8
	return []uint8{16}
}

func (o *ADD_SP_r8) String() string { // 0xe8
	return "ADD" + " " + "SP" /* operand1 */ + " " + fmt.Sprintf("SP+%X", o.operand2)
}

func (o *ADD_SP_r8) SymbolicString() string { // 0xe8
	return "ADD SP,r8"
}

type JP_HLPtr struct {
	addr         string      // 0xe9
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *JP_HLPtr) Write(w io.Writer) (int, error) { // 0xe9
	var b []byte

	b = append(b, 0xe9)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xe9,
	})
	if err != nil {
		return written, err
	}

	// (HL) is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *JP_HLPtr) Length() uint8 { // 0xe9
	return 1
}

func (o *JP_HLPtr) cycles() []uint8 { // 0xe9
	return []uint8{4}
}

func (o *JP_HLPtr) String() string { // 0xe9
	return "JP" + " " + "(HL)" /* operand1 */
}

func (o *JP_HLPtr) SymbolicString() string { // 0xe9
	return "JP (HL)"
}

type LD_a16Deref_A struct {
	addr         string      // 0xea
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_a16Deref_A) Write(w io.Writer) (int, error) { // 0xea
	var b []byte

	b = append(b, 0xea)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xea,
	})
	if err != nil {
		return written, err
	}

	err = writeImmediate16BitAddress(w, o.operand1)
	written += 2
	// A is implicit in this instruction.

	return written, err
}

func (o *LD_a16Deref_A) Length() uint8 { // 0xea
	return 3
}

func (o *LD_a16Deref_A) cycles() []uint8 { // 0xea
	return []uint8{16}
}

func (o *LD_a16Deref_A) String() string { // 0xea
	return "LD" + " " + fmt.Sprintf("%X", o.operand1) + " " + "A" /* operand2 */
}

func (o *LD_a16Deref_A) SymbolicString() string { // 0xea
	return "LD (a16),A"
}

type XOR_d8 struct {
	addr         string      // 0xee
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *XOR_d8) Write(w io.Writer) (int, error) { // 0xee
	var b []byte

	b = append(b, 0xee)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xee,
	})
	if err != nil {
		return written, err
	}

	err = writeImmediate8BitData(w, o.operand1)
	written += 1
	//  is implicit in this instruction.

	return written, err
}

func (o *XOR_d8) Length() uint8 { // 0xee
	return 2
}

func (o *XOR_d8) cycles() []uint8 { // 0xee
	return []uint8{8}
}

func (o *XOR_d8) String() string { // 0xee
	return "XOR" + " " + fmt.Sprintf("%X", o.operand1)
}

func (o *XOR_d8) SymbolicString() string { // 0xee
	return "XOR d8"
}

type RST_28H struct {
	addr         string      // 0xef
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RST_28H) Write(w io.Writer) (int, error) { // 0xef
	var b []byte

	b = append(b, 0xef)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xef,
	})
	if err != nil {
		return written, err
	}

	// 28H is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RST_28H) Length() uint8 { // 0xef
	return 1
}

func (o *RST_28H) cycles() []uint8 { // 0xef
	return []uint8{16}
}

func (o *RST_28H) String() string { // 0xef
	return "RST" + " " + "28H" /* operand1 */
}

func (o *RST_28H) SymbolicString() string { // 0xef
	return "RST 28H"
}

type RRCA struct {
	addr         string      // 0xf
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RRCA) Write(w io.Writer) (int, error) { // 0xf
	var b []byte

	b = append(b, 0xf)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xf,
	})
	if err != nil {
		return written, err
	}

	//  is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RRCA) Length() uint8 { // 0xf
	return 1
}

func (o *RRCA) cycles() []uint8 { // 0xf
	return []uint8{4}
}

func (o *RRCA) String() string { // 0xf
	return "RRCA"
}

func (o *RRCA) SymbolicString() string { // 0xf
	return "RRCA"
}

type LDH_A_a8Deref struct {
	addr         string      // 0xf0
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LDH_A_a8Deref) Write(w io.Writer) (int, error) { // 0xf0
	var b []byte

	b = append(b, 0xf0)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xf0,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	err = writeImmediate8BitAddress(w, o.operand2)
	written += 1

	return written, err
}

func (o *LDH_A_a8Deref) Length() uint8 { // 0xf0
	return 2
}

func (o *LDH_A_a8Deref) cycles() []uint8 { // 0xf0
	return []uint8{12}
}

func (o *LDH_A_a8Deref) String() string { // 0xf0
	return "LDH" + " " + "A" /* operand1 */ + " " + fmt.Sprintf("(%X)", o.operand2)
}

func (o *LDH_A_a8Deref) SymbolicString() string { // 0xf0
	return "LDH A,(a8)"
}

type POP_AF struct {
	addr         string      // 0xf1
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *POP_AF) Write(w io.Writer) (int, error) { // 0xf1
	var b []byte

	b = append(b, 0xf1)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xf1,
	})
	if err != nil {
		return written, err
	}

	// AF is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *POP_AF) Length() uint8 { // 0xf1
	return 1
}

func (o *POP_AF) cycles() []uint8 { // 0xf1
	return []uint8{12}
}

func (o *POP_AF) String() string { // 0xf1
	return "POP" + " " + "AF" /* operand1 */
}

func (o *POP_AF) SymbolicString() string { // 0xf1
	return "POP AF"
}

type LD_A_CDeref struct {
	addr         string      // 0xf2
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_A_CDeref) Write(w io.Writer) (int, error) { // 0xf2
	var b []byte

	b = append(b, 0xf2)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xf2,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	// (C) is implicit in this instruction.

	return written, err
}

func (o *LD_A_CDeref) Length() uint8 { // 0xf2
	return 1
}

func (o *LD_A_CDeref) cycles() []uint8 { // 0xf2
	return []uint8{8}
}

func (o *LD_A_CDeref) String() string { // 0xf2
	return "LD" + " " + "A" /* operand1 */ + " " + "(C)" /* operand2 */
}

func (o *LD_A_CDeref) SymbolicString() string { // 0xf2
	return "LD A,(C)"
}

type DI struct {
	addr         string      // 0xf3
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *DI) Write(w io.Writer) (int, error) { // 0xf3
	var b []byte

	b = append(b, 0xf3)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xf3,
	})
	if err != nil {
		return written, err
	}

	//  is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *DI) Length() uint8 { // 0xf3
	return 1
}

func (o *DI) cycles() []uint8 { // 0xf3
	return []uint8{4}
}

func (o *DI) String() string { // 0xf3
	return "DI"
}

func (o *DI) SymbolicString() string { // 0xf3
	return "DI"
}

type PUSH_AF struct {
	addr         string      // 0xf5
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *PUSH_AF) Write(w io.Writer) (int, error) { // 0xf5
	var b []byte

	b = append(b, 0xf5)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xf5,
	})
	if err != nil {
		return written, err
	}

	// AF is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *PUSH_AF) Length() uint8 { // 0xf5
	return 1
}

func (o *PUSH_AF) cycles() []uint8 { // 0xf5
	return []uint8{16}
}

func (o *PUSH_AF) String() string { // 0xf5
	return "PUSH" + " " + "AF" /* operand1 */
}

func (o *PUSH_AF) SymbolicString() string { // 0xf5
	return "PUSH AF"
}

type OR_d8 struct {
	addr         string      // 0xf6
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *OR_d8) Write(w io.Writer) (int, error) { // 0xf6
	var b []byte

	b = append(b, 0xf6)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xf6,
	})
	if err != nil {
		return written, err
	}

	err = writeImmediate8BitData(w, o.operand1)
	written += 1
	//  is implicit in this instruction.

	return written, err
}

func (o *OR_d8) Length() uint8 { // 0xf6
	return 2
}

func (o *OR_d8) cycles() []uint8 { // 0xf6
	return []uint8{8}
}

func (o *OR_d8) String() string { // 0xf6
	return "OR" + " " + fmt.Sprintf("%X", o.operand1)
}

func (o *OR_d8) SymbolicString() string { // 0xf6
	return "OR d8"
}

type RST_30H struct {
	addr         string      // 0xf7
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RST_30H) Write(w io.Writer) (int, error) { // 0xf7
	var b []byte

	b = append(b, 0xf7)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xf7,
	})
	if err != nil {
		return written, err
	}

	// 30H is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RST_30H) Length() uint8 { // 0xf7
	return 1
}

func (o *RST_30H) cycles() []uint8 { // 0xf7
	return []uint8{16}
}

func (o *RST_30H) String() string { // 0xf7
	return "RST" + " " + "30H" /* operand1 */
}

func (o *RST_30H) SymbolicString() string { // 0xf7
	return "RST 30H"
}

type LD_HL_SP_plus_r8 struct {
	addr         string      // 0xf8
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_HL_SP_plus_r8) Write(w io.Writer) (int, error) { // 0xf8
	var b []byte

	b = append(b, 0xf8)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xf8,
	})
	if err != nil {
		return written, err
	}

	// HL is implicit in this instruction.
	err = writeImmediateSigned8BitData(w, o.operand2)
	written += 1

	return written, err
}

func (o *LD_HL_SP_plus_r8) Length() uint8 { // 0xf8
	return 2
}

func (o *LD_HL_SP_plus_r8) cycles() []uint8 { // 0xf8
	return []uint8{12}
}

func (o *LD_HL_SP_plus_r8) String() string { // 0xf8
	return "LD" + " " + "HL" /* operand1 */ + " " + fmt.Sprintf("SP+%X", o.operand2)
}

func (o *LD_HL_SP_plus_r8) SymbolicString() string { // 0xf8
	return "LD HL,SP+r8"
}

type LD_SP_HL struct {
	addr         string      // 0xf9
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_SP_HL) Write(w io.Writer) (int, error) { // 0xf9
	var b []byte

	b = append(b, 0xf9)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xf9,
	})
	if err != nil {
		return written, err
	}

	// SP is implicit in this instruction.
	// HL is implicit in this instruction.

	return written, err
}

func (o *LD_SP_HL) Length() uint8 { // 0xf9
	return 1
}

func (o *LD_SP_HL) cycles() []uint8 { // 0xf9
	return []uint8{8}
}

func (o *LD_SP_HL) String() string { // 0xf9
	return "LD" + " " + "SP" /* operand1 */ + " " + "HL" /* operand2 */
}

func (o *LD_SP_HL) SymbolicString() string { // 0xf9
	return "LD SP,HL"
}

type LD_A_a16Deref struct {
	addr         string      // 0xfa
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *LD_A_a16Deref) Write(w io.Writer) (int, error) { // 0xfa
	var b []byte

	b = append(b, 0xfa)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xfa,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	err = writeImmediate16BitAddress(w, o.operand2)
	written += 2

	return written, err
}

func (o *LD_A_a16Deref) Length() uint8 { // 0xfa
	return 3
}

func (o *LD_A_a16Deref) cycles() []uint8 { // 0xfa
	return []uint8{16}
}

func (o *LD_A_a16Deref) String() string { // 0xfa
	return "LD" + " " + "A" /* operand1 */ + " " + fmt.Sprintf("%X", o.operand2)
}

func (o *LD_A_a16Deref) SymbolicString() string { // 0xfa
	return "LD A,(a16)"
}

type EI struct {
	addr         string      // 0xfb
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *EI) Write(w io.Writer) (int, error) { // 0xfb
	var b []byte

	b = append(b, 0xfb)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xfb,
	})
	if err != nil {
		return written, err
	}

	//  is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *EI) Length() uint8 { // 0xfb
	return 1
}

func (o *EI) cycles() []uint8 { // 0xfb
	return []uint8{4}
}

func (o *EI) String() string { // 0xfb
	return "EI"
}

func (o *EI) SymbolicString() string { // 0xfb
	return "EI"
}

type CP_d8 struct {
	addr         string      // 0xfe
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *CP_d8) Write(w io.Writer) (int, error) { // 0xfe
	var b []byte

	b = append(b, 0xfe)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xfe,
	})
	if err != nil {
		return written, err
	}

	err = writeImmediate8BitData(w, o.operand1)
	written += 1
	//  is implicit in this instruction.

	return written, err
}

func (o *CP_d8) Length() uint8 { // 0xfe
	return 2
}

func (o *CP_d8) cycles() []uint8 { // 0xfe
	return []uint8{8}
}

func (o *CP_d8) String() string { // 0xfe
	return "CP" + " " + fmt.Sprintf("%X", o.operand1)
}

func (o *CP_d8) SymbolicString() string { // 0xfe
	return "CP d8"
}

type RST_38H struct {
	addr         string      // 0xff
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RST_38H) Write(w io.Writer) (int, error) { // 0xff
	var b []byte

	b = append(b, 0xff)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{

		0xff,
	})
	if err != nil {
		return written, err
	}

	// 38H is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RST_38H) Length() uint8 { // 0xff
	return 1
}

func (o *RST_38H) cycles() []uint8 { // 0xff
	return []uint8{16}
}

func (o *RST_38H) String() string { // 0xff
	return "RST" + " " + "38H" /* operand1 */
}

func (o *RST_38H) SymbolicString() string { // 0xff
	return "RST 38H"
}

type RLC_B struct {
	addr         string      // 0x0
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RLC_B) Write(w io.Writer) (int, error) { // 0x0
	var b []byte

	b = append(b, 0x0)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x0,
	})
	if err != nil {
		return written, err
	}

	// B is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RLC_B) Length() uint8 { // 0x0
	return 2
}

func (o *RLC_B) cycles() []uint8 { // 0x0
	return []uint8{8}
}

func (o *RLC_B) String() string { // 0x0
	return "RLC" + " " + "B" /* operand1 */
}

func (o *RLC_B) SymbolicString() string { // 0x0
	return "RLC B"
}

type RLC_C struct {
	addr         string      // 0x1
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RLC_C) Write(w io.Writer) (int, error) { // 0x1
	var b []byte

	b = append(b, 0x1)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x1,
	})
	if err != nil {
		return written, err
	}

	// C is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RLC_C) Length() uint8 { // 0x1
	return 2
}

func (o *RLC_C) cycles() []uint8 { // 0x1
	return []uint8{8}
}

func (o *RLC_C) String() string { // 0x1
	return "RLC" + " " + "C" /* operand1 */
}

func (o *RLC_C) SymbolicString() string { // 0x1
	return "RLC C"
}

type RL_B struct {
	addr         string      // 0x10
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RL_B) Write(w io.Writer) (int, error) { // 0x10
	var b []byte

	b = append(b, 0x10)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x10,
	})
	if err != nil {
		return written, err
	}

	// B is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RL_B) Length() uint8 { // 0x10
	return 2
}

func (o *RL_B) cycles() []uint8 { // 0x10
	return []uint8{8}
}

func (o *RL_B) String() string { // 0x10
	return "RL" + " " + "B" /* operand1 */
}

func (o *RL_B) SymbolicString() string { // 0x10
	return "RL B"
}

type RL_C struct {
	addr         string      // 0x11
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RL_C) Write(w io.Writer) (int, error) { // 0x11
	var b []byte

	b = append(b, 0x11)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x11,
	})
	if err != nil {
		return written, err
	}

	// C is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RL_C) Length() uint8 { // 0x11
	return 2
}

func (o *RL_C) cycles() []uint8 { // 0x11
	return []uint8{8}
}

func (o *RL_C) String() string { // 0x11
	return "RL" + " " + "C" /* operand1 */
}

func (o *RL_C) SymbolicString() string { // 0x11
	return "RL C"
}

type RL_D struct {
	addr         string      // 0x12
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RL_D) Write(w io.Writer) (int, error) { // 0x12
	var b []byte

	b = append(b, 0x12)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x12,
	})
	if err != nil {
		return written, err
	}

	// D is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RL_D) Length() uint8 { // 0x12
	return 2
}

func (o *RL_D) cycles() []uint8 { // 0x12
	return []uint8{8}
}

func (o *RL_D) String() string { // 0x12
	return "RL" + " " + "D" /* operand1 */
}

func (o *RL_D) SymbolicString() string { // 0x12
	return "RL D"
}

type RL_E struct {
	addr         string      // 0x13
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RL_E) Write(w io.Writer) (int, error) { // 0x13
	var b []byte

	b = append(b, 0x13)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x13,
	})
	if err != nil {
		return written, err
	}

	// E is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RL_E) Length() uint8 { // 0x13
	return 2
}

func (o *RL_E) cycles() []uint8 { // 0x13
	return []uint8{8}
}

func (o *RL_E) String() string { // 0x13
	return "RL" + " " + "E" /* operand1 */
}

func (o *RL_E) SymbolicString() string { // 0x13
	return "RL E"
}

type RL_H struct {
	addr         string      // 0x14
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RL_H) Write(w io.Writer) (int, error) { // 0x14
	var b []byte

	b = append(b, 0x14)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x14,
	})
	if err != nil {
		return written, err
	}

	// H is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RL_H) Length() uint8 { // 0x14
	return 2
}

func (o *RL_H) cycles() []uint8 { // 0x14
	return []uint8{8}
}

func (o *RL_H) String() string { // 0x14
	return "RL" + " " + "H" /* operand1 */
}

func (o *RL_H) SymbolicString() string { // 0x14
	return "RL H"
}

type RL_L struct {
	addr         string      // 0x15
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RL_L) Write(w io.Writer) (int, error) { // 0x15
	var b []byte

	b = append(b, 0x15)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x15,
	})
	if err != nil {
		return written, err
	}

	// L is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RL_L) Length() uint8 { // 0x15
	return 2
}

func (o *RL_L) cycles() []uint8 { // 0x15
	return []uint8{8}
}

func (o *RL_L) String() string { // 0x15
	return "RL" + " " + "L" /* operand1 */
}

func (o *RL_L) SymbolicString() string { // 0x15
	return "RL L"
}

type RL_HLPtr struct {
	addr         string      // 0x16
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RL_HLPtr) Write(w io.Writer) (int, error) { // 0x16
	var b []byte

	b = append(b, 0x16)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x16,
	})
	if err != nil {
		return written, err
	}

	// (HL) is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RL_HLPtr) Length() uint8 { // 0x16
	return 2
}

func (o *RL_HLPtr) cycles() []uint8 { // 0x16
	return []uint8{16}
}

func (o *RL_HLPtr) String() string { // 0x16
	return "RL" + " " + "(HL)" /* operand1 */
}

func (o *RL_HLPtr) SymbolicString() string { // 0x16
	return "RL (HL)"
}

type RL_A struct {
	addr         string      // 0x17
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RL_A) Write(w io.Writer) (int, error) { // 0x17
	var b []byte

	b = append(b, 0x17)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x17,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RL_A) Length() uint8 { // 0x17
	return 2
}

func (o *RL_A) cycles() []uint8 { // 0x17
	return []uint8{8}
}

func (o *RL_A) String() string { // 0x17
	return "RL" + " " + "A" /* operand1 */
}

func (o *RL_A) SymbolicString() string { // 0x17
	return "RL A"
}

type RR_B struct {
	addr         string      // 0x18
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RR_B) Write(w io.Writer) (int, error) { // 0x18
	var b []byte

	b = append(b, 0x18)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x18,
	})
	if err != nil {
		return written, err
	}

	// B is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RR_B) Length() uint8 { // 0x18
	return 2
}

func (o *RR_B) cycles() []uint8 { // 0x18
	return []uint8{8}
}

func (o *RR_B) String() string { // 0x18
	return "RR" + " " + "B" /* operand1 */
}

func (o *RR_B) SymbolicString() string { // 0x18
	return "RR B"
}

type RR_C struct {
	addr         string      // 0x19
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RR_C) Write(w io.Writer) (int, error) { // 0x19
	var b []byte

	b = append(b, 0x19)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x19,
	})
	if err != nil {
		return written, err
	}

	// C is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RR_C) Length() uint8 { // 0x19
	return 2
}

func (o *RR_C) cycles() []uint8 { // 0x19
	return []uint8{8}
}

func (o *RR_C) String() string { // 0x19
	return "RR" + " " + "C" /* operand1 */
}

func (o *RR_C) SymbolicString() string { // 0x19
	return "RR C"
}

type RR_D struct {
	addr         string      // 0x1a
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RR_D) Write(w io.Writer) (int, error) { // 0x1a
	var b []byte

	b = append(b, 0x1a)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x1a,
	})
	if err != nil {
		return written, err
	}

	// D is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RR_D) Length() uint8 { // 0x1a
	return 2
}

func (o *RR_D) cycles() []uint8 { // 0x1a
	return []uint8{8}
}

func (o *RR_D) String() string { // 0x1a
	return "RR" + " " + "D" /* operand1 */
}

func (o *RR_D) SymbolicString() string { // 0x1a
	return "RR D"
}

type RR_E struct {
	addr         string      // 0x1b
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RR_E) Write(w io.Writer) (int, error) { // 0x1b
	var b []byte

	b = append(b, 0x1b)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x1b,
	})
	if err != nil {
		return written, err
	}

	// E is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RR_E) Length() uint8 { // 0x1b
	return 2
}

func (o *RR_E) cycles() []uint8 { // 0x1b
	return []uint8{8}
}

func (o *RR_E) String() string { // 0x1b
	return "RR" + " " + "E" /* operand1 */
}

func (o *RR_E) SymbolicString() string { // 0x1b
	return "RR E"
}

type RR_H struct {
	addr         string      // 0x1c
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RR_H) Write(w io.Writer) (int, error) { // 0x1c
	var b []byte

	b = append(b, 0x1c)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x1c,
	})
	if err != nil {
		return written, err
	}

	// H is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RR_H) Length() uint8 { // 0x1c
	return 2
}

func (o *RR_H) cycles() []uint8 { // 0x1c
	return []uint8{8}
}

func (o *RR_H) String() string { // 0x1c
	return "RR" + " " + "H" /* operand1 */
}

func (o *RR_H) SymbolicString() string { // 0x1c
	return "RR H"
}

type RR_L struct {
	addr         string      // 0x1d
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RR_L) Write(w io.Writer) (int, error) { // 0x1d
	var b []byte

	b = append(b, 0x1d)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x1d,
	})
	if err != nil {
		return written, err
	}

	// L is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RR_L) Length() uint8 { // 0x1d
	return 2
}

func (o *RR_L) cycles() []uint8 { // 0x1d
	return []uint8{8}
}

func (o *RR_L) String() string { // 0x1d
	return "RR" + " " + "L" /* operand1 */
}

func (o *RR_L) SymbolicString() string { // 0x1d
	return "RR L"
}

type RR_HLPtr struct {
	addr         string      // 0x1e
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RR_HLPtr) Write(w io.Writer) (int, error) { // 0x1e
	var b []byte

	b = append(b, 0x1e)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x1e,
	})
	if err != nil {
		return written, err
	}

	// (HL) is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RR_HLPtr) Length() uint8 { // 0x1e
	return 2
}

func (o *RR_HLPtr) cycles() []uint8 { // 0x1e
	return []uint8{16}
}

func (o *RR_HLPtr) String() string { // 0x1e
	return "RR" + " " + "(HL)" /* operand1 */
}

func (o *RR_HLPtr) SymbolicString() string { // 0x1e
	return "RR (HL)"
}

type RR_A struct {
	addr         string      // 0x1f
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RR_A) Write(w io.Writer) (int, error) { // 0x1f
	var b []byte

	b = append(b, 0x1f)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x1f,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RR_A) Length() uint8 { // 0x1f
	return 2
}

func (o *RR_A) cycles() []uint8 { // 0x1f
	return []uint8{8}
}

func (o *RR_A) String() string { // 0x1f
	return "RR" + " " + "A" /* operand1 */
}

func (o *RR_A) SymbolicString() string { // 0x1f
	return "RR A"
}

type RLC_D struct {
	addr         string      // 0x2
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RLC_D) Write(w io.Writer) (int, error) { // 0x2
	var b []byte

	b = append(b, 0x2)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x2,
	})
	if err != nil {
		return written, err
	}

	// D is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RLC_D) Length() uint8 { // 0x2
	return 2
}

func (o *RLC_D) cycles() []uint8 { // 0x2
	return []uint8{8}
}

func (o *RLC_D) String() string { // 0x2
	return "RLC" + " " + "D" /* operand1 */
}

func (o *RLC_D) SymbolicString() string { // 0x2
	return "RLC D"
}

type SLA_B struct {
	addr         string      // 0x20
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SLA_B) Write(w io.Writer) (int, error) { // 0x20
	var b []byte

	b = append(b, 0x20)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x20,
	})
	if err != nil {
		return written, err
	}

	// B is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SLA_B) Length() uint8 { // 0x20
	return 2
}

func (o *SLA_B) cycles() []uint8 { // 0x20
	return []uint8{8}
}

func (o *SLA_B) String() string { // 0x20
	return "SLA" + " " + "B" /* operand1 */
}

func (o *SLA_B) SymbolicString() string { // 0x20
	return "SLA B"
}

type SLA_C struct {
	addr         string      // 0x21
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SLA_C) Write(w io.Writer) (int, error) { // 0x21
	var b []byte

	b = append(b, 0x21)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x21,
	})
	if err != nil {
		return written, err
	}

	// C is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SLA_C) Length() uint8 { // 0x21
	return 2
}

func (o *SLA_C) cycles() []uint8 { // 0x21
	return []uint8{8}
}

func (o *SLA_C) String() string { // 0x21
	return "SLA" + " " + "C" /* operand1 */
}

func (o *SLA_C) SymbolicString() string { // 0x21
	return "SLA C"
}

type SLA_D struct {
	addr         string      // 0x22
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SLA_D) Write(w io.Writer) (int, error) { // 0x22
	var b []byte

	b = append(b, 0x22)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x22,
	})
	if err != nil {
		return written, err
	}

	// D is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SLA_D) Length() uint8 { // 0x22
	return 2
}

func (o *SLA_D) cycles() []uint8 { // 0x22
	return []uint8{8}
}

func (o *SLA_D) String() string { // 0x22
	return "SLA" + " " + "D" /* operand1 */
}

func (o *SLA_D) SymbolicString() string { // 0x22
	return "SLA D"
}

type SLA_E struct {
	addr         string      // 0x23
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SLA_E) Write(w io.Writer) (int, error) { // 0x23
	var b []byte

	b = append(b, 0x23)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x23,
	})
	if err != nil {
		return written, err
	}

	// E is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SLA_E) Length() uint8 { // 0x23
	return 2
}

func (o *SLA_E) cycles() []uint8 { // 0x23
	return []uint8{8}
}

func (o *SLA_E) String() string { // 0x23
	return "SLA" + " " + "E" /* operand1 */
}

func (o *SLA_E) SymbolicString() string { // 0x23
	return "SLA E"
}

type SLA_H struct {
	addr         string      // 0x24
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SLA_H) Write(w io.Writer) (int, error) { // 0x24
	var b []byte

	b = append(b, 0x24)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x24,
	})
	if err != nil {
		return written, err
	}

	// H is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SLA_H) Length() uint8 { // 0x24
	return 2
}

func (o *SLA_H) cycles() []uint8 { // 0x24
	return []uint8{8}
}

func (o *SLA_H) String() string { // 0x24
	return "SLA" + " " + "H" /* operand1 */
}

func (o *SLA_H) SymbolicString() string { // 0x24
	return "SLA H"
}

type SLA_L struct {
	addr         string      // 0x25
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SLA_L) Write(w io.Writer) (int, error) { // 0x25
	var b []byte

	b = append(b, 0x25)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x25,
	})
	if err != nil {
		return written, err
	}

	// L is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SLA_L) Length() uint8 { // 0x25
	return 2
}

func (o *SLA_L) cycles() []uint8 { // 0x25
	return []uint8{8}
}

func (o *SLA_L) String() string { // 0x25
	return "SLA" + " " + "L" /* operand1 */
}

func (o *SLA_L) SymbolicString() string { // 0x25
	return "SLA L"
}

type SLA_HLPtr struct {
	addr         string      // 0x26
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SLA_HLPtr) Write(w io.Writer) (int, error) { // 0x26
	var b []byte

	b = append(b, 0x26)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x26,
	})
	if err != nil {
		return written, err
	}

	// (HL) is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SLA_HLPtr) Length() uint8 { // 0x26
	return 2
}

func (o *SLA_HLPtr) cycles() []uint8 { // 0x26
	return []uint8{16}
}

func (o *SLA_HLPtr) String() string { // 0x26
	return "SLA" + " " + "(HL)" /* operand1 */
}

func (o *SLA_HLPtr) SymbolicString() string { // 0x26
	return "SLA (HL)"
}

type SLA_A struct {
	addr         string      // 0x27
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SLA_A) Write(w io.Writer) (int, error) { // 0x27
	var b []byte

	b = append(b, 0x27)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x27,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SLA_A) Length() uint8 { // 0x27
	return 2
}

func (o *SLA_A) cycles() []uint8 { // 0x27
	return []uint8{8}
}

func (o *SLA_A) String() string { // 0x27
	return "SLA" + " " + "A" /* operand1 */
}

func (o *SLA_A) SymbolicString() string { // 0x27
	return "SLA A"
}

type SRA_B struct {
	addr         string      // 0x28
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SRA_B) Write(w io.Writer) (int, error) { // 0x28
	var b []byte

	b = append(b, 0x28)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x28,
	})
	if err != nil {
		return written, err
	}

	// B is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SRA_B) Length() uint8 { // 0x28
	return 2
}

func (o *SRA_B) cycles() []uint8 { // 0x28
	return []uint8{8}
}

func (o *SRA_B) String() string { // 0x28
	return "SRA" + " " + "B" /* operand1 */
}

func (o *SRA_B) SymbolicString() string { // 0x28
	return "SRA B"
}

type SRA_C struct {
	addr         string      // 0x29
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SRA_C) Write(w io.Writer) (int, error) { // 0x29
	var b []byte

	b = append(b, 0x29)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x29,
	})
	if err != nil {
		return written, err
	}

	// C is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SRA_C) Length() uint8 { // 0x29
	return 2
}

func (o *SRA_C) cycles() []uint8 { // 0x29
	return []uint8{8}
}

func (o *SRA_C) String() string { // 0x29
	return "SRA" + " " + "C" /* operand1 */
}

func (o *SRA_C) SymbolicString() string { // 0x29
	return "SRA C"
}

type SRA_D struct {
	addr         string      // 0x2a
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SRA_D) Write(w io.Writer) (int, error) { // 0x2a
	var b []byte

	b = append(b, 0x2a)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x2a,
	})
	if err != nil {
		return written, err
	}

	// D is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SRA_D) Length() uint8 { // 0x2a
	return 2
}

func (o *SRA_D) cycles() []uint8 { // 0x2a
	return []uint8{8}
}

func (o *SRA_D) String() string { // 0x2a
	return "SRA" + " " + "D" /* operand1 */
}

func (o *SRA_D) SymbolicString() string { // 0x2a
	return "SRA D"
}

type SRA_E struct {
	addr         string      // 0x2b
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SRA_E) Write(w io.Writer) (int, error) { // 0x2b
	var b []byte

	b = append(b, 0x2b)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x2b,
	})
	if err != nil {
		return written, err
	}

	// E is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SRA_E) Length() uint8 { // 0x2b
	return 2
}

func (o *SRA_E) cycles() []uint8 { // 0x2b
	return []uint8{8}
}

func (o *SRA_E) String() string { // 0x2b
	return "SRA" + " " + "E" /* operand1 */
}

func (o *SRA_E) SymbolicString() string { // 0x2b
	return "SRA E"
}

type SRA_H struct {
	addr         string      // 0x2c
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SRA_H) Write(w io.Writer) (int, error) { // 0x2c
	var b []byte

	b = append(b, 0x2c)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x2c,
	})
	if err != nil {
		return written, err
	}

	// H is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SRA_H) Length() uint8 { // 0x2c
	return 2
}

func (o *SRA_H) cycles() []uint8 { // 0x2c
	return []uint8{8}
}

func (o *SRA_H) String() string { // 0x2c
	return "SRA" + " " + "H" /* operand1 */
}

func (o *SRA_H) SymbolicString() string { // 0x2c
	return "SRA H"
}

type SRA_L struct {
	addr         string      // 0x2d
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SRA_L) Write(w io.Writer) (int, error) { // 0x2d
	var b []byte

	b = append(b, 0x2d)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x2d,
	})
	if err != nil {
		return written, err
	}

	// L is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SRA_L) Length() uint8 { // 0x2d
	return 2
}

func (o *SRA_L) cycles() []uint8 { // 0x2d
	return []uint8{8}
}

func (o *SRA_L) String() string { // 0x2d
	return "SRA" + " " + "L" /* operand1 */
}

func (o *SRA_L) SymbolicString() string { // 0x2d
	return "SRA L"
}

type SRA_HLPtr struct {
	addr         string      // 0x2e
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SRA_HLPtr) Write(w io.Writer) (int, error) { // 0x2e
	var b []byte

	b = append(b, 0x2e)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x2e,
	})
	if err != nil {
		return written, err
	}

	// (HL) is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SRA_HLPtr) Length() uint8 { // 0x2e
	return 2
}

func (o *SRA_HLPtr) cycles() []uint8 { // 0x2e
	return []uint8{16}
}

func (o *SRA_HLPtr) String() string { // 0x2e
	return "SRA" + " " + "(HL)" /* operand1 */
}

func (o *SRA_HLPtr) SymbolicString() string { // 0x2e
	return "SRA (HL)"
}

type SRA_A struct {
	addr         string      // 0x2f
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SRA_A) Write(w io.Writer) (int, error) { // 0x2f
	var b []byte

	b = append(b, 0x2f)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x2f,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SRA_A) Length() uint8 { // 0x2f
	return 2
}

func (o *SRA_A) cycles() []uint8 { // 0x2f
	return []uint8{8}
}

func (o *SRA_A) String() string { // 0x2f
	return "SRA" + " " + "A" /* operand1 */
}

func (o *SRA_A) SymbolicString() string { // 0x2f
	return "SRA A"
}

type RLC_E struct {
	addr         string      // 0x3
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RLC_E) Write(w io.Writer) (int, error) { // 0x3
	var b []byte

	b = append(b, 0x3)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x3,
	})
	if err != nil {
		return written, err
	}

	// E is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RLC_E) Length() uint8 { // 0x3
	return 2
}

func (o *RLC_E) cycles() []uint8 { // 0x3
	return []uint8{8}
}

func (o *RLC_E) String() string { // 0x3
	return "RLC" + " " + "E" /* operand1 */
}

func (o *RLC_E) SymbolicString() string { // 0x3
	return "RLC E"
}

type SWAP_B struct {
	addr         string      // 0x30
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SWAP_B) Write(w io.Writer) (int, error) { // 0x30
	var b []byte

	b = append(b, 0x30)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x30,
	})
	if err != nil {
		return written, err
	}

	// B is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SWAP_B) Length() uint8 { // 0x30
	return 2
}

func (o *SWAP_B) cycles() []uint8 { // 0x30
	return []uint8{8}
}

func (o *SWAP_B) String() string { // 0x30
	return "SWAP" + " " + "B" /* operand1 */
}

func (o *SWAP_B) SymbolicString() string { // 0x30
	return "SWAP B"
}

type SWAP_C struct {
	addr         string      // 0x31
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SWAP_C) Write(w io.Writer) (int, error) { // 0x31
	var b []byte

	b = append(b, 0x31)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x31,
	})
	if err != nil {
		return written, err
	}

	// C is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SWAP_C) Length() uint8 { // 0x31
	return 2
}

func (o *SWAP_C) cycles() []uint8 { // 0x31
	return []uint8{8}
}

func (o *SWAP_C) String() string { // 0x31
	return "SWAP" + " " + "C" /* operand1 */
}

func (o *SWAP_C) SymbolicString() string { // 0x31
	return "SWAP C"
}

type SWAP_D struct {
	addr         string      // 0x32
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SWAP_D) Write(w io.Writer) (int, error) { // 0x32
	var b []byte

	b = append(b, 0x32)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x32,
	})
	if err != nil {
		return written, err
	}

	// D is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SWAP_D) Length() uint8 { // 0x32
	return 2
}

func (o *SWAP_D) cycles() []uint8 { // 0x32
	return []uint8{8}
}

func (o *SWAP_D) String() string { // 0x32
	return "SWAP" + " " + "D" /* operand1 */
}

func (o *SWAP_D) SymbolicString() string { // 0x32
	return "SWAP D"
}

type SWAP_E struct {
	addr         string      // 0x33
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SWAP_E) Write(w io.Writer) (int, error) { // 0x33
	var b []byte

	b = append(b, 0x33)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x33,
	})
	if err != nil {
		return written, err
	}

	// E is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SWAP_E) Length() uint8 { // 0x33
	return 2
}

func (o *SWAP_E) cycles() []uint8 { // 0x33
	return []uint8{8}
}

func (o *SWAP_E) String() string { // 0x33
	return "SWAP" + " " + "E" /* operand1 */
}

func (o *SWAP_E) SymbolicString() string { // 0x33
	return "SWAP E"
}

type SWAP_H struct {
	addr         string      // 0x34
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SWAP_H) Write(w io.Writer) (int, error) { // 0x34
	var b []byte

	b = append(b, 0x34)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x34,
	})
	if err != nil {
		return written, err
	}

	// H is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SWAP_H) Length() uint8 { // 0x34
	return 2
}

func (o *SWAP_H) cycles() []uint8 { // 0x34
	return []uint8{8}
}

func (o *SWAP_H) String() string { // 0x34
	return "SWAP" + " " + "H" /* operand1 */
}

func (o *SWAP_H) SymbolicString() string { // 0x34
	return "SWAP H"
}

type SWAP_L struct {
	addr         string      // 0x35
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SWAP_L) Write(w io.Writer) (int, error) { // 0x35
	var b []byte

	b = append(b, 0x35)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x35,
	})
	if err != nil {
		return written, err
	}

	// L is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SWAP_L) Length() uint8 { // 0x35
	return 2
}

func (o *SWAP_L) cycles() []uint8 { // 0x35
	return []uint8{8}
}

func (o *SWAP_L) String() string { // 0x35
	return "SWAP" + " " + "L" /* operand1 */
}

func (o *SWAP_L) SymbolicString() string { // 0x35
	return "SWAP L"
}

type SWAP_HLPtr struct {
	addr         string      // 0x36
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SWAP_HLPtr) Write(w io.Writer) (int, error) { // 0x36
	var b []byte

	b = append(b, 0x36)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x36,
	})
	if err != nil {
		return written, err
	}

	// (HL) is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SWAP_HLPtr) Length() uint8 { // 0x36
	return 2
}

func (o *SWAP_HLPtr) cycles() []uint8 { // 0x36
	return []uint8{16}
}

func (o *SWAP_HLPtr) String() string { // 0x36
	return "SWAP" + " " + "(HL)" /* operand1 */
}

func (o *SWAP_HLPtr) SymbolicString() string { // 0x36
	return "SWAP (HL)"
}

type SWAP_A struct {
	addr         string      // 0x37
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SWAP_A) Write(w io.Writer) (int, error) { // 0x37
	var b []byte

	b = append(b, 0x37)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x37,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SWAP_A) Length() uint8 { // 0x37
	return 2
}

func (o *SWAP_A) cycles() []uint8 { // 0x37
	return []uint8{8}
}

func (o *SWAP_A) String() string { // 0x37
	return "SWAP" + " " + "A" /* operand1 */
}

func (o *SWAP_A) SymbolicString() string { // 0x37
	return "SWAP A"
}

type SRL_B struct {
	addr         string      // 0x38
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SRL_B) Write(w io.Writer) (int, error) { // 0x38
	var b []byte

	b = append(b, 0x38)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x38,
	})
	if err != nil {
		return written, err
	}

	// B is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SRL_B) Length() uint8 { // 0x38
	return 2
}

func (o *SRL_B) cycles() []uint8 { // 0x38
	return []uint8{8}
}

func (o *SRL_B) String() string { // 0x38
	return "SRL" + " " + "B" /* operand1 */
}

func (o *SRL_B) SymbolicString() string { // 0x38
	return "SRL B"
}

type SRL_C struct {
	addr         string      // 0x39
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SRL_C) Write(w io.Writer) (int, error) { // 0x39
	var b []byte

	b = append(b, 0x39)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x39,
	})
	if err != nil {
		return written, err
	}

	// C is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SRL_C) Length() uint8 { // 0x39
	return 2
}

func (o *SRL_C) cycles() []uint8 { // 0x39
	return []uint8{8}
}

func (o *SRL_C) String() string { // 0x39
	return "SRL" + " " + "C" /* operand1 */
}

func (o *SRL_C) SymbolicString() string { // 0x39
	return "SRL C"
}

type SRL_D struct {
	addr         string      // 0x3a
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SRL_D) Write(w io.Writer) (int, error) { // 0x3a
	var b []byte

	b = append(b, 0x3a)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x3a,
	})
	if err != nil {
		return written, err
	}

	// D is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SRL_D) Length() uint8 { // 0x3a
	return 2
}

func (o *SRL_D) cycles() []uint8 { // 0x3a
	return []uint8{8}
}

func (o *SRL_D) String() string { // 0x3a
	return "SRL" + " " + "D" /* operand1 */
}

func (o *SRL_D) SymbolicString() string { // 0x3a
	return "SRL D"
}

type SRL_E struct {
	addr         string      // 0x3b
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SRL_E) Write(w io.Writer) (int, error) { // 0x3b
	var b []byte

	b = append(b, 0x3b)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x3b,
	})
	if err != nil {
		return written, err
	}

	// E is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SRL_E) Length() uint8 { // 0x3b
	return 2
}

func (o *SRL_E) cycles() []uint8 { // 0x3b
	return []uint8{8}
}

func (o *SRL_E) String() string { // 0x3b
	return "SRL" + " " + "E" /* operand1 */
}

func (o *SRL_E) SymbolicString() string { // 0x3b
	return "SRL E"
}

type SRL_H struct {
	addr         string      // 0x3c
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SRL_H) Write(w io.Writer) (int, error) { // 0x3c
	var b []byte

	b = append(b, 0x3c)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x3c,
	})
	if err != nil {
		return written, err
	}

	// H is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SRL_H) Length() uint8 { // 0x3c
	return 2
}

func (o *SRL_H) cycles() []uint8 { // 0x3c
	return []uint8{8}
}

func (o *SRL_H) String() string { // 0x3c
	return "SRL" + " " + "H" /* operand1 */
}

func (o *SRL_H) SymbolicString() string { // 0x3c
	return "SRL H"
}

type SRL_L struct {
	addr         string      // 0x3d
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SRL_L) Write(w io.Writer) (int, error) { // 0x3d
	var b []byte

	b = append(b, 0x3d)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x3d,
	})
	if err != nil {
		return written, err
	}

	// L is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SRL_L) Length() uint8 { // 0x3d
	return 2
}

func (o *SRL_L) cycles() []uint8 { // 0x3d
	return []uint8{8}
}

func (o *SRL_L) String() string { // 0x3d
	return "SRL" + " " + "L" /* operand1 */
}

func (o *SRL_L) SymbolicString() string { // 0x3d
	return "SRL L"
}

type SRL_HLPtr struct {
	addr         string      // 0x3e
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SRL_HLPtr) Write(w io.Writer) (int, error) { // 0x3e
	var b []byte

	b = append(b, 0x3e)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x3e,
	})
	if err != nil {
		return written, err
	}

	// (HL) is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SRL_HLPtr) Length() uint8 { // 0x3e
	return 2
}

func (o *SRL_HLPtr) cycles() []uint8 { // 0x3e
	return []uint8{16}
}

func (o *SRL_HLPtr) String() string { // 0x3e
	return "SRL" + " " + "(HL)" /* operand1 */
}

func (o *SRL_HLPtr) SymbolicString() string { // 0x3e
	return "SRL (HL)"
}

type SRL_A struct {
	addr         string      // 0x3f
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SRL_A) Write(w io.Writer) (int, error) { // 0x3f
	var b []byte

	b = append(b, 0x3f)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x3f,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *SRL_A) Length() uint8 { // 0x3f
	return 2
}

func (o *SRL_A) cycles() []uint8 { // 0x3f
	return []uint8{8}
}

func (o *SRL_A) String() string { // 0x3f
	return "SRL" + " " + "A" /* operand1 */
}

func (o *SRL_A) SymbolicString() string { // 0x3f
	return "SRL A"
}

type RLC_H struct {
	addr         string      // 0x4
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RLC_H) Write(w io.Writer) (int, error) { // 0x4
	var b []byte

	b = append(b, 0x4)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x4,
	})
	if err != nil {
		return written, err
	}

	// H is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RLC_H) Length() uint8 { // 0x4
	return 2
}

func (o *RLC_H) cycles() []uint8 { // 0x4
	return []uint8{8}
}

func (o *RLC_H) String() string { // 0x4
	return "RLC" + " " + "H" /* operand1 */
}

func (o *RLC_H) SymbolicString() string { // 0x4
	return "RLC H"
}

type BIT_0_B struct {
	addr         string      // 0x40
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_0_B) Write(w io.Writer) (int, error) { // 0x40
	var b []byte

	b = append(b, 0x40)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x40,
	})
	if err != nil {
		return written, err
	}

	// 0 is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *BIT_0_B) Length() uint8 { // 0x40
	return 2
}

func (o *BIT_0_B) cycles() []uint8 { // 0x40
	return []uint8{8}
}

func (o *BIT_0_B) String() string { // 0x40
	return "BIT" + " " + "0" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *BIT_0_B) SymbolicString() string { // 0x40
	return "BIT 0,B"
}

type BIT_0_C struct {
	addr         string      // 0x41
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_0_C) Write(w io.Writer) (int, error) { // 0x41
	var b []byte

	b = append(b, 0x41)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x41,
	})
	if err != nil {
		return written, err
	}

	// 0 is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *BIT_0_C) Length() uint8 { // 0x41
	return 2
}

func (o *BIT_0_C) cycles() []uint8 { // 0x41
	return []uint8{8}
}

func (o *BIT_0_C) String() string { // 0x41
	return "BIT" + " " + "0" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *BIT_0_C) SymbolicString() string { // 0x41
	return "BIT 0,C"
}

type BIT_0_D struct {
	addr         string      // 0x42
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_0_D) Write(w io.Writer) (int, error) { // 0x42
	var b []byte

	b = append(b, 0x42)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x42,
	})
	if err != nil {
		return written, err
	}

	// 0 is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *BIT_0_D) Length() uint8 { // 0x42
	return 2
}

func (o *BIT_0_D) cycles() []uint8 { // 0x42
	return []uint8{8}
}

func (o *BIT_0_D) String() string { // 0x42
	return "BIT" + " " + "0" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *BIT_0_D) SymbolicString() string { // 0x42
	return "BIT 0,D"
}

type BIT_0_E struct {
	addr         string      // 0x43
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_0_E) Write(w io.Writer) (int, error) { // 0x43
	var b []byte

	b = append(b, 0x43)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x43,
	})
	if err != nil {
		return written, err
	}

	// 0 is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *BIT_0_E) Length() uint8 { // 0x43
	return 2
}

func (o *BIT_0_E) cycles() []uint8 { // 0x43
	return []uint8{8}
}

func (o *BIT_0_E) String() string { // 0x43
	return "BIT" + " " + "0" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *BIT_0_E) SymbolicString() string { // 0x43
	return "BIT 0,E"
}

type BIT_0_H struct {
	addr         string      // 0x44
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_0_H) Write(w io.Writer) (int, error) { // 0x44
	var b []byte

	b = append(b, 0x44)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x44,
	})
	if err != nil {
		return written, err
	}

	// 0 is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *BIT_0_H) Length() uint8 { // 0x44
	return 2
}

func (o *BIT_0_H) cycles() []uint8 { // 0x44
	return []uint8{8}
}

func (o *BIT_0_H) String() string { // 0x44
	return "BIT" + " " + "0" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *BIT_0_H) SymbolicString() string { // 0x44
	return "BIT 0,H"
}

type BIT_0_L struct {
	addr         string      // 0x45
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_0_L) Write(w io.Writer) (int, error) { // 0x45
	var b []byte

	b = append(b, 0x45)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x45,
	})
	if err != nil {
		return written, err
	}

	// 0 is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *BIT_0_L) Length() uint8 { // 0x45
	return 2
}

func (o *BIT_0_L) cycles() []uint8 { // 0x45
	return []uint8{8}
}

func (o *BIT_0_L) String() string { // 0x45
	return "BIT" + " " + "0" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *BIT_0_L) SymbolicString() string { // 0x45
	return "BIT 0,L"
}

type BIT_0_HLPtr struct {
	addr         string      // 0x46
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_0_HLPtr) Write(w io.Writer) (int, error) { // 0x46
	var b []byte

	b = append(b, 0x46)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x46,
	})
	if err != nil {
		return written, err
	}

	// 0 is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *BIT_0_HLPtr) Length() uint8 { // 0x46
	return 2
}

func (o *BIT_0_HLPtr) cycles() []uint8 { // 0x46
	return []uint8{16}
}

func (o *BIT_0_HLPtr) String() string { // 0x46
	return "BIT" + " " + "0" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *BIT_0_HLPtr) SymbolicString() string { // 0x46
	return "BIT 0,(HL)"
}

type BIT_0_A struct {
	addr         string      // 0x47
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_0_A) Write(w io.Writer) (int, error) { // 0x47
	var b []byte

	b = append(b, 0x47)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x47,
	})
	if err != nil {
		return written, err
	}

	// 0 is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *BIT_0_A) Length() uint8 { // 0x47
	return 2
}

func (o *BIT_0_A) cycles() []uint8 { // 0x47
	return []uint8{8}
}

func (o *BIT_0_A) String() string { // 0x47
	return "BIT" + " " + "0" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *BIT_0_A) SymbolicString() string { // 0x47
	return "BIT 0,A"
}

type BIT_1_B struct {
	addr         string      // 0x48
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_1_B) Write(w io.Writer) (int, error) { // 0x48
	var b []byte

	b = append(b, 0x48)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x48,
	})
	if err != nil {
		return written, err
	}

	// 1 is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *BIT_1_B) Length() uint8 { // 0x48
	return 2
}

func (o *BIT_1_B) cycles() []uint8 { // 0x48
	return []uint8{8}
}

func (o *BIT_1_B) String() string { // 0x48
	return "BIT" + " " + "1" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *BIT_1_B) SymbolicString() string { // 0x48
	return "BIT 1,B"
}

type BIT_1_C struct {
	addr         string      // 0x49
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_1_C) Write(w io.Writer) (int, error) { // 0x49
	var b []byte

	b = append(b, 0x49)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x49,
	})
	if err != nil {
		return written, err
	}

	// 1 is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *BIT_1_C) Length() uint8 { // 0x49
	return 2
}

func (o *BIT_1_C) cycles() []uint8 { // 0x49
	return []uint8{8}
}

func (o *BIT_1_C) String() string { // 0x49
	return "BIT" + " " + "1" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *BIT_1_C) SymbolicString() string { // 0x49
	return "BIT 1,C"
}

type BIT_1_D struct {
	addr         string      // 0x4a
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_1_D) Write(w io.Writer) (int, error) { // 0x4a
	var b []byte

	b = append(b, 0x4a)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x4a,
	})
	if err != nil {
		return written, err
	}

	// 1 is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *BIT_1_D) Length() uint8 { // 0x4a
	return 2
}

func (o *BIT_1_D) cycles() []uint8 { // 0x4a
	return []uint8{8}
}

func (o *BIT_1_D) String() string { // 0x4a
	return "BIT" + " " + "1" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *BIT_1_D) SymbolicString() string { // 0x4a
	return "BIT 1,D"
}

type BIT_1_E struct {
	addr         string      // 0x4b
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_1_E) Write(w io.Writer) (int, error) { // 0x4b
	var b []byte

	b = append(b, 0x4b)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x4b,
	})
	if err != nil {
		return written, err
	}

	// 1 is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *BIT_1_E) Length() uint8 { // 0x4b
	return 2
}

func (o *BIT_1_E) cycles() []uint8 { // 0x4b
	return []uint8{8}
}

func (o *BIT_1_E) String() string { // 0x4b
	return "BIT" + " " + "1" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *BIT_1_E) SymbolicString() string { // 0x4b
	return "BIT 1,E"
}

type BIT_1_H struct {
	addr         string      // 0x4c
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_1_H) Write(w io.Writer) (int, error) { // 0x4c
	var b []byte

	b = append(b, 0x4c)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x4c,
	})
	if err != nil {
		return written, err
	}

	// 1 is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *BIT_1_H) Length() uint8 { // 0x4c
	return 2
}

func (o *BIT_1_H) cycles() []uint8 { // 0x4c
	return []uint8{8}
}

func (o *BIT_1_H) String() string { // 0x4c
	return "BIT" + " " + "1" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *BIT_1_H) SymbolicString() string { // 0x4c
	return "BIT 1,H"
}

type BIT_1_L struct {
	addr         string      // 0x4d
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_1_L) Write(w io.Writer) (int, error) { // 0x4d
	var b []byte

	b = append(b, 0x4d)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x4d,
	})
	if err != nil {
		return written, err
	}

	// 1 is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *BIT_1_L) Length() uint8 { // 0x4d
	return 2
}

func (o *BIT_1_L) cycles() []uint8 { // 0x4d
	return []uint8{8}
}

func (o *BIT_1_L) String() string { // 0x4d
	return "BIT" + " " + "1" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *BIT_1_L) SymbolicString() string { // 0x4d
	return "BIT 1,L"
}

type BIT_1_HLPtr struct {
	addr         string      // 0x4e
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_1_HLPtr) Write(w io.Writer) (int, error) { // 0x4e
	var b []byte

	b = append(b, 0x4e)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x4e,
	})
	if err != nil {
		return written, err
	}

	// 1 is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *BIT_1_HLPtr) Length() uint8 { // 0x4e
	return 2
}

func (o *BIT_1_HLPtr) cycles() []uint8 { // 0x4e
	return []uint8{16}
}

func (o *BIT_1_HLPtr) String() string { // 0x4e
	return "BIT" + " " + "1" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *BIT_1_HLPtr) SymbolicString() string { // 0x4e
	return "BIT 1,(HL)"
}

type BIT_1_A struct {
	addr         string      // 0x4f
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_1_A) Write(w io.Writer) (int, error) { // 0x4f
	var b []byte

	b = append(b, 0x4f)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x4f,
	})
	if err != nil {
		return written, err
	}

	// 1 is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *BIT_1_A) Length() uint8 { // 0x4f
	return 2
}

func (o *BIT_1_A) cycles() []uint8 { // 0x4f
	return []uint8{8}
}

func (o *BIT_1_A) String() string { // 0x4f
	return "BIT" + " " + "1" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *BIT_1_A) SymbolicString() string { // 0x4f
	return "BIT 1,A"
}

type RLC_L struct {
	addr         string      // 0x5
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RLC_L) Write(w io.Writer) (int, error) { // 0x5
	var b []byte

	b = append(b, 0x5)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x5,
	})
	if err != nil {
		return written, err
	}

	// L is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RLC_L) Length() uint8 { // 0x5
	return 2
}

func (o *RLC_L) cycles() []uint8 { // 0x5
	return []uint8{8}
}

func (o *RLC_L) String() string { // 0x5
	return "RLC" + " " + "L" /* operand1 */
}

func (o *RLC_L) SymbolicString() string { // 0x5
	return "RLC L"
}

type BIT_2_B struct {
	addr         string      // 0x50
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_2_B) Write(w io.Writer) (int, error) { // 0x50
	var b []byte

	b = append(b, 0x50)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x50,
	})
	if err != nil {
		return written, err
	}

	// 2 is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *BIT_2_B) Length() uint8 { // 0x50
	return 2
}

func (o *BIT_2_B) cycles() []uint8 { // 0x50
	return []uint8{8}
}

func (o *BIT_2_B) String() string { // 0x50
	return "BIT" + " " + "2" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *BIT_2_B) SymbolicString() string { // 0x50
	return "BIT 2,B"
}

type BIT_2_C struct {
	addr         string      // 0x51
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_2_C) Write(w io.Writer) (int, error) { // 0x51
	var b []byte

	b = append(b, 0x51)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x51,
	})
	if err != nil {
		return written, err
	}

	// 2 is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *BIT_2_C) Length() uint8 { // 0x51
	return 2
}

func (o *BIT_2_C) cycles() []uint8 { // 0x51
	return []uint8{8}
}

func (o *BIT_2_C) String() string { // 0x51
	return "BIT" + " " + "2" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *BIT_2_C) SymbolicString() string { // 0x51
	return "BIT 2,C"
}

type BIT_2_D struct {
	addr         string      // 0x52
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_2_D) Write(w io.Writer) (int, error) { // 0x52
	var b []byte

	b = append(b, 0x52)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x52,
	})
	if err != nil {
		return written, err
	}

	// 2 is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *BIT_2_D) Length() uint8 { // 0x52
	return 2
}

func (o *BIT_2_D) cycles() []uint8 { // 0x52
	return []uint8{8}
}

func (o *BIT_2_D) String() string { // 0x52
	return "BIT" + " " + "2" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *BIT_2_D) SymbolicString() string { // 0x52
	return "BIT 2,D"
}

type BIT_2_E struct {
	addr         string      // 0x53
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_2_E) Write(w io.Writer) (int, error) { // 0x53
	var b []byte

	b = append(b, 0x53)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x53,
	})
	if err != nil {
		return written, err
	}

	// 2 is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *BIT_2_E) Length() uint8 { // 0x53
	return 2
}

func (o *BIT_2_E) cycles() []uint8 { // 0x53
	return []uint8{8}
}

func (o *BIT_2_E) String() string { // 0x53
	return "BIT" + " " + "2" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *BIT_2_E) SymbolicString() string { // 0x53
	return "BIT 2,E"
}

type BIT_2_H struct {
	addr         string      // 0x54
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_2_H) Write(w io.Writer) (int, error) { // 0x54
	var b []byte

	b = append(b, 0x54)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x54,
	})
	if err != nil {
		return written, err
	}

	// 2 is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *BIT_2_H) Length() uint8 { // 0x54
	return 2
}

func (o *BIT_2_H) cycles() []uint8 { // 0x54
	return []uint8{8}
}

func (o *BIT_2_H) String() string { // 0x54
	return "BIT" + " " + "2" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *BIT_2_H) SymbolicString() string { // 0x54
	return "BIT 2,H"
}

type BIT_2_L struct {
	addr         string      // 0x55
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_2_L) Write(w io.Writer) (int, error) { // 0x55
	var b []byte

	b = append(b, 0x55)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x55,
	})
	if err != nil {
		return written, err
	}

	// 2 is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *BIT_2_L) Length() uint8 { // 0x55
	return 2
}

func (o *BIT_2_L) cycles() []uint8 { // 0x55
	return []uint8{8}
}

func (o *BIT_2_L) String() string { // 0x55
	return "BIT" + " " + "2" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *BIT_2_L) SymbolicString() string { // 0x55
	return "BIT 2,L"
}

type BIT_2_HLPtr struct {
	addr         string      // 0x56
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_2_HLPtr) Write(w io.Writer) (int, error) { // 0x56
	var b []byte

	b = append(b, 0x56)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x56,
	})
	if err != nil {
		return written, err
	}

	// 2 is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *BIT_2_HLPtr) Length() uint8 { // 0x56
	return 2
}

func (o *BIT_2_HLPtr) cycles() []uint8 { // 0x56
	return []uint8{16}
}

func (o *BIT_2_HLPtr) String() string { // 0x56
	return "BIT" + " " + "2" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *BIT_2_HLPtr) SymbolicString() string { // 0x56
	return "BIT 2,(HL)"
}

type BIT_2_A struct {
	addr         string      // 0x57
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_2_A) Write(w io.Writer) (int, error) { // 0x57
	var b []byte

	b = append(b, 0x57)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x57,
	})
	if err != nil {
		return written, err
	}

	// 2 is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *BIT_2_A) Length() uint8 { // 0x57
	return 2
}

func (o *BIT_2_A) cycles() []uint8 { // 0x57
	return []uint8{8}
}

func (o *BIT_2_A) String() string { // 0x57
	return "BIT" + " " + "2" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *BIT_2_A) SymbolicString() string { // 0x57
	return "BIT 2,A"
}

type BIT_3_B struct {
	addr         string      // 0x58
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_3_B) Write(w io.Writer) (int, error) { // 0x58
	var b []byte

	b = append(b, 0x58)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x58,
	})
	if err != nil {
		return written, err
	}

	// 3 is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *BIT_3_B) Length() uint8 { // 0x58
	return 2
}

func (o *BIT_3_B) cycles() []uint8 { // 0x58
	return []uint8{8}
}

func (o *BIT_3_B) String() string { // 0x58
	return "BIT" + " " + "3" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *BIT_3_B) SymbolicString() string { // 0x58
	return "BIT 3,B"
}

type BIT_3_C struct {
	addr         string      // 0x59
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_3_C) Write(w io.Writer) (int, error) { // 0x59
	var b []byte

	b = append(b, 0x59)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x59,
	})
	if err != nil {
		return written, err
	}

	// 3 is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *BIT_3_C) Length() uint8 { // 0x59
	return 2
}

func (o *BIT_3_C) cycles() []uint8 { // 0x59
	return []uint8{8}
}

func (o *BIT_3_C) String() string { // 0x59
	return "BIT" + " " + "3" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *BIT_3_C) SymbolicString() string { // 0x59
	return "BIT 3,C"
}

type BIT_3_D struct {
	addr         string      // 0x5a
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_3_D) Write(w io.Writer) (int, error) { // 0x5a
	var b []byte

	b = append(b, 0x5a)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x5a,
	})
	if err != nil {
		return written, err
	}

	// 3 is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *BIT_3_D) Length() uint8 { // 0x5a
	return 2
}

func (o *BIT_3_D) cycles() []uint8 { // 0x5a
	return []uint8{8}
}

func (o *BIT_3_D) String() string { // 0x5a
	return "BIT" + " " + "3" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *BIT_3_D) SymbolicString() string { // 0x5a
	return "BIT 3,D"
}

type BIT_3_E struct {
	addr         string      // 0x5b
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_3_E) Write(w io.Writer) (int, error) { // 0x5b
	var b []byte

	b = append(b, 0x5b)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x5b,
	})
	if err != nil {
		return written, err
	}

	// 3 is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *BIT_3_E) Length() uint8 { // 0x5b
	return 2
}

func (o *BIT_3_E) cycles() []uint8 { // 0x5b
	return []uint8{8}
}

func (o *BIT_3_E) String() string { // 0x5b
	return "BIT" + " " + "3" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *BIT_3_E) SymbolicString() string { // 0x5b
	return "BIT 3,E"
}

type BIT_3_H struct {
	addr         string      // 0x5c
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_3_H) Write(w io.Writer) (int, error) { // 0x5c
	var b []byte

	b = append(b, 0x5c)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x5c,
	})
	if err != nil {
		return written, err
	}

	// 3 is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *BIT_3_H) Length() uint8 { // 0x5c
	return 2
}

func (o *BIT_3_H) cycles() []uint8 { // 0x5c
	return []uint8{8}
}

func (o *BIT_3_H) String() string { // 0x5c
	return "BIT" + " " + "3" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *BIT_3_H) SymbolicString() string { // 0x5c
	return "BIT 3,H"
}

type BIT_3_L struct {
	addr         string      // 0x5d
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_3_L) Write(w io.Writer) (int, error) { // 0x5d
	var b []byte

	b = append(b, 0x5d)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x5d,
	})
	if err != nil {
		return written, err
	}

	// 3 is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *BIT_3_L) Length() uint8 { // 0x5d
	return 2
}

func (o *BIT_3_L) cycles() []uint8 { // 0x5d
	return []uint8{8}
}

func (o *BIT_3_L) String() string { // 0x5d
	return "BIT" + " " + "3" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *BIT_3_L) SymbolicString() string { // 0x5d
	return "BIT 3,L"
}

type BIT_3_HLPtr struct {
	addr         string      // 0x5e
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_3_HLPtr) Write(w io.Writer) (int, error) { // 0x5e
	var b []byte

	b = append(b, 0x5e)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x5e,
	})
	if err != nil {
		return written, err
	}

	// 3 is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *BIT_3_HLPtr) Length() uint8 { // 0x5e
	return 2
}

func (o *BIT_3_HLPtr) cycles() []uint8 { // 0x5e
	return []uint8{16}
}

func (o *BIT_3_HLPtr) String() string { // 0x5e
	return "BIT" + " " + "3" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *BIT_3_HLPtr) SymbolicString() string { // 0x5e
	return "BIT 3,(HL)"
}

type BIT_3_A struct {
	addr         string      // 0x5f
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_3_A) Write(w io.Writer) (int, error) { // 0x5f
	var b []byte

	b = append(b, 0x5f)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x5f,
	})
	if err != nil {
		return written, err
	}

	// 3 is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *BIT_3_A) Length() uint8 { // 0x5f
	return 2
}

func (o *BIT_3_A) cycles() []uint8 { // 0x5f
	return []uint8{8}
}

func (o *BIT_3_A) String() string { // 0x5f
	return "BIT" + " " + "3" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *BIT_3_A) SymbolicString() string { // 0x5f
	return "BIT 3,A"
}

type RLC_HLPtr struct {
	addr         string      // 0x6
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RLC_HLPtr) Write(w io.Writer) (int, error) { // 0x6
	var b []byte

	b = append(b, 0x6)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x6,
	})
	if err != nil {
		return written, err
	}

	// (HL) is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RLC_HLPtr) Length() uint8 { // 0x6
	return 2
}

func (o *RLC_HLPtr) cycles() []uint8 { // 0x6
	return []uint8{16}
}

func (o *RLC_HLPtr) String() string { // 0x6
	return "RLC" + " " + "(HL)" /* operand1 */
}

func (o *RLC_HLPtr) SymbolicString() string { // 0x6
	return "RLC (HL)"
}

type BIT_4_B struct {
	addr         string      // 0x60
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_4_B) Write(w io.Writer) (int, error) { // 0x60
	var b []byte

	b = append(b, 0x60)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x60,
	})
	if err != nil {
		return written, err
	}

	// 4 is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *BIT_4_B) Length() uint8 { // 0x60
	return 2
}

func (o *BIT_4_B) cycles() []uint8 { // 0x60
	return []uint8{8}
}

func (o *BIT_4_B) String() string { // 0x60
	return "BIT" + " " + "4" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *BIT_4_B) SymbolicString() string { // 0x60
	return "BIT 4,B"
}

type BIT_4_C struct {
	addr         string      // 0x61
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_4_C) Write(w io.Writer) (int, error) { // 0x61
	var b []byte

	b = append(b, 0x61)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x61,
	})
	if err != nil {
		return written, err
	}

	// 4 is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *BIT_4_C) Length() uint8 { // 0x61
	return 2
}

func (o *BIT_4_C) cycles() []uint8 { // 0x61
	return []uint8{8}
}

func (o *BIT_4_C) String() string { // 0x61
	return "BIT" + " " + "4" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *BIT_4_C) SymbolicString() string { // 0x61
	return "BIT 4,C"
}

type BIT_4_D struct {
	addr         string      // 0x62
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_4_D) Write(w io.Writer) (int, error) { // 0x62
	var b []byte

	b = append(b, 0x62)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x62,
	})
	if err != nil {
		return written, err
	}

	// 4 is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *BIT_4_D) Length() uint8 { // 0x62
	return 2
}

func (o *BIT_4_D) cycles() []uint8 { // 0x62
	return []uint8{8}
}

func (o *BIT_4_D) String() string { // 0x62
	return "BIT" + " " + "4" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *BIT_4_D) SymbolicString() string { // 0x62
	return "BIT 4,D"
}

type BIT_4_E struct {
	addr         string      // 0x63
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_4_E) Write(w io.Writer) (int, error) { // 0x63
	var b []byte

	b = append(b, 0x63)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x63,
	})
	if err != nil {
		return written, err
	}

	// 4 is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *BIT_4_E) Length() uint8 { // 0x63
	return 2
}

func (o *BIT_4_E) cycles() []uint8 { // 0x63
	return []uint8{8}
}

func (o *BIT_4_E) String() string { // 0x63
	return "BIT" + " " + "4" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *BIT_4_E) SymbolicString() string { // 0x63
	return "BIT 4,E"
}

type BIT_4_H struct {
	addr         string      // 0x64
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_4_H) Write(w io.Writer) (int, error) { // 0x64
	var b []byte

	b = append(b, 0x64)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x64,
	})
	if err != nil {
		return written, err
	}

	// 4 is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *BIT_4_H) Length() uint8 { // 0x64
	return 2
}

func (o *BIT_4_H) cycles() []uint8 { // 0x64
	return []uint8{8}
}

func (o *BIT_4_H) String() string { // 0x64
	return "BIT" + " " + "4" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *BIT_4_H) SymbolicString() string { // 0x64
	return "BIT 4,H"
}

type BIT_4_L struct {
	addr         string      // 0x65
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_4_L) Write(w io.Writer) (int, error) { // 0x65
	var b []byte

	b = append(b, 0x65)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x65,
	})
	if err != nil {
		return written, err
	}

	// 4 is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *BIT_4_L) Length() uint8 { // 0x65
	return 2
}

func (o *BIT_4_L) cycles() []uint8 { // 0x65
	return []uint8{8}
}

func (o *BIT_4_L) String() string { // 0x65
	return "BIT" + " " + "4" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *BIT_4_L) SymbolicString() string { // 0x65
	return "BIT 4,L"
}

type BIT_4_HLPtr struct {
	addr         string      // 0x66
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_4_HLPtr) Write(w io.Writer) (int, error) { // 0x66
	var b []byte

	b = append(b, 0x66)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x66,
	})
	if err != nil {
		return written, err
	}

	// 4 is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *BIT_4_HLPtr) Length() uint8 { // 0x66
	return 2
}

func (o *BIT_4_HLPtr) cycles() []uint8 { // 0x66
	return []uint8{16}
}

func (o *BIT_4_HLPtr) String() string { // 0x66
	return "BIT" + " " + "4" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *BIT_4_HLPtr) SymbolicString() string { // 0x66
	return "BIT 4,(HL)"
}

type BIT_4_A struct {
	addr         string      // 0x67
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_4_A) Write(w io.Writer) (int, error) { // 0x67
	var b []byte

	b = append(b, 0x67)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x67,
	})
	if err != nil {
		return written, err
	}

	// 4 is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *BIT_4_A) Length() uint8 { // 0x67
	return 2
}

func (o *BIT_4_A) cycles() []uint8 { // 0x67
	return []uint8{8}
}

func (o *BIT_4_A) String() string { // 0x67
	return "BIT" + " " + "4" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *BIT_4_A) SymbolicString() string { // 0x67
	return "BIT 4,A"
}

type BIT_5_B struct {
	addr         string      // 0x68
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_5_B) Write(w io.Writer) (int, error) { // 0x68
	var b []byte

	b = append(b, 0x68)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x68,
	})
	if err != nil {
		return written, err
	}

	// 5 is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *BIT_5_B) Length() uint8 { // 0x68
	return 2
}

func (o *BIT_5_B) cycles() []uint8 { // 0x68
	return []uint8{8}
}

func (o *BIT_5_B) String() string { // 0x68
	return "BIT" + " " + "5" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *BIT_5_B) SymbolicString() string { // 0x68
	return "BIT 5,B"
}

type BIT_5_C struct {
	addr         string      // 0x69
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_5_C) Write(w io.Writer) (int, error) { // 0x69
	var b []byte

	b = append(b, 0x69)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x69,
	})
	if err != nil {
		return written, err
	}

	// 5 is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *BIT_5_C) Length() uint8 { // 0x69
	return 2
}

func (o *BIT_5_C) cycles() []uint8 { // 0x69
	return []uint8{8}
}

func (o *BIT_5_C) String() string { // 0x69
	return "BIT" + " " + "5" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *BIT_5_C) SymbolicString() string { // 0x69
	return "BIT 5,C"
}

type BIT_5_D struct {
	addr         string      // 0x6a
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_5_D) Write(w io.Writer) (int, error) { // 0x6a
	var b []byte

	b = append(b, 0x6a)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x6a,
	})
	if err != nil {
		return written, err
	}

	// 5 is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *BIT_5_D) Length() uint8 { // 0x6a
	return 2
}

func (o *BIT_5_D) cycles() []uint8 { // 0x6a
	return []uint8{8}
}

func (o *BIT_5_D) String() string { // 0x6a
	return "BIT" + " " + "5" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *BIT_5_D) SymbolicString() string { // 0x6a
	return "BIT 5,D"
}

type BIT_5_E struct {
	addr         string      // 0x6b
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_5_E) Write(w io.Writer) (int, error) { // 0x6b
	var b []byte

	b = append(b, 0x6b)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x6b,
	})
	if err != nil {
		return written, err
	}

	// 5 is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *BIT_5_E) Length() uint8 { // 0x6b
	return 2
}

func (o *BIT_5_E) cycles() []uint8 { // 0x6b
	return []uint8{8}
}

func (o *BIT_5_E) String() string { // 0x6b
	return "BIT" + " " + "5" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *BIT_5_E) SymbolicString() string { // 0x6b
	return "BIT 5,E"
}

type BIT_5_H struct {
	addr         string      // 0x6c
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_5_H) Write(w io.Writer) (int, error) { // 0x6c
	var b []byte

	b = append(b, 0x6c)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x6c,
	})
	if err != nil {
		return written, err
	}

	// 5 is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *BIT_5_H) Length() uint8 { // 0x6c
	return 2
}

func (o *BIT_5_H) cycles() []uint8 { // 0x6c
	return []uint8{8}
}

func (o *BIT_5_H) String() string { // 0x6c
	return "BIT" + " " + "5" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *BIT_5_H) SymbolicString() string { // 0x6c
	return "BIT 5,H"
}

type BIT_5_L struct {
	addr         string      // 0x6d
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_5_L) Write(w io.Writer) (int, error) { // 0x6d
	var b []byte

	b = append(b, 0x6d)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x6d,
	})
	if err != nil {
		return written, err
	}

	// 5 is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *BIT_5_L) Length() uint8 { // 0x6d
	return 2
}

func (o *BIT_5_L) cycles() []uint8 { // 0x6d
	return []uint8{8}
}

func (o *BIT_5_L) String() string { // 0x6d
	return "BIT" + " " + "5" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *BIT_5_L) SymbolicString() string { // 0x6d
	return "BIT 5,L"
}

type BIT_5_HLPtr struct {
	addr         string      // 0x6e
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_5_HLPtr) Write(w io.Writer) (int, error) { // 0x6e
	var b []byte

	b = append(b, 0x6e)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x6e,
	})
	if err != nil {
		return written, err
	}

	// 5 is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *BIT_5_HLPtr) Length() uint8 { // 0x6e
	return 2
}

func (o *BIT_5_HLPtr) cycles() []uint8 { // 0x6e
	return []uint8{16}
}

func (o *BIT_5_HLPtr) String() string { // 0x6e
	return "BIT" + " " + "5" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *BIT_5_HLPtr) SymbolicString() string { // 0x6e
	return "BIT 5,(HL)"
}

type BIT_5_A struct {
	addr         string      // 0x6f
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_5_A) Write(w io.Writer) (int, error) { // 0x6f
	var b []byte

	b = append(b, 0x6f)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x6f,
	})
	if err != nil {
		return written, err
	}

	// 5 is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *BIT_5_A) Length() uint8 { // 0x6f
	return 2
}

func (o *BIT_5_A) cycles() []uint8 { // 0x6f
	return []uint8{8}
}

func (o *BIT_5_A) String() string { // 0x6f
	return "BIT" + " " + "5" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *BIT_5_A) SymbolicString() string { // 0x6f
	return "BIT 5,A"
}

type RLC_A struct {
	addr         string      // 0x7
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RLC_A) Write(w io.Writer) (int, error) { // 0x7
	var b []byte

	b = append(b, 0x7)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x7,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RLC_A) Length() uint8 { // 0x7
	return 2
}

func (o *RLC_A) cycles() []uint8 { // 0x7
	return []uint8{8}
}

func (o *RLC_A) String() string { // 0x7
	return "RLC" + " " + "A" /* operand1 */
}

func (o *RLC_A) SymbolicString() string { // 0x7
	return "RLC A"
}

type BIT_6_B struct {
	addr         string      // 0x70
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_6_B) Write(w io.Writer) (int, error) { // 0x70
	var b []byte

	b = append(b, 0x70)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x70,
	})
	if err != nil {
		return written, err
	}

	// 6 is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *BIT_6_B) Length() uint8 { // 0x70
	return 2
}

func (o *BIT_6_B) cycles() []uint8 { // 0x70
	return []uint8{8}
}

func (o *BIT_6_B) String() string { // 0x70
	return "BIT" + " " + "6" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *BIT_6_B) SymbolicString() string { // 0x70
	return "BIT 6,B"
}

type BIT_6_C struct {
	addr         string      // 0x71
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_6_C) Write(w io.Writer) (int, error) { // 0x71
	var b []byte

	b = append(b, 0x71)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x71,
	})
	if err != nil {
		return written, err
	}

	// 6 is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *BIT_6_C) Length() uint8 { // 0x71
	return 2
}

func (o *BIT_6_C) cycles() []uint8 { // 0x71
	return []uint8{8}
}

func (o *BIT_6_C) String() string { // 0x71
	return "BIT" + " " + "6" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *BIT_6_C) SymbolicString() string { // 0x71
	return "BIT 6,C"
}

type BIT_6_D struct {
	addr         string      // 0x72
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_6_D) Write(w io.Writer) (int, error) { // 0x72
	var b []byte

	b = append(b, 0x72)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x72,
	})
	if err != nil {
		return written, err
	}

	// 6 is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *BIT_6_D) Length() uint8 { // 0x72
	return 2
}

func (o *BIT_6_D) cycles() []uint8 { // 0x72
	return []uint8{8}
}

func (o *BIT_6_D) String() string { // 0x72
	return "BIT" + " " + "6" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *BIT_6_D) SymbolicString() string { // 0x72
	return "BIT 6,D"
}

type BIT_6_E struct {
	addr         string      // 0x73
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_6_E) Write(w io.Writer) (int, error) { // 0x73
	var b []byte

	b = append(b, 0x73)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x73,
	})
	if err != nil {
		return written, err
	}

	// 6 is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *BIT_6_E) Length() uint8 { // 0x73
	return 2
}

func (o *BIT_6_E) cycles() []uint8 { // 0x73
	return []uint8{8}
}

func (o *BIT_6_E) String() string { // 0x73
	return "BIT" + " " + "6" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *BIT_6_E) SymbolicString() string { // 0x73
	return "BIT 6,E"
}

type BIT_6_H struct {
	addr         string      // 0x74
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_6_H) Write(w io.Writer) (int, error) { // 0x74
	var b []byte

	b = append(b, 0x74)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x74,
	})
	if err != nil {
		return written, err
	}

	// 6 is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *BIT_6_H) Length() uint8 { // 0x74
	return 2
}

func (o *BIT_6_H) cycles() []uint8 { // 0x74
	return []uint8{8}
}

func (o *BIT_6_H) String() string { // 0x74
	return "BIT" + " " + "6" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *BIT_6_H) SymbolicString() string { // 0x74
	return "BIT 6,H"
}

type BIT_6_L struct {
	addr         string      // 0x75
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_6_L) Write(w io.Writer) (int, error) { // 0x75
	var b []byte

	b = append(b, 0x75)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x75,
	})
	if err != nil {
		return written, err
	}

	// 6 is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *BIT_6_L) Length() uint8 { // 0x75
	return 2
}

func (o *BIT_6_L) cycles() []uint8 { // 0x75
	return []uint8{8}
}

func (o *BIT_6_L) String() string { // 0x75
	return "BIT" + " " + "6" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *BIT_6_L) SymbolicString() string { // 0x75
	return "BIT 6,L"
}

type BIT_6_HLPtr struct {
	addr         string      // 0x76
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_6_HLPtr) Write(w io.Writer) (int, error) { // 0x76
	var b []byte

	b = append(b, 0x76)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x76,
	})
	if err != nil {
		return written, err
	}

	// 6 is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *BIT_6_HLPtr) Length() uint8 { // 0x76
	return 2
}

func (o *BIT_6_HLPtr) cycles() []uint8 { // 0x76
	return []uint8{16}
}

func (o *BIT_6_HLPtr) String() string { // 0x76
	return "BIT" + " " + "6" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *BIT_6_HLPtr) SymbolicString() string { // 0x76
	return "BIT 6,(HL)"
}

type BIT_6_A struct {
	addr         string      // 0x77
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_6_A) Write(w io.Writer) (int, error) { // 0x77
	var b []byte

	b = append(b, 0x77)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x77,
	})
	if err != nil {
		return written, err
	}

	// 6 is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *BIT_6_A) Length() uint8 { // 0x77
	return 2
}

func (o *BIT_6_A) cycles() []uint8 { // 0x77
	return []uint8{8}
}

func (o *BIT_6_A) String() string { // 0x77
	return "BIT" + " " + "6" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *BIT_6_A) SymbolicString() string { // 0x77
	return "BIT 6,A"
}

type BIT_7_B struct {
	addr         string      // 0x78
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_7_B) Write(w io.Writer) (int, error) { // 0x78
	var b []byte

	b = append(b, 0x78)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x78,
	})
	if err != nil {
		return written, err
	}

	// 7 is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *BIT_7_B) Length() uint8 { // 0x78
	return 2
}

func (o *BIT_7_B) cycles() []uint8 { // 0x78
	return []uint8{8}
}

func (o *BIT_7_B) String() string { // 0x78
	return "BIT" + " " + "7" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *BIT_7_B) SymbolicString() string { // 0x78
	return "BIT 7,B"
}

type BIT_7_C struct {
	addr         string      // 0x79
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_7_C) Write(w io.Writer) (int, error) { // 0x79
	var b []byte

	b = append(b, 0x79)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x79,
	})
	if err != nil {
		return written, err
	}

	// 7 is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *BIT_7_C) Length() uint8 { // 0x79
	return 2
}

func (o *BIT_7_C) cycles() []uint8 { // 0x79
	return []uint8{8}
}

func (o *BIT_7_C) String() string { // 0x79
	return "BIT" + " " + "7" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *BIT_7_C) SymbolicString() string { // 0x79
	return "BIT 7,C"
}

type BIT_7_D struct {
	addr         string      // 0x7a
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_7_D) Write(w io.Writer) (int, error) { // 0x7a
	var b []byte

	b = append(b, 0x7a)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x7a,
	})
	if err != nil {
		return written, err
	}

	// 7 is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *BIT_7_D) Length() uint8 { // 0x7a
	return 2
}

func (o *BIT_7_D) cycles() []uint8 { // 0x7a
	return []uint8{8}
}

func (o *BIT_7_D) String() string { // 0x7a
	return "BIT" + " " + "7" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *BIT_7_D) SymbolicString() string { // 0x7a
	return "BIT 7,D"
}

type BIT_7_E struct {
	addr         string      // 0x7b
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_7_E) Write(w io.Writer) (int, error) { // 0x7b
	var b []byte

	b = append(b, 0x7b)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x7b,
	})
	if err != nil {
		return written, err
	}

	// 7 is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *BIT_7_E) Length() uint8 { // 0x7b
	return 2
}

func (o *BIT_7_E) cycles() []uint8 { // 0x7b
	return []uint8{8}
}

func (o *BIT_7_E) String() string { // 0x7b
	return "BIT" + " " + "7" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *BIT_7_E) SymbolicString() string { // 0x7b
	return "BIT 7,E"
}

type BIT_7_H struct {
	addr         string      // 0x7c
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_7_H) Write(w io.Writer) (int, error) { // 0x7c
	var b []byte

	b = append(b, 0x7c)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x7c,
	})
	if err != nil {
		return written, err
	}

	// 7 is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *BIT_7_H) Length() uint8 { // 0x7c
	return 2
}

func (o *BIT_7_H) cycles() []uint8 { // 0x7c
	return []uint8{8}
}

func (o *BIT_7_H) String() string { // 0x7c
	return "BIT" + " " + "7" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *BIT_7_H) SymbolicString() string { // 0x7c
	return "BIT 7,H"
}

type BIT_7_L struct {
	addr         string      // 0x7d
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_7_L) Write(w io.Writer) (int, error) { // 0x7d
	var b []byte

	b = append(b, 0x7d)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x7d,
	})
	if err != nil {
		return written, err
	}

	// 7 is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *BIT_7_L) Length() uint8 { // 0x7d
	return 2
}

func (o *BIT_7_L) cycles() []uint8 { // 0x7d
	return []uint8{8}
}

func (o *BIT_7_L) String() string { // 0x7d
	return "BIT" + " " + "7" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *BIT_7_L) SymbolicString() string { // 0x7d
	return "BIT 7,L"
}

type BIT_7_HLPtr struct {
	addr         string      // 0x7e
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_7_HLPtr) Write(w io.Writer) (int, error) { // 0x7e
	var b []byte

	b = append(b, 0x7e)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x7e,
	})
	if err != nil {
		return written, err
	}

	// 7 is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *BIT_7_HLPtr) Length() uint8 { // 0x7e
	return 2
}

func (o *BIT_7_HLPtr) cycles() []uint8 { // 0x7e
	return []uint8{16}
}

func (o *BIT_7_HLPtr) String() string { // 0x7e
	return "BIT" + " " + "7" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *BIT_7_HLPtr) SymbolicString() string { // 0x7e
	return "BIT 7,(HL)"
}

type BIT_7_A struct {
	addr         string      // 0x7f
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *BIT_7_A) Write(w io.Writer) (int, error) { // 0x7f
	var b []byte

	b = append(b, 0x7f)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x7f,
	})
	if err != nil {
		return written, err
	}

	// 7 is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *BIT_7_A) Length() uint8 { // 0x7f
	return 2
}

func (o *BIT_7_A) cycles() []uint8 { // 0x7f
	return []uint8{8}
}

func (o *BIT_7_A) String() string { // 0x7f
	return "BIT" + " " + "7" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *BIT_7_A) SymbolicString() string { // 0x7f
	return "BIT 7,A"
}

type RRC_B struct {
	addr         string      // 0x8
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RRC_B) Write(w io.Writer) (int, error) { // 0x8
	var b []byte

	b = append(b, 0x8)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x8,
	})
	if err != nil {
		return written, err
	}

	// B is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RRC_B) Length() uint8 { // 0x8
	return 2
}

func (o *RRC_B) cycles() []uint8 { // 0x8
	return []uint8{8}
}

func (o *RRC_B) String() string { // 0x8
	return "RRC" + " " + "B" /* operand1 */
}

func (o *RRC_B) SymbolicString() string { // 0x8
	return "RRC B"
}

type RES_0_B struct {
	addr         string      // 0x80
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_0_B) Write(w io.Writer) (int, error) { // 0x80
	var b []byte

	b = append(b, 0x80)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x80,
	})
	if err != nil {
		return written, err
	}

	// 0 is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *RES_0_B) Length() uint8 { // 0x80
	return 2
}

func (o *RES_0_B) cycles() []uint8 { // 0x80
	return []uint8{8}
}

func (o *RES_0_B) String() string { // 0x80
	return "RES" + " " + "0" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *RES_0_B) SymbolicString() string { // 0x80
	return "RES 0,B"
}

type RES_0_C struct {
	addr         string      // 0x81
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_0_C) Write(w io.Writer) (int, error) { // 0x81
	var b []byte

	b = append(b, 0x81)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x81,
	})
	if err != nil {
		return written, err
	}

	// 0 is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *RES_0_C) Length() uint8 { // 0x81
	return 2
}

func (o *RES_0_C) cycles() []uint8 { // 0x81
	return []uint8{8}
}

func (o *RES_0_C) String() string { // 0x81
	return "RES" + " " + "0" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *RES_0_C) SymbolicString() string { // 0x81
	return "RES 0,C"
}

type RES_0_D struct {
	addr         string      // 0x82
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_0_D) Write(w io.Writer) (int, error) { // 0x82
	var b []byte

	b = append(b, 0x82)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x82,
	})
	if err != nil {
		return written, err
	}

	// 0 is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *RES_0_D) Length() uint8 { // 0x82
	return 2
}

func (o *RES_0_D) cycles() []uint8 { // 0x82
	return []uint8{8}
}

func (o *RES_0_D) String() string { // 0x82
	return "RES" + " " + "0" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *RES_0_D) SymbolicString() string { // 0x82
	return "RES 0,D"
}

type RES_0_E struct {
	addr         string      // 0x83
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_0_E) Write(w io.Writer) (int, error) { // 0x83
	var b []byte

	b = append(b, 0x83)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x83,
	})
	if err != nil {
		return written, err
	}

	// 0 is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *RES_0_E) Length() uint8 { // 0x83
	return 2
}

func (o *RES_0_E) cycles() []uint8 { // 0x83
	return []uint8{8}
}

func (o *RES_0_E) String() string { // 0x83
	return "RES" + " " + "0" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *RES_0_E) SymbolicString() string { // 0x83
	return "RES 0,E"
}

type RES_0_H struct {
	addr         string      // 0x84
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_0_H) Write(w io.Writer) (int, error) { // 0x84
	var b []byte

	b = append(b, 0x84)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x84,
	})
	if err != nil {
		return written, err
	}

	// 0 is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *RES_0_H) Length() uint8 { // 0x84
	return 2
}

func (o *RES_0_H) cycles() []uint8 { // 0x84
	return []uint8{8}
}

func (o *RES_0_H) String() string { // 0x84
	return "RES" + " " + "0" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *RES_0_H) SymbolicString() string { // 0x84
	return "RES 0,H"
}

type RES_0_L struct {
	addr         string      // 0x85
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_0_L) Write(w io.Writer) (int, error) { // 0x85
	var b []byte

	b = append(b, 0x85)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x85,
	})
	if err != nil {
		return written, err
	}

	// 0 is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *RES_0_L) Length() uint8 { // 0x85
	return 2
}

func (o *RES_0_L) cycles() []uint8 { // 0x85
	return []uint8{8}
}

func (o *RES_0_L) String() string { // 0x85
	return "RES" + " " + "0" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *RES_0_L) SymbolicString() string { // 0x85
	return "RES 0,L"
}

type RES_0_HLPtr struct {
	addr         string      // 0x86
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_0_HLPtr) Write(w io.Writer) (int, error) { // 0x86
	var b []byte

	b = append(b, 0x86)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x86,
	})
	if err != nil {
		return written, err
	}

	// 0 is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *RES_0_HLPtr) Length() uint8 { // 0x86
	return 2
}

func (o *RES_0_HLPtr) cycles() []uint8 { // 0x86
	return []uint8{16}
}

func (o *RES_0_HLPtr) String() string { // 0x86
	return "RES" + " " + "0" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *RES_0_HLPtr) SymbolicString() string { // 0x86
	return "RES 0,(HL)"
}

type RES_0_A struct {
	addr         string      // 0x87
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_0_A) Write(w io.Writer) (int, error) { // 0x87
	var b []byte

	b = append(b, 0x87)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x87,
	})
	if err != nil {
		return written, err
	}

	// 0 is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *RES_0_A) Length() uint8 { // 0x87
	return 2
}

func (o *RES_0_A) cycles() []uint8 { // 0x87
	return []uint8{8}
}

func (o *RES_0_A) String() string { // 0x87
	return "RES" + " " + "0" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *RES_0_A) SymbolicString() string { // 0x87
	return "RES 0,A"
}

type RES_1_B struct {
	addr         string      // 0x88
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_1_B) Write(w io.Writer) (int, error) { // 0x88
	var b []byte

	b = append(b, 0x88)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x88,
	})
	if err != nil {
		return written, err
	}

	// 1 is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *RES_1_B) Length() uint8 { // 0x88
	return 2
}

func (o *RES_1_B) cycles() []uint8 { // 0x88
	return []uint8{8}
}

func (o *RES_1_B) String() string { // 0x88
	return "RES" + " " + "1" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *RES_1_B) SymbolicString() string { // 0x88
	return "RES 1,B"
}

type RES_1_C struct {
	addr         string      // 0x89
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_1_C) Write(w io.Writer) (int, error) { // 0x89
	var b []byte

	b = append(b, 0x89)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x89,
	})
	if err != nil {
		return written, err
	}

	// 1 is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *RES_1_C) Length() uint8 { // 0x89
	return 2
}

func (o *RES_1_C) cycles() []uint8 { // 0x89
	return []uint8{8}
}

func (o *RES_1_C) String() string { // 0x89
	return "RES" + " " + "1" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *RES_1_C) SymbolicString() string { // 0x89
	return "RES 1,C"
}

type RES_1_D struct {
	addr         string      // 0x8a
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_1_D) Write(w io.Writer) (int, error) { // 0x8a
	var b []byte

	b = append(b, 0x8a)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x8a,
	})
	if err != nil {
		return written, err
	}

	// 1 is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *RES_1_D) Length() uint8 { // 0x8a
	return 2
}

func (o *RES_1_D) cycles() []uint8 { // 0x8a
	return []uint8{8}
}

func (o *RES_1_D) String() string { // 0x8a
	return "RES" + " " + "1" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *RES_1_D) SymbolicString() string { // 0x8a
	return "RES 1,D"
}

type RES_1_E struct {
	addr         string      // 0x8b
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_1_E) Write(w io.Writer) (int, error) { // 0x8b
	var b []byte

	b = append(b, 0x8b)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x8b,
	})
	if err != nil {
		return written, err
	}

	// 1 is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *RES_1_E) Length() uint8 { // 0x8b
	return 2
}

func (o *RES_1_E) cycles() []uint8 { // 0x8b
	return []uint8{8}
}

func (o *RES_1_E) String() string { // 0x8b
	return "RES" + " " + "1" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *RES_1_E) SymbolicString() string { // 0x8b
	return "RES 1,E"
}

type RES_1_H struct {
	addr         string      // 0x8c
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_1_H) Write(w io.Writer) (int, error) { // 0x8c
	var b []byte

	b = append(b, 0x8c)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x8c,
	})
	if err != nil {
		return written, err
	}

	// 1 is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *RES_1_H) Length() uint8 { // 0x8c
	return 2
}

func (o *RES_1_H) cycles() []uint8 { // 0x8c
	return []uint8{8}
}

func (o *RES_1_H) String() string { // 0x8c
	return "RES" + " " + "1" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *RES_1_H) SymbolicString() string { // 0x8c
	return "RES 1,H"
}

type RES_1_L struct {
	addr         string      // 0x8d
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_1_L) Write(w io.Writer) (int, error) { // 0x8d
	var b []byte

	b = append(b, 0x8d)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x8d,
	})
	if err != nil {
		return written, err
	}

	// 1 is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *RES_1_L) Length() uint8 { // 0x8d
	return 2
}

func (o *RES_1_L) cycles() []uint8 { // 0x8d
	return []uint8{8}
}

func (o *RES_1_L) String() string { // 0x8d
	return "RES" + " " + "1" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *RES_1_L) SymbolicString() string { // 0x8d
	return "RES 1,L"
}

type RES_1_HLPtr struct {
	addr         string      // 0x8e
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_1_HLPtr) Write(w io.Writer) (int, error) { // 0x8e
	var b []byte

	b = append(b, 0x8e)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x8e,
	})
	if err != nil {
		return written, err
	}

	// 1 is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *RES_1_HLPtr) Length() uint8 { // 0x8e
	return 2
}

func (o *RES_1_HLPtr) cycles() []uint8 { // 0x8e
	return []uint8{16}
}

func (o *RES_1_HLPtr) String() string { // 0x8e
	return "RES" + " " + "1" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *RES_1_HLPtr) SymbolicString() string { // 0x8e
	return "RES 1,(HL)"
}

type RES_1_A struct {
	addr         string      // 0x8f
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_1_A) Write(w io.Writer) (int, error) { // 0x8f
	var b []byte

	b = append(b, 0x8f)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x8f,
	})
	if err != nil {
		return written, err
	}

	// 1 is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *RES_1_A) Length() uint8 { // 0x8f
	return 2
}

func (o *RES_1_A) cycles() []uint8 { // 0x8f
	return []uint8{8}
}

func (o *RES_1_A) String() string { // 0x8f
	return "RES" + " " + "1" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *RES_1_A) SymbolicString() string { // 0x8f
	return "RES 1,A"
}

type RRC_C struct {
	addr         string      // 0x9
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RRC_C) Write(w io.Writer) (int, error) { // 0x9
	var b []byte

	b = append(b, 0x9)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x9,
	})
	if err != nil {
		return written, err
	}

	// C is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RRC_C) Length() uint8 { // 0x9
	return 2
}

func (o *RRC_C) cycles() []uint8 { // 0x9
	return []uint8{8}
}

func (o *RRC_C) String() string { // 0x9
	return "RRC" + " " + "C" /* operand1 */
}

func (o *RRC_C) SymbolicString() string { // 0x9
	return "RRC C"
}

type RES_2_B struct {
	addr         string      // 0x90
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_2_B) Write(w io.Writer) (int, error) { // 0x90
	var b []byte

	b = append(b, 0x90)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x90,
	})
	if err != nil {
		return written, err
	}

	// 2 is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *RES_2_B) Length() uint8 { // 0x90
	return 2
}

func (o *RES_2_B) cycles() []uint8 { // 0x90
	return []uint8{8}
}

func (o *RES_2_B) String() string { // 0x90
	return "RES" + " " + "2" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *RES_2_B) SymbolicString() string { // 0x90
	return "RES 2,B"
}

type RES_2_C struct {
	addr         string      // 0x91
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_2_C) Write(w io.Writer) (int, error) { // 0x91
	var b []byte

	b = append(b, 0x91)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x91,
	})
	if err != nil {
		return written, err
	}

	// 2 is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *RES_2_C) Length() uint8 { // 0x91
	return 2
}

func (o *RES_2_C) cycles() []uint8 { // 0x91
	return []uint8{8}
}

func (o *RES_2_C) String() string { // 0x91
	return "RES" + " " + "2" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *RES_2_C) SymbolicString() string { // 0x91
	return "RES 2,C"
}

type RES_2_D struct {
	addr         string      // 0x92
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_2_D) Write(w io.Writer) (int, error) { // 0x92
	var b []byte

	b = append(b, 0x92)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x92,
	})
	if err != nil {
		return written, err
	}

	// 2 is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *RES_2_D) Length() uint8 { // 0x92
	return 2
}

func (o *RES_2_D) cycles() []uint8 { // 0x92
	return []uint8{8}
}

func (o *RES_2_D) String() string { // 0x92
	return "RES" + " " + "2" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *RES_2_D) SymbolicString() string { // 0x92
	return "RES 2,D"
}

type RES_2_E struct {
	addr         string      // 0x93
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_2_E) Write(w io.Writer) (int, error) { // 0x93
	var b []byte

	b = append(b, 0x93)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x93,
	})
	if err != nil {
		return written, err
	}

	// 2 is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *RES_2_E) Length() uint8 { // 0x93
	return 2
}

func (o *RES_2_E) cycles() []uint8 { // 0x93
	return []uint8{8}
}

func (o *RES_2_E) String() string { // 0x93
	return "RES" + " " + "2" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *RES_2_E) SymbolicString() string { // 0x93
	return "RES 2,E"
}

type RES_2_H struct {
	addr         string      // 0x94
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_2_H) Write(w io.Writer) (int, error) { // 0x94
	var b []byte

	b = append(b, 0x94)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x94,
	})
	if err != nil {
		return written, err
	}

	// 2 is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *RES_2_H) Length() uint8 { // 0x94
	return 2
}

func (o *RES_2_H) cycles() []uint8 { // 0x94
	return []uint8{8}
}

func (o *RES_2_H) String() string { // 0x94
	return "RES" + " " + "2" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *RES_2_H) SymbolicString() string { // 0x94
	return "RES 2,H"
}

type RES_2_L struct {
	addr         string      // 0x95
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_2_L) Write(w io.Writer) (int, error) { // 0x95
	var b []byte

	b = append(b, 0x95)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x95,
	})
	if err != nil {
		return written, err
	}

	// 2 is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *RES_2_L) Length() uint8 { // 0x95
	return 2
}

func (o *RES_2_L) cycles() []uint8 { // 0x95
	return []uint8{8}
}

func (o *RES_2_L) String() string { // 0x95
	return "RES" + " " + "2" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *RES_2_L) SymbolicString() string { // 0x95
	return "RES 2,L"
}

type RES_2_HLPtr struct {
	addr         string      // 0x96
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_2_HLPtr) Write(w io.Writer) (int, error) { // 0x96
	var b []byte

	b = append(b, 0x96)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x96,
	})
	if err != nil {
		return written, err
	}

	// 2 is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *RES_2_HLPtr) Length() uint8 { // 0x96
	return 2
}

func (o *RES_2_HLPtr) cycles() []uint8 { // 0x96
	return []uint8{16}
}

func (o *RES_2_HLPtr) String() string { // 0x96
	return "RES" + " " + "2" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *RES_2_HLPtr) SymbolicString() string { // 0x96
	return "RES 2,(HL)"
}

type RES_2_A struct {
	addr         string      // 0x97
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_2_A) Write(w io.Writer) (int, error) { // 0x97
	var b []byte

	b = append(b, 0x97)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x97,
	})
	if err != nil {
		return written, err
	}

	// 2 is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *RES_2_A) Length() uint8 { // 0x97
	return 2
}

func (o *RES_2_A) cycles() []uint8 { // 0x97
	return []uint8{8}
}

func (o *RES_2_A) String() string { // 0x97
	return "RES" + " " + "2" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *RES_2_A) SymbolicString() string { // 0x97
	return "RES 2,A"
}

type RES_3_B struct {
	addr         string      // 0x98
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_3_B) Write(w io.Writer) (int, error) { // 0x98
	var b []byte

	b = append(b, 0x98)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x98,
	})
	if err != nil {
		return written, err
	}

	// 3 is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *RES_3_B) Length() uint8 { // 0x98
	return 2
}

func (o *RES_3_B) cycles() []uint8 { // 0x98
	return []uint8{8}
}

func (o *RES_3_B) String() string { // 0x98
	return "RES" + " " + "3" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *RES_3_B) SymbolicString() string { // 0x98
	return "RES 3,B"
}

type RES_3_C struct {
	addr         string      // 0x99
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_3_C) Write(w io.Writer) (int, error) { // 0x99
	var b []byte

	b = append(b, 0x99)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x99,
	})
	if err != nil {
		return written, err
	}

	// 3 is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *RES_3_C) Length() uint8 { // 0x99
	return 2
}

func (o *RES_3_C) cycles() []uint8 { // 0x99
	return []uint8{8}
}

func (o *RES_3_C) String() string { // 0x99
	return "RES" + " " + "3" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *RES_3_C) SymbolicString() string { // 0x99
	return "RES 3,C"
}

type RES_3_D struct {
	addr         string      // 0x9a
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_3_D) Write(w io.Writer) (int, error) { // 0x9a
	var b []byte

	b = append(b, 0x9a)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x9a,
	})
	if err != nil {
		return written, err
	}

	// 3 is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *RES_3_D) Length() uint8 { // 0x9a
	return 2
}

func (o *RES_3_D) cycles() []uint8 { // 0x9a
	return []uint8{8}
}

func (o *RES_3_D) String() string { // 0x9a
	return "RES" + " " + "3" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *RES_3_D) SymbolicString() string { // 0x9a
	return "RES 3,D"
}

type RES_3_E struct {
	addr         string      // 0x9b
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_3_E) Write(w io.Writer) (int, error) { // 0x9b
	var b []byte

	b = append(b, 0x9b)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x9b,
	})
	if err != nil {
		return written, err
	}

	// 3 is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *RES_3_E) Length() uint8 { // 0x9b
	return 2
}

func (o *RES_3_E) cycles() []uint8 { // 0x9b
	return []uint8{8}
}

func (o *RES_3_E) String() string { // 0x9b
	return "RES" + " " + "3" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *RES_3_E) SymbolicString() string { // 0x9b
	return "RES 3,E"
}

type RES_3_H struct {
	addr         string      // 0x9c
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_3_H) Write(w io.Writer) (int, error) { // 0x9c
	var b []byte

	b = append(b, 0x9c)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x9c,
	})
	if err != nil {
		return written, err
	}

	// 3 is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *RES_3_H) Length() uint8 { // 0x9c
	return 2
}

func (o *RES_3_H) cycles() []uint8 { // 0x9c
	return []uint8{8}
}

func (o *RES_3_H) String() string { // 0x9c
	return "RES" + " " + "3" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *RES_3_H) SymbolicString() string { // 0x9c
	return "RES 3,H"
}

type RES_3_L struct {
	addr         string      // 0x9d
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_3_L) Write(w io.Writer) (int, error) { // 0x9d
	var b []byte

	b = append(b, 0x9d)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x9d,
	})
	if err != nil {
		return written, err
	}

	// 3 is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *RES_3_L) Length() uint8 { // 0x9d
	return 2
}

func (o *RES_3_L) cycles() []uint8 { // 0x9d
	return []uint8{8}
}

func (o *RES_3_L) String() string { // 0x9d
	return "RES" + " " + "3" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *RES_3_L) SymbolicString() string { // 0x9d
	return "RES 3,L"
}

type RES_3_HLPtr struct {
	addr         string      // 0x9e
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_3_HLPtr) Write(w io.Writer) (int, error) { // 0x9e
	var b []byte

	b = append(b, 0x9e)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x9e,
	})
	if err != nil {
		return written, err
	}

	// 3 is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *RES_3_HLPtr) Length() uint8 { // 0x9e
	return 2
}

func (o *RES_3_HLPtr) cycles() []uint8 { // 0x9e
	return []uint8{16}
}

func (o *RES_3_HLPtr) String() string { // 0x9e
	return "RES" + " " + "3" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *RES_3_HLPtr) SymbolicString() string { // 0x9e
	return "RES 3,(HL)"
}

type RES_3_A struct {
	addr         string      // 0x9f
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_3_A) Write(w io.Writer) (int, error) { // 0x9f
	var b []byte

	b = append(b, 0x9f)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0x9f,
	})
	if err != nil {
		return written, err
	}

	// 3 is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *RES_3_A) Length() uint8 { // 0x9f
	return 2
}

func (o *RES_3_A) cycles() []uint8 { // 0x9f
	return []uint8{8}
}

func (o *RES_3_A) String() string { // 0x9f
	return "RES" + " " + "3" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *RES_3_A) SymbolicString() string { // 0x9f
	return "RES 3,A"
}

type RRC_D struct {
	addr         string      // 0xa
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RRC_D) Write(w io.Writer) (int, error) { // 0xa
	var b []byte

	b = append(b, 0xa)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xa,
	})
	if err != nil {
		return written, err
	}

	// D is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RRC_D) Length() uint8 { // 0xa
	return 2
}

func (o *RRC_D) cycles() []uint8 { // 0xa
	return []uint8{8}
}

func (o *RRC_D) String() string { // 0xa
	return "RRC" + " " + "D" /* operand1 */
}

func (o *RRC_D) SymbolicString() string { // 0xa
	return "RRC D"
}

type RES_4_B struct {
	addr         string      // 0xa0
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_4_B) Write(w io.Writer) (int, error) { // 0xa0
	var b []byte

	b = append(b, 0xa0)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xa0,
	})
	if err != nil {
		return written, err
	}

	// 4 is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *RES_4_B) Length() uint8 { // 0xa0
	return 2
}

func (o *RES_4_B) cycles() []uint8 { // 0xa0
	return []uint8{8}
}

func (o *RES_4_B) String() string { // 0xa0
	return "RES" + " " + "4" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *RES_4_B) SymbolicString() string { // 0xa0
	return "RES 4,B"
}

type RES_4_C struct {
	addr         string      // 0xa1
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_4_C) Write(w io.Writer) (int, error) { // 0xa1
	var b []byte

	b = append(b, 0xa1)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xa1,
	})
	if err != nil {
		return written, err
	}

	// 4 is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *RES_4_C) Length() uint8 { // 0xa1
	return 2
}

func (o *RES_4_C) cycles() []uint8 { // 0xa1
	return []uint8{8}
}

func (o *RES_4_C) String() string { // 0xa1
	return "RES" + " " + "4" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *RES_4_C) SymbolicString() string { // 0xa1
	return "RES 4,C"
}

type RES_4_D struct {
	addr         string      // 0xa2
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_4_D) Write(w io.Writer) (int, error) { // 0xa2
	var b []byte

	b = append(b, 0xa2)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xa2,
	})
	if err != nil {
		return written, err
	}

	// 4 is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *RES_4_D) Length() uint8 { // 0xa2
	return 2
}

func (o *RES_4_D) cycles() []uint8 { // 0xa2
	return []uint8{8}
}

func (o *RES_4_D) String() string { // 0xa2
	return "RES" + " " + "4" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *RES_4_D) SymbolicString() string { // 0xa2
	return "RES 4,D"
}

type RES_4_E struct {
	addr         string      // 0xa3
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_4_E) Write(w io.Writer) (int, error) { // 0xa3
	var b []byte

	b = append(b, 0xa3)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xa3,
	})
	if err != nil {
		return written, err
	}

	// 4 is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *RES_4_E) Length() uint8 { // 0xa3
	return 2
}

func (o *RES_4_E) cycles() []uint8 { // 0xa3
	return []uint8{8}
}

func (o *RES_4_E) String() string { // 0xa3
	return "RES" + " " + "4" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *RES_4_E) SymbolicString() string { // 0xa3
	return "RES 4,E"
}

type RES_4_H struct {
	addr         string      // 0xa4
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_4_H) Write(w io.Writer) (int, error) { // 0xa4
	var b []byte

	b = append(b, 0xa4)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xa4,
	})
	if err != nil {
		return written, err
	}

	// 4 is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *RES_4_H) Length() uint8 { // 0xa4
	return 2
}

func (o *RES_4_H) cycles() []uint8 { // 0xa4
	return []uint8{8}
}

func (o *RES_4_H) String() string { // 0xa4
	return "RES" + " " + "4" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *RES_4_H) SymbolicString() string { // 0xa4
	return "RES 4,H"
}

type RES_4_L struct {
	addr         string      // 0xa5
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_4_L) Write(w io.Writer) (int, error) { // 0xa5
	var b []byte

	b = append(b, 0xa5)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xa5,
	})
	if err != nil {
		return written, err
	}

	// 4 is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *RES_4_L) Length() uint8 { // 0xa5
	return 2
}

func (o *RES_4_L) cycles() []uint8 { // 0xa5
	return []uint8{8}
}

func (o *RES_4_L) String() string { // 0xa5
	return "RES" + " " + "4" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *RES_4_L) SymbolicString() string { // 0xa5
	return "RES 4,L"
}

type RES_4_HLPtr struct {
	addr         string      // 0xa6
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_4_HLPtr) Write(w io.Writer) (int, error) { // 0xa6
	var b []byte

	b = append(b, 0xa6)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xa6,
	})
	if err != nil {
		return written, err
	}

	// 4 is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *RES_4_HLPtr) Length() uint8 { // 0xa6
	return 2
}

func (o *RES_4_HLPtr) cycles() []uint8 { // 0xa6
	return []uint8{16}
}

func (o *RES_4_HLPtr) String() string { // 0xa6
	return "RES" + " " + "4" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *RES_4_HLPtr) SymbolicString() string { // 0xa6
	return "RES 4,(HL)"
}

type RES_4_A struct {
	addr         string      // 0xa7
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_4_A) Write(w io.Writer) (int, error) { // 0xa7
	var b []byte

	b = append(b, 0xa7)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xa7,
	})
	if err != nil {
		return written, err
	}

	// 4 is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *RES_4_A) Length() uint8 { // 0xa7
	return 2
}

func (o *RES_4_A) cycles() []uint8 { // 0xa7
	return []uint8{8}
}

func (o *RES_4_A) String() string { // 0xa7
	return "RES" + " " + "4" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *RES_4_A) SymbolicString() string { // 0xa7
	return "RES 4,A"
}

type RES_5_B struct {
	addr         string      // 0xa8
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_5_B) Write(w io.Writer) (int, error) { // 0xa8
	var b []byte

	b = append(b, 0xa8)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xa8,
	})
	if err != nil {
		return written, err
	}

	// 5 is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *RES_5_B) Length() uint8 { // 0xa8
	return 2
}

func (o *RES_5_B) cycles() []uint8 { // 0xa8
	return []uint8{8}
}

func (o *RES_5_B) String() string { // 0xa8
	return "RES" + " " + "5" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *RES_5_B) SymbolicString() string { // 0xa8
	return "RES 5,B"
}

type RES_5_C struct {
	addr         string      // 0xa9
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_5_C) Write(w io.Writer) (int, error) { // 0xa9
	var b []byte

	b = append(b, 0xa9)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xa9,
	})
	if err != nil {
		return written, err
	}

	// 5 is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *RES_5_C) Length() uint8 { // 0xa9
	return 2
}

func (o *RES_5_C) cycles() []uint8 { // 0xa9
	return []uint8{8}
}

func (o *RES_5_C) String() string { // 0xa9
	return "RES" + " " + "5" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *RES_5_C) SymbolicString() string { // 0xa9
	return "RES 5,C"
}

type RES_5_D struct {
	addr         string      // 0xaa
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_5_D) Write(w io.Writer) (int, error) { // 0xaa
	var b []byte

	b = append(b, 0xaa)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xaa,
	})
	if err != nil {
		return written, err
	}

	// 5 is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *RES_5_D) Length() uint8 { // 0xaa
	return 2
}

func (o *RES_5_D) cycles() []uint8 { // 0xaa
	return []uint8{8}
}

func (o *RES_5_D) String() string { // 0xaa
	return "RES" + " " + "5" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *RES_5_D) SymbolicString() string { // 0xaa
	return "RES 5,D"
}

type RES_5_E struct {
	addr         string      // 0xab
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_5_E) Write(w io.Writer) (int, error) { // 0xab
	var b []byte

	b = append(b, 0xab)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xab,
	})
	if err != nil {
		return written, err
	}

	// 5 is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *RES_5_E) Length() uint8 { // 0xab
	return 2
}

func (o *RES_5_E) cycles() []uint8 { // 0xab
	return []uint8{8}
}

func (o *RES_5_E) String() string { // 0xab
	return "RES" + " " + "5" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *RES_5_E) SymbolicString() string { // 0xab
	return "RES 5,E"
}

type RES_5_H struct {
	addr         string      // 0xac
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_5_H) Write(w io.Writer) (int, error) { // 0xac
	var b []byte

	b = append(b, 0xac)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xac,
	})
	if err != nil {
		return written, err
	}

	// 5 is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *RES_5_H) Length() uint8 { // 0xac
	return 2
}

func (o *RES_5_H) cycles() []uint8 { // 0xac
	return []uint8{8}
}

func (o *RES_5_H) String() string { // 0xac
	return "RES" + " " + "5" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *RES_5_H) SymbolicString() string { // 0xac
	return "RES 5,H"
}

type RES_5_L struct {
	addr         string      // 0xad
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_5_L) Write(w io.Writer) (int, error) { // 0xad
	var b []byte

	b = append(b, 0xad)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xad,
	})
	if err != nil {
		return written, err
	}

	// 5 is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *RES_5_L) Length() uint8 { // 0xad
	return 2
}

func (o *RES_5_L) cycles() []uint8 { // 0xad
	return []uint8{8}
}

func (o *RES_5_L) String() string { // 0xad
	return "RES" + " " + "5" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *RES_5_L) SymbolicString() string { // 0xad
	return "RES 5,L"
}

type RES_5_HLPtr struct {
	addr         string      // 0xae
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_5_HLPtr) Write(w io.Writer) (int, error) { // 0xae
	var b []byte

	b = append(b, 0xae)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xae,
	})
	if err != nil {
		return written, err
	}

	// 5 is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *RES_5_HLPtr) Length() uint8 { // 0xae
	return 2
}

func (o *RES_5_HLPtr) cycles() []uint8 { // 0xae
	return []uint8{16}
}

func (o *RES_5_HLPtr) String() string { // 0xae
	return "RES" + " " + "5" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *RES_5_HLPtr) SymbolicString() string { // 0xae
	return "RES 5,(HL)"
}

type RES_5_A struct {
	addr         string      // 0xaf
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_5_A) Write(w io.Writer) (int, error) { // 0xaf
	var b []byte

	b = append(b, 0xaf)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xaf,
	})
	if err != nil {
		return written, err
	}

	// 5 is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *RES_5_A) Length() uint8 { // 0xaf
	return 2
}

func (o *RES_5_A) cycles() []uint8 { // 0xaf
	return []uint8{8}
}

func (o *RES_5_A) String() string { // 0xaf
	return "RES" + " " + "5" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *RES_5_A) SymbolicString() string { // 0xaf
	return "RES 5,A"
}

type RRC_E struct {
	addr         string      // 0xb
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RRC_E) Write(w io.Writer) (int, error) { // 0xb
	var b []byte

	b = append(b, 0xb)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xb,
	})
	if err != nil {
		return written, err
	}

	// E is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RRC_E) Length() uint8 { // 0xb
	return 2
}

func (o *RRC_E) cycles() []uint8 { // 0xb
	return []uint8{8}
}

func (o *RRC_E) String() string { // 0xb
	return "RRC" + " " + "E" /* operand1 */
}

func (o *RRC_E) SymbolicString() string { // 0xb
	return "RRC E"
}

type RES_6_B struct {
	addr         string      // 0xb0
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_6_B) Write(w io.Writer) (int, error) { // 0xb0
	var b []byte

	b = append(b, 0xb0)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xb0,
	})
	if err != nil {
		return written, err
	}

	// 6 is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *RES_6_B) Length() uint8 { // 0xb0
	return 2
}

func (o *RES_6_B) cycles() []uint8 { // 0xb0
	return []uint8{8}
}

func (o *RES_6_B) String() string { // 0xb0
	return "RES" + " " + "6" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *RES_6_B) SymbolicString() string { // 0xb0
	return "RES 6,B"
}

type RES_6_C struct {
	addr         string      // 0xb1
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_6_C) Write(w io.Writer) (int, error) { // 0xb1
	var b []byte

	b = append(b, 0xb1)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xb1,
	})
	if err != nil {
		return written, err
	}

	// 6 is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *RES_6_C) Length() uint8 { // 0xb1
	return 2
}

func (o *RES_6_C) cycles() []uint8 { // 0xb1
	return []uint8{8}
}

func (o *RES_6_C) String() string { // 0xb1
	return "RES" + " " + "6" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *RES_6_C) SymbolicString() string { // 0xb1
	return "RES 6,C"
}

type RES_6_D struct {
	addr         string      // 0xb2
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_6_D) Write(w io.Writer) (int, error) { // 0xb2
	var b []byte

	b = append(b, 0xb2)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xb2,
	})
	if err != nil {
		return written, err
	}

	// 6 is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *RES_6_D) Length() uint8 { // 0xb2
	return 2
}

func (o *RES_6_D) cycles() []uint8 { // 0xb2
	return []uint8{8}
}

func (o *RES_6_D) String() string { // 0xb2
	return "RES" + " " + "6" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *RES_6_D) SymbolicString() string { // 0xb2
	return "RES 6,D"
}

type RES_6_E struct {
	addr         string      // 0xb3
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_6_E) Write(w io.Writer) (int, error) { // 0xb3
	var b []byte

	b = append(b, 0xb3)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xb3,
	})
	if err != nil {
		return written, err
	}

	// 6 is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *RES_6_E) Length() uint8 { // 0xb3
	return 2
}

func (o *RES_6_E) cycles() []uint8 { // 0xb3
	return []uint8{8}
}

func (o *RES_6_E) String() string { // 0xb3
	return "RES" + " " + "6" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *RES_6_E) SymbolicString() string { // 0xb3
	return "RES 6,E"
}

type RES_6_H struct {
	addr         string      // 0xb4
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_6_H) Write(w io.Writer) (int, error) { // 0xb4
	var b []byte

	b = append(b, 0xb4)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xb4,
	})
	if err != nil {
		return written, err
	}

	// 6 is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *RES_6_H) Length() uint8 { // 0xb4
	return 2
}

func (o *RES_6_H) cycles() []uint8 { // 0xb4
	return []uint8{8}
}

func (o *RES_6_H) String() string { // 0xb4
	return "RES" + " " + "6" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *RES_6_H) SymbolicString() string { // 0xb4
	return "RES 6,H"
}

type RES_6_L struct {
	addr         string      // 0xb5
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_6_L) Write(w io.Writer) (int, error) { // 0xb5
	var b []byte

	b = append(b, 0xb5)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xb5,
	})
	if err != nil {
		return written, err
	}

	// 6 is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *RES_6_L) Length() uint8 { // 0xb5
	return 2
}

func (o *RES_6_L) cycles() []uint8 { // 0xb5
	return []uint8{8}
}

func (o *RES_6_L) String() string { // 0xb5
	return "RES" + " " + "6" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *RES_6_L) SymbolicString() string { // 0xb5
	return "RES 6,L"
}

type RES_6_HLPtr struct {
	addr         string      // 0xb6
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_6_HLPtr) Write(w io.Writer) (int, error) { // 0xb6
	var b []byte

	b = append(b, 0xb6)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xb6,
	})
	if err != nil {
		return written, err
	}

	// 6 is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *RES_6_HLPtr) Length() uint8 { // 0xb6
	return 2
}

func (o *RES_6_HLPtr) cycles() []uint8 { // 0xb6
	return []uint8{16}
}

func (o *RES_6_HLPtr) String() string { // 0xb6
	return "RES" + " " + "6" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *RES_6_HLPtr) SymbolicString() string { // 0xb6
	return "RES 6,(HL)"
}

type RES_6_A struct {
	addr         string      // 0xb7
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_6_A) Write(w io.Writer) (int, error) { // 0xb7
	var b []byte

	b = append(b, 0xb7)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xb7,
	})
	if err != nil {
		return written, err
	}

	// 6 is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *RES_6_A) Length() uint8 { // 0xb7
	return 2
}

func (o *RES_6_A) cycles() []uint8 { // 0xb7
	return []uint8{8}
}

func (o *RES_6_A) String() string { // 0xb7
	return "RES" + " " + "6" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *RES_6_A) SymbolicString() string { // 0xb7
	return "RES 6,A"
}

type RES_7_B struct {
	addr         string      // 0xb8
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_7_B) Write(w io.Writer) (int, error) { // 0xb8
	var b []byte

	b = append(b, 0xb8)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xb8,
	})
	if err != nil {
		return written, err
	}

	// 7 is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *RES_7_B) Length() uint8 { // 0xb8
	return 2
}

func (o *RES_7_B) cycles() []uint8 { // 0xb8
	return []uint8{8}
}

func (o *RES_7_B) String() string { // 0xb8
	return "RES" + " " + "7" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *RES_7_B) SymbolicString() string { // 0xb8
	return "RES 7,B"
}

type RES_7_C struct {
	addr         string      // 0xb9
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_7_C) Write(w io.Writer) (int, error) { // 0xb9
	var b []byte

	b = append(b, 0xb9)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xb9,
	})
	if err != nil {
		return written, err
	}

	// 7 is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *RES_7_C) Length() uint8 { // 0xb9
	return 2
}

func (o *RES_7_C) cycles() []uint8 { // 0xb9
	return []uint8{8}
}

func (o *RES_7_C) String() string { // 0xb9
	return "RES" + " " + "7" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *RES_7_C) SymbolicString() string { // 0xb9
	return "RES 7,C"
}

type RES_7_D struct {
	addr         string      // 0xba
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_7_D) Write(w io.Writer) (int, error) { // 0xba
	var b []byte

	b = append(b, 0xba)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xba,
	})
	if err != nil {
		return written, err
	}

	// 7 is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *RES_7_D) Length() uint8 { // 0xba
	return 2
}

func (o *RES_7_D) cycles() []uint8 { // 0xba
	return []uint8{8}
}

func (o *RES_7_D) String() string { // 0xba
	return "RES" + " " + "7" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *RES_7_D) SymbolicString() string { // 0xba
	return "RES 7,D"
}

type RES_7_E struct {
	addr         string      // 0xbb
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_7_E) Write(w io.Writer) (int, error) { // 0xbb
	var b []byte

	b = append(b, 0xbb)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xbb,
	})
	if err != nil {
		return written, err
	}

	// 7 is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *RES_7_E) Length() uint8 { // 0xbb
	return 2
}

func (o *RES_7_E) cycles() []uint8 { // 0xbb
	return []uint8{8}
}

func (o *RES_7_E) String() string { // 0xbb
	return "RES" + " " + "7" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *RES_7_E) SymbolicString() string { // 0xbb
	return "RES 7,E"
}

type RES_7_H struct {
	addr         string      // 0xbc
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_7_H) Write(w io.Writer) (int, error) { // 0xbc
	var b []byte

	b = append(b, 0xbc)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xbc,
	})
	if err != nil {
		return written, err
	}

	// 7 is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *RES_7_H) Length() uint8 { // 0xbc
	return 2
}

func (o *RES_7_H) cycles() []uint8 { // 0xbc
	return []uint8{8}
}

func (o *RES_7_H) String() string { // 0xbc
	return "RES" + " " + "7" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *RES_7_H) SymbolicString() string { // 0xbc
	return "RES 7,H"
}

type RES_7_L struct {
	addr         string      // 0xbd
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_7_L) Write(w io.Writer) (int, error) { // 0xbd
	var b []byte

	b = append(b, 0xbd)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xbd,
	})
	if err != nil {
		return written, err
	}

	// 7 is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *RES_7_L) Length() uint8 { // 0xbd
	return 2
}

func (o *RES_7_L) cycles() []uint8 { // 0xbd
	return []uint8{8}
}

func (o *RES_7_L) String() string { // 0xbd
	return "RES" + " " + "7" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *RES_7_L) SymbolicString() string { // 0xbd
	return "RES 7,L"
}

type RES_7_HLPtr struct {
	addr         string      // 0xbe
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_7_HLPtr) Write(w io.Writer) (int, error) { // 0xbe
	var b []byte

	b = append(b, 0xbe)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xbe,
	})
	if err != nil {
		return written, err
	}

	// 7 is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *RES_7_HLPtr) Length() uint8 { // 0xbe
	return 2
}

func (o *RES_7_HLPtr) cycles() []uint8 { // 0xbe
	return []uint8{16}
}

func (o *RES_7_HLPtr) String() string { // 0xbe
	return "RES" + " " + "7" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *RES_7_HLPtr) SymbolicString() string { // 0xbe
	return "RES 7,(HL)"
}

type RES_7_A struct {
	addr         string      // 0xbf
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RES_7_A) Write(w io.Writer) (int, error) { // 0xbf
	var b []byte

	b = append(b, 0xbf)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xbf,
	})
	if err != nil {
		return written, err
	}

	// 7 is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *RES_7_A) Length() uint8 { // 0xbf
	return 2
}

func (o *RES_7_A) cycles() []uint8 { // 0xbf
	return []uint8{8}
}

func (o *RES_7_A) String() string { // 0xbf
	return "RES" + " " + "7" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *RES_7_A) SymbolicString() string { // 0xbf
	return "RES 7,A"
}

type RRC_H struct {
	addr         string      // 0xc
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RRC_H) Write(w io.Writer) (int, error) { // 0xc
	var b []byte

	b = append(b, 0xc)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xc,
	})
	if err != nil {
		return written, err
	}

	// H is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RRC_H) Length() uint8 { // 0xc
	return 2
}

func (o *RRC_H) cycles() []uint8 { // 0xc
	return []uint8{8}
}

func (o *RRC_H) String() string { // 0xc
	return "RRC" + " " + "H" /* operand1 */
}

func (o *RRC_H) SymbolicString() string { // 0xc
	return "RRC H"
}

type SET_0_B struct {
	addr         string      // 0xc0
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_0_B) Write(w io.Writer) (int, error) { // 0xc0
	var b []byte

	b = append(b, 0xc0)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xc0,
	})
	if err != nil {
		return written, err
	}

	// 0 is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *SET_0_B) Length() uint8 { // 0xc0
	return 2
}

func (o *SET_0_B) cycles() []uint8 { // 0xc0
	return []uint8{8}
}

func (o *SET_0_B) String() string { // 0xc0
	return "SET" + " " + "0" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *SET_0_B) SymbolicString() string { // 0xc0
	return "SET 0,B"
}

type SET_0_C struct {
	addr         string      // 0xc1
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_0_C) Write(w io.Writer) (int, error) { // 0xc1
	var b []byte

	b = append(b, 0xc1)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xc1,
	})
	if err != nil {
		return written, err
	}

	// 0 is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *SET_0_C) Length() uint8 { // 0xc1
	return 2
}

func (o *SET_0_C) cycles() []uint8 { // 0xc1
	return []uint8{8}
}

func (o *SET_0_C) String() string { // 0xc1
	return "SET" + " " + "0" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *SET_0_C) SymbolicString() string { // 0xc1
	return "SET 0,C"
}

type SET_0_D struct {
	addr         string      // 0xc2
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_0_D) Write(w io.Writer) (int, error) { // 0xc2
	var b []byte

	b = append(b, 0xc2)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xc2,
	})
	if err != nil {
		return written, err
	}

	// 0 is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *SET_0_D) Length() uint8 { // 0xc2
	return 2
}

func (o *SET_0_D) cycles() []uint8 { // 0xc2
	return []uint8{8}
}

func (o *SET_0_D) String() string { // 0xc2
	return "SET" + " " + "0" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *SET_0_D) SymbolicString() string { // 0xc2
	return "SET 0,D"
}

type SET_0_E struct {
	addr         string      // 0xc3
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_0_E) Write(w io.Writer) (int, error) { // 0xc3
	var b []byte

	b = append(b, 0xc3)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xc3,
	})
	if err != nil {
		return written, err
	}

	// 0 is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *SET_0_E) Length() uint8 { // 0xc3
	return 2
}

func (o *SET_0_E) cycles() []uint8 { // 0xc3
	return []uint8{8}
}

func (o *SET_0_E) String() string { // 0xc3
	return "SET" + " " + "0" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *SET_0_E) SymbolicString() string { // 0xc3
	return "SET 0,E"
}

type SET_0_H struct {
	addr         string      // 0xc4
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_0_H) Write(w io.Writer) (int, error) { // 0xc4
	var b []byte

	b = append(b, 0xc4)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xc4,
	})
	if err != nil {
		return written, err
	}

	// 0 is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *SET_0_H) Length() uint8 { // 0xc4
	return 2
}

func (o *SET_0_H) cycles() []uint8 { // 0xc4
	return []uint8{8}
}

func (o *SET_0_H) String() string { // 0xc4
	return "SET" + " " + "0" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *SET_0_H) SymbolicString() string { // 0xc4
	return "SET 0,H"
}

type SET_0_L struct {
	addr         string      // 0xc5
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_0_L) Write(w io.Writer) (int, error) { // 0xc5
	var b []byte

	b = append(b, 0xc5)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xc5,
	})
	if err != nil {
		return written, err
	}

	// 0 is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *SET_0_L) Length() uint8 { // 0xc5
	return 2
}

func (o *SET_0_L) cycles() []uint8 { // 0xc5
	return []uint8{8}
}

func (o *SET_0_L) String() string { // 0xc5
	return "SET" + " " + "0" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *SET_0_L) SymbolicString() string { // 0xc5
	return "SET 0,L"
}

type SET_0_HLPtr struct {
	addr         string      // 0xc6
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_0_HLPtr) Write(w io.Writer) (int, error) { // 0xc6
	var b []byte

	b = append(b, 0xc6)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xc6,
	})
	if err != nil {
		return written, err
	}

	// 0 is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *SET_0_HLPtr) Length() uint8 { // 0xc6
	return 2
}

func (o *SET_0_HLPtr) cycles() []uint8 { // 0xc6
	return []uint8{16}
}

func (o *SET_0_HLPtr) String() string { // 0xc6
	return "SET" + " " + "0" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *SET_0_HLPtr) SymbolicString() string { // 0xc6
	return "SET 0,(HL)"
}

type SET_0_A struct {
	addr         string      // 0xc7
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_0_A) Write(w io.Writer) (int, error) { // 0xc7
	var b []byte

	b = append(b, 0xc7)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xc7,
	})
	if err != nil {
		return written, err
	}

	// 0 is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *SET_0_A) Length() uint8 { // 0xc7
	return 2
}

func (o *SET_0_A) cycles() []uint8 { // 0xc7
	return []uint8{8}
}

func (o *SET_0_A) String() string { // 0xc7
	return "SET" + " " + "0" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *SET_0_A) SymbolicString() string { // 0xc7
	return "SET 0,A"
}

type SET_1_B struct {
	addr         string      // 0xc8
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_1_B) Write(w io.Writer) (int, error) { // 0xc8
	var b []byte

	b = append(b, 0xc8)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xc8,
	})
	if err != nil {
		return written, err
	}

	// 1 is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *SET_1_B) Length() uint8 { // 0xc8
	return 2
}

func (o *SET_1_B) cycles() []uint8 { // 0xc8
	return []uint8{8}
}

func (o *SET_1_B) String() string { // 0xc8
	return "SET" + " " + "1" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *SET_1_B) SymbolicString() string { // 0xc8
	return "SET 1,B"
}

type SET_1_C struct {
	addr         string      // 0xc9
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_1_C) Write(w io.Writer) (int, error) { // 0xc9
	var b []byte

	b = append(b, 0xc9)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xc9,
	})
	if err != nil {
		return written, err
	}

	// 1 is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *SET_1_C) Length() uint8 { // 0xc9
	return 2
}

func (o *SET_1_C) cycles() []uint8 { // 0xc9
	return []uint8{8}
}

func (o *SET_1_C) String() string { // 0xc9
	return "SET" + " " + "1" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *SET_1_C) SymbolicString() string { // 0xc9
	return "SET 1,C"
}

type SET_1_D struct {
	addr         string      // 0xca
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_1_D) Write(w io.Writer) (int, error) { // 0xca
	var b []byte

	b = append(b, 0xca)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xca,
	})
	if err != nil {
		return written, err
	}

	// 1 is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *SET_1_D) Length() uint8 { // 0xca
	return 2
}

func (o *SET_1_D) cycles() []uint8 { // 0xca
	return []uint8{8}
}

func (o *SET_1_D) String() string { // 0xca
	return "SET" + " " + "1" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *SET_1_D) SymbolicString() string { // 0xca
	return "SET 1,D"
}

type SET_1_E struct {
	addr         string      // 0xcb
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_1_E) Write(w io.Writer) (int, error) { // 0xcb
	var b []byte

	b = append(b, 0xcb)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xcb,
	})
	if err != nil {
		return written, err
	}

	// 1 is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *SET_1_E) Length() uint8 { // 0xcb
	return 2
}

func (o *SET_1_E) cycles() []uint8 { // 0xcb
	return []uint8{8}
}

func (o *SET_1_E) String() string { // 0xcb
	return "SET" + " " + "1" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *SET_1_E) SymbolicString() string { // 0xcb
	return "SET 1,E"
}

type SET_1_H struct {
	addr         string      // 0xcc
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_1_H) Write(w io.Writer) (int, error) { // 0xcc
	var b []byte

	b = append(b, 0xcc)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xcc,
	})
	if err != nil {
		return written, err
	}

	// 1 is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *SET_1_H) Length() uint8 { // 0xcc
	return 2
}

func (o *SET_1_H) cycles() []uint8 { // 0xcc
	return []uint8{8}
}

func (o *SET_1_H) String() string { // 0xcc
	return "SET" + " " + "1" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *SET_1_H) SymbolicString() string { // 0xcc
	return "SET 1,H"
}

type SET_1_L struct {
	addr         string      // 0xcd
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_1_L) Write(w io.Writer) (int, error) { // 0xcd
	var b []byte

	b = append(b, 0xcd)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xcd,
	})
	if err != nil {
		return written, err
	}

	// 1 is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *SET_1_L) Length() uint8 { // 0xcd
	return 2
}

func (o *SET_1_L) cycles() []uint8 { // 0xcd
	return []uint8{8}
}

func (o *SET_1_L) String() string { // 0xcd
	return "SET" + " " + "1" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *SET_1_L) SymbolicString() string { // 0xcd
	return "SET 1,L"
}

type SET_1_HLPtr struct {
	addr         string      // 0xce
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_1_HLPtr) Write(w io.Writer) (int, error) { // 0xce
	var b []byte

	b = append(b, 0xce)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xce,
	})
	if err != nil {
		return written, err
	}

	// 1 is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *SET_1_HLPtr) Length() uint8 { // 0xce
	return 2
}

func (o *SET_1_HLPtr) cycles() []uint8 { // 0xce
	return []uint8{16}
}

func (o *SET_1_HLPtr) String() string { // 0xce
	return "SET" + " " + "1" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *SET_1_HLPtr) SymbolicString() string { // 0xce
	return "SET 1,(HL)"
}

type SET_1_A struct {
	addr         string      // 0xcf
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_1_A) Write(w io.Writer) (int, error) { // 0xcf
	var b []byte

	b = append(b, 0xcf)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xcf,
	})
	if err != nil {
		return written, err
	}

	// 1 is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *SET_1_A) Length() uint8 { // 0xcf
	return 2
}

func (o *SET_1_A) cycles() []uint8 { // 0xcf
	return []uint8{8}
}

func (o *SET_1_A) String() string { // 0xcf
	return "SET" + " " + "1" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *SET_1_A) SymbolicString() string { // 0xcf
	return "SET 1,A"
}

type RRC_L struct {
	addr         string      // 0xd
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RRC_L) Write(w io.Writer) (int, error) { // 0xd
	var b []byte

	b = append(b, 0xd)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xd,
	})
	if err != nil {
		return written, err
	}

	// L is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RRC_L) Length() uint8 { // 0xd
	return 2
}

func (o *RRC_L) cycles() []uint8 { // 0xd
	return []uint8{8}
}

func (o *RRC_L) String() string { // 0xd
	return "RRC" + " " + "L" /* operand1 */
}

func (o *RRC_L) SymbolicString() string { // 0xd
	return "RRC L"
}

type SET_2_B struct {
	addr         string      // 0xd0
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_2_B) Write(w io.Writer) (int, error) { // 0xd0
	var b []byte

	b = append(b, 0xd0)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xd0,
	})
	if err != nil {
		return written, err
	}

	// 2 is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *SET_2_B) Length() uint8 { // 0xd0
	return 2
}

func (o *SET_2_B) cycles() []uint8 { // 0xd0
	return []uint8{8}
}

func (o *SET_2_B) String() string { // 0xd0
	return "SET" + " " + "2" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *SET_2_B) SymbolicString() string { // 0xd0
	return "SET 2,B"
}

type SET_2_C struct {
	addr         string      // 0xd1
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_2_C) Write(w io.Writer) (int, error) { // 0xd1
	var b []byte

	b = append(b, 0xd1)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xd1,
	})
	if err != nil {
		return written, err
	}

	// 2 is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *SET_2_C) Length() uint8 { // 0xd1
	return 2
}

func (o *SET_2_C) cycles() []uint8 { // 0xd1
	return []uint8{8}
}

func (o *SET_2_C) String() string { // 0xd1
	return "SET" + " " + "2" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *SET_2_C) SymbolicString() string { // 0xd1
	return "SET 2,C"
}

type SET_2_D struct {
	addr         string      // 0xd2
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_2_D) Write(w io.Writer) (int, error) { // 0xd2
	var b []byte

	b = append(b, 0xd2)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xd2,
	})
	if err != nil {
		return written, err
	}

	// 2 is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *SET_2_D) Length() uint8 { // 0xd2
	return 2
}

func (o *SET_2_D) cycles() []uint8 { // 0xd2
	return []uint8{8}
}

func (o *SET_2_D) String() string { // 0xd2
	return "SET" + " " + "2" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *SET_2_D) SymbolicString() string { // 0xd2
	return "SET 2,D"
}

type SET_2_E struct {
	addr         string      // 0xd3
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_2_E) Write(w io.Writer) (int, error) { // 0xd3
	var b []byte

	b = append(b, 0xd3)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xd3,
	})
	if err != nil {
		return written, err
	}

	// 2 is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *SET_2_E) Length() uint8 { // 0xd3
	return 2
}

func (o *SET_2_E) cycles() []uint8 { // 0xd3
	return []uint8{8}
}

func (o *SET_2_E) String() string { // 0xd3
	return "SET" + " " + "2" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *SET_2_E) SymbolicString() string { // 0xd3
	return "SET 2,E"
}

type SET_2_H struct {
	addr         string      // 0xd4
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_2_H) Write(w io.Writer) (int, error) { // 0xd4
	var b []byte

	b = append(b, 0xd4)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xd4,
	})
	if err != nil {
		return written, err
	}

	// 2 is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *SET_2_H) Length() uint8 { // 0xd4
	return 2
}

func (o *SET_2_H) cycles() []uint8 { // 0xd4
	return []uint8{8}
}

func (o *SET_2_H) String() string { // 0xd4
	return "SET" + " " + "2" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *SET_2_H) SymbolicString() string { // 0xd4
	return "SET 2,H"
}

type SET_2_L struct {
	addr         string      // 0xd5
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_2_L) Write(w io.Writer) (int, error) { // 0xd5
	var b []byte

	b = append(b, 0xd5)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xd5,
	})
	if err != nil {
		return written, err
	}

	// 2 is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *SET_2_L) Length() uint8 { // 0xd5
	return 2
}

func (o *SET_2_L) cycles() []uint8 { // 0xd5
	return []uint8{8}
}

func (o *SET_2_L) String() string { // 0xd5
	return "SET" + " " + "2" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *SET_2_L) SymbolicString() string { // 0xd5
	return "SET 2,L"
}

type SET_2_HLPtr struct {
	addr         string      // 0xd6
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_2_HLPtr) Write(w io.Writer) (int, error) { // 0xd6
	var b []byte

	b = append(b, 0xd6)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xd6,
	})
	if err != nil {
		return written, err
	}

	// 2 is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *SET_2_HLPtr) Length() uint8 { // 0xd6
	return 2
}

func (o *SET_2_HLPtr) cycles() []uint8 { // 0xd6
	return []uint8{16}
}

func (o *SET_2_HLPtr) String() string { // 0xd6
	return "SET" + " " + "2" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *SET_2_HLPtr) SymbolicString() string { // 0xd6
	return "SET 2,(HL)"
}

type SET_2_A struct {
	addr         string      // 0xd7
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_2_A) Write(w io.Writer) (int, error) { // 0xd7
	var b []byte

	b = append(b, 0xd7)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xd7,
	})
	if err != nil {
		return written, err
	}

	// 2 is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *SET_2_A) Length() uint8 { // 0xd7
	return 2
}

func (o *SET_2_A) cycles() []uint8 { // 0xd7
	return []uint8{8}
}

func (o *SET_2_A) String() string { // 0xd7
	return "SET" + " " + "2" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *SET_2_A) SymbolicString() string { // 0xd7
	return "SET 2,A"
}

type SET_3_B struct {
	addr         string      // 0xd8
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_3_B) Write(w io.Writer) (int, error) { // 0xd8
	var b []byte

	b = append(b, 0xd8)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xd8,
	})
	if err != nil {
		return written, err
	}

	// 3 is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *SET_3_B) Length() uint8 { // 0xd8
	return 2
}

func (o *SET_3_B) cycles() []uint8 { // 0xd8
	return []uint8{8}
}

func (o *SET_3_B) String() string { // 0xd8
	return "SET" + " " + "3" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *SET_3_B) SymbolicString() string { // 0xd8
	return "SET 3,B"
}

type SET_3_C struct {
	addr         string      // 0xd9
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_3_C) Write(w io.Writer) (int, error) { // 0xd9
	var b []byte

	b = append(b, 0xd9)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xd9,
	})
	if err != nil {
		return written, err
	}

	// 3 is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *SET_3_C) Length() uint8 { // 0xd9
	return 2
}

func (o *SET_3_C) cycles() []uint8 { // 0xd9
	return []uint8{8}
}

func (o *SET_3_C) String() string { // 0xd9
	return "SET" + " " + "3" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *SET_3_C) SymbolicString() string { // 0xd9
	return "SET 3,C"
}

type SET_3_D struct {
	addr         string      // 0xda
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_3_D) Write(w io.Writer) (int, error) { // 0xda
	var b []byte

	b = append(b, 0xda)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xda,
	})
	if err != nil {
		return written, err
	}

	// 3 is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *SET_3_D) Length() uint8 { // 0xda
	return 2
}

func (o *SET_3_D) cycles() []uint8 { // 0xda
	return []uint8{8}
}

func (o *SET_3_D) String() string { // 0xda
	return "SET" + " " + "3" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *SET_3_D) SymbolicString() string { // 0xda
	return "SET 3,D"
}

type SET_3_E struct {
	addr         string      // 0xdb
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_3_E) Write(w io.Writer) (int, error) { // 0xdb
	var b []byte

	b = append(b, 0xdb)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xdb,
	})
	if err != nil {
		return written, err
	}

	// 3 is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *SET_3_E) Length() uint8 { // 0xdb
	return 2
}

func (o *SET_3_E) cycles() []uint8 { // 0xdb
	return []uint8{8}
}

func (o *SET_3_E) String() string { // 0xdb
	return "SET" + " " + "3" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *SET_3_E) SymbolicString() string { // 0xdb
	return "SET 3,E"
}

type SET_3_H struct {
	addr         string      // 0xdc
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_3_H) Write(w io.Writer) (int, error) { // 0xdc
	var b []byte

	b = append(b, 0xdc)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xdc,
	})
	if err != nil {
		return written, err
	}

	// 3 is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *SET_3_H) Length() uint8 { // 0xdc
	return 2
}

func (o *SET_3_H) cycles() []uint8 { // 0xdc
	return []uint8{8}
}

func (o *SET_3_H) String() string { // 0xdc
	return "SET" + " " + "3" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *SET_3_H) SymbolicString() string { // 0xdc
	return "SET 3,H"
}

type SET_3_L struct {
	addr         string      // 0xdd
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_3_L) Write(w io.Writer) (int, error) { // 0xdd
	var b []byte

	b = append(b, 0xdd)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xdd,
	})
	if err != nil {
		return written, err
	}

	// 3 is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *SET_3_L) Length() uint8 { // 0xdd
	return 2
}

func (o *SET_3_L) cycles() []uint8 { // 0xdd
	return []uint8{8}
}

func (o *SET_3_L) String() string { // 0xdd
	return "SET" + " " + "3" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *SET_3_L) SymbolicString() string { // 0xdd
	return "SET 3,L"
}

type SET_3_HLPtr struct {
	addr         string      // 0xde
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_3_HLPtr) Write(w io.Writer) (int, error) { // 0xde
	var b []byte

	b = append(b, 0xde)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xde,
	})
	if err != nil {
		return written, err
	}

	// 3 is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *SET_3_HLPtr) Length() uint8 { // 0xde
	return 2
}

func (o *SET_3_HLPtr) cycles() []uint8 { // 0xde
	return []uint8{16}
}

func (o *SET_3_HLPtr) String() string { // 0xde
	return "SET" + " " + "3" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *SET_3_HLPtr) SymbolicString() string { // 0xde
	return "SET 3,(HL)"
}

type SET_3_A struct {
	addr         string      // 0xdf
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_3_A) Write(w io.Writer) (int, error) { // 0xdf
	var b []byte

	b = append(b, 0xdf)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xdf,
	})
	if err != nil {
		return written, err
	}

	// 3 is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *SET_3_A) Length() uint8 { // 0xdf
	return 2
}

func (o *SET_3_A) cycles() []uint8 { // 0xdf
	return []uint8{8}
}

func (o *SET_3_A) String() string { // 0xdf
	return "SET" + " " + "3" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *SET_3_A) SymbolicString() string { // 0xdf
	return "SET 3,A"
}

type RRC_HLPtr struct {
	addr         string      // 0xe
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RRC_HLPtr) Write(w io.Writer) (int, error) { // 0xe
	var b []byte

	b = append(b, 0xe)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xe,
	})
	if err != nil {
		return written, err
	}

	// (HL) is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RRC_HLPtr) Length() uint8 { // 0xe
	return 2
}

func (o *RRC_HLPtr) cycles() []uint8 { // 0xe
	return []uint8{16}
}

func (o *RRC_HLPtr) String() string { // 0xe
	return "RRC" + " " + "(HL)" /* operand1 */
}

func (o *RRC_HLPtr) SymbolicString() string { // 0xe
	return "RRC (HL)"
}

type SET_4_B struct {
	addr         string      // 0xe0
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_4_B) Write(w io.Writer) (int, error) { // 0xe0
	var b []byte

	b = append(b, 0xe0)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xe0,
	})
	if err != nil {
		return written, err
	}

	// 4 is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *SET_4_B) Length() uint8 { // 0xe0
	return 2
}

func (o *SET_4_B) cycles() []uint8 { // 0xe0
	return []uint8{8}
}

func (o *SET_4_B) String() string { // 0xe0
	return "SET" + " " + "4" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *SET_4_B) SymbolicString() string { // 0xe0
	return "SET 4,B"
}

type SET_4_C struct {
	addr         string      // 0xe1
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_4_C) Write(w io.Writer) (int, error) { // 0xe1
	var b []byte

	b = append(b, 0xe1)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xe1,
	})
	if err != nil {
		return written, err
	}

	// 4 is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *SET_4_C) Length() uint8 { // 0xe1
	return 2
}

func (o *SET_4_C) cycles() []uint8 { // 0xe1
	return []uint8{8}
}

func (o *SET_4_C) String() string { // 0xe1
	return "SET" + " " + "4" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *SET_4_C) SymbolicString() string { // 0xe1
	return "SET 4,C"
}

type SET_4_D struct {
	addr         string      // 0xe2
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_4_D) Write(w io.Writer) (int, error) { // 0xe2
	var b []byte

	b = append(b, 0xe2)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xe2,
	})
	if err != nil {
		return written, err
	}

	// 4 is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *SET_4_D) Length() uint8 { // 0xe2
	return 2
}

func (o *SET_4_D) cycles() []uint8 { // 0xe2
	return []uint8{8}
}

func (o *SET_4_D) String() string { // 0xe2
	return "SET" + " " + "4" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *SET_4_D) SymbolicString() string { // 0xe2
	return "SET 4,D"
}

type SET_4_E struct {
	addr         string      // 0xe3
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_4_E) Write(w io.Writer) (int, error) { // 0xe3
	var b []byte

	b = append(b, 0xe3)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xe3,
	})
	if err != nil {
		return written, err
	}

	// 4 is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *SET_4_E) Length() uint8 { // 0xe3
	return 2
}

func (o *SET_4_E) cycles() []uint8 { // 0xe3
	return []uint8{8}
}

func (o *SET_4_E) String() string { // 0xe3
	return "SET" + " " + "4" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *SET_4_E) SymbolicString() string { // 0xe3
	return "SET 4,E"
}

type SET_4_H struct {
	addr         string      // 0xe4
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_4_H) Write(w io.Writer) (int, error) { // 0xe4
	var b []byte

	b = append(b, 0xe4)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xe4,
	})
	if err != nil {
		return written, err
	}

	// 4 is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *SET_4_H) Length() uint8 { // 0xe4
	return 2
}

func (o *SET_4_H) cycles() []uint8 { // 0xe4
	return []uint8{8}
}

func (o *SET_4_H) String() string { // 0xe4
	return "SET" + " " + "4" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *SET_4_H) SymbolicString() string { // 0xe4
	return "SET 4,H"
}

type SET_4_L struct {
	addr         string      // 0xe5
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_4_L) Write(w io.Writer) (int, error) { // 0xe5
	var b []byte

	b = append(b, 0xe5)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xe5,
	})
	if err != nil {
		return written, err
	}

	// 4 is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *SET_4_L) Length() uint8 { // 0xe5
	return 2
}

func (o *SET_4_L) cycles() []uint8 { // 0xe5
	return []uint8{8}
}

func (o *SET_4_L) String() string { // 0xe5
	return "SET" + " " + "4" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *SET_4_L) SymbolicString() string { // 0xe5
	return "SET 4,L"
}

type SET_4_HLPtr struct {
	addr         string      // 0xe6
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_4_HLPtr) Write(w io.Writer) (int, error) { // 0xe6
	var b []byte

	b = append(b, 0xe6)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xe6,
	})
	if err != nil {
		return written, err
	}

	// 4 is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *SET_4_HLPtr) Length() uint8 { // 0xe6
	return 2
}

func (o *SET_4_HLPtr) cycles() []uint8 { // 0xe6
	return []uint8{16}
}

func (o *SET_4_HLPtr) String() string { // 0xe6
	return "SET" + " " + "4" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *SET_4_HLPtr) SymbolicString() string { // 0xe6
	return "SET 4,(HL)"
}

type SET_4_A struct {
	addr         string      // 0xe7
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_4_A) Write(w io.Writer) (int, error) { // 0xe7
	var b []byte

	b = append(b, 0xe7)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xe7,
	})
	if err != nil {
		return written, err
	}

	// 4 is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *SET_4_A) Length() uint8 { // 0xe7
	return 2
}

func (o *SET_4_A) cycles() []uint8 { // 0xe7
	return []uint8{8}
}

func (o *SET_4_A) String() string { // 0xe7
	return "SET" + " " + "4" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *SET_4_A) SymbolicString() string { // 0xe7
	return "SET 4,A"
}

type SET_5_B struct {
	addr         string      // 0xe8
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_5_B) Write(w io.Writer) (int, error) { // 0xe8
	var b []byte

	b = append(b, 0xe8)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xe8,
	})
	if err != nil {
		return written, err
	}

	// 5 is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *SET_5_B) Length() uint8 { // 0xe8
	return 2
}

func (o *SET_5_B) cycles() []uint8 { // 0xe8
	return []uint8{8}
}

func (o *SET_5_B) String() string { // 0xe8
	return "SET" + " " + "5" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *SET_5_B) SymbolicString() string { // 0xe8
	return "SET 5,B"
}

type SET_5_C struct {
	addr         string      // 0xe9
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_5_C) Write(w io.Writer) (int, error) { // 0xe9
	var b []byte

	b = append(b, 0xe9)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xe9,
	})
	if err != nil {
		return written, err
	}

	// 5 is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *SET_5_C) Length() uint8 { // 0xe9
	return 2
}

func (o *SET_5_C) cycles() []uint8 { // 0xe9
	return []uint8{8}
}

func (o *SET_5_C) String() string { // 0xe9
	return "SET" + " " + "5" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *SET_5_C) SymbolicString() string { // 0xe9
	return "SET 5,C"
}

type SET_5_D struct {
	addr         string      // 0xea
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_5_D) Write(w io.Writer) (int, error) { // 0xea
	var b []byte

	b = append(b, 0xea)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xea,
	})
	if err != nil {
		return written, err
	}

	// 5 is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *SET_5_D) Length() uint8 { // 0xea
	return 2
}

func (o *SET_5_D) cycles() []uint8 { // 0xea
	return []uint8{8}
}

func (o *SET_5_D) String() string { // 0xea
	return "SET" + " " + "5" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *SET_5_D) SymbolicString() string { // 0xea
	return "SET 5,D"
}

type SET_5_E struct {
	addr         string      // 0xeb
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_5_E) Write(w io.Writer) (int, error) { // 0xeb
	var b []byte

	b = append(b, 0xeb)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xeb,
	})
	if err != nil {
		return written, err
	}

	// 5 is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *SET_5_E) Length() uint8 { // 0xeb
	return 2
}

func (o *SET_5_E) cycles() []uint8 { // 0xeb
	return []uint8{8}
}

func (o *SET_5_E) String() string { // 0xeb
	return "SET" + " " + "5" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *SET_5_E) SymbolicString() string { // 0xeb
	return "SET 5,E"
}

type SET_5_H struct {
	addr         string      // 0xec
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_5_H) Write(w io.Writer) (int, error) { // 0xec
	var b []byte

	b = append(b, 0xec)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xec,
	})
	if err != nil {
		return written, err
	}

	// 5 is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *SET_5_H) Length() uint8 { // 0xec
	return 2
}

func (o *SET_5_H) cycles() []uint8 { // 0xec
	return []uint8{8}
}

func (o *SET_5_H) String() string { // 0xec
	return "SET" + " " + "5" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *SET_5_H) SymbolicString() string { // 0xec
	return "SET 5,H"
}

type SET_5_L struct {
	addr         string      // 0xed
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_5_L) Write(w io.Writer) (int, error) { // 0xed
	var b []byte

	b = append(b, 0xed)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xed,
	})
	if err != nil {
		return written, err
	}

	// 5 is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *SET_5_L) Length() uint8 { // 0xed
	return 2
}

func (o *SET_5_L) cycles() []uint8 { // 0xed
	return []uint8{8}
}

func (o *SET_5_L) String() string { // 0xed
	return "SET" + " " + "5" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *SET_5_L) SymbolicString() string { // 0xed
	return "SET 5,L"
}

type SET_5_HLPtr struct {
	addr         string      // 0xee
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_5_HLPtr) Write(w io.Writer) (int, error) { // 0xee
	var b []byte

	b = append(b, 0xee)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xee,
	})
	if err != nil {
		return written, err
	}

	// 5 is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *SET_5_HLPtr) Length() uint8 { // 0xee
	return 2
}

func (o *SET_5_HLPtr) cycles() []uint8 { // 0xee
	return []uint8{16}
}

func (o *SET_5_HLPtr) String() string { // 0xee
	return "SET" + " " + "5" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *SET_5_HLPtr) SymbolicString() string { // 0xee
	return "SET 5,(HL)"
}

type SET_5_A struct {
	addr         string      // 0xef
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_5_A) Write(w io.Writer) (int, error) { // 0xef
	var b []byte

	b = append(b, 0xef)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xef,
	})
	if err != nil {
		return written, err
	}

	// 5 is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *SET_5_A) Length() uint8 { // 0xef
	return 2
}

func (o *SET_5_A) cycles() []uint8 { // 0xef
	return []uint8{8}
}

func (o *SET_5_A) String() string { // 0xef
	return "SET" + " " + "5" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *SET_5_A) SymbolicString() string { // 0xef
	return "SET 5,A"
}

type RRC_A struct {
	addr         string      // 0xf
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *RRC_A) Write(w io.Writer) (int, error) { // 0xf
	var b []byte

	b = append(b, 0xf)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xf,
	})
	if err != nil {
		return written, err
	}

	// A is implicit in this instruction.
	//  is implicit in this instruction.

	return written, err
}

func (o *RRC_A) Length() uint8 { // 0xf
	return 2
}

func (o *RRC_A) cycles() []uint8 { // 0xf
	return []uint8{8}
}

func (o *RRC_A) String() string { // 0xf
	return "RRC" + " " + "A" /* operand1 */
}

func (o *RRC_A) SymbolicString() string { // 0xf
	return "RRC A"
}

type SET_6_B struct {
	addr         string      // 0xf0
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_6_B) Write(w io.Writer) (int, error) { // 0xf0
	var b []byte

	b = append(b, 0xf0)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xf0,
	})
	if err != nil {
		return written, err
	}

	// 6 is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *SET_6_B) Length() uint8 { // 0xf0
	return 2
}

func (o *SET_6_B) cycles() []uint8 { // 0xf0
	return []uint8{8}
}

func (o *SET_6_B) String() string { // 0xf0
	return "SET" + " " + "6" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *SET_6_B) SymbolicString() string { // 0xf0
	return "SET 6,B"
}

type SET_6_C struct {
	addr         string      // 0xf1
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_6_C) Write(w io.Writer) (int, error) { // 0xf1
	var b []byte

	b = append(b, 0xf1)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xf1,
	})
	if err != nil {
		return written, err
	}

	// 6 is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *SET_6_C) Length() uint8 { // 0xf1
	return 2
}

func (o *SET_6_C) cycles() []uint8 { // 0xf1
	return []uint8{8}
}

func (o *SET_6_C) String() string { // 0xf1
	return "SET" + " " + "6" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *SET_6_C) SymbolicString() string { // 0xf1
	return "SET 6,C"
}

type SET_6_D struct {
	addr         string      // 0xf2
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_6_D) Write(w io.Writer) (int, error) { // 0xf2
	var b []byte

	b = append(b, 0xf2)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xf2,
	})
	if err != nil {
		return written, err
	}

	// 6 is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *SET_6_D) Length() uint8 { // 0xf2
	return 2
}

func (o *SET_6_D) cycles() []uint8 { // 0xf2
	return []uint8{8}
}

func (o *SET_6_D) String() string { // 0xf2
	return "SET" + " " + "6" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *SET_6_D) SymbolicString() string { // 0xf2
	return "SET 6,D"
}

type SET_6_E struct {
	addr         string      // 0xf3
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_6_E) Write(w io.Writer) (int, error) { // 0xf3
	var b []byte

	b = append(b, 0xf3)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xf3,
	})
	if err != nil {
		return written, err
	}

	// 6 is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *SET_6_E) Length() uint8 { // 0xf3
	return 2
}

func (o *SET_6_E) cycles() []uint8 { // 0xf3
	return []uint8{8}
}

func (o *SET_6_E) String() string { // 0xf3
	return "SET" + " " + "6" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *SET_6_E) SymbolicString() string { // 0xf3
	return "SET 6,E"
}

type SET_6_H struct {
	addr         string      // 0xf4
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_6_H) Write(w io.Writer) (int, error) { // 0xf4
	var b []byte

	b = append(b, 0xf4)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xf4,
	})
	if err != nil {
		return written, err
	}

	// 6 is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *SET_6_H) Length() uint8 { // 0xf4
	return 2
}

func (o *SET_6_H) cycles() []uint8 { // 0xf4
	return []uint8{8}
}

func (o *SET_6_H) String() string { // 0xf4
	return "SET" + " " + "6" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *SET_6_H) SymbolicString() string { // 0xf4
	return "SET 6,H"
}

type SET_6_L struct {
	addr         string      // 0xf5
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_6_L) Write(w io.Writer) (int, error) { // 0xf5
	var b []byte

	b = append(b, 0xf5)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xf5,
	})
	if err != nil {
		return written, err
	}

	// 6 is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *SET_6_L) Length() uint8 { // 0xf5
	return 2
}

func (o *SET_6_L) cycles() []uint8 { // 0xf5
	return []uint8{8}
}

func (o *SET_6_L) String() string { // 0xf5
	return "SET" + " " + "6" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *SET_6_L) SymbolicString() string { // 0xf5
	return "SET 6,L"
}

type SET_6_HLPtr struct {
	addr         string      // 0xf6
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_6_HLPtr) Write(w io.Writer) (int, error) { // 0xf6
	var b []byte

	b = append(b, 0xf6)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xf6,
	})
	if err != nil {
		return written, err
	}

	// 6 is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *SET_6_HLPtr) Length() uint8 { // 0xf6
	return 2
}

func (o *SET_6_HLPtr) cycles() []uint8 { // 0xf6
	return []uint8{16}
}

func (o *SET_6_HLPtr) String() string { // 0xf6
	return "SET" + " " + "6" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *SET_6_HLPtr) SymbolicString() string { // 0xf6
	return "SET 6,(HL)"
}

type SET_6_A struct {
	addr         string      // 0xf7
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_6_A) Write(w io.Writer) (int, error) { // 0xf7
	var b []byte

	b = append(b, 0xf7)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xf7,
	})
	if err != nil {
		return written, err
	}

	// 6 is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *SET_6_A) Length() uint8 { // 0xf7
	return 2
}

func (o *SET_6_A) cycles() []uint8 { // 0xf7
	return []uint8{8}
}

func (o *SET_6_A) String() string { // 0xf7
	return "SET" + " " + "6" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *SET_6_A) SymbolicString() string { // 0xf7
	return "SET 6,A"
}

type SET_7_B struct {
	addr         string      // 0xf8
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_7_B) Write(w io.Writer) (int, error) { // 0xf8
	var b []byte

	b = append(b, 0xf8)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xf8,
	})
	if err != nil {
		return written, err
	}

	// 7 is implicit in this instruction.
	// B is implicit in this instruction.

	return written, err
}

func (o *SET_7_B) Length() uint8 { // 0xf8
	return 2
}

func (o *SET_7_B) cycles() []uint8 { // 0xf8
	return []uint8{8}
}

func (o *SET_7_B) String() string { // 0xf8
	return "SET" + " " + "7" /* operand1 */ + " " + "B" /* operand2 */
}

func (o *SET_7_B) SymbolicString() string { // 0xf8
	return "SET 7,B"
}

type SET_7_C struct {
	addr         string      // 0xf9
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_7_C) Write(w io.Writer) (int, error) { // 0xf9
	var b []byte

	b = append(b, 0xf9)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xf9,
	})
	if err != nil {
		return written, err
	}

	// 7 is implicit in this instruction.
	// C is implicit in this instruction.

	return written, err
}

func (o *SET_7_C) Length() uint8 { // 0xf9
	return 2
}

func (o *SET_7_C) cycles() []uint8 { // 0xf9
	return []uint8{8}
}

func (o *SET_7_C) String() string { // 0xf9
	return "SET" + " " + "7" /* operand1 */ + " " + "C" /* operand2 */
}

func (o *SET_7_C) SymbolicString() string { // 0xf9
	return "SET 7,C"
}

type SET_7_D struct {
	addr         string      // 0xfa
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_7_D) Write(w io.Writer) (int, error) { // 0xfa
	var b []byte

	b = append(b, 0xfa)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xfa,
	})
	if err != nil {
		return written, err
	}

	// 7 is implicit in this instruction.
	// D is implicit in this instruction.

	return written, err
}

func (o *SET_7_D) Length() uint8 { // 0xfa
	return 2
}

func (o *SET_7_D) cycles() []uint8 { // 0xfa
	return []uint8{8}
}

func (o *SET_7_D) String() string { // 0xfa
	return "SET" + " " + "7" /* operand1 */ + " " + "D" /* operand2 */
}

func (o *SET_7_D) SymbolicString() string { // 0xfa
	return "SET 7,D"
}

type SET_7_E struct {
	addr         string      // 0xfb
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_7_E) Write(w io.Writer) (int, error) { // 0xfb
	var b []byte

	b = append(b, 0xfb)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xfb,
	})
	if err != nil {
		return written, err
	}

	// 7 is implicit in this instruction.
	// E is implicit in this instruction.

	return written, err
}

func (o *SET_7_E) Length() uint8 { // 0xfb
	return 2
}

func (o *SET_7_E) cycles() []uint8 { // 0xfb
	return []uint8{8}
}

func (o *SET_7_E) String() string { // 0xfb
	return "SET" + " " + "7" /* operand1 */ + " " + "E" /* operand2 */
}

func (o *SET_7_E) SymbolicString() string { // 0xfb
	return "SET 7,E"
}

type SET_7_H struct {
	addr         string      // 0xfc
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_7_H) Write(w io.Writer) (int, error) { // 0xfc
	var b []byte

	b = append(b, 0xfc)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xfc,
	})
	if err != nil {
		return written, err
	}

	// 7 is implicit in this instruction.
	// H is implicit in this instruction.

	return written, err
}

func (o *SET_7_H) Length() uint8 { // 0xfc
	return 2
}

func (o *SET_7_H) cycles() []uint8 { // 0xfc
	return []uint8{8}
}

func (o *SET_7_H) String() string { // 0xfc
	return "SET" + " " + "7" /* operand1 */ + " " + "H" /* operand2 */
}

func (o *SET_7_H) SymbolicString() string { // 0xfc
	return "SET 7,H"
}

type SET_7_L struct {
	addr         string      // 0xfd
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_7_L) Write(w io.Writer) (int, error) { // 0xfd
	var b []byte

	b = append(b, 0xfd)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xfd,
	})
	if err != nil {
		return written, err
	}

	// 7 is implicit in this instruction.
	// L is implicit in this instruction.

	return written, err
}

func (o *SET_7_L) Length() uint8 { // 0xfd
	return 2
}

func (o *SET_7_L) cycles() []uint8 { // 0xfd
	return []uint8{8}
}

func (o *SET_7_L) String() string { // 0xfd
	return "SET" + " " + "7" /* operand1 */ + " " + "L" /* operand2 */
}

func (o *SET_7_L) SymbolicString() string { // 0xfd
	return "SET 7,L"
}

type SET_7_HLPtr struct {
	addr         string      // 0xfe
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_7_HLPtr) Write(w io.Writer) (int, error) { // 0xfe
	var b []byte

	b = append(b, 0xfe)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xfe,
	})
	if err != nil {
		return written, err
	}

	// 7 is implicit in this instruction.
	// (HL) is implicit in this instruction.

	return written, err
}

func (o *SET_7_HLPtr) Length() uint8 { // 0xfe
	return 2
}

func (o *SET_7_HLPtr) cycles() []uint8 { // 0xfe
	return []uint8{16}
}

func (o *SET_7_HLPtr) String() string { // 0xfe
	return "SET" + " " + "7" /* operand1 */ + " " + "(HL)" /* operand2 */
}

func (o *SET_7_HLPtr) SymbolicString() string { // 0xfe
	return "SET 7,(HL)"
}

type SET_7_A struct {
	addr         string      // 0xff
	operand1     interface{} // literal, or reg name, or (HL)...
	operand2     interface{} // same
	isCbPrefixed bool
}

func (o *SET_7_A) Write(w io.Writer) (int, error) { // 0xff
	var b []byte

	b = append(b, 0xff)

	var v int64
	_ = v // suppress unused variable warning in the case we have no operands

	written, err := w.Write([]byte{
		0xCB,
		0xff,
	})
	if err != nil {
		return written, err
	}

	// 7 is implicit in this instruction.
	// A is implicit in this instruction.

	return written, err
}

func (o *SET_7_A) Length() uint8 { // 0xff
	return 2
}

func (o *SET_7_A) cycles() []uint8 { // 0xff
	return []uint8{8}
}

func (o *SET_7_A) String() string { // 0xff
	return "SET" + " " + "7" /* operand1 */ + " " + "A" /* operand2 */
}

func (o *SET_7_A) SymbolicString() string { // 0xff
	return "SET 7,A"
}

// ReadOpCode returns an executable opcode by taking an io.Reader
// and reading a single instruction off it. If there is no more data
// returns undelying io.Reader's EOF error type.
func ReadInstruction(data io.Reader) (Instruction, error) {

	d := make([]byte, 1)
	_, err := data.Read(d)
	if err != nil {
		return nil, err
	}
	switch d[0] {

	case 0x0: // 0x0 - NOP
		o := &NOP{}

		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x1: // 0x1 - LD
		o := &LD_BC_d16{}

		{
			s := "BC" // BC is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s, err := readImmediate16BitData(data)
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x10: // 0x10 - STOP
		o := &STOP_0{}

		{
			s := "0" // 0 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x11: // 0x11 - LD
		o := &LD_DE_d16{}

		{
			s := "DE" // DE is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s, err := readImmediate16BitData(data)
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x12: // 0x12 - LD
		o := &LD_DEDeref_A{}

		{
			s := "(DE)" // (DE) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x13: // 0x13 - INC
		o := &INC_DE{}

		{
			s := "DE" // DE is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x14: // 0x14 - INC
		o := &INC_D{}

		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x15: // 0x15 - DEC
		o := &DEC_D{}

		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x16: // 0x16 - LD
		o := &LD_D_d8{}

		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s, err := readImmediate8BitData(data)
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x17: // 0x17 - RLA
		o := &RLA{}

		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x18: // 0x18 - JR
		o := &JR_r8{}

		{
			s, err := readImmediateSigned8BitData(data)
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x19: // 0x19 - ADD
		o := &ADD_HL_DE{}

		{
			s := "HL" // HL is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "DE" // DE is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x1a: // 0x1a - LD
		o := &LD_A_DEDeref{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(DE)" // (DE) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x1b: // 0x1b - DEC
		o := &DEC_DE{}

		{
			s := "DE" // DE is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x1c: // 0x1c - INC
		o := &INC_E{}

		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x1d: // 0x1d - DEC
		o := &DEC_E{}

		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x1e: // 0x1e - LD
		o := &LD_E_d8{}

		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s, err := readImmediate8BitData(data)
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x1f: // 0x1f - RRA
		o := &RRA{}

		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x2: // 0x2 - LD
		o := &LD_BCDeref_A{}

		{
			s := "(BC)" // (BC) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x20: // 0x20 - JR
		o := &JR_NZ_r8{}

		{
			s := "NZ" // NZ is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s, err := readImmediateSigned8BitData(data)
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x21: // 0x21 - LD
		o := &LD_HL_d16{}

		{
			s := "HL" // HL is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s, err := readImmediate16BitData(data)
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x22: // 0x22 - LD
		o := &LD_HLPtrInc_A{}

		{
			s := "(HL+)" // (HL+) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x23: // 0x23 - INC
		o := &INC_HL{}

		{
			s := "HL" // HL is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x24: // 0x24 - INC
		o := &INC_H{}

		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x25: // 0x25 - DEC
		o := &DEC_H{}

		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x26: // 0x26 - LD
		o := &LD_H_d8{}

		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s, err := readImmediate8BitData(data)
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x27: // 0x27 - DAA
		o := &DAA{}

		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x28: // 0x28 - JR
		o := &JR_Z_r8{}

		{
			s := "Z" // Z is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s, err := readImmediateSigned8BitData(data)
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x29: // 0x29 - ADD
		o := &ADD_HL_HL{}

		{
			s := "HL" // HL is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "HL" // HL is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x2a: // 0x2a - LD
		o := &LD_A_HLPtrInc{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL+)" // (HL+) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x2b: // 0x2b - DEC
		o := &DEC_HL{}

		{
			s := "HL" // HL is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x2c: // 0x2c - INC
		o := &INC_L{}

		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x2d: // 0x2d - DEC
		o := &DEC_L{}

		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x2e: // 0x2e - LD
		o := &LD_L_d8{}

		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s, err := readImmediate8BitData(data)
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x2f: // 0x2f - CPL
		o := &CPL{}

		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x3: // 0x3 - INC
		o := &INC_BC{}

		{
			s := "BC" // BC is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x30: // 0x30 - JR
		o := &JR_NC_r8{}

		{
			s := "NC" // NC is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s, err := readImmediateSigned8BitData(data)
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x31: // 0x31 - LD
		o := &LD_SP_d16{}

		{
			s := "SP" // SP is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s, err := readImmediate16BitData(data)
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x32: // 0x32 - LD
		o := &LD_HLPtrDec_A{}

		{
			s := "(HL-)" // (HL-) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x33: // 0x33 - INC
		o := &INC_SP{}

		{
			s := "SP" // SP is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x34: // 0x34 - INC
		o := &INC_HLPtr{}

		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x35: // 0x35 - DEC
		o := &DEC_HLPtr{}

		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x36: // 0x36 - LD
		o := &LD_HLPtr_d8{}

		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s, err := readImmediate8BitData(data)
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x37: // 0x37 - SCF
		o := &SCF{}

		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x38: // 0x38 - JR
		o := &JR_C_r8{}

		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s, err := readImmediateSigned8BitData(data)
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x39: // 0x39 - ADD
		o := &ADD_HL_SP{}

		{
			s := "HL" // HL is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "SP" // SP is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x3a: // 0x3a - LD
		o := &LD_A_HLPtrDec{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL-)" // (HL-) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x3b: // 0x3b - DEC
		o := &DEC_SP{}

		{
			s := "SP" // SP is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x3c: // 0x3c - INC
		o := &INC_A{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x3d: // 0x3d - DEC
		o := &DEC_A{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x3e: // 0x3e - LD
		o := &LD_A_d8{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s, err := readImmediate8BitData(data)
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x3f: // 0x3f - CCF
		o := &CCF{}

		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x4: // 0x4 - INC
		o := &INC_B{}

		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x40: // 0x40 - LD
		o := &LD_B_B{}

		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x41: // 0x41 - LD
		o := &LD_B_C{}

		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x42: // 0x42 - LD
		o := &LD_B_D{}

		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x43: // 0x43 - LD
		o := &LD_B_E{}

		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x44: // 0x44 - LD
		o := &LD_B_H{}

		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x45: // 0x45 - LD
		o := &LD_B_L{}

		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x46: // 0x46 - LD
		o := &LD_B_HLPtr{}

		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x47: // 0x47 - LD
		o := &LD_B_A{}

		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x48: // 0x48 - LD
		o := &LD_C_B{}

		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x49: // 0x49 - LD
		o := &LD_C_C{}

		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x4a: // 0x4a - LD
		o := &LD_C_D{}

		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x4b: // 0x4b - LD
		o := &LD_C_E{}

		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x4c: // 0x4c - LD
		o := &LD_C_H{}

		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x4d: // 0x4d - LD
		o := &LD_C_L{}

		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x4e: // 0x4e - LD
		o := &LD_C_HLPtr{}

		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x4f: // 0x4f - LD
		o := &LD_C_A{}

		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x5: // 0x5 - DEC
		o := &DEC_B{}

		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x50: // 0x50 - LD
		o := &LD_D_B{}

		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x51: // 0x51 - LD
		o := &LD_D_C{}

		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x52: // 0x52 - LD
		o := &LD_D_D{}

		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x53: // 0x53 - LD
		o := &LD_D_E{}

		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x54: // 0x54 - LD
		o := &LD_D_H{}

		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x55: // 0x55 - LD
		o := &LD_D_L{}

		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x56: // 0x56 - LD
		o := &LD_D_HLPtr{}

		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x57: // 0x57 - LD
		o := &LD_D_A{}

		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x58: // 0x58 - LD
		o := &LD_E_B{}

		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x59: // 0x59 - LD
		o := &LD_E_C{}

		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x5a: // 0x5a - LD
		o := &LD_E_D{}

		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x5b: // 0x5b - LD
		o := &LD_E_E{}

		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x5c: // 0x5c - LD
		o := &LD_E_H{}

		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x5d: // 0x5d - LD
		o := &LD_E_L{}

		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x5e: // 0x5e - LD
		o := &LD_E_HLPtr{}

		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x5f: // 0x5f - LD
		o := &LD_E_A{}

		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x6: // 0x6 - LD
		o := &LD_B_d8{}

		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s, err := readImmediate8BitData(data)
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x60: // 0x60 - LD
		o := &LD_H_B{}

		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x61: // 0x61 - LD
		o := &LD_H_C{}

		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x62: // 0x62 - LD
		o := &LD_H_D{}

		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x63: // 0x63 - LD
		o := &LD_H_E{}

		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x64: // 0x64 - LD
		o := &LD_H_H{}

		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x65: // 0x65 - LD
		o := &LD_H_L{}

		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x66: // 0x66 - LD
		o := &LD_H_HLPtr{}

		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x67: // 0x67 - LD
		o := &LD_H_A{}

		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x68: // 0x68 - LD
		o := &LD_L_B{}

		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x69: // 0x69 - LD
		o := &LD_L_C{}

		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x6a: // 0x6a - LD
		o := &LD_L_D{}

		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x6b: // 0x6b - LD
		o := &LD_L_E{}

		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x6c: // 0x6c - LD
		o := &LD_L_H{}

		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x6d: // 0x6d - LD
		o := &LD_L_L{}

		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x6e: // 0x6e - LD
		o := &LD_L_HLPtr{}

		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x6f: // 0x6f - LD
		o := &LD_L_A{}

		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x7: // 0x7 - RLCA
		o := &RLCA{}

		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x70: // 0x70 - LD
		o := &LD_HLPtr_B{}

		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x71: // 0x71 - LD
		o := &LD_HLPtr_C{}

		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x72: // 0x72 - LD
		o := &LD_HLPtr_D{}

		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x73: // 0x73 - LD
		o := &LD_HLPtr_E{}

		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x74: // 0x74 - LD
		o := &LD_HLPtr_H{}

		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x75: // 0x75 - LD
		o := &LD_HLPtr_L{}

		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x76: // 0x76 - HALT
		o := &HALT{}

		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x77: // 0x77 - LD
		o := &LD_HLPtr_A{}

		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x78: // 0x78 - LD
		o := &LD_A_B{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x79: // 0x79 - LD
		o := &LD_A_C{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x7a: // 0x7a - LD
		o := &LD_A_D{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x7b: // 0x7b - LD
		o := &LD_A_E{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x7c: // 0x7c - LD
		o := &LD_A_H{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x7d: // 0x7d - LD
		o := &LD_A_L{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x7e: // 0x7e - LD
		o := &LD_A_HLPtr{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x7f: // 0x7f - LD
		o := &LD_A_A{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x8: // 0x8 - LD
		o := &LD_a16Deref_SP{}

		{
			s, err := readImmediate16BitAddress(data)
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "SP" // SP is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x80: // 0x80 - ADD
		o := &ADD_A_B{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x81: // 0x81 - ADD
		o := &ADD_A_C{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x82: // 0x82 - ADD
		o := &ADD_A_D{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x83: // 0x83 - ADD
		o := &ADD_A_E{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x84: // 0x84 - ADD
		o := &ADD_A_H{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x85: // 0x85 - ADD
		o := &ADD_A_L{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x86: // 0x86 - ADD
		o := &ADD_A_HLPtr{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x87: // 0x87 - ADD
		o := &ADD_A_A{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x88: // 0x88 - ADC
		o := &ADC_A_B{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x89: // 0x89 - ADC
		o := &ADC_A_C{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x8a: // 0x8a - ADC
		o := &ADC_A_D{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x8b: // 0x8b - ADC
		o := &ADC_A_E{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x8c: // 0x8c - ADC
		o := &ADC_A_H{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x8d: // 0x8d - ADC
		o := &ADC_A_L{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x8e: // 0x8e - ADC
		o := &ADC_A_HLPtr{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x8f: // 0x8f - ADC
		o := &ADC_A_A{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x9: // 0x9 - ADD
		o := &ADD_HL_BC{}

		{
			s := "HL" // HL is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "BC" // BC is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x90: // 0x90 - SUB
		o := &SUB_B{}

		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x91: // 0x91 - SUB
		o := &SUB_C{}

		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x92: // 0x92 - SUB
		o := &SUB_D{}

		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x93: // 0x93 - SUB
		o := &SUB_E{}

		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x94: // 0x94 - SUB
		o := &SUB_H{}

		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x95: // 0x95 - SUB
		o := &SUB_L{}

		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x96: // 0x96 - SUB
		o := &SUB_HLPtr{}

		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x97: // 0x97 - SUB
		o := &SUB_A{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x98: // 0x98 - SBC
		o := &SBC_A_B{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x99: // 0x99 - SBC
		o := &SBC_A_C{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x9a: // 0x9a - SBC
		o := &SBC_A_D{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x9b: // 0x9b - SBC
		o := &SBC_A_E{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x9c: // 0x9c - SBC
		o := &SBC_A_H{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x9d: // 0x9d - SBC
		o := &SBC_A_L{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x9e: // 0x9e - SBC
		o := &SBC_A_HLPtr{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x9f: // 0x9f - SBC
		o := &SBC_A_A{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xa: // 0xa - LD
		o := &LD_A_BCDeref{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(BC)" // (BC) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xa0: // 0xa0 - AND
		o := &AND_B{}

		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xa1: // 0xa1 - AND
		o := &AND_C{}

		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xa2: // 0xa2 - AND
		o := &AND_D{}

		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xa3: // 0xa3 - AND
		o := &AND_E{}

		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xa4: // 0xa4 - AND
		o := &AND_H{}

		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xa5: // 0xa5 - AND
		o := &AND_L{}

		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xa6: // 0xa6 - AND
		o := &AND_HLPtr{}

		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xa7: // 0xa7 - AND
		o := &AND_A{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xa8: // 0xa8 - XOR
		o := &XOR_B{}

		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xa9: // 0xa9 - XOR
		o := &XOR_C{}

		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xaa: // 0xaa - XOR
		o := &XOR_D{}

		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xab: // 0xab - XOR
		o := &XOR_E{}

		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xac: // 0xac - XOR
		o := &XOR_H{}

		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xad: // 0xad - XOR
		o := &XOR_L{}

		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xae: // 0xae - XOR
		o := &XOR_HLPtr{}

		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xaf: // 0xaf - XOR
		o := &XOR_A{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xb: // 0xb - DEC
		o := &DEC_BC{}

		{
			s := "BC" // BC is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xb0: // 0xb0 - OR
		o := &OR_B{}

		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xb1: // 0xb1 - OR
		o := &OR_C{}

		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xb2: // 0xb2 - OR
		o := &OR_D{}

		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xb3: // 0xb3 - OR
		o := &OR_E{}

		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xb4: // 0xb4 - OR
		o := &OR_H{}

		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xb5: // 0xb5 - OR
		o := &OR_L{}

		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xb6: // 0xb6 - OR
		o := &OR_HLPtr{}

		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xb7: // 0xb7 - OR
		o := &OR_A{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xb8: // 0xb8 - CP
		o := &CP_B{}

		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xb9: // 0xb9 - CP
		o := &CP_C{}

		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xba: // 0xba - CP
		o := &CP_D{}

		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xbb: // 0xbb - CP
		o := &CP_E{}

		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xbc: // 0xbc - CP
		o := &CP_H{}

		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xbd: // 0xbd - CP
		o := &CP_L{}

		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xbe: // 0xbe - CP
		o := &CP_HLPtr{}

		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xbf: // 0xbf - CP
		o := &CP_A{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xc: // 0xc - INC
		o := &INC_C{}

		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xc0: // 0xc0 - RET
		o := &RET_NZ{}

		{
			s := "NZ" // NZ is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xc1: // 0xc1 - POP
		o := &POP_BC{}

		{
			s := "BC" // BC is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xc2: // 0xc2 - JP
		o := &JP_NZ_a16{}

		{
			s := "NZ" // NZ is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s, err := readImmediate16BitAddress(data)
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xc3: // 0xc3 - JP
		o := &JP_a16{}

		{
			s, err := readImmediate16BitAddress(data)
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xc4: // 0xc4 - CALL
		o := &CALL_NZ_a16{}

		{
			s := "NZ" // NZ is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s, err := readImmediate16BitAddress(data)
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xc5: // 0xc5 - PUSH
		o := &PUSH_BC{}

		{
			s := "BC" // BC is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xc6: // 0xc6 - ADD
		o := &ADD_A_d8{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s, err := readImmediate8BitData(data)
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xc7: // 0xc7 - RST
		o := &RST_00H{}

		{
			s := "00H" // 00H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xc8: // 0xc8 - RET
		o := &RET_Z{}

		{
			s := "Z" // Z is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xc9: // 0xc9 - RET
		o := &RET{}

		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xca: // 0xca - JP
		o := &JP_Z_a16{}

		{
			s := "Z" // Z is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s, err := readImmediate16BitAddress(data)
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xcb: // 0xcb - PREFIX
		return readCBPrefixedInstruction(data)

	case 0xcc: // 0xcc - CALL
		o := &CALL_Z_a16{}

		{
			s := "Z" // Z is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s, err := readImmediate16BitAddress(data)
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xcd: // 0xcd - CALL
		o := &CALL_a16{}

		{
			s, err := readImmediate16BitAddress(data)
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xce: // 0xce - ADC
		o := &ADC_A_d8{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s, err := readImmediate8BitData(data)
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xcf: // 0xcf - RST
		o := &RST_08H{}

		{
			s := "08H" // 08H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xd: // 0xd - DEC
		o := &DEC_C{}

		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xd0: // 0xd0 - RET
		o := &RET_NC{}

		{
			s := "NC" // NC is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xd1: // 0xd1 - POP
		o := &POP_DE{}

		{
			s := "DE" // DE is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xd2: // 0xd2 - JP
		o := &JP_NC_a16{}

		{
			s := "NC" // NC is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s, err := readImmediate16BitAddress(data)
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xd4: // 0xd4 - CALL
		o := &CALL_NC_a16{}

		{
			s := "NC" // NC is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s, err := readImmediate16BitAddress(data)
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xd5: // 0xd5 - PUSH
		o := &PUSH_DE{}

		{
			s := "DE" // DE is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xd6: // 0xd6 - SUB
		o := &SUB_d8{}

		{
			s, err := readImmediate8BitData(data)
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xd7: // 0xd7 - RST
		o := &RST_10H{}

		{
			s := "10H" // 10H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xd8: // 0xd8 - RET
		o := &RET_C{}

		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xd9: // 0xd9 - RETI
		o := &RETI{}

		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xda: // 0xda - JP
		o := &JP_C_a16{}

		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s, err := readImmediate16BitAddress(data)
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xdc: // 0xdc - CALL
		o := &CALL_C_a16{}

		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s, err := readImmediate16BitAddress(data)
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xde: // 0xde - SBC
		o := &SBC_A_d8{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s, err := readImmediate8BitData(data)
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xdf: // 0xdf - RST
		o := &RST_18H{}

		{
			s := "18H" // 18H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xe: // 0xe - LD
		o := &LD_C_d8{}

		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s, err := readImmediate8BitData(data)
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xe0: // 0xe0 - LDH
		o := &LDH_a8Deref_A{}

		{
			s, err := readImmediate8BitAddress(data)
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xe1: // 0xe1 - POP
		o := &POP_HL{}

		{
			s := "HL" // HL is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xe2: // 0xe2 - LD
		o := &LD_CDeref_A{}

		{
			s := "(C)" // (C) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xe5: // 0xe5 - PUSH
		o := &PUSH_HL{}

		{
			s := "HL" // HL is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xe6: // 0xe6 - AND
		o := &AND_d8{}

		{
			s, err := readImmediate8BitData(data)
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xe7: // 0xe7 - RST
		o := &RST_20H{}

		{
			s := "20H" // 20H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xe8: // 0xe8 - ADD
		o := &ADD_SP_r8{}

		{
			s := "SP" // SP is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s, err := readImmediateSigned8BitData(data)
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xe9: // 0xe9 - JP
		o := &JP_HLPtr{}

		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xea: // 0xea - LD
		o := &LD_a16Deref_A{}

		{
			s, err := readImmediate16BitAddress(data)
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xee: // 0xee - XOR
		o := &XOR_d8{}

		{
			s, err := readImmediate8BitData(data)
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xef: // 0xef - RST
		o := &RST_28H{}

		{
			s := "28H" // 28H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xf: // 0xf - RRCA
		o := &RRCA{}

		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xf0: // 0xf0 - LDH
		o := &LDH_A_a8Deref{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s, err := readImmediate8BitAddress(data)
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xf1: // 0xf1 - POP
		o := &POP_AF{}

		{
			s := "AF" // AF is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xf2: // 0xf2 - LD
		o := &LD_A_CDeref{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(C)" // (C) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xf3: // 0xf3 - DI
		o := &DI{}

		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xf5: // 0xf5 - PUSH
		o := &PUSH_AF{}

		{
			s := "AF" // AF is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xf6: // 0xf6 - OR
		o := &OR_d8{}

		{
			s, err := readImmediate8BitData(data)
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xf7: // 0xf7 - RST
		o := &RST_30H{}

		{
			s := "30H" // 30H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xf8: // 0xf8 - LD
		o := &LD_HL_SP_plus_r8{}

		{
			s := "HL" // HL is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s, err := readImmediateSigned8BitData(data)
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xf9: // 0xf9 - LD
		o := &LD_SP_HL{}

		{
			s := "SP" // SP is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "HL" // HL is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xfa: // 0xfa - LD
		o := &LD_A_a16Deref{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s, err := readImmediate16BitAddress(data)
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xfb: // 0xfb - EI
		o := &EI{}

		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xfe: // 0xfe - CP
		o := &CP_d8{}

		{
			s, err := readImmediate8BitData(data)
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xff: // 0xff - RST
		o := &RST_38H{}

		{
			s := "38H" // 38H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	default:
		return nil, fmt.Errorf("the proposed opcode (dec %d, hex %x) doesn't exist: %v", d[0], d[0], ErrNoOpCode)
	}

}

func readCBPrefixedInstruction(data io.Reader) (Instruction, error) {

	d := make([]byte, 1)
	_, err := data.Read(d)
	if err != nil {
		return nil, err
	}
	switch d[0] {

	case 0x0: // 0x0 - RLC
		o := &RLC_B{}

		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x1: // 0x1 - RLC
		o := &RLC_C{}

		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x10: // 0x10 - RL
		o := &RL_B{}

		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x11: // 0x11 - RL
		o := &RL_C{}

		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x12: // 0x12 - RL
		o := &RL_D{}

		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x13: // 0x13 - RL
		o := &RL_E{}

		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x14: // 0x14 - RL
		o := &RL_H{}

		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x15: // 0x15 - RL
		o := &RL_L{}

		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x16: // 0x16 - RL
		o := &RL_HLPtr{}

		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x17: // 0x17 - RL
		o := &RL_A{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x18: // 0x18 - RR
		o := &RR_B{}

		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x19: // 0x19 - RR
		o := &RR_C{}

		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x1a: // 0x1a - RR
		o := &RR_D{}

		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x1b: // 0x1b - RR
		o := &RR_E{}

		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x1c: // 0x1c - RR
		o := &RR_H{}

		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x1d: // 0x1d - RR
		o := &RR_L{}

		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x1e: // 0x1e - RR
		o := &RR_HLPtr{}

		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x1f: // 0x1f - RR
		o := &RR_A{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x2: // 0x2 - RLC
		o := &RLC_D{}

		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x20: // 0x20 - SLA
		o := &SLA_B{}

		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x21: // 0x21 - SLA
		o := &SLA_C{}

		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x22: // 0x22 - SLA
		o := &SLA_D{}

		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x23: // 0x23 - SLA
		o := &SLA_E{}

		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x24: // 0x24 - SLA
		o := &SLA_H{}

		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x25: // 0x25 - SLA
		o := &SLA_L{}

		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x26: // 0x26 - SLA
		o := &SLA_HLPtr{}

		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x27: // 0x27 - SLA
		o := &SLA_A{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x28: // 0x28 - SRA
		o := &SRA_B{}

		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x29: // 0x29 - SRA
		o := &SRA_C{}

		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x2a: // 0x2a - SRA
		o := &SRA_D{}

		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x2b: // 0x2b - SRA
		o := &SRA_E{}

		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x2c: // 0x2c - SRA
		o := &SRA_H{}

		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x2d: // 0x2d - SRA
		o := &SRA_L{}

		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x2e: // 0x2e - SRA
		o := &SRA_HLPtr{}

		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x2f: // 0x2f - SRA
		o := &SRA_A{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x3: // 0x3 - RLC
		o := &RLC_E{}

		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x30: // 0x30 - SWAP
		o := &SWAP_B{}

		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x31: // 0x31 - SWAP
		o := &SWAP_C{}

		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x32: // 0x32 - SWAP
		o := &SWAP_D{}

		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x33: // 0x33 - SWAP
		o := &SWAP_E{}

		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x34: // 0x34 - SWAP
		o := &SWAP_H{}

		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x35: // 0x35 - SWAP
		o := &SWAP_L{}

		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x36: // 0x36 - SWAP
		o := &SWAP_HLPtr{}

		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x37: // 0x37 - SWAP
		o := &SWAP_A{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x38: // 0x38 - SRL
		o := &SRL_B{}

		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x39: // 0x39 - SRL
		o := &SRL_C{}

		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x3a: // 0x3a - SRL
		o := &SRL_D{}

		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x3b: // 0x3b - SRL
		o := &SRL_E{}

		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x3c: // 0x3c - SRL
		o := &SRL_H{}

		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x3d: // 0x3d - SRL
		o := &SRL_L{}

		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x3e: // 0x3e - SRL
		o := &SRL_HLPtr{}

		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x3f: // 0x3f - SRL
		o := &SRL_A{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x4: // 0x4 - RLC
		o := &RLC_H{}

		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x40: // 0x40 - BIT
		o := &BIT_0_B{}

		{
			s := "0" // 0 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x41: // 0x41 - BIT
		o := &BIT_0_C{}

		{
			s := "0" // 0 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x42: // 0x42 - BIT
		o := &BIT_0_D{}

		{
			s := "0" // 0 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x43: // 0x43 - BIT
		o := &BIT_0_E{}

		{
			s := "0" // 0 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x44: // 0x44 - BIT
		o := &BIT_0_H{}

		{
			s := "0" // 0 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x45: // 0x45 - BIT
		o := &BIT_0_L{}

		{
			s := "0" // 0 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x46: // 0x46 - BIT
		o := &BIT_0_HLPtr{}

		{
			s := "0" // 0 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x47: // 0x47 - BIT
		o := &BIT_0_A{}

		{
			s := "0" // 0 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x48: // 0x48 - BIT
		o := &BIT_1_B{}

		{
			s := "1" // 1 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x49: // 0x49 - BIT
		o := &BIT_1_C{}

		{
			s := "1" // 1 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x4a: // 0x4a - BIT
		o := &BIT_1_D{}

		{
			s := "1" // 1 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x4b: // 0x4b - BIT
		o := &BIT_1_E{}

		{
			s := "1" // 1 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x4c: // 0x4c - BIT
		o := &BIT_1_H{}

		{
			s := "1" // 1 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x4d: // 0x4d - BIT
		o := &BIT_1_L{}

		{
			s := "1" // 1 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x4e: // 0x4e - BIT
		o := &BIT_1_HLPtr{}

		{
			s := "1" // 1 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x4f: // 0x4f - BIT
		o := &BIT_1_A{}

		{
			s := "1" // 1 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x5: // 0x5 - RLC
		o := &RLC_L{}

		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x50: // 0x50 - BIT
		o := &BIT_2_B{}

		{
			s := "2" // 2 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x51: // 0x51 - BIT
		o := &BIT_2_C{}

		{
			s := "2" // 2 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x52: // 0x52 - BIT
		o := &BIT_2_D{}

		{
			s := "2" // 2 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x53: // 0x53 - BIT
		o := &BIT_2_E{}

		{
			s := "2" // 2 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x54: // 0x54 - BIT
		o := &BIT_2_H{}

		{
			s := "2" // 2 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x55: // 0x55 - BIT
		o := &BIT_2_L{}

		{
			s := "2" // 2 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x56: // 0x56 - BIT
		o := &BIT_2_HLPtr{}

		{
			s := "2" // 2 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x57: // 0x57 - BIT
		o := &BIT_2_A{}

		{
			s := "2" // 2 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x58: // 0x58 - BIT
		o := &BIT_3_B{}

		{
			s := "3" // 3 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x59: // 0x59 - BIT
		o := &BIT_3_C{}

		{
			s := "3" // 3 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x5a: // 0x5a - BIT
		o := &BIT_3_D{}

		{
			s := "3" // 3 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x5b: // 0x5b - BIT
		o := &BIT_3_E{}

		{
			s := "3" // 3 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x5c: // 0x5c - BIT
		o := &BIT_3_H{}

		{
			s := "3" // 3 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x5d: // 0x5d - BIT
		o := &BIT_3_L{}

		{
			s := "3" // 3 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x5e: // 0x5e - BIT
		o := &BIT_3_HLPtr{}

		{
			s := "3" // 3 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x5f: // 0x5f - BIT
		o := &BIT_3_A{}

		{
			s := "3" // 3 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x6: // 0x6 - RLC
		o := &RLC_HLPtr{}

		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x60: // 0x60 - BIT
		o := &BIT_4_B{}

		{
			s := "4" // 4 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x61: // 0x61 - BIT
		o := &BIT_4_C{}

		{
			s := "4" // 4 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x62: // 0x62 - BIT
		o := &BIT_4_D{}

		{
			s := "4" // 4 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x63: // 0x63 - BIT
		o := &BIT_4_E{}

		{
			s := "4" // 4 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x64: // 0x64 - BIT
		o := &BIT_4_H{}

		{
			s := "4" // 4 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x65: // 0x65 - BIT
		o := &BIT_4_L{}

		{
			s := "4" // 4 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x66: // 0x66 - BIT
		o := &BIT_4_HLPtr{}

		{
			s := "4" // 4 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x67: // 0x67 - BIT
		o := &BIT_4_A{}

		{
			s := "4" // 4 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x68: // 0x68 - BIT
		o := &BIT_5_B{}

		{
			s := "5" // 5 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x69: // 0x69 - BIT
		o := &BIT_5_C{}

		{
			s := "5" // 5 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x6a: // 0x6a - BIT
		o := &BIT_5_D{}

		{
			s := "5" // 5 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x6b: // 0x6b - BIT
		o := &BIT_5_E{}

		{
			s := "5" // 5 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x6c: // 0x6c - BIT
		o := &BIT_5_H{}

		{
			s := "5" // 5 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x6d: // 0x6d - BIT
		o := &BIT_5_L{}

		{
			s := "5" // 5 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x6e: // 0x6e - BIT
		o := &BIT_5_HLPtr{}

		{
			s := "5" // 5 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x6f: // 0x6f - BIT
		o := &BIT_5_A{}

		{
			s := "5" // 5 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x7: // 0x7 - RLC
		o := &RLC_A{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x70: // 0x70 - BIT
		o := &BIT_6_B{}

		{
			s := "6" // 6 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x71: // 0x71 - BIT
		o := &BIT_6_C{}

		{
			s := "6" // 6 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x72: // 0x72 - BIT
		o := &BIT_6_D{}

		{
			s := "6" // 6 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x73: // 0x73 - BIT
		o := &BIT_6_E{}

		{
			s := "6" // 6 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x74: // 0x74 - BIT
		o := &BIT_6_H{}

		{
			s := "6" // 6 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x75: // 0x75 - BIT
		o := &BIT_6_L{}

		{
			s := "6" // 6 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x76: // 0x76 - BIT
		o := &BIT_6_HLPtr{}

		{
			s := "6" // 6 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x77: // 0x77 - BIT
		o := &BIT_6_A{}

		{
			s := "6" // 6 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x78: // 0x78 - BIT
		o := &BIT_7_B{}

		{
			s := "7" // 7 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x79: // 0x79 - BIT
		o := &BIT_7_C{}

		{
			s := "7" // 7 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x7a: // 0x7a - BIT
		o := &BIT_7_D{}

		{
			s := "7" // 7 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x7b: // 0x7b - BIT
		o := &BIT_7_E{}

		{
			s := "7" // 7 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x7c: // 0x7c - BIT
		o := &BIT_7_H{}

		{
			s := "7" // 7 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x7d: // 0x7d - BIT
		o := &BIT_7_L{}

		{
			s := "7" // 7 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x7e: // 0x7e - BIT
		o := &BIT_7_HLPtr{}

		{
			s := "7" // 7 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x7f: // 0x7f - BIT
		o := &BIT_7_A{}

		{
			s := "7" // 7 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x8: // 0x8 - RRC
		o := &RRC_B{}

		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x80: // 0x80 - RES
		o := &RES_0_B{}

		{
			s := "0" // 0 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x81: // 0x81 - RES
		o := &RES_0_C{}

		{
			s := "0" // 0 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x82: // 0x82 - RES
		o := &RES_0_D{}

		{
			s := "0" // 0 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x83: // 0x83 - RES
		o := &RES_0_E{}

		{
			s := "0" // 0 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x84: // 0x84 - RES
		o := &RES_0_H{}

		{
			s := "0" // 0 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x85: // 0x85 - RES
		o := &RES_0_L{}

		{
			s := "0" // 0 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x86: // 0x86 - RES
		o := &RES_0_HLPtr{}

		{
			s := "0" // 0 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x87: // 0x87 - RES
		o := &RES_0_A{}

		{
			s := "0" // 0 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x88: // 0x88 - RES
		o := &RES_1_B{}

		{
			s := "1" // 1 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x89: // 0x89 - RES
		o := &RES_1_C{}

		{
			s := "1" // 1 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x8a: // 0x8a - RES
		o := &RES_1_D{}

		{
			s := "1" // 1 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x8b: // 0x8b - RES
		o := &RES_1_E{}

		{
			s := "1" // 1 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x8c: // 0x8c - RES
		o := &RES_1_H{}

		{
			s := "1" // 1 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x8d: // 0x8d - RES
		o := &RES_1_L{}

		{
			s := "1" // 1 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x8e: // 0x8e - RES
		o := &RES_1_HLPtr{}

		{
			s := "1" // 1 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x8f: // 0x8f - RES
		o := &RES_1_A{}

		{
			s := "1" // 1 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x9: // 0x9 - RRC
		o := &RRC_C{}

		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x90: // 0x90 - RES
		o := &RES_2_B{}

		{
			s := "2" // 2 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x91: // 0x91 - RES
		o := &RES_2_C{}

		{
			s := "2" // 2 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x92: // 0x92 - RES
		o := &RES_2_D{}

		{
			s := "2" // 2 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x93: // 0x93 - RES
		o := &RES_2_E{}

		{
			s := "2" // 2 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x94: // 0x94 - RES
		o := &RES_2_H{}

		{
			s := "2" // 2 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x95: // 0x95 - RES
		o := &RES_2_L{}

		{
			s := "2" // 2 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x96: // 0x96 - RES
		o := &RES_2_HLPtr{}

		{
			s := "2" // 2 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x97: // 0x97 - RES
		o := &RES_2_A{}

		{
			s := "2" // 2 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x98: // 0x98 - RES
		o := &RES_3_B{}

		{
			s := "3" // 3 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x99: // 0x99 - RES
		o := &RES_3_C{}

		{
			s := "3" // 3 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x9a: // 0x9a - RES
		o := &RES_3_D{}

		{
			s := "3" // 3 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x9b: // 0x9b - RES
		o := &RES_3_E{}

		{
			s := "3" // 3 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x9c: // 0x9c - RES
		o := &RES_3_H{}

		{
			s := "3" // 3 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x9d: // 0x9d - RES
		o := &RES_3_L{}

		{
			s := "3" // 3 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x9e: // 0x9e - RES
		o := &RES_3_HLPtr{}

		{
			s := "3" // 3 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0x9f: // 0x9f - RES
		o := &RES_3_A{}

		{
			s := "3" // 3 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xa: // 0xa - RRC
		o := &RRC_D{}

		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xa0: // 0xa0 - RES
		o := &RES_4_B{}

		{
			s := "4" // 4 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xa1: // 0xa1 - RES
		o := &RES_4_C{}

		{
			s := "4" // 4 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xa2: // 0xa2 - RES
		o := &RES_4_D{}

		{
			s := "4" // 4 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xa3: // 0xa3 - RES
		o := &RES_4_E{}

		{
			s := "4" // 4 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xa4: // 0xa4 - RES
		o := &RES_4_H{}

		{
			s := "4" // 4 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xa5: // 0xa5 - RES
		o := &RES_4_L{}

		{
			s := "4" // 4 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xa6: // 0xa6 - RES
		o := &RES_4_HLPtr{}

		{
			s := "4" // 4 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xa7: // 0xa7 - RES
		o := &RES_4_A{}

		{
			s := "4" // 4 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xa8: // 0xa8 - RES
		o := &RES_5_B{}

		{
			s := "5" // 5 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xa9: // 0xa9 - RES
		o := &RES_5_C{}

		{
			s := "5" // 5 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xaa: // 0xaa - RES
		o := &RES_5_D{}

		{
			s := "5" // 5 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xab: // 0xab - RES
		o := &RES_5_E{}

		{
			s := "5" // 5 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xac: // 0xac - RES
		o := &RES_5_H{}

		{
			s := "5" // 5 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xad: // 0xad - RES
		o := &RES_5_L{}

		{
			s := "5" // 5 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xae: // 0xae - RES
		o := &RES_5_HLPtr{}

		{
			s := "5" // 5 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xaf: // 0xaf - RES
		o := &RES_5_A{}

		{
			s := "5" // 5 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xb: // 0xb - RRC
		o := &RRC_E{}

		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xb0: // 0xb0 - RES
		o := &RES_6_B{}

		{
			s := "6" // 6 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xb1: // 0xb1 - RES
		o := &RES_6_C{}

		{
			s := "6" // 6 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xb2: // 0xb2 - RES
		o := &RES_6_D{}

		{
			s := "6" // 6 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xb3: // 0xb3 - RES
		o := &RES_6_E{}

		{
			s := "6" // 6 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xb4: // 0xb4 - RES
		o := &RES_6_H{}

		{
			s := "6" // 6 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xb5: // 0xb5 - RES
		o := &RES_6_L{}

		{
			s := "6" // 6 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xb6: // 0xb6 - RES
		o := &RES_6_HLPtr{}

		{
			s := "6" // 6 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xb7: // 0xb7 - RES
		o := &RES_6_A{}

		{
			s := "6" // 6 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xb8: // 0xb8 - RES
		o := &RES_7_B{}

		{
			s := "7" // 7 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xb9: // 0xb9 - RES
		o := &RES_7_C{}

		{
			s := "7" // 7 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xba: // 0xba - RES
		o := &RES_7_D{}

		{
			s := "7" // 7 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xbb: // 0xbb - RES
		o := &RES_7_E{}

		{
			s := "7" // 7 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xbc: // 0xbc - RES
		o := &RES_7_H{}

		{
			s := "7" // 7 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xbd: // 0xbd - RES
		o := &RES_7_L{}

		{
			s := "7" // 7 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xbe: // 0xbe - RES
		o := &RES_7_HLPtr{}

		{
			s := "7" // 7 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xbf: // 0xbf - RES
		o := &RES_7_A{}

		{
			s := "7" // 7 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xc: // 0xc - RRC
		o := &RRC_H{}

		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xc0: // 0xc0 - SET
		o := &SET_0_B{}

		{
			s := "0" // 0 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xc1: // 0xc1 - SET
		o := &SET_0_C{}

		{
			s := "0" // 0 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xc2: // 0xc2 - SET
		o := &SET_0_D{}

		{
			s := "0" // 0 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xc3: // 0xc3 - SET
		o := &SET_0_E{}

		{
			s := "0" // 0 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xc4: // 0xc4 - SET
		o := &SET_0_H{}

		{
			s := "0" // 0 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xc5: // 0xc5 - SET
		o := &SET_0_L{}

		{
			s := "0" // 0 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xc6: // 0xc6 - SET
		o := &SET_0_HLPtr{}

		{
			s := "0" // 0 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xc7: // 0xc7 - SET
		o := &SET_0_A{}

		{
			s := "0" // 0 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xc8: // 0xc8 - SET
		o := &SET_1_B{}

		{
			s := "1" // 1 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xc9: // 0xc9 - SET
		o := &SET_1_C{}

		{
			s := "1" // 1 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xca: // 0xca - SET
		o := &SET_1_D{}

		{
			s := "1" // 1 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xcb: // 0xcb - SET
		o := &SET_1_E{}

		{
			s := "1" // 1 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xcc: // 0xcc - SET
		o := &SET_1_H{}

		{
			s := "1" // 1 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xcd: // 0xcd - SET
		o := &SET_1_L{}

		{
			s := "1" // 1 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xce: // 0xce - SET
		o := &SET_1_HLPtr{}

		{
			s := "1" // 1 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xcf: // 0xcf - SET
		o := &SET_1_A{}

		{
			s := "1" // 1 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xd: // 0xd - RRC
		o := &RRC_L{}

		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xd0: // 0xd0 - SET
		o := &SET_2_B{}

		{
			s := "2" // 2 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xd1: // 0xd1 - SET
		o := &SET_2_C{}

		{
			s := "2" // 2 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xd2: // 0xd2 - SET
		o := &SET_2_D{}

		{
			s := "2" // 2 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xd3: // 0xd3 - SET
		o := &SET_2_E{}

		{
			s := "2" // 2 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xd4: // 0xd4 - SET
		o := &SET_2_H{}

		{
			s := "2" // 2 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xd5: // 0xd5 - SET
		o := &SET_2_L{}

		{
			s := "2" // 2 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xd6: // 0xd6 - SET
		o := &SET_2_HLPtr{}

		{
			s := "2" // 2 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xd7: // 0xd7 - SET
		o := &SET_2_A{}

		{
			s := "2" // 2 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xd8: // 0xd8 - SET
		o := &SET_3_B{}

		{
			s := "3" // 3 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xd9: // 0xd9 - SET
		o := &SET_3_C{}

		{
			s := "3" // 3 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xda: // 0xda - SET
		o := &SET_3_D{}

		{
			s := "3" // 3 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xdb: // 0xdb - SET
		o := &SET_3_E{}

		{
			s := "3" // 3 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xdc: // 0xdc - SET
		o := &SET_3_H{}

		{
			s := "3" // 3 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xdd: // 0xdd - SET
		o := &SET_3_L{}

		{
			s := "3" // 3 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xde: // 0xde - SET
		o := &SET_3_HLPtr{}

		{
			s := "3" // 3 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xdf: // 0xdf - SET
		o := &SET_3_A{}

		{
			s := "3" // 3 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xe: // 0xe - RRC
		o := &RRC_HLPtr{}

		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xe0: // 0xe0 - SET
		o := &SET_4_B{}

		{
			s := "4" // 4 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xe1: // 0xe1 - SET
		o := &SET_4_C{}

		{
			s := "4" // 4 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xe2: // 0xe2 - SET
		o := &SET_4_D{}

		{
			s := "4" // 4 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xe3: // 0xe3 - SET
		o := &SET_4_E{}

		{
			s := "4" // 4 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xe4: // 0xe4 - SET
		o := &SET_4_H{}

		{
			s := "4" // 4 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xe5: // 0xe5 - SET
		o := &SET_4_L{}

		{
			s := "4" // 4 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xe6: // 0xe6 - SET
		o := &SET_4_HLPtr{}

		{
			s := "4" // 4 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xe7: // 0xe7 - SET
		o := &SET_4_A{}

		{
			s := "4" // 4 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xe8: // 0xe8 - SET
		o := &SET_5_B{}

		{
			s := "5" // 5 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xe9: // 0xe9 - SET
		o := &SET_5_C{}

		{
			s := "5" // 5 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xea: // 0xea - SET
		o := &SET_5_D{}

		{
			s := "5" // 5 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xeb: // 0xeb - SET
		o := &SET_5_E{}

		{
			s := "5" // 5 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xec: // 0xec - SET
		o := &SET_5_H{}

		{
			s := "5" // 5 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xed: // 0xed - SET
		o := &SET_5_L{}

		{
			s := "5" // 5 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xee: // 0xee - SET
		o := &SET_5_HLPtr{}

		{
			s := "5" // 5 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xef: // 0xef - SET
		o := &SET_5_A{}

		{
			s := "5" // 5 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xf: // 0xf - RRC
		o := &RRC_A{}

		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "" //  is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xf0: // 0xf0 - SET
		o := &SET_6_B{}

		{
			s := "6" // 6 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xf1: // 0xf1 - SET
		o := &SET_6_C{}

		{
			s := "6" // 6 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xf2: // 0xf2 - SET
		o := &SET_6_D{}

		{
			s := "6" // 6 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xf3: // 0xf3 - SET
		o := &SET_6_E{}

		{
			s := "6" // 6 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xf4: // 0xf4 - SET
		o := &SET_6_H{}

		{
			s := "6" // 6 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xf5: // 0xf5 - SET
		o := &SET_6_L{}

		{
			s := "6" // 6 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xf6: // 0xf6 - SET
		o := &SET_6_HLPtr{}

		{
			s := "6" // 6 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xf7: // 0xf7 - SET
		o := &SET_6_A{}

		{
			s := "6" // 6 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xf8: // 0xf8 - SET
		o := &SET_7_B{}

		{
			s := "7" // 7 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "B" // B is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xf9: // 0xf9 - SET
		o := &SET_7_C{}

		{
			s := "7" // 7 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "C" // C is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xfa: // 0xfa - SET
		o := &SET_7_D{}

		{
			s := "7" // 7 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "D" // D is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xfb: // 0xfb - SET
		o := &SET_7_E{}

		{
			s := "7" // 7 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "E" // E is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xfc: // 0xfc - SET
		o := &SET_7_H{}

		{
			s := "7" // 7 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "H" // H is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xfd: // 0xfd - SET
		o := &SET_7_L{}

		{
			s := "7" // 7 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "L" // L is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xfe: // 0xfe - SET
		o := &SET_7_HLPtr{}

		{
			s := "7" // 7 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "(HL)" // (HL) is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	case 0xff: // 0xff - SET
		o := &SET_7_A{}

		{
			s := "7" // 7 is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand1 = s
		}
		{
			s := "A" // A is implicit in this instruction.
			if err != nil {
				return nil, err
			}
			o.operand2 = s
		}

		return o, nil

	default:
		return nil, fmt.Errorf("the proposed opcode (dec %d, hex %x) doesn't exist: %v", d[0], d[0], ErrNoOpCode)
	}

}
