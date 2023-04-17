package main

func main() {
	var array [5]int
	array[0] = 1

	var arrayLiteral [5]int = [5]int{1, 2, 3, 4}
	arrayLiteral[0] = 1

	for i, v := range arrayLiteral {
		println(i, v)
	}
}
