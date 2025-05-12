package level

import (
	"container/heap"
	"errors"
)

func (l *Level) GreedyBestFirstSearch(start, goal [2]int32) ([][2]int32, error) {
	openSet := &PriorityQueue{}
	heap.Init(openSet)
	heap.Push(openSet, &PriorityQueueItem{Cell: start, Priority: heuristic(start, goal)})

	cameFrom := make(map[[2]int32][2]int32)
	visited := make(map[[2]int32]bool)

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

		if visited[posHash(current)] {
			continue
		}
		visited[posHash(current)] = true

		for neighborHash := range adjMap[posHash(current)] {
			neighbor, _ := l.Graph.Vertex(neighborHash)

			if visited[neighborHash] {
				continue
			}

			priority := heuristic(neighbor, goal)
			heap.Push(openSet, &PriorityQueueItem{Cell: neighbor, Priority: priority})
			cameFrom[neighborHash] = current
		}
	}

	return nil, errors.New("did not found a path")
}
