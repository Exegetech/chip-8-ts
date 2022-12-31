package registers

import (
	"github.com/exegetech/chip8/core/pkg/constants"
)

type Registers struct {
	V     [constants.NUMBER_OF_REGISTERS]uint8
	I     uint16
	ST    uint8
	DT    uint8
	PC    uint16
	SP    int8
	Stack [constants.STACK_DEEP]uint16
}

func New() *Registers {
	return &Registers{
		V:     [constants.NUMBER_OF_REGISTERS]uint8{},
		I:     0,
		ST:    0,
		DT:    0,
		PC:    constants.LOAD_PROGRAM_ADDRESS,
		SP:    -1,
		Stack: [constants.STACK_DEEP]uint16{},
	}
}

func (r *Registers) Reset() {
	for i, _ := range r.V {
		r.V[i] = 0
	}

	r.I = 0
	r.ST = 0
	r.DT = 0
	r.PC = constants.LOAD_PROGRAM_ADDRESS
	r.SP = -1

	for i, _ := range r.Stack {
		r.Stack[i] = 0
	}
}

func (r *Registers) StackPush(addr uint16) {
	r.SP += 1
	if r.SP >= 16 {
		panic("Stack overflow")
	}

	r.Stack[r.SP] = addr
}

func (r *Registers) StackPop() uint16 {
	value := r.Stack[r.SP]
	r.SP -= 1
	if r.SP < -1 {
		panic("Stack underflow")
	}

	return value
}

func (r *Registers) GetPC() uint16 {
	return r.PC
}

func (r *Registers) SetPC(pc uint16) {
	r.PC = pc
}

func (r *Registers) SkipPC() {
	r.PC += 2
}

func (r *Registers) SetV(idx int, val uint8) {
	if idx >= constants.NUMBER_OF_REGISTERS {
		panic("No such register")
	}

	r.V[idx] = val
}

func (r *Registers) GetV(idx int) uint8 {
	if idx >= constants.NUMBER_OF_REGISTERS {
		panic("No such register")
	}

	return r.V[idx]
}

func (r *Registers) SetI(val uint16) {
	r.I = val
}

func (r *Registers) GetI() uint16 {
	return r.I
}

func (r *Registers) SetDT(val uint8) {
	r.DT = val
}

func (r *Registers) GetDT() uint8 {
	return r.DT
}

func (r *Registers) SetST(val uint8) {
	r.ST = val
}

func (r *Registers) GetST() uint8 {
	return r.ST
}
