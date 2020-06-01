package vm

// TODO this has problems because many of the opcodes have the same mnemonic,
// so we have to think about how to disambiguate them or if we even want to
type Op int

type OpCode interface {
	Execute(*vm.VM) 
}


type RLCA_ struct {}

func (o *RLCA_) Execute(v *vm.VM) {
	
}

func (o *RLCA_) String() string {
	return "RLCA"
}

type RRA_ struct {}

func (o *RRA_) Execute(v *vm.VM) {
	
}

func (o *RRA_) String() string {
	return "RRA"
}

type SBC_A_A struct {}

func (o *SBC_A_A) Execute(v *vm.VM) {
	
}

func (o *SBC_A_A) String() string {
	return "SBC A A"
}

type CP_d8 struct {}

func (o *CP_d8) Execute(v *vm.VM) {
	
}

func (o *CP_d8) String() string {
	return "CP d8"
}

type INC_H struct {}

func (o *INC_H) Execute(v *vm.VM) {
	
}

func (o *INC_H) String() string {
	return "INC H"
}

type ADD_A_L struct {}

func (o *ADD_A_L) Execute(v *vm.VM) {
	
}

func (o *ADD_A_L) String() string {
	return "ADD A L"
}

type SBC_A_D struct {}

func (o *SBC_A_D) Execute(v *vm.VM) {
	
}

func (o *SBC_A_D) String() string {
	return "SBC A D"
}

type RST_10H struct {}

func (o *RST_10H) Execute(v *vm.VM) {
	
}

func (o *RST_10H) String() string {
	return "RST 10H"
}

type ADD_HL_DE struct {}

func (o *ADD_HL_DE) Execute(v *vm.VM) {
	
}

func (o *ADD_HL_DE) String() string {
	return "ADD HL DE"
}

type LD_C_B struct {}

func (o *LD_C_B) Execute(v *vm.VM) {
	v.C(v.B())
}

func (o *LD_C_B) String() string {
	return "LD C B"
}

type CALL_a16 struct {}

func (o *CALL_a16) Execute(v *vm.VM) {
	
}

func (o *CALL_a16) String() string {
	return "CALL a16"
}

type LDH_a8Deref_A struct {}

func (o *LDH_a8Deref_A) Execute(v *vm.VM) {
	
}

func (o *LDH_a8Deref_A) String() string {
	return "LDH (a8) A"
}

type RST_30H struct {}

func (o *RST_30H) Execute(v *vm.VM) {
	
}

func (o *RST_30H) String() string {
	return "RST 30H"
}

type LD_B_E struct {}

func (o *LD_B_E) Execute(v *vm.VM) {
	v.B(v.E())
}

func (o *LD_B_E) String() string {
	return "LD B E"
}

type ADC_A_D struct {}

func (o *ADC_A_D) Execute(v *vm.VM) {
	
}

func (o *ADC_A_D) String() string {
	return "ADC A D"
}

type EI_ struct {}

func (o *EI_) Execute(v *vm.VM) {
	
}

func (o *EI_) String() string {
	return "EI"
}

type STOP_0 struct {}

func (o *STOP_0) Execute(v *vm.VM) {
	
}

func (o *STOP_0) String() string {
	return "STOP 0"
}

type LD_L_d8 struct {}

func (o *LD_L_d8) Execute(v *vm.VM) {
	v.L(v.d8())
}

func (o *LD_L_d8) String() string {
	return "LD L d8"
}

type LD_D_B struct {}

func (o *LD_D_B) Execute(v *vm.VM) {
	v.D(v.B())
}

func (o *LD_D_B) String() string {
	return "LD D B"
}

type LD_D_A struct {}

func (o *LD_D_A) Execute(v *vm.VM) {
	v.D(v.A())
}

func (o *LD_D_A) String() string {
	return "LD D A"
}

type JP_HLPtr struct {}

func (o *JP_HLPtr) Execute(v *vm.VM) {
	
}

func (o *JP_HLPtr) String() string {
	return "JP (HL)"
}

type INC_A struct {}

func (o *INC_A) Execute(v *vm.VM) {
	
}

func (o *INC_A) String() string {
	return "INC A"
}

type LD_D_H struct {}

func (o *LD_D_H) Execute(v *vm.VM) {
	v.D(v.H())
}

func (o *LD_D_H) String() string {
	return "LD D H"
}

type LD_E_A struct {}

func (o *LD_E_A) Execute(v *vm.VM) {
	v.E(v.A())
}

func (o *LD_E_A) String() string {
	return "LD E A"
}

type LD_H_d8 struct {}

func (o *LD_H_d8) Execute(v *vm.VM) {
	v.H(v.d8())
}

func (o *LD_H_d8) String() string {
	return "LD H d8"
}

type ADD_A_B struct {}

func (o *ADD_A_B) Execute(v *vm.VM) {
	
}

func (o *ADD_A_B) String() string {
	return "ADD A B"
}

type ADC_A_A struct {}

func (o *ADC_A_A) Execute(v *vm.VM) {
	
}

func (o *ADC_A_A) String() string {
	return "ADC A A"
}

type LD_C_H struct {}

func (o *LD_C_H) Execute(v *vm.VM) {
	v.C(v.H())
}

func (o *LD_C_H) String() string {
	return "LD C H"
}

type LD_D_HLPtr struct {}

func (o *LD_D_HLPtr) Execute(v *vm.VM) {
	v.D(v.(HL)())
}

func (o *LD_D_HLPtr) String() string {
	return "LD D (HL)"
}

type LD_L_A struct {}

func (o *LD_L_A) Execute(v *vm.VM) {
	v.L(v.A())
}

func (o *LD_L_A) String() string {
	return "LD L A"
}

type LD_CDeref_A struct {}

func (o *LD_CDeref_A) Execute(v *vm.VM) {
	v.(C)(v.A())
}

func (o *LD_CDeref_A) String() string {
	return "LD (C) A"
}

type LD_D_d8 struct {}

func (o *LD_D_d8) Execute(v *vm.VM) {
	v.D(v.d8())
}

func (o *LD_D_d8) String() string {
	return "LD D d8"
}

type RLA_ struct {}

func (o *RLA_) Execute(v *vm.VM) {
	
}

func (o *RLA_) String() string {
	return "RLA"
}

type LD_L_E struct {}

func (o *LD_L_E) Execute(v *vm.VM) {
	v.L(v.E())
}

func (o *LD_L_E) String() string {
	return "LD L E"
}

type LD_A_C struct {}

func (o *LD_A_C) Execute(v *vm.VM) {
	v.A(v.C())
}

func (o *LD_A_C) String() string {
	return "LD A C"
}

type INC_BC struct {}

func (o *INC_BC) Execute(v *vm.VM) {
	
}

func (o *INC_BC) String() string {
	return "INC BC"
}

type OR_B struct {}

func (o *OR_B) Execute(v *vm.VM) {
	
}

func (o *OR_B) String() string {
	return "OR B"
}

type POP_AF struct {}

func (o *POP_AF) Execute(v *vm.VM) {
	
}

func (o *POP_AF) String() string {
	return "POP AF"
}

type LD_D_E struct {}

func (o *LD_D_E) Execute(v *vm.VM) {
	v.D(v.E())
}

func (o *LD_D_E) String() string {
	return "LD D E"
}

type JP_a16 struct {}

func (o *JP_a16) Execute(v *vm.VM) {
	
}

func (o *JP_a16) String() string {
	return "JP a16"
}

type LD_SP_d16 struct {}

func (o *LD_SP_d16) Execute(v *vm.VM) {
	v.SP(v.d16())
}

func (o *LD_SP_d16) String() string {
	return "LD SP d16"
}

type LD_H_E struct {}

func (o *LD_H_E) Execute(v *vm.VM) {
	v.H(v.E())
}

func (o *LD_H_E) String() string {
	return "LD H E"
}

type ADD_A_D struct {}

func (o *ADD_A_D) Execute(v *vm.VM) {
	
}

func (o *ADD_A_D) String() string {
	return "ADD A D"
}

type SBC_A_E struct {}

func (o *SBC_A_E) Execute(v *vm.VM) {
	
}

func (o *SBC_A_E) String() string {
	return "SBC A E"
}

type LD_A_H struct {}

func (o *LD_A_H) Execute(v *vm.VM) {
	v.A(v.H())
}

func (o *LD_A_H) String() string {
	return "LD A H"
}

type SBC_A_HLPtr struct {}

func (o *SBC_A_HLPtr) Execute(v *vm.VM) {
	
}

func (o *SBC_A_HLPtr) String() string {
	return "SBC A (HL)"
}

type RETI_ struct {}

func (o *RETI_) Execute(v *vm.VM) {
	
}

func (o *RETI_) String() string {
	return "RETI"
}

type CALL_C_a16 struct {}

func (o *CALL_C_a16) Execute(v *vm.VM) {
	
}

func (o *CALL_C_a16) String() string {
	return "CALL C a16"
}

type DEC_HL struct {}

func (o *DEC_HL) Execute(v *vm.VM) {
	
}

func (o *DEC_HL) String() string {
	return "DEC HL"
}

type OR_H struct {}

func (o *OR_H) Execute(v *vm.VM) {
	
}

func (o *OR_H) String() string {
	return "OR H"
}

type POP_BC struct {}

func (o *POP_BC) Execute(v *vm.VM) {
	
}

func (o *POP_BC) String() string {
	return "POP BC"
}

type DEC_SP struct {}

func (o *DEC_SP) Execute(v *vm.VM) {
	
}

func (o *DEC_SP) String() string {
	return "DEC SP"
}

type LD_B_H struct {}

func (o *LD_B_H) Execute(v *vm.VM) {
	v.B(v.H())
}

func (o *LD_B_H) String() string {
	return "LD B H"
}

type XOR_D struct {}

func (o *XOR_D) Execute(v *vm.VM) {
	
}

func (o *XOR_D) String() string {
	return "XOR D"
}

type OR_HLPtr struct {}

func (o *OR_HLPtr) Execute(v *vm.VM) {
	
}

func (o *OR_HLPtr) String() string {
	return "OR (HL)"
}

type INC_D struct {}

func (o *INC_D) Execute(v *vm.VM) {
	
}

func (o *INC_D) String() string {
	return "INC D"
}

type INC_L struct {}

func (o *INC_L) Execute(v *vm.VM) {
	
}

func (o *INC_L) String() string {
	return "INC L"
}

type ADC_A_E struct {}

func (o *ADC_A_E) Execute(v *vm.VM) {
	
}

func (o *ADC_A_E) String() string {
	return "ADC A E"
}

type SBC_A_B struct {}

func (o *SBC_A_B) Execute(v *vm.VM) {
	
}

func (o *SBC_A_B) String() string {
	return "SBC A B"
}

type AND_HLPtr struct {}

func (o *AND_HLPtr) Execute(v *vm.VM) {
	
}

func (o *AND_HLPtr) String() string {
	return "AND (HL)"
}

type ADD_HL_HL struct {}

func (o *ADD_HL_HL) Execute(v *vm.VM) {
	
}

func (o *ADD_HL_HL) String() string {
	return "ADD HL HL"
}

type LD_A_E struct {}

func (o *LD_A_E) Execute(v *vm.VM) {
	v.A(v.E())
}

