package util

import (
	"bytes"
	"image"
	_ "image/png"
	"os"
)

func NewImageFromFile(path, fileName string) image.Image {
	filePath := path + fileName

	data, err := os.ReadFile(filePath)
	if err != nil {
		panic("error: reading file " + fileName + " failed: " + err.Error())
	}

	img, _, err := image.Decode(bytes.NewBuffer(data))
	if err != nil {
		panic("error: decoding image " + fileName + " failed: " + err.Error())
	}

	return img
}
