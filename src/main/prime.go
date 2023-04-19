package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("prime numbers:")
	fmt.Println(2)

	for n := uint64(9999999999999999999); n > 0; n += 2 {
		prime := true

		r := uint64(math.Sqrt(float64(n))) + 1
		for i := uint64(3); i < r; i += 2 {
			if n%i == 0 {
				prime = false
				break
			}
		}

		if prime {
			fmt.Println(n)
		}
	}
}
