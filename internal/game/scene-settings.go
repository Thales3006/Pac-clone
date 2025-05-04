package game

import (
	"pac-clone/internal/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) HandleSettings() {
	rl.ClearBackground(rl.RayWhite)

	ui.NewComponent([]ui.Element{
		&ui.Label{
			Text: "Settings",
		},
		&ui.Label{
			Text: "Not much to see here, still a Work in Progress (WIP)",
		},
	}).
		Use(g.center(300, 200))

	if rl.IsKeyPressed(rl.KeyEscape) {
		if g.Level.Current == "" {
			g.currentScene = MainMenu
		} else {
			g.currentScene = Pause
		}

	}
}
