package ui

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ListView struct {
	Options     []string
	Focus       *int32
	ScrollIndex *int32
	Active      int32
}

func NewListView(options []string) *ListView {
	return &ListView{
		Options:     options,
		Focus:       new(int32),
		ScrollIndex: new(int32),
		Active:      -1,
	}
}

func (l *ListView) Use(bounds rl.Rectangle) {
	l.Active = gui.ListViewEx(bounds, l.Options, l.Focus, l.ScrollIndex, l.Active)
}
