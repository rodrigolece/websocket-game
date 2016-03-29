package main

import (
	"encoding/json"
	// "log"
	"github.com/gorilla/websocket"
	"math"
)

const (
	// maxPlayerSpeed = 3. CURRENTLY NOT ENFORCED
	velScaling = 1e-2
	turnAngle = math.Pi/30
)

type player struct {
	ws     *websocket.Conn
	output chan []byte
	id     string

	*control // Embedded so we can read directly its varibles

	gas *gas
	pos *vector
	vel *vector
}

func newPlayer(ws *websocket.Conn) *player {
	self := &player{}
	self.ws = ws

	if ws != nil {
		self.output = make(chan []byte, 256) // buffered so it doesn't block
	} else {
		self.output = nil
	}

	self.control = newControl()

	self.pos = newPos()
	self.vel = randVector()
	self.vel.multiply(velScaling)

	return self
}

func (self *player) send(event []byte) {
	if self.output == nil {
		return
	}

	self.output <- event
}

func (self *player) identity() {
	if self.id != "" {
		self.send(identityEvent(self.id))
	}
}

// func (self *player) Serialize() (buf []byte) {
//
// }

func (self *player) reader() {
	for {
		_, event, err := self.ws.ReadMessage()
		if err != nil {
			break
		}
		json.Unmarshal(event, self.control)
		// log.Printf("%s -> %s\n", self.ws.RemoteAddr(), event)
		/* go handleWsEvent(c, j)
		No longer necessary, tick uses directly the new sate */
	}
	self.ws.Close()
	// We need to remove a player more carefully!
}

func (self *player) writer() {
	for event := range self.output {
		err := self.ws.WriteMessage(websocket.TextMessage, event)
		if err != nil {
			break
		}
		// log.Printf("%s <- %s\n", self.ws.RemoteAddr(), event)
	}
	self.ws.Close()
}

func (self *player) update() {
	event := updateEvent(self)

	for p := range self.gas.players {
		if p.ws != nil && p.output != nil { // Not sent to bots
			if self.isNear(p) {
				p.send(event)
			}
		}
	}
}

func (self *player) isNear(other *player) bool {
	// Calculation goes here
	return true
}

func (self *player) tick() {
	angle := 0.0

	// Acceletaton
	if self.Accel > 0 {
		self.vel.multiply(1.02)
	} else if self.Accel < 0 {
		self.vel.multiply(0.98)
	}

	// Rotation
	if self.Turn > 0 {
		angle = -turnAngle
	} else if self.Turn < 0 {
		angle = turnAngle
	}

	var c = math.Cos(angle)
	var s = math.Sin(angle)
	self.vel[0]  = c * self.vel[0] - s * self.vel[1]
	self.vel[1] = s * self.vel[0] + c * self.vel[1]

	// And we move the particle, taking into account reflection at the borders
	futurex := self.pos[0] + self.vel[0]
    futurey := self.pos[1] + self.vel[1]

    if (futurex + radiusParticle > lx || futurex - radiusParticle < 0) {
        self.vel[0] *= -1;
    }
    if (futurey + radiusParticle > ly || futurey - radiusParticle < 0) {
        self.vel[1] *= -1;
    }

	self.pos.add(self.vel)
}
