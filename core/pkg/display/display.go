package display

import (
	"syscall/js"

	"github.com/exegetech/chip8/core/pkg/constants"
	"github.com/exegetech/chip8/core/pkg/memory"
)

type Display struct {
	screen      js.Value
	frameBuffer [][]uint8
	memory      *memory.Memory
}

func New(memory *memory.Memory) *Display {
	screen := js.Global().Get("Display").New(
		constants.DISPLAY_WIDTH,
		constants.DISPLAY_HEIGHT,
		constants.DISPLAY_SCALING,
	)

	frameBuffer := make([][]uint8, constants.DISPLAY_HEIGHT)
	for h := 0; h < constants.DISPLAY_HEIGHT; h++ {
		row := make([]uint8, constants.DISPLAY_WIDTH)
		frameBuffer[h] = row
	}

	display := &Display{
		screen:      screen,
		frameBuffer: frameBuffer,
		memory:      memory,
	}

	display.Reset()

	return display
}

func (d *Display) Reset() {
	for h := 0; h < constants.DISPLAY_HEIGHT; h++ {
		for w := 0; w < constants.DISPLAY_WIDTH; w++ {
			d.frameBuffer[h][w] = 0
		}
	}

	d.drawBuffer()
}

func (d *Display) DrawSprite(w, h int, spriteAddr uint16, spriteHeight int) bool {
	pixelCollision := false
	for lh := 0; lh < spriteHeight; lh++ {
		line := d.memory.GetMemoryAt(spriteAddr + uint16(lh))

		for lw := 0; lw < constants.CHAR_SET_WIDTH; lw++ {
			pixelIsOn := d.checkBitValueAt(line, lw)
			ph := (h + lh) % constants.DISPLAY_HEIGHT
			pw := (w + lw) % constants.DISPLAY_WIDTH

			if !pixelIsOn {
				continue
			}

			if d.frameBuffer[ph][pw] > 0 {
				pixelCollision = true
			}

			d.frameBuffer[ph][pw] ^= 1
		}
	}

	d.drawBuffer()

	return pixelCollision
}

func (d *Display) drawBuffer() {
	for h := 0; h < constants.DISPLAY_HEIGHT; h++ {
		for w := 0; w < constants.DISPLAY_WIDTH; w++ {
			d.drawPixel(h, w, d.frameBuffer[h][w])
		}
	}
}

func (d *Display) drawPixel(h, w int, on uint8) {
	color := constants.DISPLAY_BG_COLOR
	if on > 0 {
		color = constants.DISPLAY_COLOR
	}

	d.screen.Call("drawPixel", h, w, color)
}

func (d *Display) checkBitValueAt(bit uint8, position int) bool {
	mask := uint8(0b10000000) >> position
	masked := mask & bit
	if masked == 0 {
		return false
	}

	return true
}
