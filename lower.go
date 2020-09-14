package main

import (
	"math"
)

func first_stage_fuel(seconds int) float64{
	if seconds == 0 {
		return 0
	}
	return float64((empty_lower - full_lower) / seconds)
}

func first_stage_flight(seconds int, weight, vel, alt float64) (float64, float64) {
	diff := first_stage_fuel(161)
	new_weight := weight - math.Abs(diff)
	new_velocity := step(weight, diff, vel, float64(exit_lower))
	if (seconds == 0) {
		return new_velocity, new_weight
	}
	return first_stage_flight(seconds - 1, new_weight, new_velocity, alt)
}
