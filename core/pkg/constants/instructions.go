package constants

const (
	CLS         string = "CLS"
	RET                = "RET"
	JP_NNN             = "JP_NNN"
	CALL_NNN           = "CALL_NNN"
	SE_VX_KK           = "SE_VX_KK"
	SNE_VX_KK          = "SNE_VX_KK"
	SE_VX_VY           = "SE_VX_VY"
	LD_VX_KK           = "LD_VX_KK"
	ADD_VX_KK          = "ADD_VX_KK"
	LD_VX_VY           = "LD_VX_VY"
	OR_VX_VY           = "OR_VX_VY"
	AND_VX_VY          = "AND_VX_VY"
	XOR_VX_VY          = "XOR_VX_VY"
	ADD_VX_VY          = "ADD_VX_VY"
	SUB_VX_VY          = "SUB_VX_VY"
	SHR_VX_VY          = "SHR_VX_VY"
	SUBN_VX_VY         = "SUBN_VX_VY"
	SHL_VX_VY          = "SHL_VX_VY"
	SNE_VX_VY          = "SNE_VX_VY"
	LD_I_NNN           = "LD_I_NNN"
	JP_V0_NNN          = "JP_V0_NNN"
	RND_VX_KK          = "RND_VX_KK"
	DRW_VX_VY_N        = "DRW_VX_VY_N"
	SKP_VX             = "SKP_VX"
	SKNP_VX            = "SKNP_VX"
	LD_VX_DT           = "LD_VX_DT"
	LD_VX_K            = "LD_VX_K"
	LD_DT_VX           = "LD_DT_VX"
	LD_ST_VX           = "LD_ST_VX"
	ADD_I_VX           = "ADD_I_VX"
	LD_F_VX            = "LD_F_VX"
	LD_B_VX            = "LD_B_VX"
	LD_I_VX            = "LD_I_VX"
	LD_VX_I            = "LD_VX_I"
)

const (
	MASK_1 uint16 = 0xffff
	MASK_2        = 0xf000
	MASK_3        = 0xf00f
	MASK_4        = 0xf0ff
)

type ArgumentMask struct {
	Mask       uint16
	RightShift int
}

var MASK_NNN ArgumentMask = ArgumentMask{
	Mask:       0x0fff,
	RightShift: 0,
}

var MASK_N ArgumentMask = ArgumentMask{
	Mask:       0x000f,
	RightShift: 0,
}

var MASK_X ArgumentMask = ArgumentMask{
	Mask:       0x0fff,
	RightShift: 8,
}

var MASK_Y ArgumentMask = ArgumentMask{
	Mask:       0x00f0,
	RightShift: 4,
}

var MASK_KK ArgumentMask = ArgumentMask{
	Mask:       0x00ff,
	RightShift: 0,
}

type InstructionSet struct {
	Key           int
	Id            string
	Pattern       uint16
	Mask          uint16
	ArgumentMasks []ArgumentMask
}

