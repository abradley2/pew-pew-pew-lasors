package main

import (
	"go-game/lib"
	"image"
	"math"

	"github.com/disintegration/gift"
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
	explosionSprites [50]*ebiten.Image
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
			if t.Exploding {
				if t.ExplosionFrame == len(explosionSprites) {
					t.Exploding = false
					t.ExplosionFrame = 0
					continue
				}
				op.GeoM.Reset()
				op.GeoM.Rotate(math.Pi)
				op.GeoM.Translate(t.Xpos, t.Ypos)
				screen.DrawImage(explosionSprites[t.ExplosionFrame], op)
				t.ExplosionFrame++
			}
			continue
		}
		op.GeoM.Reset()
		op.GeoM.Rotate(math.Pi)
		op.GeoM.Translate(t.Xpos, t.Ypos)
		screen.DrawImage(tieSprite, op)
		if t.ShotRequested {
			fireZeMissiles("tie", t.Xpos-(t.Width/2), t.Ypos, 1)
			t.ShotRequested = false
		}
	}

	for i := 0; i < len(xwings); i++ {
		x := xwings[i]
		if x.Active != true {
			if x.Exploding {
				if x.ExplosionFrame == len(explosionSprites) {
					x.Exploding = false
					x.ExplosionFrame = 0
					continue
				}
				op.GeoM.Reset()
				op.GeoM.Rotate(0)
				op.GeoM.Translate(x.Xpos, x.Ypos)
				screen.DrawImage(explosionSprites[x.ExplosionFrame], op)
				x.ExplosionFrame++
			}
			continue
		}
		op.GeoM.Reset()
		op.GeoM.Rotate(0)
		op.GeoM.Translate(x.Xpos, x.Ypos)
		screen.DrawImage(xwingSprite, op)
		if x.ShotRequested {
			fireZeMissiles("xwing", x.Xpos+(x.Width/2), x.Ypos, -1)
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

		//
		if m.Team == "xwing" {
			for _, t := range ties {
				if t.Active && checkCollision(t, m) {
					t.Remove()
					t.Explode()
					m.Remove()
				}
			}
			screen.DrawImage(redLazorSprite, op)
		}

		if m.Team == "tie" {
			for _, x := range xwings {
				if x.Active && checkCollision(x, m) {
					x.Remove()
					x.Explode()
					m.Remove()
				}
			}
			screen.DrawImage(greenLazorSprite, op)
		}
	}

	return nil
}

func checkCollision(entity1 lib.Entity, entity2 lib.Entity) bool {
	x1, y1, w1, h1, _ := entity1.GetCoords()
	x2, y2, w2, h2, _ := entity2.GetCoords()
	var collision bool

	if x1 < x2+(w2) &&
		x1+(w1) > x2 &&
		y1 < y2+(h2) &&
		(h1)+y1 > y2 {
		collision = true
	}

	return collision
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
	explosionImg := *lib.Images["/assets/explosion.png"]
	explosionSpriteSheet, _ := ebiten.NewImageFromImage(explosionImg, ebiten.FilterDefault)

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
		w, h := redLazorSprite.Size()
		createMissile := new(lib.Missile)
		createMissile.Width = float64(w)
		createMissile.Height = float64(h)
		missiles[i] = createMissile
	}

	// create the explosion sprite sheet
	// seperate this out into a convenience function later
	w, h := explosionSpriteSheet.Size()

	heightCut := h / 5
	widthCut := w / 5
	row := 0
	column := 0

	for i := 0; i < len(explosionSprites); i += 2 {
		genImg := image.NewRGBA(image.Rect(0, 0, widthCut, heightCut))
		rect := image.Rect(column*widthCut, row*heightCut, column*widthCut+widthCut, row*heightCut+heightCut)
		filter := gift.New(gift.Crop(rect))

		filter.Draw(genImg, explosionImg)

		bigImg := image.NewRGBA(image.Rect(0, 0, 100, 100))
		bigFilter := gift.New(gift.Resize(100, 100, gift.LanczosResampling))
		bigFilter.Draw(bigImg, genImg)

		resultImg, _ := ebiten.NewImageFromImage(bigImg, ebiten.FilterDefault)
		explosionSprites[i] = resultImg
		explosionSprites[i+1] = resultImg
		column++
		if column == 6 {
			column = 0
			row++
		}
	}

	ebiten.Run(update, lib.GameWidth, lib.GameHeight, 0.5, "Hello world!")
}
