package main

func fib(n int) int {
	if n <= 1 {
		return n
	}

	first := 0
	second := 1

	for i := 2; i <= n; i++ {
		first, second = second, first+second
	}

	return second
}
