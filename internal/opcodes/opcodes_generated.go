package opcodes

import (
	"errors"
	"fmt"
	"io"

	"github.com/vsinha/vm/internal/vm"
)

var ErrNoOpCode = errors.New("no opcode with that address exists")

// TODO this has problems because many of the opcodes have the same mnemonic,
// so we have to think about how to disambiguate them or if we even want to
type Op int

type OpCode interface {
	Execute(*vm.VM)
}

type NOP_ struct {
	operand1 string
	operand2 string
}

func (o *NOP_) Execute(v *vm.VM) {
}

func (o *NOP_) String() string {
	return "NOP"
}

type LD_BC_d16 struct {
	operand1 string
	operand2 string
}

func (o *LD_BC_d16) Execute(v *vm.VM) {
}

func (o *LD_BC_d16) String() string {
	return "LD BC o.operand1"
}

type STOP_0 struct {
	operand1 string
	operand2 string
}

func (o *STOP_0) Execute(v *vm.VM) {
}

func (o *STOP_0) String() string {
	return "STOP 0"
}

type LD_DE_d16 struct {
	operand1 string
	operand2 string
}

func (o *LD_DE_d16) Execute(v *vm.VM) {
}

func (o *LD_DE_d16) String() string {
	return "LD DE o.operand1"
}

type LD_DEDeref_A struct {
	operand1 string
	operand2 string
}

func (o *LD_DEDeref_A) Execute(v *vm.VM) {
}

func (o *LD_DEDeref_A) String() string {
	return "LD (DE) A"
}

type INC_DE struct {
	operand1 string
	operand2 string
}

func (o *INC_DE) Execute(v *vm.VM) {
}

func (o *INC_DE) String() string {
	return "INC DE"
}

type INC_D struct {
	operand1 string
	operand2 string
}

func (o *INC_D) Execute(v *vm.VM) {
}

func (o *INC_D) String() string {
	return "INC D"
}

type DEC_D struct {
	operand1 string
	operand2 string
}

func (o *DEC_D) Execute(v *vm.VM) {
}

func (o *DEC_D) String() string {
	return "DEC D"
}

type LD_D_d8 struct {
	operand1 string
	operand2 string
}

func (o *LD_D_d8) Execute(v *vm.VM) {
}

func (o *LD_D_d8) String() string {
	return "LD D d8"
}

type RLA_ struct {
	operand1 string
	operand2 string
}

func (o *RLA_) Execute(v *vm.VM) {
}

func (o *RLA_) String() string {
	return "RLA"
}

type JR_r8 struct {
	operand1 string
	operand2 string
}

func (o *JR_r8) Execute(v *vm.VM) {
}

func (o *JR_r8) String() string {
	return "JR r8"
}

type ADD_HL_DE struct {
	operand1 string
	operand2 string
}

func (o *ADD_HL_DE) Execute(v *vm.VM) {
}

func (o *ADD_HL_DE) String() string {
	return "ADD HL DE"
}

type LD_A_DEDeref struct {
	operand1 string
	operand2 string
}

func (o *LD_A_DEDeref) Execute(v *vm.VM) {
}

func (o *LD_A_DEDeref) String() string {
	return "LD A (DE)"
}

type DEC_DE struct {
	operand1 string
	operand2 string
}

func (o *DEC_DE) Execute(v *vm.VM) {
}

func (o *DEC_DE) String() string {
	return "DEC DE"
}

type INC_E struct {
	operand1 string
	operand2 string
}

func (o *INC_E) Execute(v *vm.VM) {
}

func (o *INC_E) String() string {
	return "INC E"
}

type DEC_E struct {
	operand1 string
	operand2 string
}

func (o *DEC_E) Execute(v *vm.VM) {
}

func (o *DEC_E) String() string {
	return "DEC E"
}

type LD_E_d8 struct {
	operand1 string
	operand2 string
}

func (o *LD_E_d8) Execute(v *vm.VM) {
}

func (o *LD_E_d8) String() string {
	return "LD E d8"
}

type RRA_ struct {
	operand1 string
	operand2 string
}

func (o *RRA_) Execute(v *vm.VM) {
}

func (o *RRA_) String() string {
	return "RRA"
}

type LD_BCDeref_A struct {
	operand1 string
	operand2 string
}

func (o *LD_BCDeref_A) Execute(v *vm.VM) {
}

func (o *LD_BCDeref_A) String() string {
	return "LD (BC) A"
}

type JR_NZ_r8 struct {
	operand1 string
	operand2 string
}

func (o *JR_NZ_r8) Execute(v *vm.VM) {
}

func (o *JR_NZ_r8) String() string {
	return "JR NZ r8"
}

type LD_HL_d16 struct {
	operand1 string
	operand2 string
}

func (o *LD_HL_d16) Execute(v *vm.VM) {
}

func (o *LD_HL_d16) String() string {
	return "LD HL o.operand1"
}

type LD_HLPtrInc_A struct {
	operand1 string
	operand2 string
}

func (o *LD_HLPtrInc_A) Execute(v *vm.VM) {
}

func (o *LD_HLPtrInc_A) String() string {
	return "LD (HL+) A"
}

type INC_HL struct {
	operand1 string
	operand2 string
}

func (o *INC_HL) Execute(v *vm.VM) {
}

func (o *INC_HL) String() string {
	return "INC HL"
}

type INC_H struct {
	operand1 string
	operand2 string
}

func (o *INC_H) Execute(v *vm.VM) {
}

func (o *INC_H) String() string {
	return "INC H"
}

type DEC_H struct {
	operand1 string
	operand2 string
}

func (o *DEC_H) Execute(v *vm.VM) {
}

func (o *DEC_H) String() string {
	return "DEC H"
}

type LD_H_d8 struct {
	operand1 string
	operand2 string
}

func (o *LD_H_d8) Execute(v *vm.VM) {
}

func (o *LD_H_d8) String() string {
	return "LD H d8"
}

type DAA_ struct {
	operand1 string
	operand2 string
}

func (o *DAA_) Execute(v *vm.VM) {
}

func (o *DAA_) String() string {
	return "DAA"
}

type JR_Z_r8 struct {
	operand1 string
	operand2 string
}

func (o *JR_Z_r8) Execute(v *vm.VM) {
}

func (o *JR_Z_r8) String() string {
	return "JR Z r8"
}

type ADD_HL_HL struct {
	operand1 string
	operand2 string
}

func (o *ADD_HL_HL) Execute(v *vm.VM) {
}

func (o *ADD_HL_HL) String() string {
	return "ADD HL HL"
}

type LD_A_HLPtrInc struct {
	operand1 string
	operand2 string
}

func (o *LD_A_HLPtrInc) Execute(v *vm.VM) {
}

func (o *LD_A_HLPtrInc) String() string {
	return "LD A (HL+)"
}

type DEC_HL struct {
	operand1 string
	operand2 string
}

func (o *DEC_HL) Execute(v *vm.VM) {
}

func (o *DEC_HL) String() string {
	return "DEC HL"
}

type INC_L struct {
	operand1 string
	operand2 string
}

func (o *INC_L) Execute(v *vm.VM) {
}

func (o *INC_L) String() string {
	return "INC L"
}

type DEC_L struct {
	operand1 string
	operand2 string
}

func (o *DEC_L) Execute(v *vm.VM) {
}

func (o *DEC_L) String() string {
	return "DEC L"
}

type LD_L_d8 struct {
	operand1 string
	operand2 string
}

func (o *LD_L_d8) Execute(v *vm.VM) {
}

func (o *LD_L_d8) String() string {
	return "LD L d8"
}

type CPL_ struct {
	operand1 string
	operand2 string
}

func (o *CPL_) Execute(v *vm.VM) {
}

func (o *CPL_) String() string {
	return "CPL"
}

type INC_BC struct {
	operand1 string
	operand2 string
}

func (o *INC_BC) Execute(v *vm.VM) {
}

func (o *INC_BC) String() string {
	return "INC BC"
}

type JR_NC_r8 struct {
	operand1 string
	operand2 string
}

func (o *JR_NC_r8) Execute(v *vm.VM) {
}

func (o *JR_NC_r8) String() string {
	return "JR NC r8"
}

type LD_SP_d16 struct {
	operand1 string
	operand2 string
}

func (o *LD_SP_d16) Execute(v *vm.VM) {
}

func (o *LD_SP_d16) String() string {
	return "LD SP o.operand1"
}

type LD_HLPtrDec_A struct {
	operand1 string
	operand2 string
}

func (o *LD_HLPtrDec_A) Execute(v *vm.VM) {
}

func (o *LD_HLPtrDec_A) String() string {
	return "LD (HL-) A"
}

type INC_SP struct {
	operand1 string
	operand2 string
}

func (o *INC_SP) Execute(v *vm.VM) {
}

func (o *INC_SP) String() string {
	return "INC SP"
}

type INC_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *INC_HLPtr) Execute(v *vm.VM) {
}

func (o *INC_HLPtr) String() string {
	return "INC (HL)"
}

type DEC_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *DEC_HLPtr) Execute(v *vm.VM) {
}

func (o *DEC_HLPtr) String() string {
	return "DEC (HL)"
}

type LD_HLPtr_d8 struct {
	operand1 string
	operand2 string
}

func (o *LD_HLPtr_d8) Execute(v *vm.VM) {
}

func (o *LD_HLPtr_d8) String() string {
	return "LD (HL) d8"
}

type SCF_ struct {
	operand1 string
	operand2 string
}

func (o *SCF_) Execute(v *vm.VM) {
}

func (o *SCF_) String() string {
	return "SCF"
}

type JR_C_r8 struct {
	operand1 string
	operand2 string
}

func (o *JR_C_r8) Execute(v *vm.VM) {
}

func (o *JR_C_r8) String() string {
	return "JR C r8"
}

type ADD_HL_SP struct {
	operand1 string
	operand2 string
}

func (o *ADD_HL_SP) Execute(v *vm.VM) {
}

func (o *ADD_HL_SP) String() string {
	return "ADD HL SP"
}

type LD_A_HLPtrDec struct {
	operand1 string
	operand2 string
}

func (o *LD_A_HLPtrDec) Execute(v *vm.VM) {
}

func (o *LD_A_HLPtrDec) String() string {
	return "LD A (HL-)"
}

type DEC_SP struct {
	operand1 string
	operand2 string
}

func (o *DEC_SP) Execute(v *vm.VM) {
}

func (o *DEC_SP) String() string {
	return "DEC SP"
}

type INC_A struct {
	operand1 string
	operand2 string
}

func (o *INC_A) Execute(v *vm.VM) {
}

func (o *INC_A) String() string {
	return "INC A"
}

type DEC_A struct {
	operand1 string
	operand2 string
}

func (o *DEC_A) Execute(v *vm.VM) {
}

func (o *DEC_A) String() string {
	return "DEC A"
}

type LD_A_d8 struct {
	operand1 string
	operand2 string
}

func (o *LD_A_d8) Execute(v *vm.VM) {
}

func (o *LD_A_d8) String() string {
	return "LD A d8"
}

type CCF_ struct {
	operand1 string
	operand2 string
}

func (o *CCF_) Execute(v *vm.VM) {
}

func (o *CCF_) String() string {
	return "CCF"
}

type INC_B struct {
	operand1 string
	operand2 string
}

func (o *INC_B) Execute(v *vm.VM) {
}

func (o *INC_B) String() string {
	return "INC B"
}

type LD_B_B struct {
	operand1 string
	operand2 string
}

func (o *LD_B_B) Execute(v *vm.VM) {
}

func (o *LD_B_B) String() string {
	return "LD B B"
}

type LD_B_C struct {
	operand1 string
	operand2 string
}

func (o *LD_B_C) Execute(v *vm.VM) {
}

func (o *LD_B_C) String() string {
	return "LD B C"
}

type LD_B_D struct {
	operand1 string
	operand2 string
}

func (o *LD_B_D) Execute(v *vm.VM) {
}

func (o *LD_B_D) String() string {
	return "LD B D"
}

type LD_B_E struct {
	operand1 string
	operand2 string
}

func (o *LD_B_E) Execute(v *vm.VM) {
}

func (o *LD_B_E) String() string {
	return "LD B E"
}

type LD_B_H struct {
	operand1 string
	operand2 string
}

func (o *LD_B_H) Execute(v *vm.VM) {
}

func (o *LD_B_H) String() string {
	return "LD B H"
}

type LD_B_L struct {
	operand1 string
	operand2 string
}

func (o *LD_B_L) Execute(v *vm.VM) {
}

func (o *LD_B_L) String() string {
	return "LD B L"
}

type LD_B_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *LD_B_HLPtr) Execute(v *vm.VM) {
}

func (o *LD_B_HLPtr) String() string {
	return "LD B (HL)"
}

type LD_B_A struct {
	operand1 string
	operand2 string
}

func (o *LD_B_A) Execute(v *vm.VM) {
}

func (o *LD_B_A) String() string {
	return "LD B A"
}

type LD_C_B struct {
	operand1 string
	operand2 string
}

func (o *LD_C_B) Execute(v *vm.VM) {
}

func (o *LD_C_B) String() string {
	return "LD C B"
}

type LD_C_C struct {
	operand1 string
	operand2 string
}

func (o *LD_C_C) Execute(v *vm.VM) {
}

func (o *LD_C_C) String() string {
	return "LD C C"
}

type LD_C_D struct {
	operand1 string
	operand2 string
}

func (o *LD_C_D) Execute(v *vm.VM) {
}

func (o *LD_C_D) String() string {
	return "LD C D"
}

type LD_C_E struct {
	operand1 string
	operand2 string
}

func (o *LD_C_E) Execute(v *vm.VM) {
}

func (o *LD_C_E) String() string {
	return "LD C E"
}

type LD_C_H struct {
	operand1 string
	operand2 string
}

func (o *LD_C_H) Execute(v *vm.VM) {
}

func (o *LD_C_H) String() string {
	return "LD C H"
}

type LD_C_L struct {
	operand1 string
	operand2 string
}

func (o *LD_C_L) Execute(v *vm.VM) {
}

func (o *LD_C_L) String() string {
	return "LD C L"
}

type LD_C_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *LD_C_HLPtr) Execute(v *vm.VM) {
}

func (o *LD_C_HLPtr) String() string {
	return "LD C (HL)"
}

type LD_C_A struct {
	operand1 string
	operand2 string
}

func (o *LD_C_A) Execute(v *vm.VM) {
}

func (o *LD_C_A) String() string {
	return "LD C A"
}

type DEC_B struct {
	operand1 string
	operand2 string
}

func (o *DEC_B) Execute(v *vm.VM) {
}

func (o *DEC_B) String() string {
	return "DEC B"
}

type LD_D_B struct {
	operand1 string
	operand2 string
}

func (o *LD_D_B) Execute(v *vm.VM) {
}

func (o *LD_D_B) String() string {
	return "LD D B"
}

type LD_D_C struct {
	operand1 string
	operand2 string
}

func (o *LD_D_C) Execute(v *vm.VM) {
}

func (o *LD_D_C) String() string {
	return "LD D C"
}

type LD_D_D struct {
	operand1 string
	operand2 string
}

func (o *LD_D_D) Execute(v *vm.VM) {
}

func (o *LD_D_D) String() string {
	return "LD D D"
}

type LD_D_E struct {
	operand1 string
	operand2 string
}

func (o *LD_D_E) Execute(v *vm.VM) {
}

func (o *LD_D_E) String() string {
	return "LD D E"
}

type LD_D_H struct {
	operand1 string
	operand2 string
}

func (o *LD_D_H) Execute(v *vm.VM) {
}

func (o *LD_D_H) String() string {
	return "LD D H"
}

type LD_D_L struct {
	operand1 string
	operand2 string
}

func (o *LD_D_L) Execute(v *vm.VM) {
}

func (o *LD_D_L) String() string {
	return "LD D L"
}

type LD_D_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *LD_D_HLPtr) Execute(v *vm.VM) {
}

func (o *LD_D_HLPtr) String() string {
	return "LD D (HL)"
}

type LD_D_A struct {
	operand1 string
	operand2 string
}

func (o *LD_D_A) Execute(v *vm.VM) {
}

func (o *LD_D_A) String() string {
	return "LD D A"
}

type LD_E_B struct {
	operand1 string
	operand2 string
}

func (o *LD_E_B) Execute(v *vm.VM) {
}

func (o *LD_E_B) String() string {
	return "LD E B"
}

type LD_E_C struct {
	operand1 string
	operand2 string
}

func (o *LD_E_C) Execute(v *vm.VM) {
}

func (o *LD_E_C) String() string {
	return "LD E C"
}

type LD_E_D struct {
	operand1 string
	operand2 string
}

func (o *LD_E_D) Execute(v *vm.VM) {
}

func (o *LD_E_D) String() string {
	return "LD E D"
}

type LD_E_E struct {
	operand1 string
	operand2 string
}

func (o *LD_E_E) Execute(v *vm.VM) {
}

func (o *LD_E_E) String() string {
	return "LD E E"
}

