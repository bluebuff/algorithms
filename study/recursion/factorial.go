package recursion

/*
阶乘
*/
func f(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return n * f(n-1)
}
