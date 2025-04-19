package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type Component struct {
	elements []Element
}

func NewComponent(elements []Element) *Component {
	return &Component{elements}
}

func (c *Component) Use(bounds rl.Rectangle) {
	cell := rl.Rectangle{
		X:      bounds.X,
		Y:      bounds.Y,
		Width:  bounds.Width,
		Height: bounds.Height / float32(len(c.elements)),
	}

	for i, element := range c.elements {
		cell.Y = bounds.Y + cell.Height*float32(i)
		element.Use(cell)
	}
}
