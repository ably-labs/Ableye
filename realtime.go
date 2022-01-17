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
	context        context.Context
	client         *ably.Realtime
	channel        *ably.RealtimeChannel
	presence       []*ably.PresenceMessage
	unsubscribeAll *func()
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
func createRealtimeClient(id connectionID) error {

	newClient, err := ably.NewRealtime(ably.WithKey(config.Cfg.Key))
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
		connections[id].unsubscribeAll = nil
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
func subscribeAll(id connectionID) (func(), error) {
	unsubscribe, err := connections[id].channel.SubscribeAll(connections[id].context, printAblyMessage)

	if err != nil {
		return nil, err
	}
	log.Println(subscribeAllSuccess)
	return unsubscribe, nil
}

// unsubscribeAll calls a connections unsubscribe all function if it exists.
func unsubscribeAll(id connectionID) {
	if connections[id].unsubscribeAll != nil {
		unsubscribeFunc := *connections[id].unsubscribeAll
		unsubscribeFunc()
		log.Println(unsubscribeAllSuccess)
	}

}

// func publishToChannel(ctx context.Context) error {
// 	//return channel.Publish(ctx, "EventName1", "EventData1")
// 	return nil
// }

func printAblyMessage(msg *ably.Message) {
	fmt.Printf("Received message: name=%s data=%v\n", msg.Name, msg.Data)
}

// announcePresence announces the presence of a client to a channel.
func announcePresence(id connectionID) error {

	// Set timeout to be default timeout
	ctx, _ := context.WithTimeout(connections[id].context, defaultTimeout)
	err := connections[id].channel.Presence.Enter(ctx, "presence data")
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// getPresence sets the presence info text box to presence information.
func getPresence(id connectionID, presenceInfo *text.Text) {
	var buffer bytes.Buffer

	log.Println(startGetPresence)

	// Set timeout to be default timeout
	ctx, _ := context.WithTimeout(connections[id].context, defaultTimeout)

	presenceMessages, _ := connections[id].channel.Presence.Get(ctx)

	for _, v := range presenceMessages {
		if v != nil {
			buffer.WriteString(v.ID)
			buffer.WriteString(" ")
			buffer.WriteString(v.ConnectionID)
			buffer.WriteString(" ")
			buffer.WriteString(v.Name)
			buffer.WriteString(" ")
			buffer.WriteString(v.ClientID)
			buffer.WriteString(" ")
			buffer.WriteString(v.Message.String())
			buffer.WriteString(" ")
		}
	}
	presence := buffer.String()

	presenceInfo.SetText(fmt.Sprintf("Presence: %s", presence))
	log.Println(completeGetPresence)
}
