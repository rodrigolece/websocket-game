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
    // El resultado de letters[i] es de tipo uint8 (byte)
    b := make([]byte, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}
