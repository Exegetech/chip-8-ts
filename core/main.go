package main

import (
	"syscall/js"

	"github.com/exegetech/chip8/core/pkg/cpu"
)

func main() {
	cpu := cpu.New()
	js.Global().Set("cpu_reset", resetWrapper(cpu))
	js.Global().Set("cpu_loadRom", loadRomWrapper(cpu))
	js.Global().Set("cpu_cycle", cycleWrapper(cpu))
	<-make(chan bool)
}

func resetWrapper(cpu *cpu.Cpu) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		cpu.Reset()
		return nil
	})
}

func loadRomWrapper(cpu *cpu.Cpu) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		data := js.Global().Get("Uint8Array").New(args[0])
		dst := make([]byte, data.Get("length").Int())
		js.CopyBytesToGo(dst, data)

		cpu.LoadRom(dst)
		return nil
	})
}

func cycleWrapper(cpu *cpu.Cpu) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		cpu.Cycle()
		return nil
	})
}
