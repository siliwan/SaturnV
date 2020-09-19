package main

import (
	"fmt"
	"bufio"
	"os"
	"os/exec"
	"time"
)

func launch() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()

	//Debug toggle for countdown
	if false {
		fmt.Println("- - - - - - - -")
		fmt.Println("Rocket launch...")
		fmt.Println("- - - - - - - -")

		fmt.Println("\n")
		time.Sleep(1000 * time.Millisecond)

		for i := 10; i >= 0; i-- {
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()

			number := fmt.Sprintf("%d", i)

			screen := exec.Command("figlet", number)
			screen.Stdout = os.Stdout
			screen.Run()

			time.Sleep(1000 * time.Millisecond)
		}

		clear := exec.Command("clear")
		clear.Stdout = os.Stdout
		clear.Run()
	}

	screen := exec.Command("figlet", "LIFTOFF!")
	screen.Stdout = os.Stdout
	screen.Run()

	fmt.Println("Flight commencing, please stand by.")

	for i := 1; i < 30; i++ {
		fmt.Print(".")
		time.Sleep(75 * time.Millisecond)
	}

	final := exec.Command("clear")
	final.Stdout = os.Stdout
	final.Run()

	compute()

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	os.Exit(0)
}

func compute() {
	stages := rocket.Stages
	liftoff_weight := rocket.Payload_weight
	total_steps := 0
	for x := range stages {
		liftoff_weight += stages[x].Full_weight
		total_steps += stages[x].Burn_time
	}

	var current_velocity, current_altitude, current_weight float64
	current_velocity, current_altitude = 0, 0
	current_weight = float64(liftoff_weight)
	/*
	for x := range stages {
		current_velocity, current_weight, current_altitude = compute_stage(stages[x], current_velocity, current_altitude, current_weight)
		current_weight = current_weight - float64(stages[x].Empty_weight)
		//fmt.Println(current_velocity, current_weight, current_altitude)
	}
*/
	current_velocity, current_weight, current_altitude = compute_stage(stages[0], current_velocity, current_altitude, current_weight)
	result.Steps = total_steps
	result.Velocity = current_velocity
	result.Weight = current_weight
	result.Altitude = current_altitude
	results()
}

// End of run: Stage gets dropped, subtract weight.
func compute_stage(stage Stage, velocity, altitude, weight float64) (float64, float64, float64){
	return compute_step(stage.Burn_time - 1, weight, velocity, altitude, stage)
}

// length key must equal length values
func store(step int, keys []string, values []float64){
	for i := range keys {
		line := fmt.Sprintf("%d,%s,%f\n", step, keys[i], values[i])
		write_to_db(line)
	}
}
