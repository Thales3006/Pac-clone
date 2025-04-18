package game

import (
	"fmt"
	"pac-clone/internal/ui"

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

	elements := []ui.Element{
		&ui.Button{
			Text: "Play",
			OnClick: func() {
				g.currentScene = Level
			},
		},
		&ui.Button{
			Text: "Settings",
			OnClick: func() {
				fmt.Print("aaaa")
			},
		},
		&ui.Button{
			Text: "Exit",
			OnClick: func() {
				rl.CloseWindow()
			},
		},
	}

	g.component(rl.Rectangle{X: (float32(g.Width) - 300) / 2, Y: (float32(g.Height) - 200) / 2, Width: 300, Height: 200}, elements)
}

func (g *Game) HandleLevel() {
	rl.ClearBackground(rl.RayWhite)
	g.level.Render(rl.Rectangle{X: 200, Y: 200, Width: float32(g.Width), Height: float32(g.Height)})

	if rl.IsKeyDown(rl.KeyEscape) {
		g.currentScene = MainMenu
	}
}
