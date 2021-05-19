package main

import (
	"fmt"

	obj "github.com/ELares/obj/pkg"
)

func main() {
	ip := obj.NewParser()

	// Object with random values
	randomArray := []interface{}{"123", 54, struct{}{}, make(chan int), nil, "-488", -238, 44.342342}

	// Array to hold all the conversions
	intArray := make([]int, 0)

	for _, o := range randomArray {
		n := ip.ToInt(o, -1)
		intArray = append(intArray, n)
	}

	// Print results
	fmt.Printf("%v\n", intArray)
}
