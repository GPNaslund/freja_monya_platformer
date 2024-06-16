package util

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	SpriteSheet *ebiten.Image
	Frame0X     int
	Frame0Y     int
	FrameWidth  int
	FrameHeight int
	FrameCount  int
	FrameSpeed  int
}

func (a *Animation) GetFrame(updateTick int) *ebiten.Image {
	frameIndex := (updateTick / a.FrameSpeed) % a.FrameCount
	sizeX, sizeY := a.Frame0X+frameIndex*a.FrameWidth, a.Frame0Y
	return a.SpriteSheet.SubImage(image.Rect(sizeX, sizeY, sizeX+a.FrameWidth, sizeY+a.FrameHeight)).(*ebiten.Image)
}
