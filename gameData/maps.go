package gameData

import (
	"bufio"
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
			var config MapConfig

			for _, file := range mapFiles {
				switch file.Name() {
				case "level.png":
					img := util.NewImageFromFile(directory+mapDir.Name()+"/", file.Name())
					level = ebiten.NewImageFromImage(img)
				case "material.png":
					img := util.NewImageFromFile(directory+mapDir.Name()+"/", file.Name())
					material = ebiten.NewImageFromImage(img)
				case "config.cfg":
					config = ParseMapConfigFile(directory + mapDir.Name() + "/" + file.Name())
				}
			}

			if level != nil && material != nil {
				// initialize map for complete maps
				maps[mapDir.Name()] = Map{
					Level:    level,
					Material: material,
					Config:   config,
				}
			}
		}
	}

	return maps
}

func ParseMapConfigFile(filepath string) MapConfig {
	// load map config file
	configFile, err := os.Open(filepath)
	if err != nil {
		panic("error: reading map config failed: " + err.Error())
	}
	defer configFile.Close()

	scanner := bufio.NewScanner(configFile)
	scanner.Split(bufio.ScanWords)

	// TODO
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

	return MapConfig{}
}
