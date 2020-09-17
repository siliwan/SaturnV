package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)


type Rocket struct {
	Name string `json:"name"`
	Stages []Stage `json:"stages"`
	Payload_weight int `json:"payload_weight"`
}

type Stage struct {
	Name string `json:"name"`
	Full_weight int `json:"full_weight"`
	Empty_weight int `json:"empty_weight"`
	Exit_velocity int `json:"exit_velocity"`
	Burn_time int `json:"burn_time"`
}

func getRockets() []Rocket{
	files := getRocketFiles()
	rockets := []Rocket{}
	for _, f := range files {
		var rocket Rocket
		content := extract(f)
		json.Unmarshal([]byte(content), &rocket)
		rockets = append(rockets, rocket)
	}
	return rockets
}

func extract(file string) []byte {
	content, err := ioutil.ReadFile(file)
    if err != nil {
        fmt.Println(err)
    }
	return content
}
