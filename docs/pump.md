# pump

this document describes the general approach to setting up a water pump.

## Requirements

- water pump (3.3V DC motor);
- a single-channel relay module (Model: JQC-3FF-S-Z).

## Setup

### Pump
- Connect `-` to the `GND` of the breadboard;
- Connect `+` to the `NC` of the relay module.

### Relay module
- Connect `COM` to the `+` of the breadboard;
- Connect `GND` to the `GND` of the breadboard;
- Connect `S` (Signal) of the relay module to any digital pin on the `D` of the microcontroller (e.g., `D7`).

### Controller

- Connect `GND` to the breadboard `GND`;
- Connect `VIN` to the breadboard `+` (`VIN`).