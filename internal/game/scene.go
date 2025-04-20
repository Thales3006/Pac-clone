package game

type Scene uint8

const (
	MainMenu Scene = iota
	SelectionMenu
	Pause
	Settings
	Level
)

func (g *Game) HandleScene() {
	switch g.currentScene {
	case MainMenu:
		g.HandleMainMenu()
	case SelectionMenu:
		g.HandleSelectionMenu()
	case Pause:
		g.HandlePause()
	case Level:
		g.HandleLevel()
	case Settings:
		g.HandleSettings()
	}
}
