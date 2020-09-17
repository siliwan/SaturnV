package main

import (
	"math"
)

func second_stage_fuel(seconds int) float64{
	if seconds == 0 {
		return 0
	}
	return float64((empty_center - full_center) / seconds)
}

func second_stage_flight(seconds int, weight, vel float64) (float64, float64) {
	diff := second_stage_fuel(50)
	new_weight := weight - math.Abs(diff)
	new_velocity := step(weight, diff, vel, float64(exit_center))
	if (seconds == 0) {
		return new_velocity, new_weight
	}
	return second_stage_flight(seconds - 1, new_weight, new_velocity)
}
