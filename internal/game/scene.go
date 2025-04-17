package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Scene uint8

const (
	MainMenu Scene = iota
	Pause
	Level
)

func (g *Game) RenderScene() {
	switch g.currentScene {
	case MainMenu:
		g.RenderMainMenu()
	case Level:
		g.level.Render()

	}
}

func (g *Game) RenderMainMenu() {
	rl.DrawRectangle(0, 0, 300, 300, rl.White)
}
