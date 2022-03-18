# Ableye

This is a demo project created by Rosie Hamilton using Go and Ably Realtime.

Ableye is a visualisation tool which can be used to explore the ably-go SDK. 

Similar to how Postman can be used to explore and test APIs, Ableye can be used to explore and test the ably-go SDK.

[This blog post](https://ably.com/blog/ablyeye-how-we-visualized-an-ably-sdk-with-go-and-ebiten) explains why and how this project was created. 

## Setup 

If you don't have one already, [create a new Ably account](https://ably.com/sign-up)

Create a new environment variable called `ABLY_PRIVATE_KEY` and set it to your Ably private key.

## Build and run

Building Ableye requires a minimum version of Go `1.17` to be installed.

Build the Ableye executable by navigating to the project root in a terminal window and running the command `go build`.

Run the executable with `./Ableye`

## Exploring Ably Realtime

A new client can be created by left mouse clicking on either the `Realtime Client` or the `Rest client` buttons. Up to 4 clients can be created simultaneously.

### Realtime Client

After creating a realtime client, a channel can be set by inputting a channel name then left mouse clicking on the `Set Channel` button.

Once a channel has been initialised, it can be subscribed to by clicking the `Subscribe All` button. After subscribing to a channel a window will appear to display events in realtime. While subscribed, an `Unsubscribe` button will be displayed which can be used to unsubscribe. It is also possible to attach or detach from the channel at any time by clicking on the `Attach` or `Detach` buttons.

Presence can be interacted with using the `Enter`, `Get` and `Leave` buttons.

Message name and message data can be input and published to the channel using the `Publish` button.

### Rest Client

The rest client supports a smaller number of features than the realtime client.

After creating a rest client, a channel can be set by inputting a channel name then left mouse clicking on the `Set Channel` button.

It is possibly to get the channel presence by clicking on the `Get` button. 

Message name and message data can be input and published to the channel using the `Publish` button.

Please note that the rest client does not support entering and leaving a channel. The rest client also does not support subscribing to a channel.
