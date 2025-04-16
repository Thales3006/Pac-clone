package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) Render() {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	g.level.Render()

}

func (l *Level) Render() {
	for i := range l.grid {
		for j := range l.grid[i] {

			cell := l.grid[i][j]

			switch cell {
			case None:
				rl.DrawRectangle(int32(j*32), 100+int32(i*32), 32, 32, rl.Black)

			default:
				if cell&Left == Left {
					rl.DrawRectangle(int32(j*32), 100+int32(i*32), 4, 32, rl.Green)
				}
				if cell&Right == Right {
					rl.DrawRectangle(int32(j*32)+28, 100+int32(i*32), 4, 32, rl.Orange)
				}
				if cell&Up == Up {
					rl.DrawRectangle(int32(j*32), 100+int32(i*32), 32, 4, rl.Purple)
				}
				if cell&Down == Down {
					rl.DrawRectangle(int32(j*32), 100+int32(i*32)+28, 32, 4, rl.Gray)
				}
			}
		}
	}
}
