package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/ably-labs/rosie-demo/config"
)

func drawDebugText(screen *ebiten.Image) {
	// Draw cursor co-ordinates on screen
	x, y := ebiten.CursorPosition()
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("X: %d, Y: %d", x, y), screenWidth/2, 0)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Version: %s", config.Cfg.Version), screenWidth-(len(config.Cfg.Version)*20), 0)
}
