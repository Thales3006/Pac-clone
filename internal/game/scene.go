package game

import (
	"pac-clone/internal/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Scene uint8

const (
	MainMenu Scene = iota
	SelectionMenu
	Pause
	Settings
	Level
)

func (g *Game) HandleScene() {
	switch g.currentScene {
	case MainMenu:
		g.HandleMainMenu()
	case SelectionMenu:
		g.HandleSelectionMenu()
	case Level:
		g.HandleLevel()
	case Settings:
		g.HandleSettings()
	}
}

func (g *Game) HandleLevel() {
	rl.ClearBackground(rl.RayWhite)

	elements := []ui.Element{
		g.level,
	}

	g.component(rl.Rectangle{X: (float32(g.Width) - 800) / 2, Y: (float32(g.Height) - 800) / 2, Width: 800, Height: 800}, elements)

	if rl.IsKeyDown(rl.KeyEscape) {
		g.currentScene = MainMenu
	}
}
