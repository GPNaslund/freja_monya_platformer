package level1

import (
	"bytes"
	"image"
	"log"
	"math"

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
	camera       util.Camera
}

func NewLevel() *Level {
	lvl := Level{}
	lvl.loadImage()
	lvl.layers = lvl1map
	lvl.collisionBox = util.CollisionBox{
		Position: &util.Vector{
			X: 0,
			Y: screenHeight - tileSize,
		},
		Width:  screenWidth * 2,
		Height: tileSize,
	}
	monya := player.NewMonya(&util.Vector{X: 0, Y: screenHeight - 50}, screenWidth/2)
	lvl.monya = monya
	lvl.camera = util.Camera{
		Width:  screenWidth,
		Height: screenHeight,
		MaxX:   screenWidth * 2,
		MaxY:   screenHeight,
	}
	lvl.camera.CenterOnEntity(monya.GetXAndYCoordinates())
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

func (l *Level) Update() error {
	l.monya.Update()
	l.monya.IsOnGround = l.collisionBox.IsColliding(l.monya.CollisionBox)
	l.camera.CenterOnEntity(l.monya.GetXAndYCoordinates())
	return nil
}

func (l *Level) Draw(screen *ebiten.Image) {
	startX := int(math.Floor(l.camera.X / tileSize))
	endX := int(math.Ceil((l.camera.X + l.camera.Width) / tileSize))
	startY := int(math.Floor(l.camera.Y / tileSize))
	endY := int(math.Ceil((l.camera.Y + l.camera.Height) / tileSize))

	for y := startY; y < endY; y++ {
		if y < 0 || y >= len(l.layers) {
			continue
		}
		for x := startX; x < endX; x++ {
			if x < 0 || x >= len(l.layers[y]) {
				continue
			}
			tile := l.layers[y][x]
			options := &ebiten.DrawImageOptions{}
			options.GeoM.Translate(float64(x*tileSize)-l.camera.X, float64(y*tileSize)-l.camera.Y)

			sx := (tile % (l.tileset.Bounds().Dx() / tileSize)) * tileSize
			sy := (tile / (l.tileset.Bounds().Dx() / tileSize)) * tileSize
			screen.DrawImage(l.tileset.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), options)
		}
	}
	l.collisionBox.DebugTiles(screen, l.camera.X, l.camera.Y)
	l.monya.Draw(screen, true)
}
