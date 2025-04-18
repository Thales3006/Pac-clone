package ui

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Label struct {
	text string
}

func (l *Label) Use(bounds rl.Rectangle) {
	gui.Label(bounds, l.text)
}
