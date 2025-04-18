package ui

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
	text   string
	action func(bool)
}

func (b *Button) Use(bounds rl.Rectangle) {
	b.action(gui.Button(bounds, b.text))
}
