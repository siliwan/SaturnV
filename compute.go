package main

import (
	"math"
	"fmt"
)

func diff(stage Stage) float64 {
	return float64((stage.Empty_weight - stage.Full_weight) / stage.Burn_time)
}

func compute_step(seconds int, weight, vel, alt float64, stage Stage) (float64, float64, float64){
	diff := diff(stage)
	new_weight := weight - math.Abs(diff)
	ideal_vel := step(new_weight, diff, vel, float64(stage.Exit_velocity))
	new_vel := vel + ((ideal_vel - vel) - 9.81)
	fmt.Println(seconds, ideal_vel, new_vel)
	if (seconds == 0) {
		return new_vel, new_weight, alt + new_vel
	}
	return compute_step(seconds - 1, new_weight, new_vel, alt + new_vel, stage)
}
