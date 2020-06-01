package vm

import (
	"errors"
	"fmt"

	"github.com/vsinha/vm/internal/vm"
)

var ErrNoOpCode = errors.New("no opcode with that address exists")

// TODO this has problems because many of the opcodes have the same mnemonic,
// so we have to think about how to disambiguate them or if we even want to
type Op int

type OpCode interface {
	Execute(*vm.VM) 
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

type DEC_A struct {
	operand1 string
	operand2 string
}

func (o *DEC_A) Execute(v *vm.VM) {
}

func (o *DEC_A) String() string {
	return "DEC A"
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

type LD_A_L struct {
	operand1 string
	operand2 string
}

func (o *LD_A_L) Execute(v *vm.VM) {
}

func (o *LD_A_L) String() string {
	return "LD A L"
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

type LD_C_d8 struct {
	operand1 string
	operand2 string
}

func (o *LD_C_d8) Execute(v *vm.VM) {
}

func (o *LD_C_d8) String() string {
	return "LD C d8"
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

type DAA_ struct {
	operand1 string
	operand2 string
}

func (o *DAA_) Execute(v *vm.VM) {
}

func (o *DAA_) String() string {
	return "DAA"
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

type SUB_E struct {
	operand1 string
	operand2 string
}

func (o *SUB_E) Execute(v *vm.VM) {
}

func (o *SUB_E) String() string {
	return "SUB E"
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

type LD_E_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *LD_E_HLPtr) Execute(v *vm.VM) {
}

func (o *LD_E_HLPtr) String() string {
	return "LD E (HL)"
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

type ADD_HL_DE struct {
	operand1 string
	operand2 string
}

func (o *ADD_HL_DE) Execute(v *vm.VM) {
}

func (o *ADD_HL_DE) String() string {
	return "ADD HL DE"
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

type SBC_A_d8 struct {
	operand1 string
	operand2 string
}

func (o *SBC_A_d8) Execute(v *vm.VM) {
}

func (o *SBC_A_d8) String() string {
	return "SBC A d8"
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

type LD_D_A struct {
	operand1 string
	operand2 string
}

func (o *LD_D_A) Execute(v *vm.VM) {
}

func (o *LD_D_A) String() string {
	return "LD D A"
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

type LD_BCDeref_A struct {
	operand1 string
	operand2 string
}

func (o *LD_BCDeref_A) Execute(v *vm.VM) {
}

func (o *LD_BCDeref_A) String() string {
	return "LD (BC) A"
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

type LD_D_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *LD_D_HLPtr) Execute(v *vm.VM) {
}

func (o *LD_D_HLPtr) String() string {
	return "LD D (HL)"
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

type CP_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *CP_HLPtr) Execute(v *vm.VM) {
}

func (o *CP_HLPtr) String() string {
	return "CP (HL)"
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

type DEC_L struct {
	operand1 string
	operand2 string
}

func (o *DEC_L) Execute(v *vm.VM) {
}

func (o *DEC_L) String() string {
	return "DEC L"
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

type POP_BC struct {
	operand1 string
	operand2 string
}

func (o *POP_BC) Execute(v *vm.VM) {
}

func (o *POP_BC) String() string {
	return "POP BC"
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

type INC_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *INC_HLPtr) Execute(v *vm.VM) {
}

func (o *INC_HLPtr) String() string {
	return "INC (HL)"
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

type SBC_A_D struct {
	operand1 string
	operand2 string
}

func (o *SBC_A_D) Execute(v *vm.VM) {
}

func (o *SBC_A_D) String() string {
	return "SBC A D"
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

type SUB_C struct {
	operand1 string
	operand2 string
}

func (o *SUB_C) Execute(v *vm.VM) {
}

func (o *SUB_C) String() string {
	return "SUB C"
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

type JP_NZ_a16 struct {
	operand1 string
	operand2 string
}

func (o *JP_NZ_a16) Execute(v *vm.VM) {
}

func (o *JP_NZ_a16) String() string {
	return "JP NZ o.operand1"
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

type DEC_E struct {
	operand1 string
	operand2 string
}

func (o *DEC_E) Execute(v *vm.VM) {
}

func (o *DEC_E) String() string {
	return "DEC E"
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

type LD_L_D struct {
	operand1 string
	operand2 string
}

func (o *LD_L_D) Execute(v *vm.VM) {
}

func (o *LD_L_D) String() string {
	return "LD L D"
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

type ADD_A_A struct {
	operand1 string
	operand2 string
}

func (o *ADD_A_A) Execute(v *vm.VM) {
}

func (o *ADD_A_A) String() string {
	return "ADD A A"
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

type INC_H struct {
	operand1 string
	operand2 string
}

func (o *INC_H) Execute(v *vm.VM) {
}

func (o *INC_H) String() string {
	return "INC H"
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

type AND_L struct {
	operand1 string
	operand2 string
}

func (o *AND_L) Execute(v *vm.VM) {
}

func (o *AND_L) String() string {
	return "AND L"
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

type INC_HL struct {
	operand1 string
	operand2 string
}

func (o *INC_HL) Execute(v *vm.VM) {
}

func (o *INC_HL) String() string {
	return "INC HL"
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

type LD_L_B struct {
	operand1 string
	operand2 string
}

func (o *LD_L_B) Execute(v *vm.VM) {
}

func (o *LD_L_B) String() string {
	return "LD L B"
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

type ADC_A_B struct {
	operand1 string
	operand2 string
}

func (o *ADC_A_B) Execute(v *vm.VM) {
}

func (o *ADC_A_B) String() string {
	return "ADC A B"
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

type SBC_A_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *SBC_A_HLPtr) Execute(v *vm.VM) {
}

func (o *SBC_A_HLPtr) String() string {
	return "SBC A (HL)"
}

type LD_a16Deref_A struct {
	operand1 string
	operand2 string
}

func (o *LD_a16Deref_A) Execute(v *vm.VM) {
}

func (o *LD_a16Deref_A) String() string {
	return "LD ("+o.operand1+") A"
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

type LD_E_E struct {
	operand1 string
	operand2 string
}

func (o *LD_E_E) Execute(v *vm.VM) {
}

func (o *LD_E_E) String() string {
	return "LD E E"
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

type CP_C struct {
	operand1 string
	operand2 string
}

func (o *CP_C) Execute(v *vm.VM) {
}

func (o *CP_C) String() string {
	return "CP C"
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

type JR_NZ_r8 struct {
	operand1 string
	operand2 string
}

func (o *JR_NZ_r8) Execute(v *vm.VM) {
}

func (o *JR_NZ_r8) String() string {
	return "JR NZ r8"
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

type LD_HLPtrDec_A struct {
	operand1 string
	operand2 string
}

func (o *LD_HLPtrDec_A) Execute(v *vm.VM) {
}

func (o *LD_HLPtrDec_A) String() string {
	return "LD (HL-) A"
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

type LD_E_B struct {
	operand1 string
	operand2 string
}

func (o *LD_E_B) Execute(v *vm.VM) {
}

func (o *LD_E_B) String() string {
	return "LD E B"
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

type LD_A_CDeref struct {
	operand1 string
	operand2 string
}

func (o *LD_A_CDeref) Execute(v *vm.VM) {
}

func (o *LD_A_CDeref) String() string {
	return "LD A (C)"
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

type RRA_ struct {
	operand1 string
	operand2 string
}

func (o *RRA_) Execute(v *vm.VM) {
}

func (o *RRA_) String() string {
	return "RRA"
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

type LD_D_L struct {
	operand1 string
	operand2 string
}

func (o *LD_D_L) Execute(v *vm.VM) {
}

func (o *LD_D_L) String() string {
	return "LD D L"
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

type ADD_A_D struct {
	operand1 string
	operand2 string
}

func (o *ADD_A_D) Execute(v *vm.VM) {
}

func (o *ADD_A_D) String() string {
	return "ADD A D"
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

type LD_SP_HL struct {
	operand1 string
	operand2 string
}

func (o *LD_SP_HL) Execute(v *vm.VM) {
}

func (o *LD_SP_HL) String() string {
	return "LD SP HL"
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

type LD_DE_d16 struct {
	operand1 string
	operand2 string
}

func (o *LD_DE_d16) Execute(v *vm.VM) {
}

func (o *LD_DE_d16) String() string {
	return "LD DE o.operand1"
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

type XOR_D struct {
	operand1 string
	operand2 string
}

func (o *XOR_D) Execute(v *vm.VM) {
}

func (o *XOR_D) String() string {
	return "XOR D"
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

type JP_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *JP_HLPtr) Execute(v *vm.VM) {
}

func (o *JP_HLPtr) String() string {
	return "JP (HL)"
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

type OR_A struct {
	operand1 string
	operand2 string
}

func (o *OR_A) Execute(v *vm.VM) {
}

func (o *OR_A) String() string {
	return "OR A"
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

type JP_a16 struct {
	operand1 string
	operand2 string
}

func (o *JP_a16) Execute(v *vm.VM) {
}

func (o *JP_a16) String() string {
	return "JP o.operand1"
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

type LD_B_B struct {
	operand1 string
	operand2 string
}

func (o *LD_B_B) Execute(v *vm.VM) {
}

func (o *LD_B_B) String() string {
	return "LD B B"
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

type ADC_A_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *ADC_A_HLPtr) Execute(v *vm.VM) {
}

func (o *ADC_A_HLPtr) String() string {
	return "ADC A (HL)"
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

type XOR_A struct {
	operand1 string
	operand2 string
}

func (o *XOR_A) Execute(v *vm.VM) {
}

func (o *XOR_A) String() string {
	return "XOR A"
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

type LD_A_DEDeref struct {
	operand1 string
	operand2 string
}

func (o *LD_A_DEDeref) Execute(v *vm.VM) {
}

func (o *LD_A_DEDeref) String() string {
	return "LD A (DE)"
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

type LD_H_d8 struct {
	operand1 string
	operand2 string
}

func (o *LD_H_d8) Execute(v *vm.VM) {
}

func (o *LD_H_d8) String() string {
	return "LD H d8"
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

type OR_E struct {
	operand1 string
	operand2 string
}

func (o *OR_E) Execute(v *vm.VM) {
}

func (o *OR_E) String() string {
	return "OR E"
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

type LD_A_H struct {
	operand1 string
	operand2 string
}

func (o *LD_A_H) Execute(v *vm.VM) {
}

func (o *LD_A_H) String() string {
	return "LD A H"
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

type LD_a16Deref_SP struct {
	operand1 string
	operand2 string
}

func (o *LD_a16Deref_SP) Execute(v *vm.VM) {
}

func (o *LD_a16Deref_SP) String() string {
	return "LD ("+o.operand1+") SP"
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

type LD_D_B struct {
	operand1 string
	operand2 string
}

func (o *LD_D_B) Execute(v *vm.VM) {
}

func (o *LD_D_B) String() string {
	return "LD D B"
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

type AND_D struct {
	operand1 string
	operand2 string
}

func (o *AND_D) Execute(v *vm.VM) {
}

func (o *AND_D) String() string {
	return "AND D"
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

type LD_B_A struct {
	operand1 string
	operand2 string
}

func (o *LD_B_A) Execute(v *vm.VM) {
}

func (o *LD_B_A) String() string {
	return "LD B A"
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

type OR_C struct {
	operand1 string
	operand2 string
}

func (o *OR_C) Execute(v *vm.VM) {
}

func (o *OR_C) String() string {
	return "OR C"
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

type JR_r8 struct {
	operand1 string
	operand2 string
}

func (o *JR_r8) Execute(v *vm.VM) {
}

func (o *JR_r8) String() string {
	return "JR r8"
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

type AND_E struct {
	operand1 string
	operand2 string
}

func (o *AND_E) Execute(v *vm.VM) {
}

func (o *AND_E) String() string {
	return "AND E"
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

type INC_DE struct {
	operand1 string
	operand2 string
}

func (o *INC_DE) Execute(v *vm.VM) {
}

func (o *INC_DE) String() string {
	return "INC DE"
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

type ADC_A_C struct {
	operand1 string
	operand2 string
}

func (o *ADC_A_C) Execute(v *vm.VM) {
}

func (o *ADC_A_C) String() string {
	return "ADC A C"
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

type RRCA_ struct {
	operand1 string
	operand2 string
}

func (o *RRCA_) Execute(v *vm.VM) {
}

func (o *RRCA_) String() string {
	return "RRCA"
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

type SUB_L struct {
	operand1 string
	operand2 string
}

func (o *SUB_L) Execute(v *vm.VM) {
}

func (o *SUB_L) String() string {
	return "SUB L"
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

type AND_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *AND_HLPtr) Execute(v *vm.VM) {
}

func (o *AND_HLPtr) String() string {
	return "AND (HL)"
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

type LD_H_A struct {
	operand1 string
	operand2 string
}

func (o *LD_H_A) Execute(v *vm.VM) {
}

func (o *LD_H_A) String() string {
	return "LD H A"
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

type CALL_Z_a16 struct {
	operand1 string
	operand2 string
}

func (o *CALL_Z_a16) Execute(v *vm.VM) {
}

func (o *CALL_Z_a16) String() string {
	return "CALL Z o.operand1"
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

type RST_10H struct {
	operand1 string
	operand2 string
}

func (o *RST_10H) Execute(v *vm.VM) {
}

func (o *RST_10H) String() string {
	return "RST 10H"
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

type RLA_ struct {
	operand1 string
	operand2 string
}

func (o *RLA_) Execute(v *vm.VM) {
}

func (o *RLA_) String() string {
	return "RLA"
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

type DEC_SP struct {
	operand1 string
	operand2 string
}

func (o *DEC_SP) Execute(v *vm.VM) {
}

func (o *DEC_SP) String() string {
	return "DEC SP"
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

type SBC_A_A struct {
	operand1 string
	operand2 string
}

func (o *SBC_A_A) Execute(v *vm.VM) {
}

func (o *SBC_A_A) String() string {
	return "SBC A A"
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

type OR_H struct {
	operand1 string
	operand2 string
}

func (o *OR_H) Execute(v *vm.VM) {
}

func (o *OR_H) String() string {
	return "OR H"
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

type LD_CDeref_A struct {
	operand1 string
	operand2 string
}

func (o *LD_CDeref_A) Execute(v *vm.VM) {
}

func (o *LD_CDeref_A) String() string {
	return "LD (C) A"
}

type LD_A_a16Deref struct {
	operand1 string
	operand2 string
}

func (o *LD_A_a16Deref) Execute(v *vm.VM) {
}

func (o *LD_A_a16Deref) String() string {
	return "LD A ("+o.operand1+")"
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

type LDH_a8Deref_A struct {
	operand1 string
	operand2 string
}

func (o *LDH_a8Deref_A) Execute(v *vm.VM) {
}

func (o *LDH_a8Deref_A) String() string {
	return "LDH (a8) A"
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

type INC_A struct {
	operand1 string
	operand2 string
}

func (o *INC_A) Execute(v *vm.VM) {
}

func (o *INC_A) String() string {
	return "INC A"
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

type LD_H_L struct {
	operand1 string
	operand2 string
}

func (o *LD_H_L) Execute(v *vm.VM) {
}

func (o *LD_H_L) String() string {
	return "LD H L"
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

type CP_H struct {
	operand1 string
	operand2 string
}

func (o *CP_H) Execute(v *vm.VM) {
}

func (o *CP_H) String() string {
	return "CP H"
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

type OR_L struct {
	operand1 string
	operand2 string
}

func (o *OR_L) Execute(v *vm.VM) {
}

func (o *OR_L) String() string {
	return "OR L"
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

type LD_HL_SP_plus_r8 struct {
	operand1 string
	operand2 string
}

func (o *LD_HL_SP_plus_r8) Execute(v *vm.VM) {
}

func (o *LD_HL_SP_plus_r8) String() string {
	return "LD HL SP+r8"
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

type LD_L_A struct {
	operand1 string
	operand2 string
}

func (o *LD_L_A) Execute(v *vm.VM) {
}

func (o *LD_L_A) String() string {
	return "LD L A"
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

type JP_NC_a16 struct {
	operand1 string
	operand2 string
}

func (o *JP_NC_a16) Execute(v *vm.VM) {
}

func (o *JP_NC_a16) String() string {
	return "JP NC o.operand1"
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

type INC_L struct {
	operand1 string
	operand2 string
}

func (o *INC_L) Execute(v *vm.VM) {
}

func (o *INC_L) String() string {
	return "INC L"
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

type SBC_A_B struct {
	operand1 string
	operand2 string
}

func (o *SBC_A_B) Execute(v *vm.VM) {
}

func (o *SBC_A_B) String() string {
	return "SBC A B"
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

type CP_d8 struct {
	operand1 string
	operand2 string
}

func (o *CP_d8) Execute(v *vm.VM) {
}

func (o *CP_d8) String() string {
	return "CP d8"
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

type LD_B_H struct {
	operand1 string
	operand2 string
}

func (o *LD_B_H) Execute(v *vm.VM) {
}

func (o *LD_B_H) String() string {
	return "LD B H"
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

type CALL_C_a16 struct {
	operand1 string
	operand2 string
}

func (o *CALL_C_a16) Execute(v *vm.VM) {
}

func (o *CALL_C_a16) String() string {
	return "CALL C o.operand1"
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

type LD_L_L struct {
	operand1 string
	operand2 string
}

func (o *LD_L_L) Execute(v *vm.VM) {
}

func (o *LD_L_L) String() string {
	return "LD L L"
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

type CP_L struct {
	operand1 string
	operand2 string
}

func (o *CP_L) Execute(v *vm.VM) {
}

func (o *CP_L) String() string {
	return "CP L"
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

type HALT_ struct {
	operand1 string
	operand2 string
}

func (o *HALT_) Execute(v *vm.VM) {
}

func (o *HALT_) String() string {
	return "HALT"
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

type XOR_B struct {
	operand1 string
	operand2 string
}

func (o *XOR_B) Execute(v *vm.VM) {
}

func (o *XOR_B) String() string {
	return "XOR B"
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

type DEC_B struct {
	operand1 string
	operand2 string
}

func (o *DEC_B) Execute(v *vm.VM) {
}

func (o *DEC_B) String() string {
	return "DEC B"
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

type ADC_A_D struct {
	operand1 string
	operand2 string
}

func (o *ADC_A_D) Execute(v *vm.VM) {
}

func (o *ADC_A_D) String() string {
	return "ADC A D"
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

type LD_A_d8 struct {
	operand1 string
	operand2 string
}

func (o *LD_A_d8) Execute(v *vm.VM) {
}

func (o *LD_A_d8) String() string {
	return "LD A d8"
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

type RET_Z struct {
	operand1 string
	operand2 string
}

func (o *RET_Z) Execute(v *vm.VM) {
}

func (o *RET_Z) String() string {
	return "RET Z"
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

type INC_B struct {
	operand1 string
	operand2 string
}

func (o *INC_B) Execute(v *vm.VM) {
}

func (o *INC_B) String() string {
	return "INC B"
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

type AND_B struct {
	operand1 string
	operand2 string
}

func (o *AND_B) Execute(v *vm.VM) {
}

func (o *AND_B) String() string {
	return "AND B"
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

type DEC_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *DEC_HLPtr) Execute(v *vm.VM) {
}

func (o *DEC_HLPtr) String() string {
	return "DEC (HL)"
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

type SUB_A struct {
	operand1 string
	operand2 string
}

func (o *SUB_A) Execute(v *vm.VM) {
}

func (o *SUB_A) String() string {
	return "SUB A"
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

type LD_A_D struct {
	operand1 string
	operand2 string
}

func (o *LD_A_D) Execute(v *vm.VM) {
}

func (o *LD_A_D) String() string {
	return "LD A D"
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

type NOP_ struct {
	operand1 string
	operand2 string
}

func (o *NOP_) Execute(v *vm.VM) {
}

func (o *NOP_) String() string {
	return "NOP"
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

type LD_L_H struct {
	operand1 string
	operand2 string
}

func (o *LD_L_H) Execute(v *vm.VM) {
}

func (o *LD_L_H) String() string {
	return "LD L H"
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

type OR_D struct {
	operand1 string
	operand2 string
}

func (o *OR_D) Execute(v *vm.VM) {
}

func (o *OR_D) String() string {
	return "OR D"
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

type SUB_d8 struct {
	operand1 string
	operand2 string
}

func (o *SUB_d8) Execute(v *vm.VM) {
}

func (o *SUB_d8) String() string {
	return "SUB d8"
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

type PUSH_AF struct {
	operand1 string
	operand2 string
}

func (o *PUSH_AF) Execute(v *vm.VM) {
}

func (o *PUSH_AF) String() string {
	return "PUSH AF"
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

type INC_SP struct {
	operand1 string
	operand2 string
}

func (o *INC_SP) Execute(v *vm.VM) {
}

func (o *INC_SP) String() string {
	return "INC SP"
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

type ADD_A_H struct {
	operand1 string
	operand2 string
}

func (o *ADD_A_H) Execute(v *vm.VM) {
}

func (o *ADD_A_H) String() string {
	return "ADD A H"
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

type CALL_a16 struct {
	operand1 string
	operand2 string
}

func (o *CALL_a16) Execute(v *vm.VM) {
}

func (o *CALL_a16) String() string {
	return "CALL o.operand1"
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

type LD_D_E struct {
	operand1 string
	operand2 string
}

func (o *LD_D_E) Execute(v *vm.VM) {
}

func (o *LD_D_E) String() string {
	return "LD D E"
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

type ADC_A_d8 struct {
	operand1 string
	operand2 string
}

func (o *ADC_A_d8) Execute(v *vm.VM) {
}

func (o *ADC_A_d8) String() string {
	return "ADC A d8"
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

type LD_BC_d16 struct {
	operand1 string
	operand2 string
}

func (o *LD_BC_d16) Execute(v *vm.VM) {
}

func (o *LD_BC_d16) String() string {
	return "LD BC o.operand1"
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

type RET_ struct {
	operand1 string
	operand2 string
}

func (o *RET_) Execute(v *vm.VM) {
}

func (o *RET_) String() string {
	return "RET"
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

type INC_E struct {
	operand1 string
	operand2 string
}

func (o *INC_E) Execute(v *vm.VM) {
}

func (o *INC_E) String() string {
	return "INC E"
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

type XOR_L struct {
	operand1 string
	operand2 string
}

func (o *XOR_L) Execute(v *vm.VM) {
}

func (o *XOR_L) String() string {
	return "XOR L"
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

type POP_AF struct {
	operand1 string
	operand2 string
}

func (o *POP_AF) Execute(v *vm.VM) {
}

func (o *POP_AF) String() string {
	return "POP AF"
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

type AND_d8 struct {
	operand1 string
	operand2 string
}

func (o *AND_d8) Execute(v *vm.VM) {
}

func (o *AND_d8) String() string {
	return "AND d8"
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

type LD_C_E struct {
	operand1 string
	operand2 string
}

func (o *LD_C_E) Execute(v *vm.VM) {
}

func (o *LD_C_E) String() string {
	return "LD C E"
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

type SUB_H struct {
	operand1 string
	operand2 string
}

func (o *SUB_H) Execute(v *vm.VM) {
}

func (o *SUB_H) String() string {
	return "SUB H"
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

type LD_B_E struct {
	operand1 string
	operand2 string
}

func (o *LD_B_E) Execute(v *vm.VM) {
}

func (o *LD_B_E) String() string {
	return "LD B E"
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

type XOR_C struct {
	operand1 string
	operand2 string
}

func (o *XOR_C) Execute(v *vm.VM) {
}

func (o *XOR_C) String() string {
	return "XOR C"
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

type CP_A struct {
	operand1 string
	operand2 string
}

func (o *CP_A) Execute(v *vm.VM) {
}

func (o *CP_A) String() string {
	return "CP A"
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

type AND_A struct {
	operand1 string
	operand2 string
}

func (o *AND_A) Execute(v *vm.VM) {
}

func (o *AND_A) String() string {
	return "AND A"
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

type LD_D_D struct {
	operand1 string
	operand2 string
}

func (o *LD_D_D) Execute(v *vm.VM) {
}

func (o *LD_D_D) String() string {
	return "LD D D"
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

type SET_5_E struct {
	operand1 string
	operand2 string
}

func (o *SET_5_E) Execute(v *vm.VM) {
}

func (o *SET_5_E) String() string {
	return "SET 5 E"
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

type RES_5_H struct {
	operand1 string
	operand2 string
}

func (o *RES_5_H) Execute(v *vm.VM) {
}

func (o *RES_5_H) String() string {
	return "RES 5 H"
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

type BIT_2_C struct {
	operand1 string
	operand2 string
}

func (o *BIT_2_C) Execute(v *vm.VM) {
}

func (o *BIT_2_C) String() string {
	return "BIT 2 C"
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

type BIT_6_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *BIT_6_HLPtr) Execute(v *vm.VM) {
}

func (o *BIT_6_HLPtr) String() string {
	return "BIT 6 (HL)"
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

type RES_5_E struct {
	operand1 string
	operand2 string
}

func (o *RES_5_E) Execute(v *vm.VM) {
}

func (o *RES_5_E) String() string {
	return "RES 5 E"
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

type SWAP_C struct {
	operand1 string
	operand2 string
}

func (o *SWAP_C) Execute(v *vm.VM) {
}

func (o *SWAP_C) String() string {
	return "SWAP C"
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

type SET_4_D struct {
	operand1 string
	operand2 string
}

func (o *SET_4_D) Execute(v *vm.VM) {
}

func (o *SET_4_D) String() string {
	return "SET 4 D"
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

type SET_5_L struct {
	operand1 string
	operand2 string
}

func (o *SET_5_L) Execute(v *vm.VM) {
}

func (o *SET_5_L) String() string {
	return "SET 5 L"
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

type BIT_5_D struct {
	operand1 string
	operand2 string
}

func (o *BIT_5_D) Execute(v *vm.VM) {
}

func (o *BIT_5_D) String() string {
	return "BIT 5 D"
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

type SET_6_L struct {
	operand1 string
	operand2 string
}

func (o *SET_6_L) Execute(v *vm.VM) {
}

func (o *SET_6_L) String() string {
	return "SET 6 L"
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

type BIT_1_H struct {
	operand1 string
	operand2 string
}

func (o *BIT_1_H) Execute(v *vm.VM) {
}

func (o *BIT_1_H) String() string {
	return "BIT 1 H"
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

type RES_3_B struct {
	operand1 string
	operand2 string
}

func (o *RES_3_B) Execute(v *vm.VM) {
}

func (o *RES_3_B) String() string {
	return "RES 3 B"
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

type SWAP_H struct {
	operand1 string
	operand2 string
}

func (o *SWAP_H) Execute(v *vm.VM) {
}

func (o *SWAP_H) String() string {
	return "SWAP H"
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

type SRA_D struct {
	operand1 string
	operand2 string
}

func (o *SRA_D) Execute(v *vm.VM) {
}

func (o *SRA_D) String() string {
	return "SRA D"
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

type BIT_0_L struct {
	operand1 string
	operand2 string
}

func (o *BIT_0_L) Execute(v *vm.VM) {
}

func (o *BIT_0_L) String() string {
	return "BIT 0 L"
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

type SET_4_E struct {
	operand1 string
	operand2 string
}

func (o *SET_4_E) Execute(v *vm.VM) {
}

func (o *SET_4_E) String() string {
	return "SET 4 E"
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

type SLA_B struct {
	operand1 string
	operand2 string
}

func (o *SLA_B) Execute(v *vm.VM) {
}

func (o *SLA_B) String() string {
	return "SLA B"
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

type SET_7_D struct {
	operand1 string
	operand2 string
}

func (o *SET_7_D) Execute(v *vm.VM) {
}

func (o *SET_7_D) String() string {
	return "SET 7 D"
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

type RL_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *RL_HLPtr) Execute(v *vm.VM) {
}

func (o *RL_HLPtr) String() string {
	return "RL (HL)"
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

type SET_1_H struct {
	operand1 string
	operand2 string
}

func (o *SET_1_H) Execute(v *vm.VM) {
}

func (o *SET_1_H) String() string {
	return "SET 1 H"
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

type BIT_0_B struct {
	operand1 string
	operand2 string
}

func (o *BIT_0_B) Execute(v *vm.VM) {
}

func (o *BIT_0_B) String() string {
	return "BIT 0 B"
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

type BIT_2_B struct {
	operand1 string
	operand2 string
}

func (o *BIT_2_B) Execute(v *vm.VM) {
}

func (o *BIT_2_B) String() string {
	return "BIT 2 B"
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

type BIT_6_A struct {
	operand1 string
	operand2 string
}

func (o *BIT_6_A) Execute(v *vm.VM) {
}

func (o *BIT_6_A) String() string {
	return "BIT 6 A"
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

type RLC_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *RLC_HLPtr) Execute(v *vm.VM) {
}

func (o *RLC_HLPtr) String() string {
	return "RLC (HL)"
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

type BIT_2_D struct {
	operand1 string
	operand2 string
}

func (o *BIT_2_D) Execute(v *vm.VM) {
}

func (o *BIT_2_D) String() string {
	return "BIT 2 D"
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

type RES_3_C struct {
	operand1 string
	operand2 string
}

func (o *RES_3_C) Execute(v *vm.VM) {
}

func (o *RES_3_C) String() string {
	return "RES 3 C"
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

type RES_1_B struct {
	operand1 string
	operand2 string
}

func (o *RES_1_B) Execute(v *vm.VM) {
}

func (o *RES_1_B) String() string {
	return "RES 1 B"
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

type RES_2_E struct {
	operand1 string
	operand2 string
}

func (o *RES_2_E) Execute(v *vm.VM) {
}

func (o *RES_2_E) String() string {
	return "RES 2 E"
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

type RES_6_L struct {
	operand1 string
	operand2 string
}

func (o *RES_6_L) Execute(v *vm.VM) {
}

func (o *RES_6_L) String() string {
	return "RES 6 L"
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

type BIT_0_E struct {
	operand1 string
	operand2 string
}

func (o *BIT_0_E) Execute(v *vm.VM) {
}

func (o *BIT_0_E) String() string {
	return "BIT 0 E"
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

type SET_5_H struct {
	operand1 string
	operand2 string
}

func (o *SET_5_H) Execute(v *vm.VM) {
}

func (o *SET_5_H) String() string {
	return "SET 5 H"
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

type SET_3_E struct {
	operand1 string
	operand2 string
}

func (o *SET_3_E) Execute(v *vm.VM) {
}

func (o *SET_3_E) String() string {
	return "SET 3 E"
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

type RES_5_C struct {
	operand1 string
	operand2 string
}

func (o *RES_5_C) Execute(v *vm.VM) {
}

func (o *RES_5_C) String() string {
	return "RES 5 C"
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

type RES_0_L struct {
	operand1 string
	operand2 string
}

func (o *RES_0_L) Execute(v *vm.VM) {
}

func (o *RES_0_L) String() string {
	return "RES 0 L"
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

type RES_3_H struct {
	operand1 string
	operand2 string
}

func (o *RES_3_H) Execute(v *vm.VM) {
}

func (o *RES_3_H) String() string {
	return "RES 3 H"
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

type SWAP_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *SWAP_HLPtr) Execute(v *vm.VM) {
}

func (o *SWAP_HLPtr) String() string {
	return "SWAP (HL)"
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

type SWAP_A struct {
	operand1 string
	operand2 string
}

func (o *SWAP_A) Execute(v *vm.VM) {
}

func (o *SWAP_A) String() string {
	return "SWAP A"
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

type BIT_7_B struct {
	operand1 string
	operand2 string
}

func (o *BIT_7_B) Execute(v *vm.VM) {
}

func (o *BIT_7_B) String() string {
	return "BIT 7 B"
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

type SET_0_H struct {
	operand1 string
	operand2 string
}

func (o *SET_0_H) Execute(v *vm.VM) {
}

func (o *SET_0_H) String() string {
	return "SET 0 H"
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

type RLC_L struct {
	operand1 string
	operand2 string
}

func (o *RLC_L) Execute(v *vm.VM) {
}

func (o *RLC_L) String() string {
	return "RLC L"
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

type BIT_3_H struct {
	operand1 string
	operand2 string
}

func (o *BIT_3_H) Execute(v *vm.VM) {
}

func (o *BIT_3_H) String() string {
	return "BIT 3 H"
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

type RES_0_H struct {
	operand1 string
	operand2 string
}

func (o *RES_0_H) Execute(v *vm.VM) {
}

func (o *RES_0_H) String() string {
	return "RES 0 H"
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

type BIT_6_L struct {
	operand1 string
	operand2 string
}

func (o *BIT_6_L) Execute(v *vm.VM) {
}

func (o *BIT_6_L) String() string {
	return "BIT 6 L"
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

type RES_5_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *RES_5_HLPtr) Execute(v *vm.VM) {
}

func (o *RES_5_HLPtr) String() string {
	return "RES 5 (HL)"
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

type SET_3_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *SET_3_HLPtr) Execute(v *vm.VM) {
}

func (o *SET_3_HLPtr) String() string {
	return "SET 3 (HL)"
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

type SRL_L struct {
	operand1 string
	operand2 string
}

func (o *SRL_L) Execute(v *vm.VM) {
}

func (o *SRL_L) String() string {
	return "SRL L"
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

type SLA_A struct {
	operand1 string
	operand2 string
}

func (o *SLA_A) Execute(v *vm.VM) {
}

func (o *SLA_A) String() string {
	return "SLA A"
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

type SET_1_D struct {
	operand1 string
	operand2 string
}

func (o *SET_1_D) Execute(v *vm.VM) {
}

func (o *SET_1_D) String() string {
	return "SET 1 D"
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

type RES_2_L struct {
	operand1 string
	operand2 string
}

func (o *RES_2_L) Execute(v *vm.VM) {
}

func (o *RES_2_L) String() string {
	return "RES 2 L"
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

type SRL_D struct {
	operand1 string
	operand2 string
}

func (o *SRL_D) Execute(v *vm.VM) {
}

func (o *SRL_D) String() string {
	return "SRL D"
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

type RES_7_D struct {
	operand1 string
	operand2 string
}

func (o *RES_7_D) Execute(v *vm.VM) {
}

func (o *RES_7_D) String() string {
	return "RES 7 D"
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

type SRA_C struct {
	operand1 string
	operand2 string
}

func (o *SRA_C) Execute(v *vm.VM) {
}

func (o *SRA_C) String() string {
	return "SRA C"
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

type RES_0_E struct {
	operand1 string
	operand2 string
}

func (o *RES_0_E) Execute(v *vm.VM) {
}

func (o *RES_0_E) String() string {
	return "RES 0 E"
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

type RRC_C struct {
	operand1 string
	operand2 string
}

func (o *RRC_C) Execute(v *vm.VM) {
}

func (o *RRC_C) String() string {
	return "RRC C"
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

type SET_4_C struct {
	operand1 string
	operand2 string
}

func (o *SET_4_C) Execute(v *vm.VM) {
}

func (o *SET_4_C) String() string {
	return "SET 4 C"
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

type RES_1_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *RES_1_HLPtr) Execute(v *vm.VM) {
}

func (o *RES_1_HLPtr) String() string {
	return "RES 1 (HL)"
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

type SET_2_L struct {
	operand1 string
	operand2 string
}

func (o *SET_2_L) Execute(v *vm.VM) {
}

func (o *SET_2_L) String() string {
	return "SET 2 L"
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

type RES_7_H struct {
	operand1 string
	operand2 string
}

func (o *RES_7_H) Execute(v *vm.VM) {
}

func (o *RES_7_H) String() string {
	return "RES 7 H"
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

type BIT_5_L struct {
	operand1 string
	operand2 string
}

func (o *BIT_5_L) Execute(v *vm.VM) {
}

func (o *BIT_5_L) String() string {
	return "BIT 5 L"
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

type SET_7_L struct {
	operand1 string
	operand2 string
}

func (o *SET_7_L) Execute(v *vm.VM) {
}

func (o *SET_7_L) String() string {
	return "SET 7 L"
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

type SRL_B struct {
	operand1 string
	operand2 string
}

func (o *SRL_B) Execute(v *vm.VM) {
}

func (o *SRL_B) String() string {
	return "SRL B"
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

type RES_2_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *RES_2_HLPtr) Execute(v *vm.VM) {
}

func (o *RES_2_HLPtr) String() string {
	return "RES 2 (HL)"
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

type SRL_H struct {
	operand1 string
	operand2 string
}

func (o *SRL_H) Execute(v *vm.VM) {
}

func (o *SRL_H) String() string {
	return "SRL H"
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

type BIT_0_C struct {
	operand1 string
	operand2 string
}

func (o *BIT_0_C) Execute(v *vm.VM) {
}

func (o *BIT_0_C) String() string {
	return "BIT 0 C"
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

type BIT_5_H struct {
	operand1 string
	operand2 string
}

func (o *BIT_5_H) Execute(v *vm.VM) {
}

func (o *BIT_5_H) String() string {
	return "BIT 5 H"
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

type SET_3_L struct {
	operand1 string
	operand2 string
}

func (o *SET_3_L) Execute(v *vm.VM) {
}

func (o *SET_3_L) String() string {
	return "SET 3 L"
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

type SET_6_D struct {
	operand1 string
	operand2 string
}

func (o *SET_6_D) Execute(v *vm.VM) {
}

func (o *SET_6_D) String() string {
	return "SET 6 D"
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

type BIT_3_A struct {
	operand1 string
	operand2 string
}

func (o *BIT_3_A) Execute(v *vm.VM) {
}

func (o *BIT_3_A) String() string {
	return "BIT 3 A"
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

type BIT_4_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *BIT_4_HLPtr) Execute(v *vm.VM) {
}

func (o *BIT_4_HLPtr) String() string {
	return "BIT 4 (HL)"
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

type RES_2_C struct {
	operand1 string
	operand2 string
}

func (o *RES_2_C) Execute(v *vm.VM) {
}

func (o *RES_2_C) String() string {
	return "RES 2 C"
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

type RES_5_D struct {
	operand1 string
	operand2 string
}

func (o *RES_5_D) Execute(v *vm.VM) {
}

func (o *RES_5_D) String() string {
	return "RES 5 D"
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

type BIT_2_L struct {
	operand1 string
	operand2 string
}

func (o *BIT_2_L) Execute(v *vm.VM) {
}

func (o *BIT_2_L) String() string {
	return "BIT 2 L"
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

type RES_6_A struct {
	operand1 string
	operand2 string
}

func (o *RES_6_A) Execute(v *vm.VM) {
}

func (o *RES_6_A) String() string {
	return "RES 6 A"
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

type SLA_E struct {
	operand1 string
	operand2 string
}

func (o *SLA_E) Execute(v *vm.VM) {
}

func (o *SLA_E) String() string {
	return "SLA E"
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

type SET_1_L struct {
	operand1 string
	operand2 string
}

func (o *SET_1_L) Execute(v *vm.VM) {
}

func (o *SET_1_L) String() string {
	return "SET 1 L"
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

type RRC_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *RRC_HLPtr) Execute(v *vm.VM) {
}

func (o *RRC_HLPtr) String() string {
	return "RRC (HL)"
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

type SET_0_A struct {
	operand1 string
	operand2 string
}

func (o *SET_0_A) Execute(v *vm.VM) {
}

func (o *SET_0_A) String() string {
	return "SET 0 A"
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

type BIT_0_A struct {
	operand1 string
	operand2 string
}

func (o *BIT_0_A) Execute(v *vm.VM) {
}

func (o *BIT_0_A) String() string {
	return "BIT 0 A"
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

type RES_6_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *RES_6_HLPtr) Execute(v *vm.VM) {
}

func (o *RES_6_HLPtr) String() string {
	return "RES 6 (HL)"
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

type SET_0_D struct {
	operand1 string
	operand2 string
}

func (o *SET_0_D) Execute(v *vm.VM) {
}

func (o *SET_0_D) String() string {
	return "SET 0 D"
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

type BIT_2_A struct {
	operand1 string
	operand2 string
}

func (o *BIT_2_A) Execute(v *vm.VM) {
}

func (o *BIT_2_A) String() string {
	return "BIT 2 A"
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

type SWAP_L struct {
	operand1 string
	operand2 string
}

func (o *SWAP_L) Execute(v *vm.VM) {
}

func (o *SWAP_L) String() string {
	return "SWAP L"
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

type RES_0_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *RES_0_HLPtr) Execute(v *vm.VM) {
}

func (o *RES_0_HLPtr) String() string {
	return "RES 0 (HL)"
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

type RES_7_L struct {
	operand1 string
	operand2 string
}

func (o *RES_7_L) Execute(v *vm.VM) {
}

func (o *RES_7_L) String() string {
	return "RES 7 L"
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

type SET_7_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *SET_7_HLPtr) Execute(v *vm.VM) {
}

func (o *SET_7_HLPtr) String() string {
	return "SET 7 (HL)"
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

type BIT_7_L struct {
	operand1 string
	operand2 string
}

func (o *BIT_7_L) Execute(v *vm.VM) {
}

func (o *BIT_7_L) String() string {
	return "BIT 7 L"
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

type BIT_5_B struct {
	operand1 string
	operand2 string
}

func (o *BIT_5_B) Execute(v *vm.VM) {
}

func (o *BIT_5_B) String() string {
	return "BIT 5 B"
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

type SWAP_D struct {
	operand1 string
	operand2 string
}

func (o *SWAP_D) Execute(v *vm.VM) {
}

func (o *SWAP_D) String() string {
	return "SWAP D"
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

type RES_2_D struct {
	operand1 string
	operand2 string
}

func (o *RES_2_D) Execute(v *vm.VM) {
}

func (o *RES_2_D) String() string {
	return "RES 2 D"
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

type RR_L struct {
	operand1 string
	operand2 string
}

func (o *RR_L) Execute(v *vm.VM) {
}

func (o *RR_L) String() string {
	return "RR L"
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

type SRA_L struct {
	operand1 string
	operand2 string
}

func (o *SRA_L) Execute(v *vm.VM) {
}

func (o *SRA_L) String() string {
	return "SRA L"
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

type SET_1_B struct {
	operand1 string
	operand2 string
}

func (o *SET_1_B) Execute(v *vm.VM) {
}

func (o *SET_1_B) String() string {
	return "SET 1 B"
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

type RR_H struct {
	operand1 string
	operand2 string
}

func (o *RR_H) Execute(v *vm.VM) {
}

func (o *RR_H) String() string {
	return "RR H"
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

type BIT_7_C struct {
	operand1 string
	operand2 string
}

func (o *BIT_7_C) Execute(v *vm.VM) {
}

func (o *BIT_7_C) String() string {
	return "BIT 7 C"
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

type SET_1_A struct {
	operand1 string
	operand2 string
}

func (o *SET_1_A) Execute(v *vm.VM) {
}

func (o *SET_1_A) String() string {
	return "SET 1 A"
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

type RRC_H struct {
	operand1 string
	operand2 string
}

func (o *RRC_H) Execute(v *vm.VM) {
}

func (o *RRC_H) String() string {
	return "RRC H"
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

type BIT_6_D struct {
	operand1 string
	operand2 string
}

func (o *BIT_6_D) Execute(v *vm.VM) {
}

func (o *BIT_6_D) String() string {
	return "BIT 6 D"
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

type RES_4_A struct {
	operand1 string
	operand2 string
}

func (o *RES_4_A) Execute(v *vm.VM) {
}

func (o *RES_4_A) String() string {
	return "RES 4 A"
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

type SWAP_E struct {
	operand1 string
	operand2 string
}

func (o *SWAP_E) Execute(v *vm.VM) {
}

func (o *SWAP_E) String() string {
	return "SWAP E"
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

type SET_2_H struct {
	operand1 string
	operand2 string
}

func (o *SET_2_H) Execute(v *vm.VM) {
}

func (o *SET_2_H) String() string {
	return "SET 2 H"
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

type BIT_4_L struct {
	operand1 string
	operand2 string
}

func (o *BIT_4_L) Execute(v *vm.VM) {
}

func (o *BIT_4_L) String() string {
	return "BIT 4 L"
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

type RES_4_E struct {
	operand1 string
	operand2 string
}

func (o *RES_4_E) Execute(v *vm.VM) {
}

func (o *RES_4_E) String() string {
	return "RES 4 E"
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

type SET_6_C struct {
	operand1 string
	operand2 string
}

func (o *SET_6_C) Execute(v *vm.VM) {
}

func (o *SET_6_C) String() string {
	return "SET 6 C"
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

type RES_0_A struct {
	operand1 string
	operand2 string
}

func (o *RES_0_A) Execute(v *vm.VM) {
}

func (o *RES_0_A) String() string {
	return "RES 0 A"
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

type RES_1_C struct {
	operand1 string
	operand2 string
}

func (o *RES_1_C) Execute(v *vm.VM) {
}

func (o *RES_1_C) String() string {
	return "RES 1 C"
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

type SET_1_E struct {
	operand1 string
	operand2 string
}

func (o *SET_1_E) Execute(v *vm.VM) {
}

func (o *SET_1_E) String() string {
	return "SET 1 E"
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

type SRA_H struct {
	operand1 string
	operand2 string
}

func (o *SRA_H) Execute(v *vm.VM) {
}

func (o *SRA_H) String() string {
	return "SRA H"
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

type SET_7_C struct {
	operand1 string
	operand2 string
}

func (o *SET_7_C) Execute(v *vm.VM) {
}

func (o *SET_7_C) String() string {
	return "SET 7 C"
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

type SET_1_C struct {
	operand1 string
	operand2 string
}

func (o *SET_1_C) Execute(v *vm.VM) {
}

func (o *SET_1_C) String() string {
	return "SET 1 C"
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

type RRC_E struct {
	operand1 string
	operand2 string
}

func (o *RRC_E) Execute(v *vm.VM) {
}

func (o *RRC_E) String() string {
	return "RRC E"
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

type SET_4_L struct {
	operand1 string
	operand2 string
}

func (o *SET_4_L) Execute(v *vm.VM) {
}

func (o *SET_4_L) String() string {
	return "SET 4 L"
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

type BIT_6_E struct {
	operand1 string
	operand2 string
}

func (o *BIT_6_E) Execute(v *vm.VM) {
}

func (o *BIT_6_E) String() string {
	return "BIT 6 E"
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

type SET_6_B struct {
	operand1 string
	operand2 string
}

func (o *SET_6_B) Execute(v *vm.VM) {
}

func (o *SET_6_B) String() string {
	return "SET 6 B"
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

type RES_1_D struct {
	operand1 string
	operand2 string
}

func (o *RES_1_D) Execute(v *vm.VM) {
}

func (o *RES_1_D) String() string {
	return "RES 1 D"
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

type RL_E struct {
	operand1 string
	operand2 string
}

func (o *RL_E) Execute(v *vm.VM) {
}

func (o *RL_E) String() string {
	return "RL E"
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

type RES_4_B struct {
	operand1 string
	operand2 string
}

func (o *RES_4_B) Execute(v *vm.VM) {
}

func (o *RES_4_B) String() string {
	return "RES 4 B"
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

type SET_2_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *SET_2_HLPtr) Execute(v *vm.VM) {
}

func (o *SET_2_HLPtr) String() string {
	return "SET 2 (HL)"
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

type SET_5_B struct {
	operand1 string
	operand2 string
}

func (o *SET_5_B) Execute(v *vm.VM) {
}

func (o *SET_5_B) String() string {
	return "SET 5 B"
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

type SRL_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *SRL_HLPtr) Execute(v *vm.VM) {
}

func (o *SRL_HLPtr) String() string {
	return "SRL (HL)"
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

type SET_0_L struct {
	operand1 string
	operand2 string
}

func (o *SET_0_L) Execute(v *vm.VM) {
}

func (o *SET_0_L) String() string {
	return "SET 0 L"
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

type RES_3_E struct {
	operand1 string
	operand2 string
}

func (o *RES_3_E) Execute(v *vm.VM) {
}

func (o *RES_3_E) String() string {
	return "RES 3 E"
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

type SET_3_B struct {
	operand1 string
	operand2 string
}

func (o *SET_3_B) Execute(v *vm.VM) {
}

func (o *SET_3_B) String() string {
	return "SET 3 B"
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

type RLC_D struct {
	operand1 string
	operand2 string
}

func (o *RLC_D) Execute(v *vm.VM) {
}

func (o *RLC_D) String() string {
	return "RLC D"
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

type RES_3_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *RES_3_HLPtr) Execute(v *vm.VM) {
}

func (o *RES_3_HLPtr) String() string {
	return "RES 3 (HL)"
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

type RES_0_B struct {
	operand1 string
	operand2 string
}

func (o *RES_0_B) Execute(v *vm.VM) {
}

func (o *RES_0_B) String() string {
	return "RES 0 B"
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

type BIT_6_B struct {
	operand1 string
	operand2 string
}

func (o *BIT_6_B) Execute(v *vm.VM) {
}

func (o *BIT_6_B) String() string {
	return "BIT 6 B"
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

type SLA_HLPtr struct {
	operand1 string
	operand2 string
}

func (o *SLA_HLPtr) Execute(v *vm.VM) {
}

func (o *SLA_HLPtr) String() string {
	return "SLA (HL)"
}


// ReadOpCode returns an executable opcode by taking an io.Reader
// and reading a single instruction off it. If there is no more data
// returns undelying io.Reader's EOF error type.
func ReadOpCode(data io.Reader) (*OpCode, error) {
	var d [1]byte
	err, _ := data.Read(d)
	if err != nil {
		return nil, err
	}

	switch d[0] {
	
	case 0x21: // LD
		o := &LD_HL_d16{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand2 = d
		
		return o, nil
	
	case 0x3d: // DEC
		o := &DEC_A{}
		return o, nil
	
	case 0x79: // LD
		o := &LD_A_C{}
		return o, nil
	
	case 0x7d: // LD
		o := &LD_A_L{}
		return o, nil
	
	case 0xd1: // POP
		o := &POP_DE{}
		return o, nil
	
	case 0xe: // LD
		o := &LD_C_d8{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x25: // DEC
		o := &DEC_H{}
		return o, nil
	
	case 0x27: // DAA
		o := &DAA_{}
		return o, nil
	
	case 0x85: // ADD
		o := &ADD_A_L{}
		return o, nil
	
	case 0x93: // SUB
		o := &SUB_E{}
		return o, nil
	
	case 0xa4: // AND
		o := &AND_H{}
		return o, nil
	
	case 0x5e: // LD
		o := &LD_E_HLPtr{}
		return o, nil
	
	case 0x74: // LD
		o := &LD_HLPtr_H{}
		return o, nil
	
	case 0x19: // ADD
		o := &ADD_HL_DE{}
		return o, nil
	
	case 0xba: // CP
		o := &CP_D{}
		return o, nil
	
	case 0xde: // SBC
		o := &SBC_A_d8{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x46: // LD
		o := &LD_B_HLPtr{}
		return o, nil
	
	case 0x57: // LD
		o := &LD_D_A{}
		return o, nil
	
	case 0x6b: // LD
		o := &LD_L_E{}
		return o, nil
	
	case 0x2: // LD
		o := &LD_BCDeref_A{}
		return o, nil
	
	case 0xd: // DEC
		o := &DEC_C{}
		return o, nil
	
	case 0x56: // LD
		o := &LD_D_HLPtr{}
		return o, nil
	
	case 0x66: // LD
		o := &LD_H_HLPtr{}
		return o, nil
	
	case 0xbe: // CP
		o := &CP_HLPtr{}
		return o, nil
	
	case 0x1b: // DEC
		o := &DEC_DE{}
		return o, nil
	
	case 0x2d: // DEC
		o := &DEC_L{}
		return o, nil
	
	case 0x83: // ADD
		o := &ADD_A_E{}
		return o, nil
	
	case 0xc1: // POP
		o := &POP_BC{}
		return o, nil
	
	case 0xc4: // CALL
		o := &CALL_NZ_a16{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand2 = d
		
		return o, nil
	
	case 0x34: // INC
		o := &INC_HLPtr{}
		return o, nil
	
	case 0x90: // SUB
		o := &SUB_B{}
		return o, nil
	
	case 0x9a: // SBC
		o := &SBC_A_D{}
		return o, nil
	
	case 0xbb: // CP
		o := &CP_E{}
		return o, nil
	
	case 0x91: // SUB
		o := &SUB_C{}
		return o, nil
	
	case 0xf0: // LDH
		o := &LDH_A_a8Deref{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xc2: // JP
		o := &JP_NZ_a16{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand2 = d
		
		return o, nil
	
	case 0xf7: // RST
		o := &RST_30H{}
		return o, nil
	
	case 0x1d: // DEC
		o := &DEC_E{}
		return o, nil
	
	case 0x5c: // LD
		o := &LD_E_H{}
		return o, nil
	
	case 0x6a: // LD
		o := &LD_L_D{}
		return o, nil
	
	case 0x77: // LD
		o := &LD_HLPtr_A{}
		return o, nil
	
	case 0x87: // ADD
		o := &ADD_A_A{}
		return o, nil
	
	case 0xc0: // RET
		o := &RET_NZ{}
		return o, nil
	
	case 0x24: // INC
		o := &INC_H{}
		return o, nil
	
	case 0x72: // LD
		o := &LD_HLPtr_D{}
		return o, nil
	
	case 0xa5: // AND
		o := &AND_L{}
		return o, nil
	
	case 0xd0: // RET
		o := &RET_NC{}
		return o, nil
	
	case 0x23: // INC
		o := &INC_HL{}
		return o, nil
	
	case 0x49: // LD
		o := &LD_C_C{}
		return o, nil
	
	case 0x68: // LD
		o := &LD_L_B{}
		return o, nil
	
	case 0xd5: // PUSH
		o := &PUSH_DE{}
		return o, nil
	
	case 0x88: // ADC
		o := &ADC_A_B{}
		return o, nil
	
	case 0x99: // SBC
		o := &SBC_A_C{}
		return o, nil
	
	case 0x9e: // SBC
		o := &SBC_A_HLPtr{}
		return o, nil
	
	case 0xea: // LD
		o := &LD_a16Deref_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand2 = d
		
		return o, nil
	
	case 0x42: // LD
		o := &LD_B_D{}
		return o, nil
	
	case 0x5b: // LD
		o := &LD_E_E{}
		return o, nil
	
	case 0x9b: // SBC
		o := &SBC_A_E{}
		return o, nil
	
	case 0xb9: // CP
		o := &CP_C{}
		return o, nil
	
	case 0x7e: // LD
		o := &LD_A_HLPtr{}
		return o, nil
	
	case 0x20: // JR
		o := &JR_NZ_r8{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x2a: // LD
		o := &LD_A_HLPtrInc{}
		return o, nil
	
	case 0x32: // LD
		o := &LD_HLPtrDec_A{}
		return o, nil
	
	case 0x4a: // LD
		o := &LD_C_D{}
		return o, nil
	
	case 0x58: // LD
		o := &LD_E_B{}
		return o, nil
	
	case 0x5d: // LD
		o := &LD_E_L{}
		return o, nil
	
	case 0xf2: // LD
		o := &LD_A_CDeref{}
		return o, nil
	
	case 0xf6: // OR
		o := &OR_d8{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x1f: // RRA
		o := &RRA_{}
		return o, nil
	
	case 0x3f: // CCF
		o := &CCF_{}
		return o, nil
	
	case 0x55: // LD
		o := &LD_D_L{}
		return o, nil
	
	case 0x75: // LD
		o := &LD_HLPtr_L{}
		return o, nil
	
	case 0x82: // ADD
		o := &ADD_A_D{}
		return o, nil
	
	case 0x96: // SUB
		o := &SUB_HLPtr{}
		return o, nil
	
	case 0xf9: // LD
		o := &LD_SP_HL{}
		return o, nil
	
	case 0xc: // INC
		o := &INC_C{}
		return o, nil
	
	case 0x11: // LD
		o := &LD_DE_d16{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand2 = d
		
		return o, nil
	
	case 0x29: // ADD
		o := &ADD_HL_HL{}
		return o, nil
	
	case 0xaa: // XOR
		o := &XOR_D{}
		return o, nil
	
	case 0xcb: // PREFIX
		o := &PREFIX_CB{}
		return o, nil
	
	case 0xe9: // JP
		o := &JP_HLPtr{}
		return o, nil
	
	case 0x15: // DEC
		o := &DEC_D{}
		return o, nil
	
	case 0x62: // LD
		o := &LD_H_D{}
		return o, nil
	
	case 0x63: // LD
		o := &LD_H_E{}
		return o, nil
	
	case 0xb7: // OR
		o := &OR_A{}
		return o, nil
	
	case 0xd8: // RET
		o := &RET_C{}
		return o, nil
	
	case 0xc3: // JP
		o := &JP_a16{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand2 = d
		
		return o, nil
	
	case 0x2e: // LD
		o := &LD_L_d8{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x40: // LD
		o := &LD_B_B{}
		return o, nil
	
	case 0x73: // LD
		o := &LD_HLPtr_E{}
		return o, nil
	
	case 0x8e: // ADC
		o := &ADC_A_HLPtr{}
		return o, nil
	
	case 0xab: // XOR
		o := &XOR_E{}
		return o, nil
	
	case 0xaf: // XOR
		o := &XOR_A{}
		return o, nil
	
	case 0x14: // INC
		o := &INC_D{}
		return o, nil
	
	case 0x1a: // LD
		o := &LD_A_DEDeref{}
		return o, nil
	
	case 0x4e: // LD
		o := &LD_C_HLPtr{}
		return o, nil
	
	case 0x26: // LD
		o := &LD_H_d8{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xa1: // AND
		o := &AND_C{}
		return o, nil
	
	case 0xb3: // OR
		o := &OR_E{}
		return o, nil
	
	case 0x28: // JR
		o := &JR_Z_r8{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x7c: // LD
		o := &LD_A_H{}
		return o, nil
	
	case 0xf3: // DI
		o := &DI_{}
		return o, nil
	
	case 0x8: // LD
		o := &LD_a16Deref_SP{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand2 = d
		
		return o, nil
	
	case 0x41: // LD
		o := &LD_B_C{}
		return o, nil
	
	case 0x50: // LD
		o := &LD_D_B{}
		return o, nil
	
	case 0x7b: // LD
		o := &LD_A_E{}
		return o, nil
	
	case 0xa2: // AND
		o := &AND_D{}
		return o, nil
	
	case 0x3a: // LD
		o := &LD_A_HLPtrDec{}
		return o, nil
	
	case 0x47: // LD
		o := &LD_B_A{}
		return o, nil
	
	case 0x8b: // ADC
		o := &ADC_A_E{}
		return o, nil
	
	case 0xb1: // OR
		o := &OR_C{}
		return o, nil
	
	case 0xe5: // PUSH
		o := &PUSH_HL{}
		return o, nil
	
	case 0x18: // JR
		o := &JR_r8{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x7f: // LD
		o := &LD_A_A{}
		return o, nil
	
	case 0xa3: // AND
		o := &AND_E{}
		return o, nil
	
	case 0x9: // ADD
		o := &ADD_HL_BC{}
		return o, nil
	
	case 0x13: // INC
		o := &INC_DE{}
		return o, nil
	
	case 0x4c: // LD
		o := &LD_C_H{}
		return o, nil
	
	case 0x89: // ADC
		o := &ADC_A_C{}
		return o, nil
	
	case 0xd4: // CALL
		o := &CALL_NC_a16{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand2 = d
		
		return o, nil
	
	case 0xf: // RRCA
		o := &RRCA_{}
		return o, nil
	
	case 0x5f: // LD
		o := &LD_E_A{}
		return o, nil
	
	case 0x95: // SUB
		o := &SUB_L{}
		return o, nil
	
	case 0x9d: // SBC
		o := &SBC_A_L{}
		return o, nil
	
	case 0xa6: // AND
		o := &AND_HLPtr{}
		return o, nil
	
	case 0xae: // XOR
		o := &XOR_HLPtr{}
		return o, nil
	
	case 0x59: // LD
		o := &LD_E_C{}
		return o, nil
	
	case 0x5a: // LD
		o := &LD_E_D{}
		return o, nil
	
	case 0x67: // LD
		o := &LD_H_A{}
		return o, nil
	
	case 0xb0: // OR
		o := &OR_B{}
		return o, nil
	
	case 0xcc: // CALL
		o := &CALL_Z_a16{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand2 = d
		
		return o, nil
	
	case 0x81: // ADD
		o := &ADD_A_C{}
		return o, nil
	
	case 0xd7: // RST
		o := &RST_10H{}
		return o, nil
	
	case 0xe1: // POP
		o := &POP_HL{}
		return o, nil
	
	case 0x17: // RLA
		o := &RLA_{}
		return o, nil
	
	case 0x22: // LD
		o := &LD_HLPtrInc_A{}
		return o, nil
	
	case 0x3b: // DEC
		o := &DEC_SP{}
		return o, nil
	
	case 0x9c: // SBC
		o := &SBC_A_H{}
		return o, nil
	
	case 0x9f: // SBC
		o := &SBC_A_A{}
		return o, nil
	
	case 0x4d: // LD
		o := &LD_C_L{}
		return o, nil
	
	case 0xb4: // OR
		o := &OR_H{}
		return o, nil
	
	case 0xda: // JP
		o := &JP_C_a16{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand2 = d
		
		return o, nil
	
	case 0xe2: // LD
		o := &LD_CDeref_A{}
		return o, nil
	
	case 0xfa: // LD
		o := &LD_A_a16Deref{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand2 = d
		
		return o, nil
	
	case 0xcf: // RST
		o := &RST_08H{}
		return o, nil
	
	case 0xe0: // LDH
		o := &LDH_a8Deref_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x7: // RLCA
		o := &RLCA_{}
		return o, nil
	
	case 0x3c: // INC
		o := &INC_A{}
		return o, nil
	
	case 0x45: // LD
		o := &LD_B_L{}
		return o, nil
	
	case 0x65: // LD
		o := &LD_H_L{}
		return o, nil
	
	case 0x80: // ADD
		o := &ADD_A_B{}
		return o, nil
	
	case 0xbc: // CP
		o := &CP_H{}
		return o, nil
	
	case 0x70: // LD
		o := &LD_HLPtr_B{}
		return o, nil
	
	case 0xb5: // OR
		o := &OR_L{}
		return o, nil
	
	case 0xc7: // RST
		o := &RST_00H{}
		return o, nil
	
	case 0xf8: // LD
		o := &LD_HL_SP_plus_r8{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x39: // ADD
		o := &ADD_HL_SP{}
		return o, nil
	
	case 0x6f: // LD
		o := &LD_L_A{}
		return o, nil
	
	case 0x92: // SUB
		o := &SUB_D{}
		return o, nil
	
	case 0xd2: // JP
		o := &JP_NC_a16{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand2 = d
		
		return o, nil
	
	case 0xff: // RST
		o := &RST_38H{}
		return o, nil
	
	case 0x2c: // INC
		o := &INC_L{}
		return o, nil
	
	case 0x78: // LD
		o := &LD_A_B{}
		return o, nil
	
	case 0x98: // SBC
		o := &SBC_A_B{}
		return o, nil
	
	case 0xdf: // RST
		o := &RST_18H{}
		return o, nil
	
	case 0xfe: // CP
		o := &CP_d8{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x10: // STOP
		o := &STOP_0{}
		return o, nil
	
	case 0x44: // LD
		o := &LD_B_H{}
		return o, nil
	
	case 0x48: // LD
		o := &LD_C_B{}
		return o, nil
	
	case 0xdc: // CALL
		o := &CALL_C_a16{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand2 = d
		
		return o, nil
	
	case 0xe7: // RST
		o := &RST_20H{}
		return o, nil
	
	case 0xe8: // ADD
		o := &ADD_SP_r8{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x6d: // LD
		o := &LD_L_L{}
		return o, nil
	
	case 0x8f: // ADC
		o := &ADC_A_A{}
		return o, nil
	
	case 0xbd: // CP
		o := &CP_L{}
		return o, nil
	
	case 0x2f: // CPL
		o := &CPL_{}
		return o, nil
	
	case 0x76: // HALT
		o := &HALT_{}
		return o, nil
	
	case 0xa: // LD
		o := &LD_A_BCDeref{}
		return o, nil
	
	case 0xa8: // XOR
		o := &XOR_B{}
		return o, nil
	
	case 0x3: // INC
		o := &INC_BC{}
		return o, nil
	
	case 0x5: // DEC
		o := &DEC_B{}
		return o, nil
	
	case 0x36: // LD
		o := &LD_HLPtr_d8{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x8a: // ADC
		o := &ADC_A_D{}
		return o, nil
	
	case 0x6: // LD
		o := &LD_B_d8{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x3e: // LD
		o := &LD_A_d8{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x69: // LD
		o := &LD_L_C{}
		return o, nil
	
	case 0xc8: // RET
		o := &RET_Z{}
		return o, nil
	
	case 0xef: // RST
		o := &RST_28H{}
		return o, nil
	
	case 0x4: // INC
		o := &INC_B{}
		return o, nil
	
	case 0x71: // LD
		o := &LD_HLPtr_C{}
		return o, nil
	
	case 0xa0: // AND
		o := &AND_B{}
		return o, nil
	
	case 0xb: // DEC
		o := &DEC_BC{}
		return o, nil
	
	case 0x35: // DEC
		o := &DEC_HLPtr{}
		return o, nil
	
	case 0x8c: // ADC
		o := &ADC_A_H{}
		return o, nil
	
	case 0x97: // SUB
		o := &SUB_A{}
		return o, nil
	
	case 0x4f: // LD
		o := &LD_C_A{}
		return o, nil
	
	case 0x7a: // LD
		o := &LD_A_D{}
		return o, nil
	
	case 0xca: // JP
		o := &JP_Z_a16{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand2 = d
		
		return o, nil
	
	case 0x0: // NOP
		o := &NOP_{}
		return o, nil
	
	case 0x30: // JR
		o := &JR_NC_r8{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x6c: // LD
		o := &LD_L_H{}
		return o, nil
	
	case 0x8d: // ADC
		o := &ADC_A_L{}
		return o, nil
	
	case 0xb2: // OR
		o := &OR_D{}
		return o, nil
	
	case 0xb8: // CP
		o := &CP_B{}
		return o, nil
	
	case 0xd6: // SUB
		o := &SUB_d8{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xd9: // RETI
		o := &RETI_{}
		return o, nil
	
	case 0xf5: // PUSH
		o := &PUSH_AF{}
		return o, nil
	
	case 0x12: // LD
		o := &LD_DEDeref_A{}
		return o, nil
	
	case 0x33: // INC
		o := &INC_SP{}
		return o, nil
	
	case 0x64: // LD
		o := &LD_H_H{}
		return o, nil
	
	case 0x84: // ADD
		o := &ADD_A_H{}
		return o, nil
	
	case 0xc6: // ADD
		o := &ADD_A_d8{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xcd: // CALL
		o := &CALL_a16{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand2 = d
		
		return o, nil
	
	case 0x51: // LD
		o := &LD_D_C{}
		return o, nil
	
	case 0x53: // LD
		o := &LD_D_E{}
		return o, nil
	
	case 0x60: // LD
		o := &LD_H_B{}
		return o, nil
	
	case 0x61: // LD
		o := &LD_H_C{}
		return o, nil
	
	case 0xce: // ADC
		o := &ADC_A_d8{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xee: // XOR
		o := &XOR_d8{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x1: // LD
		o := &LD_BC_d16{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand2 = d
		
		return o, nil
	
	case 0x2b: // DEC
		o := &DEC_HL{}
		return o, nil
	
	case 0xc9: // RET
		o := &RET_{}
		return o, nil
	
	case 0xfb: // EI
		o := &EI_{}
		return o, nil
	
	case 0x1c: // INC
		o := &INC_E{}
		return o, nil
	
	case 0x37: // SCF
		o := &SCF_{}
		return o, nil
	
	case 0x38: // JR
		o := &JR_C_r8{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xad: // XOR
		o := &XOR_L{}
		return o, nil
	
	case 0xb6: // OR
		o := &OR_HLPtr{}
		return o, nil
	
	case 0xf1: // POP
		o := &POP_AF{}
		return o, nil
	
	case 0x6e: // LD
		o := &LD_L_HLPtr{}
		return o, nil
	
	case 0xe6: // AND
		o := &AND_d8{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x16: // LD
		o := &LD_D_d8{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x4b: // LD
		o := &LD_C_E{}
		return o, nil
	
	case 0x86: // ADD
		o := &ADD_A_HLPtr{}
		return o, nil
	
	case 0x94: // SUB
		o := &SUB_H{}
		return o, nil
	
	case 0x1e: // LD
		o := &LD_E_d8{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x43: // LD
		o := &LD_B_E{}
		return o, nil
	
	case 0x54: // LD
		o := &LD_D_H{}
		return o, nil
	
	case 0xa9: // XOR
		o := &XOR_C{}
		return o, nil
	
	case 0xac: // XOR
		o := &XOR_H{}
		return o, nil
	
	case 0xbf: // CP
		o := &CP_A{}
		return o, nil
	
	case 0x31: // LD
		o := &LD_SP_d16{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand2 = d
		
		return o, nil
	
	case 0xa7: // AND
		o := &AND_A{}
		return o, nil
	
	case 0xc5: // PUSH
		o := &PUSH_BC{}
		return o, nil
	
	case 0x52: // LD
		o := &LD_D_D{}
		return o, nil
	
	case 0xbe: // RES
		o := &RES_7_HLPtr{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xeb: // SET
		o := &SET_5_E{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x7b: // BIT
		o := &BIT_7_E{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xac: // RES
		o := &RES_5_H{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x48: // BIT
		o := &BIT_1_B{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x51: // BIT
		o := &BIT_2_C{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x6f: // BIT
		o := &BIT_5_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x76: // BIT
		o := &BIT_6_HLPtr{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xa1: // RES
		o := &RES_4_C{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xab: // RES
		o := &RES_5_E{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x15: // RL
		o := &RL_L{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x31: // SWAP
		o := &SWAP_C{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xb8: // RES
		o := &RES_7_B{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xe2: // SET
		o := &SET_4_D{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x82: // RES
		o := &RES_0_D{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xed: // SET
		o := &SET_5_L{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x17: // RL
		o := &RL_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x6a: // BIT
		o := &BIT_5_D{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xd3: // SET
		o := &SET_2_E{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xf5: // SET
		o := &SET_6_L{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x46: // BIT
		o := &BIT_0_HLPtr{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x4c: // BIT
		o := &BIT_1_H{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x6e: // BIT
		o := &BIT_5_HLPtr{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x98: // RES
		o := &RES_3_B{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x2b: // SRA
		o := &SRA_E{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x34: // SWAP
		o := &SWAP_H{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x71: // BIT
		o := &BIT_6_C{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x2a: // SRA
		o := &SRA_D{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x2e: // SRA
		o := &SRA_HLPtr{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x45: // BIT
		o := &BIT_0_L{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xb2: // RES
		o := &RES_6_D{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xe3: // SET
		o := &SET_4_E{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x1: // RLC
		o := &RLC_C{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x20: // SLA
		o := &SLA_B{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x7f: // BIT
		o := &BIT_7_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xfa: // SET
		o := &SET_7_D{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x0: // RLC
		o := &RLC_B{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x16: // RL
		o := &RL_HLPtr{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x4b: // BIT
		o := &BIT_1_E{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x4e: // BIT
		o := &BIT_1_HLPtr{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x4f: // BIT
		o := &BIT_1_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xcc: // SET
		o := &SET_1_H{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x3: // RLC
		o := &RLC_E{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x40: // BIT
		o := &BIT_0_B{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xad: // RES
		o := &RES_5_L{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x50: // BIT
		o := &BIT_2_B{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xa8: // RES
		o := &RES_5_B{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x77: // BIT
		o := &BIT_6_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xbf: // RES
		o := &RES_7_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x6: // RLC
		o := &RLC_HLPtr{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x3f: // SRL
		o := &SRL_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x52: // BIT
		o := &BIT_2_D{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x1b: // RR
		o := &RR_E{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x99: // RES
		o := &RES_3_C{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x5a: // BIT
		o := &BIT_3_D{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x88: // RES
		o := &RES_1_B{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x8d: // RES
		o := &RES_1_L{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x93: // RES
		o := &RES_2_E{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xa2: // RES
		o := &RES_4_D{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xb5: // RES
		o := &RES_6_L{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x1f: // RR
		o := &RR_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x43: // BIT
		o := &BIT_0_E{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xc1: // SET
		o := &SET_0_C{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xec: // SET
		o := &SET_5_H{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x3b: // SRL
		o := &SRL_E{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xdb: // SET
		o := &SET_3_E{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xee: // SET
		o := &SET_5_HLPtr{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xa9: // RES
		o := &RES_5_C{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x39: // SRL
		o := &SRL_C{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x85: // RES
		o := &RES_0_L{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x69: // BIT
		o := &BIT_5_C{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x9c: // RES
		o := &RES_3_H{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x28: // SRA
		o := &SRA_B{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x36: // SWAP
		o := &SWAP_HLPtr{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xf7: // SET
		o := &SET_6_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x37: // SWAP
		o := &SWAP_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x81: // RES
		o := &RES_0_C{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x78: // BIT
		o := &BIT_7_B{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x97: // RES
		o := &RES_2_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xc4: // SET
		o := &SET_0_H{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xd0: // SET
		o := &SET_2_B{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x5: // RLC
		o := &RLC_L{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x1a: // RR
		o := &RR_D{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x5c: // BIT
		o := &BIT_3_H{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x8c: // RES
		o := &RES_1_H{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x84: // RES
		o := &RES_0_H{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xb1: // RES
		o := &RES_6_C{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x75: // BIT
		o := &BIT_6_L{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x7c: // BIT
		o := &BIT_7_H{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xae: // RES
		o := &RES_5_HLPtr{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xc3: // SET
		o := &SET_0_E{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xde: // SET
		o := &SET_3_HLPtr{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x18: // RR
		o := &RR_B{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x3d: // SRL
		o := &SRL_L{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xfc: // SET
		o := &SET_7_H{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x27: // SLA
		o := &SLA_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xea: // SET
		o := &SET_5_D{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xca: // SET
		o := &SET_1_D{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x58: // BIT
		o := &BIT_3_B{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x95: // RES
		o := &RES_2_L{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xd: // RRC
		o := &RRC_L{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x3a: // SRL
		o := &SRL_D{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xa4: // RES
		o := &RES_4_H{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xba: // RES
		o := &RES_7_D{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xf3: // SET
		o := &SET_6_E{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x29: // SRA
		o := &SRA_C{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x30: // SWAP
		o := &SWAP_B{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x83: // RES
		o := &RES_0_E{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x9d: // RES
		o := &RES_3_L{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x9: // RRC
		o := &RRC_C{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x5d: // BIT
		o := &BIT_3_L{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xe1: // SET
		o := &SET_4_C{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xfb: // SET
		o := &SET_7_E{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x8e: // RES
		o := &RES_1_HLPtr{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xb0: // RES
		o := &RES_6_B{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xd5: // SET
		o := &SET_2_L{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x62: // BIT
		o := &BIT_4_D{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xbc: // RES
		o := &RES_7_H{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x61: // BIT
		o := &BIT_4_C{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x6d: // BIT
		o := &BIT_5_L{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xaf: // RES
		o := &RES_5_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xfd: // SET
		o := &SET_7_L{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x7: // RLC
		o := &RLC_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x38: // SRL
		o := &SRL_B{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x4d: // BIT
		o := &BIT_1_L{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x96: // RES
		o := &RES_2_HLPtr{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x19: // RR
		o := &RR_C{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x3c: // SRL
		o := &SRL_H{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xf4: // SET
		o := &SET_6_H{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x41: // BIT
		o := &BIT_0_C{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xda: // SET
		o := &SET_3_D{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x6c: // BIT
		o := &BIT_5_H{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xa6: // RES
		o := &RES_4_HLPtr{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xdd: // SET
		o := &SET_3_L{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xe0: // SET
		o := &SET_4_B{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xf2: // SET
		o := &SET_6_D{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x11: // RL
		o := &RL_C{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x5f: // BIT
		o := &BIT_3_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x5b: // BIT
		o := &BIT_3_E{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x66: // BIT
		o := &BIT_4_HLPtr{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x6b: // BIT
		o := &BIT_5_E{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x91: // RES
		o := &RES_2_C{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x94: // RES
		o := &RES_2_H{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xaa: // RES
		o := &RES_5_D{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x2f: // SRA
		o := &SRA_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x55: // BIT
		o := &BIT_2_L{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xb4: // RES
		o := &RES_6_H{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xb7: // RES
		o := &RES_6_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x1e: // RR
		o := &RR_HLPtr{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x23: // SLA
		o := &SLA_E{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x25: // SLA
		o := &SLA_L{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xcd: // SET
		o := &SET_1_L{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xd7: // SET
		o := &SET_2_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xe: // RRC
		o := &RRC_HLPtr{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x14: // RL
		o := &RL_H{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xc7: // SET
		o := &SET_0_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xa: // RRC
		o := &RRC_D{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x47: // BIT
		o := &BIT_0_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x7e: // BIT
		o := &BIT_7_HLPtr{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xb6: // RES
		o := &RES_6_HLPtr{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xc0: // SET
		o := &SET_0_B{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xc2: // SET
		o := &SET_0_D{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x24: // SLA
		o := &SLA_H{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x57: // BIT
		o := &BIT_2_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xdf: // SET
		o := &SET_3_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x35: // SWAP
		o := &SWAP_L{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xd1: // SET
		o := &SET_2_C{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x86: // RES
		o := &RES_0_HLPtr{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x90: // RES
		o := &RES_2_B{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xbd: // RES
		o := &RES_7_L{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xd9: // SET
		o := &SET_3_C{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xfe: // SET
		o := &SET_7_HLPtr{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x8: // RRC
		o := &RRC_B{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x7d: // BIT
		o := &BIT_7_L{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xdc: // SET
		o := &SET_3_H{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x68: // BIT
		o := &BIT_5_B{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x74: // BIT
		o := &BIT_6_H{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x32: // SWAP
		o := &SWAP_D{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x67: // BIT
		o := &BIT_4_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x92: // RES
		o := &RES_2_D{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xd2: // SET
		o := &SET_2_D{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x1d: // RR
		o := &RR_L{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x59: // BIT
		o := &BIT_3_C{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x2d: // SRA
		o := &SRA_L{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x64: // BIT
		o := &BIT_4_H{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xc8: // SET
		o := &SET_1_B{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xf8: // SET
		o := &SET_7_B{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x1c: // RR
		o := &RR_H{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x21: // SLA
		o := &SLA_C{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x79: // BIT
		o := &BIT_7_C{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xa5: // RES
		o := &RES_4_L{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xcf: // SET
		o := &SET_1_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x4: // RLC
		o := &RLC_H{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xc: // RRC
		o := &RRC_H{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x49: // BIT
		o := &BIT_1_C{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x72: // BIT
		o := &BIT_6_D{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x8f: // RES
		o := &RES_1_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xa7: // RES
		o := &RES_4_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x22: // SLA
		o := &SLA_D{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x33: // SWAP
		o := &SWAP_E{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xb9: // RES
		o := &RES_7_C{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xd4: // SET
		o := &SET_2_H{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x54: // BIT
		o := &BIT_2_H{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x65: // BIT
		o := &BIT_4_L{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x8b: // RES
		o := &RES_1_E{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xa3: // RES
		o := &RES_4_E{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xe7: // SET
		o := &SET_4_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xf1: // SET
		o := &SET_6_C{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x63: // BIT
		o := &BIT_4_E{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x87: // RES
		o := &RES_0_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x56: // BIT
		o := &BIT_2_HLPtr{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x89: // RES
		o := &RES_1_C{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x9f: // RES
		o := &RES_3_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xcb: // SET
		o := &SET_1_E{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x12: // RL
		o := &RL_D{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x2c: // SRA
		o := &SRA_H{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xe6: // SET
		o := &SET_4_HLPtr{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xf9: // SET
		o := &SET_7_C{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xf: // RRC
		o := &RRC_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xc9: // SET
		o := &SET_1_C{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xe9: // SET
		o := &SET_5_C{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xb: // RRC
		o := &RRC_E{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xc6: // SET
		o := &SET_0_HLPtr{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xe5: // SET
		o := &SET_4_L{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xff: // SET
		o := &SET_7_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x73: // BIT
		o := &BIT_6_E{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xbb: // RES
		o := &RES_7_E{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xf0: // SET
		o := &SET_6_B{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x5e: // BIT
		o := &BIT_3_HLPtr{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x8a: // RES
		o := &RES_1_D{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xb3: // RES
		o := &RES_6_E{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x13: // RL
		o := &RL_E{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x44: // BIT
		o := &BIT_0_H{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xa0: // RES
		o := &RES_4_B{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xce: // SET
		o := &SET_1_HLPtr{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xd6: // SET
		o := &SET_2_HLPtr{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xe4: // SET
		o := &SET_4_H{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xe8: // SET
		o := &SET_5_B{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xf6: // SET
		o := &SET_6_HLPtr{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x3e: // SRL
		o := &SRL_HLPtr{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x9a: // RES
		o := &RES_3_D{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xc5: // SET
		o := &SET_0_L{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x53: // BIT
		o := &BIT_2_E{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x9b: // RES
		o := &RES_3_E{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x60: // BIT
		o := &BIT_4_B{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xd8: // SET
		o := &SET_3_B{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0xef: // SET
		o := &SET_5_A{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x2: // RLC
		o := &RLC_D{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x4a: // BIT
		o := &BIT_1_D{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x9e: // RES
		o := &RES_3_HLPtr{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x7a: // BIT
		o := &BIT_7_D{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x80: // RES
		o := &RES_0_B{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x42: // BIT
		o := &BIT_0_D{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x70: // BIT
		o := &BIT_6_B{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x10: // RL
		o := &RL_B{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	case 0x26: // SLA
		o := &SLA_HLPtr{}
		err, _ = data.Read(d)
		if err != nil {
			return nil, err
		}
		o.Operand1 = d
		
		return o, nil
	
	default:
		return nil, fmt.Errorf("the proposed opcode (dec %d, hex %x) doesn't exist: %v", d[0], d[0], ErrNoOpCode)
	}
}