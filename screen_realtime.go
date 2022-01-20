package main

import (
	"fmt"

	"github.com/ably-labs/rosie-demo/button"
	colour "github.com/ably-labs/rosie-demo/colours"
	font "github.com/ably-labs/rosie-demo/fonts"
	"github.com/ably-labs/rosie-demo/text"
	"github.com/ably-labs/rosie-demo/textbox"
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
	announcePresence    button.Button
	getPresence         button.Button
	leavePresence       button.Button
	eventInfo           text.Text
	channelNameLabel    text.Text
	channelNameInput    textbox.TextBox
	messageNameLabel    text.Text
	messageNameInput    textbox.TextBox
	messageDataLabel    text.Text
	messageDataInput    textbox.TextBox
	channelPublish      button.Button
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
	connectionA.channelName = text.NewText("", colour.Yellow, font.MplusSmallFont, 0, 0)
	connectionA.channelStatus = text.NewText("", colour.Yellow, font.MplusSmallFont, 0, 0)
	connectionA.channelSubscribeAll = button.NewButton(125, 30, subscribeAllText, 12, 20, colour.Black, font.MplusSmallFont, colour.Yellow, 0, 0)
	connectionA.presenceInfo = text.NewText("", colour.Cyan, font.MplusSmallFont, 0, 0)
	connectionA.announcePresence = button.NewButton(100, 30, announcePresenceText, 12, 20, colour.Black, font.MplusSmallFont, colour.Cyan, 0, 0)
	connectionA.getPresence = button.NewButton(50, 30, getPresenceText, 12, 20, colour.Black, font.MplusSmallFont, colour.Cyan, 0, 0)
	connectionA.leavePresence = button.NewButton(70, 30, leavePresenceText, 12, 20, colour.Black, font.MplusSmallFont, colour.Cyan, 0, 0)
	connectionA.eventInfo = text.NewText("", colour.White, font.MplusSmallFont, 0, 0)
	connectionA.channelNameLabel = text.NewText(fmt.Sprintf("%s :", channelNameText), colour.Green, font.MplusSmallFont, 0, 0)
	connectionA.channelNameInput = textbox.NewTextBox(200, 36, 4, defaultChannelName, 9, 12, 22, colour.Green, font.MplusSmallFont, colour.Black, colour.Green, 0, 0)
	connectionA.setChannel = button.NewButton(150, 35, setChannelText, 22, 22, colour.Black, font.MplusSmallFont, colour.Green, 0, 0)
	connectionA.messageNameLabel = text.NewText(fmt.Sprintf("%s :", messageNameText), colour.Magenta, font.MplusSmallFont, 0, 0)
	connectionA.messageNameInput = textbox.NewTextBox(200, 36, 4, defaultMessageName, 12, 12, 22, colour.Magenta, font.MplusSmallFont, colour.Black, colour.Magenta, 0, 0)
	connectionA.messageDataLabel = text.NewText(fmt.Sprintf("%s :", messageDataText), colour.Magenta, font.MplusSmallFont, 0, 0)
	connectionA.messageDataInput = textbox.NewTextBox(200, 36, 4, defaultMessageData, 12, 12, 22, colour.Magenta, font.MplusSmallFont, colour.Black, colour.Magenta, 0, 0)
	connectionA.channelPublish = button.NewButton(80, 30, publishText, 12, 20, colour.Black, font.MplusSmallFont, colour.Magenta, 0, 0)

	//Create Connection B elements
	connectionB.id = clientB
	connectionB.createClient = button.NewButton(150, 35, createClientText, 22, 22, colour.Black, font.MplusSmallFont, colour.Yellow, screenWidth/2, screenHeight/6)
	connectionB.closeClient = button.NewButton(35, 35, "X", 12, 22, colour.Black, font.MplusSmallFont, colour.Red, (screenWidth)-45, screenHeight/6)
	connectionB.channelName = text.NewText("", colour.Yellow, font.MplusSmallFont, 0, 0)
	connectionB.channelStatus = text.NewText("", colour.Yellow, font.MplusSmallFont, 0, 0)
	connectionB.channelPublish = button.NewButton(80, 30, publishText, 12, 20, colour.Black, font.MplusSmallFont, colour.Yellow, 0, 0)
	connectionB.channelSubscribeAll = button.NewButton(125, 30, subscribeAllText, 12, 20, colour.Black, font.MplusSmallFont, colour.Yellow, 0, 0)
	connectionB.presenceInfo = text.NewText("", colour.Cyan, font.MplusSmallFont, 0, 0)
	connectionB.announcePresence = button.NewButton(100, 30, announcePresenceText, 12, 20, colour.Black, font.MplusSmallFont, colour.Cyan, 0, 0)
	connectionB.getPresence = button.NewButton(50, 30, getPresenceText, 12, 20, colour.Black, font.MplusSmallFont, colour.Cyan, 0, 0)
	connectionB.leavePresence = button.NewButton(70, 30, leavePresenceText, 12, 20, colour.Black, font.MplusSmallFont, colour.Cyan, 0, 0)
	connectionB.eventInfo = text.NewText("", colour.White, font.MplusSmallFont, 0, 0)
	connectionB.channelNameLabel = text.NewText(fmt.Sprintf("%s :", channelNameText), colour.Green, font.MplusSmallFont, 0, 0)
	connectionB.channelNameInput = textbox.NewTextBox(200, 36, 4, defaultChannelName, 9, 12, 22, colour.Green, font.MplusSmallFont, colour.Black, colour.Green, 0, 0)
	connectionB.setChannel = button.NewButton(150, 35, setChannelText, 22, 22, colour.Black, font.MplusSmallFont, colour.Green, 0, 0)
	connectionB.messageNameLabel = text.NewText(fmt.Sprintf("%s :", messageNameText), colour.Magenta, font.MplusSmallFont, 0, 0)
	connectionB.messageNameInput = textbox.NewTextBox(200, 36, 4, defaultMessageName, 18, 12, 22, colour.Magenta, font.MplusSmallFont, colour.Black, colour.Magenta, 0, 0)
	connectionB.messageDataLabel = text.NewText(fmt.Sprintf("%s :", messageDataText), colour.Magenta, font.MplusSmallFont, 0, 0)
	connectionB.messageDataInput = textbox.NewTextBox(200, 36, 4, defaultMessageData, 18, 12, 22, colour.Magenta, font.MplusSmallFont, colour.Black, colour.Magenta, 0, 0)
	connectionB.channelPublish = button.NewButton(80, 30, publishText, 12, 20, colour.Black, font.MplusSmallFont, colour.Magenta, 0, 0)

}