func (o *LD_A_E) String() string {
	return "LD A E"
}

type PREFIX_CB struct {}

func (o *PREFIX_CB) Execute(v *vm.VM) {
	
}

func (o *PREFIX_CB) String() string {
	return "PREFIX CB"
}

type LD_A_CDeref struct {}

func (o *LD_A_CDeref) Execute(v *vm.VM) {
	v.A(v.(C)())
}

func (o *LD_A_CDeref) String() string {
	return "LD A (C)"
}

type INC_B struct {}

func (o *INC_B) Execute(v *vm.VM) {
	
}

func (o *INC_B) String() string {
	return "INC B"
}

type LD_H_L struct {}

func (o *LD_H_L) Execute(v *vm.VM) {
	v.H(v.L())
}

func (o *LD_H_L) String() string {
	return "LD H L"
}

type CP_A struct {}

func (o *CP_A) Execute(v *vm.VM) {
	
}

func (o *CP_A) String() string {
	return "CP A"
}

type LD_HLPtrInc_A struct {}

func (o *LD_HLPtrInc_A) Execute(v *vm.VM) {
	v.(HL+)(v.A())
}

func (o *LD_HLPtrInc_A) String() string {
	return "LD (HL+) A"
}

type LD_B_B struct {}

func (o *LD_B_B) Execute(v *vm.VM) {
	v.B(v.B())
}

func (o *LD_B_B) String() string {
	return "LD B B"
}

type RET_C struct {}

func (o *RET_C) Execute(v *vm.VM) {
	
}

func (o *RET_C) String() string {
	return "RET C"
}

type LD_BC_d16 struct {}

func (o *LD_BC_d16) Execute(v *vm.VM) {
	v.BC(v.d16())
}

func (o *LD_BC_d16) String() string {
	return "LD BC d16"
}

type INC_C struct {}

func (o *INC_C) Execute(v *vm.VM) {
	
}

func (o *INC_C) String() string {
	return "INC C"
}

type CPL_ struct {}

func (o *CPL_) Execute(v *vm.VM) {
	
}

func (o *CPL_) String() string {
	return "CPL"
}

type LD_A_D struct {}

func (o *LD_A_D) Execute(v *vm.VM) {
	v.A(v.D())
}

func (o *LD_A_D) String() string {
	return "LD A D"
}

type LD_A_A struct {}

func (o *LD_A_A) Execute(v *vm.VM) {
	v.A(v.A())
}

func (o *LD_A_A) String() string {
	return "LD A A"
}

type LD_E_B struct {}

func (o *LD_E_B) Execute(v *vm.VM) {
	v.E(v.B())
}

func (o *LD_E_B) String() string {
	return "LD E B"
}

type LD_HLPtr_H struct {}

func (o *LD_HLPtr_H) Execute(v *vm.VM) {
	v.(HL)(v.H())
}

func (o *LD_HLPtr_H) String() string {
	return "LD (HL) H"
}

type SUB_HLPtr struct {}

func (o *SUB_HLPtr) Execute(v *vm.VM) {
	
}

func (o *SUB_HLPtr) String() string {
	return "SUB (HL)"
}

type AND_D struct {}

func (o *AND_D) Execute(v *vm.VM) {
	
}

func (o *AND_D) String() string {
	return "AND D"
}

type OR_C struct {}

func (o *OR_C) Execute(v *vm.VM) {
	
}

func (o *OR_C) String() string {
	return "OR C"
}

type LD_A_BCDeref struct {}

func (o *LD_A_BCDeref) Execute(v *vm.VM) {
	v.A(v.(BC)())
}

func (o *LD_A_BCDeref) String() string {
	return "LD A (BC)"
}

type LD_HLPtr_D struct {}

func (o *LD_HLPtr_D) Execute(v *vm.VM) {
	v.(HL)(v.D())
}

func (o *LD_HLPtr_D) String() string {
	return "LD (HL) D"
}

type PUSH_AF struct {}

func (o *PUSH_AF) Execute(v *vm.VM) {
	
}

func (o *PUSH_AF) String() string {
	return "PUSH AF"
}

type LD_BCDeref_A struct {}

func (o *LD_BCDeref_A) Execute(v *vm.VM) {
	v.(BC)(v.A())
}

func (o *LD_BCDeref_A) String() string {
	return "LD (BC) A"
}

type ADD_HL_SP struct {}

func (o *ADD_HL_SP) Execute(v *vm.VM) {
	
}

func (o *ADD_HL_SP) String() string {
	return "ADD HL SP"
}

type SUB_C struct {}

func (o *SUB_C) Execute(v *vm.VM) {
	
}

func (o *SUB_C) String() string {
	return "SUB C"
}

type LD_SP_HL struct {}

func (o *LD_SP_HL) Execute(v *vm.VM) {
	v.SP(v.HL())
}

func (o *LD_SP_HL) String() string {
	return "LD SP HL"
}

type LD_A_d8 struct {}

func (o *LD_A_d8) Execute(v *vm.VM) {
	v.A(v.d8())
}

func (o *LD_A_d8) String() string {
	return "LD A d8"
}

type ADD_A_C struct {}

func (o *ADD_A_C) Execute(v *vm.VM) {
	
}

func (o *ADD_A_C) String() string {
	return "ADD A C"
}

type JP_NC_a16 struct {}

func (o *JP_NC_a16) Execute(v *vm.VM) {
	
}

func (o *JP_NC_a16) String() string {
	return "JP NC a16"
}

type LD_C_D struct {}

func (o *LD_C_D) Execute(v *vm.VM) {
	v.C(v.D())
}

func (o *LD_C_D) String() string {
	return "LD C D"
}

type SUB_A struct {}

func (o *SUB_A) Execute(v *vm.VM) {
	
}

func (o *SUB_A) String() string {
	return "SUB A"
}

type LD_E_C struct {}

func (o *LD_E_C) Execute(v *vm.VM) {
	v.E(v.C())
}

func (o *LD_E_C) String() string {
	return "LD E C"
}

type LD_E_D struct {}

func (o *LD_E_D) Execute(v *vm.VM) {
	v.E(v.D())
}

func (o *LD_E_D) String() string {
	return "LD E D"
}

type LD_H_C struct {}

func (o *LD_H_C) Execute(v *vm.VM) {
	v.H(v.C())
}

func (o *LD_H_C) String() string {
	return "LD H C"
}

type LD_L_H struct {}

func (o *LD_L_H) Execute(v *vm.VM) {
	v.L(v.H())
}

func (o *LD_L_H) String() string {
	return "LD L H"
}

type LD_HLPtr_C struct {}

func (o *LD_HLPtr_C) Execute(v *vm.VM) {
	v.(HL)(v.C())
}

func (o *LD_HLPtr_C) String() string {
	return "LD (HL) C"
}

type ADC_A_C struct {}

func (o *ADC_A_C) Execute(v *vm.VM) {
	
}

func (o *ADC_A_C) String() string {
	return "ADC A C"
}

type OR_L struct {}

func (o *OR_L) Execute(v *vm.VM) {
	
}

func (o *OR_L) String() string {
	return "OR L"
}

type RST_08H struct {}

func (o *RST_08H) Execute(v *vm.VM) {
	
}

func (o *RST_08H) String() string {
	return "RST 08H"
}

type DEC_D struct {}

func (o *DEC_D) Execute(v *vm.VM) {
	
}

func (o *DEC_D) String() string {
	return "DEC D"
}

type JR_Z_r8 struct {}

func (o *JR_Z_r8) Execute(v *vm.VM) {
	
}

func (o *JR_Z_r8) String() string {
	return "JR Z r8"
}

type LD_E_H struct {}

func (o *LD_E_H) Execute(v *vm.VM) {
	v.E(v.H())
}

func (o *LD_E_H) String() string {
	return "LD E H"
}

type SBC_A_H struct {}

func (o *SBC_A_H) Execute(v *vm.VM) {
	
}

func (o *SBC_A_H) String() string {
	return "SBC A H"
}

type OR_A struct {}

func (o *OR_A) Execute(v *vm.VM) {
	
}

func (o *OR_A) String() string {
	return "OR A"
}

type RET_Z struct {}

func (o *RET_Z) Execute(v *vm.VM) {
	
}

func (o *RET_Z) String() string {
	return "RET Z"
}

type ADD_HL_BC struct {}

func (o *ADD_HL_BC) Execute(v *vm.VM) {
	
}

func (o *ADD_HL_BC) String() string {
	return "ADD HL BC"
}

type JR_r8 struct {}

func (o *JR_r8) Execute(v *vm.VM) {
	
}

func (o *JR_r8) String() string {
	return "JR r8"
}

type LD_A_DEDeref struct {}

func (o *LD_A_DEDeref) Execute(v *vm.VM) {
	v.A(v.(DE)())
}

func (o *LD_A_DEDeref) String() string {
	return "LD A (DE)"
}

type INC_HLPtr struct {}

func (o *INC_HLPtr) Execute(v *vm.VM) {
	
}

func (o *INC_HLPtr) String() string {
	return "INC (HL)"
}

type INC_DE struct {}

func (o *INC_DE) Execute(v *vm.VM) {
	
}

func (o *INC_DE) String() string {
	return "INC DE"
}

type LD_B_C struct {}

func (o *LD_B_C) Execute(v *vm.VM) {
	v.B(v.C())
}

func (o *LD_B_C) String() string {
	return "LD B C"
}

type ADD_A_A struct {}

func (o *ADD_A_A) Execute(v *vm.VM) {
	
}

func (o *ADD_A_A) String() string {
	return "ADD A A"
}

type ADC_A_L struct {}

func (o *ADC_A_L) Execute(v *vm.VM) {
	
}

func (o *ADC_A_L) String() string {
	return "ADC A L"
}

type RST_38H struct {}

func (o *RST_38H) Execute(v *vm.VM) {
	
}

func (o *RST_38H) String() string {
	return "RST 38H"
}

type DEC_L struct {}

func (o *DEC_L) Execute(v *vm.VM) {
	
}

func (o *DEC_L) String() string {
	return "DEC L"
}

type XOR_L struct {}

func (o *XOR_L) Execute(v *vm.VM) {
	
}

func (o *XOR_L) String() string {
	return "XOR L"
}

type OR_D struct {}

func (o *OR_D) Execute(v *vm.VM) {
	
}

func (o *OR_D) String() string {
	return "OR D"
}

type RRCA_ struct {}

func (o *RRCA_) Execute(v *vm.VM) {
	
}

func (o *RRCA_) String() string {
	return "RRCA"
}

type LD_D_L struct {}

func (o *LD_D_L) Execute(v *vm.VM) {
	v.D(v.L())
}

func (o *LD_D_L) String() string {
	return "LD D L"
}

type XOR_C struct {}

func (o *XOR_C) Execute(v *vm.VM) {
	
}

func (o *XOR_C) String() string {
	return "XOR C"
}

type CP_C struct {}

func (o *CP_C) Execute(v *vm.VM) {
	
}

func (o *CP_C) String() string {
	return "CP C"
}

type CALL_Z_a16 struct {}

func (o *CALL_Z_a16) Execute(v *vm.VM) {
	
}

func (o *CALL_Z_a16) String() string {
	return "CALL Z a16"
}

type LD_B_A struct {}

func (o *LD_B_A) Execute(v *vm.VM) {
	v.B(v.A())
}

