package game

import (
	"pac-clone/internal/entities"
	"pac-clone/internal/level"
	mv "pac-clone/internal/movement"
	"pac-clone/internal/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) HandleLevel() {
	rl.ClearBackground(rl.RayWhite)

	if g.Level.Current == "" {
		if err := g.Level.Load("level1.json"); err != nil {
			Error := ui.NewError(err.Error(), func() { g.currentScene = MainMenu })
			Error.Use(g.center(400, 200))
		}
		return
	}

	Draw(rl.Rectangle{
		X:      (float32(g.Width) - 800) / 2,
		Y:      (float32(g.Height) - 800) / 2,
		Width:  800,
		Height: 800,
	},
		g.Level,
		g.Player)

	deltaTime := rl.GetFrameTime()
	mv.UpdateEntity(&g.Player.Entity, g.Level, mv.HandleInput(g.Control, g.Player, g.Level), deltaTime)
	mv.UpdateLevel(g.Level, &g.Player.Entity)

	if rl.IsKeyPressed(rl.KeyEscape) {
		g.currentScene = Pause
	}
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
			cellRect.Y = bounds.Y + float32(i)*cellRect.Height

			switch cell {
			case level.Wall:
				rl.DrawRectangleRec(cellRect, rl.Blue)
			case level.Point:
				rl.DrawRectangle(
					int32(cellRect.X+0.3*cellRect.Width),
					int32(cellRect.Y+0.3*cellRect.Height),
					int32(0.3*cellRect.Width),
					int32(0.3*cellRect.Height),
					rl.White)
			}
		}
	}

	rl.DrawRectangleRec(rl.Rectangle{
		X:      bounds.X + cellRect.Width*(p.X+(1-p.Width)/2),
		Y:      bounds.Y + cellRect.Height*(p.Y+(1-p.Height)/2),
		Width:  cellRect.Width * p.Width,
		Height: cellRect.Height * p.Height,
	},
		rl.Yellow)
}
