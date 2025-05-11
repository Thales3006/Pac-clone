package level

func (l *Level) LoadDefault() {
	l.Grid = [][]Cell{
		{Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall},
		{Wall, Point, Point, Point, Point, Point, Point, Point, Point, Wall, Point, Point, Point, Point, Point, Point, Point, Point, Wall},
		{Wall, Point, Wall, Wall, Point, Wall, Wall, Wall, Point, Wall, Point, Wall, Wall, Wall, Point, Wall, Wall, Point, Wall},
		{Wall, Point, Point, Point, Point, Point, Point, Point, Point, Point, Point, Point, Point, Point, Point, Point, Point, Point, Wall},
		{Wall, Point, Wall, Wall, Point, Wall, Point, Wall, Wall, Wall, Wall, Wall, Point, Wall, Point, Wall, Wall, Point, Wall},
		{Wall, Point, Point, Point, Point, Wall, Point, Point, Point, Wall, Point, Point, Point, Wall, Point, Point, Point, Point, Wall},
		{Wall, Wall, Wall, Wall, Point, Wall, Wall, Wall, Empty, Wall, Empty, Wall, Wall, Wall, Point, Wall, Wall, Wall, Wall},
		{Empty, Empty, Empty, Wall, Point, Wall, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Wall, Point, Wall, Empty, Empty, Empty},
		{Wall, Wall, Wall, Wall, Point, Wall, Empty, Wall, Wall, Door, Wall, Wall, Empty, Wall, Point, Wall, Wall, Wall, Wall},
		{Point, Point, Point, Point, Point, Empty, Empty, Wall, Empty, Empty, Empty, Wall, Empty, Empty, Point, Point, Point, Point, Point},
		{Wall, Wall, Wall, Wall, Point, Wall, Empty, Wall, Wall, Wall, Wall, Wall, Empty, Wall, Point, Wall, Wall, Wall, Wall},
		{Empty, Empty, Empty, Wall, Point, Wall, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Wall, Point, Wall, Empty, Empty, Empty},
		{Wall, Wall, Wall, Wall, Point, Wall, Empty, Wall, Wall, Wall, Wall, Wall, Empty, Wall, Point, Wall, Wall, Wall, Wall},
		{Wall, Point, Point, Point, Point, Point, Point, Point, Point, Wall, Point, Point, Point, Point, Point, Point, Point, Point, Wall},
		{Wall, Point, Wall, Wall, Point, Wall, Wall, Wall, Point, Wall, Point, Wall, Wall, Wall, Point, Wall, Wall, Point, Wall},
		{Wall, Point, Point, Wall, Point, Point, Point, Point, Point, Point, Point, Point, Point, Point, Point, Wall, Point, Point, Wall},
		{Wall, Wall, Point, Wall, Point, Wall, Point, Wall, Wall, Wall, Wall, Wall, Point, Wall, Point, Wall, Point, Wall, Wall},
		{Wall, Point, Point, Point, Point, Wall, Point, Point, Point, Wall, Point, Point, Point, Wall, Point, Point, Point, Point, Wall},
		{Wall, Point, Wall, Wall, Wall, Wall, Wall, Wall, Point, Wall, Point, Wall, Wall, Wall, Wall, Wall, Wall, Point, Wall},
		{Wall, Point, Point, Point, Point, Point, Point, Point, Point, Point, Point, Point, Point, Point, Point, Point, Point, Point, Wall},
		{Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall},
	}
	l.Width = int32(len(l.Grid[0]))
	l.Height = int32(len(l.Grid))
	l.Current = "Default.json"
	l.SpawnGhost = [2]int32{8, 9}
	l.SpawnPlayer = [2]int32{1, 1}
}
