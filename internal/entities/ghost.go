package entities

import rl "github.com/gen2brain/raylib-go/raylib"

type Ghost struct {
	Entity
}

func NewGhost(xis float32, epi float32, larg float32, alt float32) *Ghost {
	return &Ghost{
		Entity: Entity{
			Rectangle: rl.Rectangle{
				X:      xis,
				Y:      epi,
				Width:  larg,
				Height: alt,
			},
			Speed: 3,
		},
	}
}
