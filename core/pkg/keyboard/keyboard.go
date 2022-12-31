package keyboard

import (
	"syscall/js"
)

type Keyboard struct {
	keyboard js.Value
}

func New() *Keyboard {
	keyboard := js.Global().Get("Keyboard").New()
	return &Keyboard{keyboard}
}

func (k *Keyboard) Reset() {
	k.keyboard.Call("reset")
}

func (k *Keyboard) IsKeyDown(idx int) bool {
	return k.keyboard.Call("isKeyDown", idx).Bool()
}

func (k *Keyboard) HasKeyDown() int {
	return k.keyboard.Call("hasKeyDown").Int()
}
