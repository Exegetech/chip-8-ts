package memory

import (
	"github.com/exegetech/chip8/core/pkg/constants"
)

type Memory struct {
	memory [constants.MEMORY_SIZE]uint8
}

func New() *Memory {
	memory := [constants.MEMORY_SIZE]uint8{}
	return &Memory{memory: memory}
}

func (m *Memory) Reset() {
	for i, _ := range m.memory {
		m.memory[i] = uint8(0)
	}
}

func (m *Memory) FillMemory(addr uint16, data []uint8) {
	for i, d := range data {
		m.memory[int(addr)+i] = d
	}
}

func (m *Memory) GetMemoryAt(addr uint16) uint8 {
	idx := int(addr)
	if idx >= len(m.memory) {
		panic("Memory out of bound")
	}

	return m.memory[idx]
}

func (m *Memory) SetMemoryAt(addr uint16, data uint8) {
	idx := int(addr)
	if idx >= len(m.memory) {
		panic("Memory out of bound")
	}

	m.memory[idx] = data
}
