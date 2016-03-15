package main

import (
    "fmt"
)

type vector [2]float64

func (self *vector) multiply(factor float64) {
    for i := range self {
        self[i] *= factor
    }
}

func (self *vector) add(other *vector) {
    for i := range self {
        self[i] += other[i]
    }
}

func (self *vector) String() string {
	return fmt.Sprintf("[%v, %v]", self[0], self[1])
}
