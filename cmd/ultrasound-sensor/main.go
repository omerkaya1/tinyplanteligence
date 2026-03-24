package main

import (
	"machine"

	"github.com/omerkaya1/tinyplanteligence/internal/driver/ultrasound"
)

func main() {
	// init periferal
	machine.InitADC()

	us := ultrasound.NewUltrasoundSensor(machine.D2, machine.D3, 100)

	var lastDistance int64
	for {
		distance := us.GetDistance()

		if distance != lastDistance {
			println(distance)
			lastDistance = distance
			continue
		}
	}
}