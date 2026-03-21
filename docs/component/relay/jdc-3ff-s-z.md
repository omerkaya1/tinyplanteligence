# Mechanical relay (Model: JQC-3FF-S-Z)

## Purpose

This module allows a low-voltage control signal (e.g., 5V from a microcontroller) to switch a higher voltage or higher current load safely and electrically isolated.

## Inventory

The board includes:

- A relay (switching element)
- A driver transistor
- A flyback diode (D1) for protection
- A 3-pin control header (typically VCC, GND, IN)
- A 3-terminal screw block labeled:
    - NC (Normally Closed)
    - COM (Common)
    - NO (Normally Open)

## Suitable if:

- The pump operates within the relay’s ratings:
    - Typically 10A @ 250V AC or 10A @ 30V DC
- The pump voltage/current is within those limits
    - You provide an appropriate external power source for the pump (relay does NOT power it)

## Important considerations:
- If it’s a DC pump, ensure:
    - Voltage ≤ ~30V DC
    - Current draw ≤ relay rating (preferably <70% for safety margin)

- If it’s an inductive load (most pumps are):
    - Add a flyback diode across the pump terminals (for DC pumps)

- The relay is mechanical, so:
    - Not ideal for very frequent switching
    - Has limited lifespan (~100k–1M cycles)