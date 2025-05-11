package movement

import (
	ent "pac-clone/internal/entities"
	"pac-clone/internal/level"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var dirVectors = map[ent.Direction]rl.Vector2{
	ent.None:  {X: 0, Y: 0},
	ent.Up:    {X: 0, Y: -1},
	ent.Down:  {X: 0, Y: 1},
	ent.Left:  {X: -1, Y: 0},
	ent.Right: {X: 1, Y: 0},
}

func mod(a, m int32) int32 {
	return (a%m + m) % m
}

func modf(a, m float32) float32 {
	for a >= m {
		a -= m
	}
	for a < 0 {
		a += m
	}
	return a
}

func UpdateEntity(e *ent.Entity, l *level.Level, delta float32) {
	if e.Direction == ent.None {
		e.Direction = e.DesiredDir
	}

	if willPassCenter(e, delta) {
		x := e.X
		y := e.Y
		closest := closestCenter(e, l, delta)
		e.X = float32(closest[0])
		e.Y = float32(closest[1])
		if canMove(e, l, e.DesiredDir, delta) {
			e.Direction = e.DesiredDir
		} else {
			e.X = x
			e.Y = y
		}
	}

	dir := dirVectors[e.Direction]

	if canMove(e, l, e.Direction, delta) {
		e.X = modf(e.X+dir.X*e.Speed*delta, float32(l.Width))
		e.Y = modf(e.Y+dir.Y*e.Speed*delta, float32(l.Height))

	} else {
		closest := closestCenter(e, l, delta)
		e.X = float32(closest[0])
		e.Y = float32(closest[1])
		e.Direction = ent.None
	}

}

func willPassCenter(e *ent.Entity, delta float32) bool {
	d := dirVectors[e.Direction]

	return containsCenter(e.X, d.X*e.Speed*delta) || containsCenter(e.Y, d.Y*e.Speed*delta)
}

func closestCenter(e *ent.Entity, l *level.Level, delta float32) [2]int32 {
	d := dirVectors[e.Direction]

	nextX := e.X + d.X*e.Speed*delta
	nextY := e.Y + d.Y*e.Speed*delta

	return [2]int32{
		int32(modf(nextX+0.5, float32(l.Width))),
		int32(modf(nextY+0.5, float32(l.Height))),
	}
}

func containsCenter(f, D float32) bool {
	begin := float32(int32(f))
	end := float32(int32(f + D))
	if D < 0 {
		begin, end = end, begin
	}
	return end > begin
}

func canMove(e *ent.Entity, l *level.Level, dir ent.Direction, delta float32) bool {
	if dir == ent.None {
		return true
	}

	offset := dirVectors[dir]

	nextX := int32(modf(e.X+offset.X*e.Speed*delta, float32(l.Width)))
	nextY := int32(modf(e.Y+offset.Y*e.Speed*delta, float32(l.Height)))

	nextXS := int32(modf(e.X+0.99+offset.X*e.Speed*delta, float32(l.Width)))
	nextYS := int32(modf(e.Y+0.99+offset.Y*e.Speed*delta, float32(l.Height)))

	wall := l.Grid[nextY][nextX] != level.Wall && l.Grid[nextYS][nextXS] != level.Wall
	if e.Door {
		return wall
	}
	return wall && l.Grid[nextY][nextX] != level.Door && l.Grid[nextYS][nextXS] != level.Door

}
