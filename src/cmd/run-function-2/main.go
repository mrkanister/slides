package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	file, err := os.Open("some-file")
	if err != nil {
		return err
	}
	defer file.Close()

	var data []int
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return fmt.Errorf("decode data: %v", err)
	}

	// ...
	return nil
}
