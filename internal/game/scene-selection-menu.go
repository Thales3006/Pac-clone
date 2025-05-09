package game

import (
	"fmt"
	"os"
	"pac-clone/internal/ui"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	selectionMenu_loaded bool         = false
	levels               []ui.Element = []ui.Element{
		&ui.Label{
			Text: "Selection Menu",
		},
	}
)

func (g *Game) HandleSelectionMenu() {
	rl.ClearBackground(rl.RayWhite)

	if !selectionMenu_loaded {
		g.loadSelectionMenu()
		selectionMenu_loaded = true
	}

	ui.NewComponent(levels).Use(g.center(300, 200))

	if rl.IsKeyPressed(rl.KeyEscape) {
		g.currentScene = MainMenu
	}
}

func (g *Game) loadSelectionMenu() {

	var counter int32 = 1
	for {
		filename := fmt.Sprintf("%s%d%s", "levels/level", counter, ".json")
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			break
		}
		levels = append(levels, &ui.Button{
			Text: strconv.Itoa(int(counter)),
			OnClick: func() {
				g.Level.Load(filename)
				g.currentScene = Level
			},
		})
		counter++
	}
}
