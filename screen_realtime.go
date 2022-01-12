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
	infoBar button.Button

	//Client A
	createClientAButton button.Button
	closeClientAButton  button.Button
	setChannelAButton   button.Button

	//Client B
	createClientBButton button.Button
	closeClientBButton  button.Button

	// createChannelButton    button.Button
	// publishToChannelButton button.Button
)

func initialiseRealtimeScreen() {
	infoBar = button.NewButton(screenWidth, 35, "information bar", 22, 22, colour.Black, font.MplusSmallFont, colour.White, 0, 25)

	//Client A
	createClientAButton = button.NewButton(200, 35, createClientText, 22, 22, colour.Black, font.MplusSmallFont, colour.Yellow, 0, screenHeight/6)
	closeClientAButton = button.NewButton(35, 35, "X", 12, 22, colour.Black, font.MplusSmallFont, colour.Red, (screenWidth/2)-45, screenHeight/6)
	setChannelAButton = button.NewButton(200, 35, setChannelText, 22, 22, colour.Black, font.MplusSmallFont, colour.Yellow, 201, screenHeight/6)

	//Client B
	createClientBButton = button.NewButton(200, 35, createClientText, 22, 22, colour.Black, font.MplusSmallFont, colour.Yellow, screenWidth/2, screenHeight/6)
	closeClientBButton = button.NewButton(35, 35, "X", 12, 22, colour.Black, font.MplusSmallFont, colour.Red, (screenWidth)-45, screenHeight/6)
	// createChannelButton = button.NewButton(200, 35, "create channel", 22, 22, colour.Black, font.MplusSmallFont, colour.Yellow, 201, screenHeight/6)
	// publishToChannelButton = button.NewButton(200, 35, "publish msg", 22, 22, colour.Black, font.MplusSmallFont, colour.Yellow, 402, screenHeight/6)
}

func drawRealtimeScreen(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "Ably Realtime", 0, 0)
	infoBar.Draw(screen)

	//Client A
	createClientAButton.Draw(screen)

	// if client A has been created
	if connections[clientA] != nil && connections[clientA].client != nil {
		drawAreaUnderButton(screen, createClientAButton)

		setChannelAButton.Draw(screen)
		closeClientAButton.Draw(screen)
	}

	//Client B
	createClientBButton.Draw(screen)

	// if client B has been created
	if connections[clientB] != nil && connections[clientB].client != nil {
		drawAreaUnderButton(screen, createClientBButton)
		closeClientBButton.Draw(screen)
	}

	// createChannelButton.Draw(screen)
	// publishToChannelButton.Draw(screen)

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

	// Handle mouseover interaction with create client A button while client A does not exist.
	if createClientAButton.IsMouseOver() && connections[clientA] == nil {
		createClientAButton.SetBgColour(colour.Green)
	} 
	
	if !createClientAButton.IsMouseOver() && connections[clientA] == nil {
		createClientAButton.SetBgColour(colour.Yellow)
	}

	// Handle mouse click on create client A button
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && createClientAButton.IsMouseOver() {
		if connections[clientA] == nil {
			if err := createRealtimeClient(clientA); err != nil {
				infoBar.SetText(err.Error())
			}
			infoBar.SetText(createRealtimeClientSuccess)
			createClientAButton.SetBgColour(colour.Green)
			createClientAButton.SetText("Client A")
		}
	}


	// Handle mouse click on set channel A button
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && setChannelAButton.IsMouseOver() {
		setChannel()
	
	}

	// Handle mouseover interaction with close client A button
	if closeClientAButton.IsMouseOver() {
		closeClientAButton.SetTextColour(colour.White)
	} else {
		closeClientAButton.SetTextColour(colour.Black)
	}

	// Handle mouse click on close client A button
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && closeClientAButton.IsMouseOver() {
		closeRealtimeClient(clientA)
		infoBar.SetText(closeRealtimeClientSuccess)
		createClientAButton.SetBgColour(colour.Yellow)
		createClientAButton.SetText(createClientText)
	}

	// Handle mouseover interaction with create client B button while client B does not exist.
	if createClientBButton.IsMouseOver() && connections[clientB]  == nil {
		createClientBButton.SetBgColour(colour.Green)
	} 
	
	if !createClientBButton.IsMouseOver() && connections[clientB] == nil{
		createClientBButton.SetBgColour(colour.Yellow)
	}

	// Handle mouse click on create client B button
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && createClientBButton.IsMouseOver() {
		if connections[clientB]  == nil {
			if err := createRealtimeClient(clientB); err != nil {
				infoBar.SetText(err.Error())
			}
			infoBar.SetText(createRealtimeClientSuccess)
			createClientBButton.SetBgColour(colour.Green)
			createClientBButton.SetText("Client B")
		}
	}

	// Handle mouseover interaction with close client B button
	if closeClientBButton.IsMouseOver() {
		closeClientBButton.SetTextColour(colour.White)
	} else {
		closeClientBButton.SetTextColour(colour.Black)
	}

	// Handle mouse click on close client B button
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && closeClientBButton.IsMouseOver() {
		closeRealtimeClient(clientB)
		infoBar.SetText(closeRealtimeClientSuccess)
		createClientBButton.SetBgColour(colour.Yellow)
		createClientBButton.SetText(createClientText)
	}

	// // Handle mouse click on create channel button
	// if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && createChannelButton.IsMouseOver() {
	// 	getChannel()
	// 	createChannelButton.SetBgColour(colour.Green)
	// }

	// // Handle mouse click on publish to channel button
	// if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && publishToChannelButton.IsMouseOver() {
	// 	if channel == nil {
	// 		publishToChannelButton.SetBgColour(colour.Magenta)
	// 	} else {
	// 		err := publishToChannel(ctx)
	// 		if err != nil {
	// 			publishToChannelButton.SetBgColour(colour.Red)
	// 		}
	// 		publishToChannelButton.SetBgColour(colour.Green)
	// 	}
	// }

}
