package main

import (
	"encoding/json"
)


type Settings struct {
	Profile Profile `json:"profile"`
	Planets []Planet `json:"planets"`
}

type Planet struct {
	Name string `json:"name"`
	Gravity float64 `json:"gravity"`
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
