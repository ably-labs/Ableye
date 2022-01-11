package main

import (
	"context"
	"github.com/ably-labs/rosie-demo/button"
	colour "github.com/ably-labs/rosie-demo/colours"
	font "github.com/ably-labs/rosie-demo/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// The elements of the realtime screen.
var (
	infoBar                button.Button
	createClientButton     button.Button
	createChannelButton    button.Button
	publishToChannelButton button.Button
)

func initialiseRealtimeScreen() {
	infoBar = button.NewButton(screenWidth, 35, "information bar", 22, 22, colour.Black, font.MplusSmallFont, colour.White, 0, 25)
	createClientButton = button.NewButton(200, 100, "create client", 25, 50, colour.Black, font.MplusNormalFont, colour.Yellow, 0, screenHeight/6)
	createChannelButton = button.NewButton(200, 100, "create channel", 25, 50, colour.Black, font.MplusNormalFont, colour.Yellow, 201, screenHeight/6)
	publishToChannelButton = button.NewButton(200, 100, "publish msg", 25, 50, colour.Black, font.MplusNormalFont, colour.Yellow, 402, screenHeight/6)
}

func drawRealtimeScreen(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "Ably Realtime", 0, 0)
	infoBar.Draw(screen)
	createClientButton.Draw(screen)
	createChannelButton.Draw(screen)
	publishToChannelButton.Draw(screen)
}

func updateRealtimeScreen(ctx context.Context) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		state = titleScreen
	}

	// Handle mouse click on create client button
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && createClientButton.IsMouseOver() {
		if err := createRealtimeClient(); err != nil {
			createClientButton.SetBgColour(colour.Red)
		}
		infoBar.SetText(createRealtimeClientSuccess)
		createClientButton.SetBgColour(colour.Green)
	}

	// Handle mouse click on create channel button
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && createChannelButton.IsMouseOver() {
		getChannel()
		createChannelButton.SetBgColour(colour.Green)
	}

	// Handle mouse click on publish to channel button
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && publishToChannelButton.IsMouseOver() {
		if channel == nil {
			publishToChannelButton.SetBgColour(colour.Magenta)
		} else {
			err := publishToChannel(ctx)
			if err != nil {
				publishToChannelButton.SetBgColour(colour.Red)
			}
			publishToChannelButton.SetBgColour(colour.Green)
		}
	}

}
