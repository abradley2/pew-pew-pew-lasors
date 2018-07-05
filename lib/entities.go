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

// Remove : remove the tie fighter
func (t *Tie) Remove() {
	t.SpawnQueued = false
	t.Active = false
}

// Remove : remove the xwing
func (x *Xwing) Remove() {
	x.SpawnQueued = false
	x.Active = false
}

// Update : update the tie fighter
func (t *Tie) Update() {
	if !t.Active {
		if t.SpawnQueued {
			return
		}
		t.SpawnQueued = true
		time.AfterFunc(time.Duration(7000*rand.Float64())*time.Millisecond, t.Spawn)
		return
	}
	if t.Ypos > float64(GameHeight)+100 {
		t.Remove()
	}
	t.Ypos += 10
}

// Update : update the xwing
func (x *Xwing) Update() {
	if !x.Active {
		if x.SpawnQueued {
			return
		}
		x.SpawnQueued = true
		time.AfterFunc(time.Duration(7000*rand.Float64())*time.Millisecond, x.Spawn)
		return
	}
	if x.Ypos < -100 {
		x.Remove()
	}
	x.Ypos -= 10
}

// Spawn : spawn the tie fighter
func (t *Tie) Spawn() {
	w, h := t.Sprite.Size()
	t.SpawnQueued = false
	t.Active = true
	t.Xpos = (rand.Float64() * float64(GameWidth-w)) + float64(w)
	t.Ypos = 0 - float64(h)
}

// Spawn : spawn the xwing
func (x *Xwing) Spawn() {
	w, h := x.Sprite.Size()
	x.SpawnQueued = false
	x.Active = true
	x.Xpos = (rand.Float64() * float64(GameWidth-w)) + float64(w)
	x.Ypos = float64(GameHeight + h)
}
