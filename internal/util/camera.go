package util

type Camera struct {
	X, Y          float64
	MaxX, MaxY    float64
	Width, Height float64
}

func (c *Camera) CenterOnEntity(entityX, entityY float64) {
	newX := entityX - float64(c.Width)/2
	newY := entityY - float64(c.Height)/2

	// Clamp the X cooridante to valid values
	if newX < 0 {
		newX = 0
	}

	// Clamp the Y coordinate to valid values
	if newY < 0 {
		newY = 0
	}

	if newY > c.MaxY-c.Height {
		newY = c.MaxY - c.Height
	}

	// Set the new values
	c.X = newX
	c.Y = newY
}
