package main

import (
	"time"

	"github.com/135yshr/tinygo-keeb-ex/pkg/color"
	"github.com/135yshr/tinygo-keeb-ex/pkg/pin"

	pio "github.com/tinygo-org/pio/rp2-pio"
)

func main() {
	s, _ := pio.PIO0.ClaimStateMachine()

	led := pin.NewLED(s)
	led.Set(color.White)

	for {
		time.Sleep(time.Millisecond * 500)
		led.Set(color.Black)
		time.Sleep(time.Millisecond * 500)
		led.Set(color.White)
		time.Sleep(time.Millisecond * 500)
		led.Set(color.Black)
		time.Sleep(time.Millisecond * 500)
		led.Set(color.Red)
	}
}
