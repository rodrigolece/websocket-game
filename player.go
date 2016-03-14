package main

import (
	"encoding/json"
	"log"
	"github.com/gorilla/websocket"
)

const (
	maxPlayerSpeed = 3.
)

type player struct {
	ws     *websocket.Conn
	output chan []byte
	id     string
	// addr    *net.Addr

	// Antes se creaba en reader pero sólo se necesita uno por jugador
	*wsEvent // Embedded para entrar directo a las variables

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

	self.wsEvent = &wsEvent{}

	// Valores de prueba para empezar
	self.pos = &vector{0, 0}
	self.vel = &vector{0.1, 0.2}
	self.vel.multiply(frameS) // incluye el intervalo de tiempo

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
		json.Unmarshal(event, self.wsEvent)
		log.Printf("%s -> %s\n", self.ws.RemoteAddr(), event)
		// go handleWsEvent(c, j) esto ahora lo hace cada función?
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
		// log.Println("Succesfully created new player.")
	}
	self.ws.Close()
}

func (self *player) tick() {
	self.pos.add(self.vel)
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
	// Meter aquí in cálculo
	return true
}
