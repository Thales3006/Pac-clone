package movement

import (
	ent "pac-clone/internal/entities"
	"pac-clone/internal/level"
)

func UpdateLevel(l *level.Level, e *ent.Entity) {
	closest := closestCenter(e, l, 0)
	if l.Grid[closest[1]][closest[0]] == level.Point {
		l.Grid[closest[1]][closest[0]] = level.Empty
	}
}
