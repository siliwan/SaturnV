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

	for i := 1; i < 50; i++ {
		fmt.Print(".")
		time.Sleep(100 * time.Millisecond)
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

}

// length key must equal length values
func store(step int, keys []string, values []float64){
	for i := range keys {
		line := fmt.Sprintf("%d,%s,%f\n", step, keys[i], values[i])
		write_to_db(line)
	}
}
