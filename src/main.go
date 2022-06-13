package main

import (
	"fmt"
	"math/rand"
)

var rock, scissors int = 42, 55

// If they are all the same type then you can just put the type at the end
// You do have to specify a type though
func shippingCalc(x, y, z int) int {
	return x + y + z
}

func main() {
	fmt.Println("Hello World")
	fmt.Println("A random number example ", rand.Intn(100))
	fmt.Println(shippingCalc(1, 2, 3))
	fmt.Println(rock, scissors)
}
