package movement

import (
	"fmt"
	"pac-clone/internal/entities"
	"pac-clone/internal/level"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func HandleAI(p *entities.Player, ghosts []*entities.Ghost, l *level.Level) {
	for _, ghost := range ghosts {
		dir := direction(
			rl.Vector2{X: ghost.X, Y: ghost.Y},
			rl.Vector2{X: p.X, Y: p.Y},
			l)
		ghost.DesiredDir = dir
		fmt.Println(dir)
	}

}

func direction(a rl.Vector2, b rl.Vector2, l *level.Level) entities.Direction {
	path, _ := level.AStar(l.Graph, [2]int32{int32(a.Y), int32(a.X)}, [2]int32{int32(b.Y), int32(b.X)})

	next := [2]int32{int32(a.Y), int32(a.X)}
	if len(path) > 1 {
		next = path[1]
	}

	dir := rl.Vector2{X: float32(next[1]) - float32(int32(a.X)), Y: float32(next[0]) - float32(int32(a.Y))}
	fmt.Println(dir)

	switch dir {
	case rl.Vector2{X: 1, Y: 0}:
		return entities.Right
	case rl.Vector2{X: -1, Y: 0}:
		return entities.Left
	case rl.Vector2{X: 0, Y: 1}:
		return entities.Down
	case rl.Vector2{X: 0, Y: -1}:
		return entities.Up
	default:
		return entities.None
	}
}
