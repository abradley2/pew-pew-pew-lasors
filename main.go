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
type Xwings [80]*lib.Xwing

// Ties : contains a slice of 40 tie entities
type Ties [80]*lib.Tie

// Missiles ; contains a slice of 200 missile entities
type Missiles [400]*lib.Missile

var (
	redLazorSprite   *ebiten.Image
	greenLazorSprite *ebiten.Image
	xwingSprite      *ebiten.Image
	tieSprite        *ebiten.Image
	xwings           Xwings
	ties             Ties
	missiles         Missiles
	op               = &ebiten.DrawImageOptions{}
)

func (x Xwings) updateEntity() {
	for _, x := range x {
		x.Update()
	}
}

func (t Ties) updateEntity() {
	for i := 0; i < len(t); i++ {
		t[i].Update()
	}
}

func (m Missiles) updateEntity() {
	for i := 0; i < len(m); i++ {
		m[i].Update()
	}
}

func update(screen *ebiten.Image) error {
	if ebiten.IsRunningSlowly() {
		return nil
	}

	xwings.updateEntity()
	ties.updateEntity()
	missiles.updateEntity()

	for i := 0; i < len(ties); i++ {
		t := ties[i]
		if t.Active != true {
			continue
		}
		op.GeoM.Reset()
		op.GeoM.Rotate(math.Pi)
		op.GeoM.Translate(t.Xpos, t.Ypos)
		screen.DrawImage(tieSprite, op)
		if t.ShotRequested {
			fireZeMissiles("tie", t.Xpos, t.Ypos, 1)
			t.ShotRequested = false
		}
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
		if x.ShotRequested {
			fireZeMissiles("xwing", x.Xpos, x.Ypos, -1)
			x.ShotRequested = false
		}
	}

	for i := 0; i < len(missiles); i++ {
		m := missiles[i]
		if m.Active != true {
			continue
		}
		op.GeoM.Reset()
		op.GeoM.Rotate(0)
		op.GeoM.Translate(m.Xpos, m.Ypos)
		var sprite *ebiten.Image

		switch m.Team {
		case "xwing":
			sprite = redLazorSprite
		case "tie":
			sprite = greenLazorSprite
		}

		screen.DrawImage(sprite, op)
	}

	return nil
}

func fireZeMissiles(team string, xPos float64, yPos float64, yVel float64) {
	var found bool
	var idx int
	for found == false && idx < len(missiles) {
		m := missiles[idx]
		if !m.Active {
			found = true
			m.Spawn(team, xPos, yPos, yVel)
		}
		idx++
	}
}

func main() {
	redLazorSprite, _ = ebiten.NewImageFromImage(*lib.Images["/assets/lazor-red.png"], ebiten.FilterDefault)
	greenLazorSprite, _ = ebiten.NewImageFromImage(*lib.Images["/assets/lazor-green.png"], ebiten.FilterDefault)
	xwingSprite, _ = ebiten.NewImageFromImage(*lib.Images["/assets/xwing-smol.png"], ebiten.FilterDefault)
	tieSprite, _ = ebiten.NewImageFromImage(*lib.Images["/assets/tie-smol.png"], ebiten.FilterDefault)

	for i := range xwings {
		w, h := xwingSprite.Size()
		createXwing := new(lib.Xwing)
		createXwing.Active = false
		createXwing.Sprite = xwingSprite
		createXwing.Width = float64(w)
		createXwing.Height = float64(h)
		xwings[i] = createXwing
	}
	for i := range ties {
		w, h := tieSprite.Size()
		createTie := new(lib.Tie)
		createTie.Active = false
		createTie.Sprite = tieSprite
		createTie.Width = float64(w)
		createTie.Height = float64(h)
		ties[i] = createTie
	}
	for i := range missiles {
		createMissile := new(lib.Missile)
		missiles[i] = createMissile
	}

	ebiten.Run(update, lib.GameWidth, lib.GameHeight, 0.5, "Hello world!")
}
