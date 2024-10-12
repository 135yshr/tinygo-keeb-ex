package main

import (
	"time"

	"github.com/135yshr/tinygo-keeb-ex/pkg/color"
	"github.com/135yshr/tinygo-keeb-ex/pkg/pin"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
	"tinygo.org/x/tinyfont/gophers"
)

func main() {
	display := pin.NewDisplay()
	display.Clear()
	time.Sleep(50 * time.Millisecond)

	for _, r := range []pin.Rotate{pin.Rotate90, pin.Rotate270} {
		display.Clear()
		rotDisplay := display.Rotated(r)

		tinyfont.WriteLine(rotDisplay, &freemono.Bold9pt7b, 5, 10, "hello", color.White)
		tinyfont.WriteLine(rotDisplay, &gophers.Regular58pt, 10, 70, "B", color.White)
		tinyfont.WriteLine(rotDisplay, &gophers.Regular58pt, 10, 110, "H", color.White)
		display.Display()
		time.Sleep(5 * time.Second)
	}
}
