package main

import (
	"time"

	"github.com/135yshr/tinygo-keeb-ex/pkg/color"
	"github.com/135yshr/tinygo-keeb-ex/pkg/pin"

	"tinygo.org/x/drivers"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
	"tinygo.org/x/tinyfont/gophers"
)

func main() {
	disp := pin.NewDisplay()
	for _, r := range []drivers.Rotation{drivers.Rotation0, drivers.Rotation90, drivers.Rotation180, drivers.Rotation270} {
		display := disp.Rotate(r)
		display.Clear()
		time.Sleep(500 * time.Millisecond)

		tinyfont.WriteLine(display, &freemono.Bold9pt7b, 5, 10, "hello", color.White)
		tinyfont.WriteLine(display, &gophers.Regular58pt, 10, 70, "B", color.White)
		tinyfont.WriteLine(display, &gophers.Regular58pt, 10, 110, "H", color.White)
		display.Display()
		time.Sleep(5 * time.Second)
	}
}
