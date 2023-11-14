package main

import (
	"fmt"
	"go-gusanos/game"
	"go-gusanos/player"
	"go-gusanos/weapon"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	fmt.Println("gusanos!")
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("go-gusanos")

	game := &game.Game{
		Players: []player.Worm{player.New(weapon.WeaponsList{})},
	}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
