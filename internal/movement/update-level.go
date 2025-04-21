package movement

import (
	ent "pac-clone/internal/entities"
	"pac-clone/internal/level"
)

func UpdateLevel(l *level.Level, e *ent.Entity) {
	if l.Grid[int32(e.Y)][int32(e.X)] == level.Point {
		l.Grid[int32(e.Y)][int32(e.X)] = level.Empty
	}
}
