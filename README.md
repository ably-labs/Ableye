# rosie-demo

This is my demo project using Go and Ably Realtime.

I want to create a tool which can be used to do exploratory testing of Go SDK
Similar to how Postman can be used to exploratory testing of an API


## Work in progress

Work in progress project name: Ableye

Completed:
* Screen to select which client (realtime or rest)
* Can create realtime clients
* Can close realtime clients
* Can set a default channel for a realtime client
* Clients can subscribe to all messages on a channel
* Clients can unsubscribe from all messages on a channel
* Clients can announce their channel presence 
* Clients can get channel presence 
* Clients can remove channel presence (leave)
* Clients can publish a message to a channel
* Clients can display messages received on a channel

TODO:

* Give clients a way to change the data that they publish to a channel
* Increase the number of connections that can be created from 2 to 4
* Replicate functionality for the realtime client for the rest client
* Display the SDK version on the title screen
* Add some nice graphics

## Setup 
If you don't have one already, [create a new Ably account](https://ably.com/sign-up)

Create a new environment variable called `ABLY_PRIVATE_KEY` and set it to your Ably private key.


## Things to investigate further in the SDK.

### Missing pointer safety
When attempting to create a new client and passing nil to the constructor for a new realtime client 
 ```
	_, err := ably.NewRealtime(nil)
 ```
This causes the SDK to panic when it tries apply default options. Stack trace
```
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x4384130]

goroutine 50 [running]:
github.com/ably/ably-go/ably.applyOptionsWithDefaults({0xc0002edbd0, 0x1, 0xc0002edb10})
        /Users/rosiehamilton/go/pkg/mod/github.com/ably/ably-go@v1.2.3/ably/options.go:822 +0x110
github.com/ably/ably-go/ably.NewREST({0xc0002edbd0, 0x0, 0xd3})
        /Users/rosiehamilton/go/pkg/mod/github.com/ably/ably-go@v1.2.3/ably/rest_client.go:118 +0x25
github.com/ably/ably-go/ably.NewRealtime({0xc0002edbd0, 0x1, 0x1})
        /Users/rosiehamilton/go/pkg/mod/github.com/ably/ably-go@v1.2.3/ably/realtime_client.go:23 +0x52
```

### Channel State
Once a channel has subscribed to all messages, its state changes from `INITIALIZED` to `ATTACHED`.
If the unsubscribe function is then called, the channel state remains `ATTACHED` even though its been unsubscribed.
I was expecting the state to change to something else as it's no longer attached.

### Auto Documentation / Code Comments

It looks like the tool tip for `channel.Presence.Get` says `Get returns a list of current members on the channel, attaching the channel first is needed.` However it seems that trying to get presence on a channel which has status `INITIALIZED` will automatically change the channel status to `ATTACHED` so the statement that attaching the channel first might not be true. 

### Announce Presence
Trying to annouce presence using the example code from the readme `err = channel.Presence.Enter(ctx, "presence data")` 
results in the following error `[ErrorInfo : code=91000 unable to enter presence channel (no clientId) statusCode=0] See https://help.ably.io/error/91000`

Note that following the link in the message `https://help.ably.io/error/91000` is a dead link with no information is available for that error code.

Most of the time, the first time a client announces presence with `presence.Enter` it silently fails. Only on rare occasions does announcing work the first time the client announces. However, announcing always seems to work the second time the client announces. This should be investigated further.

### Get Presence
It feels like sometimes an async call to get presence returns presence, but sometimes it does not. Could there be a race condition bug here?

### ably-go readme observations
The example code for announcing presence on a channel using `channel.Presence.Enter`, assigns an error to an existing variable with `=` operator, this should probably change to assign it to a new variable with `:=` operator. This information should also be update to say that to use this method, a clientID must be set on the client.

### missing message fields
* Ably messages have a field called `Timestamp` of type `int64`. It appears that once a message is received, it's timestamp is always `0`. Was expecting the timestamp to not have a nil value. Was also expecting this to be of type `time.Time`. 
* The `ConnectionID` field appears to always be empty
* The `Encoding` field appears to always be empty
* The `ID` field appears to always be empty
* The `Extras` field appears to always be sent as an empty slice of map.
