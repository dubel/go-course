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

	mergedFirstAndSecond := append(firstSlice, secondSlice...)
	mergedFirstAndSecondAndThird := append(mergedFirstAndSecond, thirdSlice...)
	mergedFirstAndSecondAndThirdAndFourth := append(mergedFirstAndSecondAndThird, fourthSlice...)

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

func mergeTwoSortedSlices(firstSlice []int, firstSize int, secondSlice []int, secondSize int) []int {

	if firstSize == 0 {
		for k := 0; k < secondSize; k++ {
			firstSlice[k] = secondSlice[k]
		}
		return firstSlice
	}
	firstSlice = pushElementToEndOfSlice(firstSlice, firstSize)
	i := secondSize
	j := 0
	for k := 0; k < firstSize+secondSize; k++ {
		if i < firstSize+secondSize && j < secondSize {
			if firstSlice[i] < secondSlice[j] {
				firstSlice[k] = firstSlice[i]
				i++
			} else {
				firstSlice[k] = secondSlice[j]
				j++
			}
		} else if j < secondSize {
			firstSlice[k] = secondSlice[j]
			j++
		}
	}
	return firstSlice

}

func pushElementToEndOfSlice(sliceOfInt []int, m int) []int {
	length := len(sliceOfInt)
	k := length
	for i := m - 1; i >= 0; i-- {
		sliceOfInt[k-1] = sliceOfInt[i]
		k--
	}
	return sliceOfInt
}
