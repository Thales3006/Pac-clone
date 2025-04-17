package game

import (
	"fmt"
	ent "pac-clone/internal/entities"
	lvl "pac-clone/internal/level"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	level  *lvl.Level
	player *ent.Player
	ghosts []*ent.Ghost
}

func NewGame() *Game {
	return &Game{
		level:  lvl.LoadLevel(),
		player: ent.NewPlayer(),
		ghosts: []*ent.Ghost{ent.NewGhost()},
	}
}

func (g *Game) Run() {
	rl.InitWindow(900, 900, "Pac-clone")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		g.Render()
	}
}

func (g *Game) Close() {
	fmt.Println("Closing!")
}

func (g *Game) Update() {
	fmt.Println("Closing!")
}

func (g *Game) Render() {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	if (gui.Button(rl.Rectangle{X: 24, Y: 24, Width: 120, Height: 30}, "#191#Show Message")) {
		fmt.Println("fodase")
	}
	g.level.Render()

}

type Screen uint8

const (
	MainMenu Screen = iota
	ScreenMenu
	ScreenPause
	ScreenLevel
)

func (g *Game) RenderScreen(current Screen) {
	switch current {
	case ScreenLevel:
		g.level.Render()

	}
}
