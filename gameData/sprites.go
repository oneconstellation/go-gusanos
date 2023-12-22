package gameData

import (
	"fmt"
	"go-gusanos/util"
	"image"
	"log"
	"os"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Image         *ebiten.Image
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
		newImg := util.CutOffset(ebiten.NewImageFromImage(img), 1, 1)

		sprites[file.Name()] = Sprite{
			Image:         newImg,
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
	x2, isX2Ok := s.SplitPointsX[row]
	y2, isY2Ok := s.SplitPointsY[col]

	var rowStart, colStart int

	if row-1 < 0 {
		rowStart = 0
	} else {
		rowStart = row - 1
	}

	if col-1 < 0 {
		colStart = 0
	} else {
		colStart = col - 1
	}

	x1, isX1Ok := s.SplitPointsX[rowStart]
	y1, isY1Ok := s.SplitPointsY[colStart]

	if !isX2Ok || !isY2Ok {
		return s.Image
	}

	if isX1Ok && isY1Ok {
		fmt.Println(x1, y1, x2, y2)
		return s.Image.SubImage(image.Rect(x1, y1, x2, y2)).(*ebiten.Image)
	}

	return s.Image.SubImage(image.Rect(0, 0, x2, y2)).(*ebiten.Image)
}
