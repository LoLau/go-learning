package src

import (
	"fmt"
)

func main() {

/*	// define a variable
	var x int

	//define a group of variable
	var (
		a, b int
		c    float64
	)

	//define variable and assign value
	d := 1
	y := int64(1)

	//define a const variable
	const f = 64

	//define a group of variable
	const (
		e = 1
		g = 2
	)*/

	var a = 1

	f := 1

	if a < 10 {
		println("true")
	} else {
		println("false")
	}

	switch a {
		case 10:
			println()
		case 20:
			println()
		default:
			println()
	}

	for a := 0; a < 10; a ++ {
		fmt.Printf("%d ", a)
	}

	z := 1
	for z < 2 {
		z ++
		fmt.Printf("%d", z)
	}

	println()
	// numbers := [6]int {1, 2, 3, 5}
	//var numbers = []float32 {1.0, 2.0, 3.0}
	var numbers [10] int
	for i := 0; i < 10 ; i++ {
		numbers[i] = i
	}

	for i, x := range numbers {
		fmt.Printf("第%d位的值为%d\n", i, x)
	}

	//goto
	var y = 10
	LOOP:
		for y < 20 {
			if y ==20  {
				a ++
				goto LOOP
			}
			fmt.Printf("%f\n ", y)
			y ++
		}

	/* endless loop
	for {
		fmt.Printf("true")
	}*/

	println()

	println(a, f)
}
