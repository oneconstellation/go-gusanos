package util

import (
	"bytes"
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
	bounds := img.Bounds().Size()
	return img.SubImage(image.Rect(x, y, bounds.X, bounds.Y)).(*ebiten.Image)
}
