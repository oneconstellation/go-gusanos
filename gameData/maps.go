package gameData

import (
	"go-gusanos/util"
	"io/fs"
	"log"
	"path/filepath"

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
	filepath.Walk(directory, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			panic("error: reading mod directory failed: " + err.Error())
		}

		// process map files here

		return nil
	})

	return maps
}
