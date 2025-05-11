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
	custom               bool         = false
	levels               []ui.Element = []ui.Element{
		&ui.Label{
			Text: "Selection Menu",
			Size: 30,
		},
	}
)

func (g *Game) HandleSelectionMenu() {
	rl.ClearBackground(rl.RayWhite)

	if !selectionMenu_loaded {
		g.loadSelectionMenu()
	}

	ui.NewComponent(levels).Use(g.center(300, 200))

	ui.NewComponent([]ui.Element{
		&ui.Button{
			Text: "Custom",
			OnClick: func() {
				custom = true
				selectionMenu_loaded = false
			},
		},
	}).
		Use(rl.Rectangle{
			X:      float32(g.Width) - 200,
			Y:      float32(g.Height) - 100,
			Width:  200,
			Height: 100,
		})

	if rl.IsKeyPressed(rl.KeyEscape) {
		g.currentScene = MainMenu
		custom = false
		selectionMenu_loaded = false
	}
}

func (g *Game) loadSelectionMenu() {
	levels = []ui.Element{}
	var counter int32 = 1
	for {

		var filename string
		if !custom {
			filename = fmt.Sprintf("%s%d%s", "levels/level", counter, ".json")
		} else {
			filename = fmt.Sprintf("%s%d%s", "levels/custom/level", counter, ".json")
		}
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
	selectionMenu_loaded = true
}
