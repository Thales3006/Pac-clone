package game

import (
	ent "pac-clone/internal/entities"
	lvl "pac-clone/internal/level"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	level        *lvl.Level
	player       *ent.Player
	ghosts       []*ent.Ghost
	currentScene Scene
	Width        int32
	Height       int32
}

func New() *Game {
	return &Game{
		level:        lvl.LoadLevel(),
		player:       ent.NewPlayer(),
		ghosts:       []*ent.Ghost{ent.NewGhost()},
		currentScene: MainMenu,
		Width:        int32(rl.GetScreenWidth()),
		Height:       int32(rl.GetScreenHeight()),
	}
}

func (g *Game) Run() {
	//rl.SetConfigFlags(rl.FlagWindowResizable)

	rl.InitWindow(900, 900, "Pac-clone")
	defer rl.CloseWindow()

	rl.SetExitKey(rl.KeyNull)
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		g.HandleFrame()
	}
}

func (g *Game) HandleFrame() {
	g.Width = int32(rl.GetScreenWidth())
	g.Height = int32(rl.GetScreenHeight())

	rl.BeginDrawing()
	defer rl.EndDrawing()

	g.HandleScene()
}
