package main

import (
	"math/rand"

	"github.com/megaspace/simulation/api"
	"github.com/megaspace/simulation/engine"
)

func main() {

	rand.Seed(1)

	go api.StartServer(":50052")

	game := engine.NewGame()
	game.Start()
}
