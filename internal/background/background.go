package background

import (
	"fmt"

	"spaceship/internal/object"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var images = map[int]string {
	1:"internal/assets/space.png",
}

type Background struct {
	image *ebiten.Image
}

func New() object.Object {
	b := new(Background)
	return b
}

func (b *Background) SetImage(level int) error {
	var err error

	switch(level) {
	case 1:
		b.image, _, err = ebitenutil.NewImageFromFile(images[level])
	default:
		err = fmt.Errorf("level does not exist")
	}

	return err
}

func (b *Background) GetImage() *ebiten.Image {
	return b.image
}

func (b *Background) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0.0, 0.0)
	screen.DrawImage(b.image, op)
}

func (b *Background) Update() error {
	return nil
}