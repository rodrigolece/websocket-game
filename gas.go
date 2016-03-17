package main

import (
    "time"
)

const (
    lx = 1.
    ly = 1.
    radiusParticle = 1 / 50
)

type gas struct {
    // numPlayers   int
    players     map[*player]bool
    ids         map[string]*player

    offset      [2]float64
    boxSize     [2]float64

    remove      chan *player
}

func newGas() *gas {
    self := &gas{}
    self.players = map[*player]bool{}
    self.ids = map[string]*player{}

    self.boxSize = [2]float64{lx, ly}
    self.offset = [2]float64{0, 0}

    self.remove = make(chan *player)

    go self.run()

    return self
}

func (self *gas) run() {

    var start, timeElapsed, sleep int64

    for {
        start = time.Now().UnixNano()
        // Actualizamos posición de jugadores y les avisamos
        for p := range self.players {
            p.tick()
            p.update()
        }

        // Quitamos jugadores
        removing := true
        for removing {
            select {
            case p := <-self.remove:
                self.removePlayer(p)
            default:
                removing = false
            }
        }

        timeElapsed = time.Now().UnixNano() - start
        sleep = frameNs - timeElapsed

        // Necesario para que corra a un número constante de fps
        time.Sleep(time.Duration(sleep) * time.Nanosecond)
        /* time.Sleep recibe como argumento una duración. time.Nanosecond es
        de tipo Duration. */
    }
}

func (self *gas) broadcast(event []byte) {
    for player, _ := range self.players {
        if player.ws != nil { // bots tienen nil
            player.send(event)
        }
    }
}

func (self *gas) addPlayer(p *player) {
    var event []byte

    id := self.newId()

    self.players[p] = true
    self.ids[id] = p

    p.id = id
    p.gas = self

    // Anunciamos nuevo jugador
    event = createPlayerEvent(p)
    self.broadcast(event)

    // Al nuevo jugador le anunciamos los jugadores presentes
    for otherId, other := range self.ids {
        if otherId != id {
            event = createPlayerEvent(other)
            p.send(event)
        }
    }
}

func (self *gas) removePlayer(p *player) {
    delete(self.players, p)
    delete(self.ids, p.id)
    // close(p.send)
    p.gas = nil
}


func (self *gas) newId() string {
    var id string
    for {
        id = randString(4)
        if _, ok := self.ids[id] ; ok == false {
            // Sólo regresamos la id si no está en uso
            return id
        }
    }
}
