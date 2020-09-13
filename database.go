package main

import (
	"fmt"
	"os"
)

func WriteDB(step, value_1, value_2 float64) (int, error) {
	return 0, nil
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