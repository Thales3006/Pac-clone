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
			level.Cell{Y: uint8(ghost.X), X: uint8(ghost.Y), Side: level.Empty},
			level.Cell{Y: uint8(p.X), X: uint8(p.Y), Side: level.Empty},
			l,
		)
		ghost.DesiredDir = dir
		fmt.Println(dir)
	}

}

func direction(a level.Cell, b level.Cell, l *level.Level) entities.Direction {
	path, _ := level.AStar(l.Graph, a, b)

	next := level.Cell{X: uint8(a.Y), Y: uint8(a.X)}
	if len(path) > 1 {
		next = path[1]
	}

	dir := rl.Vector2{X: float32(next.Y) - float32(a.Y), Y: float32(next.X) - float32(a.X)}
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
