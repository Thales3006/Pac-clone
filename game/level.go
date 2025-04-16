package game

type Level struct {
	grid [][]uint8
}

type Side uint8

const (
	None  = 0
	Right = 1 << iota
	Left
	Up
	Down
	All = Right | Left | Up | Down
)

func LoadLevel() *Level {
	l := new(Level)
	l.grid = make([][]uint8, 28)

	for i := range l.grid {
		l.grid[i] = make([]uint8, 36)
		for j := range l.grid[i] {
			//totally arbitrary setup
			switch {
			case i == j:
				l.grid[i][j] = All
			case i == 0:
				l.grid[i][j] = Right | Left
			case j == 0:
				l.grid[i][j] = Up | Left

			}
		}
	}
	return l
}
