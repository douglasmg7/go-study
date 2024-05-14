package main

import (
	"encoding/json"
	"os"
	"strings"
)

func main() {
	// File name
	file_name := "File 1"
	file_name = strings.ToLower(strings.ReplaceAll(file_name, " ", "_")) + ".json"

	// Data
	person := struct {
		Name string `json:"nome"`
		Age  int    `json:"idade"`
	}{
		"Maria",
		48,
	}

	// json, err := json.Marshal(person)
	json, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		panic(`Error marshaling struct.`)
	}

	err = os.WriteFile(file_name, json, 0644)
	if err != nil {
		panic(`Error writing to file.`)
	}

}
