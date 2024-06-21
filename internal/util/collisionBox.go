package util

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type CollisionBox struct {
	Position  *Vector
	Width     float64
	Height    float64
	rectangle *ebiten.Image
}

func (c *CollisionBox) IsColliding(other CollisionBox) bool {
	return c.Position.X < other.Position.X+other.Width &&
		c.Position.X+c.Width > other.Position.X &&
		c.Position.Y < other.Position.Y+other.Height &&
		c.Position.Y+c.Height > other.Position.Y
}

func (c *CollisionBox) DebugEntity(screen *ebiten.Image, maxX float64) {
	if c.rectangle == nil {
		c.rectangle = ebiten.NewImage(int(c.Width), int(c.Height))
		c.rectangle.Fill(color.RGBA{255, 0, 0, 128})
	}
	options := &ebiten.DrawImageOptions{}
	if c.Position.X > maxX {
		options.GeoM.Translate(maxX, c.Position.Y)
	} else {
		options.GeoM.Translate(c.Position.X, c.Position.Y)
	}
	screen.DrawImage(c.rectangle, options)
}

func (c *CollisionBox) DebugTiles(screen *ebiten.Image, cameraX, cameraY float64) {
	if c.rectangle == nil {
		c.rectangle = ebiten.NewImage(int(c.Width), int(c.Height))
		c.rectangle.Fill(color.RGBA{255, 0, 0, 128})
	}
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(c.Position.X-cameraX, c.Position.Y-cameraY)
	screen.DrawImage(c.rectangle, options)
}
