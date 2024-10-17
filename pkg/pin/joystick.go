package pin

import (
	"machine"
	"time"
)

type JoystickPressedEvent func()
type JoystickReleasedEvent func()

type Joystick struct {
	ax machine.ADC
	ay machine.ADC
	btn machine.Pin

	pressedEvents []JoystickPressedEvent
	releasedEvents []JoystickReleasedEvent
}

func NewJoystick() *Joystick {
	machine.InitADC()

	ax := machine.ADC{Pin: machine.GPIO29}
	ax.Configure(machine.ADCConfig{})
	ay := machine.ADC{Pin: machine.GPIO28}
	ay.Configure(machine.ADCConfig{})

	btn := machine.GPIO0
	btn.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	return &Joystick{
		ax: ax,
		ay: ay,
		btn: btn,
	}
}

func (j *Joystick) Position() (x, y uint16) {
	return j.ax.Get(), j.ay.Get()
}

func (j *Joystick) IsPressed() bool {
	return !j.btn.Get()
}

func (j *Joystick) AddPressedEvent(f JoystickPressedEvent) {
	j.pressedEvents = append(j.pressedEvents, f)
}

func (j *Joystick) ClearPressedEvents() {
	j.pressedEvents = nil
}

func (j *Joystick) AddReleasedEvent(f JoystickReleasedEvent) {
	j.releasedEvents = append(j.releasedEvents, f)
}

func (j *Joystick) ClearReleasedEvents() {
	j.releasedEvents = nil
}

func (j *Joystick) ThreadPressed() {
	for {
		if !j.btn.Get() {
			for _, f := range j.pressedEvents {
				f()
			}
			time.Sleep(10 * time.Millisecond)

			for !j.btn.Get() {}
			for _, f := range j.releasedEvents {
				f()
			}
			time.Sleep(10 * time.Millisecond)
		}
	}
}
