package pin

import (
	"machine"

	"tinygo.org/x/drivers/encoders"
)

type Rotary struct {
	enc *encoders.QuadratureDevice
}

func NewRotary(precision int) *Rotary {
	enc := encoders.NewQuadratureViaInterrupt(
		machine.GPIO3,
		machine.GPIO4,
	)
	enc.Configure(encoders.QuadratureConfig{
		Precision: precision,
	})

	return &Rotary{
		enc: enc,
	}
}

func (r *Rotary) Position() int {
	return r.enc.Position()
}
