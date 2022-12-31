package instructions

import (
	"golang.org/x/exp/slices"

	"github.com/exegetech/chip8/core/pkg/constants"
)

func Disassemble(opcode uint16) constants.Instruction {
	instructionIdx := slices.IndexFunc(constants.INSTRUCTION_SET, func(i constants.InstructionSet) bool {
		masked := uint16(opcode) & i.Mask
		return masked == i.Pattern
	})

	instruction := constants.INSTRUCTION_SET[instructionIdx]

	args := make([]uint16, len(instruction.ArgumentMasks))
	for i, argMask := range instruction.ArgumentMasks {
		masked := uint16(opcode) & argMask.Mask
		arg := masked >> argMask.RightShift
		args[i] = arg
	}

	return constants.Instruction{
		Instruction: instruction.Id,
		Args:        args,
	}
}
