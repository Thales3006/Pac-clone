package game

type Level struct {
	Grid [][]uint8
}

func LoadLevel() *Level {
	l := new(Level)
	l.Grid = make([][]uint8, 28)
	for i := range l.Grid {
		l.Grid[i] = make([]uint8, 36)
	}

	//totally arbitrary testing
	for i := range l.Grid {
		for j := range l.Grid[i] {
			if i == j {
				l.Grid[i][j] = 1
			}
		}
	}
	return l
}
