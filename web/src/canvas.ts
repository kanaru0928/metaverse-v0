import { GameView } from "./game/gameView";

export function setupCanvas() {
  const canvas = document.getElementById("main-canvas");
  if (!(canvas instanceof HTMLCanvasElement)) {
    throw new Error("Canvas element not found");
  }
  const ctx = canvas.getContext("2d");
  if (!ctx) {
    throw new Error("Could not get canvas context");
  }

  const gameView = new GameView(ctx);
  gameView.init();
}
