package game

import (
	"pac-clone/internal/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) component(container rl.Rectangle, elements []ui.Element) {
	cell := rl.Rectangle{
		X:      container.X,
		Y:      container.Y,
		Width:  container.Width,
		Height: container.Height / float32(len(elements)),
	}

	for i, element := range elements {

		cell.Y = container.Y + cell.Height*float32(i)
		element.Use(cell)
	}
}
