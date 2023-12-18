package game

import (
	"go-gusanos/gameData"
	"go-gusanos/player"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Data    *gameData.GameDataRepository
	Players []player.Worm
	count   int
}

func (g *Game) Update() error {
	// update state
	g.count++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// render the screen
	// render the selected (where?) map
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(g.Data.Maps["poo"].Level, op)

	frame := (g.count / 10) % 4

	// render player
	for _, player := range g.Players {
		player.Render(screen, frame)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
