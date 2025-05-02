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

	selectionMenu := ui.NewComponent([]ui.Element{
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
	})

	selectionMenu.Use(rl.Rectangle{X: (float32(g.Width) - 300) / 2, Y: (float32(g.Height) - 200) / 2, Width: 300, Height: 200})

	if rl.IsKeyPressed(rl.KeyEscape) {
		g.currentScene = MainMenu
	}
}
