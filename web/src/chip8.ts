import { CLOCKS_PER_TIME_UNIT } from './constants'
import { sleep, fetchRom } from './utils'

export class Chip8 {
  private cpuReady: boolean
  private romIsLoaded: boolean
  private isRunning: boolean

  public constructor() {
    this.cpuReady = false;
    this.romIsLoaded = false;
    this.isRunning = false;
  }

  public async initialize(): Promise<void> {
    const go = new Go();
    const { instance } = await WebAssembly.instantiateStreaming(
      fetch('cpu.wasm'),
      go.importObject,
    )

    go.run(instance)
    this.ready = true;
  }

  public async reset(): Promise<void> {
    this.isRunning = false;
    // Flush event to avoid this.run method still executing 
    await sleep()

    this.romIsLoaded = false;

    cpu_reset();
  }

  public async loadRom(romUrl: string): Promise<void> {
    if (!this.ready) {
      throw new Error('Chip8 not ready')
    }

    await this.reset()

    if (!romUrl) {
      return;
    }

    const rom = await fetchRom(romUrl);
    cpu_loadRom(rom);

    this.romIsLoaded = true;
  }

  public async run(): Promise<void> {
    if (!this.ready) {
      throw new Error('Chip8 not ready')
    }

    if (!this.romIsLoaded) {
      throw new Error('No ROM loaded')
    }

    this.isRunning = true;

    while (this.isRunning) {
      await sleep()

      for (let i = 0; i < CLOCKS_PER_TIME_UNIT; ++i) {
        cpu_cycle()
      }
    }
  }
}
