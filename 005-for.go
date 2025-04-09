package main

import "fmt"

func main() {

	// First Method
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Second Method
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// Third method
	for i := range 3 {
		fmt.Println(i)
	}

	// Infinite Loop
	for {
		fmt.Println("Loop...")
		break
	}

	// Odd / Even Loop
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
