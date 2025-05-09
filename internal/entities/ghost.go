package entities

import rl "github.com/gen2brain/raylib-go/raylib"

type Personality uint8

const (
	Blinky Personality = iota
	Pinky
	Inky
	Clyde
)

type State uint8

const (
	Chase State = iota
	Scared
	Dead
	CScared
)

type Ghost struct {
	Entity
	Personality Personality
	State       State
}

func NewGhost(personality Personality) *Ghost {
	return &Ghost{
		Entity: Entity{
			Rectangle: rl.Rectangle{
				X:      8,
				Y:      9,
				Width:  1,
				Height: 1,
			},
			Speed: 3,
		},
		Personality: personality,
	}
}
