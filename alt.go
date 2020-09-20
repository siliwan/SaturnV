package main

import (
	"math"
)

func altitude(delta_vel, seconds float64) float64{
	return delta_vel * seconds
}

func gravity(vel, time float64) float64 {
	return vel / time  - 10
}

func res(vel, alt float64) float64 {
	return 0.5 * 78.53975 * 0.75 * pluft(alt) * vel * vel
}

func pluft(alt float64) float64 {
	// Due to math rounding error set the value to zero above the karman line (100km).
	if alt >= 100000 {
		return 0
	}
	v := 1.01325 * math.Exp(-((1.247015 * alt * planet.Gravity) / 1.01325))
	return v
}
