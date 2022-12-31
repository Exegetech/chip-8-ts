package cpu

import (
	"fmt"
	"strconv"
	"time"
	"github.com/exegetech/chip8/core/pkg/constants"
	"github.com/exegetech/chip8/core/pkg/display"
	"github.com/exegetech/chip8/core/pkg/instructions"
	"github.com/exegetech/chip8/core/pkg/keyboard"
	"github.com/exegetech/chip8/core/pkg/memory"
	"github.com/exegetech/chip8/core/pkg/registers"
	"github.com/exegetech/chip8/core/pkg/sound"
	"math/rand"
)

const DEBUG = false

type Cpu struct {
	memory    *memory.Memory
	registers *registers.Registers
	display   *display.Display
	sound     *sound.Sound
	keyboard  *keyboard.Keyboard
}

func New() *Cpu {
	memory := memory.New()
	memory.FillMemory(constants.CHAR_SET_ADDRESS, constants.CHAR_SET_SPRITES)

	registers := registers.New()

	display := display.New(memory)
	sound := sound.New()

	keyboard := keyboard.New()

	return &Cpu{
		memory:    memory,
		registers: registers,
		display:   display,
		sound:     sound,
		keyboard:  keyboard,
	}
}

func (c *Cpu) Reset() {
	c.sound.Reset()
	c.keyboard.Reset()
	c.display.Reset()

	c.registers.Reset()
	c.memory.Reset()
	c.memory.FillMemory(constants.CHAR_SET_ADDRESS, constants.CHAR_SET_SPRITES)
}

func (c *Cpu) LoadRom(rom []uint8) {
	c.memory.FillMemory(constants.LOAD_PROGRAM_ADDRESS, rom)
}

func (c *Cpu) Cycle() {
	currInstructionAddr := c.registers.GetPC()
	firstByte := c.memory.GetMemoryAt(currInstructionAddr)
	secondByte := c.memory.GetMemoryAt(currInstructionAddr + 1)
	opcode := instructions.CalculateOpcode(firstByte, secondByte)

	if DEBUG {
		fmt.Println("----------------------------")
		fmt.Println("PC", strconv.FormatInt(int64(currInstructionAddr-0x200), 16))
	}

	c.execute(opcode)
	c.updateTimers()
}

func (c *Cpu) execute(opcode uint16) {
	disassembled := instructions.Disassemble(opcode)
	instruction := disassembled.Instruction
	args := disassembled.Args

	if DEBUG {
		fmt.Println("OPCODE", strconv.FormatInt(int64(opcode), 16))
		fmt.Println("INSTRUCTION", instruction)
		fmt.Println("ARGS", args)
	}

	c.registers.SkipPC()

	switch instruction {
	case constants.CLS:
		c.execute_CLS()

	case constants.RET:
		c.execute_RET()

	case constants.JP_NNN:
		c.execute_JP_NNN(args[0])

	case constants.CALL_NNN:
		c.execute_CALL_NNN(args[0])

	case constants.SE_VX_KK:
		c.execute_SE_VX_KK(int(args[0]), uint8(args[1]))

	case constants.SNE_VX_KK:
		c.execute_SNE_VX_KK(int(args[0]), uint8(args[1]))

	case constants.SE_VX_VY:
		c.execute_SE_VX_VY(int(args[0]), int(args[1]))

	case constants.LD_VX_KK:
		c.execute_LD_VX_KK(int(args[0]), uint8(args[1]))

	case constants.ADD_VX_KK:
		c.execute_ADD_VX_KK(int(args[0]), uint8(args[1]))

	case constants.LD_VX_VY:
		c.execute_LD_VX_VY(int(args[0]), int(args[1]))

	case constants.OR_VX_VY:
		c.execute_OR_VX_VY(int(args[0]), int(args[1]))

	case constants.AND_VX_VY:
		c.execute_AND_VX_VY(int(args[0]), int(args[1]))

	case constants.XOR_VX_VY:
		c.execute_XOR_VX_VY(int(args[0]), int(args[1]))

	case constants.ADD_VX_VY:
		c.execute_ADD_VX_VY(int(args[0]), int(args[1]))

	case constants.SUB_VX_VY:
		c.execute_SUB_VX_VY(int(args[0]), int(args[1]))

	case constants.SHR_VX_VY:
		c.execute_SHR_VX_VY(int(args[0]), int(args[1]))

	case constants.SUBN_VX_VY:
		c.execute_SUBN_VX_VY(int(args[0]), int(args[1]))

	case constants.SHL_VX_VY:
		c.execute_SHL_VX_VY(int(args[0]), int(args[1]))

	case constants.SNE_VX_VY:
		c.execute_SNE_VX_VY(int(args[0]), int(args[1]))

	case constants.LD_I_NNN:
		c.execute_LD_I_NNN(uint16(args[0]))

	case constants.JP_V0_NNN:
		c.execute_JP_V0_NNN(uint16(args[0]))

	case constants.RND_VX_KK:
		c.execute_RND_VX_KK(int(args[0]), uint8(args[1]))

	case constants.DRW_VX_VY_N:
		c.execute_DRW_VX_VY_N(int(args[0]), int(args[1]), int(args[2]))

	case constants.SKP_VX:
		c.execute_SKP_VX(int(args[0]))

	case constants.SKNP_VX:
		c.execute_SKNP_VX(int(args[0]))

	case constants.LD_VX_DT:
		c.execute_LD_VX_DT(int(args[0]))

	case constants.LD_VX_K:
		c.execute_LD_VX_K(int(args[0]))

	case constants.LD_DT_VX:
		c.execute_LD_DT_VX(int(args[0]))

	case constants.LD_ST_VX:
		c.execute_LD_ST_VX(int(args[0]))

	case constants.ADD_I_VX:
		c.execute_ADD_I_VX(int(args[0]))

	case constants.LD_F_VX:
		c.execute_LD_F_VX(int(args[0]))

	case constants.LD_B_VX:
		c.execute_LD_B_VX(int(args[0]))

	case constants.LD_I_VX:
		c.execute_LD_I_VX(int(args[0]))

	case constants.LD_VX_I:
		c.execute_LD_VX_I(int(args[0]))

	default:
		panic("Instruction not found")
	}
}

