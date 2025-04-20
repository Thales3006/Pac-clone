package movement

import (
	"math"
	ent "pac-clone/internal/entities"
	"pac-clone/internal/level"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var dirVectors = map[ent.Direction]rl.Vector2{
	ent.Up:    {X: 0, Y: 1},
	ent.Down:  {X: 0, Y: -1},
	ent.Left:  {X: -1, Y: 0},
	ent.Right: {X: 1, Y: 0},
}

func UpdateEntity(e *ent.Entity, l *level.Level, delta float32) {
	if e.Direction == ent.None {
		return
	}

	if !canMove(e, l, delta) {
		e.Direction = ent.None
		e.X = float32(int32(e.X))
		e.Y = float32(int32(e.Y))
		return
	}

	dir := dirVectors[e.Direction]
	e.X += dir.X * e.Speed * delta
	e.Y += dir.Y * e.Speed * delta
}

func isCenteredOnTile(e *ent.Entity) bool {
	eps := 0.1
	return math.Abs(float64(e.X-float32(int(e.X)))) < eps && math.Abs(float64(e.Y-float32(int(e.Y)))) < eps
}

func canMove(e *ent.Entity, l *level.Level, delta float32) bool {
	if e.Direction == ent.None {
		return false
	}

	offset := dirVectors[e.Direction]

	nextX := uint8(e.X + offset.X*e.Speed*delta)
	nextY := uint8(e.Y + offset.Y*e.Speed*delta)

	if nextX >= l.Width || nextY >= l.Height {
		return false
	}
	return l.Grid[nextX][nextY] == level.Empty

}
