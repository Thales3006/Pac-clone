package main

import "pac-clone/internal/game"

func main() {
	myGame := game.NewGame()
	myGame.Run()
	myGame.Close()
}
