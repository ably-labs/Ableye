package main

import (
	"context"
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

type connectionType int

const (
	realtime connectionType = iota
	rest
)

func (c connectionType) string() string {
	switch int(c) {
	case 0:
		return "Realtime Client"
	case 1:
		return "Rest Client"
	}
	return ""
}

// connection represents a connection to the Ably platform.
type connection struct {
	context         context.Context
	connectionType  *connectionType
	restClient      *ably.REST
	realtimeClient  *ably.Realtime
	realtimeChannel *ably.RealtimeChannel
	restChannel     *ably.RESTChannel
	unsubscribe     *func()
}

// newRealtimeConnection is a contructor to create a new realtime connection.
func newRealtimeConnection(client *ably.Realtime, conType connectionType) connection {
	ctx := context.Background()
	return connection{
		context:        ctx,
		connectionType: &conType,
		realtimeClient: client,
	}
}

// newRestConnection is a contructor to create a new REST connection.
func newRestConnection(client *ably.REST, conType connectionType) connection {
	ctx := context.Background()
	return connection{
		context:        ctx,
		connectionType: &conType,
		restClient:     client,
	}
}
