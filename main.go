package main

import (
	"fmt"
	"go-gusanos/game"
	"go-gusanos/gameData"
	"go-gusanos/player"
	"go-gusanos/weapon"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	fmt.Println("gusanos!")
	ebiten.SetWindowSize(1280, 960)
	ebiten.SetWindowTitle("go-gusanos")

	gameData := gameData.New("default")

	game := &game.Game{
		Players: []player.Worm{player.New(weapon.WeaponsList{})},
		Data:    &gameData,
	}

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
