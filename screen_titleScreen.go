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
	restButton     button.Button
)

func initialiseTitleScreen() {
	realtimeButton = button.NewButton(200, 100, "Ably Realtime", 25, 50, colour.White, font.MplusNormalFont, colour.Red, screenWidth/4, screenHeight/2)
	restButton = button.NewButton(200, 100, "Ably Rest", 35, 50, colour.White, font.MplusNormalFont, colour.Red, (screenWidth/4)+(screenWidth/3), screenHeight/2)
}

func drawTitleScreen(screen *ebiten.Image) {
	realtimeButton.Draw(screen)
	restButton.Draw(screen)
}

func updateTitleScreen() {

	if realtimeButton.IsMouseOver() {
		realtimeButton.SetBgColour(colour.Magenta)
	} else {
		realtimeButton.SetBgColour(colour.Red)
	}

	if restButton.IsMouseOver() {
		restButton.SetBgColour(colour.Magenta)
	} else {
		restButton.SetBgColour(colour.Red)
	}

	// Handle mouse click on realtime button
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && realtimeButton.IsMouseOver() {
		state = realtimeScreen
	}

	// Handle mouse click on rest button
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && restButton.IsMouseOver() {
		state = restScreen
	}
}
