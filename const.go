package main

import (
	"time"
)

// Constant global variables.
const (
	screenWidth    = 1366
	screenHeight   = 768
	defaultTimeout = 5 * time.Second

	/////////////////////////
	// Title Screen.
	////////////////////////

	titleText    = "Ableye"
	taglineText  = "A visualisation tool for exploring Ably."
	realtimeText = "Explore Ably Realtime"

	/////////////////////////
	// Ably Realtime Screen.
	/////////////////////////

	// Default input text box contents.
	defaultChannelName = "test"
	defaultMessageName = "message"
	defaultMessageData = "data"

	// Button text
	createClientText     = "Realtime Client"
	createRestClientText = "Rest Client"
	setChannelText       = "Set Channel"
	publishText          = "Publish"
	detachText           = "Detach"
	attachText           = "Attach"
	subscribeAllText     = "Subscribe All"
	unsubscribeText      = "Unsubscribe"
	enterPresenceText    = "Enter"
	getPresenceText      = "Get"
	leavePresenceText    = "Leave"

	// Display text
	channelNameText = "Channel"
	messageNameText = "Name"
	messageDataText = "Data"
	eventInfoText   = "Events will be displayed here in realtime."

	/////////////////////////
	// Log messages.
	/////////////////////////

	createRealtimeClientSuccess = "Successfully created a new realtime client."
	createRestClientSuccess     = "Successfully created a new rest client."
	closeRealtimeClientSuccess  = "Successfully closed a realtime client."
	closeRestClientSuccess      = "Successfully closed a rest client."
	setChannelSuccess           = "Successfully set channel for realtime client."
	detachChannelSuccess        = "Successfully detached from channel."
	attachChannelSuccess        = "Successfuly attached to channel."
	publishSuccess              = "Successfully published to channel."
	subscribeAllSuccess         = "Successfully subscribed to all channel messages."
	unsubscribeSuccess          = "Successfully unsubscribed from channel messages."
	enterPresenceSuccess        = "Successfully entered presence."
	leavePresenceSuccess        = "Successfully removed presence."

	// Async processes
	startGetPresence    = "go routine started to get presence."
	completeGetPresence = "go routine completed to get presence."
)
