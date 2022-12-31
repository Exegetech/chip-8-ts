package instructions_test

import (
	"testing"

	"github.com/exegetech/chip8/core/pkg/instructions"
	"github.com/stretchr/testify/assert"
)

func TestCalculateOpCode(t *testing.T) {
	t.Run("Calculate OpCode", func(t *testing.T) {
		result := instructions.CalculateOpcode(
			0x12,
			0x4E,
		)

		expected := uint16(0x124e)

		assert.Equal(t, expected, result)
	})
}
