package pin

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/ssd1306"
)

type Display struct {
	disp ssd1306.Device
}

func NewDisplay() *Display {
	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: 2.8 * machine.MHz,
		SDA:       machine.GPIO12,
		SCL:       machine.GPIO13,
	})

	disp := ssd1306.NewI2C(machine.I2C0)
	disp.Configure(ssd1306.Config{
		Address: 0x3C,
		Width:   128,
		Height:  64,
	})

	return &Display{
		disp: disp,
	}
}

func (d *Display) SetPixel(x, y int16, c color.RGBA) {
	d.disp.SetPixel(x, y, c)
}

func (d *Display) Display() error {
	return d.disp.Display()
}

func (d *Display) Clear() {
	d.disp.ClearDisplay()
}
