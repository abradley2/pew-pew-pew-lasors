package lib

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

// Entity : basis wrapper for all different type of ships
type Entity interface {
	Update()
	Spawn()
	Remove()
}

// Xwing : contains rebel scum
type Xwing struct {
	SpawnQueued                           bool
	ShotQueued                            bool
	Sprite                                *ebiten.Image
	Active                                bool
	Xpos, Ypos, Xvel, Yvel, Width, Height float64
}

// Tie : fighter of choice for our brave troops
type Tie struct {
	SpawnQueued                           bool
	ShotQueued                            bool
	Sprite                                *ebiten.Image
	Active                                bool
	Xpos, Ypos, Xvel, Yvel, Width, Height float64
}

// Update : update the xwing
func (x *Xwing) Update() {

}

// Remove : remove the tie fighter
func (t *Tie) Remove() {
	t.SpawnQueued = false
	t.Active = false
}

// Update : update the tie fighter
func (t *Tie) Update() {
	_, h := t.Sprite.Size()
	if !t.Active {
		if t.SpawnQueued {
			return
		}
		t.SpawnQueued = true
		time.AfterFunc(time.Duration(7000*rand.Float64())*time.Millisecond, t.Spawn)
		return
	}
	if t.Ypos > float64(GameHeight)+float64(h) {
		t.Remove()
	}
	t.Ypos += 10
}

// Spawn : spawn the tie fighter
func (t *Tie) Spawn() {
	_, h := t.Sprite.Size()
	t.SpawnQueued = false
	t.Active = true
	t.Xpos = rand.Float64() * GameWidth
	t.Ypos = 0 - float64(h)
}
