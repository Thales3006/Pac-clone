package level

import (
	"container/heap"
	"errors"
	"math"

	"github.com/dominikbraun/graph"
)

type PriorityQueueItem struct {
	Cell     Cell
	Priority float64
	Index    int
}

type PriorityQueue []*PriorityQueueItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*PriorityQueueItem)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func heuristic(a, b Cell) float64 {
	return math.Abs(float64(a.X)-float64(b.X)) + math.Abs(float64(a.Y)-float64(b.Y))
}

func AStar(g graph.Graph[[2]uint8, Cell], start, goal Cell) ([]Cell, error) {
	openSet := &PriorityQueue{}
	heap.Init(openSet)
	heap.Push(openSet, &PriorityQueueItem{Cell: start, Priority: 0})

	cameFrom := make(map[[2]uint8]Cell)
	costSoFar := make(map[[2]uint8]float64)
	costSoFar[posHash(start)] = 0

	adjMap, err := g.AdjacencyMap()
	if err != nil {
		return nil, errors.New("unable to get adjacency map")
	}

	for openSet.Len() > 0 {
		current := heap.Pop(openSet).(*PriorityQueueItem).Cell

		if posHash(current) == posHash(goal) {
			path := []Cell{current}
			for current != start {
				current = cameFrom[posHash(current)]
				path = append([]Cell{current}, path...)
			}
			return path, nil
		}

		neighborsMap := adjMap[posHash(current)]
		for neighborHash := range neighborsMap {
			newCost := costSoFar[posHash(current)] + 1
			neighbor, _ := g.Vertex(neighborHash)

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
