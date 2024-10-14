package pin

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/encoders"
)

type RotaryPushButtonEvent func()

type RotaryReleasedEvent func()

type Rotary struct {
	enc *encoders.QuadratureDevice
	btn machine.Pin

	pressedEvents []RotaryPushButtonEvent
	releasedEvents []RotaryReleasedEvent
}

func NewRotary(precision int) *Rotary {
	enc := encoders.NewQuadratureViaInterrupt(
		machine.GPIO3,
		machine.GPIO4,
	)
	enc.Configure(encoders.QuadratureConfig{
		Precision: precision,
	})

	btn := machine.GPIO2
	btn.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	return &Rotary{
		enc: enc,
		btn: btn,
	}
}

func (r *Rotary) Position() int {
	return r.enc.Position()
}

func (r *Rotary) AddPressedEvent(f RotaryPushButtonEvent) {
	r.pressedEvents = append(r.pressedEvents, f)
}

func (r *Rotary) ClearPressedEvents() {
	r.pressedEvents = []RotaryPushButtonEvent{}
}

func (r *Rotary) AddReleasedEvent(f RotaryReleasedEvent) {
	r.releasedEvents = append(r.releasedEvents, f)
}

func (r *Rotary) ClearReleasedEvents() {
	r.releasedEvents = []RotaryReleasedEvent{}
}

func (r *Rotary) ThreadPressed() {
	for {
		if !r.btn.Get() {
			for _, f := range r.pressedEvents {
				f()
			}
			time.Sleep(10 * time.Millisecond)

			for !r.btn.Get() {}
			for _, f := range r.releasedEvents {
				f()
			}
			time.Sleep(10 * time.Millisecond)
		}
	}
}
