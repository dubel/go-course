package main

func main() {
	//find primes below 100 and print them

	for i := 2; i < 100; i++ {
		isPrime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
			}
			if isPrime {
				println(i)
			}

		}
	}
}
