package util

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewImageFromFile(path, fileName string) *ebiten.Image {
	filePath := path + fileName

	data, err := os.ReadFile(filePath)
	if err != nil {
		panic("error: reading file " + fileName + " failed: " + err.Error())
	}

	img, _, err := image.Decode(bytes.NewBuffer(data))
	if err != nil {
		panic("error: decoding image " + fileName + " failed: " + err.Error())
	}

	decodedImage := ebiten.NewImageFromImage(img)

	return decodedImage
}

func CutOffset(img *ebiten.Image, x, y int) *ebiten.Image {
	size := img.Bounds().Size()
	return img.SubImage(image.Rect(x, y, size.X, size.Y)).(*ebiten.Image)
}

type SpriteMapMeta struct {
	AnchorPoints []int
	SplitPoints  []int
}

func PrepareSpriteMap(img *ebiten.Image) (*ebiten.Image, SpriteMapMeta) {
	size := img.Bounds().Size()

	for i := 0; i < size.X; i++ {
		pixel := img.At(0, i)

		fmt.Println(pixel.RGBA())
		if r, g, b, a := pixel.RGBA(); r == 255 && g == 0 && b == 0 && a == 255 {
			fmt.Println("is red!")
		}
	}

	newImg := CutOffset(img, 1, 1)

	return newImg, SpriteMapMeta{
		AnchorPoints: []int{},
		SplitPoints:  []int{},
	}
}
