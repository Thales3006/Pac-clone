package game

import (
	"pac-clone/internal/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) HandleMainMenu() {
	rl.ClearBackground(rl.RayWhite)

	elements := []ui.Element{
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
	}

	g.component(rl.Rectangle{
		X:      (float32(g.Width) - 300) / 2,
		Y:      (float32(g.Height) - 200) / 2,
		Width:  300,
		Height: 200,
	},
		elements)
}

func (g *Game) HandleSelectionMenu() {
	rl.ClearBackground(rl.RayWhite)

	elements := []ui.Element{
		&ui.Label{
			Text: "Selection Menu",
		},
		&ui.Button{
			Text: "1",
			OnClick: func() {
				g.player.LevelCurrent = 1
				g.currentScene = Level
			},
		},
		&ui.Button{
			Text: "2",
			OnClick: func() {
				g.player.LevelCurrent = 2
				g.currentScene = Level
			},
		},
		&ui.Button{
			Text: "3",
			OnClick: func() {
				g.player.LevelCurrent = 3
				g.currentScene = Level
			},
		},
	}

	g.component(rl.Rectangle{X: (float32(g.Width) - 300) / 2, Y: (float32(g.Height) - 200) / 2, Width: 300, Height: 200}, elements)

	if rl.IsKeyPressed(rl.KeyEscape) {
		g.currentScene = MainMenu
	}
}
