package main

import (
	"machine"
	"time"

	"github.com/omerkaya1/tinyplanteligence/internal/driver/buzzer"
	"github.com/omerkaya1/tinyplanteligence/internal/driver/servo"
)

func main() {
	// init periferal
	machine.InitADC()

	// there's no way we can indicate error should something go wrong,
	// unless we use the debugger or monitor the serial port.
	// to properly indicate the error, we are going to use the onboard LED.

	s, err := servo.New(machine.Timer1, machine.D10)
	if err != nil {
		b := buzzer.New(machine.LED)
		beep := buzzer.BeepParams{
			Count:    5,
			Duration: time.Millisecond * 300,
			Delay:    time.Millisecond * 500,
		}
		for {
			b.Beep(beep)
			time.Sleep(time.Second)
		}
	}

	for ; ; time.Sleep(time.Second) {
		s.Turn(servo.TurnParams{Cycle: servo.Left})

		time.Sleep(time.Second)

		s.Turn(servo.TurnParams{Cycle: servo.Center})

		time.Sleep(time.Second)

		s.Turn(servo.TurnParams{Cycle: servo.Right})
	}
}
