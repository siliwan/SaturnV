package main

import (
	"math"
)

// Saturn V r = 5 -> A ~ 78.5
// Saturn V estim. C_w value = 0.52

func res(vel, alt float64) float64 {
	return 0.5 * 78.5 * 0.52 * density(alt) * math.Pow(vel, 2)
}

func density(alt float64) float64 {
	//Schweizer formel https://wind-data.ch/tools/luftdichte.php
	return 1.247015 * math.Exp(-0.000104 * alt)
}
