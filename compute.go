package main

import (
	"math"
)

func diff(stage Stage) float64 {
	return float64((stage.Empty_weight - stage.Full_weight) / stage.Burn_time)
}

func compute_step(seconds int, weight, vel, alt float64, stage Stage) (float64, float64, float64){
	diff := diff(stage)
	new_weight := weight - math.Abs(diff)
	new_velocity := step(new_weight, diff, vel, float64(stage.Exit_velocity))
	added := altitude(new_velocity, float64(stage.Exit_velocity) / float64(stage.Burn_time))
	if (seconds == 0) {
		return new_velocity, new_weight, alt + added
	}
	return compute_step(seconds - 1, new_weight, new_velocity, alt + added, stage)
}
