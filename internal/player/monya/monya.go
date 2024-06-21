package player

import (
	"fmt"
	_ "image/png"
	"log"

	"github.com/gpnaslund/freja_monya_platformer/internal/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	walkingSpeed = 3
	jumpSpeed    = 10
	gravity      = 1.5
)

type Monya struct {
	count           int
	position        *util.Vector
	maxDrawX        float64
	velocity        *util.Velocity
	CollisionBox    util.CollisionBox
	idleAnimation   util.Animation
	walkAnimation   util.Animation
	IsOnGround      bool
	facingBackwards bool
}

func NewMonya(position *util.Vector, maxDrawX float64) *Monya {
	monya := &Monya{
		count:    0,
		position: position,
		maxDrawX: maxDrawX,
		velocity: &util.Velocity{
			X: 0,
			Y: 0,
		},
		CollisionBox: util.CollisionBox{
			Position: &util.Vector{
				X: position.X - 25/2,
				Y: position.Y - 19/2,
			},
			Width:  25,
			Height: 19,
		},
		IsOnGround:      false,
		facingBackwards: false,
	}
	monya.createIdleAnimation()
	monya.createWalkAnimation()
	return monya
}

func (m *Monya) Update() error {
	m.count++
	m.handleGravity()
	m.handleMovement()
	m.position.X += float64(m.velocity.X)
	m.position.Y += float64(m.velocity.Y)
	m.CollisionBox.Position.X = m.position.X - 25/2
	m.CollisionBox.Position.Y = m.position.Y - 19/2
	return nil
}

func (m *Monya) Draw(screen *ebiten.Image, debug bool) {
	if m.velocity.X == 0 && m.velocity.Y == 0 {
		m.drawIdleAnimation(screen, m.facingBackwards)
	} else {
		m.drawWalkAnimation(screen, m.facingBackwards)
	}
	if debug {
		m.CollisionBox.DebugEntity(screen, m.maxDrawX-25/2)
		debugStr := fmt.Sprintf("Monya; X: %f, Y: %f", m.position.X, m.position.Y)
		ebitenutil.DebugPrint(screen, debugStr)
	}
}

func (m *Monya) GetXAndYCoordinates() (x, y float64) {
	return m.position.X, m.position.Y
}

func (m *Monya) handleGravity() {
	if m.IsOnGround {
		m.velocity.Y = 0
	} else if m.velocity.Y < gravity {
		m.velocity.Y += gravity
	}
}

func (m *Monya) handleMovement() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		m.velocity.X = walkingSpeed
		m.facingBackwards = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		m.velocity.X = -walkingSpeed
		m.facingBackwards = true
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if m.IsOnGround {
			m.velocity.Y = -jumpSpeed
		}
	}
	if !ebiten.IsKeyPressed(ebiten.KeyArrowRight) && !ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		m.velocity.X = 0
	}
}

func (m *Monya) drawIdleAnimation(screen *ebiten.Image, flipSprite bool) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(-float64(m.idleAnimation.FrameWidth)/2, -float64(m.idleAnimation.FrameHeight)/2)
	if flipSprite {
		options.GeoM.Scale(-1, 1)
	}
	if m.position.X > m.maxDrawX {
		options.GeoM.Translate(m.maxDrawX, m.position.Y)
	} else {
		options.GeoM.Translate(m.position.X, m.position.Y)
	}
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
	if m.position.X > m.maxDrawX {
		options.GeoM.Translate(m.maxDrawX, m.position.Y)
	} else {
		options.GeoM.Translate(m.position.X, m.position.Y)
	}
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
