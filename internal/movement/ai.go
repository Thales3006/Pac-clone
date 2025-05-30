package movement

import (
	"math"
	"math/rand"
	ent "pac-clone/internal/entities"
	"pac-clone/internal/level"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	dirs = []ent.Direction{ent.None, ent.Up, ent.Down, ent.Left, ent.Right}
)

func HandleAI(player *ent.Player, ghosts []*ent.Ghost, l *level.Level) {
	for _, ghost := range ghosts {
		dir := ent.None
		g := rl.Vector2{X: ghost.X, Y: ghost.Y}
		p := rl.Vector2{X: player.X, Y: player.Y}

		ghost.Speed = 3
		switch {
		case ghost.State == ent.Dead:
			dir = direction(g, rl.Vector2{X: float32(l.SpawnGhost[1]), Y: float32(l.SpawnGhost[0])}, l)
			if dir == ent.None {
				ghost.State = ent.Chase
				ghost.Door = false
				ghost.Wait.Reset()
			}
		case ghost.State == ent.Scared:
			if ghost.Personality != ent.Clyde {
				dir = direction(g, l.FindFarthest(p), l)
			} else {
				dir = directionGreedy(g, l.FindFarthest(p), l)
			}
			ghost.Speed = 2
			if player.Powerfull.Done() {
				ghost.State = ent.Chase
			}
		case !ghost.Wait.Done():
			dir = ent.None
		default:
			if !ghost.Door {
				ghost.Door = true
			}

			switch ghost.Personality {
			case ent.Blinky:
				dir = direction(g, p, l)
			case ent.Pinky:
				dir = direction(g, PinkyPredict(player, l), l)
			case ent.Inky:
				dir = direction(g, InkyMean(p, ghosts), l)
			case ent.Clyde:
				if distance(g, p) < 4 {
					ghost.State = ent.CScared
				}
				if ghost.State == ent.CScared {
					dir = directionGreedy(g, rl.Vector2{X: float32(l.SpawnGhost[1]), Y: float32(l.SpawnGhost[0])}, l)
				} else {
					dir = directionGreedy(g, p, l)
				}
				if distance(g, rl.Vector2{X: float32(l.SpawnGhost[1]), Y: float32(l.SpawnGhost[0])}) < 3 {
					ghost.State = ent.Chase
				}
			}
		}

		if dir == ent.None {
			dir = ent.Direction(dirs[rand.Intn(len(dirs))])
		}

		ghost.DesiredDir = dir
	}

}

func direction(a rl.Vector2, b rl.Vector2, l *level.Level) ent.Direction {
	path, _ := l.AStar([2]int32{int32(a.Y + 0.5), int32(a.X + 0.5)}, [2]int32{int32(b.Y + 0.5), int32(b.X + 0.5)})

	next := [2]int32{int32(a.Y + 0.5), int32(a.X + 0.5)}
	if len(path) > 1 {
		next = path[1]
	}

	dir := rl.Vector2{X: float32(next[1]) - float32(int32(a.X+0.5)), Y: float32(next[0]) - float32(int32(a.Y+0.5))}

	switch dir {
	case rl.Vector2{X: 1, Y: 0}:
		return ent.Right
	case rl.Vector2{X: -1, Y: 0}:
		return ent.Left
	case rl.Vector2{X: 0, Y: 1}:
		return ent.Down
	case rl.Vector2{X: 0, Y: -1}:
		return ent.Up
	default:
		return ent.None
	}
}

func directionGreedy(a rl.Vector2, b rl.Vector2, l *level.Level) ent.Direction {
	path, _ := l.GreedyBestFirstSearch([2]int32{int32(a.Y + 0.5), int32(a.X + 0.5)}, [2]int32{int32(b.Y + 0.5), int32(b.X + 0.5)})

	next := [2]int32{int32(a.Y + 0.5), int32(a.X + 0.5)}
	if len(path) > 1 {
		next = path[1]
	}

	dir := rl.Vector2{X: float32(next[1]) - float32(int32(a.X+0.5)), Y: float32(next[0]) - float32(int32(a.Y+0.5))}

	switch dir {
	case rl.Vector2{X: 1, Y: 0}:
		return ent.Right
	case rl.Vector2{X: -1, Y: 0}:
		return ent.Left
	case rl.Vector2{X: 0, Y: 1}:
		return ent.Down
	case rl.Vector2{X: 0, Y: -1}:
		return ent.Up
	default:
		return ent.None
	}
}

func distance(a, b rl.Vector2) float32 {
	vec := rl.Vector2Subtract(b, a)
	return float32(math.Sqrt(float64(vec.X*vec.X + vec.Y*vec.Y)))
}

func InkyMean(n rl.Vector2, ghosts []*ent.Ghost) rl.Vector2 {
	mean := rl.Vector2{X: 0, Y: 0}
	counter := 0

	for _, ghost := range ghosts {
		if ghost.Personality == ent.Pinky || ghost.Personality == ent.Blinky {
			mean.X += ghost.X
			mean.Y += ghost.Y
			counter++
		}
	}

	if counter == 0 {
		return n
	}
	mean.X /= float32(counter)
	mean.Y /= float32(counter)
	return mean
}

func PinkyPredict(p *ent.Player, l *level.Level) rl.Vector2 {
	coord := [2]int32{int32(p.Y), int32(p.X)}
	next := [2]int32{int32(dirVectors[p.Direction].Y), int32(dirVectors[p.Direction].X)}

	for i := 0; i < 5; i++ {
		if l.Grid[Mod(coord[0]+next[0], l.Height)][Mod(coord[1]+next[1], l.Width)] == level.Wall {
			break
		}
		next[0] += int32(dirVectors[p.Direction].Y)
		next[1] += int32(dirVectors[p.Direction].X)
	}

	predict := rl.Vector2{X: float32(coord[1] + next[1]), Y: float32(coord[0] + next[0])}

	return predict
}
