package game

import (
	"pac-clone/internal/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) HandleSettings() {
	rl.ClearBackground(rl.RayWhite)

	elements := []ui.Element{
		&ui.Label{
			Text: "Settings",
		},
		&ui.Label{
			Text: "Not much to see here, still a Work in Progress (WIP)",
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
		if g.player.LevelCurrent == 0 {
			g.currentScene = MainMenu
		} else {
			g.currentScene = Level
		}

	}
}
