package main

import (
	"fmt"
	"go-gusanos/data"
	"go-gusanos/game"
	"go-gusanos/player"
	"go-gusanos/weapon"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	fmt.Println("gusanos!")
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("go-gusanos")

	gameData := data.GameDataRepository{}

	gameData.LoadSprites("default")

	game := &game.Game{
		Players: []player.Worm{player.New(weapon.WeaponsList{})},
	}

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
