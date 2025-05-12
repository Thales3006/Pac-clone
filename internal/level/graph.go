package level

import (
	"strconv"

	"github.com/dominikbraun/graph"
)

func posHash(c [2]int32) [2]int32 {
	return c
}

func mod(a, m int32) int32 {
	return (a%m + m) % m
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
			l.Graph.AddEdge([2]int32{int32(i), int32(j)}, [2]int32{int32(i), mod(int32(j+1), l.Width)})
			l.Graph.AddEdge([2]int32{int32(i), int32(j)}, [2]int32{int32(i), mod(int32(j-1), l.Width)})
			l.Graph.AddEdge([2]int32{int32(i), int32(j)}, [2]int32{mod(int32(i+1), l.Height), int32(j)})
			l.Graph.AddEdge([2]int32{int32(i), int32(j)}, [2]int32{mod(int32(i-1), l.Height), int32(j)})
		}
	}

	//file, _ := os.Create("./graph.gv")
	//_ = draw.DOT(l.Graph, file)
}
