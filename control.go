package main

// This struct is used to read the events from the keyboard arrows
type control struct {
    // -1, 0 y 1: bool does not work because it can only hold two states
    Accel   int    `json:"accel"`
    Turn    int    `json:"turn"`
	Name string `json:"name"`
}

func newControl() *control {
	self := &control{}
	return self
}
