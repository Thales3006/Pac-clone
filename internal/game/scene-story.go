package game

import (
	"pac-clone/internal/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) HandleSettings() {
	rl.ClearBackground(rl.RayWhite)

	ui.NewComponent([]ui.Element{
		&ui.Label{
			Text: "Story",
			Size: 30,
		},
		&ui.Label{
			Text: "Pac-man is trying to make his city a better place \n" +
				"by cleaning the streets, but the ghosts \n" +
				"don't want that to happen. \n" +
				"Help Pac-man recycle all the trash.",
			Size: 20,
		},
		&ui.Button{
			Text: "Back",
			OnClick: func() {
				g.currentScene = MainMenu
			},
		},
	}).
		Use(g.center(600, 300))

	if rl.IsKeyPressed(rl.KeyEscape) {
		if g.Level.Current == "" {
			g.currentScene = MainMenu
		} else {
			g.currentScene = Pause
		}

	}
}
