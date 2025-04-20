package game

import (
	"pac-clone/internal/level"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Key int

const (
	KeyRight Key = iota
	KeyLeft
	KeyUp
	KeyDown
)

func newControl() map[Key]int32 {
	return map[Key]int32{
		KeyRight: rl.KeyD,
		KeyLeft:  rl.KeyA,
		KeyUp:    rl.KeyW,
		KeyDown:  rl.KeyS,
	}
}

func (g *Game) Movement() {
	cell := g.Level.Grid[int32(g.Player.Y)][int32(g.Player.X)]

	if rl.IsKeyDown(g.Control[KeyRight]) && cell&level.Right != level.Right {
		g.Player.X += g.Player.Vel
	}
	if rl.IsKeyDown(g.Control[KeyLeft]) && cell&level.Left != level.Left {
		g.Player.X -= g.Player.Vel
	}
	if rl.IsKeyDown(g.Control[KeyUp]) && cell&level.Up != level.Up {
		g.Player.Y += g.Player.Vel
	}
	if rl.IsKeyDown(g.Control[KeyDown]) && cell&level.Down != level.Down {
		g.Player.Y -= g.Player.Vel
	}
}