type LD_E_H struct {
	operand1 string
	operand2 string
}

func (o *LD_E_H) Execute(v *vm.VM) {
}

func (o *LD_E_H) String() string {
	return "LD E H"
}

type LD_E_L struct {
	operand1 string
	operand2 string
}

func (o *LD_E_L) Execute(v *vm.VM) {
}

func (o *LD_E_L) String() string {
	return "LD E L"
}

type LD_E_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *LD_E_HLPtr) Execute(v *vm.VM) {
}

func (o *LD_E_HLPtr) String() string {
	return "LD E (HL)"
}

type LD_E_A struct {
	operand1 string
	operand2 string
}

func (o *LD_E_A) Execute(v *vm.VM) {
}

func (o *LD_E_A) String() string {
	return "LD E A"
}

type LD_B_d8 struct {
	operand1 string
	operand2 string
}

func (o *LD_B_d8) Execute(v *vm.VM) {
}

func (o *LD_B_d8) String() string {
	return "LD B d8"
}

type LD_H_B struct {
	operand1 string
	operand2 string
}

func (o *LD_H_B) Execute(v *vm.VM) {
}

func (o *LD_H_B) String() string {
	return "LD H B"
}

type LD_H_C struct {
	operand1 string
	operand2 string
}

func (o *LD_H_C) Execute(v *vm.VM) {
}

func (o *LD_H_C) String() string {
	return "LD H C"
}

type LD_H_D struct {
	operand1 string
	operand2 string
}

func (o *LD_H_D) Execute(v *vm.VM) {
}

func (o *LD_H_D) String() string {
	return "LD H D"
}

type LD_H_E struct {
	operand1 string
	operand2 string
}

func (o *LD_H_E) Execute(v *vm.VM) {
}

func (o *LD_H_E) String() string {
	return "LD H E"
}

type LD_H_H struct {
	operand1 string
	operand2 string
}

func (o *LD_H_H) Execute(v *vm.VM) {
}

func (o *LD_H_H) String() string {
	return "LD H H"
}

type LD_H_L struct {
	operand1 string
	operand2 string
}

func (o *LD_H_L) Execute(v *vm.VM) {
}

func (o *LD_H_L) String() string {
	return "LD H L"
}

type LD_H_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *LD_H_HLPtr) Execute(v *vm.VM) {
}

func (o *LD_H_HLPtr) String() string {
	return "LD H (HL)"
}

type LD_H_A struct {
	operand1 string
	operand2 string
}

func (o *LD_H_A) Execute(v *vm.VM) {
}

func (o *LD_H_A) String() string {
	return "LD H A"
}

type LD_L_B struct {
	operand1 string
	operand2 string
}

func (o *LD_L_B) Execute(v *vm.VM) {
}

func (o *LD_L_B) String() string {
	return "LD L B"
}

type LD_L_C struct {
	operand1 string
	operand2 string
}

func (o *LD_L_C) Execute(v *vm.VM) {
}

func (o *LD_L_C) String() string {
	return "LD L C"
}

type LD_L_D struct {
	operand1 string
	operand2 string
}

func (o *LD_L_D) Execute(v *vm.VM) {
}

func (o *LD_L_D) String() string {
	return "LD L D"
}

type LD_L_E struct {
	operand1 string
	operand2 string
}

func (o *LD_L_E) Execute(v *vm.VM) {
}

func (o *LD_L_E) String() string {
	return "LD L E"
}

type LD_L_H struct {
	operand1 string
	operand2 string
}

func (o *LD_L_H) Execute(v *vm.VM) {
}

func (o *LD_L_H) String() string {
	return "LD L H"
}

type LD_L_L struct {
	operand1 string
	operand2 string
}

func (o *LD_L_L) Execute(v *vm.VM) {
}

func (o *LD_L_L) String() string {
	return "LD L L"
}

type LD_L_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *LD_L_HLPtr) Execute(v *vm.VM) {
}

func (o *LD_L_HLPtr) String() string {
	return "LD L (HL)"
}

type LD_L_A struct {
	operand1 string
	operand2 string
}

func (o *LD_L_A) Execute(v *vm.VM) {
}

func (o *LD_L_A) String() string {
	return "LD L A"
}

type RLCA_ struct {
	operand1 string
	operand2 string
}

func (o *RLCA_) Execute(v *vm.VM) {
}

func (o *RLCA_) String() string {
	return "RLCA"
}

type LD_HLPtr_B struct {
	operand1 string
	operand2 string
}

func (o *LD_HLPtr_B) Execute(v *vm.VM) {
}

func (o *LD_HLPtr_B) String() string {
	return "LD (HL) B"
}

type LD_HLPtr_C struct {
	operand1 string
	operand2 string
}

func (o *LD_HLPtr_C) Execute(v *vm.VM) {
}

func (o *LD_HLPtr_C) String() string {
	return "LD (HL) C"
}

type LD_HLPtr_D struct {
	operand1 string
	operand2 string
}

func (o *LD_HLPtr_D) Execute(v *vm.VM) {
}

func (o *LD_HLPtr_D) String() string {
	return "LD (HL) D"
}

type LD_HLPtr_E struct {
	operand1 string
	operand2 string
}

func (o *LD_HLPtr_E) Execute(v *vm.VM) {
}

func (o *LD_HLPtr_E) String() string {
	return "LD (HL) E"
}

type LD_HLPtr_H struct {
	operand1 string
	operand2 string
}

func (o *LD_HLPtr_H) Execute(v *vm.VM) {
}

func (o *LD_HLPtr_H) String() string {
	return "LD (HL) H"
}

type LD_HLPtr_L struct {
	operand1 string
	operand2 string
}

func (o *LD_HLPtr_L) Execute(v *vm.VM) {
}

func (o *LD_HLPtr_L) String() string {
	return "LD (HL) L"
}

type HALT_ struct {
	operand1 string
	operand2 string
}

func (o *HALT_) Execute(v *vm.VM) {
}

func (o *HALT_) String() string {
	return "HALT"
}

type LD_HLPtr_A struct {
	operand1 string
	operand2 string
}

func (o *LD_HLPtr_A) Execute(v *vm.VM) {
}

func (o *LD_HLPtr_A) String() string {
	return "LD (HL) A"
}

type LD_A_B struct {
	operand1 string
	operand2 string
}

func (o *LD_A_B) Execute(v *vm.VM) {
}

func (o *LD_A_B) String() string {
	return "LD A B"
}

type LD_A_C struct {
	operand1 string
	operand2 string
}

func (o *LD_A_C) Execute(v *vm.VM) {
}

func (o *LD_A_C) String() string {
	return "LD A C"
}

type LD_A_D struct {
	operand1 string
	operand2 string
}

func (o *LD_A_D) Execute(v *vm.VM) {
}

func (o *LD_A_D) String() string {
	return "LD A D"
}

type LD_A_E struct {
	operand1 string
	operand2 string
}

func (o *LD_A_E) Execute(v *vm.VM) {
}

func (o *LD_A_E) String() string {
	return "LD A E"
}

type LD_A_H struct {
	operand1 string
	operand2 string
}

func (o *LD_A_H) Execute(v *vm.VM) {
}

func (o *LD_A_H) String() string {
	return "LD A H"
}

type LD_A_L struct {
	operand1 string
	operand2 string
}

func (o *LD_A_L) Execute(v *vm.VM) {
}

func (o *LD_A_L) String() string {
	return "LD A L"
}

type LD_A_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *LD_A_HLPtr) Execute(v *vm.VM) {
}

func (o *LD_A_HLPtr) String() string {
	return "LD A (HL)"
}

type LD_A_A struct {
	operand1 string
	operand2 string
}

func (o *LD_A_A) Execute(v *vm.VM) {
}

func (o *LD_A_A) String() string {
	return "LD A A"
}

type LD_a16Deref_SP struct {
	operand1 string
	operand2 string
}

func (o *LD_a16Deref_SP) Execute(v *vm.VM) {
}

func (o *LD_a16Deref_SP) String() string {
	return "LD (" + o.operand1 + ") SP"
}

type ADD_A_B struct {
	operand1 string
	operand2 string
}

func (o *ADD_A_B) Execute(v *vm.VM) {
}

func (o *ADD_A_B) String() string {
	return "ADD A B"
}

type ADD_A_C struct {
	operand1 string
	operand2 string
}

func (o *ADD_A_C) Execute(v *vm.VM) {
}

func (o *ADD_A_C) String() string {
	return "ADD A C"
}

type ADD_A_D struct {
	operand1 string
	operand2 string
}

func (o *ADD_A_D) Execute(v *vm.VM) {
}

func (o *ADD_A_D) String() string {
	return "ADD A D"
}

type ADD_A_E struct {
	operand1 string
	operand2 string
}

func (o *ADD_A_E) Execute(v *vm.VM) {
}

func (o *ADD_A_E) String() string {
	return "ADD A E"
}

type ADD_A_H struct {
	operand1 string
	operand2 string
}

func (o *ADD_A_H) Execute(v *vm.VM) {
}

func (o *ADD_A_H) String() string {
	return "ADD A H"
}

type ADD_A_L struct {
	operand1 string
	operand2 string
}

func (o *ADD_A_L) Execute(v *vm.VM) {
}

func (o *ADD_A_L) String() string {
	return "ADD A L"
}

type ADD_A_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *ADD_A_HLPtr) Execute(v *vm.VM) {
}

func (o *ADD_A_HLPtr) String() string {
	return "ADD A (HL)"
}

type ADD_A_A struct {
	operand1 string
	operand2 string
}

func (o *ADD_A_A) Execute(v *vm.VM) {
}

func (o *ADD_A_A) String() string {
	return "ADD A A"
}

type ADC_A_B struct {
	operand1 string
	operand2 string
}

func (o *ADC_A_B) Execute(v *vm.VM) {
}

func (o *ADC_A_B) String() string {
	return "ADC A B"
}

type ADC_A_C struct {
	operand1 string
	operand2 string
}

func (o *ADC_A_C) Execute(v *vm.VM) {
}

func (o *ADC_A_C) String() string {
	return "ADC A C"
}

type ADC_A_D struct {
	operand1 string
	operand2 string
}

func (o *ADC_A_D) Execute(v *vm.VM) {
}

func (o *ADC_A_D) String() string {
	return "ADC A D"
}

type ADC_A_E struct {
	operand1 string
	operand2 string
}

func (o *ADC_A_E) Execute(v *vm.VM) {
}

func (o *ADC_A_E) String() string {
	return "ADC A E"
}

type ADC_A_H struct {
	operand1 string
	operand2 string
}

func (o *ADC_A_H) Execute(v *vm.VM) {
}

func (o *ADC_A_H) String() string {
	return "ADC A H"
}

type ADC_A_L struct {
	operand1 string
	operand2 string
}

func (o *ADC_A_L) Execute(v *vm.VM) {
}

func (o *ADC_A_L) String() string {
	return "ADC A L"
}

type ADC_A_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *ADC_A_HLPtr) Execute(v *vm.VM) {
}

func (o *ADC_A_HLPtr) String() string {
	return "ADC A (HL)"
}

type ADC_A_A struct {
	operand1 string
	operand2 string
}

func (o *ADC_A_A) Execute(v *vm.VM) {
}

func (o *ADC_A_A) String() string {
	return "ADC A A"
}

type ADD_HL_BC struct {
	operand1 string
	operand2 string
}

func (o *ADD_HL_BC) Execute(v *vm.VM) {
}

func (o *ADD_HL_BC) String() string {
	return "ADD HL BC"
}

type SUB_B struct {
	operand1 string
	operand2 string
}

func (o *SUB_B) Execute(v *vm.VM) {
}

func (o *SUB_B) String() string {
	return "SUB B"
}

type SUB_C struct {
	operand1 string
	operand2 string
}

func (o *SUB_C) Execute(v *vm.VM) {
}

func (o *SUB_C) String() string {
	return "SUB C"
}

type SUB_D struct {
	operand1 string
	operand2 string
}

func (o *SUB_D) Execute(v *vm.VM) {
}

func (o *SUB_D) String() string {
	return "SUB D"
}

type SUB_E struct {
	operand1 string
	operand2 string
}

func (o *SUB_E) Execute(v *vm.VM) {
}

func (o *SUB_E) String() string {
	return "SUB E"
}

type SUB_H struct {
	operand1 string
	operand2 string
}

func (o *SUB_H) Execute(v *vm.VM) {
}

func (o *SUB_H) String() string {
	return "SUB H"
}

type SUB_L struct {
	operand1 string
	operand2 string
}

func (o *SUB_L) Execute(v *vm.VM) {
}

func (o *SUB_L) String() string {
	return "SUB L"
}

type SUB_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *SUB_HLPtr) Execute(v *vm.VM) {
}

func (o *SUB_HLPtr) String() string {
	return "SUB (HL)"
}

type SUB_A struct {
	operand1 string
	operand2 string
}

func (o *SUB_A) Execute(v *vm.VM) {
}

func (o *SUB_A) String() string {
	return "SUB A"
}

type SBC_A_B struct {
	operand1 string
	operand2 string
}

func (o *SBC_A_B) Execute(v *vm.VM) {
}

func (o *SBC_A_B) String() string {
	return "SBC A B"
}

type SBC_A_C struct {
	operand1 string
	operand2 string
}

func (o *SBC_A_C) Execute(v *vm.VM) {
}

func (o *SBC_A_C) String() string {
	return "SBC A C"
}

type SBC_A_D struct {
	operand1 string
	operand2 string
}

func (o *SBC_A_D) Execute(v *vm.VM) {
}

func (o *SBC_A_D) String() string {
	return "SBC A D"
}

type SBC_A_E struct {
	operand1 string
	operand2 string
}

func (o *SBC_A_E) Execute(v *vm.VM) {
}

func (o *SBC_A_E) String() string {
	return "SBC A E"
}

type SBC_A_H struct {
	operand1 string
	operand2 string
}

func (o *SBC_A_H) Execute(v *vm.VM) {
}

func (o *SBC_A_H) String() string {
	return "SBC A H"
}

type SBC_A_L struct {
	operand1 string
	operand2 string
}

func (o *SBC_A_L) Execute(v *vm.VM) {
}

func (o *SBC_A_L) String() string {
	return "SBC A L"
}

type SBC_A_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *SBC_A_HLPtr) Execute(v *vm.VM) {
}

func (o *SBC_A_HLPtr) String() string {
	return "SBC A (HL)"
}

type SBC_A_A struct {
	operand1 string
	operand2 string
}

func (o *SBC_A_A) Execute(v *vm.VM) {
}

func (o *SBC_A_A) String() string {
	return "SBC A A"
}

type LD_A_BCDeref struct {
	operand1 string
	operand2 string
}

func (o *LD_A_BCDeref) Execute(v *vm.VM) {
}

func (o *LD_A_BCDeref) String() string {
	return "LD A (BC)"
}

type AND_B struct {
	operand1 string
	operand2 string
}

func (o *AND_B) Execute(v *vm.VM) {
}

func (o *AND_B) String() string {
	return "AND B"
}

type AND_C struct {
	operand1 string
	operand2 string
}

func (o *AND_C) Execute(v *vm.VM) {
}

func (o *AND_C) String() string {
	return "AND C"
}

type AND_D struct {
	operand1 string
	operand2 string
}

func (o *AND_D) Execute(v *vm.VM) {
}

func (o *AND_D) String() string {
	return "AND D"
}

type AND_E struct {
	operand1 string
	operand2 string
}

func (o *AND_E) Execute(v *vm.VM) {
}

func (o *AND_E) String() string {
	return "AND E"
}

type AND_H struct {
	operand1 string
	operand2 string
}

func (o *AND_H) Execute(v *vm.VM) {
}

func (o *AND_H) String() string {
	return "AND H"
}

type AND_L struct {
	operand1 string
	operand2 string
}

func (o *AND_L) Execute(v *vm.VM) {
}

func (o *AND_L) String() string {
	return "AND L"
}

type AND_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *AND_HLPtr) Execute(v *vm.VM) {
}

func (o *AND_HLPtr) String() string {
	return "AND (HL)"
}

type AND_A struct {
	operand1 string
	operand2 string
}

func (o *AND_A) Execute(v *vm.VM) {
}

func (o *AND_A) String() string {
	return "AND A"
}

type XOR_B struct {
	operand1 string
	operand2 string
}

func (o *XOR_B) Execute(v *vm.VM) {
}

func (o *XOR_B) String() string {
	return "XOR B"
}

type XOR_C struct {
	operand1 string
	operand2 string
}

func (o *XOR_C) Execute(v *vm.VM) {
}

func (o *XOR_C) String() string {
	return "XOR C"
}

type XOR_D struct {
	operand1 string
	operand2 string
}

func (o *XOR_D) Execute(v *vm.VM) {
}

func (o *XOR_D) String() string {
	return "XOR D"
}

type XOR_E struct {
	operand1 string
	operand2 string
}

func (o *XOR_E) Execute(v *vm.VM) {
}

func (o *XOR_E) String() string {
	return "XOR E"
}

