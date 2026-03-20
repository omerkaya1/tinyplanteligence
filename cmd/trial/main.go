package main

import (
	"machine"
	"time"

	"github.com/omerkaya1/tinyplanteligence/internal/driver/buzzer"
	soilmoisture "github.com/omerkaya1/tinyplanteligence/internal/driver/sensor/soil-moisture"
)

const (
	operationInterval = time.Minute
)

func main() {
	// init periferal
	machine.InitADC()

	// machine.LED.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// first sensor
	ms1 := soilmoisture.New(soilmoisture.Params{
		DryThreshold: soilmoisture.HW080Dry,
		WetThreshold: soilmoisture.HW080Wet,
		Voltage:      machine.D2,
		CtrlPin:      machine.ADC2,
	})
	// second sensor
	ms2 := soilmoisture.New(soilmoisture.Params{
		DryThreshold: soilmoisture.HW080Dry,
		WetThreshold: soilmoisture.HW080Wet,
		Voltage:      machine.D3,
		CtrlPin:      machine.ADC3,
	})

	// buzzer
	// NOTE: instead of a buzzer, a LED is used for now
	beeper := buzzer.New(machine.LED)

	// water pump
	// TODO(omerkaya1): add water pump
	// pump := pump.New(machine.D5)
	// _ = pump
	// ...

	for {
		checkSensor(ms1, beeper, 1)
		checkSensor(ms2, beeper, 2)

		time.Sleep(operationInterval)
	}
}

const (
	readSensorTimeout = 20 * time.Millisecond
)

var (
	defaultBeepParams = buzzer.BeepParams{
		Count:    1,
		Duration: 2 * time.Second,
		Delay:    time.Second,
	}
)

func checkSensor(s *soilmoisture.Sensor, b *buzzer.Buzzer, times int8) {
	s.On()
	time.Sleep(readSensorTimeout)
	val2 := s.Read()
	s.Off()

	switch val2 {
	case soilmoisture.CompletelyDry, soilmoisture.Dry, soilmoisture.SlightlyDry:
		b.Beep(defaultBeepParams)
		b.Beep(buzzer.BeepParams{Count: times, Duration: time.Millisecond * 300, Delay: time.Second})
		b.Beep(defaultBeepParams)
	}
}
