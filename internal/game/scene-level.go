package game

import (
	"image/color"
	"math"
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
	tile         rl.Texture2D
	start        *utils.Timer
)

type Connection int32

const (
	None       Connection = 0
	Horizontal Connection = iota
	Vertical
)

func (g *Game) HandleLevel() {
	rl.ClearBackground(backgroundColor)

	g.PlotMainTheme()

	g.Draw(g.center(float32(g.Height-150), float32(g.Height-150)))

	if g.Level.Points == 0 {
		ui.NewPopup("Parabéns!", "You Venceu!",
			func() { g.currentScene = SelectionMenu },
			[]*ui.Pair{{
				Button: "Selecione outro nível",
				OnClick: func() {
					g.currentScene = SelectionMenu
					g.unloadLevel()
					if g.Level.Required == g.levelUnlocked {
						g.levelUnlocked++
					}
				}},
			}).Use(g.center(400, 200))
		return
	}

	if g.Player.Health == 0 {
		ui.NewPopup("Desculpa!", "Você perdeu!",
			func() { g.currentScene = SelectionMenu },
			[]*ui.Pair{
				{
					Button: "Selecione outro nível",
					OnClick: func() {
						g.currentScene = SelectionMenu
						g.unloadLevel()
					}},
				{
					Button: "Menu Principal",
					OnClick: func() {
						g.currentScene = MainMenu
						g.unloadLevel()
					}},
			}).Use(g.center(400, 200))
		return
	}

	if !start.Done() {
		rl.DrawText("Esperar...", (g.Width-150)/2, 0, 30, rl.Gray)
		return
	}
	rl.DrawText("Jogar!", (g.Width-150)/2, 0, 30, rl.Green)

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

	tile = rl.LoadTexture("assets/tilemap.png")

	g.ResetPositions()
	g.Player.Direction = entities.None
	g.Player.Health = 5
	g.Player.Score = 0
	g.Player.IsDead = false

	level_loaded = true
	start = utils.NewTimer(1 * time.Second)
}

func (g *Game) unloadLevel() {
	g.Level.Unload()

	g.Player.Direction = entities.None
	level_loaded = false
}

