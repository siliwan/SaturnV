package main

import (
	"math"
)

func diff(stage Stage) float64 {
	return float64((stage.Full_weight - stage.Empty_weight) / stage.Burn_time)
}

func compute_step(seconds int, weight, vel, alt float64, stage Stage) (float64, float64, float64){
	diff := diff(stage)
	new_weight := weight - math.Abs(diff)
	new_velocity := step(new_weight, diff, vel, float64(stage.Exit_velocity))
	added := 1.0 //altitude(new_velocity, float64(burn_lower) / 50)
	if (seconds == 0) {
		return new_velocity, new_weight, alt + added
	}
	return first_stage_flight(seconds - 1, new_weight, new_velocity, alt + added)
}
