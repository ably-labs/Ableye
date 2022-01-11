package main

import (
	"context"
	"github.com/ably-labs/rosie-demo/config"
	ably "github.com/ably/ably-go/ably"
	"log"
)


func init(){
	// Initialise a map to store a client for each clientID. 
	clients = make(map[clientID]*ably.Realtime)
}

var (

	clients map[clientID]*ably.Realtime
	channel *ably.RealtimeChannel
)

type clientID int

const (
	clientA clientID = iota
	clientB
	clientC
	clientD
)

func createRealtimeClient(id clientID) error {

	if clients[id] == nil {
		newClient, err := ably.NewRealtime(ably.WithKey(config.Cfg.Key))
		if err != nil {
			return err
		}

		clients[id] = newClient
		log.Println(createRealtimeClientSuccess)
	}

	return nil
}

func getChannel() {
//	channel = client.Channels.Get("test")
}

func publishToChannel(ctx context.Context) error {
	//return channel.Publish(ctx, "EventName1", "EventData1")
	return nil
}
