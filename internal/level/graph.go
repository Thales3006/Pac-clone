package level

import (
	"os"
	"strconv"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

type Cell struct {
	X    uint8
	Y    uint8
	Side Side
}

func posHash(c Cell) [2]uint8 {
	return [2]uint8{c.X, c.Y}
}

func (l *Level) generateGraph() {
	l.Graph = graph.New(posHash)
	for i := range l.Grid {
		for j, cell := range l.Grid[i] {
			if cell != Wall {
				l.Graph.AddVertex(Cell{uint8(i), uint8(j), cell},
					graph.VertexAttribute("pos", strconv.Itoa(j)+","+strconv.Itoa(int(l.Height-1)-i)+"!"))
			}
		}
	}

	for i := range l.Grid {
		for j, cell := range l.Grid[i] {
			if cell == Wall {
				continue
			}
			if _, err := l.Graph.Vertex([2]uint8{uint8(i), uint8(j + 1)}); err == nil {
				l.Graph.AddEdge([2]uint8{uint8(i), uint8(j)}, [2]uint8{uint8(i), uint8(j + 1)})
			}
			if _, err := l.Graph.Vertex([2]uint8{uint8(i), uint8(j - 1)}); err == nil {
				l.Graph.AddEdge([2]uint8{uint8(i), uint8(j)}, [2]uint8{uint8(i), uint8(j - 1)})
			}
			if _, err := l.Graph.Vertex([2]uint8{uint8(i + 1), uint8(j)}); err == nil {
				l.Graph.AddEdge([2]uint8{uint8(i), uint8(j)}, [2]uint8{uint8(i + 1), uint8(j)})
			}
			if _, err := l.Graph.Vertex([2]uint8{uint8(i - 1), uint8(j)}); err == nil {
				l.Graph.AddEdge([2]uint8{uint8(i), uint8(j)}, [2]uint8{uint8(i - 1), uint8(j)})
			}
		}
	}

	file, _ := os.Create("./graph.gv")
	_ = draw.DOT(l.Graph, file)
}
