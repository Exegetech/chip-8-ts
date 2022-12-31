import { useState } from "react";

import { ROMS } from '../constants'

export interface Props {
  onSelect: (rom: string) => void;
}

export function RomSelector(props: Props) {
  const [selected, setSelected] = useState('');

  function handleChange(event): void {
    setSelected(event.target.value);
    props.onSelect(event.target.value);
  }

  return (
    <div className="form-group mb-2">
      <select
        className="form-select"
        value={selected}
        onChange={handleChange}
      >
        <option value={''} key={'empty'}>
          Select a ROM
        </option>

        {ROMS.map((rom) => (
          <option value={rom} key={rom}>
            {rom}
          </option>
        ))}
      </select>
    </div>
  )
}
