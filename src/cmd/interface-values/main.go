package main

import "fmt"

type cat interface{}

type concreteCat struct{}

func getCat() cat {
	var c1 *concreteCat
	// ... forget to assign res
	fmt.Println("c1 == nil:", c1 == nil) // HL
	return c1
}

func main() {
	c2 := getCat()
	fmt.Println("c2 == nil:", c2 == nil) // HL
}
