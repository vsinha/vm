package opcodes

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/vsinha/vm/internal/memory"
	"github.com/vsinha/vm/internal/registers"
)

// ErrUnimplemented is the unimplemented error
var ErrUnimplemented = errors.New("this opcode has not yet been implemented")

// ErrHalt is the control flow error for when the VM should be stoppped.
var ErrHalt = errors.New("HALTED")

// ErrOperatorNotValidType is returned when the in memory representation of the
// operator doesn't match the opcode's required type.
var ErrOperatorNotValidType = errors.New("operator was not the expected type for this opcode")

type ExecutionResult struct {
	Cycles uint8

	// DidSetPC tells the VM to not touch the PC value after executing
	// this instruction, whatever needed to be done has been handled in
	// Execute. The default of false will tell the VM to increment the PC
	// based off the length of the instruction.
	DidSetPC bool
}

// Instruction represents an instruction
type Instruction interface {
	Execute(vm) (ExecutionResult, error)
	String() string
	Write(io.Writer) (int, error)

	cycles() []uint8

	// Length describes the number of bytes in memory that were used to describe this instruction.
	Length() uint8
}

type vm interface {
	Mem() memory.Memory
	Reg() *registers.Registers
}

var endianness = binary.BigEndian

func readImmediate16BitAddress(r io.Reader) (uint16, error) {
	return readImmediate16BitData(r)
}
func readImmediate16BitData(r io.Reader) (uint16, error) {
	var v uint16
	err := binary.Read(r, endianness, &v)
	return v, err
}
func readImmediate8BitAddress(r io.Reader) (uint8, error) {
	return readImmediate8BitData(r)
}
func readImmediate8BitData(r io.Reader) (uint8, error) {
	var v uint8
	err := binary.Read(r, endianness, &v)
	return v, err
}
func readImmediateSigned8BitData(r io.Reader) (int8, error) {
	var v int8
	err := binary.Read(r, endianness, &v)
	return v, err
}

