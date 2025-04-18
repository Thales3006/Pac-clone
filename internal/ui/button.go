package ui

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
	Text   string
	Action func()
}

func (b *Button) Use(bounds rl.Rectangle) {
	if gui.Button(bounds, b.Text) {
		b.Action()
	}
}
