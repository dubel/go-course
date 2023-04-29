package main

import (
	"fmt"
	"strconv"
	"strings"
)

func BubbleSort(s []int) {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s)-1; j++ {
			if s[j] > s[j+1] {
				Swap(s, j)
			}
		}
	}
}

func Swap(s []int, i int) {
	if i >= len(s)-1 {
		return
	}
	temp := s[i]
	s[i] = s[i+1]
	s[i+1] = temp
}

func main() {
	fmt.Println("Please enter a sequence of up to 10 integers that would be sorted by Bubble Sort.")
	inputNumbersSlice := make([]int, 0, 10)
	var inputNumberAsString string
	for i := 0; i < 10; i++ {
		fmt.Println("Please input integer value or 'X' to end the program execution:")
		tokensNumber, err := fmt.Scan(&inputNumberAsString)
		if strings.ToUpper(inputNumberAsString) == "X" {
			break
		}
		inputNumberAsInt, err := strconv.Atoi(inputNumberAsString)
		if err != nil {
			fmt.Println("Error during conversion")
			continue
		}
		if err == nil && tokensNumber == 1 {
			inputNumbersSlice = append(inputNumbersSlice, inputNumberAsInt)
			BubbleSort(inputNumbersSlice)
			fmt.Println("Slice sorted after adding new value:", inputNumbersSlice)
		} else {
			break
		}
	}
	fmt.Println("Program execution finished. Final content of sorted slice:", inputNumbersSlice)
}
