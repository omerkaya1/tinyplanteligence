package soilmoisture

import "machine"

// Calibrated results gathered from the HW080 Sensor.
const (
	HW080Dry = 64352 // the value for a completely dry sensor.
	HW080Wet = 25250 // the value sensor shows when it is completely submerged in water.
)

const totalCategories = 6

// MoistureLevel represents the level of moisture in the soil.
type MoistureLevel uint8

const (
	CompletelyDry = iota
	Dry
	SlightlyDry
	Moist
	VeryMoist
	Water
)

type Sensor struct {
	category     uint16
	dryThreshold uint16
	wetThreshold uint16
	adc          *machine.ADC
	voltage      machine.Pin
}

type Params struct {
	DryThreshold uint16
	WetThreshold uint16
	Voltage      machine.Pin
	CtrlPin      machine.Pin
}

func New(params Params) *Sensor {
	// derive categories
	category := (params.DryThreshold - params.WetThreshold) / totalCategories
	s := Sensor{
		category:     category,
		dryThreshold: params.DryThreshold,
		wetThreshold: params.WetThreshold,
		voltage:      params.Voltage,
		adc:          &machine.ADC{Pin: params.CtrlPin},
	}
	s.voltage.Configure(machine.PinConfig{Mode: machine.PinOutput})
	s.adc.Configure(machine.ADCConfig{})
	return &s
}

func (s *Sensor) On() {
	s.voltage.High()
}

func (s *Sensor) Off() {
	s.voltage.Low()
}

func (s *Sensor) Read() MoistureLevel {
	val := s.adc.Get()

	println(val)

	switch {
	case val >= s.dryThreshold:
		return CompletelyDry
	case val >= s.dryThreshold-s.category:
		return Dry
	case val >= s.dryThreshold-s.category*2:
		return SlightlyDry
	case val >= s.dryThreshold-s.category*3:
		return Moist
	case val >= s.dryThreshold-s.category*4:
		return VeryMoist
	default:
		return Water
	}
}
