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
}

func New() *Game {
	return &Game{
		level:        lvl.LoadLevel(),
		player:       ent.NewPlayer(),
		ghosts:       []*ent.Ghost{ent.NewGhost()},
		currentScene: Level,
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

func (g *Game) Render() {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	g.RenderScene()
}
