package game

import (
	"pac-clone/internal/entities"
	"pac-clone/internal/level"
	mv "pac-clone/internal/movement"
	"pac-clone/internal/ui"
	"pac-clone/internal/utils"
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	level_loaded bool = false
	blinkyTex    rl.Texture2D
	pinkyTex     rl.Texture2D
	inkyTex      rl.Texture2D
	clideTex     rl.Texture2D
	start        *utils.Timer
)

func (g *Game) HandleLevel() {
	rl.ClearBackground(rl.RayWhite)

	g.Draw(g.center(800, 800), true)

	if g.Level.Points == 0 {
		ui.NewPopup("Congratulations!", "You Won \nthe Level!",
			func() { g.currentScene = SelectionMenu },
			[]*ui.Pair{{
				Button: "Select other Level",
				OnClick: func() {
					g.currentScene = SelectionMenu
					g.unloadLevel()
				}},
			}).Use(g.center(400, 200))
		return
	}

	if g.Player.Health == 0 {
		ui.NewPopup("Sorry!", "You Lost",
			func() { g.currentScene = SelectionMenu },
			[]*ui.Pair{
				{
					Button: "Select other Level",
					OnClick: func() {
						g.currentScene = SelectionMenu
						g.unloadLevel()
					}},
				{
					Button: "Main Menu",
					OnClick: func() {
						g.currentScene = MainMenu
						g.unloadLevel()
					}},
			}).Use(g.center(400, 200))
		return
	}

	if !start.Done() {
		rl.DrawText("Wait...", (g.Width-150)/2, 20, 30, rl.Gray)
		return
	}
	rl.DrawText("Play!", (g.Width-150)/2, 20, 30, rl.Green)

	if g.Player.IsDead {
		g.Player.IsDead = false
		g.Player.Health--

		g.ResetPositions()

		start.Reset()
	}

	deltaTime := rl.GetFrameTime()
	mv.HandleInput(g.Control, g.Player, g.Level)
	mv.HandleAI(g.Player, g.Ghosts, g.Level)
	mv.UpdateEntity(&g.Player.Entity, g.Level, deltaTime)

	for _, ghost := range g.Ghosts {
		mv.UpdateEntity(&ghost.Entity, g.Level, deltaTime)
	}
	mv.UpdateKill(g.Player, g.Ghosts, g.Level, deltaTime)

	mv.UpdateLevel(g.Level, g.Player, g.Ghosts)

	if rl.IsKeyPressed(rl.KeyEscape) {
		g.currentScene = Pause
	}
}

func (g *Game) loadLevel() {

	if g.Level.Current == "" {
		ui.NewError("Level was not found", func() { g.currentScene = MainMenu }).
			Use(g.center(400, 200))
		return
	}

	blinkyTex = rl.LoadTexture("assets/blinky.png")
	pinkyTex = rl.LoadTexture("assets/pinky.png")
	inkyTex = rl.LoadTexture("assets/inky.png")
	clideTex = rl.LoadTexture("assets/clyde.png")

	g.ResetPositions()
	g.Player.Direction = entities.None
	g.Player.Health = 5
	g.Player.IsDead = false

	level_loaded = true
	start = utils.NewTimer(1 * time.Second)
}

func (g *Game) unloadLevel() {
	g.Level.Unload()

	blinkyTex = rl.Texture2D{}
	pinkyTex = rl.Texture2D{}
	inkyTex = rl.Texture2D{}
	clideTex = rl.Texture2D{}

	g.Player.Direction = entities.None
	level_loaded = false
}

func (g *Game) Draw(bounds rl.Rectangle, drawEntities bool) {
	rl.DrawRectangleRec(bounds, rl.Black)

	if !level_loaded {
		g.loadLevel()
	}

	rl.DrawText("Score: "+strconv.Itoa(int(g.Player.Score)), 0, 0, 30, rl.Black)

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
			case level.Power:
				rl.DrawRectangle(
					int32(cellRect.X+0.2*cellRect.Width),
					int32(cellRect.Y+0.2*cellRect.Height),
					int32(0.5*cellRect.Width),
					int32(0.5*cellRect.Height),
					rl.Orange)
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

		if ghost.State == entities.Dead {
			rl.DrawRectangleRec(rl.Rectangle{
				X:      bounds.X + cellRect.Width*(ghost.X+(1-ghost.Width)/2),
				Y:      bounds.Y + cellRect.Height*(ghost.Y+(1-ghost.Height)/2),
				Width:  cellRect.Width * ghost.Width,
				Height: cellRect.Height * ghost.Height,
			},
				rl.White)
			continue
		}

		if ghost.State == entities.Scared {
			rl.DrawRectangleRec(rl.Rectangle{
				X:      bounds.X + cellRect.Width*(ghost.X+(1-ghost.Width)/2),
				Y:      bounds.Y + cellRect.Height*(ghost.Y+(1-ghost.Height)/2),
				Width:  cellRect.Width * ghost.Width,
				Height: cellRect.Height * ghost.Height,
			},
				rl.DarkBlue)
			continue
		}

		texture := &rl.Texture2D{}
		switch ghost.Personality {
		case entities.Blinky:
			texture = &blinkyTex
		case entities.Pinky:
			texture = &pinkyTex
		case entities.Inky:
			texture = &inkyTex
		case entities.Clyde:
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

func (g *Game) ResetPositions() {
	g.Player.X = float32(g.Level.SpawnPlayer[1])
	g.Player.Y = float32(g.Level.SpawnPlayer[0])

	for _, ghost := range g.Ghosts {
		ghost.X = float32(g.Level.SpawnGhost[1])
		ghost.Y = float32(g.Level.SpawnGhost[0])
	}
}
