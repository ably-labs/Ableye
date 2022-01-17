package main

var (

	// Log messages
	createRealtimeClientSuccess = "Successfully created a new realtime client."
	closeRealtimeClientSuccess  = "Successfully closed a realtime client."
	setChannelSuccess           = "Successfully set channel for realtime client."
	subscribeAllSuccess         = "Successfully subscribed to all channel messages."
	unsubscribeAllSuccess       = "Successfully unsubscribed from all channel messages."
	announcePresenceSuccess     = "Successfully announced presence."
	leavePresenceSuccess        = "Successfully removed presence."

	// Button text
	createClientText     = "Create Client"
	setChannelText       = "Set Channel"
	subscribeAllText     = "Subscribe All"
	unsubscribeAllText   = "Unsubscribe"
	announcePresenceText = "Announce"
	getPresenceText      = "Get"
	leavePresenceText    = "Leave"

	// Async processes
	startGetPresence    = "go routine started to get presence."
	completeGetPresence = "go routine completed to get presence."
)
