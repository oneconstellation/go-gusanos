package gameData

import (
	"fmt"
	"go-gusanos/util"
	"image"
	"image/color/palette"
	_ "image/png"
	"log"
	"os"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Image         *ebiten.Image
	RawImage      image.Image
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
			r, g, b, _ := pixel.RGBA()

			// red pixel for anchor point, black for split point
			isAnchor := r == 65535 && g == 0 && b == 0
			isSplit := r == 0 && g == 0 && b == 0

			if isAnchor {
				anchorPointsX[anchorCount] = i
				anchorCount++
			}

			if isSplit {
				splitPointsX[splitCount] = i
				splitCount++
			}
		}

		anchorCount = 0
		splitCount = 0

		for i := 0; i < size.Y; i++ {
			pixel := img.At(0, i)
			r, g, b, _ := pixel.RGBA()

			// red pixel for anchor point, black for split point
			isAnchor := r == 65535 && g == 0 && b == 0
			isSplit := r == 0 && g == 0 && b == 0

			if isAnchor {
				anchorPointsY[anchorCount] = i
				anchorCount++
			}

			if isSplit {
				splitPointsY[splitCount] = i
				splitCount++
			}
		}

		fmt.Println("[" + file.Name() + "] sprites count in map: " + strconv.Itoa(len(anchorPointsX)*len(anchorPointsY)))
		fmt.Println(splitPointsX)
		fmt.Println(splitPointsY)
		fmt.Println(anchorPointsX)
		fmt.Println(anchorPointsY)

		newImg := ebiten.NewImageFromImage(img).SubImage(image.Rect(1, 1, size.X, size.Y)).(*ebiten.Image)

		sprites[file.Name()] = Sprite{
			Image:         newImg,
			RawImage:      img,
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
	// if we want row 0, col 0 - (0, 0, splitx 0, splity 0)
	// if we want row 1, col 0 - (splitx 0, 0, splitx 1, splity 1)
	// if we want row 2, col 0 - (splitx 1, 0, splitx 2, splity 1)
	// if we want row 0, col 1 - (0, splity 0, splitx 1, splity 1)
	// if we want row 1, col 1 - (splitx 0, splity 0, splitx 1, splity 1)
	// if we want row 2, col 2 - (splitx 1, splity 1, splitx 2, splity 2)
	// if we want row 2, col 3 - (splitx 1, splity 2, splitx 2, splity 3)

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

	paletted := image.NewPaletted(s.RawImage.Bounds(), palette.Plan9) // TODO, black rect instead
	cut := paletted.SubImage(image.Rect(x0, y0, x1, y1))

	return ebiten.NewImageFromImage(cut)
}
