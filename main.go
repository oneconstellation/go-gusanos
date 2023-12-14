package main

import (
	"fmt"
	"go-gusanos/game"
	"go-gusanos/gameData"
	"go-gusanos/player"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	fmt.Println("gusanos!")
	ebiten.SetWindowSize(640, 480)
	// ebiten.SetWindowSize(1280, 960)
	ebiten.SetWindowTitle("go-gusanos")
	ebiten.SetVsyncEnabled(true)

	gameData := gameData.New("default")

	game := &game.Game{
		Players: []player.Worm{player.New()},
		Data:    &gameData,
	}

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
