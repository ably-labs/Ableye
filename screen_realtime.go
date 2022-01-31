package main

import (
	"fmt"

	"github.com/ably-labs/Ableye/button"
	colour "github.com/ably-labs/Ableye/colours"
	font "github.com/ably-labs/Ableye/fonts"
	"github.com/ably-labs/Ableye/text"
	"github.com/ably-labs/Ableye/textbox"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type connectionElements struct {
	id                   connectionID
	createRealtimeClient button.Button
	createRestClient     button.Button
	clientLabel          button.Button
	closeClient          button.Button
	setChannel           button.Button
	channelName          text.Text
	channelStatus        text.Text
	channelDetach        button.Button
	channelAttach        button.Button
	channelSubscribeAll  button.Button
	presenceInfo         text.Text
	enterPresence        button.Button
	getPresence          button.Button
	leavePresence        button.Button
	eventInfo            text.Text
	channelNameLabel     text.Text
	channelNameInput     textbox.TextBox
	messageNameLabel     text.Text
	messageNameInput     textbox.TextBox
	messageDataLabel     text.Text
	messageDataInput     textbox.TextBox
	channelPublish       button.Button
}

// The elements of the realtime screen.
var (
	infoBar button.Button

	//Connection Elements
	connectionA, connectionB, connectionC, connectionD connectionElements
)

func initialiseRealtimeScreen() {
	infoBar = button.NewButton(screenWidth, 35, "", 22, 22, colour.Black, font.MplusSmallFont, colour.White, 0, 25)

	//Initialise connection elements.
	initialiseConnectionElements(&connectionA, clientA, 0, screenHeight/6)
	initialiseConnectionElements(&connectionB, clientB, screenWidth/2, screenHeight/6)
	initialiseConnectionElements(&connectionC, clientC, 0, (screenHeight/2)+75)
	initialiseConnectionElements(&connectionD, clientD, screenWidth/2, (screenHeight/2)+75)
}

// initialiseConnectionElements, creates all the button, text and text box elements
// that are required for a connection with correct colours, sizes, fonts and settings
// and saves the connectionElement information in a global variable
func initialiseConnectionElements(elements *connectionElements, id connectionID, x int, y int) {
	elements.id = id
	elements.createRealtimeClient = button.NewButton(150, 35, createClientText, 16, 22, colour.Black, font.MplusSmallFont, colour.White, x, y)
	elements.createRestClient = button.NewButton(150, 35, createRestClientText, 32, 22, colour.Black, font.MplusSmallFont, colour.White, x+151, y)
	elements.clientLabel = button.NewButton(400, 35, "", 12, 22, colour.Black, font.MplusSmallFont, colour.ZingyGreen, x, y)
	elements.closeClient = button.NewButton(35, 35, "X", 12, 22, colour.Black, font.MplusSmallFont, colour.BrightRed, 0, 0)
	elements.channelName = text.NewText("", colour.ZingyGreen, font.MplusSmallFont, 0, 0)
	elements.channelStatus = text.NewText("", colour.ZingyGreen, font.MplusSmallFont, 0, 0)
	elements.channelDetach = button.NewButton(75, 30, detachText, 10, 20, colour.Black, font.MplusSmallFont, colour.ZingyGreen, 0, 0)
	elements.channelAttach = button.NewButton(75, 30, attachText, 10, 20, colour.Black, font.MplusSmallFont, colour.ZingyGreen, 0, 0)
	elements.channelSubscribeAll = button.NewButton(115, 30, subscribeAllText, 10, 20, colour.Black, font.MplusSmallFont, colour.ZingyGreen, 0, 0)
	elements.presenceInfo = text.NewText("", colour.ElectricCyan, font.MplusSmallFont, 0, 0)
	elements.enterPresence = button.NewButton(65, 30, enterPresenceText, 12, 20, colour.Black, font.MplusSmallFont, colour.ElectricCyan, 0, 0)
	elements.getPresence = button.NewButton(50, 30, getPresenceText, 12, 20, colour.Black, font.MplusSmallFont, colour.ElectricCyan, 0, 0)
	elements.leavePresence = button.NewButton(70, 30, leavePresenceText, 12, 20, colour.Black, font.MplusSmallFont, colour.ElectricCyan, 0, 0)
	elements.eventInfo = text.NewText("", colour.White, font.MplusSmallFont, 0, 0)
	elements.channelNameLabel = text.NewText(fmt.Sprintf("%s :", channelNameText), colour.ZingyGreen, font.MplusSmallFont, 0, 0)
	elements.channelNameInput = textbox.NewTextBox(200, 36, 4, defaultChannelName, 9, 12, 22, colour.ZingyGreen, font.MplusSmallFont, colour.Black, colour.ZingyGreen, 0, 0)
	elements.setChannel = button.NewButton(150, 35, setChannelText, 22, 22, colour.Black, font.MplusSmallFont, colour.ZingyGreen, 0, 0)
	elements.messageNameLabel = text.NewText(fmt.Sprintf("%s :", messageNameText), colour.JazzyPink, font.MplusSmallFont, 0, 0)
	elements.messageNameInput = textbox.NewTextBox(200, 36, 4, defaultMessageName, 12, 12, 22, colour.JazzyPink, font.MplusSmallFont, colour.Black, colour.JazzyPink, 0, 0)
	elements.messageDataLabel = text.NewText(fmt.Sprintf("%s :", messageDataText), colour.JazzyPink, font.MplusSmallFont, 0, 0)
	elements.messageDataInput = textbox.NewTextBox(200, 36, 4, defaultMessageData, 12, 12, 22, colour.JazzyPink, font.MplusSmallFont, colour.Black, colour.JazzyPink, 0, 0)
	elements.channelPublish = button.NewButton(80, 30, publishText, 12, 20, colour.Black, font.MplusSmallFont, colour.JazzyPink, 0, 0)
}

// drawRealtimeScreen draws the realtime screen.
func drawRealtimeScreen(screen *ebiten.Image) {

	// Info bar is used to display log messages and error messages.
	infoBar.Draw(screen)

	//Draw elements for each connection.
	drawConnectionElements(screen, &connectionA)
	drawConnectionElements(screen, &connectionB)
	drawConnectionElements(screen, &connectionC)
	drawConnectionElements(screen, &connectionD)
}

// drawConnectionElements draws all the elements associated with a connection to the screen.
func drawConnectionElements(screen *ebiten.Image, elements *connectionElements) {

	id := elements.id

	// if there is no connection
	if connections[id] == nil {
		// draw the buttons that can create a connection with a client.
		elements.createRealtimeClient.Draw(screen)
		elements.createRestClient.Draw(screen)
	}

	// if client has been created
	if connections[id] != nil && (connections[id].client != nil || connections[id].restClient != nil) {

		// move the client label to the location of the create realtime button
		elements.clientLabel.X = elements.createRealtimeClient.X
		elements.clientLabel.Y = elements.createRealtimeClient.Y
		if elements.clientLabel.GetText() == "" {
			if isRealtimeClient(id) {
				elements.clientLabel.SetText(fmt.Sprintf("%s - %s", connections[id].client.Auth.ClientID(), connections[id].connectionType.string()))
			}
			if isRestClient(id) {
				elements.clientLabel.SetText(fmt.Sprintf("%s - %s", connections[id].restClient.Auth.ClientID(), connections[id].connectionType.string()))
			}
		}

		elements.clientLabel.Draw(screen)

		drawClientInfo(screen, elements.clientLabel, &elements.closeClient)

		elements.closeClient.Draw(screen)

		// if a channel has not been set for this client, draw the elements required to set the channel.
		if connections[id].channel == nil {
			drawSetChannel(screen, elements.clientLabel, elements.channelNameLabel, &elements.channelNameInput, &elements.setChannel)
		}
	}

	// if client channel has been created
	if connections[id] != nil && connections[id].channel != nil {
		drawChannelInfo(screen, elements)
	}

	// if a channel has been subscribed an unsubscribe function will be saved in memory.
	// if an unsubscribe function exists, draw event info
	if connections[id] != nil && connections[id].unsubscribe != nil {
		drawEventInfo(screen, elements.clientLabel, &elements.eventInfo)
	}
}

// drawClientInfo draws the client window and the close client button.
func drawClientInfo(screen *ebiten.Image, clientLabel button.Button, closeClient *button.Button) {
	// Elements are drawn in locations that are calculated from the location of the client label.
	button := clientLabel

	// Draw a window which is made from two overlapping images.
	ebitenutil.DrawRect(screen, float64(button.X), float64(button.Y)+float64(button.Height), (screenWidth/2)-10, screenHeight/3, colour.ZingyGreen)
	ebitenutil.DrawRect(screen, float64(button.X)+1, float64(button.Y)+float64(button.Height)+1, (screenWidth/2)-12, (screenHeight/3)-2, colour.Black)

	// Draw the close client button.
	closeClient.SetX(button.X + 638)
	closeClient.SetY(button.Y)
	closeClient.Draw(screen)
}

func drawSetChannel(screen *ebiten.Image, clientLabel button.Button, channelNameLabel text.Text, channelNameInput *textbox.TextBox, setChannel *button.Button) {
	// Elements are drawn in locations that are calculated from the location of the client label.
	button := clientLabel

	// Channel name label.
	channelNameLabel.SetX(button.X + 10)
	channelNameLabel.SetY(button.Y + button.Height + 40)
	channelNameLabel.Draw(screen)

	// Channel name input text box.
	channelNameInput.SetX(button.X + 150)
	channelNameInput.SetY(button.Y + button.Height + 18)
	channelNameInput.Draw(screen)

	// Set channel button.
	setChannel.SetX(button.X + 375)
	setChannel.SetY(button.Y + button.Height + 18)
	setChannel.Draw(screen)
}

// drawChannelInfo draws the channel window, the presence window and message controls.
func drawChannelInfo(screen *ebiten.Image, elements *connectionElements) {
	// Elements are drawn in locations that are calculated from the location of the client label.
	button := elements.clientLabel
	id := elements.id

	// Draw the channel window.
	ebitenutil.DrawRect(screen, float64(button.X)+4, float64(button.Y)+float64(button.Height)+3, (screenWidth/2)-18, screenHeight/10, colour.ZingyGreen)
	ebitenutil.DrawRect(screen, float64(button.X)+5, float64(button.Y)+float64(button.Height)+4, (screenWidth/2)-20, (screenHeight/10)-2, colour.Black)

	// Draw the channel name text box.
	elements.channelName.SetX(button.X + 10)
	elements.channelName.SetY(button.Y + button.Height + 25)
	elements.channelName.SetText(fmt.Sprintf("%s : %s", channelNameText, connections[id].channel.Name))
	elements.channelName.Draw(screen)

	// Draw the channel status text box.
	elements.channelStatus.SetX(button.X + 230)
	elements.channelStatus.SetY(button.Y + button.Height + 25)
	elements.channelStatus.SetText(fmt.Sprintf("Status : %s", connections[id].channel.State()))
	elements.channelStatus.Draw(screen)

	//Draw the channel detach button.
	elements.channelDetach.SetX(button.X + 401)
	elements.channelDetach.SetY(button.Y + button.Height + 4)
	elements.channelDetach.Draw(screen)

	//Draw the channel attach button.
	elements.channelAttach.SetX(button.X + 477)
	elements.channelAttach.SetY(button.Y + button.Height + 4)
	elements.channelAttach.Draw(screen)

	// Draw the channel subscribe/unsubscribe button.
	elements.channelSubscribeAll.SetX(button.X + 553)
	elements.channelSubscribeAll.SetY(button.Y + button.Height + 4)
	elements.channelSubscribeAll.Draw(screen)

	// Draw the presence window.
	ebitenutil.DrawRect(screen, float64(button.X)+8, float64(button.Y)+float64(button.Height)+42, (screenWidth/2)-26, screenHeight/24, colour.ElectricCyan)
	ebitenutil.DrawRect(screen, float64(button.X)+9, float64(button.Y)+float64(button.Height)+43, (screenWidth/2)-28, (screenHeight/24)-2, colour.Black)

	// If presence information is being drawn in its initisalised location.
	if elements.presenceInfo.X == 0 && elements.presenceInfo.Y == 0 {
		// Initalise the text in the presence information.
		elements.presenceInfo.SetText("Presence :")
	}

	// Draw the presence information.
	elements.presenceInfo.SetX(button.X + 12)
	elements.presenceInfo.SetY(button.Y + button.Height + 62)
	elements.presenceInfo.Draw(screen)

	// Draw the enter presence button.
	elements.enterPresence.SetX(button.X + 477)
	elements.enterPresence.SetY(button.Y + button.Height + 43)
	elements.enterPresence.Draw(screen)

	// Draw the get presence button.
	elements.getPresence.SetX(button.X + 543)
	elements.getPresence.SetY(button.Y + button.Height + 43)
	elements.getPresence.Draw(screen)

	// Draw the leave presence button.
	elements.leavePresence.SetX(button.X + 594)
	elements.leavePresence.SetY(button.Y + button.Height + 43)
	elements.leavePresence.Draw(screen)

	// Draw the message name label.
	elements.messageNameLabel.SetX(button.X + 10)
	elements.messageNameLabel.SetY(button.Y + button.Height + 105)
	elements.messageNameLabel.Draw(screen)

	// Draw the message name input text box.
	elements.messageNameInput.SetX(button.X + 10)
	elements.messageNameInput.SetY(button.Y + button.Height + 120)
	elements.messageNameInput.Draw(screen)

	// Draw the message data label.
	elements.messageDataLabel.SetX(button.X + 300)
	elements.messageDataLabel.SetY(button.Y + button.Height + 105)
	elements.messageDataLabel.Draw(screen)

	// Draw the message data input text box.
	elements.messageDataInput.SetX(button.X + 300)
	elements.messageDataInput.SetY(button.Y + button.Height + 120)
	elements.messageDataInput.Draw(screen)

	// Draw the channel publish button.
	elements.channelPublish.SetX(button.X + 550)
	elements.channelPublish.SetY(button.Y + button.Height + 125)
	elements.channelPublish.Draw(screen)
}

// drawEventInfo draws the event window.
func drawEventInfo(screen *ebiten.Image, clientLabel button.Button, eventInfo *text.Text) {

	// Elements are drawn in locations that are calculated from the location of the client label.
	button := clientLabel

	// Draw the event window.
	ebitenutil.DrawRect(screen, float64(button.X)+4, float64(button.Y)+float64(button.Height)+182, (screenWidth/2)-18, screenHeight/12, colour.White)
	ebitenutil.DrawRect(screen, float64(button.X)+5, float64(button.Y)+float64(button.Height)+183, (screenWidth/2)-20, (screenHeight/12)-2, colour.Black)

	// If event information is being drawn in its initisalised location.
	if eventInfo.X == 0 && eventInfo.Y == 0 {
		// Initalise the text in the event information.
		eventInfo.SetText(eventInfoText)
	}

	// Draw the event information.
	eventInfo.SetX(button.X + 12)
	eventInfo.SetY(button.Y + button.Height + 200)
	eventInfo.Draw(screen)
}

// updateRealtimeScreen updates the realtime screen.
func updateRealtimeScreen() {
	// If the Escape key is pressed, return to the title screen.
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		state = titleScreen
	}

	// Update elements for each connection.
	updateConnectionElements(&connectionA)
	updateConnectionElements(&connectionB)
	updateConnectionElements(&connectionC)
	updateConnectionElements(&connectionD)
}

