package level

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Level struct {
	Grid [][]Side
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
	return l
}

func (l *Level) Render() {
	for i := range l.Grid {
		for j := range l.Grid[i] {

			cell := l.Grid[i][j]

			switch cell {
			case None:
				rl.DrawRectangle(int32(j*32), 100+int32(i*32), 32, 32, rl.Black)

			default:
				if cell&Left == Left {
					rl.DrawRectangle(int32(j*32), 100+int32(i*32), 4, 32, rl.Green)
				}
				if cell&Right == Right {
					rl.DrawRectangle(int32(j*32)+28, 100+int32(i*32), 4, 32, rl.Orange)
				}
				if cell&Up == Up {
					rl.DrawRectangle(int32(j*32), 100+int32(i*32), 32, 4, rl.Purple)
				}
				if cell&Down == Down {
					rl.DrawRectangle(int32(j*32), 100+int32(i*32)+28, 32, 4, rl.Gray)
				}
			}
		}
	}
}
