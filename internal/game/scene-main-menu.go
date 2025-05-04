package game

import (
	"pac-clone/internal/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) HandleMainMenu() {
	rl.ClearBackground(rl.RayWhite)

	ui.NewComponent([]ui.Element{
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
			Text: "Editor",
			OnClick: func() {
				g.currentScene = Editor
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
				g.isRunning = false
			},
		},
	}).
		Use(g.center(300, 300))
}
