package game

import (
	"fmt"
	"pac-clone/internal/ui"
	"path/filepath"
	"sort"
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
	rl.ClearBackground(backgroundColor)

	if !selectionMenu_loaded {
		g.loadSelectionMenu()
	}

	ui.NewComponent(levels).Use(g.center(300, 200))

	ui.NewComponent([]ui.Element{
		&ui.Label{
			Text: "Níveis desbloqueados: " + strconv.Itoa(int(g.levelUnlocked)),
			Size: 30,
		},
	}).
		Use(rl.Rectangle{
			X:      20,
			Y:      20,
			Width:  200,
			Height: 100,
		})

	if !custom {
		ui.NewComponent([]ui.Element{
			&ui.Button{
				Text: "Customizados",
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
	}

	if rl.IsKeyPressed(rl.KeyEscape) {
		g.currentScene = MainMenu
		custom = false
		selectionMenu_loaded = false
	}
}

func (g *Game) loadSelectionMenu() {
	levels = []ui.Element{}
	var pattern string

	if !custom {
		pattern = "levels/*.json"
	} else {
		pattern = "levels/custom/*.json"
	}

	files, err := filepath.Glob(pattern)
	if err != nil {
		fmt.Println("Erro ao procurar arquivos JSON:", err)
		return
	}

	sort.Strings(files)

	for i, filename := range files {
		index := i + 1
		currentFilename := filename

		if !custom {
			levels = append(levels, &ui.Button{
				Text: strconv.Itoa(index),
				OnClick: func() {
					g.Level.Load(currentFilename)
					if g.levelUnlocked >= g.Level.Required {
						g.currentScene = Level
					}
				},
			})
		} else {
			levels = append(levels, &ui.Button{
				Text: filepath.Base(filename),
				OnClick: func() {
					g.Level.Load(currentFilename)
					if g.levelUnlocked >= g.Level.Required {
						g.currentScene = Level
					}
				},
			})
		}
	}

	selectionMenu_loaded = true
}
