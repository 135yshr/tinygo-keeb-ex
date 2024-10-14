package main

import (
	"time"

	"github.com/135yshr/tinygo-keeb-ex/pkg/color"
	"github.com/135yshr/tinygo-keeb-ex/pkg/pin"

	"tinygo.org/x/tinydraw"
)

func main() {
	display := pin.NewDisplay()
	display.Clear()
	time.Sleep(50 * time.Millisecond)

	for {
		tinydraw.Rectangle(display, 10, 20, 30, 40, color.White)
		display.Display()
		time.Sleep(500 * time.Millisecond)

		tinydraw.FilledCircle(display, 60, 50, 10, color.White)
		display.Display()
		time.Sleep(500 * time.Millisecond)

		tinydraw.Triangle(display, 100, 10, 80, 40, 60, 10, color.White)
		display.Display()
		time.Sleep(500 * time.Millisecond)

		display.Clear()
		time.Sleep(500 * time.Millisecond)
	}
}
