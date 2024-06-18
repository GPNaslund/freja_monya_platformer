package level1

import (
	"bytes"
	"image"
	"log"

	"github.com/gpnaslund/freja_monya_platformer/internal/levels"
	player "github.com/gpnaslund/freja_monya_platformer/internal/player/monya"
	"github.com/gpnaslund/freja_monya_platformer/internal/util"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	tileSize     = 16
	screenWidth  = 320
	screenHeight = 240
)

type Level struct {
	tileset      *ebiten.Image
	layers       [][]int
	collisionBox util.CollisionBox
	monya        *player.Monya
}

func NewLevel() *Level {
	lvl := Level{}
	lvl.loadImage()
	lvl.mapTiles()
	lvl.collisionBox = util.CollisionBox{
		Position: &util.Vector{
			X: screenWidth / 2,
			Y: screenHeight - (tileSize / 2),
		},
		Width:  screenWidth,
		Height: tileSize,
	}
	monya := player.NewMonya(&util.Vector{X: 20, Y: screenHeight - 50})
	lvl.monya = monya
	return &lvl
}

func (l *Level) loadImage() {
	fileData, err := levels.LevelAssets.ReadFile("resources/Tileset.png")
	if err != nil {
		log.Fatal("Level 1: Failed to load TilesetGround")
	}

	img, _, err := image.Decode(bytes.NewReader(fileData))
	if err != nil {
		log.Fatal("Level 1: Failed to decode TilesetGround")
	}

	l.tileset = ebiten.NewImageFromImage(img)
}

func (l *Level) mapTiles() {
	l.layers = [][]int{
		{
			48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48,
			48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48,
			48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48,
			48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48,
			48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48,
			48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48,
			48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48,
			48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48,
			48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48,
			48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48,
			48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48,
			48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48,
			48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48,
			48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		},
	}

}

func (l *Level) Update() error {
	l.monya.Update()
	l.monya.IsOnGround = l.collisionBox.IsColliding(l.monya.CollisionBox)
	return nil
}

func (l *Level) Draw(screen *ebiten.Image) {
	w := l.tileset.Bounds().Dx()
	tileXCount := w / tileSize
	const xCount = screenWidth / tileSize
	for _, layer := range l.layers {
		for i, t := range layer {
			options := &ebiten.DrawImageOptions{}
			options.GeoM.Translate(float64((i%xCount)*tileSize), float64((i/xCount)*tileSize))

			sx := (t % tileXCount) * tileSize
			sy := (t / tileXCount) * tileSize
			screen.DrawImage(l.tileset.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), options)
		}
	}
	l.monya.Draw(screen, false)
}
