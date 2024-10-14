package main

import (
	"time"

	"github.com/135yshr/tinygo-keeb-ex/pkg/pin"
)

func main() {
	r := pin.NewRotary(4)
	r.AddPressedEvent(func() {
		println("pressed")
	})
	r.AddReleasedEvent(func() {
		println("released")
	})

	go r.ThreadPressed()

	for {
		time.Sleep(100 * time.Millisecond)
	}
}
