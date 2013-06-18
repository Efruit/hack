package main

import (
	"encoding/json"
)

type Process struct{
	name string
	command string
	color string
}
type Project struct{
	name string
	global map[string]Process
	local map[string]Process
}

func jsonToProject(content []byte, project *Project) {
	var projMap  map[string]*json.RawMessage
	err := json.Unmarshal(content, &projMap)

	var local map[string]string
	err = json.Unmarshal(*projMap["local"], &local)
	if err != nil {
		return
	}

	var global map[string]string
	err = json.Unmarshal(*projMap["global"], &global)
	if err != nil {
		return
	}


	colors := []string{"g", "y", "b", "m", "c"}
	numColors := len(colors)
	count := 0

	project.local = make(map[string]Process, len(local))
	for name, command := range local {
		project.local[name] = Process{name, command, colors[count]}
		count++
		if(count > numColors){
			count = 0
		}
	}

	project.global = make(map[string]Process, len(global))
	for name, command := range global {
		project.global[name] = Process{name, command, colors[count]}
		if(count > numColors){
			count = 0
		}
	}
}