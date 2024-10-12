package main

import "github.com/135yshr/tinygo-keeb-ex/pkg/pin"

func main() {
	vol := pin.NewVolume(4)

	for oldValue := 0; ; {
		if newValue := vol.Value(); newValue != oldValue {
			println("value: ", newValue)
			oldValue = newValue
		}
	}
}
