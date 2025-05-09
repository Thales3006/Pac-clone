package game

import (
	"pac-clone/internal/entities"
	"pac-clone/internal/level"
	mv "pac-clone/internal/movement"
	"pac-clone/internal/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	level_loaded bool = false
	blinkyTex    rl.Texture2D
	pinkyTex     rl.Texture2D
	inkyTex      rl.Texture2D
	clideTex     rl.Texture2D
)

func (g *Game) HandleLevel() {
	rl.ClearBackground(rl.RayWhite)

	if g.Level.Current == "" {
		if err := g.Level.Load("level1.json"); err != nil {
			ui.NewError(err.Error(), func() { g.currentScene = MainMenu }).
				Use(g.center(400, 200))
		}
		return
	}

	g.Draw(g.center(800, 800), true)

	deltaTime := rl.GetFrameTime()
	mv.HandleInput(g.Control, g.Player, g.Level)
	mv.HandleAI(g.Player, g.Ghosts, g.Level)
	mv.UpdateEntity(&g.Player.Entity, g.Level, deltaTime)

	for _, ghost := range g.Ghosts {
		mv.UpdateEntity(&ghost.Entity, g.Level, deltaTime)
	}

	mv.UpdateLevel(g.Level, &g.Player.Entity)

	if rl.IsKeyPressed(rl.KeyEscape) {
		g.currentScene = Pause
	}
}

func (g *Game) loadLevel() {
	blinkyTex = rl.LoadTexture("assets/blinky.png")
	pinkyTex = rl.LoadTexture("assets/pinky.png")
	inkyTex = rl.LoadTexture("assets/inky.png")
	clideTex = rl.LoadTexture("assets/clyde.png")

	level_loaded = true
}

func (g *Game) Draw(bounds rl.Rectangle, drawEntities bool) {
	rl.DrawRectangleRec(bounds, rl.Black)

	if !level_loaded {
		g.loadLevel()
	}

	cellRect := rl.Rectangle{
		Width:  bounds.Width / float32(g.Level.Width),
		Height: bounds.Height / float32(g.Level.Height),
	}

	for i := range g.Level.Grid {
		for j, cell := range g.Level.Grid[i] {

			cellRect.X = bounds.X + float32(j)*cellRect.Width
			cellRect.Y = bounds.Y + float32(i)*cellRect.Height
			switch cell {
			case level.Wall:
				rl.DrawRectangleRec(cellRect, rl.Blue)
			case level.Door:
				rl.DrawRectangleRec(cellRect, rl.Gray)
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

	if !drawEntities {
		return
	}

	rl.DrawRectangleRec(rl.Rectangle{
		X:      bounds.X + cellRect.Width*(g.Player.X+(1-g.Player.Width)/2),
		Y:      bounds.Y + cellRect.Height*(g.Player.Y+(1-g.Player.Height)/2),
		Width:  cellRect.Width * g.Player.Width,
		Height: cellRect.Height * g.Player.Height,
	},
		rl.Yellow)

	for _, ghost := range g.Ghosts {
		texture := &rl.Texture2D{}
		switch ghost.Personality {
		case entities.Blinky:
			texture = &blinkyTex
		case entities.Pinky:
			texture = &pinkyTex
		case entities.Inky:
			texture = &inkyTex
		case entities.Clide:
			texture = &clideTex
		}

		texture.Width = int32(cellRect.Width * ghost.Width)
		texture.Height = int32(cellRect.Height * ghost.Height)

		srcRect := rl.NewRectangle(0, 0, float32(texture.Width), float32(texture.Height))
		position := rl.NewVector2(
			bounds.X+cellRect.Width*(ghost.X+(1-ghost.Width)/2),
			bounds.Y+cellRect.Height*(ghost.Y+(1-ghost.Height)/2),
		)
		rl.DrawTextureRec(*texture, srcRect, position, rl.White)
	}
}
