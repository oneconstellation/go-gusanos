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
	Config   []byte
}

type Maps map[string]Map

func LoadMaps(modName string) Maps {
	log.Println("loading maps...")

	maps := map[string]Map{}
	directory := util.GetModDataPath(modName) + "/maps/"

	// load files from mods/x/maps directory
	mapDirectories, err := os.ReadDir(directory)
	if err != nil {
		panic("error: reading map directory failed: " + err.Error())
	}

	for _, mapDir := range mapDirectories {
		if mapDir.IsDir() {
			mapFiles, err := os.ReadDir(directory + mapDir.Name())
			if err != nil {
				panic("error: reading map directory failed: " + err.Error())
			}

			var level, material *ebiten.Image
			var config []byte

			for _, file := range mapFiles {
				switch file.Name() {
				case "level.png":
					level = util.NewImageFromFile(directory+mapDir.Name()+"/", file.Name())
				case "material.png":
					material = util.NewImageFromFile(directory+mapDir.Name()+"/", file.Name())
				case "config.cfg":
					config, err = os.ReadFile(directory + mapDir.Name() + "/" + file.Name())
					if err != nil {
						panic("error: reading map config failed: " + err.Error())
					}
				}
			}

			if level != nil && material != nil && config != nil {
				// initialize map for complete maps
				maps[mapDir.Name()] = Map{
					Level:    level,
					Material: material,
					Config:   config,
				}
			}
		}
	}

	log.Println(maps)

	return maps
}
