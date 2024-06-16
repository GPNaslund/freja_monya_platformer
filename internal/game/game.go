package game

import (
	"image/color"

	player "github.com/gpnaslund/freja_monya_platformer/internal/player/monya"
	"github.com/gpnaslund/freja_monya_platformer/internal/util"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	monya *player.Monya
}

func NewGame() *Game {
	monya := player.NewMonya(util.Vector{X: 320 / 2, Y: 240 / 2})
	return &Game{
		monya: monya,
	}
}

func (g *Game) Update() error {
	g.monya.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	img := ebiten.NewImage(320, 240)
	img.Fill(color.White)
	screen.DrawImage(img, nil)
	g.monya.Draw(screen, true)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
