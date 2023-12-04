package gameData

import (
	"go-gusanos/util"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type MapConfig struct {
	Spawnpoints [][]int64
}

type Map struct {
	Level    *ebiten.Image
	Material *ebiten.Image
	Config   MapConfig
}

type Maps map[string]Map

func LoadMaps(modName string) Maps {
	log.Println("loading maps...")

	maps := map[string]Map{}
	directory := util.GetModDataPath(modName) + "/maps/"

	// load files from mods/x/maps directory
	files, err := os.ReadDir(directory)
	if err != nil {
		panic("error: reading mod directory failed: " + err.Error())
	}

	for _, file := range files {
		log.Println(file.Name(), file.IsDir())
	}

	return maps
}
