package game

import rl "github.com/gen2brain/raylib-go/raylib"

func (g *Game) Render() {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	g.Level.Render()

}

func (l *Level) Render() {
	for i := range l.Grid {
		for j := range l.Grid[i] {
			//this must be changed in the future
			color := rl.White
			if l.Grid[i][j] != 0 {
				color = rl.Blue
			}
			rl.DrawRectangle(int32(j*16), 100+int32(i*16), 8, 8, color)
		}
	}
}
