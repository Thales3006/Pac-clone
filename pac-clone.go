package main

import "pac-clone/game"

func main() {
	myGame := game.New()
	myGame.Run()
	myGame.Close()
}
