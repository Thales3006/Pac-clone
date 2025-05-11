package entities

import (
	"pac-clone/internal/utils"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Entity
	Health    uint8
	Powerfull utils.Timer
	IsDead    bool
	Score     int32
}

func NewPlayer() *Player {
	powerfull := *utils.NewTimer(10 * time.Second)
	powerfull.Finish()
	return &Player{
		Entity: Entity{
			Rectangle: rl.Rectangle{
				X:      1,
				Y:      1,
				Width:  0.7,
				Height: 0.7,
			},
			Door:  false,
			Speed: 3.5,
		},
		Health:    5,
		Powerfull: powerfull,
		IsDead:    false,
	}
}
