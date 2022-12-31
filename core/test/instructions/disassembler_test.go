package instructions_test

import (
	"testing"

	"github.com/exegetech/chip8/core/pkg/constants"
	"github.com/exegetech/chip8/core/pkg/instructions"
	"github.com/stretchr/testify/assert"
)

type scenario struct {
	opcode      uint16
	instruction constants.Instruction
}

var scenarios = map[string]scenario{
	"Disassembles CLS": scenario{
		opcode: 0x00e0,
		instruction: constants.Instruction{
			Instruction: constants.CLS,
			Args:        []uint16{},
		},
	},
	"Disassembles RET": scenario{
		opcode: 0x00ee,
		instruction: constants.Instruction{
			Instruction: constants.RET,
			Args:        []uint16{},
		},
	},
	"Disassembles JP_NNN": scenario{
		opcode: 0x1123,
		instruction: constants.Instruction{
			Instruction: constants.JP_NNN,
			Args:        []uint16{0x123},
		},
	},
	"Disassembles CALL_NNN": scenario{
		opcode: 0x2123,
		instruction: constants.Instruction{
			Instruction: constants.CALL_NNN,
			Args:        []uint16{0x123},
		},
	},
	"Disassembles SE_VX_KK": scenario{
		opcode: 0x3154,
		instruction: constants.Instruction{
			Instruction: constants.SE_VX_KK,
			Args:        []uint16{0x1, 0x54},
		},
	},
	"Disassembles SNE_VX_KK": scenario{
		opcode: 0x4154,
		instruction: constants.Instruction{
			Instruction: constants.SNE_VX_KK,
			Args:        []uint16{0x1, 0x54},
		},
	},
	"Disassembles SE_VX_VY": scenario{
		opcode: 0x5120,
		instruction: constants.Instruction{
			Instruction: constants.SE_VX_VY,
			Args:        []uint16{0x1, 0x2},
		},
	},
	"Disassembles LD_VX_KK": scenario{
		opcode: 0x6123,
		instruction: constants.Instruction{
			Instruction: constants.LD_VX_KK,
			Args:        []uint16{0x1, 0x23},
		},
	},
	"Disassembles ADD_VX_KK": scenario{
		opcode: 0x7123,
		instruction: constants.Instruction{
			Instruction: constants.ADD_VX_KK,
			Args:        []uint16{0x1, 0x23},
		},
	},
	"Disassembles LD_VX_VY": scenario{
		opcode: 0x8120,
		instruction: constants.Instruction{
			Instruction: constants.LD_VX_VY,
			Args:        []uint16{0x1, 0x2},
		},
	},
	"Disassembles OR_VX_VY": scenario{
		opcode: 0x8121,
		instruction: constants.Instruction{
			Instruction: constants.OR_VX_VY,
			Args:        []uint16{0x1, 0x2},
		},
	},
	"Disassembles AND_VX_VY": scenario{
		opcode: 0x8122,
		instruction: constants.Instruction{
			Instruction: constants.AND_VX_VY,
			Args:        []uint16{0x1, 0x2},
		},
	},
	"Disassembles XOR_VX_VY": scenario{
		opcode: 0x8123,
		instruction: constants.Instruction{
			Instruction: constants.XOR_VX_VY,
			Args:        []uint16{0x1, 0x2},
		},
	},
	"Disassembles ADD_VX_VY": scenario{
		opcode: 0x8124,
		instruction: constants.Instruction{
			Instruction: constants.ADD_VX_VY,
			Args:        []uint16{0x1, 0x2},
		},
	},
	"Disassembles SUB_VX_VY": scenario{
		opcode: 0x8125,
		instruction: constants.Instruction{
			Instruction: constants.SUB_VX_VY,
			Args:        []uint16{0x1, 0x2},
		},
	},
	"Disassembles SHR_VX_VY": scenario{
		opcode: 0x8126,
		instruction: constants.Instruction{
			Instruction: constants.SHR_VX_VY,
			Args:        []uint16{0x1, 0x2},
		},
	},
	"Disassembles SUBN_VX_VY": scenario{
		opcode: 0x8127,
		instruction: constants.Instruction{
			Instruction: constants.SUBN_VX_VY,
			Args:        []uint16{0x1, 0x2},
		},
	},
	"Disassembles SHL_VX_VY": scenario{
		opcode: 0x812e,
		instruction: constants.Instruction{
			Instruction: constants.SHL_VX_VY,
			Args:        []uint16{0x1, 0x2},
		},
	},
	"Disassembles SNE_VX_VY": scenario{
		opcode: 0x9120,
		instruction: constants.Instruction{
			Instruction: constants.SNE_VX_VY,
			Args:        []uint16{0x1, 0x2},
		},
	},
	"Disassembles LD_I_NNN": scenario{
		opcode: 0xa125,
		instruction: constants.Instruction{
			Instruction: constants.LD_I_NNN,
			Args:        []uint16{0x125},
		},
	},
	"Disassembles JP_V0_NNN": scenario{
		opcode: 0xb125,
		instruction: constants.Instruction{
			Instruction: constants.JP_V0_NNN,
			Args:        []uint16{0x125},
		},
	},
	"Disassembles RND_VX_KK": scenario{
		opcode: 0xc125,
		instruction: constants.Instruction{
			Instruction: constants.RND_VX_KK,
			Args:        []uint16{0x1, 0x25},
		},
	},
	"Disassembles DRW_VX_VY_N": scenario{
		opcode: 0xd125,
		instruction: constants.Instruction{
			Instruction: constants.DRW_VX_VY_N,
			Args:        []uint16{0x1, 0x2, 0x5},
		},
	},
	"Disassembles SKP_VX": scenario{
		opcode: 0xe19e,
		instruction: constants.Instruction{
			Instruction: constants.SKP_VX,
			Args:        []uint16{0x1},
		},
	},
	"Disassembles SKNP_VX": scenario{
		opcode: 0xe2a1,
		instruction: constants.Instruction{
			Instruction: constants.SKNP_VX,
			Args:        []uint16{0x2},
		},
	},
	"Disassembles LD_VX_DT": scenario{
		opcode: 0xf207,
		instruction: constants.Instruction{
			Instruction: constants.LD_VX_DT,
			Args:        []uint16{0x2},
		},
	},
	"Disassembles LD_VX_K": scenario{
		opcode: 0xf20a,
		instruction: constants.Instruction{
			Instruction: constants.LD_VX_K,
			Args:        []uint16{0x2},
		},
	},
	"Disassembles LD_DT_VX": scenario{
		opcode: 0xf215,
		instruction: constants.Instruction{
			Instruction: constants.LD_DT_VX,
			Args:        []uint16{0x2},
		},
	},
	"Disassembles LD_ST_VX": scenario{
		opcode: 0xf218,
		instruction: constants.Instruction{
			Instruction: constants.LD_ST_VX,
			Args:        []uint16{0x2},
		},
	},
	"Disassembles ADD_I_VX": scenario{
		opcode: 0xf21e,
		instruction: constants.Instruction{
			Instruction: constants.ADD_I_VX,
			Args:        []uint16{0x2},
		},
	},
	"Disassembles LD_F_VX": scenario{
		opcode: 0xf229,
		instruction: constants.Instruction{
			Instruction: constants.LD_F_VX,
			Args:        []uint16{0x2},
		},
	},
	"Disassembles LD_B_VX": scenario{
		opcode: 0xf233,
		instruction: constants.Instruction{
			Instruction: constants.LD_B_VX,
			Args:        []uint16{0x2},
		},
	},
	"Disassembles LD_I_VX": scenario{
		opcode: 0xf255,
		instruction: constants.Instruction{
			Instruction: constants.LD_I_VX,
			Args:        []uint16{0x2},
		},
	},
	"Disassembles LD_VX_I": scenario{
		opcode: 0xf265,
		instruction: constants.Instruction{
			Instruction: constants.LD_VX_I,
			Args:        []uint16{0x2},
		},
	},
}

func TestDisassemble(t *testing.T) {
	for name, scenario := range scenarios {
		t.Run(name, func(t *testing.T) {
			result := instructions.Disassemble(scenario.opcode)
			assert.Equal(t, result, scenario.instruction)
		})
	}
}
