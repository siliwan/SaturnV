package main

import (
	"fmt"
	"os"
)

func WriteDB(step, value_1, value_2 float64) (int, error) {
	return 0, nil
}

func write_to_db(str string) {
	f, err := os.OpenFile("output/flight.csv", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("FATAL FILE ERROR!")
	}
	if _, err := f.WriteString(str); err != nil {
		fmt.Println("FATAL FILE ERROR!")
	}
	if err := f.Close(); err != nil {
		fmt.Println("FATAL FILE ERROR!")
    }
}
