package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var inputNumberAsString string
	slice := make([]int, 0, 3)
	for {
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
			slice = append(slice, inputNumberAsInt)
			sort.Ints(slice)
			fmt.Println("Slice after adding new value:", slice)
		} else {
			break
		}
	}
}