type XOR_H struct {
	operand1 string
	operand2 string
}

func (o *XOR_H) Execute(v *vm.VM) {
}

func (o *XOR_H) String() string {
	return "XOR H"
}

type XOR_L struct {
	operand1 string
	operand2 string
}

func (o *XOR_L) Execute(v *vm.VM) {
}

func (o *XOR_L) String() string {
	return "XOR L"
}

type XOR_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *XOR_HLPtr) Execute(v *vm.VM) {
}

func (o *XOR_HLPtr) String() string {
	return "XOR (HL)"
}

type XOR_A struct {
	operand1 string
	operand2 string
}

func (o *XOR_A) Execute(v *vm.VM) {
}

func (o *XOR_A) String() string {
	return "XOR A"
}

type DEC_BC struct {
	operand1 string
	operand2 string
}

func (o *DEC_BC) Execute(v *vm.VM) {
}

func (o *DEC_BC) String() string {
	return "DEC BC"
}

type OR_B struct {
	operand1 string
	operand2 string
}

func (o *OR_B) Execute(v *vm.VM) {
}

func (o *OR_B) String() string {
	return "OR B"
}

type OR_C struct {
	operand1 string
	operand2 string
}

func (o *OR_C) Execute(v *vm.VM) {
}

func (o *OR_C) String() string {
	return "OR C"
}

type OR_D struct {
	operand1 string
	operand2 string
}

func (o *OR_D) Execute(v *vm.VM) {
}

func (o *OR_D) String() string {
	return "OR D"
}

type OR_E struct {
	operand1 string
	operand2 string
}

func (o *OR_E) Execute(v *vm.VM) {
}

func (o *OR_E) String() string {
	return "OR E"
}

type OR_H struct {
	operand1 string
	operand2 string
}

func (o *OR_H) Execute(v *vm.VM) {
}

func (o *OR_H) String() string {
	return "OR H"
}

type OR_L struct {
	operand1 string
	operand2 string
}

func (o *OR_L) Execute(v *vm.VM) {
}

func (o *OR_L) String() string {
	return "OR L"
}

type OR_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *OR_HLPtr) Execute(v *vm.VM) {
}

func (o *OR_HLPtr) String() string {
	return "OR (HL)"
}

type OR_A struct {
	operand1 string
	operand2 string
}

func (o *OR_A) Execute(v *vm.VM) {
}

func (o *OR_A) String() string {
	return "OR A"
}

type CP_B struct {
	operand1 string
	operand2 string
}

func (o *CP_B) Execute(v *vm.VM) {
}

func (o *CP_B) String() string {
	return "CP B"
}

type CP_C struct {
	operand1 string
	operand2 string
}

func (o *CP_C) Execute(v *vm.VM) {
}

func (o *CP_C) String() string {
	return "CP C"
}

type CP_D struct {
	operand1 string
	operand2 string
}

func (o *CP_D) Execute(v *vm.VM) {
}

func (o *CP_D) String() string {
	return "CP D"
}

type CP_E struct {
	operand1 string
	operand2 string
}

func (o *CP_E) Execute(v *vm.VM) {
}

func (o *CP_E) String() string {
	return "CP E"
}

type CP_H struct {
	operand1 string
	operand2 string
}

func (o *CP_H) Execute(v *vm.VM) {
}

func (o *CP_H) String() string {
	return "CP H"
}

type CP_L struct {
	operand1 string
	operand2 string
}

func (o *CP_L) Execute(v *vm.VM) {
}

func (o *CP_L) String() string {
	return "CP L"
}

type CP_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *CP_HLPtr) Execute(v *vm.VM) {
}

func (o *CP_HLPtr) String() string {
	return "CP (HL)"
}

type CP_A struct {
	operand1 string
	operand2 string
}

func (o *CP_A) Execute(v *vm.VM) {
}

func (o *CP_A) String() string {
	return "CP A"
}

type INC_C struct {
	operand1 string
	operand2 string
}

func (o *INC_C) Execute(v *vm.VM) {
}

func (o *INC_C) String() string {
	return "INC C"
}

type RET_NZ struct {
	operand1 string
	operand2 string
}

func (o *RET_NZ) Execute(v *vm.VM) {
}

func (o *RET_NZ) String() string {
	return "RET NZ"
}

type POP_BC struct {
	operand1 string
	operand2 string
}

func (o *POP_BC) Execute(v *vm.VM) {
}

func (o *POP_BC) String() string {
	return "POP BC"
}

type JP_NZ_a16 struct {
	operand1 string
	operand2 string
}

func (o *JP_NZ_a16) Execute(v *vm.VM) {
}

func (o *JP_NZ_a16) String() string {
	return "JP NZ o.operand1"
}

type JP_a16 struct {
	operand1 string
	operand2 string
}

func (o *JP_a16) Execute(v *vm.VM) {
}

func (o *JP_a16) String() string {
	return "JP o.operand1"
}

type CALL_NZ_a16 struct {
	operand1 string
	operand2 string
}

func (o *CALL_NZ_a16) Execute(v *vm.VM) {
}

func (o *CALL_NZ_a16) String() string {
	return "CALL NZ o.operand1"
}

type PUSH_BC struct {
	operand1 string
	operand2 string
}

func (o *PUSH_BC) Execute(v *vm.VM) {
}

func (o *PUSH_BC) String() string {
	return "PUSH BC"
}

type ADD_A_d8 struct {
	operand1 string
	operand2 string
}

func (o *ADD_A_d8) Execute(v *vm.VM) {
}

func (o *ADD_A_d8) String() string {
	return "ADD A d8"
}

type RST_00H struct {
	operand1 string
	operand2 string
}

func (o *RST_00H) Execute(v *vm.VM) {
}

func (o *RST_00H) String() string {
	return "RST 00H"
}

type RET_Z struct {
	operand1 string
	operand2 string
}

func (o *RET_Z) Execute(v *vm.VM) {
}

func (o *RET_Z) String() string {
	return "RET Z"
}

type RET_ struct {
	operand1 string
	operand2 string
}

func (o *RET_) Execute(v *vm.VM) {
}

func (o *RET_) String() string {
	return "RET"
}

type JP_Z_a16 struct {
	operand1 string
	operand2 string
}

func (o *JP_Z_a16) Execute(v *vm.VM) {
}

func (o *JP_Z_a16) String() string {
	return "JP Z o.operand1"
}

type PREFIX_CB struct {
	operand1 string
	operand2 string
}

func (o *PREFIX_CB) Execute(v *vm.VM) {
}

func (o *PREFIX_CB) String() string {
	return "PREFIX CB"
}

type CALL_Z_a16 struct {
	operand1 string
	operand2 string
}

func (o *CALL_Z_a16) Execute(v *vm.VM) {
}

func (o *CALL_Z_a16) String() string {
	return "CALL Z o.operand1"
}

type CALL_a16 struct {
	operand1 string
	operand2 string
}

func (o *CALL_a16) Execute(v *vm.VM) {
}

func (o *CALL_a16) String() string {
	return "CALL o.operand1"
}

type ADC_A_d8 struct {
	operand1 string
	operand2 string
}

func (o *ADC_A_d8) Execute(v *vm.VM) {
}

func (o *ADC_A_d8) String() string {
	return "ADC A d8"
}

type RST_08H struct {
	operand1 string
	operand2 string
}

func (o *RST_08H) Execute(v *vm.VM) {
}

func (o *RST_08H) String() string {
	return "RST 08H"
}

type DEC_C struct {
	operand1 string
	operand2 string
}

func (o *DEC_C) Execute(v *vm.VM) {
}

func (o *DEC_C) String() string {
	return "DEC C"
}

type RET_NC struct {
	operand1 string
	operand2 string
}

func (o *RET_NC) Execute(v *vm.VM) {
}

func (o *RET_NC) String() string {
	return "RET NC"
}

type POP_DE struct {
	operand1 string
	operand2 string
}

func (o *POP_DE) Execute(v *vm.VM) {
}

func (o *POP_DE) String() string {
	return "POP DE"
}

type JP_NC_a16 struct {
	operand1 string
	operand2 string
}

func (o *JP_NC_a16) Execute(v *vm.VM) {
}

func (o *JP_NC_a16) String() string {
	return "JP NC o.operand1"
}

type CALL_NC_a16 struct {
	operand1 string
	operand2 string
}

func (o *CALL_NC_a16) Execute(v *vm.VM) {
}

func (o *CALL_NC_a16) String() string {
	return "CALL NC o.operand1"
}

type PUSH_DE struct {
	operand1 string
	operand2 string
}

func (o *PUSH_DE) Execute(v *vm.VM) {
}

func (o *PUSH_DE) String() string {
	return "PUSH DE"
}

type SUB_d8 struct {
	operand1 string
	operand2 string
}

func (o *SUB_d8) Execute(v *vm.VM) {
}

func (o *SUB_d8) String() string {
	return "SUB d8"
}

type RST_10H struct {
	operand1 string
	operand2 string
}

func (o *RST_10H) Execute(v *vm.VM) {
}

func (o *RST_10H) String() string {
	return "RST 10H"
}

type RET_C struct {
	operand1 string
	operand2 string
}

func (o *RET_C) Execute(v *vm.VM) {
}

func (o *RET_C) String() string {
	return "RET C"
}

type RETI_ struct {
	operand1 string
	operand2 string
}

func (o *RETI_) Execute(v *vm.VM) {
}

func (o *RETI_) String() string {
	return "RETI"
}

type JP_C_a16 struct {
	operand1 string
	operand2 string
}

func (o *JP_C_a16) Execute(v *vm.VM) {
}

func (o *JP_C_a16) String() string {
	return "JP C o.operand1"
}

type CALL_C_a16 struct {
	operand1 string
	operand2 string
}

func (o *CALL_C_a16) Execute(v *vm.VM) {
}

func (o *CALL_C_a16) String() string {
	return "CALL C o.operand1"
}

type SBC_A_d8 struct {
	operand1 string
	operand2 string
}

func (o *SBC_A_d8) Execute(v *vm.VM) {
}

func (o *SBC_A_d8) String() string {
	return "SBC A d8"
}

type RST_18H struct {
	operand1 string
	operand2 string
}

func (o *RST_18H) Execute(v *vm.VM) {
}

func (o *RST_18H) String() string {
	return "RST 18H"
}

type LD_C_d8 struct {
	operand1 string
	operand2 string
}

func (o *LD_C_d8) Execute(v *vm.VM) {
}

func (o *LD_C_d8) String() string {
	return "LD C d8"
}

type LDH_a8Deref_A struct {
	operand1 string
	operand2 string
}

func (o *LDH_a8Deref_A) Execute(v *vm.VM) {
}

func (o *LDH_a8Deref_A) String() string {
	return "LDH (a8) A"
}

type POP_HL struct {
	operand1 string
	operand2 string
}

func (o *POP_HL) Execute(v *vm.VM) {
}

func (o *POP_HL) String() string {
	return "POP HL"
}

type LD_CDeref_A struct {
	operand1 string
	operand2 string
}

func (o *LD_CDeref_A) Execute(v *vm.VM) {
}

func (o *LD_CDeref_A) String() string {
	return "LD (C) A"
}

type PUSH_HL struct {
	operand1 string
	operand2 string
}

func (o *PUSH_HL) Execute(v *vm.VM) {
}

func (o *PUSH_HL) String() string {
	return "PUSH HL"
}

type AND_d8 struct {
	operand1 string
	operand2 string
}

func (o *AND_d8) Execute(v *vm.VM) {
}

func (o *AND_d8) String() string {
	return "AND d8"
}

type RST_20H struct {
	operand1 string
	operand2 string
}

func (o *RST_20H) Execute(v *vm.VM) {
}

func (o *RST_20H) String() string {
	return "RST 20H"
}

type ADD_SP_r8 struct {
	operand1 string
	operand2 string
}

func (o *ADD_SP_r8) Execute(v *vm.VM) {
}

func (o *ADD_SP_r8) String() string {
	return "ADD SP r8"
}

type JP_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *JP_HLPtr) Execute(v *vm.VM) {
}

func (o *JP_HLPtr) String() string {
	return "JP (HL)"
}

type LD_a16Deref_A struct {
	operand1 string
	operand2 string
}

func (o *LD_a16Deref_A) Execute(v *vm.VM) {
}

func (o *LD_a16Deref_A) String() string {
	return "LD (" + o.operand1 + ") A"
}

type XOR_d8 struct {
	operand1 string
	operand2 string
}

func (o *XOR_d8) Execute(v *vm.VM) {
}

func (o *XOR_d8) String() string {
	return "XOR d8"
}

type RST_28H struct {
	operand1 string
	operand2 string
}

func (o *RST_28H) Execute(v *vm.VM) {
}

func (o *RST_28H) String() string {
	return "RST 28H"
}

type RRCA_ struct {
	operand1 string
	operand2 string
}

func (o *RRCA_) Execute(v *vm.VM) {
}

func (o *RRCA_) String() string {
	return "RRCA"
}

type LDH_A_a8Deref struct {
	operand1 string
	operand2 string
}

func (o *LDH_A_a8Deref) Execute(v *vm.VM) {
}

func (o *LDH_A_a8Deref) String() string {
	return "LDH A (a8)"
}

type POP_AF struct {
	operand1 string
	operand2 string
}

func (o *POP_AF) Execute(v *vm.VM) {
}

func (o *POP_AF) String() string {
	return "POP AF"
}

type LD_A_CDeref struct {
	operand1 string
	operand2 string
}

func (o *LD_A_CDeref) Execute(v *vm.VM) {
}

func (o *LD_A_CDeref) String() string {
	return "LD A (C)"
}

type DI_ struct {
	operand1 string
	operand2 string
}

func (o *DI_) Execute(v *vm.VM) {
}

func (o *DI_) String() string {
	return "DI"
}

type PUSH_AF struct {
	operand1 string
	operand2 string
}

func (o *PUSH_AF) Execute(v *vm.VM) {
}

func (o *PUSH_AF) String() string {
	return "PUSH AF"
}

type OR_d8 struct {
	operand1 string
	operand2 string
}

func (o *OR_d8) Execute(v *vm.VM) {
}

func (o *OR_d8) String() string {
	return "OR d8"
}

type RST_30H struct {
	operand1 string
	operand2 string
}

func (o *RST_30H) Execute(v *vm.VM) {
}

func (o *RST_30H) String() string {
	return "RST 30H"
}

type LD_HL_SP_plus_r8 struct {
	operand1 string
	operand2 string
}

func (o *LD_HL_SP_plus_r8) Execute(v *vm.VM) {
}

func (o *LD_HL_SP_plus_r8) String() string {
	return "LD HL SP+r8"
}

type LD_SP_HL struct {
	operand1 string
	operand2 string
}

func (o *LD_SP_HL) Execute(v *vm.VM) {
}

func (o *LD_SP_HL) String() string {
	return "LD SP HL"
}

type LD_A_a16Deref struct {
	operand1 string
	operand2 string
}

func (o *LD_A_a16Deref) Execute(v *vm.VM) {
}

func (o *LD_A_a16Deref) String() string {
	return "LD A (" + o.operand1 + ")"
}

type EI_ struct {
	operand1 string
	operand2 string
}

func (o *EI_) Execute(v *vm.VM) {
}

func (o *EI_) String() string {
	return "EI"
}

type CP_d8 struct {
	operand1 string
	operand2 string
}

func (o *CP_d8) Execute(v *vm.VM) {
}

func (o *CP_d8) String() string {
	return "CP d8"
}

type RST_38H struct {
	operand1 string
	operand2 string
}

func (o *RST_38H) Execute(v *vm.VM) {
}

func (o *RST_38H) String() string {
	return "RST 38H"
}

type RLC_B struct {
	operand1 string
	operand2 string
}

func (o *RLC_B) Execute(v *vm.VM) {
}

func (o *RLC_B) String() string {
	return "RLC B"
}

type RLC_C struct {
	operand1 string
	operand2 string
}

func (o *RLC_C) Execute(v *vm.VM) {
}

func (o *RLC_C) String() string {
	return "RLC C"
}

type RL_B struct {
	operand1 string
	operand2 string
}

func (o *RL_B) Execute(v *vm.VM) {
}

func (o *RL_B) String() string {
	return "RL B"
}

type RL_C struct {
	operand1 string
	operand2 string
}

func (o *RL_C) Execute(v *vm.VM) {
}

func (o *RL_C) String() string {
	return "RL C"
}

type RL_D struct {
	operand1 string
	operand2 string
}

func (o *RL_D) Execute(v *vm.VM) {
}

func (o *RL_D) String() string {
	return "RL D"
}

type RL_E struct {
	operand1 string
	operand2 string
}

func (o *RL_E) Execute(v *vm.VM) {
}

func (o *RL_E) String() string {
	return "RL E"
}

type RL_H struct {
	operand1 string
	operand2 string
}

func (o *RL_H) Execute(v *vm.VM) {
}

func (o *RL_H) String() string {
	return "RL H"
}

type RL_L struct {
	operand1 string
	operand2 string
}

