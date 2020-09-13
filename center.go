package main

import (
	"math"
)

func second_stage_fuel(seconds int) float64{
	return float64((empty_center - full_center) / seconds)
}

func second_stage_flight(seconds int, weight, vel float64) (float64, float64) {
	diff := second_stage_fuel(387)
	new_weight := weight - math.Abs(diff)
	new_velocity := step(weight, diff, vel, float64(exit_center))
	//fmt.Println(seconds, diff, new_velocity, new_weight)
	if (seconds == 387) {
		return new_velocity, new_weight
	}
	return second_stage_flight(seconds + 1, new_weight, new_velocity)
}
