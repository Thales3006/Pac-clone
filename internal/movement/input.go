package movement

import (
	ent "pac-clone/internal/entities"
	"pac-clone/internal/level"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func HandleInput(c *Control, p *ent.Player, l *level.Level) {
	Right := rl.IsKeyDown(c.bind[KeyRight])
	Left := rl.IsKeyDown(c.bind[KeyLeft])
	Up := rl.IsKeyDown(c.bind[KeyUp])
	Down := rl.IsKeyDown(c.bind[KeyDown])

	if Right && !Left && !Up && !Down {
		p.DesiredDir = ent.Right
	}
	if !Right && Left && !Up && !Down {
		p.DesiredDir = ent.Left
	}
	if !Right && !Left && Up && !Down {
		p.DesiredDir = ent.Up
	}
	if !Right && !Left && !Up && Down {
		p.DesiredDir = ent.Down
	}
}