// updateConnectionElements updates all the elements for a connection.
func updateConnectionElements(elements *connectionElements) {
	// Client elements.
	updateCreateRealtimeClientButton(&elements.createRealtimeClient, elements.id)
	updateCreateRestClientButton(&elements.createRestClient, elements.id)
	updateCloseClientButton(&elements.closeClient, &elements.clientLabel, &elements.presenceInfo, &elements.eventInfo, elements.id)

	// Set channel elements.
	updateTextInputBox(&elements.channelNameInput, elements.id)
	updateSetChannelButton(&elements.setChannel, elements.channelNameInput.GetText(), elements.id)

	// Channel controls.
	updateChannelDetachButton(&elements.channelDetach, elements.id)
	updateChannelAttachButton(&elements.channelAttach, elements.id)
	updateSubscribeChannelButton(&elements.channelSubscribeAll, &elements.eventInfo, elements.id)

	// Presence controls.
	updateEnterPresenceButton(&elements.enterPresence, elements.id)
	updateGetPresenceButton(&elements.getPresence, &elements.presenceInfo, elements.id)
	updateLeavePresenceButton(&elements.leavePresence, elements.id)

	// Message controls.
	updateTextInputBox(&elements.messageNameInput, elements.id)
	updateTextInputBox(&elements.messageDataInput, elements.id)
	updateChannelPublishButton(&elements.channelPublish, elements.messageNameInput.GetText(), elements.messageDataInput.GetText(), elements.id)
}

