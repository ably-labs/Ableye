package main

import (
	"fmt"
	"github.com/ably-labs/rosie-demo/button"
	colour "github.com/ably-labs/rosie-demo/colours"
	font "github.com/ably-labs/rosie-demo/fonts"
	"github.com/ably-labs/rosie-demo/text"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type connectionElements struct {
	id                  connectionID
	createClient        button.Button
	closeClient         button.Button
	setChannel          button.Button
	channelName         text.Text
	channelStatus       text.Text
	channelSubscribeAll button.Button
	presenceInfo        text.Text
	announcePresence 	button.Button
	getPresence         button.Button
}

// The elements of the realtime screen.
var (
	infoBar button.Button

	//Connection A elements
	connectionA connectionElements

	//Connection B elements
	connectionB connectionElements
)

func initialiseRealtimeScreen() {
	infoBar = button.NewButton(screenWidth, 35, "information bar", 22, 22, colour.Black, font.MplusSmallFont, colour.White, 0, 25)

	//Initialise connection A elements.
	connectionA.id = clientA
	connectionA.createClient = button.NewButton(150, 35, createClientText, 22, 22, colour.Black, font.MplusSmallFont, colour.Yellow, 0, screenHeight/6)
	connectionA.closeClient = button.NewButton(35, 35, "X", 12, 22, colour.Black, font.MplusSmallFont, colour.Red, (screenWidth/2)-45, screenHeight/6)
	connectionA.setChannel = button.NewButton(150, 35, setChannelText, 22, 22, colour.Black, font.MplusSmallFont, colour.Yellow, 151, screenHeight/6)
	connectionA.channelName = text.NewText("", colour.Yellow, font.MplusSmallFont, 0, 0)
	connectionA.channelStatus = text.NewText("", colour.White, font.MplusSmallFont, 0, 0)
	connectionA.channelSubscribeAll = button.NewButton(125, 30, subscribeAllText, 12, 20, colour.Black, font.MplusSmallFont, colour.Yellow, 0, 0)
	connectionA.presenceInfo = text.NewText("", colour.Cyan, font.MplusSmallFont, 0, 0)
	connectionA.announcePresence = button.NewButton(100, 30, announcePresenceText, 12, 20, colour.Black, font.MplusSmallFont, colour.Cyan, 0, 0)
	connectionA.getPresence = button.NewButton(100, 30, getPresenceText, 32, 20, colour.Black, font.MplusSmallFont, colour.Cyan, 0, 0)

	//Create Connection B elements
	connectionB.id = clientB
	connectionB.createClient = button.NewButton(150, 35, createClientText, 22, 22, colour.Black, font.MplusSmallFont, colour.Yellow, screenWidth/2, screenHeight/6)
	connectionB.closeClient = button.NewButton(35, 35, "X", 12, 22, colour.Black, font.MplusSmallFont, colour.Red, (screenWidth)-45, screenHeight/6)
	connectionB.setChannel = button.NewButton(150, 35, setChannelText, 22, 22, colour.Black, font.MplusSmallFont, colour.Yellow, (screenWidth/2)+151, screenHeight/6)
	connectionB.channelName = text.NewText("", colour.Yellow, font.MplusSmallFont, 0, 0)
	connectionB.channelStatus = text.NewText("", colour.White, font.MplusSmallFont, 0, 0)
	connectionB.channelSubscribeAll = button.NewButton(125, 30, subscribeAllText, 12, 20, colour.Black, font.MplusSmallFont, colour.Yellow, 0, 0)
	connectionB.presenceInfo = text.NewText("", colour.Cyan, font.MplusSmallFont, 0, 0)
	connectionB.announcePresence = button.NewButton(100, 30, announcePresenceText, 12, 20, colour.Black, font.MplusSmallFont, colour.Cyan, 0, 0)
	connectionB.getPresence = button.NewButton(100, 30, getPresenceText, 32, 20, colour.Black, font.MplusSmallFont, colour.Cyan, 0, 0)

}

func drawRealtimeScreen(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "Ably Realtime", 0, 0)
	infoBar.Draw(screen)

	//Connection A elements
	drawConnectionElements(screen, &connectionA)

	//Connection B elements
	drawConnectionElements(screen, &connectionB)
}

// drawConnectionElements draws all the elements associated with a connection to the screen.
func drawConnectionElements(screen *ebiten.Image, elements *connectionElements) {

	id := elements.id

	// Draw the button to create a new client.
	elements.createClient.Draw(screen)

	// if client has been created
	if connections[id] != nil && connections[id].client != nil {
		drawClientInfo(screen, elements.createClient)
		// if a channel has not been set for this client, draw the set channel button.
		if connections[id].channel == nil {
			elements.setChannel.Draw(screen)
		}
		// draw the close client button.
		elements.closeClient.Draw(screen)
	}

	// if client channel has been created
	if connections[id] != nil && connections[id].channel != nil {
		drawChannelInfo(screen, elements)
	}
}

