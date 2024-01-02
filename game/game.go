package game

import (
	"fmt"
	"go-gusanos/gameData"
	"go-gusanos/player"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	Data    *gameData.GameDataRepository
	Players []*player.Worm
	Keys    []ebiten.Key
}

func (g *Game) Update() error {
	// update state
	g.Keys = inpututil.AppendPressedKeys(g.Keys[:0])
	for _, player := range g.Players {
		player.Update(g.Keys)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// render the screen
	// render the selected (where?) map
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(g.Data.Maps["poo"].Level, op)

	// render player
	for _, player := range g.Players {
		player.Render(screen)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("x speed: %f", player.XSpeed), 10, 10)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("y speed: %f", player.YSpeed), 10, 20)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
