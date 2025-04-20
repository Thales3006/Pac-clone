package game

import (
	"pac-clone/internal/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) HandleSettings() {
	rl.ClearBackground(rl.RayWhite)

	settings := ui.NewComponent([]ui.Element{
		&ui.Label{
			Text: "Settings",
		},
		&ui.Label{
			Text: "Not much to see here, still a Work in Progress (WIP)",
		},
	})

	settings.Use(rl.Rectangle{
		X:      (float32(g.Width) - 300) / 2,
		Y:      (float32(g.Height) - 200) / 2,
		Width:  300,
		Height: 200,
	})

	if rl.IsKeyPressed(rl.KeyEscape) {
		if g.level.Current == 0 {
			g.currentScene = MainMenu
		} else {
			g.currentScene = Pause
		}

	}
}
