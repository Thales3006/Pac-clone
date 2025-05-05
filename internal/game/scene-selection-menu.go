package game

import (
	"fmt"
	"pac-clone/internal/ui"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) HandleSelectionMenu() {
	rl.ClearBackground(rl.RayWhite)

	onClick := func(level int32) {
		if level > g.levelUnlocked {
			fmt.Println("not unlocked")
			return
		}
		g.Level.Load("level" + strconv.Itoa(int(level)) + ".json")
		g.currentScene = Level
	}

	ui.NewComponent([]ui.Element{
		&ui.Label{
			Text: "Selection Menu",
		},
		&ui.Button{
			Text:    "1",
			OnClick: func() { onClick(1) },
		},
		&ui.Button{
			Text:    "2",
			OnClick: func() { onClick(2) },
		},
		&ui.Button{
			Text:    "3",
			OnClick: func() { onClick(3) },
		},
	}).
		Use(g.center(300, 200))

	if rl.IsKeyPressed(rl.KeyEscape) {
		g.currentScene = MainMenu
	}
}
