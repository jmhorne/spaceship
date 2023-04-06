package bullet

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Bullet struct {
	image *ebiten.Image
	xPos, yPos float64
}

func New(xPos, yPos float64) (*Bullet, error) {
	b := new(Bullet)
	b.xPos, b.yPos = xPos, yPos
	err := b.SetImage(0)
	return b, err
}

func (b *Bullet) SetImage(image int) error {
	var err error

	b.image, _, err = ebitenutil.NewImageFromFile("internal/assets/bullet.png")
	
	return err
}

func (b *Bullet) GetImage() *ebiten.Image {
	return b.image
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.xPos, b.yPos)
	screen.DrawImage(b.image, op)
}

func (b *Bullet) Update() error {
	b.yPos -= 5
	return nil
}
