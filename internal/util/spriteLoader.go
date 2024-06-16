package util

import (
	"bytes"
	"embed"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadSprite(embedded embed.FS, path string) (*ebiten.Image, error) {
	fileData, err := embedded.ReadFile(path)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(bytes.NewReader(fileData))
	if err != nil {
		return nil, err
	}

	sprite := ebiten.NewImageFromImage(img)
	return sprite, nil
}
