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
	xwings Xwings
	ties   Ties
	op     = &ebiten.DrawImageOptions{}
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

	// Draw each sprite.
	// DrawImage can be called many many times, but in the implementation,
	// the actual draw call to GPU is very few since these calls satisfy
	// some conditions e.g. all the rendering sources and targets are same.
	// For more detail, see:
	// https://godoc.org/github.com/hajimehoshi/ebiten#Image.DrawImage
	for i := 0; i < len(ties); i++ {
		op.GeoM.Reset()
		screen.DrawImage(lib.EbitenImage, op)
	}

	return nil
}

func main() {
	for _, xwing := range xwings {
		xwing.Sprite = lib.Images["/assets/xwing-smol.png"]
	}
	for _, tie := range ties {
		tie.Sprite = lib.Images["/assets/tie-smol.png"]
	}

	ebiten.Run(update, lib.GameWidth, lib.GameHeight, 2, "Hello world!")
}
