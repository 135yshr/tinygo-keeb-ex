package pin

import (
	"machine"

	"tinygo.org/x/drivers/encoders"
)

type Volume struct {
	enc *encoders.QuadratureDevice
}

func NewVolume(precision int) *Volume {
	enc := encoders.NewQuadratureViaInterrupt(
		machine.GPIO3,
		machine.GPIO4,
	)
	enc.Configure(encoders.QuadratureConfig{
		Precision: precision,
	})

	return &Volume{
		enc: enc,
	}
}

func (v *Volume) Value() int {
	return v.enc.Position()
}
