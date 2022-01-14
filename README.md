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
* Can subscribe to all messages on a channel
* Can unsubscribe from all messages on a channel

TODO:

* Publish to a channel
* Display messages received on a channel
* Display channel presence information
* Add some nice graphics

## Setup 
If you don't have one already, [create a new Ably account](https://ably.com/sign-up)

Create a new environment variable called `ABLY_PRIVATE_KEY` and set it to your Ably private key.


## Things to investigate in the SDK.

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


