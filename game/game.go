package game

import (
	"fmt"
	"go-gusanos/gameData"
	"go-gusanos/player"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Data    *gameData.GameDataRepository
	Players []player.Worm
}

func (g *Game) Update() error {
	// update state
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// render the screen
	// render the selected (where?) map
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(g.Data.Maps["poo"].Level, op)

	// render players
	for _, player := range g.Players {
		screen.DrawImage(g.Data.Sprites["skin.png"], op)
		fmt.Println("added player ", player)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
