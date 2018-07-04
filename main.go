package main

import (
	"go-game/lib"

	"github.com/hajimehoshi/ebiten"
)

type entityGroup interface {
	updateEntity()
}

// Xwings : contains a slice of 40 xwing entities
type Xwings [40]lib.Xwing

// Ties : contains a slice of 40 tie entities
type Ties [40]lib.Tie

var (
	xwingSprite *ebiten.Image
	tieSprite   *ebiten.Image
	xwings      Xwings
	ties        Ties
	op          = &ebiten.DrawImageOptions{}
)

func (x Xwings) updateEntity() {
	for _, x := range x {
		x.Update()
	}
}

func (t Ties) updateEntity() {
	for _, t := range t {
		t.Update()
	}
}

func update(screen *ebiten.Image) error {
	xwings.updateEntity()
	ties.updateEntity()

	if ebiten.IsRunningSlowly() {
		return nil
	}

	for i := 0; i < len(ties); i++ {
		screen.DrawImage(xwingSprite, op)
		op.GeoM.Reset()
	}

	return nil
}

func main() {
	xwingSprite, _ = ebiten.NewImageFromImage(*lib.Images["/assets/xwing-smol.png"], ebiten.FilterDefault)
	tieSprite, _ = ebiten.NewImageFromImage(*lib.Images["/assets/tie-smol.png"], ebiten.FilterDefault)

	ebiten.Run(update, lib.GameWidth, lib.GameHeight, 2, "Hello world!")
}
