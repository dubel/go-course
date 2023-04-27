package main

func main() {
	println(foo(1, 2))
	a := [3]int{1, 2, 3}
	println(bar(a))
	println(a[0])
	println(baz(&a))
	println(a[0])
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
