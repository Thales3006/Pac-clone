package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Label struct {
	Text string
	Size float32
}

func (l *Label) Use(bounds rl.Rectangle) {
	rl.DrawTextEx(rl.GetFontDefault(), l.Text, rl.Vector2{X: bounds.X, Y: bounds.Y}, l.Size, 5, rl.Black)
}
