package lib

import (
	"github.com/hajimehoshi/ebiten"
)

// Entity : basis wrapper for all different type of ships
type Entity interface {
	Update()
}

// Xwing : contains rebel scum
type Xwing struct {
	Sprite                 *ebiten.Image
	Active                 bool
	Xpos, Ypos, Xvel, Yvel float64
	Width, Height          int
}

// Tie : fighter of choice for our brave troops
type Tie struct {
	Sprite                 *ebiten.Image
	Active                 bool
	Xpos, Ypos, Xvel, Yvel float64
	Width, Height          int
}

// Update : update the xwing
func (x *Xwing) Update() {

}

// Update : update the tie fighter
func (t *Tie) Update() {
	t.Xpos += 2
	t.Ypos += 2
}