func (c *Cpu) execute_CLS() {
	c.display.Reset()
}

func (c *Cpu) execute_RET() {
	addr := c.registers.StackPop()
	c.registers.SetPC(addr)
}

func (c *Cpu) execute_JP_NNN(nnn uint16) {
	c.registers.SetPC(nnn)
}

func (c *Cpu) execute_CALL_NNN(nnn uint16) {
	addr := c.registers.GetPC()
	c.registers.StackPush(addr)

	c.registers.SetPC(nnn)
}

func (c *Cpu) execute_SE_VX_KK(x int, kk uint8) {
	if c.registers.GetV(x) == kk {
		c.registers.SkipPC()
	}
}

func (c *Cpu) execute_SNE_VX_KK(x int, kk uint8) {
	if c.registers.GetV(x) != kk {
		c.registers.SkipPC()
	}
}

func (c *Cpu) execute_SE_VX_VY(x, y int) {
	if c.registers.GetV(x) == c.registers.GetV(y) {
		c.registers.SkipPC()
	}
}

func (c *Cpu) execute_LD_VX_KK(x int, kk uint8) {
	c.registers.SetV(x, kk)
}

func (c *Cpu) execute_ADD_VX_KK(x int, kk uint8) {
	vx := c.registers.GetV(x)
	c.registers.SetV(x, vx+kk)
}

func (c *Cpu) execute_LD_VX_VY(x, y int) {
	vy := c.registers.GetV(y)
	c.registers.SetV(x, vy)
}

func (c *Cpu) execute_OR_VX_VY(x, y int) {
	vx := c.registers.GetV(x)
	vy := c.registers.GetV(y)
	result := vx | vy

	c.registers.SetV(x, result)
}

func (c *Cpu) execute_AND_VX_VY(x, y int) {
	vx := c.registers.GetV(x)
	vy := c.registers.GetV(y)
	result := vx & vy

	c.registers.SetV(x, result)
}

func (c *Cpu) execute_XOR_VX_VY(x, y int) {
	vx := c.registers.GetV(x)
	vy := c.registers.GetV(y)
	result := vx ^ vy

	c.registers.SetV(x, result)
}

func (c *Cpu) execute_ADD_VX_VY(x, y int) {
	vx := c.registers.GetV(x)
	vy := c.registers.GetV(y)
	result := vx + vy

	if result > 0xff {
		c.registers.SetV(0xf, 1)
	} else {
		c.registers.SetV(0xf, 0)
	}

	c.registers.SetV(
		x,
		// Registers are uint8 already so
		// it will automatially wrap around
		result,
	)
}

func (c *Cpu) execute_SUB_VX_VY(x, y int) {
	vx := c.registers.GetV(x)
	vy := c.registers.GetV(y)

	if vx > vy {
		c.registers.SetV(0xf, 1)
	} else {
		c.registers.SetV(0xf, 0)
	}

	result := vx - vy
	c.registers.SetV(x, result)
}

func (c *Cpu) execute_SHR_VX_VY(x, _y int) {
	vf := c.registers.GetV(0xf)
	c.registers.SetV(0xf, vf&0b00000001)

	vx := c.registers.GetV(x)
	c.registers.SetV(x, vx>>1)
}

