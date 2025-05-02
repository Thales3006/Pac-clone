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

func UpdateEntity(e *ent.Entity, l *level.Level, delta float32) {

	dir := dirVectors[e.Direction]
	ndir := dirVectors[e.DesiredDir]

	if e.Direction == ent.None {
		e.Direction = e.DesiredDir
		return
	}

	if willPassCenter(e, delta) && e.Direction != e.DesiredDir {

		if e.Direction == ent.Right || e.Direction == ent.Left {
			temp := deltaInt(e.X, dir.X)
			e.X = float32(int32(e.X)) + (dir.X+1)/2
			if canMove(e, l, e.DesiredDir, delta) {
				e.Y += ndir.Y * (e.Speed*delta - temp)
				e.Direction = e.DesiredDir
			}
		} else {
			temp := deltaInt(e.Y, dir.Y)
			e.Y = float32(int32(e.Y)) + (dir.Y+1)/2
			if canMove(e, l, e.DesiredDir, delta) {
				e.X += ndir.X * (e.Speed*delta - temp)
				e.Direction = e.DesiredDir
			}
		}
		return
	}

	if !canMove(e, l, e.Direction, delta) {
		e.Direction = ent.None
		e.X = float32(int32(e.X))
		e.Y = float32(int32(e.Y))
		return
	}

	e.X += dir.X * e.Speed * delta
	e.Y += dir.Y * e.Speed * delta

	if e.X >= float32(l.Width) {
		e.X = 0
	}
	if e.Y >= float32(l.Height) {
		e.Y = 0
	}
	if e.X < 0 {
		e.X = float32(l.Width) - 1
	}
	if e.Y < 0 {
		e.Y = float32(l.Height) - 1
	}

}

func willPassCenter(e *ent.Entity, delta float32) bool {
	d := dirVectors[e.Direction]

	return containsCenter(e.X, d.X*e.Speed*delta) || containsCenter(e.Y, d.Y*e.Speed*delta)
}

func deltaInt(f, dir float32) float32 {
	if dir == 1 {
		return 1 - (f - float32(int32(f)))
	} else {
		return (f - float32(int32(f)))
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
		return false
	}

	offset := dirVectors[dir]

	nextX := uint8(e.X + offset.X*e.Speed*delta)
	nextY := uint8(e.Y + offset.Y*e.Speed*delta)

	nextXs := uint8(e.X + 0.99 + offset.X*e.Speed*delta)
	nextYs := uint8(e.Y + 0.99 + offset.Y*e.Speed*delta)

	if nextX >= l.Width || nextY >= l.Height || nextXs >= l.Width || nextYs >= l.Height {
		return true
	}
	return l.Grid[nextY][nextX] != level.Wall && l.Grid[nextYs][nextXs] != level.Wall

}
