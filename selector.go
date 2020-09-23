package main

import (
	"fmt"
	"bufio"
	"os"
	"os/exec"
	"time"
	"strings"
)

// Global vars for state setting. Bad behaviour, don't repliacte.
var (
	menu map[int]string
	rocket Rocket
	planet Planet
	settings Settings
)

func startup() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()

	menu = make(map[int]string)
	menu[1] = "Rockets"
	menu[2] = "Settings"
	menu[3] = "Execute"

	fmt.Println("- - - - - - - - - - - - - - - - -")
	fmt.Println("Rocket simulator loaded successfully")
	fmt.Println("- - - - - - - - - - - - - - - - -")
	fmt.Println("")
	fmt.Println("Select menu:")
	fmt.Printf("%d) %s\n", 1, menu[1])
	fmt.Printf("%d) %s\n", 2, menu[2])
	fmt.Printf("%d) %s\n", 3, menu[3])
	fmt.Println("\n")
	fmt.Println("q) Quit rocket simulator.")

	fmt.Println("")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.ReplaceAll(text, "\n", "")


	switch text {
	case "1":
		time.Sleep(75 * time.Millisecond)
		rockets_menu()
	case "2":
		time.Sleep(75 * time.Millisecond)
		settings_menu()
	case "3":
		time.Sleep(75 * time.Millisecond)
		execute_menu()
	case "q":
		os.Exit(0)
	}
}
