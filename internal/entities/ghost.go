package entities

import (
	"pac-clone/internal/utils"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

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
	Dead
	CScared
	Scared
)

type Ghost struct {
	Entity
	Personality Personality
	State       State
	Wait        utils.Timer
}

func NewGhost(personality Personality) *Ghost {
	timer := *utils.NewTimer(2 * time.Second)
	timer.Finish()
	return &Ghost{
		Entity: Entity{
			Rectangle: rl.Rectangle{
				X:      8,
				Y:      9,
				Width:  0.8,
				Height: 0.8,
			},
			Door:  true,
			Speed: 3,
		},
		Wait:        timer,
		Personality: personality,
	}
}
