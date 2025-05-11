package game

import (
	"image/color"
	"pac-clone/internal/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	loaded_MainMenu bool
	mainTitle       rl.Texture2D
	mainTheme       rl.Texture2D
	backgroundColor color.RGBA = color.RGBA{167, 247, 253, 255}
)

func (g *Game) HandleMainMenu() {
	rl.ClearBackground(backgroundColor)

	if !loaded_MainMenu {
		g.loadMainMenu()
	}

	g.PlotMainTheme()

	titleRect := g.center(500, 250)
	titleRect.Y *= 0.2

	rl.DrawTexturePro(mainTitle,
		rl.Rectangle{X: 0, Y: 0, Width: float32(mainTheme.Width), Height: float32(mainTheme.Height)},
		titleRect,
		rl.Vector2{X: 0, Y: 0},
		0,
		rl.White,
	)

	ui.NewComponent([]ui.Element{
		&ui.Label{
			Text: "",
			Size: 30,
		},
		&ui.Button{
			Text: "Jogar",
			OnClick: func() {
				g.currentScene = SelectionMenu
			},
		},
		&ui.Button{
			Text: "Editor",
			OnClick: func() {
				g.currentScene = Editor
			},
		},
		&ui.Button{
			Text: "Contexto",
			OnClick: func() {
				g.currentScene = Settings
			},
		},
		&ui.Button{
			Text: "Sair",
			OnClick: func() {
				g.isRunning = false
			},
		},
	}).
		Use(g.center(300, 300))
}

func (g *Game) loadMainMenu() {

	mainTitle = rl.LoadTexture("assets/main-title.png")
	mainTheme = rl.LoadTexture("assets/main-theme.png")

	loaded_MainMenu = true
}

func (g *Game) PlotMainTheme() {
	scale := float32(g.Height) / float32(mainTheme.Height)
	rl.DrawTexturePro(mainTheme,
		rl.Rectangle{X: 0, Y: 0, Width: float32(mainTheme.Width), Height: float32(mainTheme.Height)},
		rl.Rectangle{X: float32(g.Width)/2 - float32(mainTheme.Width)*scale/2, Y: 0, Width: float32(mainTheme.Width) * scale, Height: float32(mainTheme.Height) * scale},
		rl.Vector2{X: 0, Y: 0},
		0,
		rl.White,
	)
}
