package main

import (
	"fmt"
	"math"
	"time"
)

var (
	full_lower = 2286000
	full_center = 491000
	full_upper = 120000

	empty_lower = 135000
	empty_center = 39000
	empty_upper = 13300

	exit_lower = -2980

	total_steps = 50

	payload = 2900

)

func main(){
	getRockets()
	time.Sleep(50 * time.Millisecond)
	weight := full_lower + full_center + full_upper
	fmt.Println(first_flight(0, float64(weight), 0))
	//fmt.Println(empty_lower + full_center + full_upper)
}

func first_to_lower(seconds int) float64{
	return float64((empty_lower - full_lower) / seconds)
}

func second_stage(seconds int) float64 {
	return float64((empty_center - full_upper) / seconds)
}

func third_stage(seconds int) float64 {
	return float64((empty_upper - payload) / seconds)
}

func first_flight(seconds int, weight, vel float64) float64 {
	diff := first_to_lower(50)
	new_weight := weight - math.Abs(diff)
	new_velocity := get_velocity(seconds, weight, vel, diff)
	//fmt.Println("Step:", seconds, "New weight:", new_weight, "Vel:", new_velocity)
	if (seconds == total_steps - 1) {
		return new_velocity
	}
	return first_flight(seconds + 1, new_weight, new_velocity)
}

func first_step(m1, m2, v2 float64) float64 {
	return ((m2) / (m1 - m2) * v2)
}

func stepn(m_rakete, m_gas, vel_rakete, exit_vel float64) float64 {
	return (m_rakete * vel_rakete + m_gas * exit_vel) / (m_rakete)
}

func get_velocity(seconds int, new_weight, vel, diff float64) float64 {
	if (seconds == 0) {
				return 0
	} else {
		return stepn(new_weight, diff, vel, float64(exit_lower))
	}
}