func (c *Cpu) execute_SUBN_VX_VY(x, y int) {
	vx := c.registers.GetV(x)
	vy := c.registers.GetV(y)

	if vy > vx {
		c.registers.SetV(0xf, 1)
	} else {
		c.registers.SetV(0xf, 0)
	}

	result := vy - vx
	c.registers.SetV(x, result)
}

func (c *Cpu) execute_SHL_VX_VY(x, _y int) {
	vf := c.registers.GetV(0xf)
	c.registers.SetV(0xf, vf&0b10000000)

	vx := c.registers.GetV(x)
	c.registers.SetV(x, vx<<1)
}

func (c *Cpu) execute_SNE_VX_VY(x, y int) {
	if c.registers.GetV(x) != c.registers.GetV(y) {
		c.registers.SkipPC()
	}
}

func (c *Cpu) execute_LD_I_NNN(nnn uint16) {
	c.registers.SetI(nnn)
}

func (c *Cpu) execute_JP_V0_NNN(nnn uint16) {
	v0 := c.registers.GetV(0)
	c.registers.SetPC(uint16(v0) + nnn)
}

func (c *Cpu) execute_RND_VX_KK(x int, kk uint8) {
	min := 0x0
	max := 0xff
	random := rand.Intn(max-min) + min
	result := uint8(random) & kk
	c.registers.SetV(x, result)
}

func (c *Cpu) execute_DRW_VX_VY_N(x, y, n int) {
	i := c.registers.GetI()
	vx := c.registers.GetV(x)
	vy := c.registers.GetV(y)

	collission := c.display.DrawSprite(int(vx), int(vy), i, n)
	if collission {
		c.registers.SetV(0xf, 1)
	} else {
		c.registers.SetV(0xf, 0)
	}
}

func (c *Cpu) execute_SKP_VX(x int) {
	fmt.Println("IsKeyDown", x)
	if c.keyboard.IsKeyDown(x) {
	  fmt.Println("Key is down", x)
		c.registers.SkipPC()
	}
}

func (c *Cpu) execute_SKNP_VX(x int) {
	// fmt.Println("NOT IsKeyDown", x)
	if !c.keyboard.IsKeyDown(x) {
	  // fmt.Println("Key is NOT down", x)
		c.registers.SkipPC()
	}
}

func (c *Cpu) execute_LD_VX_DT(x int) {
	dt := c.registers.GetDT()
	c.registers.SetV(x, dt)
}

func (c *Cpu) execute_LD_VX_K(x int) {
	ch := make(chan int)
	go func() {
		keyPressedIdx := -1
		for keyPressedIdx == -1 {
			keyPressedIdx = c.keyboard.HasKeyDown()
			time.Sleep(5 * time.Millisecond)
		}

		ch <- keyPressedIdx
	}()

	keyIdx := <-ch

	fmt.Println("Keypress", keyIdx)
	c.registers.SetV(x, uint8(keyIdx))
}

func (c *Cpu) execute_LD_DT_VX(x int) {
	vx := c.registers.GetV(x)
	c.registers.SetDT(vx)
}

func (c *Cpu) execute_LD_ST_VX(x int) {
	vx := c.registers.GetV(x)
	c.registers.SetST(vx)
}

func (c *Cpu) execute_ADD_I_VX(x int) {
	i := c.registers.GetI()
	vx := c.registers.GetV(x)
	result := i + uint16(vx)
	c.registers.SetI(result)
}

func (c *Cpu) execute_LD_F_VX(x int) {
	location := x * constants.CHAR_SET_HEIGHT
	c.registers.SetI(uint16(location))
}

func (c *Cpu) execute_LD_B_VX(x int) {
	vx := c.registers.GetV(x)
	hundreds := vx / 100
	tens := (vx % 100) / 10
	ones := (vx % 100) % 10

	i := c.registers.GetI()

	c.memory.SetMemoryAt(i, hundreds)
	c.memory.SetMemoryAt(i+1, tens)
	c.memory.SetMemoryAt(i+2, ones)
}

func (c *Cpu) execute_LD_I_VX(x int) {
	i := c.registers.GetI()

	for k := 0; k <= x; k++ {
		val := c.registers.GetV(k)
		c.memory.SetMemoryAt(i+uint16(k), val)
	}
}

func (c *Cpu) execute_LD_VX_I(x int) {
	i := c.registers.GetI()

	for k := 0; k <= x; k++ {
		val := c.memory.GetMemoryAt(i + uint16(k))
		c.registers.SetV(k, val)
	}
}

func (c *Cpu) updateTimers() {
	dt := c.registers.GetDT()
	if dt > 0 {
		c.registers.SetDT(dt - 1)
	}

	st := c.registers.GetST()
	if st > 0 {
		c.sound.EnableSound()
		c.registers.SetST(st - 1)
	}

	st = c.registers.GetST()
	if st == 0 {
		c.sound.DisableSound()
	}
}
