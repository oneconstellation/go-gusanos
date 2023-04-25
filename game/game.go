package game

import "github.com/hajimehoshi/ebiten"

type Game struct{}

func (g *Game) Update(screen *ebiten.Image) error {
	// update state
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// render the screen
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
