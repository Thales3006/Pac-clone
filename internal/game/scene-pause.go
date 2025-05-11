package game

import (
	"pac-clone/internal/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) HandlePause() {
	rl.ClearBackground(rl.RayWhite)

	ui.NewComponent([]ui.Element{
		&ui.Label{
			Text: "Pause",
			Size: 30,
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

				if level_loaded {
					g.unloadLevel()
				}
				if editor_loaded {
					g.unloadEditor()
				}
			},
		},
		&ui.Button{
			Text: "Exit",
			OnClick: func() {
				g.isRunning = false
			},
		},
	}).
		Use(rl.Rectangle{
			X:      (float32(g.Width) - 300) / 2,
			Y:      (float32(g.Height) - 200) / 2,
			Width:  300,
			Height: 200,
		})

	if rl.IsKeyPressed(rl.KeyEscape) {
		g.currentScene = Level
	}
}
