package movement

import (
	ent "pac-clone/internal/entities"
	"pac-clone/internal/level"
)

func UpdateLevel(l *level.Level, p *ent.Player, ghosts []*ent.Ghost) {
	closest := closestCenter(&p.Entity, l, 0)
	if l.Grid[closest[1]][closest[0]] == level.Point {
		l.Grid[closest[1]][closest[0]] = level.Empty
		l.Points--
		p.Score += 10
	}

	if l.Grid[closest[1]][closest[0]] == level.Power {
		l.Grid[closest[1]][closest[0]] = level.Empty
		l.Points--
		p.Score += 100
		p.Powerfull.Reset()
		for _, ghost := range ghosts {
			if ghost.State != ent.Dead && ghost.Wait.Done() {
				ghost.State = ent.Scared
			}
		}
	}
}
