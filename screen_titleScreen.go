package main

import (
	"github.com/ably-labs/rosie-demo/button"
	colour "github.com/ably-labs/rosie-demo/colours"

	font "github.com/ably-labs/rosie-demo/fonts"
	"github.com/hajimehoshi/ebiten/v2"
)

// The elements of the title screen.
var (
	realtimeButton button.Button
)

func initialiseTitleScreen() {
	realtimeButton = button.NewButton(200, 100, "Ably Realtime", 25, 55, colour.White, font.MplusNormalFont, colour.BrightRed, (screenWidth/2)-100, (screenHeight/2)-100)
}

func drawTitleScreen(screen *ebiten.Image) {
	realtimeButton.Draw(screen)
}

func updateTitleScreen() {

	if realtimeButton.IsMouseOver() {
		realtimeButton.SetBgColour(colour.JazzyPink)
	} else {
		realtimeButton.SetBgColour(colour.BrightRed)
	}

	// Handle mouse click on realtime button
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && realtimeButton.IsMouseOver() {
		state = realtimeScreen
	}
}
