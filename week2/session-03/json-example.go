package main

import (
	"encoding/json"
	"fmt"
	"hacktiv8-go/session-03/model"
	"log"
	"os"
	"path/filepath"
)

func main() {
	jsonExample()
}

func jsonExample() {
	// Get the absolute path to the JSON file
	jsonPath := filepath.Join("data", "DATA.json")

	// Read JSON file content
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		log.Fatal(err)
	}

	// var person model.Person
	var people model.People

	// Unmarshal JSON data into person struct
	if err := json.Unmarshal(data, &people); err != nil {
		log.Fatal(err)
	}

	for _, person := range people.People {
		fmt.Printf("Name: %s\n", person.Name)
		fmt.Printf("Age: %d\n", person.Age)
		fmt.Printf("Email: %s\n", person.Email)
	}
}
