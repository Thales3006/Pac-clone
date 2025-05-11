package game

import (
	"pac-clone/internal/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) HandleSettings() {
	rl.ClearBackground(backgroundColor)

	ui.NewComponent([]ui.Element{
		&ui.Label{
			Text: "Contexto",
			Size: 30,
		},
		&ui.Label{
			Text: "Pac-man está tentando fazer a sua cidade\n" +
				"um lugar melhor ajudando a limpar as ruas, porém\n" +
				"os fantamas querem impedi-lo. \n" +
				"Ajude o Pac-man reciclar todo o lixo da cidade",
			Size: 20,
		},
		&ui.Button{
			Text: "Voltar",
			OnClick: func() {
				g.currentScene = MainMenu
			},
		},
	}).
		Use(g.center(600, 300))

	if rl.IsKeyPressed(rl.KeyEscape) {
		if g.Level.Current == "" {
			g.currentScene = MainMenu
		} else {
			g.currentScene = Pause
		}

	}
}
