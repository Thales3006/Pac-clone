package movement

import (
	ent "pac-clone/internal/entities"
	"pac-clone/internal/level"
)

func UpdateKill(p *ent.Player, ghosts []*ent.Ghost, l *level.Level, delta float32) {
	for _, ghost := range ghosts {
		if colision(&p.Entity, &ghost.Entity) && ghost.State != ent.Dead {
			if ghost.State == ent.Scared {
				ghost.State = ent.Dead
				p.Score += 1000
			} else {
				p.IsDead = true
				p.Score -= 250
			}
		}
	}
}

func colision(a *ent.Entity, b *ent.Entity) bool {
	return a.X < b.X+b.Width && a.X+a.Width > b.X && a.Y < b.Y+b.Height && a.Y+a.Height > b.Y
}
