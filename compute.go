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
	new_vel := vel + ((ideal_vel - vel) - planet.Gravity)
	// luft_force = luftwiderstand -> geteilt durch masse -> a
	luft_a := res(new_vel, alt + new_vel) / new_weight
	final_vel := vel + ((new_vel - vel) - luft_a)
	new_alt := alt + final_vel
	fmt.Println(seconds, luft_a, new_vel, final_vel)
	if (seconds == 0) {
		return final_vel, new_weight, new_alt
	}
	return compute_step(seconds - 1, new_weight, final_vel, new_alt, stage)
}
