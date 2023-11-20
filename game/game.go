package game

import (
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
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(g.Data.Sprites["uzumaki.png"], op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
