package util

import "os"

func GetModDataPath(modName string) string {
	directory := "mods/" + modName
	_, err := os.Stat(directory)
	if os.IsNotExist(err) {
		panic("error: mod directory does not exist: " + err.Error())
	}

	return directory
}
