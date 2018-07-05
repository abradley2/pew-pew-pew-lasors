package main

import (
	"go-game/lib"
	"math"

	"github.com/hajimehoshi/ebiten"
)

type entityGroup interface {
	updateEntity()
}

// Xwings : contains a slice of 40 xwing entities
type Xwings [50]*lib.Xwing

// Ties : contains a slice of 40 tie entities
type Ties [50]*lib.Tie

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
	for i := 0; i < len(ties); i++ {
		ties[i].Update()
	}
}

func update(screen *ebiten.Image) error {
	xwings.updateEntity()
	ties.updateEntity()

	if ebiten.IsRunningSlowly() {
		return nil
	}

	for i := 0; i < len(ties); i++ {
		t := ties[i]
		if t.Active != true {
			continue
		}
		op.GeoM.Reset()
		op.GeoM.Rotate(math.Pi)
		op.GeoM.Translate(t.Xpos, t.Ypos)
		screen.DrawImage(tieSprite, op)
	}

	for i := 0; i < len(xwings); i++ {
		x := xwings[i]
		if x.Active != true {
			continue
		}
		op.GeoM.Reset()
		op.GeoM.Rotate(0)
		op.GeoM.Translate(x.Xpos, x.Ypos)
		screen.DrawImage(xwingSprite, op)
	}

	return nil
}

func main() {
	xwingSprite, _ = ebiten.NewImageFromImage(*lib.Images["/assets/xwing-smol.png"], ebiten.FilterDefault)
	tieSprite, _ = ebiten.NewImageFromImage(*lib.Images["/assets/tie-smol.png"], ebiten.FilterDefault)

	for i := range xwings {
		createXwing := new(lib.Xwing)
		createXwing.Active = false
		createXwing.Sprite = xwingSprite
		xwings[i] = createXwing
	}
	for i := range ties {
		createTie := new(lib.Tie)
		createTie.Active = false
		createTie.Sprite = tieSprite
		ties[i] = createTie
	}

	ebiten.Run(update, lib.GameWidth, lib.GameHeight, 0.5, "Hello world!")
}
