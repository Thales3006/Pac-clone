package ui

import rl "github.com/gen2brain/raylib-go/raylib"

func Center(p rl.Rectangle) rl.Rectangle {
	p.X = (p.X - p.Width) / 2
	p.Y = (p.Y - p.Height) / 2
	return p
}
