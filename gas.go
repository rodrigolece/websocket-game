package main

import (
    "time"
)

const (
    lx = 1.
    ly = 1.
    radiusParticle = float64(1) / 50
)

// A gas holds the set of players (particles)
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
        // We update the position of the players and we send it to them
        for p := range self.players {
            p.tick()
            p.update()
        }

        // Remove players
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

        // Necessary to run at a constant rate of fps
        time.Sleep(time.Duration(sleep) * time.Nanosecond)
        /* time.Sleep takes a duration for an argument. time.Nanosecond
        is of type Duration */
    }
}

func (self *gas) broadcast(event []byte) {
    for player, _ := range self.players {
        if player.ws != nil { // bots have nil
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

    // We announce the new player
    event = createPlayerEvent(p)
    self.broadcast(event)

    // To the new player we announce existing players
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
            // We only take unused ids
            return id
        }
    }
}
