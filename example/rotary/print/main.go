package main

import "github.com/135yshr/tinygo-keeb-ex/pkg/pin"

func main() {
	rot := pin.NewRotary()

	for oldValue := 0; ; {
		if newValue := rot.Position(); newValue != oldValue {
			println("value: ", newValue)
			oldValue = newValue
		}
	}
}
