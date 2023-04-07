package main

import "fmt"

func main() {
	var inputNumber float32
	for {
		fmt.Println("Please input float value to be truncated or zero to end the program execution:")
		tokensNumber, err := fmt.Scan(&inputNumber)
		var truncatedInt = int(inputNumber)
		if truncatedInt == 0 {
			break
		}
		if err == nil && tokensNumber == 1 {
			fmt.Printf("truncated value is: %d\n", truncatedInt)
		} else {
			break
		}
	}
}
