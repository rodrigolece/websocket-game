package main

import (
    "time"
    "math/rand"
)

const letters = "abcdefghijklmnopqrstuvwxyz"

func init() {
    rand.Seed(time.Now().UnixNano())
}

// Used to generate random id
func randString(n int) string {
    b := make([]byte, n)
    // The result of letters[i] isof type uint8 (byte)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func randVector() *vector {
    return &vector{rand.Float64(), rand.Float64()}
}

func newPos() *vector {
    out := randVector()
    out.multiply(1 - 2 * radiusParticle)
    out.add(&vector{radiusParticle, radiusParticle})
    return out

    // We might want to use no overlap once the particles interact with each other

    // do {
    //     var overlap = false;
    //     var x = gas.boxSize.lx * Math.random() + radiusParticle;
    //     var y = gas.boxSize.ly * Math.random() + radiusParticle;
    //     if (n == 0) {
    //         return { x: x, y: y };
    //     }
    //     for (var i = 0; i < n; i++){
    //         var part = gas.particles[i];
    //         var dx = part.pos.x - x;
    //         var dy = part.pos.y - y;
    //         var norm = Math.sqrt( dx*dx + dy*dy );
    //         if (norm <= minD) { overlap = true; }
    //     }
    // } while (overlap) ;
}
