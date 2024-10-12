package pin

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers"
	"tinygo.org/x/drivers/ssd1306"
)

const (
	Rotate0 Rotate = iota
	Rotate90
	Rotate180
	Rotate270
)

type Rotate int

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

func (d *Display) Device() *ssd1306.Device {
	return &d.disp
}

func (d *Display) Rotated(r Rotate) *RotatedDisplay {
	return &RotatedDisplay{&d.disp, r}
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

type RotatedDisplay struct {
	drivers.Displayer
	rotate Rotate
}

func (d *RotatedDisplay) Size() (x, y int16) {
	return y, x
}

func (d *RotatedDisplay) SetPixel(x, y int16, c color.RGBA) {
	sx, sy := d.Displayer.Size()
	switch d.rotate {
	case Rotate0:
		d.Displayer.SetPixel(x, y, c)
	case Rotate90:
		d.Displayer.SetPixel(y, sy-x, c)
	case Rotate180:
		d.Displayer.SetPixel(sx-y, sy-x, c)
	case Rotate270:
		d.Displayer.SetPixel(sx-y, x, c)
	}
}