// updateCreateRealtimeClientButton contains the update logic for each realtime client creation button.
func updateCreateRealtimeClientButton(button *button.Button, id connectionID) {
	// Handle mouseover interaction with create button while a connection does not exist.
	if button.IsMouseOver() && connections[id] == nil {
		button.SetBgColour(colour.ZingyGreen)
	}

	// if the button is not moused over and there is no connection.
	if !button.IsMouseOver() && connections[id] == nil {
		button.SetBgColour(colour.White)
	}

	// Handle mouse click on a create realtime client button.
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && button.IsMouseOver() {
		if connections[id] == nil {
			if err := createRealtimeClient(id); err != nil {
				infoBar.SetText(err.Error())
			}
			infoBar.SetText(createRealtimeClientSuccess)
		}
	}
}

// updateCreateRestClientButton contains the update logic for each realtime client creation button.
func updateCreateRestClientButton(button *button.Button, id connectionID) {
	// Handle mouseover interaction with create button while a connection does not exist.
	if button.IsMouseOver() && connections[id] == nil {
		button.SetBgColour(colour.ZingyGreen)
	}

	// if the button is not moused over and there is no connection.
	if !button.IsMouseOver() && connections[id] == nil {
		button.SetBgColour(colour.White)
	}

	// Handle mouse click on a create rest client button.
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && button.IsMouseOver() {
		if connections[id] == nil {
			if err := createRestClient(id); err != nil {
				infoBar.SetText(err.Error())
			}
			infoBar.SetText(createRestClientSuccess)
		}
	}
}

