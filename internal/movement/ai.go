package movement

import (
	"fmt"
	"pac-clone/internal/entities"
	"pac-clone/internal/level"

	"github.com/dominikbraun/graph"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func HandleAI(p *entities.Player, ghosts []*entities.Ghost, l *level.Level) {
	gGrid := [2]uint8{uint8(ghosts[0].X), uint8(ghosts[0].Y)}
	pGrid := [2]uint8{uint8(p.X), uint8(p.Y)}
	path, err := graph.ShortestPath(l.Graph, gGrid, pGrid)
	fmt.Println("NOVO FRAME:")
	fmt.Println(gGrid)
	fmt.Println(pGrid)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if len(path) <= 1 {
		fmt.Println("Drento")
		return
	}
	next := path[1]

	dir := rl.Vector2{X: float32(next[0]) - float32(gGrid[0]), Y: float32(next[1]) - float32(gGrid[1])}
	fmt.Println(dir)
	fmt.Println("----------")
}
