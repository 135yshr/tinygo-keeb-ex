package pin

import (
	"image/color"
	"machine"

	pio "github.com/tinygo-org/pio/rp2-pio"
	"github.com/tinygo-org/pio/rp2-pio/piolib"
)

type LED struct {
	ws *piolib.WS2812B
}

func NewLED(s pio.StateMachine) *LED {
	ws, _ := piolib.NewWS2812B(s, machine.GPIO16)
	return &LED{
		ws: ws,
	}
}

func (l *LED) Set(c color.Color) {
	l.ws.PutColor(c)
}