// updateCloseClientButton contains the update logic for each close client button.
// a clientLabel, presenceInfo and eventInfo are passed to this function so they
// can be reset when a client is closed.
func updateCloseClientButton(closeButton *button.Button, clientLabel *button.Button, presenceInfo *text.Text, eventInfo *text.Text, id connectionID) {
	// Handle mouseover interaction with a close client button.
	if closeButton.IsMouseOver() {
		closeButton.SetTextColour(colour.White)
	} else {
		closeButton.SetTextColour(colour.Black)
	}

	// Handle mouse click on a close client button.
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && closeButton.IsMouseOver() {

		// if a realtime client is being closed.
		if isRealtimeClient(id) {
			closeRealtimeClient(id)
			infoBar.SetText(closeRealtimeClientSuccess)
		}

		// if a rest client is being closed
		if isRestClient(id) {
			closeRestClient(id)
			infoBar.SetText(closeRealtimeClientSuccess)
		}

		// Reset the client label.
		clientLabel.SetText("")

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
		button.SetBgColour(colour.ZingyGreen)
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
	// Handle mouseover interaction with a channel publish button.
	if button.IsMouseOver() {
		button.SetBgColour(colour.White)
	} else {
		button.SetBgColour(colour.JazzyPink)
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
			infoBar.SetText(publishSuccess)
		}
	}
}

