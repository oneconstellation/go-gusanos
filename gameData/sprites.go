package gameData

import (
	"bytes"
	"go-gusanos/util"
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprites map[string]*ebiten.Image

func LoadSprites(modName string) Sprites {
	log.Println("loading sprites...")

	sprites := map[string]*ebiten.Image{}
	directory := util.GetModDataPath(modName) + "/sprites/"

	// load files from mods/x/sprites directory
	files, err := os.ReadDir(directory)
	if err != nil {
		panic("error: reading mod directory failed: " + err.Error())
	}

	// convert them to []ebiten.Image
	for _, file := range files {
		filePath := directory + file.Name()

		data, err := os.ReadFile(filePath)
		if err != nil {
			panic("error: reading file " + file.Name() + " failed: " + err.Error())
		}
		log.Println(file.Name())
		img, _, err := image.Decode(bytes.NewBuffer(data))
		if err != nil {
			panic("error: decoding image " + file.Name() + " failed: " + err.Error())
		}
		decodedImage := ebiten.NewImageFromImage(img)
		sprites[file.Name()] = decodedImage
	}

	// add all sprites to repository
	return sprites
}
