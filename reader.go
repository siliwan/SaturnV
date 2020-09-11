package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func getRocketFiles() []string {
	rockets := []string{}
	files, err := ioutil.ReadDir("./")
    if err != nil {
        fmt.Println(err)
    }
    for _, f := range files {
		if filepath.Ext(f.Name()) == ".json" {
			rockets = append(rockets, f.Name())
		}
    }
	return rockets
}
