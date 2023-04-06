package game

import (
	"spaceship/internal/background"
	"spaceship/internal/object"
	"spaceship/internal/player"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	WIDTH, HEIGHT = 640, 480
)

type Game struct {
	level      int
	background object.Object
	player     object.Object
}

func New(level ...int) (*Game, error) {
	g := new(Game)

	if len(level) == 0 {
		g.level = 1
	} else {
		g.level = level[0]
	}

	g.background = background.New()
	var err error
	if g.player, err = player.New(WIDTH, HEIGHT, 4.0); err != nil {
		return nil, err
	}

	if err := g.background.SetImage(g.level); err != nil {
		return nil, err
	}

	if err := g.player.SetImage(g.level); err != nil {
		return nil, err
	}

	return g, nil
}

func (g *Game) Update() error {
	g.background.Update()
	g.player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.background.Draw(screen)
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WIDTH, HEIGHT
}
