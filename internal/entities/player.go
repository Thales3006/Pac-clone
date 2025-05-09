package entities

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	Entity
	Health uint8
}

func NewPlayer() *Player {
	return &Player{
		Entity: Entity{
			Rectangle: rl.Rectangle{
				X:      1,
				Y:      1,
				Width:  0.5,
				Height: 0.5,
			},
			Speed: 4,
		},
		Health: 5,
	}
}