func readBytesAsString(r io.Reader, n int) (string, error) {
	d := make([]byte, n)

	readBytes, err := r.Read(d)
	if err != nil {
		return "", err
	}
	if readBytes != n {
		return "", fmt.Errorf("wanted to read %d bytes but only read %d", n, readBytes)
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

func addInto(v vm, i Instruction, dest *registers.Reg, a registers.Reg, b registers.Reg) (ExecutionResult, error) {
	v.Reg().A = a + b
	// TODO set carry bits
	return ExecutionResult{Cycles: i.cycles()[0]}, nil
}

func addIntoA(v vm, i Instruction, a registers.Reg, b registers.Reg) (ExecutionResult, error) {
	return addInto(v, i, &v.Reg().A, a, b)
}

// Execute NOP instruction. 0x0
func (i *NOP) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, nil
}

// Execute LD_BC_d16 instruction. 0x1
func (i *LD_BC_d16) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute STOP_0 instruction. 0x10
func (i *STOP_0) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_DE_d16 instruction.
func (i *LD_DE_d16) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_DEDeref_A instruction.
func (i *LD_DEDeref_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute INC_DE instruction.
func (i *INC_DE) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute INC_D instruction.
func (i *INC_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute DEC_D instruction.
func (i *DEC_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_D_d8 instruction.
func (i *LD_D_d8) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RLA instruction.
func (i *RLA) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute JR_r8 instruction. JR is a relative jump between 128
// addresses forward or backwards.
func (i *JR_r8) Execute(v vm) (ExecutionResult, error) {
	// vm.Reg().PC += i.operand1.(int8)
	// return ExecutionResult{Cycles: i.cycles()[0], DidSDidSetPC: true}, ErrUnimplemented
	return ExecutionResult{}, ErrUnimplemented
}

// Execute ADD_HL_DE instruction.
func (i *ADD_HL_DE) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_A_DEDeref instruction.
func (i *LD_A_DEDeref) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute DEC_DE instruction.
func (i *DEC_DE) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute INC_E instruction.
func (i *INC_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute DEC_E instruction.
func (i *DEC_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_E_d8 instruction.
func (i *LD_E_d8) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RRA instruction.
func (i *RRA) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_BCDeref_A instruction.
func (i *LD_BCDeref_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute JR_NZ_r8 instruction.
func (i *JR_NZ_r8) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_HL_d16 instruction.
func (i *LD_HL_d16) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_HLPtrInc_A instruction.
func (i *LD_HLPtrInc_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute INC_HL instruction.
func (i *INC_HL) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute INC_H instruction.
func (i *INC_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute DEC_H instruction.
func (i *DEC_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_H_d8 instruction.
func (i *LD_H_d8) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute DAA instruction.
func (i *DAA) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute JR_Z_r8 instruction.
func (i *JR_Z_r8) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute ADD_HL_HL instruction.
func (i *ADD_HL_HL) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_A_HLPtrInc instruction.
func (i *LD_A_HLPtrInc) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute DEC_HL instruction.
func (i *DEC_HL) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute INC_L instruction.
func (i *INC_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute DEC_L instruction.
func (i *DEC_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_L_d8 instruction.
func (i *LD_L_d8) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute CPL instruction.
func (i *CPL) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute INC_BC instruction.
func (i *INC_BC) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute JR_NC_r8 instruction.
func (i *JR_NC_r8) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_SP_d16 instruction.
func (i *LD_SP_d16) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_HLPtrDec_A instruction.
func (i *LD_HLPtrDec_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute INC_SP instruction.
func (i *INC_SP) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute INC_HLPtr instruction.
func (i *INC_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute DEC_HLPtr instruction.
func (i *DEC_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_HLPtr_d8 instruction.
func (i *LD_HLPtr_d8) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SCF instruction.
func (i *SCF) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute JR_C_r8 instruction.
func (i *JR_C_r8) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute ADD_HL_SP instruction.
func (i *ADD_HL_SP) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_A_HLPtrDec instruction.
func (i *LD_A_HLPtrDec) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute DEC_SP instruction.
func (i *DEC_SP) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute INC_A instruction.
func (i *INC_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute DEC_A instruction.
func (i *DEC_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_A_d8 instruction.
func (i *LD_A_d8) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute CCF instruction.
func (i *CCF) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute INC_B instruction.
func (i *INC_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_B_B instruction.
func (i *LD_B_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_B_C instruction.
func (i *LD_B_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_B_D instruction.
func (i *LD_B_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_B_E instruction.
func (i *LD_B_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_B_H instruction.
func (i *LD_B_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_B_L instruction.
func (i *LD_B_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_B_HLPtr instruction.
func (i *LD_B_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_B_A instruction.
func (i *LD_B_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_C_B instruction.
func (i *LD_C_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_C_C instruction.
func (i *LD_C_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_C_D instruction.
func (i *LD_C_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_C_E instruction.
func (i *LD_C_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_C_H instruction.
func (i *LD_C_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_C_L instruction.
func (i *LD_C_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_C_HLPtr instruction.
func (i *LD_C_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_C_A instruction.
func (i *LD_C_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute DEC_B instruction.
func (i *DEC_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_D_B instruction.
func (i *LD_D_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_D_C instruction.
func (i *LD_D_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_D_D instruction.
func (i *LD_D_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_D_E instruction.
func (i *LD_D_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_D_H instruction.
func (i *LD_D_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_D_L instruction.
func (i *LD_D_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_D_HLPtr instruction.
func (i *LD_D_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_D_A instruction.
func (i *LD_D_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_E_B instruction.
func (i *LD_E_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_E_C instruction.
func (i *LD_E_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_E_D instruction.
func (i *LD_E_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_E_E instruction.
func (i *LD_E_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_E_H instruction.
func (i *LD_E_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_E_L instruction.
func (i *LD_E_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_E_HLPtr instruction.
func (i *LD_E_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_E_A instruction.
func (i *LD_E_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_B_d8 instruction.
func (i *LD_B_d8) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_H_B instruction.
func (i *LD_H_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_H_C instruction.
func (i *LD_H_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_H_D instruction.
func (i *LD_H_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_H_E instruction.
func (i *LD_H_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_H_H instruction.
func (i *LD_H_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_H_L instruction.
func (i *LD_H_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_H_HLPtr instruction.
func (i *LD_H_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_H_A instruction.
func (i *LD_H_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_L_B instruction.
func (i *LD_L_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_L_C instruction.
func (i *LD_L_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_L_D instruction.
func (i *LD_L_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_L_E instruction.
func (i *LD_L_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_L_H instruction.
func (i *LD_L_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_L_L instruction.
func (i *LD_L_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_L_HLPtr instruction.
func (i *LD_L_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_L_A instruction.
func (i *LD_L_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RLCA instruction.
func (i *RLCA) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_HLPtr_B instruction.
func (i *LD_HLPtr_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_HLPtr_C instruction.
func (i *LD_HLPtr_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_HLPtr_D instruction.
func (i *LD_HLPtr_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_HLPtr_E instruction.
func (i *LD_HLPtr_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_HLPtr_H instruction.
func (i *LD_HLPtr_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_HLPtr_L instruction.
func (i *LD_HLPtr_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute Halt.
func (i *HALT) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, fmt.Errorf("HALTED")
}

// Execute LD_HLPtr_A instruction.
func (i *LD_HLPtr_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_A_B instruction.
func (i *LD_A_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_A_C instruction.
func (i *LD_A_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_A_D instruction.
func (i *LD_A_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_A_E instruction.
func (i *LD_A_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_A_H instruction.
func (i *LD_A_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_A_L instruction.
func (i *LD_A_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_A_HLPtr instruction.
func (i *LD_A_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_A_A instruction.
func (i *LD_A_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_a16Deref_SP instruction.
func (i *LD_a16Deref_SP) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute ADD_A_B instruction.
func (i *ADD_A_B) Execute(v vm) (ExecutionResult, error) {
	return addIntoA(v, i, v.Reg().A, v.Reg().B)
}

// Execute ADD_A_C instruction.
func (i *ADD_A_C) Execute(v vm) (ExecutionResult, error) {
	return addIntoA(v, i, v.Reg().A, v.Reg().C)
}

// Execute ADD_A_D instruction.
func (i *ADD_A_D) Execute(v vm) (ExecutionResult, error) {
	return addIntoA(v, i, v.Reg().A, v.Reg().D)
}

// Execute ADD_A_E instruction.
func (i *ADD_A_E) Execute(v vm) (ExecutionResult, error) {
	return addIntoA(v, i, v.Reg().A, v.Reg().E)
}

// Execute ADD_A_H instruction.
func (i *ADD_A_H) Execute(v vm) (ExecutionResult, error) {
	return addIntoA(v, i, v.Reg().A, v.Reg().H)
}

// Execute ADD_A_L instruction.
func (i *ADD_A_L) Execute(v vm) (ExecutionResult, error) {
	return addIntoA(v, i, v.Reg().A, v.Reg().L)
}

// Execute ADD_A_HLPtr instruction.
func (i *ADD_A_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute ADD_A_A instruction.
func (i *ADD_A_A) Execute(v vm) (ExecutionResult, error) {
	return addIntoA(v, i, v.Reg().A, v.Reg().A)
}

// Execute ADC_A_B instruction.
func (i *ADC_A_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute ADC_A_C instruction.
func (i *ADC_A_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute ADC_A_D instruction.
func (i *ADC_A_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute ADC_A_E instruction.
func (i *ADC_A_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute ADC_A_H instruction.
func (i *ADC_A_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute ADC_A_L instruction.
func (i *ADC_A_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute ADC_A_HLPtr instruction.
func (i *ADC_A_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute ADC_A_A instruction.
func (i *ADC_A_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute ADD_HL_BC instruction.
func (i *ADD_HL_BC) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SUB_B instruction.
func (i *SUB_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SUB_C instruction.
func (i *SUB_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SUB_D instruction.
func (i *SUB_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SUB_E instruction.
func (i *SUB_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SUB_H instruction.
func (i *SUB_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SUB_L instruction.
func (i *SUB_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SUB_HLPtr instruction.
func (i *SUB_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SUB_A instruction.
func (i *SUB_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SBC_A_B instruction.
func (i *SBC_A_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SBC_A_C instruction.
func (i *SBC_A_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SBC_A_D instruction.
func (i *SBC_A_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SBC_A_E instruction.
func (i *SBC_A_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SBC_A_H instruction.
func (i *SBC_A_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SBC_A_L instruction.
func (i *SBC_A_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SBC_A_HLPtr instruction.
func (i *SBC_A_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SBC_A_A instruction.
func (i *SBC_A_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_A_BCDeref instruction.
func (i *LD_A_BCDeref) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute AND_B instruction.
func (i *AND_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute AND_C instruction.
func (i *AND_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute AND_D instruction.
func (i *AND_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute AND_E instruction.
func (i *AND_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute AND_H instruction.
func (i *AND_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute AND_L instruction.
func (i *AND_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute AND_HLPtr instruction.
func (i *AND_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute AND_A instruction.
func (i *AND_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute XOR_B instruction.
func (i *XOR_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute XOR_C instruction.
func (i *XOR_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute XOR_D instruction.
func (i *XOR_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute XOR_E instruction.
func (i *XOR_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute XOR_H instruction.
func (i *XOR_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute XOR_L instruction.
func (i *XOR_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute XOR_HLPtr instruction.
func (i *XOR_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute XOR_A instruction.
func (i *XOR_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute DEC_BC instruction.
func (i *DEC_BC) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute OR_B instruction.
func (i *OR_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute OR_C instruction.
func (i *OR_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute OR_D instruction.
func (i *OR_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute OR_E instruction.
func (i *OR_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute OR_H instruction.
func (i *OR_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute OR_L instruction.
func (i *OR_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute OR_HLPtr instruction.
func (i *OR_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute OR_A instruction.
func (i *OR_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute CP_B instruction.
func (i *CP_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute CP_C instruction.
func (i *CP_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute CP_D instruction.
func (i *CP_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute CP_E instruction.
func (i *CP_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute CP_H instruction.
func (i *CP_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute CP_L instruction.
func (i *CP_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute CP_HLPtr instruction.
func (i *CP_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute CP_A instruction.
func (i *CP_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute INC_C instruction.
func (i *INC_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RET_NZ instruction.
func (i *RET_NZ) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute POP_BC instruction.
func (i *POP_BC) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute JP_NZ_a16 instruction.
func (i *JP_NZ_a16) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute JP_a16 instruction.
func (i *JP_a16) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute CALL_NZ_a16 instruction.
func (i *CALL_NZ_a16) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute PUSH_BC instruction.
func (i *PUSH_BC) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute ADD_A_d8 instruction.
func (i *ADD_A_d8) Execute(v vm) (ExecutionResult, error) {
	d, ok := i.operand2.(uint8)
	if !ok {
		return ExecutionResult{}, ErrOperatorNotValidType
	}

	return addIntoA(v, i, v.Reg().A, registers.Reg(d))
}

// Execute RST_00H instruction.
func (i *RST_00H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RET_Z instruction.
func (i *RET_Z) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RET instruction.
func (i *RET) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute JP_Z_a16 instruction.
func (i *JP_Z_a16) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute PREFIX_CB instruction.
func (i *PREFIX_CB) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute CALL_Z_a16 instruction.
func (i *CALL_Z_a16) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute CALL_a16 instruction.
func (i *CALL_a16) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute ADC_A_d8 instruction.
func (i *ADC_A_d8) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RST_08H instruction.
func (i *RST_08H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute DEC_C instruction.
func (i *DEC_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RET_NC instruction.
func (i *RET_NC) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute POP_DE instruction.
func (i *POP_DE) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute JP_NC_a16 instruction.
func (i *JP_NC_a16) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute CALL_NC_a16 instruction.
func (i *CALL_NC_a16) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute PUSH_DE instruction.
func (i *PUSH_DE) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SUB_d8 instruction.
func (i *SUB_d8) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RST_10H instruction.
func (i *RST_10H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RET_C instruction.
func (i *RET_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RETI instruction.
func (i *RETI) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute JP_C_a16 instruction.
func (i *JP_C_a16) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute CALL_C_a16 instruction.
func (i *CALL_C_a16) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SBC_A_d8 instruction.
func (i *SBC_A_d8) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RST_18H instruction.
func (i *RST_18H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_C_d8 instruction.
func (i *LD_C_d8) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LDH_a8Deref_A instruction.
func (i *LDH_a8Deref_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute POP_HL instruction.
func (i *POP_HL) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_CDeref_A instruction.
func (i *LD_CDeref_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute PUSH_HL instruction.
func (i *PUSH_HL) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute AND_d8 instruction.
func (i *AND_d8) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RST_20H instruction.
func (i *RST_20H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute ADD_SP_r8 instruction.
func (i *ADD_SP_r8) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute JP_HLPtr instruction.
func (i *JP_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_a16Deref_A instruction.
func (i *LD_a16Deref_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute XOR_d8 instruction.
func (i *XOR_d8) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RST_28H instruction.
func (i *RST_28H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RRCA instruction.
func (i *RRCA) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LDH_A_a8Deref instruction.
func (i *LDH_A_a8Deref) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute POP_AF instruction.
func (i *POP_AF) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_A_CDeref instruction.
func (i *LD_A_CDeref) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute DI instruction.
func (i *DI) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute PUSH_AF instruction.
func (i *PUSH_AF) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute OR_d8 instruction.
func (i *OR_d8) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RST_30H instruction.
func (i *RST_30H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_HL_SP_plus_r8 instruction.
func (i *LD_HL_SP_plus_r8) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_SP_HL instruction.
func (i *LD_SP_HL) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute LD_A_a16Deref instruction.
func (i *LD_A_a16Deref) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute EI instruction.
func (i *EI) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute CP_d8 instruction.
func (i *CP_d8) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RST_38H instruction.
func (i *RST_38H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RLC_B instruction.
func (i *RLC_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RLC_C instruction.
func (i *RLC_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RL_B instruction.
func (i *RL_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RL_C instruction.
func (i *RL_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RL_D instruction.
func (i *RL_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RL_E instruction.
func (i *RL_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RL_H instruction.
func (i *RL_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RL_L instruction.
func (i *RL_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RL_HLPtr instruction.
func (i *RL_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RL_A instruction.
func (i *RL_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RR_B instruction.
func (i *RR_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RR_C instruction.
func (i *RR_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RR_D instruction.
func (i *RR_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RR_E instruction.
func (i *RR_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RR_H instruction.
func (i *RR_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RR_L instruction.
func (i *RR_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RR_HLPtr instruction.
func (i *RR_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RR_A instruction.
func (i *RR_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RLC_D instruction.
func (i *RLC_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SLA_B instruction.
func (i *SLA_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SLA_C instruction.
func (i *SLA_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SLA_D instruction.
func (i *SLA_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SLA_E instruction.
func (i *SLA_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SLA_H instruction.
func (i *SLA_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SLA_L instruction.
func (i *SLA_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SLA_HLPtr instruction.
func (i *SLA_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SLA_A instruction.
func (i *SLA_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SRA_B instruction.
func (i *SRA_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SRA_C instruction.
func (i *SRA_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SRA_D instruction.
func (i *SRA_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SRA_E instruction.
func (i *SRA_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SRA_H instruction.
func (i *SRA_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SRA_L instruction.
func (i *SRA_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SRA_HLPtr instruction.
func (i *SRA_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SRA_A instruction.
func (i *SRA_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RLC_E instruction.
func (i *RLC_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SWAP_B instruction.
func (i *SWAP_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SWAP_C instruction.
func (i *SWAP_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SWAP_D instruction.
func (i *SWAP_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SWAP_E instruction.
func (i *SWAP_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SWAP_H instruction.
func (i *SWAP_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SWAP_L instruction.
func (i *SWAP_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SWAP_HLPtr instruction.
func (i *SWAP_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SWAP_A instruction.
func (i *SWAP_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SRL_B instruction.
func (i *SRL_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SRL_C instruction.
func (i *SRL_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SRL_D instruction.
func (i *SRL_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SRL_E instruction.
func (i *SRL_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SRL_H instruction.
func (i *SRL_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SRL_L instruction.
func (i *SRL_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SRL_HLPtr instruction.
func (i *SRL_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SRL_A instruction.
func (i *SRL_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RLC_H instruction.
func (i *RLC_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_0_B instruction.
func (i *BIT_0_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_0_C instruction.
func (i *BIT_0_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_0_D instruction.
func (i *BIT_0_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_0_E instruction.
func (i *BIT_0_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_0_H instruction.
func (i *BIT_0_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_0_L instruction.
func (i *BIT_0_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_0_HLPtr instruction.
func (i *BIT_0_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_0_A instruction.
func (i *BIT_0_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_1_B instruction.
func (i *BIT_1_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_1_C instruction.
func (i *BIT_1_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_1_D instruction.
func (i *BIT_1_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_1_E instruction.
func (i *BIT_1_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_1_H instruction.
func (i *BIT_1_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_1_L instruction.
func (i *BIT_1_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_1_HLPtr instruction.
func (i *BIT_1_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_1_A instruction.
func (i *BIT_1_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RLC_L instruction.
func (i *RLC_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_2_B instruction.
func (i *BIT_2_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_2_C instruction.
func (i *BIT_2_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_2_D instruction.
func (i *BIT_2_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_2_E instruction.
func (i *BIT_2_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_2_H instruction.
func (i *BIT_2_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_2_L instruction.
func (i *BIT_2_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_2_HLPtr instruction.
func (i *BIT_2_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_2_A instruction.
func (i *BIT_2_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_3_B instruction.
func (i *BIT_3_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_3_C instruction.
func (i *BIT_3_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_3_D instruction.
func (i *BIT_3_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_3_E instruction.
func (i *BIT_3_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_3_H instruction.
func (i *BIT_3_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_3_L instruction.
func (i *BIT_3_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_3_HLPtr instruction.
func (i *BIT_3_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_3_A instruction.
func (i *BIT_3_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RLC_HLPtr instruction.
func (i *RLC_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_4_B instruction.
func (i *BIT_4_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_4_C instruction.
func (i *BIT_4_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_4_D instruction.
func (i *BIT_4_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_4_E instruction.
func (i *BIT_4_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_4_H instruction.
func (i *BIT_4_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_4_L instruction.
func (i *BIT_4_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_4_HLPtr instruction.
func (i *BIT_4_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_4_A instruction.
func (i *BIT_4_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_5_B instruction.
func (i *BIT_5_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_5_C instruction.
func (i *BIT_5_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_5_D instruction.
func (i *BIT_5_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_5_E instruction.
func (i *BIT_5_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_5_H instruction.
func (i *BIT_5_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_5_L instruction.
func (i *BIT_5_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_5_HLPtr instruction.
func (i *BIT_5_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_5_A instruction.
func (i *BIT_5_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RLC_A instruction.
func (i *RLC_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_6_B instruction.
func (i *BIT_6_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_6_C instruction.
func (i *BIT_6_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_6_D instruction.
func (i *BIT_6_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_6_E instruction.
func (i *BIT_6_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_6_H instruction.
func (i *BIT_6_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_6_L instruction.
func (i *BIT_6_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_6_HLPtr instruction.
func (i *BIT_6_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_6_A instruction.
func (i *BIT_6_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_7_B instruction.
func (i *BIT_7_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_7_C instruction.
func (i *BIT_7_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_7_D instruction.
func (i *BIT_7_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_7_E instruction.
func (i *BIT_7_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_7_H instruction.
func (i *BIT_7_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_7_L instruction.
func (i *BIT_7_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_7_HLPtr instruction.
func (i *BIT_7_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute BIT_7_A instruction.
func (i *BIT_7_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RRC_B instruction.
func (i *RRC_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_0_B instruction.
func (i *RES_0_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_0_C instruction.
func (i *RES_0_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_0_D instruction.
func (i *RES_0_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_0_E instruction.
func (i *RES_0_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_0_H instruction.
func (i *RES_0_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_0_L instruction.
func (i *RES_0_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_0_HLPtr instruction.
func (i *RES_0_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_0_A instruction.
func (i *RES_0_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_1_B instruction.
func (i *RES_1_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_1_C instruction.
func (i *RES_1_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_1_D instruction.
func (i *RES_1_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_1_E instruction.
func (i *RES_1_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_1_H instruction.
func (i *RES_1_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_1_L instruction.
func (i *RES_1_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_1_HLPtr instruction.
func (i *RES_1_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_1_A instruction.
func (i *RES_1_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RRC_C instruction.
func (i *RRC_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_2_B instruction.
func (i *RES_2_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_2_C instruction.
func (i *RES_2_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_2_D instruction.
func (i *RES_2_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_2_E instruction.
func (i *RES_2_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_2_H instruction.
func (i *RES_2_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_2_L instruction.
func (i *RES_2_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_2_HLPtr instruction.
func (i *RES_2_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_2_A instruction.
func (i *RES_2_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_3_B instruction.
func (i *RES_3_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_3_C instruction.
func (i *RES_3_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_3_D instruction.
func (i *RES_3_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_3_E instruction.
func (i *RES_3_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_3_H instruction.
func (i *RES_3_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_3_L instruction.
func (i *RES_3_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_3_HLPtr instruction.
func (i *RES_3_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_3_A instruction.
func (i *RES_3_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RRC_D instruction.
func (i *RRC_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_4_B instruction.
func (i *RES_4_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_4_C instruction.
func (i *RES_4_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_4_D instruction.
func (i *RES_4_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_4_E instruction.
func (i *RES_4_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_4_H instruction.
func (i *RES_4_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_4_L instruction.
func (i *RES_4_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_4_HLPtr instruction.
func (i *RES_4_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_4_A instruction.
func (i *RES_4_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_5_B instruction.
func (i *RES_5_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_5_C instruction.
func (i *RES_5_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_5_D instruction.
func (i *RES_5_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_5_E instruction.
func (i *RES_5_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_5_H instruction.
func (i *RES_5_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_5_L instruction.
func (i *RES_5_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_5_HLPtr instruction.
func (i *RES_5_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_5_A instruction.
func (i *RES_5_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RRC_E instruction.
func (i *RRC_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_6_B instruction.
func (i *RES_6_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_6_C instruction.
func (i *RES_6_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_6_D instruction.
func (i *RES_6_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_6_E instruction.
func (i *RES_6_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_6_H instruction.
func (i *RES_6_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_6_L instruction.
func (i *RES_6_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_6_HLPtr instruction.
func (i *RES_6_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_6_A instruction.
func (i *RES_6_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_7_B instruction.
func (i *RES_7_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_7_C instruction.
func (i *RES_7_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_7_D instruction.
func (i *RES_7_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_7_E instruction.
func (i *RES_7_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_7_H instruction.
func (i *RES_7_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_7_L instruction.
func (i *RES_7_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_7_HLPtr instruction.
func (i *RES_7_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RES_7_A instruction.
func (i *RES_7_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RRC_H instruction.
func (i *RRC_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_0_B instruction.
func (i *SET_0_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_0_C instruction.
func (i *SET_0_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_0_D instruction.
func (i *SET_0_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_0_E instruction.
func (i *SET_0_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_0_H instruction.
func (i *SET_0_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_0_L instruction.
func (i *SET_0_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_0_HLPtr instruction.
func (i *SET_0_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_0_A instruction.
func (i *SET_0_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_1_B instruction.
func (i *SET_1_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_1_C instruction.
func (i *SET_1_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_1_D instruction.
func (i *SET_1_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_1_E instruction.
func (i *SET_1_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_1_H instruction.
func (i *SET_1_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_1_L instruction.
func (i *SET_1_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_1_HLPtr instruction.
func (i *SET_1_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_1_A instruction.
func (i *SET_1_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RRC_L instruction.
func (i *RRC_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_2_B instruction.
func (i *SET_2_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_2_C instruction.
func (i *SET_2_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_2_D instruction.
func (i *SET_2_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_2_E instruction.
func (i *SET_2_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_2_H instruction.
func (i *SET_2_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_2_L instruction.
func (i *SET_2_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_2_HLPtr instruction.
func (i *SET_2_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_2_A instruction.
func (i *SET_2_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_3_B instruction.
func (i *SET_3_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_3_C instruction.
func (i *SET_3_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_3_D instruction.
func (i *SET_3_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_3_E instruction.
func (i *SET_3_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_3_H instruction.
func (i *SET_3_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_3_L instruction.
func (i *SET_3_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_3_HLPtr instruction.
func (i *SET_3_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_3_A instruction.
func (i *SET_3_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RRC_HLPtr instruction.
func (i *RRC_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_4_B instruction.
func (i *SET_4_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_4_C instruction.
func (i *SET_4_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_4_D instruction.
func (i *SET_4_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_4_E instruction.
func (i *SET_4_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_4_H instruction.
func (i *SET_4_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_4_L instruction.
func (i *SET_4_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_4_HLPtr instruction.
func (i *SET_4_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_4_A instruction.
func (i *SET_4_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_5_B instruction.
func (i *SET_5_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_5_C instruction.
func (i *SET_5_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_5_D instruction.
func (i *SET_5_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_5_E instruction.
func (i *SET_5_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_5_H instruction.
func (i *SET_5_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_5_L instruction.
func (i *SET_5_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_5_HLPtr instruction.
func (i *SET_5_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_5_A instruction.
func (i *SET_5_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute RRC_A instruction.
func (i *RRC_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_6_B instruction.
func (i *SET_6_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_6_C instruction.
func (i *SET_6_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_6_D instruction.
func (i *SET_6_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_6_E instruction.
func (i *SET_6_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_6_H instruction.
func (i *SET_6_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_6_L instruction.
func (i *SET_6_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_6_HLPtr instruction.
func (i *SET_6_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_6_A instruction.
func (i *SET_6_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_7_B instruction.
func (i *SET_7_B) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_7_C instruction.
func (i *SET_7_C) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_7_D instruction.
func (i *SET_7_D) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_7_E instruction.
func (i *SET_7_E) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_7_H instruction.
func (i *SET_7_H) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_7_L instruction.
func (i *SET_7_L) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_7_HLPtr instruction.
func (i *SET_7_HLPtr) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}

// Execute SET_7_A instruction.
func (i *SET_7_A) Execute(v vm) (ExecutionResult, error) {
	return ExecutionResult{}, ErrUnimplemented
}