func (o *LD_B_A) String() string {
	return "LD B A"
}

type LD_A_B struct {}

func (o *LD_A_B) Execute(v *vm.VM) {
	v.A(v.B())
}

func (o *LD_A_B) String() string {
	return "LD A B"
}

type AND_A struct {}

func (o *AND_A) Execute(v *vm.VM) {
	
}

func (o *AND_A) String() string {
	return "AND A"
}

type CP_E struct {}

func (o *CP_E) Execute(v *vm.VM) {
	
}

func (o *CP_E) String() string {
	return "CP E"
}

type LD_B_d8 struct {}

func (o *LD_B_d8) Execute(v *vm.VM) {
	v.B(v.d8())
}

func (o *LD_B_d8) String() string {
	return "LD B d8"
}

type DEC_C struct {}

func (o *DEC_C) Execute(v *vm.VM) {
	
}

func (o *DEC_C) String() string {
	return "DEC C"
}

type AND_L struct {}

func (o *AND_L) Execute(v *vm.VM) {
	
}

func (o *AND_L) String() string {
	return "AND L"
}

type ADC_A_d8 struct {}

func (o *ADC_A_d8) Execute(v *vm.VM) {
	
}

func (o *ADC_A_d8) String() string {
	return "ADC A d8"
}

type XOR_B struct {}

func (o *XOR_B) Execute(v *vm.VM) {
	
}

func (o *XOR_B) String() string {
	return "XOR B"
}

type ADD_A_d8 struct {}

func (o *ADD_A_d8) Execute(v *vm.VM) {
	
}

func (o *ADD_A_d8) String() string {
	return "ADD A d8"
}

type HALT_ struct {}

func (o *HALT_) Execute(v *vm.VM) {
	
}

func (o *HALT_) String() string {
	return "HALT"
}

type CP_L struct {}

func (o *CP_L) Execute(v *vm.VM) {
	
}

func (o *CP_L) String() string {
	return "CP L"
}

type LD_C_E struct {}

func (o *LD_C_E) Execute(v *vm.VM) {
	v.C(v.E())
}

func (o *LD_C_E) String() string {
	return "LD C E"
}

type LD_C_L struct {}

func (o *LD_C_L) Execute(v *vm.VM) {
	v.C(v.L())
}

func (o *LD_C_L) String() string {
	return "LD C L"
}

type LD_HLPtr_B struct {}

func (o *LD_HLPtr_B) Execute(v *vm.VM) {
	v.(HL)(v.B())
}

func (o *LD_HLPtr_B) String() string {
	return "LD (HL) B"
}

type LD_a16Deref_A struct {}

func (o *LD_a16Deref_A) Execute(v *vm.VM) {
	v.(a16)(v.A())
}

func (o *LD_a16Deref_A) String() string {
	return "LD (a16) A"
}

type JR_C_r8 struct {}

func (o *JR_C_r8) Execute(v *vm.VM) {
	
}

func (o *JR_C_r8) String() string {
	return "JR C r8"
}

type LD_B_HLPtr struct {}

func (o *LD_B_HLPtr) Execute(v *vm.VM) {
	v.B(v.(HL)())
}

func (o *LD_B_HLPtr) String() string {
	return "LD B (HL)"
}

type LD_D_D struct {}

func (o *LD_D_D) Execute(v *vm.VM) {
	v.D(v.D())
}

func (o *LD_D_D) String() string {
	return "LD D D"
}

type LD_HLPtr_A struct {}

func (o *LD_HLPtr_A) Execute(v *vm.VM) {
	v.(HL)(v.A())
}

func (o *LD_HLPtr_A) String() string {
	return "LD (HL) A"
}

type RST_00H struct {}

func (o *RST_00H) Execute(v *vm.VM) {
	
}

func (o *RST_00H) String() string {
	return "RST 00H"
}

type JP_C_a16 struct {}

func (o *JP_C_a16) Execute(v *vm.VM) {
	
}

func (o *JP_C_a16) String() string {
	return "JP C a16"
}

type JR_NC_r8 struct {}

func (o *JR_NC_r8) Execute(v *vm.VM) {
	
}

func (o *JR_NC_r8) String() string {
	return "JR NC r8"
}

type LD_H_D struct {}

func (o *LD_H_D) Execute(v *vm.VM) {
	v.H(v.D())
}

func (o *LD_H_D) String() string {
	return "LD H D"
}

type LDH_A_a8Deref struct {}

func (o *LDH_A_a8Deref) Execute(v *vm.VM) {
	
}

func (o *LDH_A_a8Deref) String() string {
	return "LDH A (a8)"
}

type DEC_DE struct {}

func (o *DEC_DE) Execute(v *vm.VM) {
	
}

func (o *DEC_DE) String() string {
	return "DEC DE"
}

type LD_HLPtr_d8 struct {}

func (o *LD_HLPtr_d8) Execute(v *vm.VM) {
	v.(HL)(v.d8())
}

func (o *LD_HLPtr_d8) String() string {
	return "LD (HL) d8"
}

type LD_E_E struct {}

func (o *LD_E_E) Execute(v *vm.VM) {
	v.E(v.E())
}

func (o *LD_E_E) String() string {
	return "LD E E"
}

type XOR_d8 struct {}

func (o *XOR_d8) Execute(v *vm.VM) {
	
}

func (o *XOR_d8) String() string {
	return "XOR d8"
}

type OR_d8 struct {}

func (o *OR_d8) Execute(v *vm.VM) {
	
}

func (o *OR_d8) String() string {
	return "OR d8"
}

type LD_E_HLPtr struct {}

func (o *LD_E_HLPtr) Execute(v *vm.VM) {
	v.E(v.(HL)())
}

func (o *LD_E_HLPtr) String() string {
	return "LD E (HL)"
}

type LD_L_HLPtr struct {}

func (o *LD_L_HLPtr) Execute(v *vm.VM) {
	v.L(v.(HL)())
}

func (o *LD_L_HLPtr) String() string {
	return "LD L (HL)"
}

type SUB_H struct {}

func (o *SUB_H) Execute(v *vm.VM) {
	
}

func (o *SUB_H) String() string {
	return "SUB H"
}

type XOR_HLPtr struct {}

func (o *XOR_HLPtr) Execute(v *vm.VM) {
	
}

func (o *XOR_HLPtr) String() string {
	return "XOR (HL)"
}

type POP_HL struct {}

func (o *POP_HL) Execute(v *vm.VM) {
	
}

func (o *POP_HL) String() string {
	return "POP HL"
}

type LD_A_HLPtrInc struct {}

func (o *LD_A_HLPtrInc) Execute(v *vm.VM) {
	v.A(v.(HL+)())
}

func (o *LD_A_HLPtrInc) String() string {
	return "LD A (HL+)"
}

type LD_C_HLPtr struct {}

func (o *LD_C_HLPtr) Execute(v *vm.VM) {
	v.C(v.(HL)())
}

func (o *LD_C_HLPtr) String() string {
	return "LD C (HL)"
}

type LD_L_B struct {}

func (o *LD_L_B) Execute(v *vm.VM) {
	v.L(v.B())
}

func (o *LD_L_B) String() string {
	return "LD L B"
}

type SUB_B struct {}

func (o *SUB_B) Execute(v *vm.VM) {
	
}

func (o *SUB_B) String() string {
	return "SUB B"
}

type JP_NZ_a16 struct {}

func (o *JP_NZ_a16) Execute(v *vm.VM) {
	
}

func (o *JP_NZ_a16) String() string {
	return "JP NZ a16"
}

type JP_Z_a16 struct {}

func (o *JP_Z_a16) Execute(v *vm.VM) {
	
}

func (o *JP_Z_a16) String() string {
	return "JP Z a16"
}

type PUSH_DE struct {}

func (o *PUSH_DE) Execute(v *vm.VM) {
	
}

func (o *PUSH_DE) String() string {
	return "PUSH DE"
}

type RST_18H struct {}

func (o *RST_18H) Execute(v *vm.VM) {
	
}

func (o *RST_18H) String() string {
	return "RST 18H"
}

type DEC_E struct {}

func (o *DEC_E) Execute(v *vm.VM) {
	
}

func (o *DEC_E) String() string {
	return "DEC E"
}

type LD_A_L struct {}

func (o *LD_A_L) Execute(v *vm.VM) {
	v.A(v.L())
}

func (o *LD_A_L) String() string {
	return "LD A L"
}

type ADC_A_H struct {}

func (o *ADC_A_H) Execute(v *vm.VM) {
	
}

func (o *ADC_A_H) String() string {
	return "ADC A H"
}

type RET_NZ struct {}

func (o *RET_NZ) Execute(v *vm.VM) {
	
}

func (o *RET_NZ) String() string {
	return "RET NZ"
}

type LD_DE_d16 struct {}

func (o *LD_DE_d16) Execute(v *vm.VM) {
	v.DE(v.d16())
}

func (o *LD_DE_d16) String() string {
	return "LD DE d16"
}

type ADD_A_HLPtr struct {}

func (o *ADD_A_HLPtr) Execute(v *vm.VM) {
	
}

func (o *ADD_A_HLPtr) String() string {
	return "ADD A (HL)"
}

type CP_B struct {}

func (o *CP_B) Execute(v *vm.VM) {
	
}

func (o *CP_B) String() string {
	return "CP B"
}

type CP_D struct {}

func (o *CP_D) Execute(v *vm.VM) {
	
}

func (o *CP_D) String() string {
	return "CP D"
}

type POP_DE struct {}

func (o *POP_DE) Execute(v *vm.VM) {
	
}

func (o *POP_DE) String() string {
	return "POP DE"
}

type LD_HL_SP+r8 struct {}

func (o *LD_HL_SP+r8) Execute(v *vm.VM) {
	v.HL(v.SP+r8())
}

func (o *LD_HL_SP+r8) String() string {
	return "LD HL SP+r8"
}

type NOP_ struct {}

func (o *NOP_) Execute(v *vm.VM) {
	
}

func (o *NOP_) String() string {
	return "NOP"
}

type DAA_ struct {}

func (o *DAA_) Execute(v *vm.VM) {
	
}

func (o *DAA_) String() string {
	return "DAA"
}

type LD_HLPtrDec_A struct {}

func (o *LD_HLPtrDec_A) Execute(v *vm.VM) {
	v.(HL-)(v.A())
}

func (o *LD_HLPtrDec_A) String() string {
	return "LD (HL-) A"
}

type LD_D_C struct {}

func (o *LD_D_C) Execute(v *vm.VM) {
	v.D(v.C())
}

func (o *LD_D_C) String() string {
	return "LD D C"
}

type ADC_A_HLPtr struct {}

func (o *ADC_A_HLPtr) Execute(v *vm.VM) {
	
}

func (o *ADC_A_HLPtr) String() string {
	return "ADC A (HL)"
}

type RET_ struct {}

func (o *RET_) Execute(v *vm.VM) {
	
}

func (o *RET_) String() string {
	return "RET"
}

type LD_B_L struct {}

func (o *LD_B_L) Execute(v *vm.VM) {
	v.B(v.L())
}

func (o *LD_B_L) String() string {
	return "LD B L"
}

type OR_E struct {}

func (o *OR_E) Execute(v *vm.VM) {
	
}