func (o *RL_L) Execute(v *vm.VM) {
}

func (o *RL_L) String() string {
	return "RL L"
}

type RL_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *RL_HLPtr) Execute(v *vm.VM) {
}

func (o *RL_HLPtr) String() string {
	return "RL (HL)"
}

type RL_A struct {
	operand1 string
	operand2 string
}

func (o *RL_A) Execute(v *vm.VM) {
}

func (o *RL_A) String() string {
	return "RL A"
}

type RR_B struct {
	operand1 string
	operand2 string
}

func (o *RR_B) Execute(v *vm.VM) {
}

func (o *RR_B) String() string {
	return "RR B"
}

type RR_C struct {
	operand1 string
	operand2 string
}

func (o *RR_C) Execute(v *vm.VM) {
}

func (o *RR_C) String() string {
	return "RR C"
}

type RR_D struct {
	operand1 string
	operand2 string
}

func (o *RR_D) Execute(v *vm.VM) {
}

func (o *RR_D) String() string {
	return "RR D"
}

type RR_E struct {
	operand1 string
	operand2 string
}

func (o *RR_E) Execute(v *vm.VM) {
}

func (o *RR_E) String() string {
	return "RR E"
}

type RR_H struct {
	operand1 string
	operand2 string
}

func (o *RR_H) Execute(v *vm.VM) {
}

func (o *RR_H) String() string {
	return "RR H"
}

type RR_L struct {
	operand1 string
	operand2 string
}

func (o *RR_L) Execute(v *vm.VM) {
}

func (o *RR_L) String() string {
	return "RR L"
}

type RR_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *RR_HLPtr) Execute(v *vm.VM) {
}

func (o *RR_HLPtr) String() string {
	return "RR (HL)"
}

type RR_A struct {
	operand1 string
	operand2 string
}

func (o *RR_A) Execute(v *vm.VM) {
}

func (o *RR_A) String() string {
	return "RR A"
}

type RLC_D struct {
	operand1 string
	operand2 string
}

func (o *RLC_D) Execute(v *vm.VM) {
}

func (o *RLC_D) String() string {
	return "RLC D"
}

type SLA_B struct {
	operand1 string
	operand2 string
}

func (o *SLA_B) Execute(v *vm.VM) {
}

func (o *SLA_B) String() string {
	return "SLA B"
}

type SLA_C struct {
	operand1 string
	operand2 string
}

func (o *SLA_C) Execute(v *vm.VM) {
}

func (o *SLA_C) String() string {
	return "SLA C"
}

type SLA_D struct {
	operand1 string
	operand2 string
}

func (o *SLA_D) Execute(v *vm.VM) {
}

func (o *SLA_D) String() string {
	return "SLA D"
}

type SLA_E struct {
	operand1 string
	operand2 string
}

func (o *SLA_E) Execute(v *vm.VM) {
}

func (o *SLA_E) String() string {
	return "SLA E"
}

type SLA_H struct {
	operand1 string
	operand2 string
}

func (o *SLA_H) Execute(v *vm.VM) {
}

func (o *SLA_H) String() string {
	return "SLA H"
}

type SLA_L struct {
	operand1 string
	operand2 string
}

func (o *SLA_L) Execute(v *vm.VM) {
}

func (o *SLA_L) String() string {
	return "SLA L"
}

type SLA_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *SLA_HLPtr) Execute(v *vm.VM) {
}

func (o *SLA_HLPtr) String() string {
	return "SLA (HL)"
}

type SLA_A struct {
	operand1 string
	operand2 string
}

func (o *SLA_A) Execute(v *vm.VM) {
}

func (o *SLA_A) String() string {
	return "SLA A"
}

type SRA_B struct {
	operand1 string
	operand2 string
}

func (o *SRA_B) Execute(v *vm.VM) {
}

func (o *SRA_B) String() string {
	return "SRA B"
}

type SRA_C struct {
	operand1 string
	operand2 string
}

func (o *SRA_C) Execute(v *vm.VM) {
}

func (o *SRA_C) String() string {
	return "SRA C"
}

type SRA_D struct {
	operand1 string
	operand2 string
}

func (o *SRA_D) Execute(v *vm.VM) {
}

func (o *SRA_D) String() string {
	return "SRA D"
}

type SRA_E struct {
	operand1 string
	operand2 string
}

func (o *SRA_E) Execute(v *vm.VM) {
}

func (o *SRA_E) String() string {
	return "SRA E"
}

type SRA_H struct {
	operand1 string
	operand2 string
}

func (o *SRA_H) Execute(v *vm.VM) {
}

func (o *SRA_H) String() string {
	return "SRA H"
}

type SRA_L struct {
	operand1 string
	operand2 string
}

func (o *SRA_L) Execute(v *vm.VM) {
}

func (o *SRA_L) String() string {
	return "SRA L"
}

type SRA_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *SRA_HLPtr) Execute(v *vm.VM) {
}

func (o *SRA_HLPtr) String() string {
	return "SRA (HL)"
}

type SRA_A struct {
	operand1 string
	operand2 string
}

func (o *SRA_A) Execute(v *vm.VM) {
}

func (o *SRA_A) String() string {
	return "SRA A"
}

type RLC_E struct {
	operand1 string
	operand2 string
}

func (o *RLC_E) Execute(v *vm.VM) {
}

func (o *RLC_E) String() string {
	return "RLC E"
}

type SWAP_B struct {
	operand1 string
	operand2 string
}

func (o *SWAP_B) Execute(v *vm.VM) {
}

func (o *SWAP_B) String() string {
	return "SWAP B"
}

type SWAP_C struct {
	operand1 string
	operand2 string
}

func (o *SWAP_C) Execute(v *vm.VM) {
}

func (o *SWAP_C) String() string {
	return "SWAP C"
}

type SWAP_D struct {
	operand1 string
	operand2 string
}

func (o *SWAP_D) Execute(v *vm.VM) {
}

func (o *SWAP_D) String() string {
	return "SWAP D"
}

type SWAP_E struct {
	operand1 string
	operand2 string
}

func (o *SWAP_E) Execute(v *vm.VM) {
}

func (o *SWAP_E) String() string {
	return "SWAP E"
}

type SWAP_H struct {
	operand1 string
	operand2 string
}

func (o *SWAP_H) Execute(v *vm.VM) {
}

func (o *SWAP_H) String() string {
	return "SWAP H"
}

type SWAP_L struct {
	operand1 string
	operand2 string
}

func (o *SWAP_L) Execute(v *vm.VM) {
}

func (o *SWAP_L) String() string {
	return "SWAP L"
}

type SWAP_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *SWAP_HLPtr) Execute(v *vm.VM) {
}

func (o *SWAP_HLPtr) String() string {
	return "SWAP (HL)"
}

type SWAP_A struct {
	operand1 string
	operand2 string
}

func (o *SWAP_A) Execute(v *vm.VM) {
}

func (o *SWAP_A) String() string {
	return "SWAP A"
}

type SRL_B struct {
	operand1 string
	operand2 string
}

func (o *SRL_B) Execute(v *vm.VM) {
}

func (o *SRL_B) String() string {
	return "SRL B"
}

type SRL_C struct {
	operand1 string
	operand2 string
}

func (o *SRL_C) Execute(v *vm.VM) {
}

func (o *SRL_C) String() string {
	return "SRL C"
}

type SRL_D struct {
	operand1 string
	operand2 string
}

func (o *SRL_D) Execute(v *vm.VM) {
}

func (o *SRL_D) String() string {
	return "SRL D"
}

type SRL_E struct {
	operand1 string
	operand2 string
}

func (o *SRL_E) Execute(v *vm.VM) {
}

func (o *SRL_E) String() string {
	return "SRL E"
}

type SRL_H struct {
	operand1 string
	operand2 string
}

func (o *SRL_H) Execute(v *vm.VM) {
}

func (o *SRL_H) String() string {
	return "SRL H"
}

type SRL_L struct {
	operand1 string
	operand2 string
}

func (o *SRL_L) Execute(v *vm.VM) {
}

func (o *SRL_L) String() string {
	return "SRL L"
}

type SRL_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *SRL_HLPtr) Execute(v *vm.VM) {
}

func (o *SRL_HLPtr) String() string {
	return "SRL (HL)"
}

type SRL_A struct {
	operand1 string
	operand2 string
}

func (o *SRL_A) Execute(v *vm.VM) {
}

func (o *SRL_A) String() string {
	return "SRL A"
}

type RLC_H struct {
	operand1 string
	operand2 string
}

func (o *RLC_H) Execute(v *vm.VM) {
}

func (o *RLC_H) String() string {
	return "RLC H"
}

type BIT_0_B struct {
	operand1 string
	operand2 string
}

func (o *BIT_0_B) Execute(v *vm.VM) {
}

func (o *BIT_0_B) String() string {
	return "BIT 0 B"
}

type BIT_0_C struct {
	operand1 string
	operand2 string
}

func (o *BIT_0_C) Execute(v *vm.VM) {
}

func (o *BIT_0_C) String() string {
	return "BIT 0 C"
}

type BIT_0_D struct {
	operand1 string
	operand2 string
}

func (o *BIT_0_D) Execute(v *vm.VM) {
}

func (o *BIT_0_D) String() string {
	return "BIT 0 D"
}

type BIT_0_E struct {
	operand1 string
	operand2 string
}

func (o *BIT_0_E) Execute(v *vm.VM) {
}

func (o *BIT_0_E) String() string {
	return "BIT 0 E"
}

type BIT_0_H struct {
	operand1 string
	operand2 string
}

func (o *BIT_0_H) Execute(v *vm.VM) {
}

func (o *BIT_0_H) String() string {
	return "BIT 0 H"
}

type BIT_0_L struct {
	operand1 string
	operand2 string
}

func (o *BIT_0_L) Execute(v *vm.VM) {
}

func (o *BIT_0_L) String() string {
	return "BIT 0 L"
}

type BIT_0_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *BIT_0_HLPtr) Execute(v *vm.VM) {
}

func (o *BIT_0_HLPtr) String() string {
	return "BIT 0 (HL)"
}

type BIT_0_A struct {
	operand1 string
	operand2 string
}

func (o *BIT_0_A) Execute(v *vm.VM) {
}

func (o *BIT_0_A) String() string {
	return "BIT 0 A"
}

type BIT_1_B struct {
	operand1 string
	operand2 string
}

func (o *BIT_1_B) Execute(v *vm.VM) {
}

func (o *BIT_1_B) String() string {
	return "BIT 1 B"
}

type BIT_1_C struct {
	operand1 string
	operand2 string
}

func (o *BIT_1_C) Execute(v *vm.VM) {
}

func (o *BIT_1_C) String() string {
	return "BIT 1 C"
}

type BIT_1_D struct {
	operand1 string
	operand2 string
}

func (o *BIT_1_D) Execute(v *vm.VM) {
}

func (o *BIT_1_D) String() string {
	return "BIT 1 D"
}

type BIT_1_E struct {
	operand1 string
	operand2 string
}

func (o *BIT_1_E) Execute(v *vm.VM) {
}

func (o *BIT_1_E) String() string {
	return "BIT 1 E"
}

type BIT_1_H struct {
	operand1 string
	operand2 string
}

func (o *BIT_1_H) Execute(v *vm.VM) {
}

func (o *BIT_1_H) String() string {
	return "BIT 1 H"
}

type BIT_1_L struct {
	operand1 string
	operand2 string
}

func (o *BIT_1_L) Execute(v *vm.VM) {
}

func (o *BIT_1_L) String() string {
	return "BIT 1 L"
}

type BIT_1_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *BIT_1_HLPtr) Execute(v *vm.VM) {
}

func (o *BIT_1_HLPtr) String() string {
	return "BIT 1 (HL)"
}

type BIT_1_A struct {
	operand1 string
	operand2 string
}

func (o *BIT_1_A) Execute(v *vm.VM) {
}

func (o *BIT_1_A) String() string {
	return "BIT 1 A"
}

type RLC_L struct {
	operand1 string
	operand2 string
}

func (o *RLC_L) Execute(v *vm.VM) {
}

func (o *RLC_L) String() string {
	return "RLC L"
}

type BIT_2_B struct {
	operand1 string
	operand2 string
}

func (o *BIT_2_B) Execute(v *vm.VM) {
}

func (o *BIT_2_B) String() string {
	return "BIT 2 B"
}

type BIT_2_C struct {
	operand1 string
	operand2 string
}

func (o *BIT_2_C) Execute(v *vm.VM) {
}

func (o *BIT_2_C) String() string {
	return "BIT 2 C"
}

type BIT_2_D struct {
	operand1 string
	operand2 string
}

func (o *BIT_2_D) Execute(v *vm.VM) {
}

func (o *BIT_2_D) String() string {
	return "BIT 2 D"
}

type BIT_2_E struct {
	operand1 string
	operand2 string
}

func (o *BIT_2_E) Execute(v *vm.VM) {
}

func (o *BIT_2_E) String() string {
	return "BIT 2 E"
}

type BIT_2_H struct {
	operand1 string
	operand2 string
}

func (o *BIT_2_H) Execute(v *vm.VM) {
}

func (o *BIT_2_H) String() string {
	return "BIT 2 H"
}

type BIT_2_L struct {
	operand1 string
	operand2 string
}

func (o *BIT_2_L) Execute(v *vm.VM) {
}

func (o *BIT_2_L) String() string {
	return "BIT 2 L"
}

type BIT_2_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *BIT_2_HLPtr) Execute(v *vm.VM) {
}

func (o *BIT_2_HLPtr) String() string {
	return "BIT 2 (HL)"
}

type BIT_2_A struct {
	operand1 string
	operand2 string
}

func (o *BIT_2_A) Execute(v *vm.VM) {
}

func (o *BIT_2_A) String() string {
	return "BIT 2 A"
}

type BIT_3_B struct {
	operand1 string
	operand2 string
}

func (o *BIT_3_B) Execute(v *vm.VM) {
}

func (o *BIT_3_B) String() string {
	return "BIT 3 B"
}

type BIT_3_C struct {
	operand1 string
	operand2 string
}

func (o *BIT_3_C) Execute(v *vm.VM) {
}

func (o *BIT_3_C) String() string {
	return "BIT 3 C"
}

type BIT_3_D struct {
	operand1 string
	operand2 string
}

func (o *BIT_3_D) Execute(v *vm.VM) {
}

func (o *BIT_3_D) String() string {
	return "BIT 3 D"
}

type BIT_3_E struct {
	operand1 string
	operand2 string
}

func (o *BIT_3_E) Execute(v *vm.VM) {
}

func (o *BIT_3_E) String() string {
	return "BIT 3 E"
}

type BIT_3_H struct {
	operand1 string
	operand2 string
}

func (o *BIT_3_H) Execute(v *vm.VM) {
}

func (o *BIT_3_H) String() string {
	return "BIT 3 H"
}

type BIT_3_L struct {
	operand1 string
	operand2 string
}

func (o *BIT_3_L) Execute(v *vm.VM) {
}

func (o *BIT_3_L) String() string {
	return "BIT 3 L"
}

type BIT_3_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *BIT_3_HLPtr) Execute(v *vm.VM) {
}

func (o *BIT_3_HLPtr) String() string {
	return "BIT 3 (HL)"
}

type BIT_3_A struct {
	operand1 string
	operand2 string
}

func (o *BIT_3_A) Execute(v *vm.VM) {
}

func (o *BIT_3_A) String() string {
	return "BIT 3 A"
}

type RLC_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *RLC_HLPtr) Execute(v *vm.VM) {
}

func (o *RLC_HLPtr) String() string {
	return "RLC (HL)"
}

type BIT_4_B struct {
	operand1 string
	operand2 string
}

func (o *BIT_4_B) Execute(v *vm.VM) {
}

func (o *BIT_4_B) String() string {
	return "BIT 4 B"
}

type BIT_4_C struct {
	operand1 string
	operand2 string
}

func (o *BIT_4_C) Execute(v *vm.VM) {
}

func (o *BIT_4_C) String() string {
	return "BIT 4 C"
}

type BIT_4_D struct {
	operand1 string
	operand2 string
}

func (o *BIT_4_D) Execute(v *vm.VM) {
}

func (o *BIT_4_D) String() string {
	return "BIT 4 D"
}

type BIT_4_E struct {
	operand1 string
	operand2 string
}

func (o *BIT_4_E) Execute(v *vm.VM) {
}

func (o *BIT_4_E) String() string {
	return "BIT 4 E"
}

type BIT_4_H struct {
	operand1 string
	operand2 string
}

func (o *BIT_4_H) Execute(v *vm.VM) {
}

func (o *BIT_4_H) String() string {
	return "BIT 4 H"
}

type BIT_4_L struct {
	operand1 string
	operand2 string
}

func (o *BIT_4_L) Execute(v *vm.VM) {
}

func (o *BIT_4_L) String() string {
	return "BIT 4 L"
}

type BIT_4_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *BIT_4_HLPtr) Execute(v *vm.VM) {
}

func (o *BIT_4_HLPtr) String() string {
	return "BIT 4 (HL)"
}

