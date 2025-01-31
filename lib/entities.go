package lib

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

// Entity : basis wrapper for all different type of ships
type Entity interface {
	Update()
	Remove()
	GetCoords() (float64, float64, float64, float64, float64)
}

// Xwing : contains rebel scum
type Xwing struct {
	Exploding                             bool
	ExplosionFrame                        int
	Team                                  string
	ShotRequested                         bool
	SpawnQueued                           bool
	ShotQueued                            bool
	Sprite                                *ebiten.Image
	Active                                bool
	Xpos, Ypos, Xvel, Yvel, Width, Height float64
}

// Tie : fighter of choice for our brave troops
type Tie struct {
	Exploding                             bool
	ExplosionFrame                        int
	Team                                  string
	ShotRequested                         bool
	SpawnQueued                           bool
	ShotQueued                            bool
	Sprite                                *ebiten.Image
	Active                                bool
	Xpos, Ypos, Xvel, Yvel, Width, Height float64
}

// Missile : a missile shot by either a tie or an xwing
type Missile struct {
	Team                                  string
	Sprite                                *ebiten.Image
	Active                                bool
	Xpos, Ypos, Xvel, Yvel, Width, Height float64
}

// Spawn : spawn a missile at the given coords
func (m *Missile) Spawn(team string, x float64, y float64, yvel float64) {
	m.Active = true
	m.Team = team
	m.Xpos = x
	m.Ypos = y
	m.Yvel = yvel
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

// Remove : remove the missile
func (m *Missile) Remove() {
	m.Active = false
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
		return
	}
	if !t.ShotQueued {
		t.ShotQueued = true
		time.AfterFunc(time.Duration(1500*rand.Float64())*time.Millisecond, t.RequestShot)
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
		return
	}
	if !x.ShotQueued {
		x.ShotQueued = true
		time.AfterFunc(time.Duration(1500*rand.Float64())*time.Millisecond, x.RequestShot)
	}
	x.Ypos -= 10
}

// Update : update a missile
func (m *Missile) Update() {
	if !m.Active {
		return
	}
	m.Ypos += (30 * m.Yvel)

	if m.Ypos < -100 || m.Ypos > GameHeight+100 || m.Xpos < -100 || m.Xpos > GameWidth+100 {
		m.Remove()
	}
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

// RequestShot : set shot requested true so game loop can add a missile to the game
func (t *Tie) RequestShot() {
	t.ShotRequested = true
	t.ShotQueued = false
}

// RequestShot : set shot requested true so game loop can add a missile to the game
func (x *Xwing) RequestShot() {
	x.ShotRequested = true
	x.ShotQueued = false
}

// GetCoords : get x, y, w, h for a tie
func (t *Tie) GetCoords() (float64, float64, float64, float64, float64) {
	return t.Xpos - t.Width, t.Ypos - t.Height, t.Width, t.Height, -1
}

// GetCoords : get x, y, w, h for an xwing
func (x *Xwing) GetCoords() (float64, float64, float64, float64, float64) {
	return x.Xpos, x.Ypos, x.Width, x.Height, 1
}

// GetCoords : get x, y, w, h for a missile
func (m *Missile) GetCoords() (float64, float64, float64, float64, float64) {
	return m.Xpos, m.Ypos, m.Width, m.Height, 1
}

// Explode : blow up an xwing
func (x *Xwing) Explode() {
	x.Exploding = true
	x.ExplosionFrame = 0
}

// Explode : blow up a tiefighter
func (t *Tie) Explode() {
	t.Exploding = true
	t.ExplosionFrame = 0
}
