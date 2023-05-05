package main

import (
	"fmt"
	"strconv"
)

func Swap(numList []int, index int) {
	temp := (numList)[index]
	(numList)[index] = (numList)[index+1]
	(numList)[index+1] = temp
}

func BubbleSort(numbers []int) {

	for i := 0; i < len(numbers); i++ {

		for j := 1; j < len(numbers); j++ {
			// need to use parentheses to specify what is dereferenced.
			if (numbers)[j-1] > (numbers)[j] {
				Swap(numbers, j-1)
			}
		}
	}
	fmt.Println(numbers)
}

func main() {
	var userInput string
	userNumbers := make([]int, 0, 10)
	for len(userNumbers) < 10 {
		fmt.Println("Enter a number, q to quit")
		fmt.Scan(&userInput)
		if userInput == "q" {
			break
		} else {
			// Convert string to int and only add it to the slice if successful
			number, err := strconv.Atoi(userInput)
			if err == nil {
				userNumbers = append(userNumbers, number)
			}

		}
	}
	BubbleSort(userNumbers)
}
