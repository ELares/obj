# obj
A repo to handle common tasks related to `interface{}` in go.

# Caster
## `int`
* Array with random items
* Convert all the items to `int`

```go
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
	intArray := make([]int, 0)

	for _, o := range randomArray {
		n := c.ToInt(o, -1)
		intArray = append(intArray, n)
	}

	// Print results
	fmt.Printf("%v\n", intArray)
}

```
### Output
```console
[123 54 -1 -1 -1 -488 -238 44]
```


---
## `int32`
* Array with random items
* Convert all the items to `int32`

```go
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
	intArray := make([]int32, 0)

	for _, o := range randomArray {
		n := c.ToInt32(o, -1)
		intArray = append(intArray, n)
	}

	// Print results
	fmt.Printf("%v\n", intArray)
}
```
### Output
```console
[123 54 -1 -1 -1 -488 -238 44]
```