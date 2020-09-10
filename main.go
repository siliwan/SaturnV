package main

import (
	"fmt"
	"math"
	"time"
	"os"
)

var (
	full_lower = 2286000
	full_center = 491000
	full_upper = 120000

	empty_lower = 135000
	empty_center = 39000
	empty_upper = 13300

	exit_lower = -2980
)

func main(){
	time.Sleep(50 * time.Millisecond)
	weight := full_lower + full_center + full_upper
	fmt.Println(first_flight(0, float64(weight), 0))
	fmt.Println(empty_lower + full_center + full_upper)
}

func first_to_lower(seconds int) float64{
	return float64((empty_lower - full_lower) / seconds)
}

func write_to_db(seconds, weight, vel float64) {
	f, err := os.OpenFile("graph.csv", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("FATAL FILE ERROR!")
	}
	s_s := fmt.Sprintf("%f", seconds)
	w_s := fmt.Sprintf("%f", weight)
	v_s := fmt.Sprintf("%f", vel)
	str := s_s + "," + w_s + "," + v_s + "\n"
	if _, err := f.WriteString(str); err != nil {
		fmt.Println("FATAL FILE ERROR!")
	}
	if err := f.Close(); err != nil {
		fmt.Println("FATAL FILE ERROR!")
    }
}

func first_flight(seconds, weight, vel float64) float64 {
	diff := first_to_lower(50)
	new_weight := weight - math.Abs(diff)
	new_velocity := get_velocity(seconds, weight, vel, diff)
	fmt.Println("Step:", seconds, "New weight:", new_weight, "Vel:", new_velocity)
	write_to_db(seconds, weight, new_velocity)
	if (seconds == 49) {
		return new_velocity
	} else {
		first_flight(seconds + 1, new_weight, new_velocity)
	}
	return 0
}

func first_step(m1, m2, v2 float64) float64 {
	return ((m2) / (m1 - m2) * v2)
}

func stepn(m_rakete, m_gas, vel_rakete float64) float64 {
	return (m_rakete * vel_rakete + m_gas * float64(exit_lower)) / (m_rakete)
}

func get_velocity(seconds, new_weight, vel, diff float64) float64 {
	if (seconds == 0) {
		/*		return 0
	} else if (seconds == 1) {*/
		return first_step(new_weight, diff, float64(exit_lower))
	} else {
		return stepn(new_weight, diff, vel)
	}
	return float64(0)
}
