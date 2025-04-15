package main

import "pac-clone/game"

func main() {
	game := new(game.Game)
	game.Run()
	game.Close()
}