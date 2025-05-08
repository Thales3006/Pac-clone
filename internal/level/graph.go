package level

import (
	"os"
	"strconv"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

func posHash(c [2]int32) [2]int32 {
	return c
}

func (l *Level) generateGraph() {
	l.Graph = graph.New(posHash)
	for i := range l.Grid {
		for j, cell := range l.Grid[i] {
			if cell != Wall {
				l.Graph.AddVertex([2]int32{int32(i), int32(j)},
					graph.VertexAttribute("pos", strconv.Itoa(j)+","+strconv.Itoa(int(l.Height-1)-i)+"!"))
			}
		}
	}

	for i := range l.Grid {
		for j, cell := range l.Grid[i] {
			if cell == Wall {
				continue
			}
			if _, err := l.Graph.Vertex([2]int32{int32(i), int32(j + 1)}); err == nil {
				l.Graph.AddEdge([2]int32{int32(i), int32(j)}, [2]int32{int32(i), int32(j + 1)})
			}
			if _, err := l.Graph.Vertex([2]int32{int32(i), int32(j - 1)}); err == nil {
				l.Graph.AddEdge([2]int32{int32(i), int32(j)}, [2]int32{int32(i), int32(j - 1)})
			}
			if _, err := l.Graph.Vertex([2]int32{int32(i + 1), int32(j)}); err == nil {
				l.Graph.AddEdge([2]int32{int32(i), int32(j)}, [2]int32{int32(i + 1), int32(j)})
			}
			if _, err := l.Graph.Vertex([2]int32{int32(i - 1), int32(j)}); err == nil {
				l.Graph.AddEdge([2]int32{int32(i), int32(j)}, [2]int32{int32(i - 1), int32(j)})
			}
		}
	}

	file, _ := os.Create("./graph.gv")
	_ = draw.DOT(l.Graph, file)
}