type BIT_4_A struct {
	operand1 string
	operand2 string
}

func (o *BIT_4_A) Execute(v *vm.VM) {
}

func (o *BIT_4_A) String() string {
	return "BIT 4 A"
}

type BIT_5_B struct {
	operand1 string
	operand2 string
}

func (o *BIT_5_B) Execute(v *vm.VM) {
}

func (o *BIT_5_B) String() string {
	return "BIT 5 B"
}

type BIT_5_C struct {
	operand1 string
	operand2 string
}

func (o *BIT_5_C) Execute(v *vm.VM) {
}

func (o *BIT_5_C) String() string {
	return "BIT 5 C"
}

type BIT_5_D struct {
	operand1 string
	operand2 string
}

func (o *BIT_5_D) Execute(v *vm.VM) {
}

func (o *BIT_5_D) String() string {
	return "BIT 5 D"
}

type BIT_5_E struct {
	operand1 string
	operand2 string
}

func (o *BIT_5_E) Execute(v *vm.VM) {
}

func (o *BIT_5_E) String() string {
	return "BIT 5 E"
}

type BIT_5_H struct {
	operand1 string
	operand2 string
}

func (o *BIT_5_H) Execute(v *vm.VM) {
}

func (o *BIT_5_H) String() string {
	return "BIT 5 H"
}

type BIT_5_L struct {
	operand1 string
	operand2 string
}

func (o *BIT_5_L) Execute(v *vm.VM) {
}

func (o *BIT_5_L) String() string {
	return "BIT 5 L"
}

type BIT_5_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *BIT_5_HLPtr) Execute(v *vm.VM) {
}

func (o *BIT_5_HLPtr) String() string {
	return "BIT 5 (HL)"
}

type BIT_5_A struct {
	operand1 string
	operand2 string
}

func (o *BIT_5_A) Execute(v *vm.VM) {
}

func (o *BIT_5_A) String() string {
	return "BIT 5 A"
}

type RLC_A struct {
	operand1 string
	operand2 string
}

func (o *RLC_A) Execute(v *vm.VM) {
}

func (o *RLC_A) String() string {
	return "RLC A"
}

type BIT_6_B struct {
	operand1 string
	operand2 string
}

func (o *BIT_6_B) Execute(v *vm.VM) {
}

func (o *BIT_6_B) String() string {
	return "BIT 6 B"
}

type BIT_6_C struct {
	operand1 string
	operand2 string
}

func (o *BIT_6_C) Execute(v *vm.VM) {
}

func (o *BIT_6_C) String() string {
	return "BIT 6 C"
}

type BIT_6_D struct {
	operand1 string
	operand2 string
}

func (o *BIT_6_D) Execute(v *vm.VM) {
}

func (o *BIT_6_D) String() string {
	return "BIT 6 D"
}

type BIT_6_E struct {
	operand1 string
	operand2 string
}

func (o *BIT_6_E) Execute(v *vm.VM) {
}

func (o *BIT_6_E) String() string {
	return "BIT 6 E"
}

type BIT_6_H struct {
	operand1 string
	operand2 string
}

func (o *BIT_6_H) Execute(v *vm.VM) {
}

func (o *BIT_6_H) String() string {
	return "BIT 6 H"
}

type BIT_6_L struct {
	operand1 string
	operand2 string
}

func (o *BIT_6_L) Execute(v *vm.VM) {
}

func (o *BIT_6_L) String() string {
	return "BIT 6 L"
}

type BIT_6_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *BIT_6_HLPtr) Execute(v *vm.VM) {
}

func (o *BIT_6_HLPtr) String() string {
	return "BIT 6 (HL)"
}

type BIT_6_A struct {
	operand1 string
	operand2 string
}

func (o *BIT_6_A) Execute(v *vm.VM) {
}

func (o *BIT_6_A) String() string {
	return "BIT 6 A"
}

type BIT_7_B struct {
	operand1 string
	operand2 string
}

func (o *BIT_7_B) Execute(v *vm.VM) {
}

func (o *BIT_7_B) String() string {
	return "BIT 7 B"
}

type BIT_7_C struct {
	operand1 string
	operand2 string
}

func (o *BIT_7_C) Execute(v *vm.VM) {
}

func (o *BIT_7_C) String() string {
	return "BIT 7 C"
}

type BIT_7_D struct {
	operand1 string
	operand2 string
}

func (o *BIT_7_D) Execute(v *vm.VM) {
}

func (o *BIT_7_D) String() string {
	return "BIT 7 D"
}

type BIT_7_E struct {
	operand1 string
	operand2 string
}

func (o *BIT_7_E) Execute(v *vm.VM) {
}

func (o *BIT_7_E) String() string {
	return "BIT 7 E"
}

type BIT_7_H struct {
	operand1 string
	operand2 string
}

func (o *BIT_7_H) Execute(v *vm.VM) {
}

func (o *BIT_7_H) String() string {
	return "BIT 7 H"
}

type BIT_7_L struct {
	operand1 string
	operand2 string
}

func (o *BIT_7_L) Execute(v *vm.VM) {
}

func (o *BIT_7_L) String() string {
	return "BIT 7 L"
}

type BIT_7_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *BIT_7_HLPtr) Execute(v *vm.VM) {
}

func (o *BIT_7_HLPtr) String() string {
	return "BIT 7 (HL)"
}

type BIT_7_A struct {
	operand1 string
	operand2 string
}

func (o *BIT_7_A) Execute(v *vm.VM) {
}

func (o *BIT_7_A) String() string {
	return "BIT 7 A"
}

type RRC_B struct {
	operand1 string
	operand2 string
}

func (o *RRC_B) Execute(v *vm.VM) {
}

func (o *RRC_B) String() string {
	return "RRC B"
}

type RES_0_B struct {
	operand1 string
	operand2 string
}

func (o *RES_0_B) Execute(v *vm.VM) {
}

func (o *RES_0_B) String() string {
	return "RES 0 B"
}

type RES_0_C struct {
	operand1 string
	operand2 string
}

func (o *RES_0_C) Execute(v *vm.VM) {
}

func (o *RES_0_C) String() string {
	return "RES 0 C"
}

type RES_0_D struct {
	operand1 string
	operand2 string
}

func (o *RES_0_D) Execute(v *vm.VM) {
}

func (o *RES_0_D) String() string {
	return "RES 0 D"
}

type RES_0_E struct {
	operand1 string
	operand2 string
}

func (o *RES_0_E) Execute(v *vm.VM) {
}

func (o *RES_0_E) String() string {
	return "RES 0 E"
}

type RES_0_H struct {
	operand1 string
	operand2 string
}

func (o *RES_0_H) Execute(v *vm.VM) {
}

func (o *RES_0_H) String() string {
	return "RES 0 H"
}

type RES_0_L struct {
	operand1 string
	operand2 string
}

func (o *RES_0_L) Execute(v *vm.VM) {
}

func (o *RES_0_L) String() string {
	return "RES 0 L"
}

type RES_0_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *RES_0_HLPtr) Execute(v *vm.VM) {
}

func (o *RES_0_HLPtr) String() string {
	return "RES 0 (HL)"
}

type RES_0_A struct {
	operand1 string
	operand2 string
}

func (o *RES_0_A) Execute(v *vm.VM) {
}

func (o *RES_0_A) String() string {
	return "RES 0 A"
}

type RES_1_B struct {
	operand1 string
	operand2 string
}

func (o *RES_1_B) Execute(v *vm.VM) {
}

func (o *RES_1_B) String() string {
	return "RES 1 B"
}

type RES_1_C struct {
	operand1 string
	operand2 string
}

func (o *RES_1_C) Execute(v *vm.VM) {
}

func (o *RES_1_C) String() string {
	return "RES 1 C"
}

type RES_1_D struct {
	operand1 string
	operand2 string
}

func (o *RES_1_D) Execute(v *vm.VM) {
}

func (o *RES_1_D) String() string {
	return "RES 1 D"
}

type RES_1_E struct {
	operand1 string
	operand2 string
}

func (o *RES_1_E) Execute(v *vm.VM) {
}

func (o *RES_1_E) String() string {
	return "RES 1 E"
}

type RES_1_H struct {
	operand1 string
	operand2 string
}

func (o *RES_1_H) Execute(v *vm.VM) {
}

func (o *RES_1_H) String() string {
	return "RES 1 H"
}

type RES_1_L struct {
	operand1 string
	operand2 string
}

func (o *RES_1_L) Execute(v *vm.VM) {
}

func (o *RES_1_L) String() string {
	return "RES 1 L"
}

type RES_1_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *RES_1_HLPtr) Execute(v *vm.VM) {
}

func (o *RES_1_HLPtr) String() string {
	return "RES 1 (HL)"
}

type RES_1_A struct {
	operand1 string
	operand2 string
}

func (o *RES_1_A) Execute(v *vm.VM) {
}

func (o *RES_1_A) String() string {
	return "RES 1 A"
}

type RRC_C struct {
	operand1 string
	operand2 string
}

func (o *RRC_C) Execute(v *vm.VM) {
}

func (o *RRC_C) String() string {
	return "RRC C"
}

type RES_2_B struct {
	operand1 string
	operand2 string
}

func (o *RES_2_B) Execute(v *vm.VM) {
}

func (o *RES_2_B) String() string {
	return "RES 2 B"
}

type RES_2_C struct {
	operand1 string
	operand2 string
}

func (o *RES_2_C) Execute(v *vm.VM) {
}

func (o *RES_2_C) String() string {
	return "RES 2 C"
}

type RES_2_D struct {
	operand1 string
	operand2 string
}

func (o *RES_2_D) Execute(v *vm.VM) {
}

func (o *RES_2_D) String() string {
	return "RES 2 D"
}

type RES_2_E struct {
	operand1 string
	operand2 string
}

func (o *RES_2_E) Execute(v *vm.VM) {
}

func (o *RES_2_E) String() string {
	return "RES 2 E"
}

type RES_2_H struct {
	operand1 string
	operand2 string
}

func (o *RES_2_H) Execute(v *vm.VM) {
}

func (o *RES_2_H) String() string {
	return "RES 2 H"
}

type RES_2_L struct {
	operand1 string
	operand2 string
}

func (o *RES_2_L) Execute(v *vm.VM) {
}

func (o *RES_2_L) String() string {
	return "RES 2 L"
}

type RES_2_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *RES_2_HLPtr) Execute(v *vm.VM) {
}

func (o *RES_2_HLPtr) String() string {
	return "RES 2 (HL)"
}

type RES_2_A struct {
	operand1 string
	operand2 string
}

func (o *RES_2_A) Execute(v *vm.VM) {
}

func (o *RES_2_A) String() string {
	return "RES 2 A"
}

type RES_3_B struct {
	operand1 string
	operand2 string
}

func (o *RES_3_B) Execute(v *vm.VM) {
}

func (o *RES_3_B) String() string {
	return "RES 3 B"
}

type RES_3_C struct {
	operand1 string
	operand2 string
}

func (o *RES_3_C) Execute(v *vm.VM) {
}

func (o *RES_3_C) String() string {
	return "RES 3 C"
}

type RES_3_D struct {
	operand1 string
	operand2 string
}

func (o *RES_3_D) Execute(v *vm.VM) {
}

func (o *RES_3_D) String() string {
	return "RES 3 D"
}

type RES_3_E struct {
	operand1 string
	operand2 string
}

func (o *RES_3_E) Execute(v *vm.VM) {
}

func (o *RES_3_E) String() string {
	return "RES 3 E"
}

type RES_3_H struct {
	operand1 string
	operand2 string
}

func (o *RES_3_H) Execute(v *vm.VM) {
}

func (o *RES_3_H) String() string {
	return "RES 3 H"
}

type RES_3_L struct {
	operand1 string
	operand2 string
}

func (o *RES_3_L) Execute(v *vm.VM) {
}

func (o *RES_3_L) String() string {
	return "RES 3 L"
}

type RES_3_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *RES_3_HLPtr) Execute(v *vm.VM) {
}

func (o *RES_3_HLPtr) String() string {
	return "RES 3 (HL)"
}

type RES_3_A struct {
	operand1 string
	operand2 string
}

func (o *RES_3_A) Execute(v *vm.VM) {
}

func (o *RES_3_A) String() string {
	return "RES 3 A"
}

type RRC_D struct {
	operand1 string
	operand2 string
}

func (o *RRC_D) Execute(v *vm.VM) {
}

func (o *RRC_D) String() string {
	return "RRC D"
}

type RES_4_B struct {
	operand1 string
	operand2 string
}

func (o *RES_4_B) Execute(v *vm.VM) {
}

func (o *RES_4_B) String() string {
	return "RES 4 B"
}

type RES_4_C struct {
	operand1 string
	operand2 string
}

func (o *RES_4_C) Execute(v *vm.VM) {
}

func (o *RES_4_C) String() string {
	return "RES 4 C"
}

type RES_4_D struct {
	operand1 string
	operand2 string
}

func (o *RES_4_D) Execute(v *vm.VM) {
}

func (o *RES_4_D) String() string {
	return "RES 4 D"
}

type RES_4_E struct {
	operand1 string
	operand2 string
}

func (o *RES_4_E) Execute(v *vm.VM) {
}

func (o *RES_4_E) String() string {
	return "RES 4 E"
}

type RES_4_H struct {
	operand1 string
	operand2 string
}

func (o *RES_4_H) Execute(v *vm.VM) {
}

func (o *RES_4_H) String() string {
	return "RES 4 H"
}

type RES_4_L struct {
	operand1 string
	operand2 string
}

func (o *RES_4_L) Execute(v *vm.VM) {
}

func (o *RES_4_L) String() string {
	return "RES 4 L"
}

type RES_4_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *RES_4_HLPtr) Execute(v *vm.VM) {
}

func (o *RES_4_HLPtr) String() string {
	return "RES 4 (HL)"
}

type RES_4_A struct {
	operand1 string
	operand2 string
}

func (o *RES_4_A) Execute(v *vm.VM) {
}

func (o *RES_4_A) String() string {
	return "RES 4 A"
}

type RES_5_B struct {
	operand1 string
	operand2 string
}

func (o *RES_5_B) Execute(v *vm.VM) {
}

func (o *RES_5_B) String() string {
	return "RES 5 B"
}

type RES_5_C struct {
	operand1 string
	operand2 string
}

func (o *RES_5_C) Execute(v *vm.VM) {
}

func (o *RES_5_C) String() string {
	return "RES 5 C"
}

type RES_5_D struct {
	operand1 string
	operand2 string
}

func (o *RES_5_D) Execute(v *vm.VM) {
}

func (o *RES_5_D) String() string {
	return "RES 5 D"
}

type RES_5_E struct {
	operand1 string
	operand2 string
}

func (o *RES_5_E) Execute(v *vm.VM) {
}

func (o *RES_5_E) String() string {
	return "RES 5 E"
}

type RES_5_H struct {
	operand1 string
	operand2 string
}

func (o *RES_5_H) Execute(v *vm.VM) {
}

func (o *RES_5_H) String() string {
	return "RES 5 H"
}

type RES_5_L struct {
	operand1 string
	operand2 string
}

func (o *RES_5_L) Execute(v *vm.VM) {
}

func (o *RES_5_L) String() string {
	return "RES 5 L"
}

type RES_5_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *RES_5_HLPtr) Execute(v *vm.VM) {
}

func (o *RES_5_HLPtr) String() string {
	return "RES 5 (HL)"
}

type RES_5_A struct {
	operand1 string
	operand2 string
}

func (o *RES_5_A) Execute(v *vm.VM) {
}

func (o *RES_5_A) String() string {
	return "RES 5 A"
}

type RRC_E struct {
	operand1 string
	operand2 string
}

func (o *RRC_E) Execute(v *vm.VM) {
}

func (o *RRC_E) String() string {
	return "RRC E"
}

type RES_6_B struct {
	operand1 string
	operand2 string
}

func (o *RES_6_B) Execute(v *vm.VM) {
}

func (o *RES_6_B) String() string {
	return "RES 6 B"
}

type RES_6_C struct {
	operand1 string
	operand2 string
}

func (o *RES_6_C) Execute(v *vm.VM) {
}

func (o *RES_6_C) String() string {
	return "RES 6 C"
}

type RES_6_D struct {
	operand1 string
	operand2 string
}

func (o *RES_6_D) Execute(v *vm.VM) {
}

func (o *RES_6_D) String() string {
	return "RES 6 D"
}

type RES_6_E struct {
	operand1 string
	operand2 string
}

func (o *RES_6_E) Execute(v *vm.VM) {
}

func (o *RES_6_E) String() string {
	return "RES 6 E"
}

type RES_6_H struct {
	operand1 string
	operand2 string
}

func (o *RES_6_H) Execute(v *vm.VM) {
}

func (o *RES_6_H) String() string {
	return "RES 6 H"
}

type RES_6_L struct {
	operand1 string
	operand2 string
}

func (o *RES_6_L) Execute(v *vm.VM) {
}

