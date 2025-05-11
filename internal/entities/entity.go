package entities

import rl "github.com/gen2brain/raylib-go/raylib"

type Direction int8

const (
	None Direction = iota
	Up
	Down
	Left
	Right
)

type Entity struct {
	rl.Rectangle
	Speed      float32
	Direction  Direction
	DesiredDir Direction
	Door       bool
}
