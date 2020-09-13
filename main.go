 package main

import (
	"fmt"
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

)

func main(){
	getRockets()
	time.Sleep(50 * time.Millisecond)
	first_weight := full_lower + full_center + full_upper + payload
	lower_vel, lower_weight := first_stage_flight(0, float64(first_weight), 0, 0)
	center_vel, center_weight := second_stage_flight(0, lower_weight - float64(empty_lower), lower_vel)
	upper_vel, upper_weight := third_stage_flight(0, center_weight - float64(empty_center), center_vel)
	fmt.Println(lower_vel, lower_weight)
	fmt.Println(center_vel, center_weight)
	fmt.Println(upper_vel, upper_weight)
}

func step(m_rakete, m_gas, vel_rakete, exit_vel float64) float64 {
	return (m_rakete * vel_rakete + m_gas * exit_vel) / (m_rakete)
}
