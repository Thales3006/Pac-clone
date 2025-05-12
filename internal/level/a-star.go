package level

import (
	"container/heap"
	"errors"
)

func (l *Level) AStar(start, goal [2]int32) ([][2]int32, error) {
	openSet := &PriorityQueue{}
	heap.Init(openSet)
	heap.Push(openSet, &PriorityQueueItem{Cell: start, Priority: 0})

	cameFrom := make(map[[2]int32][2]int32)
	costSoFar := make(map[[2]int32]float64)
	costSoFar[posHash(start)] = 0

	if currentLevel != l.Current {
		adjMap, _ = l.Graph.AdjacencyMap()
		currentLevel = l.Current
	}

	for openSet.Len() > 0 {
		current := heap.Pop(openSet).(*PriorityQueueItem).Cell

		if posHash(current) == posHash(goal) {
			path := [][2]int32{current}
			for current != start {
				current = cameFrom[posHash(current)]
				path = append([][2]int32{current}, path...)
			}
			return path, nil
		}

		neighborsMap := adjMap[posHash(current)]
		for neighborHash := range neighborsMap {
			newCost := costSoFar[posHash(current)] + 1
			neighbor, _ := l.Graph.Vertex(neighborHash)

			if oldCost, ok := costSoFar[neighborHash]; !ok || newCost < oldCost {
				costSoFar[neighborHash] = newCost
				priority := newCost + heuristic(neighbor, goal)
				heap.Push(openSet, &PriorityQueueItem{Cell: neighbor, Priority: priority})
				cameFrom[neighborHash] = current
			}
		}
	}

	return nil, errors.New("did not found a path")
}
