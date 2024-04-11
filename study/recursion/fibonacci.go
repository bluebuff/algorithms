package recursion

/*
斐波那契第n项
*/
func fibonacci(n int) int {
	return fibonacci_recursion(n)
}

func fibonacci_recursion(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibonacci_recursion(n-1) + fibonacci_recursion(n-2)
}
