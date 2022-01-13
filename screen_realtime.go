package main

import (
	"context"
	"fmt"
	"github.com/ably-labs/rosie-demo/button"
	colour "github.com/ably-labs/rosie-demo/colours"
	font "github.com/ably-labs/rosie-demo/fonts"
	"github.com/ably-labs/rosie-demo/text"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// The elements of the realtime screen.
var (
	infoBar button.Button

	//Client A
	createClientAButton button.Button
	closeClientAButton  button.Button
	setChannelAButton   button.Button
	displayChannelA     text.Text

	//Client B
	createClientBButton button.Button
	closeClientBButton  button.Button
	setChannelBButton   button.Button
	displayChannelB     text.Text
)

func initialiseRealtimeScreen() {
	infoBar = button.NewButton(screenWidth, 35, "information bar", 22, 22, colour.Black, font.MplusSmallFont, colour.White, 0, 25)

	//Client A
	createClientAButton = button.NewButton(200, 35, createClientText, 22, 22, colour.Black, font.MplusSmallFont, colour.Yellow, 0, screenHeight/6)
	closeClientAButton = button.NewButton(35, 35, "X", 12, 22, colour.Black, font.MplusSmallFont, colour.Red, (screenWidth/2)-45, screenHeight/6)
	setChannelAButton = button.NewButton(200, 35, setChannelText, 22, 22, colour.Black, font.MplusSmallFont, colour.Yellow, 201, screenHeight/6)
	displayChannelA = text.NewText("", colour.Yellow, font.MplusSmallFont, 0, 0)

	//Client B
	createClientBButton = button.NewButton(200, 35, createClientText, 22, 22, colour.Black, font.MplusSmallFont, colour.Yellow, screenWidth/2, screenHeight/6)
	closeClientBButton = button.NewButton(35, 35, "X", 12, 22, colour.Black, font.MplusSmallFont, colour.Red, (screenWidth)-45, screenHeight/6)
	setChannelBButton = button.NewButton(200, 35, setChannelText, 22, 22, colour.Black, font.MplusSmallFont, colour.Yellow, (screenWidth/2)+201, screenHeight/6)
	displayChannelB = text.NewText("", colour.Yellow, font.MplusSmallFont, 0, 0)
}

func drawRealtimeScreen(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "Ably Realtime", 0, 0)
	infoBar.Draw(screen)

	//Client A
	createClientAButton.Draw(screen)

	// if client A has been created
	if connections[clientA] != nil && connections[clientA].client != nil {
		drawClientAreaUnderButton(screen, createClientAButton)
		// if a channel has not been set for this client, draw the set channel button.
		if connections[clientA].channel == nil {
			setChannelAButton.Draw(screen)
		}
		closeClientAButton.Draw(screen)
	}

	// if client A channel has been created
	if connections[clientA] != nil && connections[clientA].channel != nil {
		drawChannelInfo(screen, createClientAButton, displayChannelA, clientA)
	}

	//Client B
	createClientBButton.Draw(screen)

	// if client B has been created
	if connections[clientB] != nil && connections[clientB].client != nil {
		drawClientAreaUnderButton(screen, createClientBButton)
		// if a channel has not been set for this client, draw the set channel button.
		if connections[clientB].channel == nil {
			setChannelBButton.Draw(screen)
		}
		closeClientBButton.Draw(screen)
	}

	// if client B channel has been created
	if connections[clientB] != nil && connections[clientB].channel != nil {
		drawChannelInfo(screen, createClientBButton, displayChannelB, clientB)
	}
}

func updateRealtimeScreen(ctx context.Context) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		state = titleScreen
	}

	updateCreateClientButton(&createClientAButton, clientA)
	updateCreateClientButton(&createClientBButton, clientB)

	updateCloseClientButton(&closeClientAButton, &createClientAButton, clientA)
	updateCloseClientButton(&closeClientBButton, &createClientBButton, clientB)

	// Handle mouse click on set channel A button
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && setChannelAButton.IsMouseOver() {
		if connections[clientA] != nil && connections[clientA].channel == nil {
			setChannel(clientA)
			infoBar.SetText(setChannelSuccess)
		}
	}

	// Handle mouse click on set channel B button
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && setChannelBButton.IsMouseOver() {
		if connections[clientB] != nil && connections[clientB].channel == nil {
			setChannel(clientB)
			infoBar.SetText(setChannelSuccess)
		}
	}
}

// drawClientAreaUnderButton draws a rectangle offset to touch the bottom edge of an existing button
func drawClientAreaUnderButton(screen *ebiten.Image, button button.Button) {
	ebitenutil.DrawRect(screen, float64(button.X), (float64(button.Y) + float64(button.Height)), (screenWidth/2)-10, screenHeight/3, colour.Green)
	ebitenutil.DrawRect(screen, float64(button.X)+1, (float64(button.Y)+float64(button.Height))+1, (screenWidth/2)-12, (screenHeight/3)-2, colour.Black)

}

// drawChannelInfo draws channel information, it's location is anchored to an existing button
func drawChannelInfo(screen *ebiten.Image, button button.Button, text text.Text, id connectionID) {
	ebitenutil.DrawRect(screen, float64(button.X)+4, (float64(button.Y)+float64(button.Height))+3, (screenWidth/2)-18, screenHeight/24, colour.Yellow)
	ebitenutil.DrawRect(screen, float64(button.X)+5, (float64(button.Y)+float64(button.Height))+4, (screenWidth/2)-20, (screenHeight/24)-2, colour.Black)
	text.SetX(button.X + 5)
	text.SetY((button.Y + button.Height) + 25)
	text.SetText(fmt.Sprintf("Channel : %s", connections[id].channel.Name))
	text.Draw(screen)
}

// updateCreateClientButton contains the update logic for each client creation button.
func updateCreateClientButton(button *button.Button, id connectionID) {

	// Handle mouseover interaction with create button while a connection does not exist.
	if button.IsMouseOver() && connections[id] == nil {
		button.SetBgColour(colour.Green)
	}

	if !button.IsMouseOver() && connections[id] == nil {
		button.SetBgColour(colour.Yellow)
	}

	// Handle mouse click on a create client button
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && button.IsMouseOver() {
		if connections[id] == nil {
			if err := createRealtimeClient(id); err != nil {
				infoBar.SetText(err.Error())
			}
			infoBar.SetText(createRealtimeClientSuccess)
			button.SetBgColour(colour.Green)
			button.SetText(id.string())
		}
	}
}

// updateCloseClientButton contains the update logic for each close client button.
// a createButton is passed to this function so it can be reset after a client is closed.
func updateCloseClientButton(closeButton *button.Button, createButton *button.Button, id connectionID) {

	// Handle mouseover interaction with a close client button
	if closeButton.IsMouseOver() {
		closeButton.SetTextColour(colour.White)
	} else {
		closeButton.SetTextColour(colour.Black)
	}

	// Handle mouse click on a close client button
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && closeButton.IsMouseOver() {
		closeRealtimeClient(id)
		infoBar.SetText(closeRealtimeClientSuccess)
		// reset the create button once a client is closed.
		createButton.SetBgColour(colour.Yellow)
		createButton.SetText(createClientText)
	}
}
