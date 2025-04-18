package ui

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
	Text   string
	Action func(bool)
}

func (b *Button) Use(bounds rl.Rectangle) {
	b.Action(gui.Button(bounds, b.Text))
}
