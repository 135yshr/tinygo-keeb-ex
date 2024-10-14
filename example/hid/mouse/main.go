package main

import (
	"time"

	"github.com/135yshr/tinygo-keeb-ex/pkg/pin"
	"machine/usb/hid/mouse"
)

func main() {
	m := mouse.Port()
	rotary := pin.NewRotary(4)
	rotary.AddPressedEvent(func() {
		m.Press(mouse.Left)
	})
	rotary.AddReleasedEvent(func() {
		m.Release(mouse.Left)
	})
	go rotary.ThreadPressed()


	for {
		time.Sleep(100 * time.Millisecond)
	}
}
