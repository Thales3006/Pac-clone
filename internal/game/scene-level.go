package game

import (
	"pac-clone/internal/entities"
	"pac-clone/internal/level"
	mv "pac-clone/internal/movement"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) HandleLevel() {
	rl.ClearBackground(rl.RayWhite)

	Draw(rl.Rectangle{
		X:      (float32(g.Width) - 800) / 2,
		Y:      (float32(g.Height) - 800) / 2,
		Width:  800,
		Height: 800,
	},
		g.Level,
		g.Player)

	if rl.IsKeyPressed(rl.KeyEscape) {
		g.currentScene = Pause
	}

	mv.HandleInput(g.Control, g.Player, g.Level)
	mv.UpdateEntity(&g.Player.Entity, g.Level, 1)
}

func Draw(bounds rl.Rectangle, l *level.Level, p *entities.Player) {
	rl.DrawRectangleRec(bounds, rl.Black)

	cellRect := rl.Rectangle{
		Width:  bounds.Width / float32(l.Width),
		Height: bounds.Height / float32(l.Height),
	}

	for i := range l.Grid {
		for j, cell := range l.Grid[i] {

			cellRect.X = bounds.X + float32(j)*cellRect.Width
			cellRect.Y = bounds.Y + float32(int(l.Height-1)-i)*cellRect.Height

			if cell == level.Wall {
				rl.DrawRectangleRec(cellRect, rl.Blue)
			}
		}
	}

	rl.DrawRectangleRec(rl.Rectangle{
		X:      bounds.X + cellRect.Width*p.X,
		Y:      bounds.Y + cellRect.Height*(float32(l.Height)-(p.Y+p.Height)),
		Width:  cellRect.Width * p.Width,
		Height: cellRect.Height * p.Height,
	},
		rl.Yellow)
}
