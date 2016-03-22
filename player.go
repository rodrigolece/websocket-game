package main

import (
	"encoding/json"
	// "log"
	"github.com/gorilla/websocket"
	"math"
)

const (
	// maxPlayerSpeed = 3.
	velScaling = 1e-2
	turnAngle = math.Pi/30
)

type player struct {
	ws     *websocket.Conn
	output chan []byte
	id     string
	// addr    *net.Addr

	*control // Embedded para entrar directo a las variables

	gas *gas
	pos *vector // Se pordrían mover a *entity
	vel *vector
}

func newPlayer(ws *websocket.Conn) *player {
	self := &player{}
	self.ws = ws

	if ws != nil {
		self.output = make(chan []byte, 256) // buffereado para que no bloquee
	} else {
		self.output = nil
	}

	self.control = newControl()

	// Valores de prueba para empezar
	self.pos = randVector()
	self.vel = randVector()
	self.vel.multiply(velScaling)
	// self.vel.multiply(frameS)

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
		Ya no es necesario porque tick usa directamente el cambio de estado */
	}
	self.ws.Close()
	// Hay más cosas que se tienen que hacer para matar a un jugador
}

func (self *player) writer() {
	for event := range self.output {
		// por este range de aquí es importante cerrar el canal en hub
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
		if p.ws != nil && p.output != nil { // No se manda a los bots
			if self.isNear(p) {
				p.send(event)
			}
		}
	}
}

func (self *player) isNear(other *player) bool {
	// Meter aquí un cálculo
	return true
}

func (self *player) tick() {
	angle := 0.0

	// Aceleración y rotación

	if self.Accel > 0 {
		self.vel.multiply(1.02)
	} else if self.Accel < 0 {
		self.vel.multiply(0.98)
	}

	if self.Turn > 0 {
		angle = -turnAngle
	} else if self.Turn < 0 {
		angle = turnAngle
	}

	var c = math.Cos(angle)
	var s = math.Sin(angle)
	self.vel[0]  = c * self.vel[0] - s * self.vel[1]
	self.vel[1] = s * self.vel[0] + c * self.vel[1]

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
