package game

import (
	"go-gusanos/gameData"
	"go-gusanos/player"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	Data    *gameData.GameDataRepository
	Players []player.Worm
	count   int
}

func (g *Game) Update() error {
	// update state
	g.count++

	// if g.count > 4 {
	// 	g.count = 0
	// }

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
	ebitenutil.DebugPrintAt(screen, strconv.Itoa(g.count), 160, 120)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
