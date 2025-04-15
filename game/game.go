package game

import "fmt"

type Game struct {
	Level  *Level
	Player *Player
	Ghosts []*Ghost
}

func New() *Game {
    return &Game{
        Level:  LoadLevel(),
        Player: NewPlayer(),
        Ghosts: []*Ghost{NewGhost()},
    }
}

func (g *Game) Run() {
	fmt.Println("Game running")
}

func (g *Game) Close() {
	fmt.Println("Closing!")
}