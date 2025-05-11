package game

import (
	"pac-clone/internal/entities"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	tileTableStreet = map[Connection]rl.Vector2{
		Horizontal: {X: 2, Y: 1},
		Vertical:   {X: 3, Y: 1},
		None:       {X: 4, Y: 1},
	}
	tileTable = map[Pick][]rl.Vector2{
		Wall:   {{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 3, Y: 0}, {X: 4, Y: 0}, {X: 5, Y: 0}, {X: 8, Y: 1}, {X: 6, Y: 0}, {X: 7, Y: 0}, {X: 8, Y: 0}, {X: 9, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}, {X: 8, Y: 1}},
		Point:  {{X: 0, Y: 2}, {X: 1, Y: 2}, {X: 2, Y: 2}, {X: 9, Y: 1}},
		Power:  {{X: 3, Y: 2}},
		Player: {{X: 3, Y: 2}, {X: 4, Y: 2}},
	}
	tileScared = map[entities.Direction][]rl.Vector2{
		entities.Down:  {{X: 8, Y: 3}, {X: 9, Y: 3}},
		entities.Left:  {{X: 8, Y: 4}, {X: 9, Y: 4}},
		entities.Up:    {{X: 8, Y: 6}, {X: 9, Y: 5}},
		entities.Right: {{X: 8, Y: 6}, {X: 9, Y: 6}},
	}
	tileDead = map[entities.Direction][]rl.Vector2{
		entities.Down:  {{X: 0, Y: 7}, {X: 1, Y: 7}},
		entities.Left:  {{X: 0, Y: 8}, {X: 1, Y: 8}},
		entities.Up:    {{X: 0, Y: 9}, {X: 1, Y: 9}},
		entities.Right: {{X: 2, Y: 7}, {X: 3, Y: 7}},
	}
	tileGhost = map[entities.Personality]map[entities.State]map[entities.Direction][]rl.Vector2{
		entities.Blinky: {
			entities.Chase: {
				entities.Down:  {{X: 0, Y: 3}, {X: 1, Y: 3}},
				entities.Left:  {{X: 0, Y: 4}, {X: 1, Y: 4}},
				entities.Up:    {{X: 0, Y: 5}, {X: 1, Y: 5}},
				entities.Right: {{X: 0, Y: 6}, {X: 1, Y: 6}},
				entities.None:  {{X: 0, Y: 3}, {X: 1, Y: 3}},
			},
			entities.Scared: tileScared,
			entities.Dead:   tileDead,
		},
		entities.Pinky: {
			entities.Chase: {
				entities.Down:  {{X: 2, Y: 3}, {X: 3, Y: 3}},
				entities.Left:  {{X: 2, Y: 4}, {X: 3, Y: 4}},
				entities.Up:    {{X: 2, Y: 5}, {X: 3, Y: 5}},
				entities.Right: {{X: 2, Y: 6}, {X: 3, Y: 6}},
				entities.None:  {{X: 2, Y: 3}, {X: 3, Y: 3}},
			},
			entities.Scared: tileScared,
			entities.Dead:   tileDead,
		},
		entities.Inky: {
			entities.Chase: {
				entities.Down:  {{X: 4, Y: 3}, {X: 5, Y: 3}},
				entities.Left:  {{X: 4, Y: 4}, {X: 5, Y: 4}},
				entities.Up:    {{X: 4, Y: 5}, {X: 5, Y: 5}},
				entities.Right: {{X: 4, Y: 6}, {X: 5, Y: 6}},
				entities.None:  {{X: 4, Y: 3}, {X: 5, Y: 3}},
			},
			entities.Scared: tileScared,
			entities.Dead:   tileDead,
		},
		entities.Clyde: {
			entities.Chase: {
				entities.Down:  {{X: 6, Y: 3}, {X: 7, Y: 3}},
				entities.Left:  {{X: 6, Y: 4}, {X: 7, Y: 4}},
				entities.Up:    {{X: 6, Y: 5}, {X: 7, Y: 5}},
				entities.Right: {{X: 6, Y: 6}, {X: 7, Y: 6}},
				entities.None:  {{X: 6, Y: 3}, {X: 7, Y: 3}},
			},
			entities.CScared: {
				entities.Down:  {{X: 6, Y: 3}, {X: 7, Y: 3}},
				entities.Left:  {{X: 6, Y: 4}, {X: 7, Y: 4}},
				entities.Up:    {{X: 6, Y: 5}, {X: 7, Y: 5}},
				entities.Right: {{X: 6, Y: 6}, {X: 7, Y: 6}},
				entities.None:  {{X: 6, Y: 3}, {X: 7, Y: 3}},
			},
			entities.Scared: tileScared,
			entities.Dead:   tileDead,
		},
	}
	degree = map[entities.Direction]float32{
		entities.Right: 0,
		entities.Left:  180,
		entities.Up:    270,
		entities.Down:  90,
	}
)
