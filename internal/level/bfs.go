package level

import (
	"github.com/dominikbraun/graph"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var adjMap map[[2]int32]map[[2]int32]graph.Edge[[2]int32]
var currentLevel string = ""

func (l *Level) FindFarthest(pos rl.Vector2) rl.Vector2 {
	start := [2]int32{int32(pos.Y), int32(pos.X)}

	visited := map[[2]int32]bool{}
	queue := [][2]int32{start}
	visited[start] = true
	var farthest [2]int32

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		farthest = current

		if currentLevel != l.Current {
			adjMap, _ = l.Graph.AdjacencyMap()
			currentLevel = l.Current
		}

		for _, neighbor := range adjMap[current] {
			if !visited[neighbor.Target] {
				visited[neighbor.Target] = true
				queue = append(queue, neighbor.Target)
			}
		}
	}

	return rl.Vector2{X: float32(farthest[1]), Y: float32(farthest[0])}
}