func drawRealtimeScreen(screen *ebiten.Image) {

	// Info bar is used to display log messages and error messages.
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

		// if a channel has not been set for this client, draw the elements required to set the channel.
		if connections[id].channel == nil {
			drawSetChannel(screen, elements.createClient, elements.channelNameLabel, &elements.channelNameInput, &elements.setChannel)
		}
		// draw the close client button.
		elements.closeClient.Draw(screen)
	}

	// if client channel has been created
	if connections[id] != nil && connections[id].channel != nil {
		drawChannelInfo(screen, elements)
	}

	// if a channel has been subscribed an unsubscribe function will be saved in memory.
	// if an unsubscribe function exists, draw event info
	if connections[id] != nil && connections[id].unsubscribe != nil {
		drawEventInfo(screen, elements.createClient, &elements.eventInfo)
	}

}

func drawSetChannel(screen *ebiten.Image, createClient button.Button, channelNameLabel text.Text, channelNameInput *textbox.TextBox, setChannel *button.Button) {

	// elements are anchored to the createClient button.
	button := createClient

	// channel name label
	channelNameLabel.SetX(button.X + 10)
	channelNameLabel.SetY(button.Y + button.Height + 40)
	channelNameLabel.Draw(screen)

	// channel name input text box
	channelNameInput.SetX(button.X + 150)
	channelNameInput.SetY(button.Y + button.Height + 18)
	channelNameInput.Draw(screen)

	// set channel button
	setChannel.SetX(button.X + 375)
	setChannel.SetY(button.Y + button.Height + 18)
	setChannel.Draw(screen)
}

