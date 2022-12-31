const NUMBER_OF_KEYS = 16;
const KEYMAP = [
  "1",
  "2",
  "3",
  "q",
  "w",
  "e",
  "a",
  "s",
  "d",
  "x",
  "z",
  "c",
  "4",
  "r",
  "f",
  "v",
];

export class Keyboard {
  private keys: boolean[];

  public constructor() {
    this.keys = new Array(NUMBER_OF_KEYS).fill(false);

    document.addEventListener("keydown", (event) => this.keydown(event));
    document.addEventListener("keyup", (event) => this.keyup(event));
  }

  private keydown(event: KeyboardEvent): void {
    const { key } = event;
    const keyIndex = KEYMAP.findIndex((mapKey) => mapKey === key.toLowerCase());

    if (keyIndex > -1) {
      this.keys[keyIndex] = true;
    }
  }

  private keyup(event: KeyboardEvent): void {
    const { key } = event;
    const keyIndex = KEYMAP.findIndex((mapKey) => mapKey === key.toLowerCase());

    if (keyIndex > -1) {
      this.keys[keyIndex] = false;
    }
  }

  public reset(): void {
    for (let i = 0; i < this.keys.length; i++) {
      this.keys[i] = false;
    }
  }

  public isKeyDown(keyIndex: number): boolean {
    return this.keys[keyIndex];
  }

  public hasKeyDown(): number {
    return this.keys.findIndex((key) => key);
  }
}
