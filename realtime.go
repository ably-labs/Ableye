package main

import (
	"context"
	"github.com/ably-labs/rosie-demo/config"
	ably "github.com/ably/ably-go/ably"
)

var (
	client  *ably.Realtime
	channel *ably.RealtimeChannel
)

func createRealtimeClient() error {
	cl, err := ably.NewRealtime(ably.WithKey(config.Cfg.Key))
	client = cl
	return err
}

func getChannel() {
	channel = client.Channels.Get("test")
}

func publishToChannel(ctx context.Context) error{
	return channel.Publish(ctx, "EventName1", "EventData1")
}
