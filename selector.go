package main

import (
	"fmt"
	"bufio"
	"os"
	"os/exec"
	"time"
)

// Global vars for state setting. Bad behaviour, don't repliacte.
var (
	menu map[int]string
	rocket Rocket
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
	fmt.Println(1, menu[1])
	fmt.Println(2, menu[2])
	fmt.Println(3, menu[3])

	fmt.Println("")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	switch text {
	case "1\n":
		time.Sleep(75 * time.Millisecond)
		rockets_menu()
	case "2\n":
		time.Sleep(75 * time.Millisecond)
		fmt.Println(getSettings())
	case "3\n":
		time.Sleep(75 * time.Millisecond)
		rockets_menu()
	}
}
