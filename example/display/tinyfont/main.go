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

	tinyfont.WriteLine(display, &freemono.Bold9pt7b, 5, 10, "hello", color.White)
	tinyfont.WriteLine(display, &gophers.Regular32pt, 5, 50, "ABCEF", color.White)
	display.Display()
}
