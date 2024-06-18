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

func (c *CollisionBox) topLeft() Vector {
	return Vector{
		X: c.Position.X - c.Width/2,
		Y: c.Position.Y - c.Height/2,
	}
}

func (c *CollisionBox) IsColliding(other CollisionBox) bool {
	thisTopLeft := c.topLeft()
	otherTopLeft := other.topLeft()

	return thisTopLeft.X < otherTopLeft.X+other.Width &&
		thisTopLeft.X+c.Width > otherTopLeft.X &&
		thisTopLeft.Y < otherTopLeft.Y+other.Height &&
		thisTopLeft.Y+c.Height > otherTopLeft.Y
}

func (c *CollisionBox) Debug(screen *ebiten.Image) {
	if c.rectangle == nil {
		c.rectangle = ebiten.NewImage(int(c.Width), int(c.Height))
		c.rectangle.Fill(color.RGBA{255, 0, 0, 128})
	}
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(-float64(c.Width)/2, -float64(c.Height)/2)
	options.GeoM.Translate(c.Position.X, c.Position.Y)
	screen.DrawImage(c.rectangle, options)
}
