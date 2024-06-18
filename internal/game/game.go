package game

import (
	"image/color"

	level1 "github.com/gpnaslund/freja_monya_platformer/internal/levels/lvl1"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

type Game struct {
	lvl1 *level1.Level
}

func NewGame() *Game {
	lvl1 := level1.NewLevel()
	return &Game{
		lvl1: lvl1,
	}
}

func (g *Game) Update() error {
	g.lvl1.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	g.lvl1.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
