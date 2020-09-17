package main

import (
	"math"
)

func third_stage_fuel(seconds int) float64{
	if seconds == 0 {
		return 0
	}
	return float64((empty_upper - full_upper) / seconds)
}

func third_stage_flight(seconds int, weight, vel float64) (float64, float64) {
	diff := third_stage_fuel(318)
	new_weight := weight - math.Abs(diff)
	new_velocity := step(weight, diff, vel, float64(exit_upper))
	if (seconds == 0) {
		return new_velocity, new_weight
	}
	return third_stage_flight(seconds - 1, new_weight, new_velocity)
}