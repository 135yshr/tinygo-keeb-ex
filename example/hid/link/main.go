package main

import (
	"time"

	"machine/usb/hid/keyboard"

	"github.com/135yshr/tinygo-keeb-ex/pkg/pin"
)

func main() {
	kb := keyboard.Port()
	rotary := pin.NewRotary(4)
	rotary.AddPressedEvent(func() {
		kb.Down(keyboard.KeyA)
	})
	rotary.AddReleasedEvent(func() {
		kb.Up(keyboard.KeyA)
	})
	go rotary.ThreadPressed()

	for {
		time.Sleep(100 * time.Millisecond)
	}
}
