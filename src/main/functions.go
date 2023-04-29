package main

func main() {
	println(foo(1, 2))
	a := [3]int{1, 2, 3}
	println(bar(a))
	println(a[0])
	println(baz(&a))
	println(a[0])
	b := []int{1, 2, 3}
	acceptSlice(b)
	println(b[0])

	println(applyFunc(increment, 1))
	println(applyFunc(decrement, 1))
	//anonymous function
	println(applyFunc(func(x int) int { return x * x }, 1))
}

func applyFunc(f func(int) int, x int) int {
	return f(x)
}

func increment(x int) int {
	return x + 1
}

func decrement(x int) int {
	return x - 1
}

func foo(x, y int) (int, int) {
	return x + y, x - y
}

func bar(x [3]int) int {
	x[0] = 666
	return x[0]
}

func baz(x *[3]int) int {
	(*x)[0] = 666
	return (*x)[0]
}

func acceptSlice(x []int) {
	x[0] = 777
}
