package main

import "fmt"

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + v0*t + s0
	}
}

func main() {

	var a, v0, s0, t float64
	fmt.Print("Enter the value of acceleration: ")
	fmt.Scan(&a)
	fmt.Print("Enter the value of initial velocity: ")
	fmt.Scan(&v0)
	fmt.Print("Enter the value of initial displacement: ")
	fmt.Scan(&s0)
	fmt.Print("Enter the time to calculate displacement for: ")
	fmt.Scan(&t)
	fn := GenDisplaceFn(a, v0, s0)
	fmt.Printf("The displacement calculated is: %f\n", fn(t))
}