func updateRealtimeScreen() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		state = titleScreen
	}

	updateCreateClientButton(&connectionA.createClient, connectionA.id)
	updateCreateClientButton(&connectionB.createClient, connectionB.id)

	updateCloseClientButton(&connectionA.closeClient, &connectionA.createClient, &connectionA.presenceInfo, connectionA.id)
	updateCloseClientButton(&connectionB.closeClient, &connectionB.createClient, &connectionB.presenceInfo, connectionB.id)

	updateSetChannelButton(&connectionA.setChannel, connectionA.id)
	updateSetChannelButton(&connectionB.setChannel, connectionB.id)

	updateSubscribeChannelButton(&connectionA.channelSubscribeAll, connectionA.id)
	updateSubscribeChannelButton(&connectionB.channelSubscribeAll, connectionB.id)

	updateAnnouncePresenceButton(&connectionA.announcePresence, connectionA.id)
	updateAnnouncePresenceButton(&connectionB.announcePresence, connectionB.id)

	updateGetPresenceButton(&connectionA.getPresence, &connectionA.presenceInfo, connectionA.id)
	updateGetPresenceButton(&connectionB.getPresence, &connectionB.presenceInfo, connectionB.id)

}

// drawClientInfo draws a rectangle that is used to display client information.
// This rectangle is anchored to an existing button.
func drawClientInfo(screen *ebiten.Image, button button.Button) {
	ebitenutil.DrawRect(screen, float64(button.X), (float64(button.Y) + float64(button.Height)), (screenWidth/2)-10, screenHeight/3, colour.Green)
	ebitenutil.DrawRect(screen, float64(button.X)+1, (float64(button.Y)+float64(button.Height))+1, (screenWidth/2)-12, (screenHeight/3)-2, colour.Black)

}

// drawChannelInfo draws channel information, it's location is anchored to an existing button
func drawChannelInfo(screen *ebiten.Image, elements *connectionElements) {
	// button is used to anchor the channel information, everything is drawn relative to the button.
	button := elements.createClient
	id := elements.id

	// channel area
	ebitenutil.DrawRect(screen, float64(button.X)+4, (float64(button.Y)+float64(button.Height))+3, (screenWidth/2)-18, screenHeight/10, colour.Yellow)
	ebitenutil.DrawRect(screen, float64(button.X)+5, (float64(button.Y)+float64(button.Height))+4, (screenWidth/2)-20, (screenHeight/10)-2, colour.Black)

	// channel name text box
	elements.channelName.SetX(button.X + 10)
	elements.channelName.SetY((button.Y + button.Height) + 25)
	elements.channelName.SetText(fmt.Sprintf("Channel : %s", connections[id].channel.Name))
	elements.channelName.Draw(screen)

	// channel status text box
	elements.channelStatus.SetX(button.X + 200)
	elements.channelStatus.SetY((button.Y + button.Height) + 25)
	elements.channelStatus.SetText(fmt.Sprintf("Status : %s", connections[id].channel.State()))
	elements.channelStatus.Draw(screen)

	// channel subscribe all button
	elements.channelSubscribeAll.SetX(button.X + 543)
	elements.channelSubscribeAll.SetY((button.Y + button.Height) + 4)
	elements.channelSubscribeAll.Draw(screen)

	// presence area
	ebitenutil.DrawRect(screen, float64(button.X)+8, (float64(button.Y)+float64(button.Height))+42, (screenWidth/2)-26, screenHeight/24, colour.Cyan)
	ebitenutil.DrawRect(screen, float64(button.X)+9, (float64(button.Y)+float64(button.Height))+43, (screenWidth/2)-28, (screenHeight/24)-2, colour.Black)

	// presence info text box
	// if presenceInfo is being drawn in its initisalised location.
	if elements.presenceInfo.X == 0 && elements.presenceInfo.Y == 0 {
		// initalise the text
		elements.presenceInfo.SetText("Presence :")
	}

	elements.presenceInfo.SetX(button.X + 12)
	elements.presenceInfo.SetY(button.Y + button.Height + 62)
	elements.presenceInfo.Draw(screen)

	// announce presence button
	elements.announcePresence.SetX(button.X + 464)
	elements.announcePresence.SetY(button.Y + button.Height + 43)
	elements.announcePresence.Draw(screen)

	// get presence button
	elements.getPresence.SetX(button.X + 565)
	elements.getPresence.SetY(button.Y + button.Height + 43)
	elements.getPresence.Draw(screen)
}

