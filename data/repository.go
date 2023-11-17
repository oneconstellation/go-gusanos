package data

import (
	"go-gusanos/util"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
)

type GameDataRepository struct {
	Sprites []ebiten.Image
}

func (g GameDataRepository) LoadSprites(modName string) {
	log.Println("loading sprites...")

	directory := util.GetModDataPath(modName) + "/sprites"

	// load files from mods/x/sprites directory
	files, err := os.ReadDir(directory)
	if err != nil {
		panic("error: reading mod directory failed: " + err.Error())
	}

	// convert them to []ebiten.Image
	for _, file := range files {
		data, err := os.ReadFile(directory + file.Name())
	}

	// add all sprites to repository
}