func updateRealtimeScreen() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		state = titleScreen
	}

	updateCreateClientButton(&connectionA.createClient, connectionA.id)
	updateCreateClientButton(&connectionB.createClient, connectionB.id)

	updateCloseClientButton(&connectionA.closeClient, &connectionA.createClient, &connectionA.presenceInfo, &connectionA.eventInfo, connectionA.id)
	updateCloseClientButton(&connectionB.closeClient, &connectionB.createClient, &connectionB.presenceInfo, &connectionB.eventInfo, connectionB.id)

	updateSetChannelButton(&connectionA.setChannel, connectionA.channelNameInput.GetText(), connectionA.id)
	updateSetChannelButton(&connectionB.setChannel, connectionB.channelNameInput.GetText(), connectionB.id)

	updateChannelPublishButton(&connectionA.channelPublish, connectionA.messageNameInput.GetText(), connectionA.messageDataInput.GetText(), connectionA.id)
	updateChannelPublishButton(&connectionB.channelPublish, connectionB.messageNameInput.GetText(), connectionB.messageDataInput.GetText(), connectionB.id)

	updateSubscribeChannelButton(&connectionA.channelSubscribeAll, &connectionA.eventInfo, connectionA.id)
	updateSubscribeChannelButton(&connectionB.channelSubscribeAll, &connectionB.eventInfo, connectionB.id)

	updateAnnouncePresenceButton(&connectionA.announcePresence, connectionA.id)
	updateAnnouncePresenceButton(&connectionB.announcePresence, connectionB.id)

	updateGetPresenceButton(&connectionA.getPresence, &connectionA.presenceInfo, connectionA.id)
	updateGetPresenceButton(&connectionB.getPresence, &connectionB.presenceInfo, connectionB.id)

	updateLeavePresenceButton(&connectionA.leavePresence, connectionA.id)
	updateLeavePresenceButton(&connectionB.leavePresence, connectionB.id)

	updateChannelNameInput(&connectionA.channelNameInput, connectionA.id)
	updateChannelNameInput(&connectionB.channelNameInput, connectionB.id)

	updateMessageNameInput(&connectionA.messageNameInput, connectionA.id)
	updateMessageNameInput(&connectionB.messageNameInput, connectionB.id)

	updateMessageDataInput(&connectionA.messageDataInput, connectionA.id)
	updateMessageDataInput(&connectionB.messageDataInput, connectionB.id)
}

// drawClientInfo draws a rectangle that is used to display client information.
// This rectangle is anchored to an existing button.
func drawClientInfo(screen *ebiten.Image, button button.Button) {
	ebitenutil.DrawRect(screen, float64(button.X), float64(button.Y)+float64(button.Height), (screenWidth/2)-10, screenHeight/3, colour.Green)
	ebitenutil.DrawRect(screen, float64(button.X)+1, float64(button.Y)+float64(button.Height)+1, (screenWidth/2)-12, (screenHeight/3)-2, colour.Black)
}

// drawChannelInfo draws channel information, it's location is anchored to an existing button
func drawChannelInfo(screen *ebiten.Image, elements *connectionElements) {
	// button is used to anchor the channel information, everything is drawn relative to the button.
	button := elements.createClient
	id := elements.id

	// channel area
	ebitenutil.DrawRect(screen, float64(button.X)+4, float64(button.Y)+float64(button.Height)+3, (screenWidth/2)-18, screenHeight/10, colour.Yellow)
	ebitenutil.DrawRect(screen, float64(button.X)+5, float64(button.Y)+float64(button.Height)+4, (screenWidth/2)-20, (screenHeight/10)-2, colour.Black)

	// channel name text box
	elements.channelName.SetX(button.X + 10)
	elements.channelName.SetY(button.Y + button.Height + 25)
	elements.channelName.SetText(fmt.Sprintf("%s : %s", channelNameText, connections[id].channel.Name))
	elements.channelName.Draw(screen)

	// channel status text box
	elements.channelStatus.SetX(button.X + 280)
	elements.channelStatus.SetY(button.Y + button.Height + 25)
	elements.channelStatus.SetText(fmt.Sprintf("Status : %s", connections[id].channel.State()))
	elements.channelStatus.Draw(screen)

	// channel subscribe all button
	elements.channelSubscribeAll.SetX(button.X + 543)
	elements.channelSubscribeAll.SetY(button.Y + button.Height + 4)
	elements.channelSubscribeAll.Draw(screen)

	// presence area
	ebitenutil.DrawRect(screen, float64(button.X)+8, float64(button.Y)+float64(button.Height)+42, (screenWidth/2)-26, screenHeight/24, colour.Cyan)
	ebitenutil.DrawRect(screen, float64(button.X)+9, float64(button.Y)+float64(button.Height)+43, (screenWidth/2)-28, (screenHeight/24)-2, colour.Black)

	// if presenceInfo is being drawn in its initisalised location.
	if elements.presenceInfo.X == 0 && elements.presenceInfo.Y == 0 {
		// initalise the text
		elements.presenceInfo.SetText("Presence :")
	}
	elements.presenceInfo.SetX(button.X + 12)
	elements.presenceInfo.SetY(button.Y + button.Height + 62)
	elements.presenceInfo.Draw(screen)

	// announce presence button
	elements.announcePresence.SetX(button.X + 442)
	elements.announcePresence.SetY(button.Y + button.Height + 43)
	elements.announcePresence.Draw(screen)

	// get presence button
	elements.getPresence.SetX(button.X + 543)
	elements.getPresence.SetY(button.Y + button.Height + 43)
	elements.getPresence.Draw(screen)

	// leave presence button
	elements.leavePresence.SetX(button.X + 594)
	elements.leavePresence.SetY(button.Y + button.Height + 43)
	elements.leavePresence.Draw(screen)

	// message name label
	elements.messageNameLabel.SetX(button.X + 10)
	elements.messageNameLabel.SetY(button.Y + button.Height + 105)
	elements.messageNameLabel.Draw(screen)

	// message name input text box
	elements.messageNameInput.SetX(button.X + 10)
	elements.messageNameInput.SetY(button.Y + button.Height + 120)
	elements.messageNameInput.Draw(screen)

	// message data label
	elements.messageDataLabel.SetX(button.X + 300)
	elements.messageDataLabel.SetY(button.Y + button.Height + 105)
	elements.messageDataLabel.Draw(screen)

	// message data input text box
	elements.messageDataInput.SetX(button.X + 300)
	elements.messageDataInput.SetY(button.Y + button.Height + 120)
	elements.messageDataInput.Draw(screen)

	// channel publish button
	elements.channelPublish.SetX(button.X + 550)
	elements.channelPublish.SetY(button.Y + button.Height + 125)
	elements.channelPublish.Draw(screen)

}

