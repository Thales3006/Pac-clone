package game

import (
	"pac-clone/internal/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) center(width float32, height float32) rl.Rectangle {
	return ui.Center(rl.Rectangle{
		X:      float32(g.Width),
		Y:      float32(g.Height),
		Width:  width,
		Height: height,
	})
}
