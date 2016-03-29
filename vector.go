package main

import (
    "fmt"
)

type vector [2]float64 // The vector used for pos and vel

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
    // A stringer method is necessary to send the vector over JSON
	return fmt.Sprintf("[%v, %v]", self[0], self[1])
}
