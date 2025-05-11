package level

import (
	"github.com/dominikbraun/graph"
)

type Level struct {
	Grid        [][]Cell                        `json:"Grid"`
	Width       int32                           `json:"-"`
	Height      int32                           `json:"-"`
	Current     string                          `json:"-"`
	Graph       graph.Graph[[2]int32, [2]int32] `json:"-"`
	Points      int32                           `json:"-"`
	SpawnPlayer [2]int32                        `json:"SpawnPlayer"`
	SpawnGhost  [2]int32                        `json:"SpawnGhost"`
	Required    int32                           `json:"Required"`
}

type Cell int32

const (
	Empty Cell = iota
	Wall
	Door
	Point
	Power
)

func NewLevel() *Level {
	return &Level{
		Grid:    nil,
		Width:   0,
		Height:  0,
		Current: "",
	}
}