// drawEventInfo draws event information, it's location is anchored to an existing button
func drawEventInfo(screen *ebiten.Image, button button.Button, eventInfo *text.Text) {
	// event area
	ebitenutil.DrawRect(screen, float64(button.X)+4, float64(button.Y)+float64(button.Height)+182, (screenWidth/2)-18, screenHeight/12, colour.White)
	ebitenutil.DrawRect(screen, float64(button.X)+5, float64(button.Y)+float64(button.Height)+183, (screenWidth/2)-20, (screenHeight/12)-2, colour.Black)

	// if event info is being drawn in its initisalised location.
	if eventInfo.X == 0 && eventInfo.Y == 0 {
		// initalise the text
		eventInfo.SetText(eventInfoText)
	}
	eventInfo.SetX(button.X + 12)
	eventInfo.SetY(button.Y + button.Height + 200)
	eventInfo.Draw(screen)
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
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && button.IsMouseOver() {
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
// a createButton, presenceInfo and eventInfo are passed to this function so they
// can be reset when a client is closed.
func updateCloseClientButton(closeButton *button.Button, createButton *button.Button, presenceInfo *text.Text, eventInfo *text.Text, id connectionID) {
	// Handle mouseover interaction with a close client button.
	if closeButton.IsMouseOver() {
		closeButton.SetTextColour(colour.White)
	} else {
		closeButton.SetTextColour(colour.Black)
	}

	// Handle mouse click on a close client button.
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && closeButton.IsMouseOver() {
		closeRealtimeClient(id)
		infoBar.SetText(closeRealtimeClientSuccess)
		// Reset the create button once a client is closed.
		createButton.SetBgColour(colour.Yellow)
		createButton.SetText(createClientText)

		// Reset the presence text once a client is closed.
		presenceInfo.Reset()

		// Reset the eventInfo text
		eventInfo.Reset()
	}
}

// updateSetChannelButton contains the update logic for each set channel button.
func updateSetChannelButton(button *button.Button, channelName string, id connectionID) {
	// Handle mouseover interaction with a set channel button.
	if button.IsMouseOver() {
		button.SetBgColour(colour.White)
	} else {
		button.SetBgColour(colour.Green)
	}

	// Handle mouse click on set channel button.
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && button.IsMouseOver() {
		// if the connection exists and does not have a channel.
		if connections[id] != nil && connections[id].channel == nil {
			setChannel(channelName, id)
			infoBar.SetText(setChannelSuccess)
		}
	}
}

// updateChannelPublishButton contains the update logic for each channel publish button.
func updateChannelPublishButton(button *button.Button, messageName string, messageData interface{}, id connectionID) {
	// Handle mouseover interaction with a leave presence button.
	if button.IsMouseOver() {
		button.SetBgColour(colour.White)
	} else {
		button.SetBgColour(colour.Magenta)
	}

	// if a connection exists that has a channel
	if connections[id] != nil && connections[id].channel != nil {

		// and this button has been clicked.
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && button.IsMouseOver() {
			err := publishToChannel(id, messageName, messageData)
			if err != nil {
				infoBar.SetText(err.Error())
				return
			}
			infoBar.SetText(publishToChannelSuccess)
		}
	}
}

//updateSubscribeChannelButton contains the logic to update a subscribe button.
// An event info text box is passed to this function so events that occur while
// subscribed can be drawn to the screen.
func updateSubscribeChannelButton(button *button.Button, eventInfo *text.Text, id connectionID) {

	// If a connection exists and no unsubscribe function is saved
	if connections[id] != nil && connections[id].unsubscribe == nil {
		button.SetText(subscribeAllText)
	}

	// Handle mouseover interaction with subscribe all button while the channel is not subscribed.
	if button.IsMouseOver() && connections[id] != nil && connections[id].unsubscribe == nil {
		button.SetBgColour(colour.White)
	}

	// if the button is not moused over and the channel is not subscribed.
	if !button.IsMouseOver() && connections[id] != nil && connections[id].unsubscribe == nil {
		button.SetBgColour(colour.Yellow)
	}

	// if the button is clicked
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && button.IsMouseOver() {

		// if a channel exists and the connection has no unsubscribe function saved
		if connections[id] != nil && connections[id].channel != nil && connections[id].unsubscribe == nil {

			unsubscribeAll, err := subscribeAll(id, eventInfo)
			if err != nil {
				infoBar.SetText(err.Error())
			}
			// Save the unsubscribe function.
			connections[id].unsubscribe = &unsubscribeAll
			infoBar.SetText(subscribeAllSuccess)

			// Change the SubscribeAll button into an Unsubscribe button.
			button.SetBgColour(colour.Red)
			button.SetText(unsubscribeText)
			return
		}

		// if there is an unsubscribe function saved
		if connections[id] != nil && connections[id].unsubscribe != nil {
			unsubscribe(id)
			infoBar.SetText(unsubscribeSuccess)
			// tear down channel unsubcribe function
			connections[id].unsubscribe = nil
			eventInfo.Reset()

			return
		}
	}
}

// updateGetPresenceButton contains the update logic for each announce presence button.
func updateAnnouncePresenceButton(button *button.Button, id connectionID) {
	// Handle mouseover interaction with an announce presence button.
	if button.IsMouseOver() {
		button.SetBgColour(colour.White)
	} else {
		button.SetBgColour(colour.Cyan)
	}

	// if a connection exists that has a channel
	if connections[id] != nil && connections[id].channel != nil {

		// and the button is clicked
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && button.IsMouseOver() {

			err := announcePresence(id)
			if err != nil {
				infoBar.SetText(err.Error())
				return
			}
			infoBar.SetText(announcePresenceSuccess)
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

	// if a connection exists that has a channel
	if connections[id] != nil && connections[id].channel != nil {

		// and the button is clicked
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && button.IsMouseOver() {
			// the call to get presence is async to prevent blocking.
			go getPresence(id, text)
		}
	}
}

// updateLeavePresenceButton contains the update logic for each leave presence button.
func updateLeavePresenceButton(button *button.Button, id connectionID) {
	// Handle mouseover interaction with a leave presence button.
	if button.IsMouseOver() {
		button.SetBgColour(colour.White)
	} else {
		button.SetBgColour(colour.Cyan)
	}

	// if a connection exists that has a channel
	if connections[id] != nil && connections[id].channel != nil {
		// and the button is clicked
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && button.IsMouseOver() {
			err := leavePresence(id)
			if err != nil {
				infoBar.SetText(err.Error())
				return
			}
			infoBar.SetText(leavePresenceSuccess)
		}
	}
}

// updateChannelNameInput contains the update logic for a channel name input text box
func updateChannelNameInput(textBox *textbox.TextBox, id connectionID) {

	// a mouse click anywhere which is not over the text box will remove focus from it.
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && !textBox.IsMouseOver() {
		textBox.RemoveFocus()
	}

	// a mouse click on the text box will set focus to it.
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && textBox.IsMouseOver() {
		textBox.SetFocus()
	}

	textBox.Update()
}

// updateMessageNameInput contains the update logic for a message name input text box
func updateMessageNameInput(textBox *textbox.TextBox, id connectionID) {

	// a mouse click anywhere which is not over the text box will remove focus from it.
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && !textBox.IsMouseOver() {
		textBox.RemoveFocus()
	}

	// a mouse click on the text box will set focus to it.
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && textBox.IsMouseOver() {
		textBox.SetFocus()
	}

	textBox.Update()
}

// updateMessageDataInput contains the update logic for a message data input text box
func updateMessageDataInput(textBox *textbox.TextBox, id connectionID) {

	// a mouse click anywhere which is not over the text box will remove focus from it.
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && !textBox.IsMouseOver() {
		textBox.RemoveFocus()
	}

	// a mouse click on the text box will set focus to it.
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && textBox.IsMouseOver() {
		textBox.SetFocus()
	}

	textBox.Update()
}
