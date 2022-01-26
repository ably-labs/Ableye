package main

import (
	"errors"
	_ "image/jpeg"
	"log"

	"github.com/ably-labs/Ableye/config"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	state gameState
)

func init() {
	state = titleScreen
}

type Game struct{}

// NewGame is a constructor for the game.
func NewGame() *Game {
	return &Game{}
}

//handleClose is called when the used closes the game window.
func handleClose() {
	connectionIDs := []connectionID{clientA, clientB, clientC, clientD}
	for _, id := range connectionIDs {
		if connections[id] != nil && connections[id].client != nil {
			connections[id].client.Close()
		}
	}
}

//Update updates the logical state.
func (g *Game) Update() error {

	if ebiten.IsWindowBeingClosed() {
		handleClose()
		// An error must be returned to trigger the window closing once closing has been handled.
		return errors.New("window has been closed")
	}
	// Handle updates for each game state.
	switch state {
	case titleScreen:
		updateTitleScreen()
	case realtimeScreen:
		updateRealtimeScreen()
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

	// Set window closing handled to true.
	ebiten.SetWindowClosingHandled(true)

	// Run the game.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
