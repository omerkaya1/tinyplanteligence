package servo

import (
	"machine"
)

const period = 20e6 // 20ms

const (
	Right  uint32 = 1000
	Center uint32 = 1500
	Left   uint32 = 2000
)

type Servo struct {
	pwm     machine.PWM
	pin     machine.Pin
	channel uint8
}

func New(pwm machine.PWM, pin machine.Pin) (Servo, error) {
	if err := pwm.Configure(machine.PWMConfig{Period: period}); err != nil {
		return Servo{}, err
	}
	pin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	channel, err := pwm.Channel(pin)
	if err != nil {
		return Servo{}, err
	}
	return Servo{
		pwm:     pwm,
		pin:     pin,
		channel: channel,
	}, nil
}

type TurnParams struct {
	Cycle uint32
}

func (s Servo) Turn(p TurnParams) {
	value := s.pwm.Top() * p.Cycle / (period / 1000)
	s.pwm.Set(s.channel, value)
}
