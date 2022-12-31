import { TIMER_60_HZ } from './constants'

export async function fetchRom(romUrl: string): Promise<Uint8Array> {
  const rom = await fetch(romUrl);
  const arrayBuffer = await rom.arrayBuffer();
  return new Uint8Array(arrayBuffer);
}

export async function sleep(ms = TIMER_60_HZ): Promise<null> {
  return new Promise((resolve, reject) => {
    setTimeout(resolve, ms);
  })
}

