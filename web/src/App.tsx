import { useState, useEffect, useRef } from 'react'

import { ROMS } from './constants'
import { Chip8 } from './chip8'
import { AppLayout } from "./components/AppLayout";
import { RomSelector } from "./components/RomSelector";
import './App.css'

function App() {
  const chip8Ref = useRef<Chip8 | null>(null);
  const [selectedRom, setSelectedRom] = useState(''); 

  useEffect(() => {
    async function initChip8() {
      const chip8 = new Chip8()
      await chip8.initialize()

      chip8Ref.current = chip8;
    }

    initChip8()
  }, [])

  async function handleOnRomSelect(romUrl: string): void {
    if (!chip8Ref.current) {
      return
    }

    const chip8 = chip8Ref.current;
    await chip8.loadRom(romUrl);
    setSelectedRom(romUrl);
  }

  async function handleRun(): void {
    if (!chip8Ref.current) {
      return
    }

    const chip8 = chip8Ref.current;
    await chip8.run();
  }

  return (
    <AppLayout>
      <div className="columns">
        <div className="column col-4">
          <RomSelector onSelect={handleOnRomSelect} />

          <button disabled={!selectedRom} onClick={handleRun}>Run</button>
        </div>

        <div className="column col-8">
          <div className="card">
            <div className="card-body">
              <canvas id="display"></canvas>
            </div>
          </div>
        </div>
      </div>
    </AppLayout>
  );
}

export default App
