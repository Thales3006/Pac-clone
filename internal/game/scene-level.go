package game

import (
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
	level_loaded    bool = false
	blinkyTex       rl.Texture2D
	pinkyTex        rl.Texture2D
	inkyTex         rl.Texture2D
	clideTex        rl.Texture2D
	tile            rl.Texture2D
	start           *utils.Timer
	tileTableStreet = map[Conection]rl.Vector2{
		Horizontal: {X: 2, Y: 1},
		Vertical:   {X: 3, Y: 1},
		None:       {X: 4, Y: 1},
	}
	tileTable = map[Pick][]rl.Vector2{
		Wall:   {{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 3, Y: 0}, {X: 4, Y: 0}, {X: 5, Y: 0}, {X: 8, Y: 1}, {X: 6, Y: 0}, {X: 7, Y: 0}, {X: 8, Y: 0}, {X: 9, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}, {X: 8, Y: 1}},
		Point:  {{X: 0, Y: 2}, {X: 1, Y: 2}, {X: 2, Y: 2}, {X: 9, Y: 1}},
		Power:  {{X: 3, Y: 2}},
		Player: {{X: 3, Y: 2}, {X: 4, Y: 2}},
	}
	degree = map[entities.Direction]float32{
		entities.Right: 0,
		entities.Left:  180,
		entities.Up:    270,
		entities.Down:  90,
	}
)

type Conection int32

const (
	None       Conection = 0
	Horizontal Conection = iota
	Vertical
)

func (g *Game) HandleLevel() {
	rl.ClearBackground(rl.RayWhite)

	g.PlotMainTheme()

	g.Draw(g.center(float32(g.Height-150), float32(g.Height-150)))

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

	blinkyTex = rl.Texture2D{}
	pinkyTex = rl.Texture2D{}
	inkyTex = rl.Texture2D{}
	clideTex = rl.Texture2D{}

	g.Player.Direction = entities.None
	level_loaded = false
}

func (g *Game) Draw(bounds rl.Rectangle) {
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
				rl.DrawRectangle(
					int32(cellRect.X+0.2*cellRect.Width),
					int32(cellRect.Y+0.2*cellRect.Height),
					int32(0.5*cellRect.Width),
					int32(0.5*cellRect.Height),
					rl.Orange)
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

		srcRect := rl.NewRectangle(0, 0, float32(cellRect.Width*ghost.Width), float32(cellRect.Height*ghost.Height))
		position := rl.NewVector2(
			bounds.X+cellRect.Width*(ghost.X+(1-ghost.Width)/2),
			bounds.Y+cellRect.Height*(ghost.Y+(1-ghost.Height)/2),
		)
		texture.Width = int32(cellRect.Width * ghost.Width)
		texture.Height = int32(cellRect.Height * ghost.Height)
		rl.DrawTextureRec(*texture, srcRect, position, rl.White)
	}
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

func getConection(l *level.Level, i, j int32) Conection {
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
