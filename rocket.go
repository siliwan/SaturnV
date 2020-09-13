package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type Rocket struct {
	Name string
}

func getRockets(){
	files := getRocketFiles()
	for _, f := range files {
		var rocket Rocket
		content := extract(f)
		json.Unmarshal([]byte(content), &rocket)
		fmt.Println(rocket.Name)
	}
}

func extract(file string) string {
	content, err := ioutil.ReadFile(file)
    if err != nil {
        fmt.Println(err)
    }
	return string(content)
}