func (o *OR_E) String() string {
	return "OR E"
}

type CALL_NC_a16 struct {}

func (o *CALL_NC_a16) Execute(v *vm.VM) {
	
}

func (o *CALL_NC_a16) String() string {
	return "CALL NC a16"
}

type LD_H_B struct {}

func (o *LD_H_B) Execute(v *vm.VM) {
	v.H(v.B())
}

func (o *LD_H_B) String() string {
	return "LD H B"
}

type LD_H_HLPtr struct {}

func (o *LD_H_HLPtr) Execute(v *vm.VM) {
	v.H(v.(HL)())
}

func (o *LD_H_HLPtr) String() string {
	return "LD H (HL)"
}

type SBC_A_d8 struct {}

func (o *SBC_A_d8) Execute(v *vm.VM) {
	
}

func (o *SBC_A_d8) String() string {
	return "SBC A d8"
}

type LD_C_d8 struct {}

func (o *LD_C_d8) Execute(v *vm.VM) {
	v.C(v.d8())
}

func (o *LD_C_d8) String() string {
	return "LD C d8"
}

type INC_HL struct {}

func (o *INC_HL) Execute(v *vm.VM) {
	
}

func (o *INC_HL) String() string {
	return "INC HL"
}

type LD_HLPtr_E struct {}

func (o *LD_HLPtr_E) Execute(v *vm.VM) {
	v.(HL)(v.E())
}

func (o *LD_HLPtr_E) String() string {
	return "LD (HL) E"
}

type LD_DEDeref_A struct {}

func (o *LD_DEDeref_A) Execute(v *vm.VM) {
	v.(DE)(v.A())
}

func (o *LD_DEDeref_A) String() string {
	return "LD (DE) A"
}

type LD_A_HLPtrDec struct {}

func (o *LD_A_HLPtrDec) Execute(v *vm.VM) {
	v.A(v.(HL-)())
}

func (o *LD_A_HLPtrDec) String() string {
	return "LD A (HL-)"
}

type LD_C_A struct {}

func (o *LD_C_A) Execute(v *vm.VM) {
	v.C(v.A())
}

func (o *LD_C_A) String() string {
	return "LD C A"
}

type SUB_L struct {}

func (o *SUB_L) Execute(v *vm.VM) {
	
}

func (o *SUB_L) String() string {
	return "SUB L"
}

type AND_E struct {}

func (o *AND_E) Execute(v *vm.VM) {
	
}

func (o *AND_E) String() string {
	return "AND E"
}

type AND_H struct {}

func (o *AND_H) Execute(v *vm.VM) {
	
}

func (o *AND_H) String() string {
	return "AND H"
}

type XOR_H struct {}

func (o *XOR_H) Execute(v *vm.VM) {
	
}

func (o *XOR_H) String() string {
	return "XOR H"
}

type RET_NC struct {}

func (o *RET_NC) Execute(v *vm.VM) {
	
}

func (o *RET_NC) String() string {
	return "RET NC"
}

type XOR_E struct {}

func (o *XOR_E) Execute(v *vm.VM) {
	
}

func (o *XOR_E) String() string {
	return "XOR E"
}

type CP_HLPtr struct {}

func (o *CP_HLPtr) Execute(v *vm.VM) {
	
}

func (o *CP_HLPtr) String() string {
	return "CP (HL)"
}

type RST_28H struct {}

func (o *RST_28H) Execute(v *vm.VM) {
	
}

func (o *RST_28H) String() string {
	return "RST 28H"
}

type LD_A_a16Deref struct {}

func (o *LD_A_a16Deref) Execute(v *vm.VM) {
	v.A(v.(a16)())
}

func (o *LD_A_a16Deref) String() string {
	return "LD A (a16)"
}

type DEC_B struct {}

func (o *DEC_B) Execute(v *vm.VM) {
	
}

func (o *DEC_B) String() string {
	return "DEC B"
}

type DEC_A struct {}

func (o *DEC_A) Execute(v *vm.VM) {
	
}

func (o *DEC_A) String() string {
	return "DEC A"
}

type LD_E_L struct {}

func (o *LD_E_L) Execute(v *vm.VM) {
	v.E(v.L())
}

func (o *LD_E_L) String() string {
	return "LD E L"
}

type LD_L_L struct {}

func (o *LD_L_L) Execute(v *vm.VM) {
	v.L(v.L())
}

func (o *LD_L_L) String() string {
	return "LD L L"
}

type XOR_A struct {}

func (o *XOR_A) Execute(v *vm.VM) {
	
}

func (o *XOR_A) String() string {
	return "XOR A"
}

type ADD_SP_r8 struct {}

func (o *ADD_SP_r8) Execute(v *vm.VM) {
	
}

func (o *ADD_SP_r8) String() string {
	return "ADD SP r8"
}

type LD_a16Deref_SP struct {}

func (o *LD_a16Deref_SP) Execute(v *vm.VM) {
	v.(a16)(v.SP())
}

func (o *LD_a16Deref_SP) String() string {
	return "LD (a16) SP"
}

type AND_B struct {}

func (o *AND_B) Execute(v *vm.VM) {
	
}

func (o *AND_B) String() string {
	return "AND B"
}

type CP_H struct {}

func (o *CP_H) Execute(v *vm.VM) {
	
}

func (o *CP_H) String() string {
	return "CP H"
}

type AND_d8 struct {}

func (o *AND_d8) Execute(v *vm.VM) {
	
}

func (o *AND_d8) String() string {
	return "AND d8"
}

type CCF_ struct {}

func (o *CCF_) Execute(v *vm.VM) {
	
}

func (o *CCF_) String() string {
	return "CCF"
}

type LD_B_D struct {}

func (o *LD_B_D) Execute(v *vm.VM) {
	v.B(v.D())
}

func (o *LD_B_D) String() string {
	return "LD B D"
}

type LD_H_A struct {}

func (o *LD_H_A) Execute(v *vm.VM) {
	v.H(v.A())
}

func (o *LD_H_A) String() string {
	return "LD H A"
}

type LD_L_D struct {}

func (o *LD_L_D) Execute(v *vm.VM) {
	v.L(v.D())
}

func (o *LD_L_D) String() string {
	return "LD L D"
}

type LD_A_HLPtr struct {}

func (o *LD_A_HLPtr) Execute(v *vm.VM) {
	v.A(v.(HL)())
}

func (o *LD_A_HLPtr) String() string {
	return "LD A (HL)"
}

type SUB_D struct {}

func (o *SUB_D) Execute(v *vm.VM) {
	
}

func (o *SUB_D) String() string {
	return "SUB D"
}

type PUSH_BC struct {}

func (o *PUSH_BC) Execute(v *vm.VM) {
	
}

func (o *PUSH_BC) String() string {
	return "PUSH BC"
}

type PUSH_HL struct {}

func (o *PUSH_HL) Execute(v *vm.VM) {
	
}

func (o *PUSH_HL) String() string {
	return "PUSH HL"
}

type INC_SP struct {}

func (o *INC_SP) Execute(v *vm.VM) {
	
}

func (o *INC_SP) String() string {
	return "INC SP"
}

type ADD_A_E struct {}

func (o *ADD_A_E) Execute(v *vm.VM) {
	
}

func (o *ADD_A_E) String() string {
	return "ADD A E"
}

type ADD_A_H struct {}

func (o *ADD_A_H) Execute(v *vm.VM) {
	
}

func (o *ADD_A_H) String() string {
	return "ADD A H"
}

type ADC_A_B struct {}

func (o *ADC_A_B) Execute(v *vm.VM) {
	
}

func (o *ADC_A_B) String() string {
	return "ADC A B"
}

type AND_C struct {}

func (o *AND_C) Execute(v *vm.VM) {
	
}

func (o *AND_C) String() string {
	return "AND C"
}

type CALL_NZ_a16 struct {}

func (o *CALL_NZ_a16) Execute(v *vm.VM) {
	
}

func (o *CALL_NZ_a16) String() string {
	return "CALL NZ a16"
}

type LD_E_d8 struct {}

func (o *LD_E_d8) Execute(v *vm.VM) {
	v.E(v.d8())
}

func (o *LD_E_d8) String() string {
	return "LD E d8"
}

type LD_L_C struct {}

func (o *LD_L_C) Execute(v *vm.VM) {
	v.L(v.C())
}

func (o *LD_L_C) String() string {
	return "LD L C"
}

type SUB_E struct {}

func (o *SUB_E) Execute(v *vm.VM) {
	
}

func (o *SUB_E) String() string {
	return "SUB E"
}

type SUB_d8 struct {}

func (o *SUB_d8) Execute(v *vm.VM) {
	
}

func (o *SUB_d8) String() string {
	return "SUB d8"
}

type DEC_HLPtr struct {}

func (o *DEC_HLPtr) Execute(v *vm.VM) {
	
}

func (o *DEC_HLPtr) String() string {
	return "DEC (HL)"
}

type SCF_ struct {}

func (o *SCF_) Execute(v *vm.VM) {
	
}

func (o *SCF_) String() string {
	return "SCF"
}

type RST_20H struct {}

func (o *RST_20H) Execute(v *vm.VM) {
	
}

func (o *RST_20H) String() string {
	return "RST 20H"
}

type INC_E struct {}

func (o *INC_E) Execute(v *vm.VM) {
	
}

func (o *INC_E) String() string {
	return "INC E"
}

type LD_HL_d16 struct {}

func (o *LD_HL_d16) Execute(v *vm.VM) {
	v.HL(v.d16())
}

func (o *LD_HL_d16) String() string {
	return "LD HL d16"
}

type LD_H_H struct {}

func (o *LD_H_H) Execute(v *vm.VM) {
	v.H(v.H())
}

func (o *LD_H_H) String() string {
	return "LD H H"
}

type LD_HLPtr_L struct {}

func (o *LD_HLPtr_L) Execute(v *vm.VM) {
	v.(HL)(v.L())
}

func (o *LD_HLPtr_L) String() string {
	return "LD (HL) L"
}

type SBC_A_C struct {}

func (o *SBC_A_C) Execute(v *vm.VM) {
	
}

func (o *SBC_A_C) String() string {
	return "SBC A C"
}

type SBC_A_L struct {}

func (o *SBC_A_L) Execute(v *vm.VM) {
	
}

func (o *SBC_A_L) String() string {
	return "SBC A L"
}

type DEC_BC struct {}

func (o *DEC_BC) Execute(v *vm.VM) {
	
}

func (o *DEC_BC) String() string {
	return "DEC BC"
}

type JR_NZ_r8 struct {}

func (o *JR_NZ_r8) Execute(v *vm.VM) {
	
}

func (o *JR_NZ_r8) String() string {
	return "JR NZ r8"
}

type DEC_H struct {}

func (o *DEC_H) Execute(v *vm.VM) {
	
}

func (o *DEC_H) String() string {
	return "DEC H"
}

type LD_C_C struct {}

func (o *LD_C_C) Execute(v *vm.VM) {
	v.C(v.C())
}

func (o *LD_C_C) String() string {
	return "LD C C"
}

type DI_ struct {}

func (o *DI_) Execute(v *vm.VM) {
	
}

func (o *DI_) String() string {
	return "DI"
}

type RES_5_H struct {}

func (o *RES_5_H) Execute(v *vm.VM) {
	
}

