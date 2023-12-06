package gameData

import (
	"go-gusanos/util"
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
		sprites[file.Name()] = util.NewImageFromFile(directory, file.Name())
	}

	// add all sprites to repository
	return sprites
}
