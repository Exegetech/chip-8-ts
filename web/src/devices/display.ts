export class Display {
  private width: number
  private height: number
  private scaling: number
  private screen: HTMLCanvasElement;
  private context: CanvasRenderingContext2D;

  public constructor(width: number, height: number, scaling: number) {
    this.scaling = scaling

    this.screen = document.getElementById("display");
    this.screen.width = width * scaling;
    this.screen.height = height * scaling;

    this.context = this.screen.getContext("2d")!;
  }

  public drawPixel(h: number, w: number, color: string): void {
    this.context.fillStyle = color;
    this.context.fillRect(w * this.scaling, h * this.scaling, this.scaling, this.scaling);
  }
}