func (g *Game) Draw(bounds rl.Rectangle) {
	rl.DrawRectangleRec(bounds, rl.Black)

	if !level_loaded {
		g.loadLevel()
	}

	rl.DrawText("Pontuação: "+strconv.Itoa(int(g.Player.Score)), 0, 0, 30, rl.Black)
	rl.DrawText("Vidas: "+strconv.Itoa(int(g.Player.Health)), 0, 40, 30, rl.Black)

	cellRect := rl.Rectangle{
		Width:  bounds.Width / float32(g.Level.Width),
		Height: bounds.Height / float32(g.Level.Height),
	}

	for i := range g.Level.Grid {
		for j, cell := range g.Level.Grid[i] {

			cellRect.X = bounds.X + float32(j)*cellRect.Width
			cellRect.Y = bounds.Y + float32(i)*cellRect.Height
			switch {
			case cell == level.Wall:
				DrawTile(tile, tileTable[Wall][mv.Mod(int32(i)+int32(j), int32(len(tileTable[Wall])))], cellRect)
			case cell == level.Door:
				rl.DrawRectangleRec(cellRect, rl.Gray)
			default:
				DrawTile(tile, tileTableStreet[getConection(g.Level, int32(i), int32(j))], cellRect)

			}

			switch cell {
			case level.Point:
				DrawTile(tile, tileTable[Point][mv.Mod(int32(i)+int32(j), int32(len(tileTable[Point])))], rl.Rectangle{
					X:      cellRect.X + (0.25+float32(math.Cos(rl.GetTime())/20))*cellRect.Width,
					Y:      cellRect.Y + (0.25+float32(math.Sin(rl.GetTime())/20))*cellRect.Height,
					Width:  0.5 * cellRect.Width,
					Height: 0.5 * cellRect.Height,
				})

			case level.Power:
				DrawTile(tile, rl.Vector2{X: 9, Y: 2}, rl.Rectangle{
					X:      cellRect.X + (float32(math.Cos(rl.GetTime())/20))*cellRect.Width,
					Y:      cellRect.Y + (float32(math.Sin(rl.GetTime())/20))*cellRect.Height,
					Width:  cellRect.Width,
					Height: cellRect.Height,
				})
			}

		}

	}

	DrawTileRot(tile, tileTable[Player][int(rl.GetTime()*4)%len(tileTable[Player])], rl.Rectangle{
		X:      bounds.X + cellRect.Width*(g.Player.X+0.5),
		Y:      bounds.Y + cellRect.Height*(g.Player.Y+0.5),
		Width:  cellRect.Width * g.Player.Width,
		Height: cellRect.Width * g.Player.Height,
	}, degree[g.Player.Direction])

	for _, ghost := range g.Ghosts {
		pos := tileGhost[ghost.Personality][ghost.State][ghost.Direction]
		if len(pos) == 0 {
			continue
		}
		DrawTile(tile, pos[int(rl.GetTime()*20)%len(pos)],
			rl.Rectangle{
				X:      bounds.X + cellRect.Width*(ghost.X+(1-ghost.Width)/2),
				Y:      bounds.Y + cellRect.Height*(ghost.Y+(1-ghost.Height)/2),
				Width:  float32(cellRect.Width * ghost.Width),
				Height: float32(cellRect.Height * ghost.Height),
			})
	}

	borderColor := color.RGBA{19, 74, 39, 255}
	rl.DrawRectangleRec(rl.Rectangle{
		X:      bounds.X - cellRect.Width,
		Y:      bounds.Y,
		Width:  cellRect.Width,
		Height: bounds.Height + cellRect.Height,
	}, borderColor)

	rl.DrawRectangleRec(rl.Rectangle{
		X:      bounds.X + bounds.Width,
		Y:      bounds.Y,
		Width:  cellRect.Width,
		Height: bounds.Height + cellRect.Height,
	}, borderColor)

	rl.DrawRectangleRec(rl.Rectangle{
		X:      bounds.X - cellRect.Width,
		Y:      bounds.Y - cellRect.Height,
		Width:  bounds.Width + cellRect.Width*2,
		Height: cellRect.Height,
	}, borderColor)

	rl.DrawRectangleRec(rl.Rectangle{
		X:      bounds.X,
		Y:      bounds.Y + bounds.Height,
		Width:  bounds.Width,
		Height: cellRect.Height,
	}, borderColor)
}

func (g *Game) ResetPositions() {
	g.Player.X = float32(g.Level.SpawnPlayer[1])
	g.Player.Y = float32(g.Level.SpawnPlayer[0])

	for _, ghost := range g.Ghosts {
		ghost.X = float32(g.Level.SpawnGhost[1])
		ghost.Y = float32(g.Level.SpawnGhost[0])
		ghost.State = entities.Chase
	}
}

func DrawTile(tileTexture rl.Texture2D, pos rl.Vector2, bounds rl.Rectangle) {
	const tileSize = 32

	source := rl.NewRectangle(
		float32(pos.X*tileSize),
		float32(pos.Y*tileSize),
		tileSize, tileSize,
	)
	rl.DrawTexturePro(tileTexture, source, bounds, rl.NewVector2(0, 0), 0, rl.White)
}

func DrawTileRot(tileTexture rl.Texture2D, pos rl.Vector2, bounds rl.Rectangle, degree float32) {
	const tileSize = 32

	source := rl.NewRectangle(
		float32(pos.X*tileSize),
		float32(pos.Y*tileSize),
		tileSize, tileSize,
	)
	rl.DrawTexturePro(tileTexture, source, bounds, rl.Vector2{X: bounds.Width * 0.5, Y: bounds.Height * 0.5}, degree, rl.White)
}

func getConection(l *level.Level, i, j int32) Connection {
	conection := None
	if c := l.Grid[mv.Mod(i-1, l.Height)][j]; c != level.Wall && c != level.Door {
		conection = Vertical
	} else if c := l.Grid[mv.Mod(i+1, l.Height)][j]; c != level.Wall && c != level.Door {
		conection = Vertical
	}

	if c := l.Grid[i][mv.Mod(j-1, l.Width)]; c != level.Wall && c != level.Door {
		if conection == Vertical {
			conection = None
		} else {
			conection = Horizontal
		}
	} else if c := l.Grid[i][mv.Mod(j+1, l.Width)]; c != level.Wall && c != level.Door {
		if conection == Vertical {
			conection = None
		} else {
			conection = Horizontal
		}
	}
	return conection
}
