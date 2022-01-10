package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// The elements of the rest screen.
var (

)

func drawRestScreen(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "Ably Rest", screenWidth/2, 50)
}

func updateRestScreen(){
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		state = titleScreen
	}
}