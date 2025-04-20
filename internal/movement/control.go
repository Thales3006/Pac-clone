package movement

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Key int

const (
	KeyRight Key = iota
	KeyLeft
	KeyUp
	KeyDown
)

type Control struct {
	bind map[Key]int32
}

func NewControl() *Control {
	return &Control{
		bind: map[Key]int32{
			KeyRight: rl.KeyD,
			KeyLeft:  rl.KeyA,
			KeyUp:    rl.KeyW,
			KeyDown:  rl.KeyS,
		},
	}
}

func (c *Control) ResetBind() {
	c.bind = map[Key]int32{
		KeyRight: rl.KeyD,
		KeyLeft:  rl.KeyA,
		KeyUp:    rl.KeyW,
		KeyDown:  rl.KeyS,
	}
}
