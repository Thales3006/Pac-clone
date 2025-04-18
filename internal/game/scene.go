package game

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Scene uint8

const (
	MainMenu Scene = iota
	Pause
	Level
)

func (g *Game) HandleScene() {
	switch g.currentScene {
	case MainMenu:
		g.HandleMainMenu()
	case Level:
		g.HandleLevel()
	}
}

func (g *Game) HandleMainMenu() {
	rl.ClearBackground(rl.RayWhite)

	if gui.Button(rl.Rectangle{X: 24, Y: 84, Width: 120, Height: 30}, "Play") {
		g.currentScene = Level
	}
	if gui.Button(rl.Rectangle{X: 24, Y: 124, Width: 120, Height: 30}, "Settings") {

	}
	if gui.Button(rl.Rectangle{X: 24, Y: 164, Width: 120, Height: 30}, "Exit") {
		rl.CloseWindow()
	}
}

func (g *Game) HandleLevel() {
	rl.ClearBackground(rl.RayWhite)
	g.level.Render(rl.Rectangle{X: 200, Y: 200, Width: float32(g.Width), Height: float32(g.Height)})

	if rl.IsKeyDown(rl.KeyEscape) {
		g.currentScene = MainMenu
	}
}
