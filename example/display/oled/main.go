package main

import (
	"time"

	"github.com/135yshr/tinygo-keeb-ex/pkg/color"
	"github.com/135yshr/tinygo-keeb-ex/pkg/pin"
)

func main() {
	display := pin.NewDisplay()
	display.Clear()
	time.Sleep(50 * time.Millisecond)

	cnt := 0
	for {
		c := color.White
		if cnt%2 == 1 {
			c = color.Black
		}
		for x := int16(0); x < 128; x += 2 {
			for y := int16(0); y < 64; y += 2 {
				display.SetPixel(x+0, y+0, c)
				display.SetPixel(x+0, y+1, c)
				display.SetPixel(x+1, y+0, c)
				display.SetPixel(x+1, y+1, c)
				display.Display()
			}
		}
		cnt++
	}
}
