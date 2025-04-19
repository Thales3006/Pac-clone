package game

import (
	"pac-clone/internal/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) HandleSelectionMenu() {
	rl.ClearBackground(rl.RayWhite)

	selectionMenu := ui.NewComponent([]ui.Element{
		&ui.Label{
			Text: "Selection Menu",
		},
		&ui.Button{
			Text: "1",
			OnClick: func() {
				g.LevelCurrent = 1
				g.currentScene = Level
			},
		},
		&ui.Button{
			Text: "2",
			OnClick: func() {
				g.LevelCurrent = 2
				g.currentScene = Level
			},
		},
		&ui.Button{
			Text: "3",
			OnClick: func() {
				g.LevelCurrent = 3
				g.currentScene = Level
			},
		},
	})

	selectionMenu.Use(rl.Rectangle{X: (float32(g.Width) - 300) / 2, Y: (float32(g.Height) - 200) / 2, Width: 300, Height: 200})

	if rl.IsKeyPressed(rl.KeyEscape) {
		g.currentScene = MainMenu
	}
}
