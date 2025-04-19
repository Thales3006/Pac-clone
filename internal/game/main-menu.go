package game

import (
	"pac-clone/internal/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) HandleMainMenu() {
	rl.ClearBackground(rl.RayWhite)

	mainMenu := ui.NewComponent([]ui.Element{
		&ui.Label{
			Text: "Main Menu",
		},
		&ui.Button{
			Text: "Play",
			OnClick: func() {
				g.currentScene = SelectionMenu
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
	})

	mainMenu.Use(rl.Rectangle{
		X:      (float32(g.Width) - 300) / 2,
		Y:      (float32(g.Height) - 200) / 2,
		Width:  300,
		Height: 200,
	})
}
