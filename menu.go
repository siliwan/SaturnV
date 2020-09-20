package main

import (
	"fmt"
	"os"
	"os/exec"
	"bufio"
	"strings"
	"strconv"
	"time"
)

func rockets_menu() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println("- - - - - - - - -")
	fmt.Println("Available rockets:")
	fmt.Println("- - - - - - - - -")
	fmt.Println("")

	rockets := getRockets()
	for i, r := range rockets {
		fmt.Printf("%d) %s \n", i + 1, r.Name)
	}
	fmt.Println("\n")
	fmt.Println("x) to return to menu.")
	fmt.Println("q) quit rocket simulator.")
	fmt.Println("")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	text = strings.ReplaceAll(text, "\n", "")

	num, err := strconv.Atoi(text)
	if err == nil {
		rocket = rockets[num - 1]
		time.Sleep(250 * time.Millisecond)
		startup()
	}

	switch text {
	case "x":
		startup()
	case "q":
		os.Exit(0)
	default:
		rockets_menu()
	}
}

func settings_menu() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()

	settings := getSettings()

	fmt.Println("- - - - - - - -")
	fmt.Println("Current settings:")
	fmt.Println("- - - - - - - -")
	fmt.Println("Loaded from 'settings.json'.")
	fmt.Println("")

	fmt.Println("Flight profile:")

	fmt.Println("Gravity enabled:", settings.Profile.Gravity)
	fmt.Println("Air resistance enabled:", settings.Profile.Air)

	fmt.Println("")

	fmt.Println("Planets:")
	fmt.Println("(Select one)")
	for i, p := range settings.Planets {
		fmt.Printf("%d) %s, Gravity: %fm/s \n", i + 1, p.Name, p.Gravity)
	}

	fmt.Println("\n")
	fmt.Println("x) to return to menu.")
	fmt.Println("q) quit rocket simulator.")
	fmt.Println("")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	text = strings.ReplaceAll(text, "\n", "")
	num, err := strconv.Atoi(text)
	if err == nil {
		planet = settings.Planets[num - 1]
		time.Sleep(250 * time.Millisecond)
		startup()
	}

	switch text {
	case "x":
		startup()
	case "q":
		os.Exit(0)
	}
}

func execute_menu() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Println("- - - - - - - - - - - - -")
	fmt.Println("Executing rocket simulation.")
	fmt.Println("- - - - - - - - - - - - -")

	if rocket.Name == "" || planet.Name == "" {
		missing_element()
	}

	fmt.Println("\n")
	fmt.Println("Configuration:")
	fmt.Println("Rocket:", rocket.Name)
	fmt.Println("Planet:", planet.Name)

	fmt.Println("\n")
	fmt.Println("Confirm configuration? (y/n)")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	text = strings.ReplaceAll(text, "\n", "")

	switch text {
	case "y":
		launch()
	case "n":
		startup()
	}

}

func missing_element() {
	fmt.Println("ERROR! Execution failed.")
	fmt.Println("You need to select rocket & planet before launch!")
	fmt.Println("\n")
	for i := 0; i < 35; i++ {
		fmt.Print(".")
		time.Sleep(100 * time.Millisecond)
	}
	startup()
}
