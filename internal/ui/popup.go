package ui

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Pair struct {
	Button  string
	OnClick func()
}

type Popup struct {
	Title   string
	Message string
	OnClose func()
	Options []*Pair
}

func NewPopup(title string, message string, onClose func(), options []*Pair) *Popup {
	return &Popup{
		Title:   title,
		Message: message,
		OnClose: onClose,
		Options: options,
	}
}

func (p *Popup) Use(bounds rl.Rectangle) {
	buttons := ""
	for i, option := range p.Options {
		if i != 0 {
			buttons += ";"
		}
		buttons += option.Button
	}
	if val := gui.MessageBox(bounds, p.Title, p.Message, buttons) - 1; val != -2 {
		if val == -1 {
			p.OnClose()
		} else {
			p.Options[val].OnClick()
		}
	}
}
