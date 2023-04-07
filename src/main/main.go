package main

import "fmt"

func main() {
	ptr := new(int)
	*ptr = 4
	text := "ala"
	text = "blo"
	integer := 5
	type Workday int
	const (
		Monday Workday = iota
		Tuesday
		Wednesday
	)

	switch integer {
	case 1:
		fmt.Printf("1!")
	case 2:
		fmt.Printf("1!")
	}
	fmt.Printf("Hello, world! %s", text)
}
