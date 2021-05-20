package main

import (
	"fmt"

	"github.com/ELares/obj/pkg/caster"
)

func main() {
	c := caster.NewCaster()

	// Object with random values
	randomArray := []interface{}{"123", 54, struct{}{}, make(chan int), nil, "-488", -238, 44.342342}

	// Array to hold all the conversions
	intArray := make([]int16, 0)

	for _, o := range randomArray {
		n := c.ToInt16(o, -1)
		intArray = append(intArray, n)
	}

	// Print results
	fmt.Printf("%v\n", intArray)
}