func (o *RES_6_L) String() string {
	return "RES 6 L"
}

type RES_6_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *RES_6_HLPtr) Execute(v *vm.VM) {
}

func (o *RES_6_HLPtr) String() string {
	return "RES 6 (HL)"
}

type RES_6_A struct {
	operand1 string
	operand2 string
}

func (o *RES_6_A) Execute(v *vm.VM) {
}

func (o *RES_6_A) String() string {
	return "RES 6 A"
}

type RES_7_B struct {
	operand1 string
	operand2 string
}

func (o *RES_7_B) Execute(v *vm.VM) {
}

func (o *RES_7_B) String() string {
	return "RES 7 B"
}

type RES_7_C struct {
	operand1 string
	operand2 string
}

func (o *RES_7_C) Execute(v *vm.VM) {
}

func (o *RES_7_C) String() string {
	return "RES 7 C"
}

type RES_7_D struct {
	operand1 string
	operand2 string
}

func (o *RES_7_D) Execute(v *vm.VM) {
}

func (o *RES_7_D) String() string {
	return "RES 7 D"
}

type RES_7_E struct {
	operand1 string
	operand2 string
}

func (o *RES_7_E) Execute(v *vm.VM) {
}

func (o *RES_7_E) String() string {
	return "RES 7 E"
}

type RES_7_H struct {
	operand1 string
	operand2 string
}

func (o *RES_7_H) Execute(v *vm.VM) {
}

func (o *RES_7_H) String() string {
	return "RES 7 H"
}

type RES_7_L struct {
	operand1 string
	operand2 string
}

func (o *RES_7_L) Execute(v *vm.VM) {
}

func (o *RES_7_L) String() string {
	return "RES 7 L"
}

type RES_7_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *RES_7_HLPtr) Execute(v *vm.VM) {
}

func (o *RES_7_HLPtr) String() string {
	return "RES 7 (HL)"
}

type RES_7_A struct {
	operand1 string
	operand2 string
}

func (o *RES_7_A) Execute(v *vm.VM) {
}

func (o *RES_7_A) String() string {
	return "RES 7 A"
}

type RRC_H struct {
	operand1 string
	operand2 string
}

func (o *RRC_H) Execute(v *vm.VM) {
}

func (o *RRC_H) String() string {
	return "RRC H"
}

type SET_0_B struct {
	operand1 string
	operand2 string
}

func (o *SET_0_B) Execute(v *vm.VM) {
}

func (o *SET_0_B) String() string {
	return "SET 0 B"
}

type SET_0_C struct {
	operand1 string
	operand2 string
}

func (o *SET_0_C) Execute(v *vm.VM) {
}

func (o *SET_0_C) String() string {
	return "SET 0 C"
}

type SET_0_D struct {
	operand1 string
	operand2 string
}

func (o *SET_0_D) Execute(v *vm.VM) {
}

func (o *SET_0_D) String() string {
	return "SET 0 D"
}

type SET_0_E struct {
	operand1 string
	operand2 string
}

func (o *SET_0_E) Execute(v *vm.VM) {
}

func (o *SET_0_E) String() string {
	return "SET 0 E"
}

type SET_0_H struct {
	operand1 string
	operand2 string
}

func (o *SET_0_H) Execute(v *vm.VM) {
}

func (o *SET_0_H) String() string {
	return "SET 0 H"
}

type SET_0_L struct {
	operand1 string
	operand2 string
}

func (o *SET_0_L) Execute(v *vm.VM) {
}

func (o *SET_0_L) String() string {
	return "SET 0 L"
}

type SET_0_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *SET_0_HLPtr) Execute(v *vm.VM) {
}

func (o *SET_0_HLPtr) String() string {
	return "SET 0 (HL)"
}

type SET_0_A struct {
	operand1 string
	operand2 string
}

func (o *SET_0_A) Execute(v *vm.VM) {
}

func (o *SET_0_A) String() string {
	return "SET 0 A"
}

type SET_1_B struct {
	operand1 string
	operand2 string
}

func (o *SET_1_B) Execute(v *vm.VM) {
}

func (o *SET_1_B) String() string {
	return "SET 1 B"
}

type SET_1_C struct {
	operand1 string
	operand2 string
}

func (o *SET_1_C) Execute(v *vm.VM) {
}

func (o *SET_1_C) String() string {
	return "SET 1 C"
}

type SET_1_D struct {
	operand1 string
	operand2 string
}

func (o *SET_1_D) Execute(v *vm.VM) {
}

func (o *SET_1_D) String() string {
	return "SET 1 D"
}

type SET_1_E struct {
	operand1 string
	operand2 string
}

func (o *SET_1_E) Execute(v *vm.VM) {
}

func (o *SET_1_E) String() string {
	return "SET 1 E"
}

type SET_1_H struct {
	operand1 string
	operand2 string
}

func (o *SET_1_H) Execute(v *vm.VM) {
}

func (o *SET_1_H) String() string {
	return "SET 1 H"
}

type SET_1_L struct {
	operand1 string
	operand2 string
}

func (o *SET_1_L) Execute(v *vm.VM) {
}

func (o *SET_1_L) String() string {
	return "SET 1 L"
}

type SET_1_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *SET_1_HLPtr) Execute(v *vm.VM) {
}

func (o *SET_1_HLPtr) String() string {
	return "SET 1 (HL)"
}

type SET_1_A struct {
	operand1 string
	operand2 string
}

func (o *SET_1_A) Execute(v *vm.VM) {
}

func (o *SET_1_A) String() string {
	return "SET 1 A"
}

type RRC_L struct {
	operand1 string
	operand2 string
}

func (o *RRC_L) Execute(v *vm.VM) {
}

func (o *RRC_L) String() string {
	return "RRC L"
}

type SET_2_B struct {
	operand1 string
	operand2 string
}

func (o *SET_2_B) Execute(v *vm.VM) {
}

func (o *SET_2_B) String() string {
	return "SET 2 B"
}

type SET_2_C struct {
	operand1 string
	operand2 string
}

func (o *SET_2_C) Execute(v *vm.VM) {
}

func (o *SET_2_C) String() string {
	return "SET 2 C"
}

type SET_2_D struct {
	operand1 string
	operand2 string
}

func (o *SET_2_D) Execute(v *vm.VM) {
}

func (o *SET_2_D) String() string {
	return "SET 2 D"
}

type SET_2_E struct {
	operand1 string
	operand2 string
}

func (o *SET_2_E) Execute(v *vm.VM) {
}

func (o *SET_2_E) String() string {
	return "SET 2 E"
}

type SET_2_H struct {
	operand1 string
	operand2 string
}

func (o *SET_2_H) Execute(v *vm.VM) {
}

func (o *SET_2_H) String() string {
	return "SET 2 H"
}

type SET_2_L struct {
	operand1 string
	operand2 string
}

func (o *SET_2_L) Execute(v *vm.VM) {
}

func (o *SET_2_L) String() string {
	return "SET 2 L"
}

type SET_2_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *SET_2_HLPtr) Execute(v *vm.VM) {
}

func (o *SET_2_HLPtr) String() string {
	return "SET 2 (HL)"
}

type SET_2_A struct {
	operand1 string
	operand2 string
}

func (o *SET_2_A) Execute(v *vm.VM) {
}

func (o *SET_2_A) String() string {
	return "SET 2 A"
}

type SET_3_B struct {
	operand1 string
	operand2 string
}

func (o *SET_3_B) Execute(v *vm.VM) {
}

func (o *SET_3_B) String() string {
	return "SET 3 B"
}

type SET_3_C struct {
	operand1 string
	operand2 string
}

func (o *SET_3_C) Execute(v *vm.VM) {
}

func (o *SET_3_C) String() string {
	return "SET 3 C"
}

type SET_3_D struct {
	operand1 string
	operand2 string
}

func (o *SET_3_D) Execute(v *vm.VM) {
}

func (o *SET_3_D) String() string {
	return "SET 3 D"
}

type SET_3_E struct {
	operand1 string
	operand2 string
}

func (o *SET_3_E) Execute(v *vm.VM) {
}

func (o *SET_3_E) String() string {
	return "SET 3 E"
}

type SET_3_H struct {
	operand1 string
	operand2 string
}

func (o *SET_3_H) Execute(v *vm.VM) {
}

func (o *SET_3_H) String() string {
	return "SET 3 H"
}

type SET_3_L struct {
	operand1 string
	operand2 string
}

func (o *SET_3_L) Execute(v *vm.VM) {
}

func (o *SET_3_L) String() string {
	return "SET 3 L"
}

type SET_3_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *SET_3_HLPtr) Execute(v *vm.VM) {
}

func (o *SET_3_HLPtr) String() string {
	return "SET 3 (HL)"
}

type SET_3_A struct {
	operand1 string
	operand2 string
}

func (o *SET_3_A) Execute(v *vm.VM) {
}

func (o *SET_3_A) String() string {
	return "SET 3 A"
}

type RRC_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *RRC_HLPtr) Execute(v *vm.VM) {
}

func (o *RRC_HLPtr) String() string {
	return "RRC (HL)"
}

type SET_4_B struct {
	operand1 string
	operand2 string
}

func (o *SET_4_B) Execute(v *vm.VM) {
}

func (o *SET_4_B) String() string {
	return "SET 4 B"
}

type SET_4_C struct {
	operand1 string
	operand2 string
}

func (o *SET_4_C) Execute(v *vm.VM) {
}

func (o *SET_4_C) String() string {
	return "SET 4 C"
}

type SET_4_D struct {
	operand1 string
	operand2 string
}

func (o *SET_4_D) Execute(v *vm.VM) {
}

func (o *SET_4_D) String() string {
	return "SET 4 D"
}

type SET_4_E struct {
	operand1 string
	operand2 string
}

func (o *SET_4_E) Execute(v *vm.VM) {
}

func (o *SET_4_E) String() string {
	return "SET 4 E"
}

type SET_4_H struct {
	operand1 string
	operand2 string
}

func (o *SET_4_H) Execute(v *vm.VM) {
}

func (o *SET_4_H) String() string {
	return "SET 4 H"
}

type SET_4_L struct {
	operand1 string
	operand2 string
}

func (o *SET_4_L) Execute(v *vm.VM) {
}

func (o *SET_4_L) String() string {
	return "SET 4 L"
}

type SET_4_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *SET_4_HLPtr) Execute(v *vm.VM) {
}

func (o *SET_4_HLPtr) String() string {
	return "SET 4 (HL)"
}

type SET_4_A struct {
	operand1 string
	operand2 string
}

func (o *SET_4_A) Execute(v *vm.VM) {
}

func (o *SET_4_A) String() string {
	return "SET 4 A"
}

type SET_5_B struct {
	operand1 string
	operand2 string
}

func (o *SET_5_B) Execute(v *vm.VM) {
}

func (o *SET_5_B) String() string {
	return "SET 5 B"
}

type SET_5_C struct {
	operand1 string
	operand2 string
}

func (o *SET_5_C) Execute(v *vm.VM) {
}

func (o *SET_5_C) String() string {
	return "SET 5 C"
}

type SET_5_D struct {
	operand1 string
	operand2 string
}

func (o *SET_5_D) Execute(v *vm.VM) {
}

func (o *SET_5_D) String() string {
	return "SET 5 D"
}

type SET_5_E struct {
	operand1 string
	operand2 string
}

func (o *SET_5_E) Execute(v *vm.VM) {
}

func (o *SET_5_E) String() string {
	return "SET 5 E"
}

type SET_5_H struct {
	operand1 string
	operand2 string
}

func (o *SET_5_H) Execute(v *vm.VM) {
}

func (o *SET_5_H) String() string {
	return "SET 5 H"
}

type SET_5_L struct {
	operand1 string
	operand2 string
}

func (o *SET_5_L) Execute(v *vm.VM) {
}

func (o *SET_5_L) String() string {
	return "SET 5 L"
}

type SET_5_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *SET_5_HLPtr) Execute(v *vm.VM) {
}

func (o *SET_5_HLPtr) String() string {
	return "SET 5 (HL)"
}

type SET_5_A struct {
	operand1 string
	operand2 string
}

func (o *SET_5_A) Execute(v *vm.VM) {
}

func (o *SET_5_A) String() string {
	return "SET 5 A"
}

type RRC_A struct {
	operand1 string
	operand2 string
}

func (o *RRC_A) Execute(v *vm.VM) {
}

func (o *RRC_A) String() string {
	return "RRC A"
}

type SET_6_B struct {
	operand1 string
	operand2 string
}

func (o *SET_6_B) Execute(v *vm.VM) {
}

func (o *SET_6_B) String() string {
	return "SET 6 B"
}

type SET_6_C struct {
	operand1 string
	operand2 string
}

func (o *SET_6_C) Execute(v *vm.VM) {
}

func (o *SET_6_C) String() string {
	return "SET 6 C"
}

type SET_6_D struct {
	operand1 string
	operand2 string
}

func (o *SET_6_D) Execute(v *vm.VM) {
}

func (o *SET_6_D) String() string {
	return "SET 6 D"
}

type SET_6_E struct {
	operand1 string
	operand2 string
}

func (o *SET_6_E) Execute(v *vm.VM) {
}

func (o *SET_6_E) String() string {
	return "SET 6 E"
}

type SET_6_H struct {
	operand1 string
	operand2 string
}

func (o *SET_6_H) Execute(v *vm.VM) {
}

func (o *SET_6_H) String() string {
	return "SET 6 H"
}

type SET_6_L struct {
	operand1 string
	operand2 string
}

func (o *SET_6_L) Execute(v *vm.VM) {
}

func (o *SET_6_L) String() string {
	return "SET 6 L"
}

type SET_6_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *SET_6_HLPtr) Execute(v *vm.VM) {
}

func (o *SET_6_HLPtr) String() string {
	return "SET 6 (HL)"
}

type SET_6_A struct {
	operand1 string
	operand2 string
}

func (o *SET_6_A) Execute(v *vm.VM) {
}

func (o *SET_6_A) String() string {
	return "SET 6 A"
}

type SET_7_B struct {
	operand1 string
	operand2 string
}

func (o *SET_7_B) Execute(v *vm.VM) {
}

func (o *SET_7_B) String() string {
	return "SET 7 B"
}

type SET_7_C struct {
	operand1 string
	operand2 string
}

func (o *SET_7_C) Execute(v *vm.VM) {
}

func (o *SET_7_C) String() string {
	return "SET 7 C"
}

type SET_7_D struct {
	operand1 string
	operand2 string
}

func (o *SET_7_D) Execute(v *vm.VM) {
}

func (o *SET_7_D) String() string {
	return "SET 7 D"
}

type SET_7_E struct {
	operand1 string
	operand2 string
}

func (o *SET_7_E) Execute(v *vm.VM) {
}

func (o *SET_7_E) String() string {
	return "SET 7 E"
}

type SET_7_H struct {
	operand1 string
	operand2 string
}

func (o *SET_7_H) Execute(v *vm.VM) {
}

func (o *SET_7_H) String() string {
	return "SET 7 H"
}

type SET_7_L struct {
	operand1 string
	operand2 string
}

func (o *SET_7_L) Execute(v *vm.VM) {
}

func (o *SET_7_L) String() string {
	return "SET 7 L"
}

type SET_7_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *SET_7_HLPtr) Execute(v *vm.VM) {
}

func (o *SET_7_HLPtr) String() string {
	return "SET 7 (HL)"
}

type SET_7_A struct {
	operand1 string
	operand2 string
}

func (o *SET_7_A) Execute(v *vm.VM) {
}

func (o *SET_7_A) String() string {
	return "SET 7 A"
}