func (o *RES_5_H) String() string {
	return "RES 5 H"
}

type RES_7_B struct {}

func (o *RES_7_B) Execute(v *vm.VM) {
	
}

func (o *RES_7_B) String() string {
	return "RES 7 B"
}

type RL_E struct {}

func (o *RL_E) Execute(v *vm.VM) {
	
}

func (o *RL_E) String() string {
	return "RL E"
}

type BIT_6_HLPtr struct {}

func (o *BIT_6_HLPtr) Execute(v *vm.VM) {
	
}

func (o *BIT_6_HLPtr) String() string {
	return "BIT 6 (HL)"
}

type RES_0_C struct {}

func (o *RES_0_C) Execute(v *vm.VM) {
	
}

func (o *RES_0_C) String() string {
	return "RES 0 C"
}

type RES_0_L struct {}

func (o *RES_0_L) Execute(v *vm.VM) {
	
}

func (o *RES_0_L) String() string {
	return "RES 0 L"
}

type RES_3_E struct {}

func (o *RES_3_E) Execute(v *vm.VM) {
	
}

func (o *RES_3_E) String() string {
	return "RES 3 E"
}

type RES_0_A struct {}

func (o *RES_0_A) Execute(v *vm.VM) {
	
}

func (o *RES_0_A) String() string {
	return "RES 0 A"
}

type RES_1_B struct {}

func (o *RES_1_B) Execute(v *vm.VM) {
	
}

func (o *RES_1_B) String() string {
	return "RES 1 B"
}

type RES_4_D struct {}

func (o *RES_4_D) Execute(v *vm.VM) {
	
}

func (o *RES_4_D) String() string {
	return "RES 4 D"
}

type SET_1_D struct {}

func (o *SET_1_D) Execute(v *vm.VM) {
	
}

func (o *SET_1_D) String() string {
	return "SET 1 D"
}

type BIT_7_C struct {}

func (o *BIT_7_C) Execute(v *vm.VM) {
	
}

func (o *BIT_7_C) String() string {
	return "BIT 7 C"
}

type SRA_A struct {}

func (o *SRA_A) Execute(v *vm.VM) {
	
}

func (o *SRA_A) String() string {
	return "SRA A"
}

type SRA_HLPtr struct {}

func (o *SRA_HLPtr) Execute(v *vm.VM) {
	
}

func (o *SRA_HLPtr) String() string {
	return "SRA (HL)"
}

type RES_5_B struct {}

func (o *RES_5_B) Execute(v *vm.VM) {
	
}

func (o *RES_5_B) String() string {
	return "RES 5 B"
}

type SET_4_D struct {}

func (o *SET_4_D) Execute(v *vm.VM) {
	
}

func (o *SET_4_D) String() string {
	return "SET 4 D"
}

type RR_E struct {}

func (o *RR_E) Execute(v *vm.VM) {
	
}

func (o *RR_E) String() string {
	return "RR E"
}

type BIT_2_B struct {}

func (o *BIT_2_B) Execute(v *vm.VM) {
	
}

func (o *BIT_2_B) String() string {
	return "BIT 2 B"
}

type BIT_3_E struct {}

func (o *BIT_3_E) Execute(v *vm.VM) {
	
}

func (o *BIT_3_E) String() string {
	return "BIT 3 E"
}

type BIT_4_H struct {}

func (o *BIT_4_H) Execute(v *vm.VM) {
	
}

func (o *BIT_4_H) String() string {
	return "BIT 4 H"
}

type RES_4_H struct {}

func (o *RES_4_H) Execute(v *vm.VM) {
	
}

func (o *RES_4_H) String() string {
	return "RES 4 H"
}

type RL_C struct {}

func (o *RL_C) Execute(v *vm.VM) {
	
}

func (o *RL_C) String() string {
	return "RL C"
}

type BIT_6_D struct {}

func (o *BIT_6_D) Execute(v *vm.VM) {
	
}

func (o *BIT_6_D) String() string {
	return "BIT 6 D"
}

type RES_3_A struct {}

func (o *RES_3_A) Execute(v *vm.VM) {
	
}

func (o *RES_3_A) String() string {
	return "RES 3 A"
}

type SET_2_B struct {}

func (o *SET_2_B) Execute(v *vm.VM) {
	
}

func (o *SET_2_B) String() string {
	return "SET 2 B"
}

type SET_7_E struct {}

func (o *SET_7_E) Execute(v *vm.VM) {
	
}

func (o *SET_7_E) String() string {
	return "SET 7 E"
}

type BIT_4_E struct {}

func (o *BIT_4_E) Execute(v *vm.VM) {
	
}

func (o *BIT_4_E) String() string {
	return "BIT 4 E"
}

type RES_6_E struct {}

func (o *RES_6_E) Execute(v *vm.VM) {
	
}

func (o *RES_6_E) String() string {
	return "RES 6 E"
}

type SWAP_C struct {}

func (o *SWAP_C) Execute(v *vm.VM) {
	
}

func (o *SWAP_C) String() string {
	return "SWAP C"
}

type SWAP_H struct {}

func (o *SWAP_H) Execute(v *vm.VM) {
	
}

func (o *SWAP_H) String() string {
	return "SWAP H"
}

type SRL_HLPtr struct {}

func (o *SRL_HLPtr) Execute(v *vm.VM) {
	
}

func (o *SRL_HLPtr) String() string {
	return "SRL (HL)"
}

type BIT_1_HLPtr struct {}

func (o *BIT_1_HLPtr) Execute(v *vm.VM) {
	
}

func (o *BIT_1_HLPtr) String() string {
	return "BIT 1 (HL)"
}

type BIT_4_HLPtr struct {}

func (o *BIT_4_HLPtr) Execute(v *vm.VM) {
	
}

func (o *BIT_4_HLPtr) String() string {
	return "BIT 4 (HL)"
}

type BIT_5_B struct {}

func (o *BIT_5_B) Execute(v *vm.VM) {
	
}

func (o *BIT_5_B) String() string {
	return "BIT 5 B"
}

type SET_4_A struct {}

func (o *SET_4_A) Execute(v *vm.VM) {
	
}

func (o *SET_4_A) String() string {
	return "SET 4 A"
}

type SET_7_L struct {}

func (o *SET_7_L) Execute(v *vm.VM) {
	
}

func (o *SET_7_L) String() string {
	return "SET 7 L"
}

type SWAP_A struct {}

func (o *SWAP_A) Execute(v *vm.VM) {
	
}

func (o *SWAP_A) String() string {
	return "SWAP A"
}

type BIT_3_HLPtr struct {}

func (o *BIT_3_HLPtr) Execute(v *vm.VM) {
	
}

func (o *BIT_3_HLPtr) String() string {
	return "BIT 3 (HL)"
}

type BIT_4_L struct {}

func (o *BIT_4_L) Execute(v *vm.VM) {
	
}

func (o *BIT_4_L) String() string {
	return "BIT 4 L"
}

type BIT_5_H struct {}

func (o *BIT_5_H) Execute(v *vm.VM) {
	
}

func (o *BIT_5_H) String() string {
	return "BIT 5 H"
}

type SET_1_L struct {}

func (o *SET_1_L) Execute(v *vm.VM) {
	
}

func (o *SET_1_L) String() string {
	return "SET 1 L"
}

type RL_A struct {}

func (o *RL_A) Execute(v *vm.VM) {
	
}

func (o *RL_A) String() string {
	return "RL A"
}

type BIT_2_E struct {}

func (o *BIT_2_E) Execute(v *vm.VM) {
	
}

func (o *BIT_2_E) String() string {
	return "BIT 2 E"
}

type BIT_2_HLPtr struct {}

func (o *BIT_2_HLPtr) Execute(v *vm.VM) {
	
}

func (o *BIT_2_HLPtr) String() string {
	return "BIT 2 (HL)"
}

type RES_1_E struct {}

func (o *RES_1_E) Execute(v *vm.VM) {
	
}

func (o *RES_1_E) String() string {
	return "RES 1 E"
}

type RES_2_L struct {}

func (o *RES_2_L) Execute(v *vm.VM) {
	
}

func (o *RES_2_L) String() string {
	return "RES 2 L"
}

type RRC_H struct {}

func (o *RRC_H) Execute(v *vm.VM) {
	
}

func (o *RRC_H) String() string {
	return "RRC H"
}

type SRA_E struct {}

func (o *SRA_E) Execute(v *vm.VM) {
	
}

func (o *SRA_E) String() string {
	return "SRA E"
}

type BIT_5_HLPtr struct {}

func (o *BIT_5_HLPtr) Execute(v *vm.VM) {
	
}

func (o *BIT_5_HLPtr) String() string {
	return "BIT 5 (HL)"
}

type SET_4_E struct {}

func (o *SET_4_E) Execute(v *vm.VM) {
	
}

func (o *SET_4_E) String() string {
	return "SET 4 E"
}

type SET_5_H struct {}

func (o *SET_5_H) Execute(v *vm.VM) {
	
}

func (o *SET_5_H) String() string {
	return "SET 5 H"
}

type RLC_B struct {}

func (o *RLC_B) Execute(v *vm.VM) {
	
}

func (o *RLC_B) String() string {
	return "RLC B"
}

type RES_1_H struct {}

func (o *RES_1_H) Execute(v *vm.VM) {
	
}

func (o *RES_1_H) String() string {
	return "RES 1 H"
}

type RES_2_E struct {}

func (o *RES_2_E) Execute(v *vm.VM) {
	
}

func (o *RES_2_E) String() string {
	return "RES 2 E"
}

type SET_6_H struct {}

func (o *SET_6_H) Execute(v *vm.VM) {
	
}

func (o *SET_6_H) String() string {
	return "SET 6 H"
}

type SRA_L struct {}

func (o *SRA_L) Execute(v *vm.VM) {
	
}

func (o *SRA_L) String() string {
	return "SRA L"
}

type SET_3_E struct {}

func (o *SET_3_E) Execute(v *vm.VM) {
	
}

func (o *SET_3_E) String() string {
	return "SET 3 E"
}

type RRC_L struct {}

func (o *RRC_L) Execute(v *vm.VM) {
	
}

func (o *RRC_L) String() string {
	return "RRC L"
}

type RL_B struct {}

func (o *RL_B) Execute(v *vm.VM) {
	
}

func (o *RL_B) String() string {
	return "RL B"
}

type RES_2_B struct {}

func (o *RES_2_B) Execute(v *vm.VM) {
	
}

func (o *RES_2_B) String() string {
	return "RES 2 B"
}

type BIT_6_B struct {}

func (o *BIT_6_B) Execute(v *vm.VM) {
	
}

func (o *BIT_6_B) String() string {
	return "BIT 6 B"
}

type RES_5_A struct {}

func (o *RES_5_A) Execute(v *vm.VM) {
	
}

func (o *RES_5_A) String() string {
	return "RES 5 A"
}

type SET_0_A struct {}

func (o *SET_0_A) Execute(v *vm.VM) {
	
}

func (o *SET_0_A) String() string {
	return "SET 0 A"
}

type SLA_C struct {}

func (o *SLA_C) Execute(v *vm.VM) {
	
}

func (o *SLA_C) String() string {
	return "SLA C"
}

type BIT_7_E struct {}

func (o *BIT_7_E) Execute(v *vm.VM) {
	
}

