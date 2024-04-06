package recursion

import "fmt"

func element(i, j int) int {
	if j == 0 || i == j {
		return 1
	}
	return element(i-1, j-1) + element(i-1, j)
}

func printParscalTriangle(n int) {
	for i := 0; i < n; i++ {
		printSpace(n, i)
		for j := 0; j <= i; j++ {
			fmt.Printf("%-4d", element(i, j))
		}
		fmt.Println()
	}
}

func printSpace(n, i int) {
	for i := (n - i - 1) * 2; i > 0; i-- {
		fmt.Print(" ")
	}
}
