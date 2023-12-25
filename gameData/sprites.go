package gameData

import (
	"fmt"
	"go-gusanos/util"
	"image"
	"image/color"
	"image/color/palette"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

type Sprite struct {
	Image         *ebiten.Image
	RawImage      *image.Paletted
	AnchorPointsX map[int]int
	AnchorPointsY map[int]int
	SplitPointsX  map[int]int
	SplitPointsY  map[int]int
}

type Sprites map[string]Sprite

func LoadSprites(modName string) Sprites {
	log.Println("loading sprites...")

	sprites := map[string]Sprite{}
	directory := util.GetModDataPath(modName) + "/sprites/"

	// load files from mods/x/sprites directory
	files, err := os.ReadDir(directory)
	if err != nil {
		panic("error: reading mod directory failed: " + err.Error())
	}

	for _, file := range files {
		img := util.NewImageFromFile(directory, file.Name())
		size := img.Bounds().Size()
		anchorPointsX := map[int]int{}
		anchorPointsY := map[int]int{}
		splitPointsX := map[int]int{}
		splitPointsY := map[int]int{}

		anchorCount := 0
		splitCount := 0

		for i := 0; i < size.X; i++ {
			pixel := img.At(i, 0)

			if isAnchorPoint(pixel) {
				anchorPointsX[anchorCount] = i
				anchorCount++
			}

			if isSplitPoint(pixel) {
				splitPointsX[splitCount] = i
				splitCount++
			}
		}

		anchorCount = 0
		splitCount = 0

		for i := 0; i < size.Y; i++ {
			pixel := img.At(0, i)

			if isAnchorPoint(pixel) {
				anchorPointsY[anchorCount] = i
				anchorCount++
			}

			if isSplitPoint(pixel) {
				splitPointsY[splitCount] = i
				splitCount++
			}
		}

		newImg := img.(SubImager).SubImage(image.Rect(1, 1, size.X, size.Y))
		paletted := toPaletted(newImg)

		sprites[file.Name()] = Sprite{
			Image:         ebiten.NewImageFromImage(paletted),
			RawImage:      paletted,
			AnchorPointsX: anchorPointsX,
			AnchorPointsY: anchorPointsY,
			SplitPointsX:  splitPointsX,
			SplitPointsY:  splitPointsY,
		}
	}

	// add all sprites to repository
	return sprites
}

func (s Sprite) GetSubSprite(row, col int) *ebiten.Image {
	var x0, x1, y0, y1 int

	if row > 0 {
		x0 = s.SplitPointsX[row-1]
		x1 = s.SplitPointsX[row]
	} else {
		x0 = 0
		x1 = s.SplitPointsX[row]
	}

	if col > 0 {
		y0 = s.SplitPointsY[col-1] + 1
		y1 = s.SplitPointsY[col] + 1
	} else {
		y0 = 0
		y1 = s.SplitPointsY[col] + 1
	}

	cut := s.RawImage.SubImage(image.Rect(x0, y0, x1, y1))
	ax, ay := s.markAnchorPoint(row, col)
	fmt.Println(ax, ay)

	return ebiten.NewImageFromImage(cut)
}

func isAnchorPoint(pixel color.Color) bool {
	r, g, b, _ := pixel.RGBA()
	return r == 65535 && g == 0 && b == 0
}

func isSplitPoint(pixel color.Color) bool {
	r, g, b, _ := pixel.RGBA()
	return r == 0 && g == 0 && b == 0
}

func isTransparentPoint(pixel color.Color) bool {
	r, g, b, _ := pixel.RGBA()
	return r == 65535 && g == 0 && b == 65535
}

func toPaletted(img image.Image) *image.Paletted {
	var palette color.Palette = palette.WebSafe
	palette = append(palette, color.Transparent)

	bounds := img.Bounds()
	newImg := image.NewPaletted(bounds, palette)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := img.At(x, y)

			if isTransparentPoint(pixel) {
				// swap magenta background for transparent
				newImg.Set(x, y, color.Transparent)
				continue
			}

			newImg.Set(x, y, newImg.Palette.Convert(pixel))
		}
	}

	return newImg
}

func (s Sprite) markAnchorPoint(row, col int) (int, int) {
	x := s.AnchorPointsX[row]
	y := s.AnchorPointsY[col]

	s.RawImage.Set(x, y, color.RGBA{R: 255, G: 0, B: 0, A: 255})

	return x, y
}
