package game

import (
	"go-gusanos/player"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Game struct {
	Players []player.Worm
}

func (g *Game) Update(screen *ebiten.Image) error {
	// update state
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// render the screen
	ebitenutil.DebugPrint(screen, "go gusanos!")
	for _, player := range g.Players {
		ebitenutil.DebugPrintAt(screen, player.Name, 0, 15)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
