package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Element interface {
	Use(rl.Rectangle)
}
