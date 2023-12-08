package gameData

import (
	"bufio"
	"go-gusanos/util"
	"log"
	"os"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

type MapConfig struct {
	Spawnpoints [][]int64
}

type Map struct {
	Level    *ebiten.Image
	Material *ebiten.Image
	Config   string
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
			var configFile *os.File
			var configText string

			for _, file := range mapFiles {
				switch file.Name() {
				case "level.png":
					level = util.NewImageFromFile(directory+mapDir.Name()+"/", file.Name())
				case "material.png":
					material = util.NewImageFromFile(directory+mapDir.Name()+"/", file.Name())
				case "config.cfg":
					// use parser here
				}
			}

			if level != nil && material != nil && configFile != nil {
				// initialize map for complete maps
				maps[mapDir.Name()] = Map{
					Level:    level,
					Material: material,
					Config:   configText,
				}
			}
		}
	}

	// log.Println(maps)

	return maps
}

func ParseMapConfigFile(filepath string) MapConfig {
	// not going to recreate whole OMFGScript at the moment,
	// just simple, naive parsing the config.cfg for spawnpoints

	// load map config file
	configFile, err := os.Open(filepath)
	if err != nil {
		panic("error: reading map config failed: " + err.Error())
	}
	defer configFile.Close()

	var unparsedLines []string = []string{}
	// scan file for spawnpoints definition
	scanner := bufio.NewScanner(configFile)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {

	})

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "spawnpoints = [") {
			continue
		}
	}

	return MapConfig{}
}