// ReadOpCode returns an executable opcode by taking an io.Reader
// and reading a single instruction off it. If there is no more data
// returns undelying io.Reader's EOF error type.
func ReadOpCode(data io.Reader) (OpCode, error) {
	d, err := readByte(data)
	if err != nil {
		return nil, err
	}

	switch d[0] {

	case 0x0: // 0x0 - NOP
		o := &NOP_{}

		var s string

		s = ""
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x1: // 0x1 - LD
		o := &LD_BC_d16{}

		var s string

		s = "BC"
		o.operand1 = s

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0x10: // 0x10 - STOP
		o := &STOP_0{}

		var s string

		s = "0"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x11: // 0x11 - LD
		o := &LD_DE_d16{}

		var s string

		s = "DE"
		o.operand1 = s

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0x12: // 0x12 - LD
		o := &LD_DEDeref_A{}

		var s string

		s = "(DE)"
		o.operand1 = s

		s = "A"
		o.operand2 = s

		return o, nil

	case 0x13: // 0x13 - INC
		o := &INC_DE{}

		var s string

		s = "DE"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x14: // 0x14 - INC
		o := &INC_D{}

		var s string

		s = "D"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x15: // 0x15 - DEC
		o := &DEC_D{}

		var s string

		s = "D"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x16: // 0x16 - LD
		o := &LD_D_d8{}

		var s string

		s = "D"
		o.operand1 = s

		s, err = readBytesAsString(data, 1)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0x17: // 0x17 - RLA
		o := &RLA_{}

		var s string

		s = ""
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x18: // 0x18 - JR
		o := &JR_r8{}

		var s string

		s = "r8"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x19: // 0x19 - ADD
		o := &ADD_HL_DE{}

		var s string

		s = "HL"
		o.operand1 = s

		s = "DE"
		o.operand2 = s

		return o, nil

	case 0x1a: // 0x1a - LD
		o := &LD_A_DEDeref{}

		var s string

		s = "A"
		o.operand1 = s

		s = "(DE)"
		o.operand2 = s

		return o, nil

	case 0x1b: // 0x1b - DEC
		o := &DEC_DE{}

		var s string

		s = "DE"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x1c: // 0x1c - INC
		o := &INC_E{}

		var s string

		s = "E"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x1d: // 0x1d - DEC
		o := &DEC_E{}

		var s string

		s = "E"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x1e: // 0x1e - LD
		o := &LD_E_d8{}

		var s string

		s = "E"
		o.operand1 = s

		s, err = readBytesAsString(data, 1)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0x1f: // 0x1f - RRA
		o := &RRA_{}

		var s string

		s = ""
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x2: // 0x2 - LD
		o := &LD_BCDeref_A{}

		var s string

		s = "(BC)"
		o.operand1 = s

		s = "A"
		o.operand2 = s

		return o, nil

	case 0x20: // 0x20 - JR
		o := &JR_NZ_r8{}

		var s string

		s = "NZ"
		o.operand1 = s

		s = "r8"
		o.operand2 = s

		return o, nil

	case 0x21: // 0x21 - LD
		o := &LD_HL_d16{}

		var s string

		s = "HL"
		o.operand1 = s

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0x22: // 0x22 - LD
		o := &LD_HLPtrInc_A{}

		var s string

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand1 = s

		s = "A"
		o.operand2 = s

		return o, nil

	case 0x23: // 0x23 - INC
		o := &INC_HL{}

		var s string

		s = "HL"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x24: // 0x24 - INC
		o := &INC_H{}

		var s string

		s = "H"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x25: // 0x25 - DEC
		o := &DEC_H{}

		var s string

		s = "H"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x26: // 0x26 - LD
		o := &LD_H_d8{}

		var s string

		s = "H"
		o.operand1 = s

		s, err = readBytesAsString(data, 1)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0x27: // 0x27 - DAA
		o := &DAA_{}

		var s string

		s = ""
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x28: // 0x28 - JR
		o := &JR_Z_r8{}

		var s string

		s = "Z"
		o.operand1 = s

		s = "r8"
		o.operand2 = s

		return o, nil

	case 0x29: // 0x29 - ADD
		o := &ADD_HL_HL{}

		var s string

		s = "HL"
		o.operand1 = s

		s = "HL"
		o.operand2 = s

		return o, nil

	case 0x2a: // 0x2a - LD
		o := &LD_A_HLPtrInc{}

		var s string

		s = "A"
		o.operand1 = s

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0x2b: // 0x2b - DEC
		o := &DEC_HL{}

		var s string

		s = "HL"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x2c: // 0x2c - INC
		o := &INC_L{}

		var s string

		s = "L"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x2d: // 0x2d - DEC
		o := &DEC_L{}

		var s string

		s = "L"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x2e: // 0x2e - LD
		o := &LD_L_d8{}

		var s string

		s = "L"
		o.operand1 = s

		s, err = readBytesAsString(data, 1)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0x2f: // 0x2f - CPL
		o := &CPL_{}

		var s string

		s = ""
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x3: // 0x3 - INC
		o := &INC_BC{}

		var s string

		s = "BC"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x30: // 0x30 - JR
		o := &JR_NC_r8{}

		var s string

		s = "NC"
		o.operand1 = s

		s = "r8"
		o.operand2 = s

		return o, nil

	case 0x31: // 0x31 - LD
		o := &LD_SP_d16{}

		var s string

		s = "SP"
		o.operand1 = s

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0x32: // 0x32 - LD
		o := &LD_HLPtrDec_A{}

		var s string

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand1 = s

		s = "A"
		o.operand2 = s

		return o, nil

	case 0x33: // 0x33 - INC
		o := &INC_SP{}

		var s string

		s = "SP"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x34: // 0x34 - INC
		o := &INC_HLPtr{}

		var s string

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x35: // 0x35 - DEC
		o := &DEC_HLPtr{}

		var s string

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x36: // 0x36 - LD
		o := &LD_HLPtr_d8{}

		var s string

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand1 = s

		s, err = readBytesAsString(data, 1)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0x37: // 0x37 - SCF
		o := &SCF_{}

		var s string

		s = ""
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x38: // 0x38 - JR
		o := &JR_C_r8{}

		var s string

		s = "C"
		o.operand1 = s

		s = "r8"
		o.operand2 = s

		return o, nil

	case 0x39: // 0x39 - ADD
		o := &ADD_HL_SP{}

		var s string

		s = "HL"
		o.operand1 = s

		s = "SP"
		o.operand2 = s

		return o, nil

	case 0x3a: // 0x3a - LD
		o := &LD_A_HLPtrDec{}

		var s string

		s = "A"
		o.operand1 = s

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0x3b: // 0x3b - DEC
		o := &DEC_SP{}

		var s string

		s = "SP"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x3c: // 0x3c - INC
		o := &INC_A{}

		var s string

		s = "A"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x3d: // 0x3d - DEC
		o := &DEC_A{}

		var s string

		s = "A"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x3e: // 0x3e - LD
		o := &LD_A_d8{}

		var s string

		s = "A"
		o.operand1 = s

		s, err = readBytesAsString(data, 1)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0x3f: // 0x3f - CCF
		o := &CCF_{}

		var s string

		s = ""
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x4: // 0x4 - INC
		o := &INC_B{}

		var s string

		s = "B"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x40: // 0x40 - LD
		o := &LD_B_B{}

		var s string

		s = "B"
		o.operand1 = s

		s = "B"
		o.operand2 = s

		return o, nil

	case 0x41: // 0x41 - LD
		o := &LD_B_C{}

		var s string

		s = "B"
		o.operand1 = s

		s = "C"
		o.operand2 = s

		return o, nil

	case 0x42: // 0x42 - LD
		o := &LD_B_D{}

		var s string

		s = "B"
		o.operand1 = s

		s = "D"
		o.operand2 = s

		return o, nil

	case 0x43: // 0x43 - LD
		o := &LD_B_E{}

		var s string

		s = "B"
		o.operand1 = s

		s = "E"
		o.operand2 = s

		return o, nil

	case 0x44: // 0x44 - LD
		o := &LD_B_H{}

		var s string

		s = "B"
		o.operand1 = s

		s = "H"
		o.operand2 = s

		return o, nil

	case 0x45: // 0x45 - LD
		o := &LD_B_L{}

		var s string

		s = "B"
		o.operand1 = s

		s = "L"
		o.operand2 = s

		return o, nil

	case 0x46: // 0x46 - LD
		o := &LD_B_HLPtr{}

		var s string

		s = "B"
		o.operand1 = s

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0x47: // 0x47 - LD
		o := &LD_B_A{}

		var s string

		s = "B"
		o.operand1 = s

		s = "A"
		o.operand2 = s

		return o, nil

	case 0x48: // 0x48 - LD
		o := &LD_C_B{}

		var s string

		s = "C"
		o.operand1 = s

		s = "B"
		o.operand2 = s

		return o, nil

	case 0x49: // 0x49 - LD
		o := &LD_C_C{}

		var s string

		s = "C"
		o.operand1 = s

		s = "C"
		o.operand2 = s

		return o, nil

	case 0x4a: // 0x4a - LD
		o := &LD_C_D{}

		var s string

		s = "C"
		o.operand1 = s

		s = "D"
		o.operand2 = s

		return o, nil

	case 0x4b: // 0x4b - LD
		o := &LD_C_E{}

		var s string

		s = "C"
		o.operand1 = s

		s = "E"
		o.operand2 = s

		return o, nil

	case 0x4c: // 0x4c - LD
		o := &LD_C_H{}

		var s string

		s = "C"
		o.operand1 = s

		s = "H"
		o.operand2 = s

		return o, nil

	case 0x4d: // 0x4d - LD
		o := &LD_C_L{}

		var s string

		s = "C"
		o.operand1 = s

		s = "L"
		o.operand2 = s

		return o, nil

	case 0x4e: // 0x4e - LD
		o := &LD_C_HLPtr{}

		var s string

		s = "C"
		o.operand1 = s

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0x4f: // 0x4f - LD
		o := &LD_C_A{}

		var s string

		s = "C"
		o.operand1 = s

		s = "A"
		o.operand2 = s

		return o, nil

	case 0x5: // 0x5 - DEC
		o := &DEC_B{}

		var s string

		s = "B"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x50: // 0x50 - LD
		o := &LD_D_B{}

		var s string

		s = "D"
		o.operand1 = s

		s = "B"
		o.operand2 = s

		return o, nil

	case 0x51: // 0x51 - LD
		o := &LD_D_C{}

		var s string

		s = "D"
		o.operand1 = s

		s = "C"
		o.operand2 = s

		return o, nil

	case 0x52: // 0x52 - LD
		o := &LD_D_D{}

		var s string

		s = "D"
		o.operand1 = s

		s = "D"
		o.operand2 = s

		return o, nil

	case 0x53: // 0x53 - LD
		o := &LD_D_E{}

		var s string

		s = "D"
		o.operand1 = s

		s = "E"
		o.operand2 = s

		return o, nil

	case 0x54: // 0x54 - LD
		o := &LD_D_H{}

		var s string

		s = "D"
		o.operand1 = s

		s = "H"
		o.operand2 = s

		return o, nil

	case 0x55: // 0x55 - LD
		o := &LD_D_L{}

		var s string

		s = "D"
		o.operand1 = s

		s = "L"
		o.operand2 = s

		return o, nil

	case 0x56: // 0x56 - LD
		o := &LD_D_HLPtr{}

		var s string

		s = "D"
		o.operand1 = s

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0x57: // 0x57 - LD
		o := &LD_D_A{}

		var s string

		s = "D"
		o.operand1 = s

		s = "A"
		o.operand2 = s

		return o, nil

	case 0x58: // 0x58 - LD
		o := &LD_E_B{}

		var s string

		s = "E"
		o.operand1 = s

		s = "B"
		o.operand2 = s

		return o, nil

	case 0x59: // 0x59 - LD
		o := &LD_E_C{}

		var s string

		s = "E"
		o.operand1 = s

		s = "C"
		o.operand2 = s

		return o, nil

	case 0x5a: // 0x5a - LD
		o := &LD_E_D{}

		var s string

		s = "E"
		o.operand1 = s

		s = "D"
		o.operand2 = s

		return o, nil

	case 0x5b: // 0x5b - LD
		o := &LD_E_E{}

		var s string

		s = "E"
		o.operand1 = s

		s = "E"
		o.operand2 = s

		return o, nil

	case 0x5c: // 0x5c - LD
		o := &LD_E_H{}

		var s string

		s = "E"
		o.operand1 = s

		s = "H"
		o.operand2 = s

		return o, nil

	case 0x5d: // 0x5d - LD
		o := &LD_E_L{}

		var s string

		s = "E"
		o.operand1 = s

		s = "L"
		o.operand2 = s

		return o, nil

	case 0x5e: // 0x5e - LD
		o := &LD_E_HLPtr{}

		var s string

		s = "E"
		o.operand1 = s

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0x5f: // 0x5f - LD
		o := &LD_E_A{}

		var s string

		s = "E"
		o.operand1 = s

		s = "A"
		o.operand2 = s

		return o, nil

	case 0x6: // 0x6 - LD
		o := &LD_B_d8{}

		var s string

		s = "B"
		o.operand1 = s

		s, err = readBytesAsString(data, 1)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0x60: // 0x60 - LD
		o := &LD_H_B{}

		var s string

		s = "H"
		o.operand1 = s

		s = "B"
		o.operand2 = s

		return o, nil

	case 0x61: // 0x61 - LD
		o := &LD_H_C{}

		var s string

		s = "H"
		o.operand1 = s

		s = "C"
		o.operand2 = s

		return o, nil

	case 0x62: // 0x62 - LD
		o := &LD_H_D{}

		var s string

		s = "H"
		o.operand1 = s

		s = "D"
		o.operand2 = s

		return o, nil

	case 0x63: // 0x63 - LD
		o := &LD_H_E{}

		var s string

		s = "H"
		o.operand1 = s

		s = "E"
		o.operand2 = s

		return o, nil

	case 0x64: // 0x64 - LD
		o := &LD_H_H{}

		var s string

		s = "H"
		o.operand1 = s

		s = "H"
		o.operand2 = s

		return o, nil

	case 0x65: // 0x65 - LD
		o := &LD_H_L{}

		var s string

		s = "H"
		o.operand1 = s

		s = "L"
		o.operand2 = s

		return o, nil

	case 0x66: // 0x66 - LD
		o := &LD_H_HLPtr{}

		var s string

		s = "H"
		o.operand1 = s

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0x67: // 0x67 - LD
		o := &LD_H_A{}

		var s string

		s = "H"
		o.operand1 = s

		s = "A"
		o.operand2 = s

		return o, nil

	case 0x68: // 0x68 - LD
		o := &LD_L_B{}

		var s string

		s = "L"
		o.operand1 = s

		s = "B"
		o.operand2 = s

		return o, nil

	case 0x69: // 0x69 - LD
		o := &LD_L_C{}

		var s string

		s = "L"
		o.operand1 = s

		s = "C"
		o.operand2 = s

		return o, nil

	case 0x6a: // 0x6a - LD
		o := &LD_L_D{}

		var s string

		s = "L"
		o.operand1 = s

		s = "D"
		o.operand2 = s

		return o, nil

	case 0x6b: // 0x6b - LD
		o := &LD_L_E{}

		var s string

		s = "L"
		o.operand1 = s

		s = "E"
		o.operand2 = s

		return o, nil

	case 0x6c: // 0x6c - LD
		o := &LD_L_H{}

		var s string

		s = "L"
		o.operand1 = s

		s = "H"
		o.operand2 = s

		return o, nil

	case 0x6d: // 0x6d - LD
		o := &LD_L_L{}

		var s string

		s = "L"
		o.operand1 = s

		s = "L"
		o.operand2 = s

		return o, nil

	case 0x6e: // 0x6e - LD
		o := &LD_L_HLPtr{}

		var s string

		s = "L"
		o.operand1 = s

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0x6f: // 0x6f - LD
		o := &LD_L_A{}

		var s string

		s = "L"
		o.operand1 = s

		s = "A"
		o.operand2 = s

		return o, nil

	case 0x7: // 0x7 - RLCA
		o := &RLCA_{}

		var s string

		s = ""
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x70: // 0x70 - LD
		o := &LD_HLPtr_B{}

		var s string

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand1 = s

		s = "B"
		o.operand2 = s

		return o, nil

	case 0x71: // 0x71 - LD
		o := &LD_HLPtr_C{}

		var s string

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand1 = s

		s = "C"
		o.operand2 = s

		return o, nil

	case 0x72: // 0x72 - LD
		o := &LD_HLPtr_D{}

		var s string

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand1 = s

		s = "D"
		o.operand2 = s

		return o, nil

	case 0x73: // 0x73 - LD
		o := &LD_HLPtr_E{}

		var s string

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand1 = s

		s = "E"
		o.operand2 = s

		return o, nil

	case 0x74: // 0x74 - LD
		o := &LD_HLPtr_H{}

		var s string

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand1 = s

		s = "H"
		o.operand2 = s

		return o, nil

	case 0x75: // 0x75 - LD
		o := &LD_HLPtr_L{}

		var s string

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand1 = s

		s = "L"
		o.operand2 = s

		return o, nil

	case 0x76: // 0x76 - HALT
		o := &HALT_{}

		var s string

		s = ""
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x77: // 0x77 - LD
		o := &LD_HLPtr_A{}

		var s string

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand1 = s

		s = "A"
		o.operand2 = s

		return o, nil

	case 0x78: // 0x78 - LD
		o := &LD_A_B{}

		var s string

		s = "A"
		o.operand1 = s

		s = "B"
		o.operand2 = s

		return o, nil

	case 0x79: // 0x79 - LD
		o := &LD_A_C{}

		var s string

		s = "A"
		o.operand1 = s

		s = "C"
		o.operand2 = s

		return o, nil

	case 0x7a: // 0x7a - LD
		o := &LD_A_D{}

		var s string

		s = "A"
		o.operand1 = s

		s = "D"
		o.operand2 = s

		return o, nil

	case 0x7b: // 0x7b - LD
		o := &LD_A_E{}

		var s string

		s = "A"
		o.operand1 = s

		s = "E"
		o.operand2 = s

		return o, nil

	case 0x7c: // 0x7c - LD
		o := &LD_A_H{}

		var s string

		s = "A"
		o.operand1 = s

		s = "H"
		o.operand2 = s

		return o, nil

	case 0x7d: // 0x7d - LD
		o := &LD_A_L{}

		var s string

		s = "A"
		o.operand1 = s

		s = "L"
		o.operand2 = s

		return o, nil

	case 0x7e: // 0x7e - LD
		o := &LD_A_HLPtr{}

		var s string

		s = "A"
		o.operand1 = s

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0x7f: // 0x7f - LD
		o := &LD_A_A{}

		var s string

		s = "A"
		o.operand1 = s

		s = "A"
		o.operand2 = s

		return o, nil

	case 0x8: // 0x8 - LD
		o := &LD_a16Deref_SP{}

		var s string

		s = "(a16)"
		o.operand1 = s

		s = "SP"
		o.operand2 = s

		return o, nil

	case 0x80: // 0x80 - ADD
		o := &ADD_A_B{}

		var s string

		s = "A"
		o.operand1 = s

		s = "B"
		o.operand2 = s

		return o, nil

	case 0x81: // 0x81 - ADD
		o := &ADD_A_C{}

		var s string

		s = "A"
		o.operand1 = s

		s = "C"
		o.operand2 = s

		return o, nil

	case 0x82: // 0x82 - ADD
		o := &ADD_A_D{}

		var s string

		s = "A"
		o.operand1 = s

		s = "D"
		o.operand2 = s

		return o, nil

	case 0x83: // 0x83 - ADD
		o := &ADD_A_E{}

		var s string

		s = "A"
		o.operand1 = s

		s = "E"
		o.operand2 = s

		return o, nil

	case 0x84: // 0x84 - ADD
		o := &ADD_A_H{}

		var s string

		s = "A"
		o.operand1 = s

		s = "H"
		o.operand2 = s

		return o, nil

	case 0x85: // 0x85 - ADD
		o := &ADD_A_L{}

		var s string

		s = "A"
		o.operand1 = s

		s = "L"
		o.operand2 = s

		return o, nil

	case 0x86: // 0x86 - ADD
		o := &ADD_A_HLPtr{}

		var s string

		s = "A"
		o.operand1 = s

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0x87: // 0x87 - ADD
		o := &ADD_A_A{}

		var s string

		s = "A"
		o.operand1 = s

		s = "A"
		o.operand2 = s

		return o, nil

	case 0x88: // 0x88 - ADC
		o := &ADC_A_B{}

		var s string

		s = "A"
		o.operand1 = s

		s = "B"
		o.operand2 = s

		return o, nil

	case 0x89: // 0x89 - ADC
		o := &ADC_A_C{}

		var s string

		s = "A"
		o.operand1 = s

		s = "C"
		o.operand2 = s

		return o, nil

	case 0x8a: // 0x8a - ADC
		o := &ADC_A_D{}

		var s string

		s = "A"
		o.operand1 = s

		s = "D"
		o.operand2 = s

		return o, nil

	case 0x8b: // 0x8b - ADC
		o := &ADC_A_E{}

		var s string

		s = "A"
		o.operand1 = s

		s = "E"
		o.operand2 = s

		return o, nil

	case 0x8c: // 0x8c - ADC
		o := &ADC_A_H{}

		var s string

		s = "A"
		o.operand1 = s

		s = "H"
		o.operand2 = s

		return o, nil

	case 0x8d: // 0x8d - ADC
		o := &ADC_A_L{}

		var s string

		s = "A"
		o.operand1 = s

		s = "L"
		o.operand2 = s

		return o, nil

	case 0x8e: // 0x8e - ADC
		o := &ADC_A_HLPtr{}

		var s string

		s = "A"
		o.operand1 = s

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0x8f: // 0x8f - ADC
		o := &ADC_A_A{}

		var s string

		s = "A"
		o.operand1 = s

		s = "A"
		o.operand2 = s

		return o, nil

	case 0x9: // 0x9 - ADD
		o := &ADD_HL_BC{}

		var s string

		s = "HL"
		o.operand1 = s

		s = "BC"
		o.operand2 = s

		return o, nil

	case 0x90: // 0x90 - SUB
		o := &SUB_B{}

		var s string

		s = "B"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x91: // 0x91 - SUB
		o := &SUB_C{}

		var s string

		s = "C"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x92: // 0x92 - SUB
		o := &SUB_D{}

		var s string

		s = "D"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x93: // 0x93 - SUB
		o := &SUB_E{}

		var s string

		s = "E"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x94: // 0x94 - SUB
		o := &SUB_H{}

		var s string

		s = "H"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x95: // 0x95 - SUB
		o := &SUB_L{}

		var s string

		s = "L"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x96: // 0x96 - SUB
		o := &SUB_HLPtr{}

		var s string

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x97: // 0x97 - SUB
		o := &SUB_A{}

		var s string

		s = "A"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0x98: // 0x98 - SBC
		o := &SBC_A_B{}

		var s string

		s = "A"
		o.operand1 = s

		s = "B"
		o.operand2 = s

		return o, nil

	case 0x99: // 0x99 - SBC
		o := &SBC_A_C{}

		var s string

		s = "A"
		o.operand1 = s

		s = "C"
		o.operand2 = s

		return o, nil

	case 0x9a: // 0x9a - SBC
		o := &SBC_A_D{}

		var s string

		s = "A"
		o.operand1 = s

		s = "D"
		o.operand2 = s

		return o, nil

	case 0x9b: // 0x9b - SBC
		o := &SBC_A_E{}

		var s string

		s = "A"
		o.operand1 = s

		s = "E"
		o.operand2 = s

		return o, nil

	case 0x9c: // 0x9c - SBC
		o := &SBC_A_H{}

		var s string

		s = "A"
		o.operand1 = s

		s = "H"
		o.operand2 = s

		return o, nil

	case 0x9d: // 0x9d - SBC
		o := &SBC_A_L{}

		var s string

		s = "A"
		o.operand1 = s

		s = "L"
		o.operand2 = s

		return o, nil

	case 0x9e: // 0x9e - SBC
		o := &SBC_A_HLPtr{}

		var s string

		s = "A"
		o.operand1 = s

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0x9f: // 0x9f - SBC
		o := &SBC_A_A{}

		var s string

		s = "A"
		o.operand1 = s

		s = "A"
		o.operand2 = s

		return o, nil

	case 0xa: // 0xa - LD
		o := &LD_A_BCDeref{}

		var s string

		s = "A"
		o.operand1 = s

		s = "(BC)"
		o.operand2 = s

		return o, nil

	case 0xa0: // 0xa0 - AND
		o := &AND_B{}

		var s string

		s = "B"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xa1: // 0xa1 - AND
		o := &AND_C{}

		var s string

		s = "C"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xa2: // 0xa2 - AND
		o := &AND_D{}

		var s string

		s = "D"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xa3: // 0xa3 - AND
		o := &AND_E{}

		var s string

		s = "E"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xa4: // 0xa4 - AND
		o := &AND_H{}

		var s string

		s = "H"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xa5: // 0xa5 - AND
		o := &AND_L{}

		var s string

		s = "L"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xa6: // 0xa6 - AND
		o := &AND_HLPtr{}

		var s string

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xa7: // 0xa7 - AND
		o := &AND_A{}

		var s string

		s = "A"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xa8: // 0xa8 - XOR
		o := &XOR_B{}

		var s string

		s = "B"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xa9: // 0xa9 - XOR
		o := &XOR_C{}

		var s string

		s = "C"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xaa: // 0xaa - XOR
		o := &XOR_D{}

		var s string

		s = "D"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xab: // 0xab - XOR
		o := &XOR_E{}

		var s string

		s = "E"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xac: // 0xac - XOR
		o := &XOR_H{}

		var s string

		s = "H"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xad: // 0xad - XOR
		o := &XOR_L{}

		var s string

		s = "L"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xae: // 0xae - XOR
		o := &XOR_HLPtr{}

		var s string

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xaf: // 0xaf - XOR
		o := &XOR_A{}

		var s string

		s = "A"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xb: // 0xb - DEC
		o := &DEC_BC{}

		var s string

		s = "BC"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xb0: // 0xb0 - OR
		o := &OR_B{}

		var s string

		s = "B"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xb1: // 0xb1 - OR
		o := &OR_C{}

		var s string

		s = "C"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xb2: // 0xb2 - OR
		o := &OR_D{}

		var s string

		s = "D"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xb3: // 0xb3 - OR
		o := &OR_E{}

		var s string

		s = "E"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xb4: // 0xb4 - OR
		o := &OR_H{}

		var s string

		s = "H"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xb5: // 0xb5 - OR
		o := &OR_L{}

		var s string

		s = "L"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xb6: // 0xb6 - OR
		o := &OR_HLPtr{}

		var s string

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xb7: // 0xb7 - OR
		o := &OR_A{}

		var s string

		s = "A"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xb8: // 0xb8 - CP
		o := &CP_B{}

		var s string

		s = "B"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xb9: // 0xb9 - CP
		o := &CP_C{}

		var s string

		s = "C"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xba: // 0xba - CP
		o := &CP_D{}

		var s string

		s = "D"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xbb: // 0xbb - CP
		o := &CP_E{}

		var s string

		s = "E"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xbc: // 0xbc - CP
		o := &CP_H{}

		var s string

		s = "H"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xbd: // 0xbd - CP
		o := &CP_L{}

		var s string

		s = "L"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xbe: // 0xbe - CP
		o := &CP_HLPtr{}

		var s string

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xbf: // 0xbf - CP
		o := &CP_A{}

		var s string

		s = "A"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xc: // 0xc - INC
		o := &INC_C{}

		var s string

		s = "C"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xc0: // 0xc0 - RET
		o := &RET_NZ{}

		var s string

		s = "NZ"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xc1: // 0xc1 - POP
		o := &POP_BC{}

		var s string

		s = "BC"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xc2: // 0xc2 - JP
		o := &JP_NZ_a16{}

		var s string

		s = "NZ"
		o.operand1 = s

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0xc3: // 0xc3 - JP
		o := &JP_a16{}

		var s string

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xc4: // 0xc4 - CALL
		o := &CALL_NZ_a16{}

		var s string

		s = "NZ"
		o.operand1 = s

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0xc5: // 0xc5 - PUSH
		o := &PUSH_BC{}

		var s string

		s = "BC"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xc6: // 0xc6 - ADD
		o := &ADD_A_d8{}

		var s string

		s = "A"
		o.operand1 = s

		s, err = readBytesAsString(data, 1)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0xc7: // 0xc7 - RST
		o := &RST_00H{}

		var s string

		s = "00H"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xc8: // 0xc8 - RET
		o := &RET_Z{}

		var s string

		s = "Z"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xc9: // 0xc9 - RET
		o := &RET_{}

		var s string

		s = ""
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xca: // 0xca - JP
		o := &JP_Z_a16{}

		var s string

		s = "Z"
		o.operand1 = s

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0xcb: // 0xcb - PREFIX
		o := &PREFIX_CB{}

		var s string

		s = "CB"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xcc: // 0xcc - CALL
		o := &CALL_Z_a16{}

		var s string

		s = "Z"
		o.operand1 = s

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0xcd: // 0xcd - CALL
		o := &CALL_a16{}

		var s string

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xce: // 0xce - ADC
		o := &ADC_A_d8{}

		var s string

		s = "A"
		o.operand1 = s

		s, err = readBytesAsString(data, 1)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0xcf: // 0xcf - RST
		o := &RST_08H{}

		var s string

		s = "08H"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xd: // 0xd - DEC
		o := &DEC_C{}

		var s string

		s = "C"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xd0: // 0xd0 - RET
		o := &RET_NC{}

		var s string

		s = "NC"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xd1: // 0xd1 - POP
		o := &POP_DE{}

		var s string

		s = "DE"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xd2: // 0xd2 - JP
		o := &JP_NC_a16{}

		var s string

		s = "NC"
		o.operand1 = s

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0xd4: // 0xd4 - CALL
		o := &CALL_NC_a16{}

		var s string

		s = "NC"
		o.operand1 = s

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0xd5: // 0xd5 - PUSH
		o := &PUSH_DE{}

		var s string

		s = "DE"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xd6: // 0xd6 - SUB
		o := &SUB_d8{}

		var s string

		s, err = readBytesAsString(data, 1)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xd7: // 0xd7 - RST
		o := &RST_10H{}

		var s string

		s = "10H"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xd8: // 0xd8 - RET
		o := &RET_C{}

		var s string

		s = "C"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xd9: // 0xd9 - RETI
		o := &RETI_{}

		var s string

		s = ""
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xda: // 0xda - JP
		o := &JP_C_a16{}

		var s string

		s = "C"
		o.operand1 = s

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0xdc: // 0xdc - CALL
		o := &CALL_C_a16{}

		var s string

		s = "C"
		o.operand1 = s

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0xde: // 0xde - SBC
		o := &SBC_A_d8{}

		var s string

		s = "A"
		o.operand1 = s

		s, err = readBytesAsString(data, 1)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0xdf: // 0xdf - RST
		o := &RST_18H{}

		var s string

		s = "18H"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xe: // 0xe - LD
		o := &LD_C_d8{}

		var s string

		s = "C"
		o.operand1 = s

		s, err = readBytesAsString(data, 1)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand2 = s

		return o, nil

	case 0xe0: // 0xe0 - LDH
		o := &LDH_a8Deref_A{}

		var s string

		s = "(a8)"
		o.operand1 = s

		s = "A"
		o.operand2 = s

		return o, nil

	case 0xe1: // 0xe1 - POP
		o := &POP_HL{}

		var s string

		s = "HL"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xe2: // 0xe2 - LD
		o := &LD_CDeref_A{}

		var s string

		s = "(C)"
		o.operand1 = s

		s = "A"
		o.operand2 = s

		return o, nil

	case 0xe5: // 0xe5 - PUSH
		o := &PUSH_HL{}

		var s string

		s = "HL"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xe6: // 0xe6 - AND
		o := &AND_d8{}

		var s string

		s, err = readBytesAsString(data, 1)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xe7: // 0xe7 - RST
		o := &RST_20H{}

		var s string

		s = "20H"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xe8: // 0xe8 - ADD
		o := &ADD_SP_r8{}

		var s string

		s = "SP"
		o.operand1 = s

		s = "r8"
		o.operand2 = s

		return o, nil

	case 0xe9: // 0xe9 - JP
		o := &JP_HLPtr{}

		var s string

		s, err = readBytesAsString(data, 2)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xea: // 0xea - LD
		o := &LD_a16Deref_A{}

		var s string

		s = "(a16)"
		o.operand1 = s

		s = "A"
		o.operand2 = s

		return o, nil

	case 0xee: // 0xee - XOR
		o := &XOR_d8{}

		var s string

		s, err = readBytesAsString(data, 1)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xef: // 0xef - RST
		o := &RST_28H{}

		var s string

		s = "28H"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xf: // 0xf - RRCA
		o := &RRCA_{}

		var s string

		s = ""
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xf0: // 0xf0 - LDH
		o := &LDH_A_a8Deref{}

		var s string

		s = "A"
		o.operand1 = s

		s = "(a8)"
		o.operand2 = s

		return o, nil

	case 0xf1: // 0xf1 - POP
		o := &POP_AF{}

		var s string

		s = "AF"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xf2: // 0xf2 - LD
		o := &LD_A_CDeref{}

		var s string

		s = "A"
		o.operand1 = s

		s = "(C)"
		o.operand2 = s

		return o, nil

	case 0xf3: // 0xf3 - DI
		o := &DI_{}

		var s string

		s = ""
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xf5: // 0xf5 - PUSH
		o := &PUSH_AF{}

		var s string

		s = "AF"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xf6: // 0xf6 - OR
		o := &OR_d8{}

		var s string

		s, err = readBytesAsString(data, 1)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xf7: // 0xf7 - RST
		o := &RST_30H{}

		var s string

		s = "30H"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xf8: // 0xf8 - LD
		o := &LD_HL_SP_plus_r8{}

		var s string

		s = "HL"
		o.operand1 = s

		s = "SP+r8"
		o.operand2 = s

		return o, nil

	case 0xf9: // 0xf9 - LD
		o := &LD_SP_HL{}

		var s string

		s = "SP"
		o.operand1 = s

		s = "HL"
		o.operand2 = s

		return o, nil

	case 0xfa: // 0xfa - LD
		o := &LD_A_a16Deref{}

		var s string

		s = "A"
		o.operand1 = s

		s = "(a16)"
		o.operand2 = s

		return o, nil

	case 0xfb: // 0xfb - EI
		o := &EI_{}

		var s string

		s = ""
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xfe: // 0xfe - CP
		o := &CP_d8{}

		var s string

		s, err = readBytesAsString(data, 1)
		if err != nil {
			return nil, fmt.Errorf("useful error message: %v", err)
		}

		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	case 0xff: // 0xff - RST
		o := &RST_38H{}

		var s string

		s = "38H"
		o.operand1 = s

		s = ""
		o.operand2 = s

		return o, nil

	default:
		return nil, fmt.Errorf("the proposed opcode (dec %d, hex %x) doesn't exist: %v", d[0], d[0], ErrNoOpCode)
	}
}
