package data

import (
	"bytes"
	"go-gusanos/util"
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameDataRepository struct {
	Sprites []ebiten.Image
}

func (g *GameDataRepository) LoadSprites(modName string) {
	log.Println("loading sprites...")

	sprites := []ebiten.Image{}
	directory := util.GetModDataPath(modName) + "/sprites/"

	// load files from mods/x/sprites directory
	files, err := os.ReadDir(directory)
	if err != nil {
		panic("error: reading mod directory failed: " + err.Error())
	}

	// convert them to []ebiten.Image
	for _, file := range files {
		data, err := os.ReadFile(directory + file.Name())
		if err != nil {
			panic("error: reading file " + file.Name() + " failed: " + err.Error())
		}
		img, _, err := image.Decode(bytes.NewReader(data))
		if err != nil {
			panic("error: decoding image " + file.Name() + " failed: " + err.Error())
		}
		decodedImage := ebiten.NewImageFromImage(img)
		sprites = append(sprites, *decodedImage)
	}

	// add all sprites to repository
	g.Sprites = sprites
}
