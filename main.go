package main

import (
	"fmt"
	"go-game/lib"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	xwings = make([]*lib.Entity, 40)
	ties   = make([]*lib.Entity, 40)
)

func update(screen *ebiten.Image) error {
	for path := range lib.Images {
		fmt.Printf("\n\nthe path is : %s\n\n", path)
	}
	ebitenutil.DebugPrint(screen, "Hello world! again")
	return nil
}

func main() {
	// lets make some slices for storing entities

	ebiten.Run(update, lib.GameWidth, lib.GameHeight, 2, "Hello world!")
}
