package game

import (
	"os"
	"pac-clone/internal/level"
	"pac-clone/internal/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	editor_loaded     bool = false
	editor_listView   *ui.ListView
	editor_file_names []string
	editor_tool       level.Side = level.Wall
)

func (g *Game) HandleEditor() {
	rl.ClearBackground(rl.RayWhite)

	if !editor_loaded {
		g.loadEditor()
	}

	if *editor_listView.ScrollIndex >= 0 && editor_file_names[*editor_listView.ScrollIndex] != g.Level.Current {
		g.Level.Unload()
		g.Level.Load(editor_file_names[*editor_listView.ScrollIndex])
	}

	g.handleEditing()

	rl.DrawText(g.Level.Current, 150, 0, 30, rl.Black)
	editor_listView.Use(rl.Rectangle{X: 0, Y: 0, Width: 100, Height: 400})

	button := &ui.Button{
		Text: "Save",
		OnClick: func() {
			g.Level.Save(g.Level.Current)
		},
	}
	button.Use(rl.Rectangle{
		X:      float32(g.Width) - 200,
		Y:      float32(g.Height) - 100,
		Width:  200,
		Height: 100,
	})

	switch rl.GetKeyPressed() {
	case rl.KeyOne:
		editor_tool = level.Wall
	case rl.KeyTwo:
		editor_tool = level.Point
	case rl.KeyThree:
		editor_tool = level.Door

	case rl.KeyEscape:
		g.currentScene = MainMenu
		editor_loaded = false
	}

}

func (g *Game) loadEditor() {
	files, err := os.ReadDir("levels")
	if err != nil {

	}

	editor_file_names = []string{}
	for _, file := range files {
		editor_file_names = append(editor_file_names, file.Name())
	}
	editor_listView = ui.NewListView(editor_file_names)

	editor_loaded = true
}

func (g *Game) handleEditing() {
	bounds := g.center(600, 600)
	g.Draw(bounds)

	cell := rl.Rectangle{
		Width:  bounds.Width / float32(g.Level.Width),
		Height: bounds.Height / float32(g.Level.Height),
	}

	if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		pos := rl.GetMousePosition()
		pos.X -= bounds.X
		pos.Y -= bounds.Y

		if pos.X >= 0 && pos.Y >= 0 && pos.X < bounds.Width && pos.Y < bounds.Height {
			g.Level.Grid[int32(pos.Y/cell.Height)][int32(pos.X/cell.Width)] = editor_tool
		}
	} else if rl.IsMouseButtonDown(rl.MouseButtonRight) {
		pos := rl.GetMousePosition()
		pos.X -= bounds.X
		pos.Y -= bounds.Y

		if pos.X >= 0 && pos.Y >= 0 && pos.X < bounds.Width && pos.Y < bounds.Height {
			g.Level.Grid[int32(pos.Y/cell.Height)][int32(pos.X/cell.Width)] = level.Empty
		}
	}

}