// updateChannelDetachhButton contains the update logic for each channel attach button.
func updateChannelDetachButton(button *button.Button, id connectionID) {
	// Handle mouseover interaction with a channel attach button.
	if button.IsMouseOver() {
		button.SetBgColour(colour.White)
	} else {
		button.SetBgColour(colour.ZingyGreen)
	}

	// if a connection exists that has a channel
	if connections[id] != nil && connections[id].channel != nil {

		// and this button has been clicked.
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && button.IsMouseOver() {
			err := detachChannel(id)
			if err != nil {
				infoBar.SetText(err.Error())
				return
			}
			infoBar.SetText(detachChannelSuccess)
		}
	}
}

// updateChannelAttachButton contains the update logic for each channel attach button.
func updateChannelAttachButton(button *button.Button, id connectionID) {
	// Handle mouseover interaction with a channel attach button.
	if button.IsMouseOver() {
		button.SetBgColour(colour.White)
	} else {
		button.SetBgColour(colour.ZingyGreen)
	}

	// if a connection exists that has a channel
	if connections[id] != nil && connections[id].channel != nil {

		// and this button has been clicked.
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && button.IsMouseOver() {
			err := attachChannel(id)
			if err != nil {
				infoBar.SetText(err.Error())
				return
			}
			infoBar.SetText(attachChannelSuccess)
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
		button.SetBgColour(colour.ZingyGreen)
	}

	// if the button is clicked
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && button.IsMouseOver() {

		// if a channel exists and the connection has no unsubscribe function saved
		if connections[id] != nil && connections[id].channel != nil && connections[id].unsubscribe == nil {

			unsubscribeAll, err := subscribeAll(id, eventInfo)
			if err != nil {
				infoBar.SetText(err.Error())
				return
			}
			// Save the unsubscribe function.
			connections[id].unsubscribe = &unsubscribeAll
			infoBar.SetText(subscribeAllSuccess)

			// Change the SubscribeAll button into an Unsubscribe button.
			button.SetBgColour(colour.BrightRed)
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

// updateEnterPresenceButton contains the update logic for each enter presence button.
func updateEnterPresenceButton(button *button.Button, id connectionID) {
	// Handle mouseover interaction with an enter presence button.
	if button.IsMouseOver() {
		button.SetBgColour(colour.White)
	} else {
		button.SetBgColour(colour.ElectricCyan)
	}

	// if a connection exists that has a channel
	if connections[id] != nil && connections[id].channel != nil {

		// and the button is clicked
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && button.IsMouseOver() {

			err := enterPresence(id)
			if err != nil {
				infoBar.SetText(err.Error())
				return
			}
			infoBar.SetText(enterPresenceSuccess)
		}
	}
}

// updateGetPresenceButton contains the update logic for each get presence button
func updateGetPresenceButton(button *button.Button, text *text.Text, id connectionID) {
	// Handle mouseover interaction with a get presence button.
	if button.IsMouseOver() {
		button.SetBgColour(colour.White)
	} else {
		button.SetBgColour(colour.ElectricCyan)
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
		button.SetBgColour(colour.ElectricCyan)
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

// updateTextInputBox contains the update logic for text input text boxs
func updateTextInputBox(textBox *textbox.TextBox, id connectionID) {
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

// isRealtimeClient is a helper function that returns true if a connection exists for an ID and it's type is realtime.
func isRealtimeClient(id connectionID) bool {
	return connections[id] != nil && connections[id].connectionType != nil && *connections[id].connectionType == realtime
}

// isRestClient is a helper function that returns true if a connection exists for an ID and it's type is rest.
func isRestClient(id connectionID) bool {
	return connections[id] != nil && connections[id].connectionType != nil && *connections[id].connectionType == rest
}
