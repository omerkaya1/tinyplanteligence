package ultrasound

import (
	"machine"
	"time"
)

const (
	soundSpeed = 343 // m/s

	retries uint8 = 5
)

type UltrasoundSensor struct {
	sensor  machine.Pin
	echo    machine.Pin
	timeout int64
}

func NewUltrasoundSensor(sensorPin, echoPin machine.Pin, maxDistance uint32) *UltrasoundSensor {
	sensorPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	echoPin.Configure(machine.PinConfig{Mode: machine.PinInput})
	return &UltrasoundSensor{
		sensor:  sensorPin,
		echo:    echoPin,
		timeout: int64(maxDistance * 20000 / soundSpeed),
	}
}

// GetDistance returns the distance in mm.
func (s UltrasoundSensor) GetDistance() int64 {
	return (s.pulse(retries) * 1715) / 10000 // adjusted for floating point
}

const interval = 23324

func (s UltrasoundSensor) pulse(retries uint8) int64 {
	t := time.Now()
	s.sensor.Low()
	time.Sleep(2 * time.Microsecond)
	s.sensor.High()
	time.Sleep(10 * time.Microsecond)
	s.sensor.Low()

	var i uint8
	for {
		if s.echo.Get() {
			t = time.Now()
			break
		}
		i++
		if i > retries {
			if time.Since(t).Microseconds() > interval {
				return 0
			}
			i = 0
		}
	}
	i = 0
	for {
		if !s.echo.Get() {
			return time.Since(t).Microseconds()
		}
		i++
		if i > retries {
			if time.Since(t).Microseconds() > interval {
				return 0
			}
			i = 0
		}
	}
}
