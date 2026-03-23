package keyboard

import "machine"

type (
	Keyboard struct {
		inputActive bool
		lastRow     int
		lastCol     int

		rows [4]machine.Pin
		cols [4]machine.Pin
	}
	Config struct {
		Rows [4]machine.Pin
		Cols [4]machine.Pin
	}
)

// keysMapping is a map of the keys on the keyboard.
var keysMapping = [4][4]string{
	{"1", "2", "3", "A"},
	{"4", "5", "6", "B"},
	{"7", "8", "9", "C"},
	{"*", "0", "#", "D"},
}

// NewKeyboard initialises the keyboard and driver, configuring the pins along the way.
func NewKeyboard(cfg Config) *Keyboard {
	for _, row := range cfg.Rows {
		row.Configure(machine.PinConfig{Mode: machine.PinOutput})
	}
	for _, col := range cfg.Cols {
		col.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	}
	return &Keyboard{
		rows: cfg.Rows,
		cols: cfg.Cols,

		inputActive: true,
		lastRow:     -1,
		lastCol:     -1,
	}
}

func (k *Keyboard) ReadKey() string {
	for i := range k.rows {
		k.rows[i].Low()
		for j := range k.cols {
			if !k.cols[j].Get() && k.inputActive {
				k.inputActive = false
				k.lastRow, k.lastCol = i, j
				return keysMapping[i][j]
			}
			if k.cols[j].Get() && !k.inputActive && k.lastRow == i && k.lastCol == j {
				k.inputActive = true
			}
		}
		k.rows[i].High()
	}
	return ""
}
