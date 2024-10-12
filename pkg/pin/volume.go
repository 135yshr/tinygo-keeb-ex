package pin

import (
	"machine"

	"tinygo.org/x/drivers/encoders"
)

type Volume struct {
	enc *encoders.QuadratureDevice
}

func NewVolume() *Volume {
	enc := encoders.NewQuadratureViaInterrupt(
		machine.GPIO3,
		machine.GPIO4,
	)
	enc.Configure(encoders.QuadratureConfig{
		Precision: 4,
	})

	return &Volume{
		enc: enc,
	}
}

func (v *Volume) Value() int {
	return v.enc.Position()
}
