package level

import (
	"pac-clone/internal/entities"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Level struct {
	Grid     [][]Side
	Width    uint8
	Height   uint8
	Unlocked uint8
	Current  uint8
}

type Side uint8

const (
	None  Side = 0
	Right Side = 1 << iota
	Left
	Up
	Down
	All = Right | Left | Up | Down
)

func LoadLevel() *Level {
	grid := make([][]Side, 20)

	for i := range grid {
		grid[i] = make([]Side, 20)
		for j := range grid[i] {
			//totally arbitrary setup
			switch {
			case i == j:
				grid[i][j] = All
			case i == 0:
				grid[i][j] = Right | Left
			case j == 0:
				grid[i][j] = Up | Left

			}
		}
	}

	return &Level{
		Grid:     grid,
		Width:    uint8(len(grid[0])),
		Height:   uint8(len(grid)),
		Unlocked: 1,
		Current:  0,
	}
}

func (l *Level) Draw(bounds rl.Rectangle, p *entities.Player) {
	rl.DrawRectangleRec(bounds, rl.Black)

	bounds.Width = bounds.Width / float32(l.Width)
	bounds.Height = bounds.Height / float32(l.Height)

	const thickness = 0.2

	wall := rl.Rectangle{
		X:      bounds.Width * (1 - thickness),
		Y:      bounds.Height * (1 - thickness),
		Width:  bounds.Width * thickness,
		Height: bounds.Height * thickness,
	}

	cellRect := rl.Rectangle{
		Width:  bounds.Width,
		Height: bounds.Height,
	}

	for i := range l.Grid {
		for j := range l.Grid[i] {

			cell := l.Grid[i][j]
			cellRect.X = bounds.X + float32(j)*bounds.Width
			cellRect.Y = bounds.Y + float32(int(l.Height)-(i+1))*bounds.Height

			if cell&Left == Left {
				rl.DrawRectangleRec(rl.Rectangle{
					X:      cellRect.X,
					Y:      cellRect.Y,
					Width:  wall.Width,
					Height: cellRect.Height,
				},
					rl.Green)
			}
			if cell&Right == Right {
				rl.DrawRectangleRec(rl.Rectangle{
					X:      cellRect.X + wall.X,
					Y:      cellRect.Y,
					Width:  wall.Width,
					Height: cellRect.Height,
				},
					rl.Orange)
			}
			if cell&Up == Up {
				rl.DrawRectangleRec(rl.Rectangle{
					X:      cellRect.X,
					Y:      cellRect.Y,
					Width:  cellRect.Width,
					Height: wall.Height,
				},
					rl.Purple)
			}
			if cell&Down == Down {
				rl.DrawRectangleRec(rl.Rectangle{
					X:      cellRect.X,
					Y:      cellRect.Y + wall.Y,
					Width:  cellRect.Width,
					Height: wall.Height,
				},
					rl.Gray)
			}
		}
	}

	rl.DrawRectangleRec(rl.Rectangle{
		X:      bounds.X + bounds.Width*p.X,
		Y:      bounds.Y + bounds.Height*(float32(l.Height)-(p.Y+1)),
		Width:  bounds.Width * p.Width,
		Height: bounds.Height * p.Height,
	},
		rl.Yellow)
}