var INSTRUCTION_SET []InstructionSet = []InstructionSet{
	InstructionSet{
		Key:           2,
		Id:            CLS,
		Pattern:       0x00e0,
		Mask:          MASK_1,
		ArgumentMasks: []ArgumentMask{},
	},
	InstructionSet{
		Key:           3,
		Id:            RET,
		Pattern:       0x00ee,
		Mask:          MASK_1,
		ArgumentMasks: []ArgumentMask{},
	},
	InstructionSet{
		Key:           4,
		Id:            JP_NNN,
		Pattern:       0x1000,
		Mask:          MASK_2,
		ArgumentMasks: []ArgumentMask{MASK_NNN},
	},
	InstructionSet{
		Key:           5,
		Id:            CALL_NNN,
		Pattern:       0x2000,
		Mask:          MASK_2,
		ArgumentMasks: []ArgumentMask{MASK_NNN},
	},
	InstructionSet{
		Key:           6,
		Id:            SE_VX_KK,
		Pattern:       0x3000,
		Mask:          MASK_2,
		ArgumentMasks: []ArgumentMask{MASK_X, MASK_KK},
	},
	InstructionSet{
		Key:           7,
		Id:            SNE_VX_KK,
		Pattern:       0x4000,
		Mask:          MASK_2,
		ArgumentMasks: []ArgumentMask{MASK_X, MASK_KK},
	},
	InstructionSet{
		Key:           8,
		Id:            SE_VX_VY,
		Pattern:       0x5000,
		Mask:          MASK_3,
		ArgumentMasks: []ArgumentMask{MASK_X, MASK_Y},
	},
	InstructionSet{
		Key:           9,
		Id:            LD_VX_KK,
		Pattern:       0x6000,
		Mask:          MASK_2,
		ArgumentMasks: []ArgumentMask{MASK_X, MASK_KK},
	},
	InstructionSet{
		Key:           10,
		Id:            ADD_VX_KK,
		Pattern:       0x7000,
		Mask:          MASK_2,
		ArgumentMasks: []ArgumentMask{MASK_X, MASK_KK},
	},
	InstructionSet{
		Key:           11,
		Id:            LD_VX_VY,
		Pattern:       0x8000,
		Mask:          MASK_3,
		ArgumentMasks: []ArgumentMask{MASK_X, MASK_Y},
	},
	InstructionSet{
		Key:           12,
		Id:            OR_VX_VY,
		Pattern:       0x8001,
		Mask:          MASK_3,
		ArgumentMasks: []ArgumentMask{MASK_X, MASK_Y},
	},
	InstructionSet{
		Key:           13,
		Id:            AND_VX_VY,
		Pattern:       0x8002,
		Mask:          MASK_3,
		ArgumentMasks: []ArgumentMask{MASK_X, MASK_Y},
	},
	InstructionSet{
		Key:           14,
		Id:            XOR_VX_VY,
		Pattern:       0x8003,
		Mask:          MASK_3,
		ArgumentMasks: []ArgumentMask{MASK_X, MASK_Y},
	},
	InstructionSet{
		Key:           15,
		Id:            ADD_VX_VY,
		Pattern:       0x8004,
		Mask:          MASK_3,
		ArgumentMasks: []ArgumentMask{MASK_X, MASK_Y},
	},
	InstructionSet{
		Key:           16,
		Id:            SUB_VX_VY,
		Pattern:       0x8005,
		Mask:          MASK_3,
		ArgumentMasks: []ArgumentMask{MASK_X, MASK_Y},
	},
	InstructionSet{
		Key:           17,
		Id:            SHR_VX_VY,
		Pattern:       0x8006,
		Mask:          MASK_3,
		ArgumentMasks: []ArgumentMask{MASK_X, MASK_Y},
	},
	InstructionSet{
		Key:           18,
		Id:            SUBN_VX_VY,
		Pattern:       0x8007,
		Mask:          MASK_3,
		ArgumentMasks: []ArgumentMask{MASK_X, MASK_Y},
	},
	InstructionSet{
		Key:           19,
		Id:            SHL_VX_VY,
		Pattern:       0x800e,
		Mask:          MASK_3,
		ArgumentMasks: []ArgumentMask{MASK_X, MASK_Y},
	},
	InstructionSet{
		Key:           20,
		Id:            SNE_VX_VY,
		Pattern:       0x9000,
		Mask:          MASK_3,
		ArgumentMasks: []ArgumentMask{MASK_X, MASK_Y},
	},
	InstructionSet{
		Key:           21,
		Id:            LD_I_NNN,
		Pattern:       0xa000,
		Mask:          MASK_2,
		ArgumentMasks: []ArgumentMask{MASK_NNN},
	},
	InstructionSet{
		Key:           22,
		Id:            JP_V0_NNN,
		Pattern:       0xb000,
		Mask:          MASK_2,
		ArgumentMasks: []ArgumentMask{MASK_NNN},
	},
	InstructionSet{
		Key:           23,
		Id:            RND_VX_KK,
		Pattern:       0xc000,
		Mask:          MASK_2,
		ArgumentMasks: []ArgumentMask{MASK_X, MASK_KK},
	},
	InstructionSet{
		Key:           24,
		Id:            DRW_VX_VY_N,
		Pattern:       0xd000,
		Mask:          MASK_2,
		ArgumentMasks: []ArgumentMask{MASK_X, MASK_Y, MASK_N},
	},
	InstructionSet{
		Key:           25,
		Id:            SKP_VX,
		Pattern:       0xe09e,
		Mask:          MASK_4,
		ArgumentMasks: []ArgumentMask{MASK_X},
	},
	InstructionSet{
		Key:           26,
		Id:            SKNP_VX,
		Pattern:       0xe0a1,
		Mask:          MASK_4,
		ArgumentMasks: []ArgumentMask{MASK_X},
	},
	InstructionSet{
		Key:           27,
		Id:            LD_VX_DT,
		Pattern:       0xf007,
		Mask:          MASK_3,
		ArgumentMasks: []ArgumentMask{MASK_X},
	},
	InstructionSet{
		Key:           28,
		Id:            LD_VX_K,
		Pattern:       0xf00a,
		Mask:          MASK_3,
		ArgumentMasks: []ArgumentMask{MASK_X},
	},
	InstructionSet{
		Key:           29,
		Id:            LD_DT_VX,
		Pattern:       0xf015,
		Mask:          MASK_4,
		ArgumentMasks: []ArgumentMask{MASK_X},
	},
	InstructionSet{
		Key:           30,
		Id:            LD_ST_VX,
		Pattern:       0xf018,
		Mask:          MASK_4,
		ArgumentMasks: []ArgumentMask{MASK_X},
	},
	InstructionSet{
		Key:           31,
		Id:            ADD_I_VX,
		Pattern:       0xf01e,
		Mask:          MASK_4,
		ArgumentMasks: []ArgumentMask{MASK_X},
	},
	InstructionSet{
		Key:           32,
		Id:            LD_F_VX,
		Pattern:       0xf029,
		Mask:          MASK_4,
		ArgumentMasks: []ArgumentMask{MASK_X},
	},
	InstructionSet{
		Key:           33,
		Id:            LD_B_VX,
		Pattern:       0xf033,
		Mask:          MASK_4,
		ArgumentMasks: []ArgumentMask{MASK_X},
	},
	InstructionSet{
		Key:           34,
		Id:            LD_I_VX,
		Pattern:       0xf055,
		Mask:          MASK_4,
		ArgumentMasks: []ArgumentMask{MASK_X},
	},
	InstructionSet{
		Key:           35,
		Id:            LD_VX_I,
		Pattern:       0xf065,
		Mask:          MASK_4,
		ArgumentMasks: []ArgumentMask{MASK_X},
	},
}

type Instruction struct {
	Instruction string
	Args        []uint16
}
