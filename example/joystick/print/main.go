package main

import (
	"fmt"
	"time"

	"github.com/135yshr/tinygo-keeb-ex/pkg/pin"
)

func main() {
	joystick := pin.NewJoystick()

	for {
		x, y := joystick.Position()
		pressed := joystick.IsPressed()
		fmt.Printf("%04X %04X %t\n", x, y, pressed)
		time.Sleep(200 * time.Millisecond)
	}
}
