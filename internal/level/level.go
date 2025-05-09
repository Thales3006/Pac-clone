package level

import (
	"github.com/dominikbraun/graph"
)

type Level struct {
	Grid    [][]Cell                        `json:"Grid"`
	Width   int32                           `json:"-"`
	Height  int32                           `json:"-"`
	Current string                          `json:"-"`
	Graph   graph.Graph[[2]int32, [2]int32] `json:"-"`
}

type Cell int32

const (
	Empty Cell = iota
	Wall
	Door
	Point
)

func NewLevel() *Level {
	return &Level{
		Grid:    nil,
		Width:   0,
		Height:  0,
		Current: "",
	}
}
