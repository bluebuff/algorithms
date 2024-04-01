package recursion

/*
冒泡排序
*/
func bubbleSort(a []int) {
	bubbleSort_recursion(a, len(a)-1)
}

func bubbleSort_recursion(a []int, j int) {
	if j == 0 {
		return
	}
	x := 0
	for i := 0; i < j; i++ {
		if a[i] > a[i+1] {
			a[i], a[i+1] = a[i+1], a[i]
			x = i
		}
	}
	bubbleSort_recursion(a, x)
}
