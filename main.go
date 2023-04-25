package main

import (
	"fmt"
	"go-gusanos/game"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	fmt.Println("gusanos!")
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("go-gusanos")

	game := &game.Game{}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
