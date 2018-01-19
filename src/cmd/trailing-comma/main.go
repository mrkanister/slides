package main

import "fmt"

func main() {
	names := []string{
		"Java",
		"Python",
		"Go"
	}

	for i, n := range names {
		fmt.Println(i, n)
	}
}
