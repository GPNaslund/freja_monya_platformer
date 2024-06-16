package player

import (
	_ "image/png"
	"log"

	"github.com/gpnaslund/freja_monya_platformer/internal/util"
	"github.com/hajimehoshi/ebiten/v2"
)

type Monya struct {
	count         int
	position      util.Vector
	collisionBox  util.CollisionBox
	idleAnimation util.Animation
	walkAnimation util.Animation
}

func NewMonya(position util.Vector) *Monya {
	monya := &Monya{
		count:    0,
		position: position,
		collisionBox: util.CollisionBox{
			Position: position,
			Width:    25,
			Height:   19,
		},
	}
	monya.createIdleAnimation()
	monya.createWalkAnimation()
	return monya
}

func (m *Monya) Update() error {
	m.count++
	return nil
}

func (m *Monya) Draw(screen *ebiten.Image, debug bool) {
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		m.drawWalkAnimation(screen, false)
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		m.drawWalkAnimation(screen, true)
	} else {
		m.drawIdleAnimation(screen)
	}

	if debug {
		m.collisionBox.Debug(screen)
	}
}

func (m *Monya) drawIdleAnimation(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(-float64(m.idleAnimation.FrameWidth)/2, -float64(m.idleAnimation.FrameHeight)/2)
	options.GeoM.Translate(m.position.X, m.position.Y)
	idleFrame := m.idleAnimation.GetFrame(m.count)
	screen.DrawImage(idleFrame, options)
}

func (m *Monya) createIdleAnimation() {
	idleSpriteSheet, err := util.LoadSprite(assets, "resources/IdleMod.png")
	if err != nil {
		log.Fatal("Failed to load Monya Idle animation")
	}
	m.idleAnimation = util.Animation{
		SpriteSheet: idleSpriteSheet,
		Frame0X:     0,
		Frame0Y:     0,
		FrameWidth:  31,
		FrameHeight: 20,
		FrameCount:  4,
		FrameSpeed:  10,
	}
}

func (m *Monya) drawWalkAnimation(screen *ebiten.Image, flipSprite bool) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(-float64(m.walkAnimation.FrameWidth)/2, -float64(m.walkAnimation.FrameHeight)/2)
	if flipSprite {
		options.GeoM.Scale(-1, 1)
	}
	options.GeoM.Translate(m.position.X, m.position.Y)
	walkingFrame := m.walkAnimation.GetFrame(m.count)
	screen.DrawImage(walkingFrame, options)
}

func (m *Monya) createWalkAnimation() {
	runningSpriteSheet, err := util.LoadSprite(assets, "resources/WalkMod.png")
	if err != nil {
		log.Fatal("Failed to load Monya Walk animation")
	}
	m.walkAnimation = util.Animation{
		SpriteSheet: runningSpriteSheet,
		Frame0X:     0,
		Frame0Y:     0,
		FrameWidth:  36,
		FrameHeight: 20,
		FrameCount:  6,
		FrameSpeed:  10,
	}
}
