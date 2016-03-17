package main

import (
    "time"
    "math/rand"
)

const letters = "abcdefghijklmnopqrstuvwxyz"

func init() {
    rand.Seed(time.Now().UnixNano())
}

func randString(n int) string {
    b := make([]byte, n)
    // El resultado de letters[i] es de tipo uint8 (byte)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func randVector() *vector {
    return &vector{rand.Float64(), rand.Float64()}
}
