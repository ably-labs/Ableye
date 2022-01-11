# rosie-demo

This is my demo project using Go and Ably Realtime.

I want to create a tool which can be used to do exploratory testing of Go SDK
Similar to how Postman can be used to exploratory testing of an API

Work in progress project name: Ableye

Want two clickable buttons to select which client is being tested

Create Read Update Delete of channels - in a visual way

Subscribe a client to a channel

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