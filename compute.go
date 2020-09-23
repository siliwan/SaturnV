package main

import (
	"math"
)

func diff(stage Stage) float64 {
	return float64((stage.Empty_weight - stage.Full_weight) / rocket.Steps)
}

func compute_step(seconds int, weight, vel, alt float64, stage Stage) (float64, float64, float64){
	diff := diff(stage)
	new_weight := weight - math.Abs(diff)
	delta_t := float64(stage.Burn_time / rocket.Steps)

	ideal_vel := step(new_weight, diff, vel, float64(stage.Exit_velocity))

	luft_a := res(ideal_vel, alt + (ideal_vel * delta_t)) / new_weight

	a := ((ideal_vel - vel) / delta_t) - planet.Gravity - luft_a

	final_vel := vel + a * delta_t

	new_alt := alt + final_vel * delta_t

	keys := []string{"Velocity", "Altitude", "Weight", "Acceleration"}
	values := []float64{final_vel, new_alt, new_weight, a}

	store(seconds, keys, values)

	if (seconds == 0) {
		key := []string{stage.Name}
		value := []float64{float64(stage.Burn_time)}
		store(-1, key, value)
		return final_vel, new_weight, new_alt
	}
	return compute_step(seconds - 1, new_weight, final_vel, new_alt, stage)
}
