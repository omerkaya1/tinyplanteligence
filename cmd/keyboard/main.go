package main

import (
	"machine"

	"github.com/omerkaya1/tinyplanteligence/internal/driver/keyboard"
)

func main() {
	// init periferal
	machine.InitADC()

	kb := keyboard.NewKeyboard(keyboard.Config{
		Rows: [4]machine.Pin{machine.D9, machine.D8, machine.D7, machine.D6},
		Cols: [4]machine.Pin{machine.D5, machine.D4, machine.D3, machine.D2},
	})

	for {
		if key := kb.ReadKey(); key != "" {
			println(key)
		}
	}
}
