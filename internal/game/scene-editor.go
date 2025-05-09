package game

import (
	"os"
	"pac-clone/internal/level"
	"pac-clone/internal/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Pick int32

const (
	Wall Pick = iota
	Point
	Door
	Player
	Ghost
)

var (
	editor_loaded     bool = false
	editor_listView   *ui.ListView
	editor_file_names []string
	editor_tool       Pick = Wall
)

func (g *Game) HandleEditor() {
	rl.ClearBackground(rl.RayWhite)

	if !editor_loaded {
		g.loadEditor()
	}

	if *editor_listView.ScrollIndex >= 0 && "levels/"+editor_file_names[*editor_listView.ScrollIndex] != g.Level.Current {
		g.Level.Unload()
		g.Level.Load("levels/" + editor_file_names[*editor_listView.ScrollIndex])
	}

	g.handleEditing()

	rl.DrawText(g.Level.Current, 150, 0, 30, rl.Black)
	editor_listView.Use(rl.Rectangle{X: 0, Y: 0, Width: 100, Height: 400})

	ui.NewComponent([]ui.Element{
		&ui.Button{
			Text: "Save",
			OnClick: func() {
				g.Level.Save(g.Level.Current)
			},
		},
		&ui.Button{
			Text: "New",
			OnClick: func() {
				g.Level.CreateFile()
				editor_loaded = false
			},
		},
	}).
		Use(rl.Rectangle{
			X:      float32(g.Width) - 200,
			Y:      float32(g.Height) - 100,
			Width:  200,
			Height: 100,
		})

	switch rl.GetKeyPressed() {
	case rl.KeyOne:
		editor_tool = Wall
	case rl.KeyTwo:
		editor_tool = Point
	case rl.KeyThree:
		editor_tool = Door
	case rl.KeyFour:
		editor_tool = Player
	case rl.KeyFive:
		editor_tool = Ghost

	case rl.KeyEscape:
		g.currentScene = MainMenu
		editor_loaded = false
	}

}

func (g *Game) loadEditor() {
	files, err := os.ReadDir("levels")
	if err != nil {
		ui.NewError(err.Error(), func() { g.currentScene = MainMenu })
		return
	}

	editor_file_names = []string{}
	for _, file := range files {
		editor_file_names = append(editor_file_names, file.Name())
	}
	editor_listView = ui.NewListView(editor_file_names)

	editor_loaded = true
}

func (g *Game) unloadEditor() {
	editor_file_names = []string{}
	g.Level.Unload()
	editor_listView = nil
	editor_loaded = false
}

func (g *Game) handleEditing() {
	bounds := g.center(600, 600)
	g.Draw(bounds, true)

	cell := rl.Rectangle{
		Width:  bounds.Width / float32(g.Level.Width),
		Height: bounds.Height / float32(g.Level.Height),
	}

	var currentText string

	switch editor_tool {
	case Wall:
		currentText = "1: Wall"
	case Point:
		currentText = "2: Point"
	case Door:
		currentText = "3: Door"
	case Player:
		currentText = "4: Player Spawn"
	case Ghost:
		currentText = "5: Ghosts Spawn"
	}
	rl.DrawText(currentText, int32(float32(g.Width)*0.2), int32(float32(g.Height)*0.1), 30, rl.Black)

	if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		pos := rl.GetMousePosition()
		pos.X -= bounds.X
		pos.Y -= bounds.Y

		if pos.X >= 0 && pos.Y >= 0 && pos.X < bounds.Width && pos.Y < bounds.Height {
			switch editor_tool {
			case Wall:
				g.Level.Grid[int32(pos.Y/cell.Height)][int32(pos.X/cell.Width)] = level.Wall
			case Point:
				g.Level.Grid[int32(pos.Y/cell.Height)][int32(pos.X/cell.Width)] = level.Point
			case Door:
				g.Level.Grid[int32(pos.Y/cell.Height)][int32(pos.X/cell.Width)] = level.Door
			case Player:
				if g.Level.Grid[int32(pos.Y/cell.Height)][int32(pos.X/cell.Width)] != level.Wall {
					g.Level.SpawnPlayer = [2]int32{int32(pos.Y / cell.Height), int32(pos.X / cell.Width)}
				}
			case Ghost:
				if g.Level.Grid[int32(pos.Y/cell.Height)][int32(pos.X/cell.Width)] != level.Wall {
					g.Level.SpawnGhost = [2]int32{int32(pos.Y / cell.Height), int32(pos.X / cell.Width)}
				}
			}
		}
	} else if rl.IsMouseButtonDown(rl.MouseButtonRight) {
		pos := rl.GetMousePosition()
		pos.X -= bounds.X
		pos.Y -= bounds.Y

		if pos.X >= 0 && pos.Y >= 0 && pos.X < bounds.Width && pos.Y < bounds.Height {
			g.Level.Grid[int32(pos.Y/cell.Height)][int32(pos.X/cell.Width)] = level.Empty
		}
	}

	g.ResetPositions()

}
