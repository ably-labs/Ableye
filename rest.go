package main

import (
	"context"
	"log"

	"github.com/ably-labs/Ableye/config"
	ably "github.com/ably/ably-go/ably"
)

// createRestClient creates a new rest client and stores it in a connection.
// A clientID is also set on the client.
func createRestClient(id connectionID) error {

	newClient, err := ably.NewREST(
		ably.WithKey(config.Cfg.Key),
		ably.WithClientID(id.string()),
	)
	if err != nil {
		return err
	}

	connection := newRestConnection(newClient, rest)
	connections[id] = &connection
	log.Println(createRestClientSuccess)

	return nil
}

// closeRestClient closes an existing realtime client and removes the connection.
func closeRestClient(id connectionID) {

	if connections[id] != nil && connections[id].restClient != nil {

		//Tear down the connection in internal memory.
		connections[id].restClient = nil
		connections[id] = nil

		log.Println(closeRestClientSuccess)
	}
}

// setRestChannel sets the rest channel to the name provided in the channel name input text box.
func restSetChannel(name string, id connectionID) {
	newChannel := connections[id].restClient.Channels.Get(name)
	connections[id].restChannel = newChannel
	log.Println(setRestChannelSuccess)
}

// publishToRestChannel publishes message name and message data to a realtime channel.
func publishToRestChannel(id connectionID, messageName string, messageData interface{}) error {

	// Set timeout to be default timeout
	ctx, cancel := context.WithTimeout(connections[id].context, defaultTimeout)
	defer cancel()

	if err := connections[id].restChannel.Publish(ctx, messageName, messageData); err != nil {
		return err
	}

	log.Println(publishSuccess)
	return nil
}
