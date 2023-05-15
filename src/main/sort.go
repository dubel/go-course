package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func main() {
	fmt.Println("Please enter a sequence of integers that would be sorted by concurrent Bubble Sort.")
	inputNumbersSlice := make([]int, 0, 10)
	var inputNumberAsString string
	for {
		fmt.Println("Note: You have input 'X' to end input process and start sorting")
		fmt.Print("Input integer and hit enter (or 'X' to start sorting):")
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
		} else {
			break
		}
	}
	fmt.Println("Slice to be sorted: ", inputNumbersSlice)
	lenOfEachSlice := len(inputNumbersSlice) / 4
	fmt.Println("Length of each slice: ", lenOfEachSlice)

	firstSlice := inputNumbersSlice[0:lenOfEachSlice]
	secondSlice := inputNumbersSlice[lenOfEachSlice:(2 * lenOfEachSlice)]
	thirdSlice := inputNumbersSlice[2*lenOfEachSlice : (3 * lenOfEachSlice)]
	fourthSlice := inputNumbersSlice[3*lenOfEachSlice : len(inputNumbersSlice)]

	var wg sync.WaitGroup
	wg.Add(4)
	go BubbleSortConcurrent(firstSlice, &wg)
	go BubbleSortConcurrent(secondSlice, &wg)
	go BubbleSortConcurrent(thirdSlice, &wg)
	go BubbleSortConcurrent(fourthSlice, &wg)
	wg.Wait()

	mergedFirstAndSecond := merge(firstSlice, secondSlice)
	mergedFirstAndSecondAndThird := merge(mergedFirstAndSecond, thirdSlice)
	mergedFirstAndSecondAndThirdAndFourth := merge(mergedFirstAndSecondAndThird, fourthSlice)

	fmt.Println("Program execution finished. Final content of sorted slice:", mergedFirstAndSecondAndThirdAndFourth)
}

// ---- sorting functions

func BubbleSortConcurrent(s []int, wg *sync.WaitGroup) {
	fmt.Println("A slice to be sorted: ", s)
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s)-1; j++ {
			if s[j] > s[j+1] {
				swap(s, j)
			}
		}
	}
	wg.Done()
}

func swap(s []int, i int) {
	if i >= len(s)-1 {
		return
	}
	temp := s[i]
	s[i] = s[i+1]
	s[i+1] = temp
}

// -- merging slices helper functions

func merge(a []int, b []int) []int {
	var final []int
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}