func (o *BIT_7_E) String() string {
	return "BIT 7 E"
}

type RES_6_L struct {}

func (o *RES_6_L) Execute(v *vm.VM) {
	
}

func (o *RES_6_L) String() string {
	return "RES 6 L"
}

type SET_2_C struct {}

func (o *SET_2_C) Execute(v *vm.VM) {
	
}

func (o *SET_2_C) String() string {
	return "SET 2 C"
}

type SET_5_E struct {}

func (o *SET_5_E) Execute(v *vm.VM) {
	
}

func (o *SET_5_E) String() string {
	return "SET 5 E"
}

type RR_D struct {}

func (o *RR_D) Execute(v *vm.VM) {
	
}

func (o *RR_D) String() string {
	return "RR D"
}

type RR_L struct {}

func (o *RR_L) Execute(v *vm.VM) {
	
}

func (o *RR_L) String() string {
	return "RR L"
}

type BIT_5_E struct {}

func (o *BIT_5_E) Execute(v *vm.VM) {
	
}

func (o *BIT_5_E) String() string {
	return "BIT 5 E"
}

type RES_5_HLPtr struct {}

func (o *RES_5_HLPtr) Execute(v *vm.VM) {
	
}

func (o *RES_5_HLPtr) String() string {
	return "RES 5 (HL)"
}

type SET_3_H struct {}

func (o *SET_3_H) Execute(v *vm.VM) {
	
}

func (o *SET_3_H) String() string {
	return "SET 3 H"
}

type RES_5_L struct {}

func (o *RES_5_L) Execute(v *vm.VM) {
	
}

func (o *RES_5_L) String() string {
	return "RES 5 L"
}

type RES_7_H struct {}

func (o *RES_7_H) Execute(v *vm.VM) {
	
}

func (o *RES_7_H) String() string {
	return "RES 7 H"
}

type SET_3_C struct {}

func (o *SET_3_C) Execute(v *vm.VM) {
	
}

func (o *SET_3_C) String() string {
	return "SET 3 C"
}

type SRA_B struct {}

func (o *SRA_B) Execute(v *vm.VM) {
	
}

func (o *SRA_B) String() string {
	return "SRA B"
}

type SWAP_E struct {}

func (o *SWAP_E) Execute(v *vm.VM) {
	
}

func (o *SWAP_E) String() string {
	return "SWAP E"
}

type SRL_C struct {}

func (o *SRL_C) Execute(v *vm.VM) {
	
}

func (o *SRL_C) String() string {
	return "SRL C"
}

type BIT_0_D struct {}

func (o *BIT_0_D) Execute(v *vm.VM) {
	
}

func (o *BIT_0_D) String() string {
	return "BIT 0 D"
}

type BIT_7_H struct {}

func (o *BIT_7_H) Execute(v *vm.VM) {
	
}

func (o *BIT_7_H) String() string {
	return "BIT 7 H"
}

type RES_2_D struct {}

func (o *RES_2_D) Execute(v *vm.VM) {
	
}

func (o *RES_2_D) String() string {
	return "RES 2 D"
}

type RES_2_H struct {}

func (o *RES_2_H) Execute(v *vm.VM) {
	
}

func (o *RES_2_H) String() string {
	return "RES 2 H"
}

type SET_4_C struct {}

func (o *SET_4_C) Execute(v *vm.VM) {
	
}

func (o *SET_4_C) String() string {
	return "SET 4 C"
}

type SLA_L struct {}

func (o *SLA_L) Execute(v *vm.VM) {
	
}

func (o *SLA_L) String() string {
	return "SLA L"
}

type BIT_5_D struct {}

func (o *BIT_5_D) Execute(v *vm.VM) {
	
}

func (o *BIT_5_D) String() string {
	return "BIT 5 D"
}

type SET_2_A struct {}

func (o *SET_2_A) Execute(v *vm.VM) {
	
}

func (o *SET_2_A) String() string {
	return "SET 2 A"
}

type RES_5_C struct {}

func (o *RES_5_C) Execute(v *vm.VM) {
	
}

func (o *RES_5_C) String() string {
	return "RES 5 C"
}

type SET_4_HLPtr struct {}

func (o *SET_4_HLPtr) Execute(v *vm.VM) {
	
}

func (o *SET_4_HLPtr) String() string {
	return "SET 4 (HL)"
}

type SET_6_B struct {}

func (o *SET_6_B) Execute(v *vm.VM) {
	
}

func (o *SET_6_B) String() string {
	return "SET 6 B"
}

type SRA_C struct {}

func (o *SRA_C) Execute(v *vm.VM) {
	
}

func (o *SRA_C) String() string {
	return "SRA C"
}

type BIT_1_C struct {}

func (o *BIT_1_C) Execute(v *vm.VM) {
	
}

func (o *BIT_1_C) String() string {
	return "BIT 1 C"
}

type BIT_4_A struct {}

func (o *BIT_4_A) Execute(v *vm.VM) {
	
}

func (o *BIT_4_A) String() string {
	return "BIT 4 A"
}

type RES_3_H struct {}

func (o *RES_3_H) Execute(v *vm.VM) {
	
}

func (o *RES_3_H) String() string {
	return "RES 3 H"
}

type RES_4_C struct {}

func (o *RES_4_C) Execute(v *vm.VM) {
	
}

func (o *RES_4_C) String() string {
	return "RES 4 C"
}

type SET_7_B struct {}

func (o *SET_7_B) Execute(v *vm.VM) {
	
}

func (o *SET_7_B) String() string {
	return "SET 7 B"
}

type RLC_E struct {}

func (o *RLC_E) Execute(v *vm.VM) {
	
}

func (o *RLC_E) String() string {
	return "RLC E"
}

type SET_5_HLPtr struct {}

func (o *SET_5_HLPtr) Execute(v *vm.VM) {
	
}

func (o *SET_5_HLPtr) String() string {
	return "SET 5 (HL)"
}

type BIT_7_HLPtr struct {}

func (o *BIT_7_HLPtr) Execute(v *vm.VM) {
	
}

func (o *BIT_7_HLPtr) String() string {
	return "BIT 7 (HL)"
}

type SET_3_L struct {}

func (o *SET_3_L) Execute(v *vm.VM) {
	
}

func (o *SET_3_L) String() string {
	return "SET 3 L"
}

type RES_6_HLPtr struct {}

func (o *RES_6_HLPtr) Execute(v *vm.VM) {
	
}

func (o *RES_6_HLPtr) String() string {
	return "RES 6 (HL)"
}

type RES_7_E struct {}

func (o *RES_7_E) Execute(v *vm.VM) {
	
}

func (o *RES_7_E) String() string {
	return "RES 7 E"
}

type SET_3_D struct {}

func (o *SET_3_D) Execute(v *vm.VM) {
	
}

func (o *SET_3_D) String() string {
	return "SET 3 D"
}

type RLC_L struct {}

func (o *RLC_L) Execute(v *vm.VM) {
	
}

func (o *RLC_L) String() string {
	return "RLC L"
}

type RR_HLPtr struct {}

func (o *RR_HLPtr) Execute(v *vm.VM) {
	
}

func (o *RR_HLPtr) String() string {
	return "RR (HL)"
}

type BIT_4_B struct {}

func (o *BIT_4_B) Execute(v *vm.VM) {
	
}

func (o *BIT_4_B) String() string {
	return "BIT 4 B"
}

type BIT_5_L struct {}

func (o *BIT_5_L) Execute(v *vm.VM) {
	
}

func (o *BIT_5_L) String() string {
	return "BIT 5 L"
}

type RES_4_L struct {}

func (o *RES_4_L) Execute(v *vm.VM) {
	
}

func (o *RES_4_L) String() string {
	return "RES 4 L"
}

type RR_C struct {}

func (o *RR_C) Execute(v *vm.VM) {
	
}

func (o *RR_C) String() string {
	return "RR C"
}

type BIT_2_H struct {}

func (o *BIT_2_H) Execute(v *vm.VM) {
	
}

func (o *BIT_2_H) String() string {
	return "BIT 2 H"
}

type RES_4_A struct {}

func (o *RES_4_A) Execute(v *vm.VM) {
	
}

func (o *RES_4_A) String() string {
	return "RES 4 A"
}

type SET_0_E struct {}

func (o *SET_0_E) Execute(v *vm.VM) {
	
}

func (o *SET_0_E) String() string {
	return "SET 0 E"
}

type SET_7_A struct {}

func (o *SET_7_A) Execute(v *vm.VM) {
	
}

func (o *SET_7_A) String() string {
	return "SET 7 A"
}

type RES_6_D struct {}

func (o *RES_6_D) Execute(v *vm.VM) {
	
}

func (o *RES_6_D) String() string {
	return "RES 6 D"
}

type SET_2_L struct {}

func (o *SET_2_L) Execute(v *vm.VM) {
	
}

func (o *SET_2_L) String() string {
	return "SET 2 L"
}

type SET_5_A struct {}

func (o *SET_5_A) Execute(v *vm.VM) {
	
}

func (o *SET_5_A) String() string {
	return "SET 5 A"
}

type RRC_HLPtr struct {}

func (o *RRC_HLPtr) Execute(v *vm.VM) {
	
}

func (o *RRC_HLPtr) String() string {
	return "RRC (HL)"
}

type RR_A struct {}

func (o *RR_A) Execute(v *vm.VM) {
	
}

func (o *RR_A) String() string {
	return "RR A"
}

type SLA_HLPtr struct {}

func (o *SLA_HLPtr) Execute(v *vm.VM) {
	
}

func (o *SLA_HLPtr) String() string {
	return "SLA (HL)"
}

type BIT_2_L struct {}

func (o *BIT_2_L) Execute(v *vm.VM) {
	
}

func (o *BIT_2_L) String() string {
	return "BIT 2 L"
}

type RES_5_E struct {}

func (o *RES_5_E) Execute(v *vm.VM) {
	
}

func (o *RES_5_E) String() string {
	return "RES 5 E"
}

type SET_7_C struct {}

func (o *SET_7_C) Execute(v *vm.VM) {
	
}

func (o *SET_7_C) String() string {
	return "SET 7 C"
}

type RL_H struct {}

func (o *RL_H) Execute(v *vm.VM) {
	
}

func (o *RL_H) String() string {
	return "RL H"
}

type RR_B struct {}

func (o *RR_B) Execute(v *vm.VM) {
	
}

func (o *RR_B) String() string {
	return "RR B"
}

type SWAP_L struct {}

func (o *SWAP_L) Execute(v *vm.VM) {
	
}

func (o *SWAP_L) String() string {
	return "SWAP L"
}

type SET_6_C struct {}

func (o *SET_6_C) Execute(v *vm.VM) {
	
}

func (o *SET_6_C) String() string {
	return "SET 6 C"
}

type SET_7_HLPtr struct {}

func (o *SET_7_HLPtr) Execute(v *vm.VM) {
	
}

func (o *SET_7_HLPtr) String() string {
	return "SET 7 (HL)"
}

type BIT_5_A struct {}

func (o *BIT_5_A) Execute(v *vm.VM) {
	
}

func (o *BIT_5_A) String() string {
	return "BIT 5 A"
}

type BIT_6_A struct {}

func (o *BIT_6_A) Execute(v *vm.VM) {
	
}