// updateCreateClientButton contains the update logic for each client creation button.
func updateCreateClientButton(button *button.Button, id connectionID) {

	// Handle mouseover interaction with create button while a connection does not exist.
	if button.IsMouseOver() && connections[id] == nil {
		button.SetBgColour(colour.Green)
	}

	// if the button is not moused over and there is no connection.
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
func updateCloseClientButton(closeButton *button.Button, createButton *button.Button, presenceInfo *text.Text, id connectionID) {
	// Handle mouseover interaction with a close client button.
	if closeButton.IsMouseOver() {
		closeButton.SetTextColour(colour.White)
	} else {
		closeButton.SetTextColour(colour.Black)
	}

	// Handle mouse click on a close client button.
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && closeButton.IsMouseOver() {
		closeRealtimeClient(id)
		infoBar.SetText(closeRealtimeClientSuccess)
		// Reset the create button once a client is closed.
		createButton.SetBgColour(colour.Yellow)
		createButton.SetText(createClientText)

		// Reset the presence text once a client is closed.
		// Return it to its initialised location.
		// Set its text to its initialised text.
		presenceInfo.SetX(0)
		presenceInfo.SetY(0)
		presenceInfo.SetText("")

	}
}

// updateSetChannelButton contains the update logic for each set channel button.
func updateSetChannelButton(button *button.Button, id connectionID) {
	// Handle mouseover interaction with a set channel button.
	if button.IsMouseOver() {
		button.SetBgColour(colour.Green)
	} else {
		button.SetBgColour(colour.Yellow)
	}

	// Handle mouse click on set channel button.
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && button.IsMouseOver() {
		// if the connection exists and does not have a channel.
		if connections[id] != nil && connections[id].channel == nil {
			setChannel(id)
			infoBar.SetText(setChannelSuccess)
		}
	}
}

func updateSubscribeChannelButton(button *button.Button, id connectionID) {

	// If a connection exists and no unsubscribeAll function is saved
	if connections[id] != nil && connections[id].unsubscribeAll == nil {
		button.SetText(subscribeAllText)
	}

	// Handle mouseover interaction with subscribe all button while the channel is not subscribed.
	if button.IsMouseOver() && connections[id] != nil && connections[id].unsubscribeAll == nil {
		button.SetBgColour(colour.White)
	}

	// if the button is not moused over and the channel is not subscribed.
	if !button.IsMouseOver() && connections[id] != nil && connections[id].unsubscribeAll == nil {
		button.SetBgColour(colour.Yellow)
	}

	// Handle mouse click on subscribe all button.
	// As this button toggles between two states, the trigger is the mouse button releasing.
	// This prevents the button from quickly toggling if the mouse button is held down.
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && button.IsMouseOver() {
		// if a channel exists and the connection has no unsubscribeAll function saved

		if connections[id].channel != nil && connections[id].unsubscribeAll == nil {

			unsubscribeAll, err := subscribeAll(id)
			if err != nil {
				infoBar.SetText(err.Error())
			}
			// Save the unsubscribe all function.
			connections[id].unsubscribeAll = &unsubscribeAll
			infoBar.SetText(subscribeAllSuccess)

			// Change the SubscribeAll button into an UnsubscribeAll button.
			button.SetBgColour(colour.Red)
			button.SetText(unsubscribeAllText)
			return
		}

		// if there is an unsubscribeAll function saved
		if connections[id].unsubscribeAll != nil {
			unsubscribeAll(id)
			infoBar.SetText(unsubscribeAllSuccess)
			connections[id].unsubscribeAll = nil
			return
		}
	}
}

// updateGetPresenceButton contains the update logic for each get presence button
func updateGetPresenceButton(button *button.Button, text *text.Text, id connectionID) {
	// Handle mouseover interaction with a get presence button.
	if button.IsMouseOver() {
		button.SetBgColour(colour.White)
	} else {
		button.SetBgColour(colour.Cyan)
	}

	// Handle mouse click on get presence button.
	// As this button has no conditional to prevent its action triggering on every frame,
	// the action to get presence is performed on mouse release.
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && button.IsMouseOver() {
		// the call to get presence is async to prevent blocking.
		go getPresence(id, text)
	}
}

// updateGetPresenceButton contains the update logic for each get presence button.
func updateAnnouncePresenceButton(button *button.Button, id connectionID) {
	// Handle mouseover interaction with an announce presence button.
	if button.IsMouseOver() {
		button.SetBgColour(colour.White)
	} else {
		button.SetBgColour(colour.Cyan)
	}

	// Handle mouse click on get presence button.
	// As this button has no conditional to prevent its action triggering on every frame,
	// the action to get presence is performed on mouse release.
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && button.IsMouseOver() {
	
		err := announcePresence(id)
		if err != nil {
			infoBar.SetText(err.Error())
		}
	}
}
