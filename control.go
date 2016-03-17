package main

type control struct {
    // -1, 0 y 1: bool no sirve porque son solo dos estados
    Accel   int    `json:"accel"`
    Turn    int    `json:"turn"`
	Name string `json:"name"`
}

func newControl() *control {
	self := &control{}
	return self
}
