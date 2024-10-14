package main

import (
	"time"

	"github.com/135yshr/tinygo-keeb-ex/pkg/color"
	"github.com/135yshr/tinygo-keeb-ex/pkg/pin"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/gophers"
)

func main() {
	display := pin.NewDisplay()
	display.Clear()
	time.Sleep(50 * time.Millisecond)

	data := []byte("ABCEF")
	for {
		display.ClearBuffer()
		data[0], data[1], data[2], data[3], data[4] = data[1], data[2], data[3], data[4], data[0]
		tinyfont.WriteLine(display, &gophers.Regular32pt, 5, 45, string(data), color.White)
		display.Display()
		time.Sleep(200 * time.Millisecond)
	}
}
