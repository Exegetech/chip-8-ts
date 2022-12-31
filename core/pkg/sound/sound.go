package sound

import (
	"syscall/js"

	"github.com/exegetech/chip8/core/pkg/constants"
)

type Sound struct {
	audio js.Value
}

func New() *Sound {
	audio := js.Global().Get("Sound").New(constants.AUDIO_VOLUME)
	return &Sound{audio}
}

func (s *Sound) Reset() {
	s.DisableSound()
}

func (s *Sound) EnableSound() {
	s.audio.Call("enableSound")
}

func (s *Sound) DisableSound() {
	s.audio.Call("disableSound")
}