func (o *BIT_6_A) String() string {
	return "BIT 6 A"
}

type SET_1_C struct {}

func (o *SET_1_C) Execute(v *vm.VM) {
	
}

func (o *SET_1_C) String() string {
	return "SET 1 C"
}

type BIT_0_A struct {}

func (o *BIT_0_A) Execute(v *vm.VM) {
	
}

func (o *BIT_0_A) String() string {
	return "BIT 0 A"
}

type BIT_1_B struct {}

func (o *BIT_1_B) Execute(v *vm.VM) {
	
}

func (o *BIT_1_B) String() string {
	return "BIT 1 B"
}

type SET_4_L struct {}

func (o *SET_4_L) Execute(v *vm.VM) {
	
}

func (o *SET_4_L) String() string {
	return "SET 4 L"
}

type RRC_E struct {}

func (o *RRC_E) Execute(v *vm.VM) {
	
}

func (o *RRC_E) String() string {
	return "RRC E"
}

type SLA_H struct {}

func (o *SLA_H) Execute(v *vm.VM) {
	
}

func (o *SLA_H) String() string {
	return "SLA H"
}

type RES_2_A struct {}

func (o *RES_2_A) Execute(v *vm.VM) {
	
}

func (o *RES_2_A) String() string {
	return "RES 2 A"
}

type RES_6_H struct {}

func (o *RES_6_H) Execute(v *vm.VM) {
	
}

func (o *RES_6_H) String() string {
	return "RES 6 H"
}

type SET_5_C struct {}

func (o *SET_5_C) Execute(v *vm.VM) {
	
}

func (o *SET_5_C) String() string {
	return "SET 5 C"
}

type BIT_0_B struct {}

func (o *BIT_0_B) Execute(v *vm.VM) {
	
}

func (o *BIT_0_B) String() string {
	return "BIT 0 B"
}

type BIT_0_E struct {}

func (o *BIT_0_E) Execute(v *vm.VM) {
	
}

func (o *BIT_0_E) String() string {
	return "BIT 0 E"
}

type RES_5_D struct {}

func (o *RES_5_D) Execute(v *vm.VM) {
	
}

func (o *RES_5_D) String() string {
	return "RES 5 D"
}

type SET_3_A struct {}

func (o *SET_3_A) Execute(v *vm.VM) {
	
}

func (o *SET_3_A) String() string {
	return "SET 3 A"
}

type RRC_C struct {}

func (o *RRC_C) Execute(v *vm.VM) {
	
}

func (o *RRC_C) String() string {
	return "RRC C"
}

type SRA_H struct {}

func (o *SRA_H) Execute(v *vm.VM) {
	
}

func (o *SRA_H) String() string {
	return "SRA H"
}

type SRL_E struct {}

func (o *SRL_E) Execute(v *vm.VM) {
	
}

func (o *SRL_E) String() string {
	return "SRL E"
}

type BIT_0_C struct {}

func (o *BIT_0_C) Execute(v *vm.VM) {
	
}

func (o *BIT_0_C) String() string {
	return "BIT 0 C"
}

type BIT_4_D struct {}

func (o *BIT_4_D) Execute(v *vm.VM) {
	
}

func (o *BIT_4_D) String() string {
	return "BIT 4 D"
}

type SET_7_H struct {}

func (o *SET_7_H) Execute(v *vm.VM) {
	
}

func (o *SET_7_H) String() string {
	return "SET 7 H"
}

type RL_HLPtr struct {}

func (o *RL_HLPtr) Execute(v *vm.VM) {
	
}

func (o *RL_HLPtr) String() string {
	return "RL (HL)"
}

type BIT_6_L struct {}

func (o *BIT_6_L) Execute(v *vm.VM) {
	
}

func (o *BIT_6_L) String() string {
	return "BIT 6 L"
}

type RES_6_C struct {}

func (o *RES_6_C) Execute(v *vm.VM) {
	
}

func (o *RES_6_C) String() string {
	return "RES 6 C"
}

type SET_5_B struct {}

func (o *SET_5_B) Execute(v *vm.VM) {
	
}

func (o *SET_5_B) String() string {
	return "SET 5 B"
}

type BIT_5_C struct {}

func (o *BIT_5_C) Execute(v *vm.VM) {
	
}

func (o *BIT_5_C) String() string {
	return "BIT 5 C"
}

type RES_0_E struct {}

func (o *RES_0_E) Execute(v *vm.VM) {
	
}

func (o *RES_0_E) String() string {
	return "RES 0 E"
}

type RES_4_HLPtr struct {}

func (o *RES_4_HLPtr) Execute(v *vm.VM) {
	
}

func (o *RES_4_HLPtr) String() string {
	return "RES 4 (HL)"
}

type RES_7_C struct {}

func (o *RES_7_C) Execute(v *vm.VM) {
	
}

func (o *RES_7_C) String() string {
	return "RES 7 C"
}

type RLC_C struct {}

func (o *RLC_C) Execute(v *vm.VM) {
	
}

func (o *RLC_C) String() string {
	return "RLC C"
}

type BIT_3_L struct {}

func (o *BIT_3_L) Execute(v *vm.VM) {
	
}

func (o *BIT_3_L) String() string {
	return "BIT 3 L"
}

type RES_1_L struct {}

func (o *RES_1_L) Execute(v *vm.VM) {
	
}

func (o *RES_1_L) String() string {
	return "RES 1 L"
}

type RLC_HLPtr struct {}

func (o *RLC_HLPtr) Execute(v *vm.VM) {
	
}

func (o *RLC_HLPtr) String() string {
	return "RLC (HL)"
}

type BIT_1_H struct {}

func (o *BIT_1_H) Execute(v *vm.VM) {
	
}

func (o *BIT_1_H) String() string {
	return "BIT 1 H"
}

type SET_1_B struct {}

func (o *SET_1_B) Execute(v *vm.VM) {
	
}

func (o *SET_1_B) String() string {
	return "SET 1 B"
}

type SET_1_A struct {}

func (o *SET_1_A) Execute(v *vm.VM) {
	
}

func (o *SET_1_A) String() string {
	return "SET 1 A"
}

type SET_6_D struct {}

func (o *SET_6_D) Execute(v *vm.VM) {
	
}

func (o *SET_6_D) String() string {
	return "SET 6 D"
}

type BIT_1_A struct {}

func (o *BIT_1_A) Execute(v *vm.VM) {
	
}

func (o *BIT_1_A) String() string {
	return "BIT 1 A"
}

type SET_6_E struct {}

func (o *SET_6_E) Execute(v *vm.VM) {
	
}

func (o *SET_6_E) String() string {
	return "SET 6 E"
}

type RRC_B struct {}

func (o *RRC_B) Execute(v *vm.VM) {
	
}

func (o *RRC_B) String() string {
	return "RRC B"
}

type SLA_E struct {}

func (o *SLA_E) Execute(v *vm.VM) {
	
}

func (o *SLA_E) String() string {
	return "SLA E"
}

type BIT_3_B struct {}

func (o *BIT_3_B) Execute(v *vm.VM) {
	
}

func (o *BIT_3_B) String() string {
	return "BIT 3 B"
}

type RES_0_H struct {}

func (o *RES_0_H) Execute(v *vm.VM) {
	
}

func (o *RES_0_H) String() string {
	return "RES 0 H"
}

type BIT_7_A struct {}

func (o *BIT_7_A) Execute(v *vm.VM) {
	
}

func (o *BIT_7_A) String() string {
	return "BIT 7 A"
}

type RES_6_A struct {}

func (o *RES_6_A) Execute(v *vm.VM) {
	
}

func (o *RES_6_A) String() string {
	return "RES 6 A"
}

type SRA_D struct {}

func (o *SRA_D) Execute(v *vm.VM) {
	
}

func (o *SRA_D) String() string {
	return "SRA D"
}

type BIT_7_D struct {}

func (o *BIT_7_D) Execute(v *vm.VM) {
	
}

func (o *BIT_7_D) String() string {
	return "BIT 7 D"
}

type RES_7_L struct {}

func (o *RES_7_L) Execute(v *vm.VM) {
	
}

func (o *RES_7_L) String() string {
	return "RES 7 L"
}

type SET_0_L struct {}

func (o *SET_0_L) Execute(v *vm.VM) {
	
}

func (o *SET_0_L) String() string {
	return "SET 0 L"
}

type SET_0_HLPtr struct {}

func (o *SET_0_HLPtr) Execute(v *vm.VM) {
	
}

func (o *SET_0_HLPtr) String() string {
	return "SET 0 (HL)"
}

type BIT_0_HLPtr struct {}

func (o *BIT_0_HLPtr) Execute(v *vm.VM) {
	
}

func (o *BIT_0_HLPtr) String() string {
	return "BIT 0 (HL)"
}

type BIT_1_L struct {}

func (o *BIT_1_L) Execute(v *vm.VM) {
	
}

func (o *BIT_1_L) String() string {
	return "BIT 1 L"
}

type RES_3_L struct {}

func (o *RES_3_L) Execute(v *vm.VM) {
	
}

func (o *RES_3_L) String() string {
	return "RES 3 L"
}

type RLC_A struct {}

func (o *RLC_A) Execute(v *vm.VM) {
	
}

func (o *RLC_A) String() string {
	return "RLC A"
}

type RR_H struct {}

func (o *RR_H) Execute(v *vm.VM) {
	
}

func (o *RR_H) String() string {
	return "RR H"
}

type SWAP_B struct {}

func (o *SWAP_B) Execute(v *vm.VM) {
	
}

func (o *SWAP_B) String() string {
	return "SWAP B"
}

type SRL_D struct {}

func (o *SRL_D) Execute(v *vm.VM) {
	
}

func (o *SRL_D) String() string {
	return "SRL D"
}

type BIT_0_L struct {}

func (o *BIT_0_L) Execute(v *vm.VM) {
	
}

func (o *BIT_0_L) String() string {
	return "BIT 0 L"
}

type RES_7_HLPtr struct {}

func (o *RES_7_HLPtr) Execute(v *vm.VM) {
	
}

func (o *RES_7_HLPtr) String() string {
	return "RES 7 (HL)"
}

type SET_0_C struct {}

func (o *SET_0_C) Execute(v *vm.VM) {
	
}

func (o *SET_0_C) String() string {
	return "SET 0 C"
}

type SET_3_B struct {}

func (o *SET_3_B) Execute(v *vm.VM) {
	
}

func (o *SET_3_B) String() string {
	return "SET 3 B"
}

type SET_7_D struct {}

func (o *SET_7_D) Execute(v *vm.VM) {
	
}

func (o *SET_7_D) String() string {
	return "SET 7 D"
}

type SLA_D struct {}

func (o *SLA_D) Execute(v *vm.VM) {
	
}

func (o *SLA_D) String() string {
	return "SLA D"
}

type BIT_6_H struct {}

func (o *BIT_6_H) Execute(v *vm.VM) {
	
}

func (o *BIT_6_H) String() string {
	return "BIT 6 H"
}

type RES_1_A struct {}

func (o *RES_1_A) Execute(v *vm.VM) {
	
}

func (o *RES_1_A) String() string {
	return "RES 1 A"
}

type RES_6_B struct {}

func (o *RES_6_B) Execute(v *vm.VM) {
	
}

