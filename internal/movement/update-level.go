package movement

import (
	"fmt"
	ent "pac-clone/internal/entities"
	"pac-clone/internal/level"
)

func UpdateLevel(l *level.Level, e *ent.Entity) {
	closest := closestCenter(e, l, 0)
	fmt.Println(l.Points)
	if l.Grid[closest[1]][closest[0]] == level.Point || l.Grid[closest[1]][closest[0]] == level.Power {
		l.Grid[closest[1]][closest[0]] = level.Empty
		l.Points--
	}
}
