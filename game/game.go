package game

import (
	"go-gusanos/gameData"
	"go-gusanos/player"
	"image/color"
	"strconv"
	"strings"

	"github.com/hajimehoshi/bitmapfont/v3"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Game struct {
	Data    *gameData.GameDataRepository
	Players []player.Worm
	Keys    []ebiten.Key
	count   int
}

func (g *Game) Update() error {
	// update state
	g.count++

	g.Keys = inpututil.AppendPressedKeys(g.Keys[:0])

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
		ebitenutil.DebugPrintAt(screen, "frame:"+strconv.Itoa(frame), 10, 60)
	}
	ebitenutil.DebugPrintAt(screen, strconv.Itoa(g.count), 160, 120)

	var keyStrs []string
	var keyNames []string
	for _, k := range g.Keys {
		keyStrs = append(keyStrs, k.String())
		if name := ebiten.KeyName(k); name != "" {
			keyNames = append(keyNames, name)
		}
	}

	// Use bitmapfont.Face instead of ebitenutil.DebugPrint, since some key names might not be printed with DebugPrint.
	text.Draw(screen, strings.Join(keyStrs, ", ")+"\n"+strings.Join(keyNames, ", "), bitmapfont.Face, 4, 12, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
