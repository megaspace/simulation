package main

import (
	"math/rand"

	"github.com/megaspace/simulation/api"
	"github.com/megaspace/simulation/game"
)

func main() {

	rand.Seed(1)

	go api.StartServer(":50052")

	game := game.New()
	go game.Start()
}
