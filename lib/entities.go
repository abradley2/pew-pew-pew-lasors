package lib

// Entity : basis wrapper for all different type of ships
type Entity interface {
	update()
}

// Xwing : contains rebel scum
type Xwing struct {
	active                 bool
	sprite                 string
	xPos, yPos, xVel, yVel float64
	width, height          int
}

// Tie : fighter of choice for our brave troops
type Tie struct {
	active                 bool
	sprite                 string
	xPos, yPos, xVel, yVel float64
	width, height          int
}

func (x Xwing) update() {

}

func (t Tie) update() {

}
