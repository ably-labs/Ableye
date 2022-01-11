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
	createClientAButton    button.Button
	createClientBButton    button.Button
	createChannelButton    button.Button
	publishToChannelButton button.Button
)

func initialiseRealtimeScreen() {
	infoBar = button.NewButton(screenWidth, 35, "information bar", 22, 22, colour.Black, font.MplusSmallFont, colour.White, 0, 25)
	createClientAButton = button.NewButton(200, 35, "create client", 22, 22, colour.Black, font.MplusSmallFont, colour.Yellow, 0, screenHeight/6)
	createClientBButton = button.NewButton(200, 35, "create client", 22, 22, colour.Black, font.MplusSmallFont, colour.Yellow, screenWidth/2, screenHeight/6)
	createChannelButton = button.NewButton(200, 35, "create channel", 22, 22, colour.Black, font.MplusSmallFont, colour.Yellow, 201, screenHeight/6)
	publishToChannelButton = button.NewButton(200, 35, "publish msg", 22, 22, colour.Black, font.MplusSmallFont, colour.Yellow, 402, screenHeight/6)
}

func drawRealtimeScreen(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "Ably Realtime", 0, 0)
	infoBar.Draw(screen)
	createClientAButton.Draw(screen)
	createClientBButton.Draw(screen)
	createChannelButton.Draw(screen)
	publishToChannelButton.Draw(screen)

	// if client A has been created
	if clients[clientA] != nil {
		drawAreaUnderButton(screen, createClientAButton)
	}

	// if client B has been created
	if clients[clientB] != nil {
		drawAreaUnderButton(screen, createClientBButton)
	}
}

// drawAreaUnderButton draws a rectangle offset to touch the bottom edge of an existing button
func drawAreaUnderButton(screen *ebiten.Image, button button.Button) {
	ebitenutil.DrawRect(screen, float64(button.X), (float64(button.Y) + float64(button.Height)), (screenWidth/2)-10, screenHeight/3, colour.Green)
	ebitenutil.DrawRect(screen, float64(button.X)+1, (float64(button.Y)+float64(button.Height))+1, (screenWidth/2)-12, (screenHeight/3)-2, colour.Black)
}

func updateRealtimeScreen(ctx context.Context) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		state = titleScreen
	}

	// Handle mouse click on create client A button
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && createClientAButton.IsMouseOver() {
		if err := createRealtimeClient(clientA); err != nil {
			createClientAButton.SetBgColour(colour.Red)
		}
		infoBar.SetText(createRealtimeClientSuccess)
		createClientAButton.SetBgColour(colour.Green)

	}

	// Handle mouse click on create client B button
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && createClientBButton.IsMouseOver() {
		if err := createRealtimeClient(clientB); err != nil {
			createClientAButton.SetBgColour(colour.Red)
		}
		infoBar.SetText(createRealtimeClientSuccess)
		createClientBButton.SetBgColour(colour.Green)
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
