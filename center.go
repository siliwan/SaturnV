package main

import (
	"math"
)

func second_stage_fuel(seconds int) float64{
	return float64((empty_center - full_center) / seconds)
}

func second_stage_flight(seconds int, weight, vel, alt float64) (float64, float64, float64) {
	diff := second_stage_fuel(50)
	new_weight := weight - math.Abs(diff)
	new_velocity := step(new_weight, diff, vel, float64(exit_center))
	added := altitude(new_velocity, float64(burn_center) / 50)
	if (seconds == 0) {
		return new_velocity, new_weight, alt + added
	}
	return second_stage_flight(seconds - 1, new_weight, new_velocity, alt + added)
}
