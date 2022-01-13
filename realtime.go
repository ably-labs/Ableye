package main

import (
	"context"
	"github.com/ably-labs/rosie-demo/config"
	ably "github.com/ably/ably-go/ably"
	"log"
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
	client  *ably.Realtime
	channel *ably.RealtimeChannel
}

// newConnection is a contructor to create a new connection.
func newConnection(client *ably.Realtime) connection {
	return connection{
		client: client,
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
		connections[id] = nil
		log.Println(closeRealtimeClientSuccess)
	}
}

func setChannel(id connectionID) {
	newChannel := connections[id].client.Channels.Get(defaultChannel)
	connections[id].channel = newChannel
	log.Println(setChannelSuccess)
}

func publishToChannel(ctx context.Context) error {
	//return channel.Publish(ctx, "EventName1", "EventData1")
	return nil
}
