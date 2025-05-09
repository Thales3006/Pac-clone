package entities

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	Entity
	Health    uint8
	Powerfull bool
}

func NewPlayer() *Player {
	return &Player{
		Entity: Entity{
			Rectangle: rl.Rectangle{
				X:      1,
				Y:      1,
				Width:  1,
				Height: 1,
			},
			Door:  false,
			Speed: 3.5,
		},
		Health:    5,
		Powerfull: false,
	}
}
