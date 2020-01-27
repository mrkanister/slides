package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	file, err := os.Open("some-file")
	if err != nil {
		log.Fatal()
	}
	defer file.Close() // HL1

	var data []int
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		log.Fatal(err) // HL1
	}

	// ...
}
