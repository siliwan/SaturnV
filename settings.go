package main

import (
	"fmt"
	//"encoding/json"
)


type Settings struct {
	Output Output `json:"output"`
	Profile Profile `json:"profile"`
	Planets []Planet `json:"planets"`
}

type Planet struct {
	Name string `json:"name"`
	Gravity string `json:"gravity"`
}

type Output struct {
	CSV bool `json:"csv"`
	Graph bool `json:"graph"`
	Text bool `json:"text"`
}

type Profile struct {
	CSV bool `json:"single"`
	Graph bool `json:"gravity"`
	Text bool `json:"air"`
}

func getSettings() Settings{
	content := extract("./settings.json")
	fmt.Println(content)
	var settings Settings
	return settings
}
