# Ableye

This is a demo project created by Rosie Hamilton using Go and Ably Realtime.

Ableye is a visualisation tool which can be used to explore the ably-go SDK.

Similar to how Postman can be used to explore and test APIs, Ableye can be used to explore and test the ably-go SDK.

## Setup 

If you don't have one already, [create a new Ably account](https://ably.com/sign-up)

Create a new environment variable called `ABLY_PRIVATE_KEY` and set it to your Ably private key.

## Build and run

Building Ableye requires a minimum version of Go `1.17` to be installed.

Build the Ableye executable by navigating to the project root in a terminal window and running the command `go build`.

Run the executable with `./Ableye`

## Exploring Ably Realtime

Left mouse click on the `New Client` buttons to create up to 4 realtime clients.

A channel can be set on a client by inputting a channel name and left mouse clicking the `Set Channel` button.

Once a channel has been initialised, the channel can be subscribed to by clicking the `Subscribe All` button. After subscribing to a channel a window will appear to display events in realtime. It is also possible to attach or detach from the channel at any time by clicking on the `Attach` or `Detach` buttons.

Presence can be interacted with using the `Announce`, `Get` and `Leave` buttons.

Message name and message data can be input and published to the channel using the `Publish` button.

## Known issues

* The `Announce` button must sometimes be clicked twice for a client to be able to see their own presence in a channel. See https://github.com/ably/ably-go/issues/436
