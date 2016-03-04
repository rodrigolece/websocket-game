package main

type vector struct {
    x, y float64
}

func (self *vector) multiply(factor float64) {
    self.x *= factor
    self.y *= factor
}

func (self *vector) add(other *vector) {
    self.x += other.x
    self.y += other.y
}
