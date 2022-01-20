package main

var (

	// Log messages
	createRealtimeClientSuccess = "Successfully created a new realtime client."
	closeRealtimeClientSuccess  = "Successfully closed a realtime client."
	setChannelSuccess           = "Successfully set channel for realtime client."
	publishToChannelSuccess     = "Successfully published to channel."
	subscribeAllSuccess         = "Successfully subscribed to all channel messages."
	unsubscribeSuccess          = "Successfully unsubscribed from channel messages."
	announcePresenceSuccess     = "Successfully announced presence."
	leavePresenceSuccess        = "Successfully removed presence."

	// Button text
	createClientText     = "Create Client"
	setChannelText       = "Set Channel"
	publishText          = "Publish"
	subscribeAllText     = "Subscribe All"
	unsubscribeText      = "Unsubscribe"
	announcePresenceText = "Announce"
	getPresenceText      = "Get"
	leavePresenceText    = "Leave"

	// Display text
	channelNameText = "Channel Name"
	messageNameText = "Message Name"
	messageDataText = "Message Data"
	eventInfoText   = "Events will be displayed here in realtime."

	// Async processes
	startGetPresence    = "go routine started to get presence."
	completeGetPresence = "go routine completed to get presence."
)
