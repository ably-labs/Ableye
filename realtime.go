package main

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/ably-labs/rosie-demo/config"
	"github.com/ably-labs/rosie-demo/text"
	ably "github.com/ably/ably-go/ably"
)

func init() {
	// Initialise a map to store connections to the ably platform.
	connections = make(map[connectionID]*connection)
}

var (
	connections map[connectionID]*connection
)

type connectionID int

func (c connectionID) string() string {
	switch int(c) {
	case 0:
		return "Client A"
	case 1:
		return "Client B"
	case 2:
		return "Client C"
	case 3:
		return "Client D"
	}
	return ""
}

const (
	clientA connectionID = iota
	clientB
	clientC
	clientD
)

// connection represents a connection to the Ably platform.
type connection struct {
	context     context.Context
	client      *ably.Realtime
	channel     *ably.RealtimeChannel
	unsubscribe *func()
}

// newConnection is a contructor to create a new connection.
func newConnection(client *ably.Realtime) connection {
	ctx := context.Background()
	return connection{
		context: ctx,
		client:  client,
	}
}

// createRealtimeClient creates a new realtime client and stores it in a connection.
// A clientID is also set on the client.
func createRealtimeClient(id connectionID) error {

	newClient, err := ably.NewRealtime(
		ably.WithKey(config.Cfg.Key),
		ably.WithClientID(id.string()),
	)

	if err != nil {
		return err
	}

	connection := newConnection(newClient)
	connections[id] = &connection
	log.Println(createRealtimeClientSuccess)

	return nil
}

// closeRealtimeClient closes an existing realtime client and removes the connection.
func closeRealtimeClient(id connectionID) {

	if connections[id] != nil && connections[id].client != nil {
		connections[id].client.Close()

		//Tear down the connection in internal memory.
		connections[id].unsubscribe = nil
		connections[id] = nil

		log.Println(closeRealtimeClientSuccess)
	}
}

// setChannel sets the channel to the default channel for a connection.
func setChannel(id connectionID) {
	newChannel := connections[id].client.Channels.Get(defaultChannel)
	connections[id].channel = newChannel
	log.Println(setChannelSuccess)
}

// subscribeAll subscribes the connection's channel to all messsages.
// once subscribe events are output to the eventInfo text box
func subscribeAll(id connectionID, eventInfo *text.Text) (func(), error) {

	handlerFunc := newEventHandler(eventInfo)
	unsubscribe, err := connections[id].channel.SubscribeAll(connections[id].context, handlerFunc)

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
			eventInfo.SetText(fmt.Sprintf("Event : %s, %s, %s,", msg.ClientID, msg.Name, msg.Data))
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

func publishToChannel(id connectionID) error {

	// Set timeout to be default timeout
	ctx, cancel := context.WithTimeout(connections[id].context, defaultTimeout)
	defer cancel()

	if err := connections[id].channel.Publish(ctx, "Event Name", "Event Data"); err != nil {
		return err
	}
	log.Println(publishToChannelSuccess)
	return nil
}

// announcePresence announces the presence of a client to a channel.
func announcePresence(id connectionID) error {

	// Set timeout to be default timeout
	ctx, cancel := context.WithTimeout(connections[id].context, defaultTimeout)
	defer cancel()

	err := connections[id].channel.Presence.Enter(ctx, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(announcePresenceSuccess)
	return nil
}

// getPresence sets the presence info text box to presence information.
func getPresence(id connectionID, presenceInfo *text.Text) {
	var buffer bytes.Buffer

	log.Println(startGetPresence)

	// Set timeout to be default timeout
	ctx, cancel := context.WithTimeout(connections[id].context, defaultTimeout)
	defer cancel()

	presenceMessages, _ := connections[id].channel.Presence.Get(ctx)

	for i, v := range presenceMessages {
		if v != nil {
			buffer.WriteString(v.ClientID)
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

	err := connections[id].channel.Presence.Leave(ctx, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(leavePresenceSuccess)
	return nil
}
