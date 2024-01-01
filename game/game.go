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
	count   int
}

func (g *Game) Update() error {
	// update state
	g.count++

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

	frame := (g.count / 10) % 4

	// render player
	for _, player := range g.Players {
		player.Render(screen, frame)
		ebitenutil.DebugPrint(screen, fmt.Sprintf("%f", player.Aim))
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
