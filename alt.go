package main

func altitude(delta_vel, seconds float64) float64{
	return delta_vel * seconds
}

func acc(vel, time float64) float64{
	return vel / time
}

func gravity(vel, time float64) float64 {
	return vel / time - 1
}
