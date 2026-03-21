package main

import (
	"machine"
	"time"

	"github.com/omerkaya1/tinyplanteligence/internal/driver/pump"
)

const operationInterval = time.Minute

func main() {
	// init periferal
	machine.InitADC()

	// water pump
	wp := pump.New(machine.D7)

	for {
		wp.Pour(pump.PourParams{
			Duration: time.Second,
			Delay:    time.Second,
			Count:    1,
		})
		time.Sleep(operationInterval)
	}
}
