package main

import (
	"encoding/json"
)


type Settings struct {
	Output Output `json:"output"`
	Profile Profile `json:"profile"`
	Planets []Planet `json:"planets"`
}

type Planet struct {
	Name string `json:"name"`
	Gravity float64 `json:"gravity"`
}

type Output struct {
	CSV bool `json:"csv"`
	Graph bool `json:"graph"`
	Text bool `json:"text"`
}

type Profile struct {
	Gravity bool `json:"gravity"`
	Air bool `json:"air"`
}

func getSettings() Settings{
	content := extract("./settings.json")
	var settings Settings
	json.Unmarshal(content, &settings)
	return settings
}
