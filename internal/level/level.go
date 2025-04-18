package level

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Level struct {
	Grid   [][]Side
	Width  uint8
	Height uint8
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
	l := new(Level)
	l.Grid = make([][]Side, 28)
	l.Height = uint8(len(l.Grid))

	for i := range l.Grid {
		l.Grid[i] = make([]Side, 36)
		for j := range l.Grid[i] {
			//totally arbitrary setup
			switch {
			case i == j:
				l.Grid[i][j] = All
			case i == 0:
				l.Grid[i][j] = Right | Left
			case j == 0:
				l.Grid[i][j] = Up | Left

			}
		}
	}
	l.Width = uint8(len(l.Grid[0]))
	return l
}

func (l *Level) Render(place rl.Rectangle) {
	place.Width = place.Width / float32(l.Width)
	place.Height = place.Height / float32(l.Height)

	const thickness = 0.2

	wall := rl.Rectangle{
		X:      place.Width * (1 - thickness),
		Y:      place.Height * (1 - thickness),
		Width:  place.Width * thickness,
		Height: place.Height * thickness,
	}

	cellRect := rl.Rectangle{
		Width:  place.Width,
		Height: place.Height,
	}

	for i := range l.Grid {
		for j := range l.Grid[i] {

			cell := l.Grid[i][j]
			cellRect.X = place.X + float32(j)*place.Width
			cellRect.Y = place.Y + float32(i)*place.Height

			switch cell {
			case None:
				rl.DrawRectangleRec(cellRect, rl.Black)

			default:
				if cell&Left == Left {
					rl.DrawRectangleRec(rl.Rectangle{X: cellRect.X, Y: cellRect.Y, Width: wall.Width, Height: cellRect.Height}, rl.Green)
				}
				if cell&Right == Right {
					rl.DrawRectangleRec(rl.Rectangle{X: cellRect.X + wall.X, Y: cellRect.Y, Width: wall.Width, Height: cellRect.Height}, rl.Orange)
				}
				if cell&Up == Up {
					rl.DrawRectangleRec(rl.Rectangle{X: cellRect.X, Y: cellRect.Y, Width: cellRect.Width, Height: wall.Height}, rl.Purple)
				}
				if cell&Down == Down {
					rl.DrawRectangleRec(rl.Rectangle{X: cellRect.X, Y: cellRect.Y + wall.Y, Width: cellRect.Width, Height: wall.Height}, rl.Gray)
				}
			}
		}
	}
}
