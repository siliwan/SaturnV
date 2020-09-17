package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"bufio"
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
		num := strconv.Itoa(i + 1)
		fmt.Println(num + ")", r.Name)
	}
	fmt.Println("")
	fmt.Println("x) to return to menu.")
	fmt.Println("q) quit rocket simulator.")
	fmt.Println("")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	switch text {
	case "x\n":
		startup()
	case "q\n":
		os.Exit(0)
	}

}
