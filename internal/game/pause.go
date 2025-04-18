package game

import (
	"pac-clone/internal/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) HandlePause() {
	rl.ClearBackground(rl.RayWhite)

	elements := []ui.Element{
		&ui.Label{
			Text: "Pause",
		},
		&ui.Button{
			Text: "Continue",
			OnClick: func() {
				g.currentScene = Level
			},
		},
		&ui.Button{
			Text: "Main Menu",
			OnClick: func() {
				g.currentScene = MainMenu
			},
		},
		&ui.Button{
			Text: "Settings",
			OnClick: func() {
				g.currentScene = Settings
			},
		},
		&ui.Button{
			Text: "Exit",
			OnClick: func() {
				g.IsRunning = false
			},
		},
	}

	g.component(rl.Rectangle{
		X:      (float32(g.Width) - 300) / 2,
		Y:      (float32(g.Height) - 200) / 2,
		Width:  300,
		Height: 200,
	},
		elements)

	if rl.IsKeyPressed(rl.KeyEscape) {
		g.currentScene = Level
	}
}
