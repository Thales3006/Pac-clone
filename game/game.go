package game

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Level  *Level
	Player *Player
	Ghosts []*Ghost
}

type Renderer interface {
	Render()
}

func New() *Game {
	return &Game{
		Level:  LoadLevel(),
		Player: NewPlayer(),
		Ghosts: []*Ghost{NewGhost()},
	}
}

func (g *Game) Run() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		g.Render()
	}
}

func (g *Game) Close() {
	fmt.Println("Closing!")
}
