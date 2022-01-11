package main

import (
	"context"
	"github.com/ably-labs/rosie-demo/config"
	ably "github.com/ably/ably-go/ably"
	"log"
)

func init() {
	// Initialise a map to store a client for each clientID.
	clients = make(map[clientID]*ably.Realtime)
}

var (
	clients map[clientID]*ably.Realtime
//	channel *ably.RealtimeChannel
)

type clientID int

const (
	clientA clientID = iota
	clientB
	clientC
	clientD
)

// createRealtimeClient creates a new realtime client and stores it in the clients map against the key of clientID.
func createRealtimeClient(id clientID) error {
	newClient, err := ably.NewRealtime(ably.WithKey(config.Cfg.Key))
	if err != nil {
		return err
	}

	clients[id] = newClient
	log.Println(createRealtimeClientSuccess)

	return nil
}

// closeRealtimeClient closes an existing realtime client and removes it from the clients map.
func closeRealtimeClient(id clientID) {
	if clients[id] != nil {
		clients[id].Close()
		clients[id] = nil
		log.Println(closeRealtimeClientSuccess)
	}
}

func setChannel() {

	// TO DO create some kind of struct (maybe a session?) to hold both a client and it's active channel.
	//	channel = client.Channels.Get("test")
}

func publishToChannel(ctx context.Context) error {
	//return channel.Publish(ctx, "EventName1", "EventData1")
	return nil
}
