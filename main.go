package main

import (
	_ "image/jpeg"
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ably-labs/rosie-demo/config"
)

var (
	state gameState
)

func init() {
	state = titleScreen
}

type Game struct {}

// NewGame is a constructor for the game.
func NewGame() *Game {
	return &Game{}
}

//Update updates the logical state.
func (g *Game) Update() error {

	// Handle updates for each game state.
	switch state {
	case titleScreen:
		updateTitleScreen()
	case realtimeScreen:
		updateRealtimeScreen()
	case restScreen:
		updateRestScreen()
	}

	return nil
}

//Draw renders the screen.
func (g *Game) Draw(screen *ebiten.Image) {

	//Draw debug elements if debug mode is on.
	if config.Cfg.DebugMode {
		drawDebugText(screen)
	}

	//Handle drawing for each game state.
	switch state {
	case titleScreen:
		drawTitleScreen(screen)
	case realtimeScreen:
		drawRealtimeScreen(screen)
	case restScreen:
		drawRestScreen(screen)
	}
}

//Layout returns the logical screen size, the screen is automatically scaled.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle(titleText)

	// initialisation
	initialiseTitleScreen()
	initialiseRealtimeScreen()

	// Create a new instance of game.
	game := NewGame()

	// Run the game.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
