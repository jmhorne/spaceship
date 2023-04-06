package object

import "github.com/hajimehoshi/ebiten/v2"

type Object interface {
	SetImage(level int) error
	GetImage() *ebiten.Image
	Draw(screen *ebiten.Image)
	Update() error
}