package player

import (
	"fmt"
	"spaceship/internal/bullet"
	"spaceship/internal/object"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var images = map[int]string{
	1: "internal/assets/spaceship.png",
}

type Player struct {
	image      *ebiten.Image
	maxX, maxY int
	xPos, yPos float64
	speed      float64
	bullets    []*bullet.Bullet
}

func New(maxX, maxY int, speed float64) (object.Object, error) {
	p := new(Player)
	p.maxX = maxX
	p.maxY = maxY

	if err := p.SetImage(1); err != nil {
		return nil, err
	}

	p.xPos = (float64(maxX) / 2) - (float64(p.image.Bounds().Dx()) / 2)
	p.yPos = (float64(maxY) / 2) + ((float64(maxY) / 2) / 2)
	p.speed = speed

	p.bullets = make([]*bullet.Bullet, 0)
	return p, nil
}

func (p *Player) SetImage(status int) error {
	var err error

	switch status {
	case 1:
		p.image, _, err = ebitenutil.NewImageFromFile(images[status])
	default:
		err = fmt.Errorf("status does not exist")
	}

	return err
}

func (p *Player) GetImage() *ebiten.Image {
	return p.image
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.xPos, p.yPos)
	screen.DrawImage(p.image, op)

	for _, b := range p.bullets {
		b.Draw(screen)
	}
}

func (p *Player) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.yPos -= p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.yPos += p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.xPos -= p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.xPos += p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		p.NewBullet()
	}

	// check bounds
	if p.xPos <= 0 {
		p.xPos = 0
	}
	if (p.xPos + float64(p.image.Bounds().Dx())) >= float64(p.maxX) {
		p.xPos = float64(p.maxX) - float64(p.image.Bounds().Dx())
	}
	if p.yPos <= 0 {
		p.yPos = 0
	}
	if (p.yPos + float64(p.image.Bounds().Dy())) >= float64(p.maxY) {
		p.yPos = float64(p.maxY) - float64(p.image.Bounds().Dy())
	}

	for _, b := range p.bullets {
		b.Update()
	}

	return nil
}

func (p *Player) NewBullet() {
	newBullet, _ := bullet.New(p.xPos+(float64(p.image.Bounds().Dx()/2)-5), p.yPos)

	p.bullets = append(p.bullets, newBullet)
}