func (o *RES_6_B) String() string {
	return "RES 6 B"
}

type SET_6_L struct {}

func (o *SET_6_L) Execute(v *vm.VM) {
	
}

func (o *SET_6_L) String() string {
	return "SET 6 L"
}

type SWAP_HLPtr struct {}

func (o *SWAP_HLPtr) Execute(v *vm.VM) {
	
}

func (o *SWAP_HLPtr) String() string {
	return "SWAP (HL)"
}

type SRL_B struct {}

func (o *SRL_B) Execute(v *vm.VM) {
	
}

func (o *SRL_B) String() string {
	return "SRL B"
}

type RES_4_E struct {}

func (o *RES_4_E) Execute(v *vm.VM) {
	
}

func (o *RES_4_E) String() string {
	return "RES 4 E"
}

type RES_7_A struct {}

func (o *RES_7_A) Execute(v *vm.VM) {
	
}

func (o *RES_7_A) String() string {
	return "RES 7 A"
}

type RL_D struct {}

func (o *RL_D) Execute(v *vm.VM) {
	
}

func (o *RL_D) String() string {
	return "RL D"
}

type SLA_B struct {}

func (o *SLA_B) Execute(v *vm.VM) {
	
}

func (o *SLA_B) String() string {
	return "SLA B"
}

type SET_0_H struct {}

func (o *SET_0_H) Execute(v *vm.VM) {
	
}

func (o *SET_0_H) String() string {
	return "SET 0 H"
}

type SET_3_HLPtr struct {}

func (o *SET_3_HLPtr) Execute(v *vm.VM) {
	
}

func (o *SET_3_HLPtr) String() string {
	return "SET 3 (HL)"
}

type SRL_H struct {}

func (o *SRL_H) Execute(v *vm.VM) {
	
}

func (o *SRL_H) String() string {
	return "SRL H"
}

type BIT_3_A struct {}

func (o *BIT_3_A) Execute(v *vm.VM) {
	
}

func (o *BIT_3_A) String() string {
	return "BIT 3 A"
}

type RES_7_D struct {}

func (o *RES_7_D) Execute(v *vm.VM) {
	
}

func (o *RES_7_D) String() string {
	return "RES 7 D"
}

type BIT_1_E struct {}

func (o *BIT_1_E) Execute(v *vm.VM) {
	
}

func (o *BIT_1_E) String() string {
	return "BIT 1 E"
}

type BIT_6_E struct {}

func (o *BIT_6_E) Execute(v *vm.VM) {
	
}

func (o *BIT_6_E) String() string {
	return "BIT 6 E"
}

type SET_1_H struct {}

func (o *SET_1_H) Execute(v *vm.VM) {
	
}

func (o *SET_1_H) String() string {
	return "SET 1 H"
}

type BIT_3_C struct {}

func (o *BIT_3_C) Execute(v *vm.VM) {
	
}

func (o *BIT_3_C) String() string {
	return "BIT 3 C"
}

type BIT_3_D struct {}

func (o *BIT_3_D) Execute(v *vm.VM) {
	
}

func (o *BIT_3_D) String() string {
	return "BIT 3 D"
}

type RES_3_C struct {}

func (o *RES_3_C) Execute(v *vm.VM) {
	
}

func (o *RES_3_C) String() string {
	return "RES 3 C"
}

type SET_0_B struct {}

func (o *SET_0_B) Execute(v *vm.VM) {
	
}

func (o *SET_0_B) String() string {
	return "SET 0 B"
}

type SET_0_D struct {}

func (o *SET_0_D) Execute(v *vm.VM) {
	
}

func (o *SET_0_D) String() string {
	return "SET 0 D"
}

type RRC_D struct {}

func (o *RRC_D) Execute(v *vm.VM) {
	
}

func (o *RRC_D) String() string {
	return "RRC D"
}

type SRL_L struct {}

func (o *SRL_L) Execute(v *vm.VM) {
	
}

func (o *SRL_L) String() string {
	return "SRL L"
}

type BIT_7_B struct {}

func (o *BIT_7_B) Execute(v *vm.VM) {
	
}

func (o *BIT_7_B) String() string {
	return "BIT 7 B"
}

type BIT_7_L struct {}

func (o *BIT_7_L) Execute(v *vm.VM) {
	
}

func (o *BIT_7_L) String() string {
	return "BIT 7 L"
}

type RES_3_B struct {}

func (o *RES_3_B) Execute(v *vm.VM) {
	
}

func (o *RES_3_B) String() string {
	return "RES 3 B"
}

type SET_4_H struct {}

func (o *SET_4_H) Execute(v *vm.VM) {
	
}

func (o *SET_4_H) String() string {
	return "SET 4 H"
}

type RES_0_D struct {}

func (o *RES_0_D) Execute(v *vm.VM) {
	
}

func (o *RES_0_D) String() string {
	return "RES 0 D"
}

type SET_2_H struct {}

func (o *SET_2_H) Execute(v *vm.VM) {
	
}

func (o *SET_2_H) String() string {
	return "SET 2 H"
}

type RL_L struct {}

func (o *RL_L) Execute(v *vm.VM) {
	
}

func (o *RL_L) String() string {
	return "RL L"
}

type BIT_2_A struct {}

func (o *BIT_2_A) Execute(v *vm.VM) {
	
}

func (o *BIT_2_A) String() string {
	return "BIT 2 A"
}

type RES_0_HLPtr struct {}

func (o *RES_0_HLPtr) Execute(v *vm.VM) {
	
}

func (o *RES_0_HLPtr) String() string {
	return "RES 0 (HL)"
}

type BIT_2_D struct {}

func (o *BIT_2_D) Execute(v *vm.VM) {
	
}

func (o *BIT_2_D) String() string {
	return "BIT 2 D"
}

type BIT_3_H struct {}

func (o *BIT_3_H) Execute(v *vm.VM) {
	
}

func (o *BIT_3_H) String() string {
	return "BIT 3 H"
}

type SET_1_HLPtr struct {}

func (o *SET_1_HLPtr) Execute(v *vm.VM) {
	
}

func (o *SET_1_HLPtr) String() string {
	return "SET 1 (HL)"
}

type SET_5_D struct {}

func (o *SET_5_D) Execute(v *vm.VM) {
	
}

func (o *SET_5_D) String() string {
	return "SET 5 D"
}

type SET_5_L struct {}

func (o *SET_5_L) Execute(v *vm.VM) {
	
}

func (o *SET_5_L) String() string {
	return "SET 5 L"
}

type RLC_D struct {}

func (o *RLC_D) Execute(v *vm.VM) {
	
}

func (o *RLC_D) String() string {
	return "RLC D"
}

type SET_6_A struct {}

func (o *SET_6_A) Execute(v *vm.VM) {
	
}

func (o *SET_6_A) String() string {
	return "SET 6 A"
}

type RLC_H struct {}

func (o *RLC_H) Execute(v *vm.VM) {
	
}

func (o *RLC_H) String() string {
	return "RLC H"
}

type RRC_A struct {}

func (o *RRC_A) Execute(v *vm.VM) {
	
}

func (o *RRC_A) String() string {
	return "RRC A"
}

type SWAP_D struct {}

func (o *SWAP_D) Execute(v *vm.VM) {
	
}

func (o *SWAP_D) String() string {
	return "SWAP D"
}

type SRL_A struct {}

func (o *SRL_A) Execute(v *vm.VM) {
	
}

func (o *SRL_A) String() string {
	return "SRL A"
}

type RES_1_C struct {}

func (o *RES_1_C) Execute(v *vm.VM) {
	
}

func (o *RES_1_C) String() string {
	return "RES 1 C"
}

type BIT_1_D struct {}

func (o *BIT_1_D) Execute(v *vm.VM) {
	
}

func (o *BIT_1_D) String() string {
	return "BIT 1 D"
}

type RES_1_HLPtr struct {}

func (o *RES_1_HLPtr) Execute(v *vm.VM) {
	
}

func (o *RES_1_HLPtr) String() string {
	return "RES 1 (HL)"
}

type SET_2_D struct {}

func (o *SET_2_D) Execute(v *vm.VM) {
	
}

func (o *SET_2_D) String() string {
	return "SET 2 D"
}

type SET_6_HLPtr struct {}

func (o *SET_6_HLPtr) Execute(v *vm.VM) {
	
}

func (o *SET_6_HLPtr) String() string {
	return "SET 6 (HL)"
}

type BIT_0_H struct {}

func (o *BIT_0_H) Execute(v *vm.VM) {
	
}

func (o *BIT_0_H) String() string {
	return "BIT 0 H"
}

type BIT_4_C struct {}

func (o *BIT_4_C) Execute(v *vm.VM) {
	
}

func (o *BIT_4_C) String() string {
	return "BIT 4 C"
}

type BIT_6_C struct {}

func (o *BIT_6_C) Execute(v *vm.VM) {
	
}

func (o *BIT_6_C) String() string {
	return "BIT 6 C"
}

type SET_4_B struct {}

func (o *SET_4_B) Execute(v *vm.VM) {
	
}

func (o *SET_4_B) String() string {
	return "SET 4 B"
}

type RES_0_B struct {}

func (o *RES_0_B) Execute(v *vm.VM) {
	
}

func (o *RES_0_B) String() string {
	return "RES 0 B"
}

type RES_2_HLPtr struct {}

func (o *RES_2_HLPtr) Execute(v *vm.VM) {
	
}

func (o *RES_2_HLPtr) String() string {
	return "RES 2 (HL)"
}

type RES_3_D struct {}

func (o *RES_3_D) Execute(v *vm.VM) {
	
}

func (o *RES_3_D) String() string {
	return "RES 3 D"
}

type SET_1_E struct {}

func (o *SET_1_E) Execute(v *vm.VM) {
	
}

func (o *SET_1_E) String() string {
	return "SET 1 E"
}

type SLA_A struct {}

func (o *SLA_A) Execute(v *vm.VM) {
	
}

func (o *SLA_A) String() string {
	return "SLA A"
}

type BIT_2_C struct {}

func (o *BIT_2_C) Execute(v *vm.VM) {
	
}

func (o *BIT_2_C) String() string {
	return "BIT 2 C"
}

type RES_2_C struct {}

func (o *RES_2_C) Execute(v *vm.VM) {
	
}

func (o *RES_2_C) String() string {
	return "RES 2 C"
}

type SET_2_E struct {}

func (o *SET_2_E) Execute(v *vm.VM) {
	
}

func (o *SET_2_E) String() string {
	return "SET 2 E"
}

type SET_2_HLPtr struct {}

func (o *SET_2_HLPtr) Execute(v *vm.VM) {
	
}

func (o *SET_2_HLPtr) String() string {
	return "SET 2 (HL)"
}

type RES_1_D struct {}

func (o *RES_1_D) Execute(v *vm.VM) {
	
}

func (o *RES_1_D) String() string {
	return "RES 1 D"
}

type RES_3_HLPtr struct {}

func (o *RES_3_HLPtr) Execute(v *vm.VM) {
	
}

func (o *RES_3_HLPtr) String() string {
	return "RES 3 (HL)"
}

type RES_4_B struct {}

func (o *RES_4_B) Execute(v *vm.VM) {
	
}

func (o *RES_4_B) String() string {
	return "RES 4 B"
}

