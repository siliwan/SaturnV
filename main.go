 package main

import (
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
	exit_center = -4160
	exit_upper = -4180

	payload = 41000

	burn_lower = 161
	burn_center = 387
	burn_upper = 319

	first_weight = full_lower + full_center + full_upper + payload

)

// S1 vel_target : ~3979

func main(){
	time.Sleep(150 * time.Millisecond)
	startup()
	/*
	lower_vel, lower_weight, lower_alt := first_stage_flight(49, float64(first_weight), 0, 0)
	center_vel, center_weight, center_alt := second_stage_flight(49, lower_weight - float64(empty_lower), lower_vel, lower_alt)
	upper_vel, upper_weight, upper_alt := third_stage_flight(49, center_weight - float64(empty_center), center_vel, center_alt)
	fmt.Println(lower_vel, lower_weight, lower_alt)
	fmt.Println(center_vel, center_weight, center_alt)
	fmt.Println(upper_vel, upper_weight, upper_alt)
	fmt.Println("Payload:", upper_weight - float64(empty_upper), "Vel:", upper_vel, "Alt:", upper_alt)
*/
}

func step(m_rakete, m_gas, vel_rakete, exit_vel float64) float64 {
	return vel_rakete + (m_gas * exit_vel / m_rakete)
	//return ((m_rakete * vel_rakete + m_gas * exit_vel) / (m_rakete))
}
