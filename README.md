#websocket-game

**websocket-game** is a minimum implementation of a massively multiplayer online
(MMO) game server, written in Go, and a simple client which takes the form of an
HTML web page. The client-server communication uses a WebSocket connection and a
homemade protocol written in JSON which we will further discuss.

**Disclaimer:** The code was developed mainly as an exercise, to practice Go and
JavaScript and to understand how WebSockets work. The implementation is naive
in many aspects because of this.

## The game
The game is still at an early stage (since most of the effort was put into
server and the client). So far every player that connects to the server is
assigned a ship or particle that has position and velocity, and moves in space. The
ship is controlled with the keyboard arrows. Certainly, this is not a game yet.
The plan is to add obstacles such as black holes and probably shooting or "eating"
capabilities.

## The server
The server is written in Go, and it heavily borrows from [this project][1].

The server serves the web page at `/`, manages JavaScript scripts at `/js` and
WebSocket connections at `/ws`.

Every WebSocket connection is associated with a `player` struct. Each `player` has
`reader` and `writer` methods to read and write on the connection. The server also
maintains the player's position and velocity, and calculates its next position. This
is used in order to send corrections to the clients (since the players' positions
will start to diverge after a while).

The server can easily be extended to use bot players, but they are not currently
implemented.

## The client
The client is inspired by [this project][2], but it is a much more simple
implementation. Basic functions are separated on different scripts.

`gas.js` holds the functions that represent the particles, and the functions used
to move them and create the animation.

`listen.js` registers the event handlers for pressing and releasing the keyboard
arrows.

Finally, `conn.js` establishes the WebSocket connection and listens for the server's
instructions (such as creating or removing a player).

The easiest way to have communication between JS scripts is through global variables.
Variables that are used in several scripts are declared before anything else in
`declarations.js`. This is far from ideal, and one of the first improvements we
need to make is writing real modules (for example with `require.js`).

## The protocol

The communication between server and clients is made possible by a homemade protocol
that was established much by trial and error, and has room for lots of improvement.

One of the most important realizations when working on the protocol is that communication
in one direction or the other does not need to use the same structure. A more typical
set of instructions is used for the server's events. They take the form of an `event`
attribute, an `id` attribute and if necessary data such as position and velocity
vectors. For example, we have the create and destroy player events:

```
{
    "event":"createPlayer",
    "id":"...",
    "pos":[..., ...],
    "vel":[..., ...]
}

{
    "event":"destroyPlayer",
    "id":"..."
}
```

The events we want to send from the client to the server are basically the arrows
events, and we can use a completely different approach.

In `control.js` we define the object

```
var control = {
    accel: 0,
    turn: 0
};
```

in which we store the current status of the keyboard arrows. If the up arrow is
pressed, we have and acceleration event and `accel` is set to `1`. If the down arrow
is pressed, on the contrary we set `accel` to `-1`. We have something equivalent
for the left and right arrows and the `turn` attribute. Every time we have a change
of state, we send `control` over the WebSocket connection and the server reads if
there is a turn or and acceleration event.

## Acknowledgements

This project was developed as part of a [menteslibres.io][3] scholarship.

[1]: https://github.com/xiam/shooter-server
[2]: https://github.com/xiam/shooter-html5
[3]: https://menteslibres.io
