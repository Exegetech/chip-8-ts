export class Sound {
  private soundEnabled: boolean;

  public constructor(volume: number) {
    this.soundEnabled = false;

    if ("AudioContext" in window) {
      const audioContext = new AudioContext();
      const masterGain = new GainNode(audioContext);

      masterGain.gain.value = volume;
      masterGain.connect(audioContext.destination);

      let soundEnabled = false;
      let oscillator: OscillatorNode;

      Object.defineProperties(this, {
        soundEnabled: {
          get: function () {
            return soundEnabled;
          },

          set: function (value) {
            if (value === soundEnabled) {
              return;
            }

            soundEnabled = value;
            if (soundEnabled) {
              oscillator = new OscillatorNode(audioContext, {
                type: "square",
              });

              oscillator.connect(masterGain);
              oscillator.start();
            } else {
              oscillator.stop();
            }
          },
        },
      });
    }
  }

  public enableSound(): void {
    this.soundEnabled = true;
  }

  public disableSound(): void {
    this.soundEnabled = false;
  }
}
