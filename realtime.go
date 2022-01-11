package main

import (
	"log"
	"context"
	"github.com/ably-labs/rosie-demo/config"
	ably "github.com/ably/ably-go/ably"
)

var (
	client  *ably.Realtime
	channel *ably.RealtimeChannel
)

func createRealtimeClient() error {
	if client == nil {
		cl, err := ably.NewRealtime(ably.WithKey(config.Cfg.Key))
		if err != nil {
			return err
		}
		client = cl
		log.Println(createRealtimeClientSuccess)
	}

	return nil
}

func getChannel() {
	channel = client.Channels.Get("test")
}

func publishToChannel(ctx context.Context) error{
	return channel.Publish(ctx, "EventName1", "EventData1")
}
