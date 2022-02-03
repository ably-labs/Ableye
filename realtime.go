package main

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/ably-labs/Ableye/config"
	"github.com/ably-labs/Ableye/text"
	ably "github.com/ably/ably-go/ably"
)

// createRealtimeClient creates a new realtime client and stores it in a connection.
// A clientID is also set on the client.
func createRealtimeClient(id connectionID) error {

	var newClient *ably.Realtime

	options := []ably.ClientOption{
		ably.WithKey(config.Cfg.Key),
		ably.WithClientID(id.string()),
	}

	if config.Cfg.DebugMode {
		options = append(options, ably.WithLogLevel(ably.LogDebug))
	}

	client, err := ably.NewRealtime(options...)
	if err != nil {
		return err
	}
	newClient = client

	connection := newRealtimeConnection(newClient, realtime)
	connections[id] = &connection
	log.Println(createRealtimeClientSuccess)

	return nil
}

// closeRealtimeClient closes an existing realtime client and removes the connection.
func closeRealtimeClient(id connectionID) {

	if connections[id] != nil && connections[id].realtimeClient != nil {
		connections[id].realtimeClient.Close()

		//Tear down the connection in internal memory.
		connections[id].unsubscribe = nil
		connections[id] = nil

		log.Println(closeRealtimeClientSuccess)
	}
}

// realtimeSetChannel sets the channel to the name provided in the channel name input text box.
func realtimeSetChannel(name string, id connectionID) {
	newChannel := connections[id].realtimeClient.Channels.Get(name)
	connections[id].realtimeChannel = newChannel
	log.Println(setRealtimeChannelSuccess)
}

// detachChannel attaches a client to a channel.
func detachChannel(id connectionID) error {
	// Set timeout to be default timeout
	ctx, cancel := context.WithTimeout(connections[id].context, defaultTimeout)
	defer cancel()

	if err := connections[id].realtimeChannel.Detach(ctx); err != nil {
		return err
	}

	log.Println(detachChannelSuccess)
	return nil
}

// attachChannel attaches a client to a channel.
func attachChannel(id connectionID) error {
	// Set timeout to be default timeout
	ctx, cancel := context.WithTimeout(connections[id].context, defaultTimeout)
	defer cancel()

	if err := connections[id].realtimeChannel.Attach(ctx); err != nil {
		return err
	}

	log.Println(attachChannelSuccess)
	return nil
}

// realtimeSubscribeAll subscribes the connection's channel to all messsages.
// once subscribe events are output to the eventInfo text box
func realtimeSubscribeAll(id connectionID, eventInfo *text.Text) (func(), error) {

	handlerFunc := newEventHandler(eventInfo)
	unsubscribe, err := connections[id].realtimeChannel.SubscribeAll(connections[id].context, handlerFunc)

	if err != nil {
		return nil, err
	}
	log.Println(subscribeAllSuccess)
	return unsubscribe, nil
}

//newEventHandler returns a function that can handle a message event.
//This pattern allows dependencies to be injected into the handler function.
func newEventHandler(eventInfo *text.Text) func(*ably.Message) {
	return func(msg *ably.Message) {
		log.Printf("Received message: name=%s data=%v\n", msg.Name, msg.Data)

		if eventInfo != nil {
			eventInfo.SetText(fmt.Sprintf("Event Received From : %s\n\n%s : %s     %s: %s", msg.ClientID, messageNameText, msg.Name, messageDataText, msg.Data))
		}
	}
}

// unsubscribe calls a connections unsubscribe function if it exists.
func unsubscribe(id connectionID) {
	if connections[id].unsubscribe != nil {
		unsubscribeFunc := *connections[id].unsubscribe
		unsubscribeFunc()
		log.Println(unsubscribeSuccess)
	}

}

// publishToRealtimeChannel publishes message name and message data to a realtime channel.
func publishToRealtimeChannel(id connectionID, messageName string, messageData interface{}) error {

	// Set timeout to be default timeout
	ctx, cancel := context.WithTimeout(connections[id].context, defaultTimeout)
	defer cancel()

	if err := connections[id].realtimeChannel.Publish(ctx, messageName, messageData); err != nil {
		return err
	}

	log.Println(publishSuccess)
	return nil
}

// enterPresence informs the channel that the client has entered the channel.
func enterPresence(id connectionID) error {
	// Set timeout to be default timeout
	ctx, cancel := context.WithTimeout(connections[id].context, defaultTimeout)
	defer cancel()

	if err := connections[id].realtimeChannel.Presence.Enter(ctx, nil); err != nil {
		log.Println(err)
		return err
	}

	log.Println(enterPresenceSuccess)
	return nil
}

// getRealtimePresence sets the presence info text box to presence information.
func getRealtimePresence(id connectionID, presenceInfo *text.Text) {
	var buffer bytes.Buffer

	log.Println(startGetPresence)

	// Set timeout to be default timeout
	ctx, cancel := context.WithTimeout(connections[id].context, defaultTimeout)
	defer cancel()

	presenceMessages, err := connections[id].realtimeChannel.Presence.Get(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	for i, msg := range presenceMessages {
		if msg != nil {
			buffer.WriteString(msg.ClientID)
			// if not the last message, add a comma and a space.
			if i != len(presenceMessages)-1 {
				buffer.WriteString(", ")
			}
		}
	}
	presence := buffer.String()

	presenceInfo.SetText(fmt.Sprintf("Presence : %s", presence))
	log.Println(completeGetPresence)
}

// leavePresence removes the presence of a client from a channel.
func leavePresence(id connectionID) error {
	// Set timeout to be default timeout
	ctx, cancel := context.WithTimeout(connections[id].context, defaultTimeout)
	defer cancel()

	err := connections[id].realtimeChannel.Presence.Leave(ctx, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(leavePresenceSuccess)
	return nil
}
