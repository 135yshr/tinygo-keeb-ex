package pin

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers"
	"tinygo.org/x/drivers/ssd1306"
)

type Display interface {
	drivers.Displayer

	Rotate(r drivers.Rotation) Display
	Rotation() drivers.Rotation
	Clear()
}

type display struct {
	disp     ssd1306.Device
	rotation drivers.Rotation
}

func NewDisplay() Display {
	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: 2.8 * machine.MHz,
		SDA:       machine.GPIO12,
		SCL:       machine.GPIO13,
	})

	i2c0 := ssd1306.NewI2C(machine.I2C0)
	i2c0.Configure(ssd1306.Config{
		Address: 0x3C,
		Width:   128,
		Height:  64,
	})

	return &display{
		disp:     i2c0,
		rotation: drivers.Rotation0,
	}
}

func NewRotatedDisplay(r drivers.Rotation) Display {
	return NewDisplay().Rotate(r)
}

func (d *display) Size() (x, y int16) {
	return d.disp.Size()
}

func (d *display) SetPixel(x, y int16, c color.RGBA) {
	d.disp.SetPixel(x, y, c)
}

func (d *display) Display() error {
	return d.disp.Display()
}

func (d *display) Clear() {
	d.disp.ClearDisplay()
}

func (d *display) Rotation() drivers.Rotation {
	return d.rotation
}

func (d *display) Rotate(r drivers.Rotation) Display {
	newDisplay := &display{
		disp:     d.disp,
		rotation: r,
	}
	switch r {
	case drivers.Rotation0:
		return &rotattion0Display{newDisplay}
	case drivers.Rotation90:
		return &rotetion90Display{newDisplay}
	case drivers.Rotation180:
		return &rotetion180Display{newDisplay}
	case drivers.Rotation270:
		return &rotetion270Display{newDisplay}
	default:
		panic("Not supported rotation")
	}
}

type rotattion0Display struct {
	*display
}

type rotetion90Display struct {
	*display
}

func (d *rotetion90Display) Size() (x, y int16) {
	return y, x
}

func (d *rotetion90Display) SetPixel(x, y int16, c color.RGBA) {
	_, sy := d.display.Size()
	d.display.SetPixel(y, sy-x, c)
}

type rotetion180Display struct {
	*display
}

func (d *rotetion180Display) SetPixel(x, y int16, c color.RGBA) {
	sx, sy := d.display.Size()
	d.display.SetPixel(sx-x, sy-y, c)
}

type rotetion270Display struct {
	*display
}

func (d *rotetion270Display) Size() (x, y int16) {
	return y, x
}

func (d *rotetion270Display) SetPixel(x, y int16, c color.RGBA) {
	sx, _ := d.display.Size()
	d.display.SetPixel(sx-y, x, c)
}
