package game

import rl "github.com/gen2brain/raylib-go/raylib"

func (g *Game) HandleEditor() {
	rl.ClearBackground(rl.RayWhite)

	if rl.IsKeyPressed(rl.KeyEscape) {
		g.currentScene = MainMenu
	}
}
