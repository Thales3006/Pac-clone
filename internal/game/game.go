package game

import (
	ent "pac-clone/internal/entities"
	lvl "pac-clone/internal/level"

	mv "pac-clone/internal/movement"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Level         *lvl.Level
	Player        *ent.Player
	Ghosts        []*ent.Ghost
	currentScene  Scene
	isRunning     bool
	Width         int32
	Height        int32
	levelUnlocked int32
	Control       *mv.Control
}

func New() *Game {
	return &Game{
		Level:        lvl.NewLevel(),
		Player:       ent.NewPlayer(),
		Ghosts:       []*ent.Ghost{ent.NewGhost()},
		currentScene: MainMenu,
		isRunning:    true,
		Width:        int32(rl.GetScreenWidth()),
		Height:       int32(rl.GetScreenHeight()),
		Control:      mv.NewControl(),
	}
}

func (g *Game) Run() {
	//rl.SetConfigFlags(rl.FlagWindowResizable)

	rl.InitWindow(900, 900, "Pac-clone")
	defer rl.CloseWindow()

	rl.SetExitKey(rl.KeyNull)
	rl.SetTargetFPS(60)

	for g.isRunning && !rl.WindowShouldClose() {
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
