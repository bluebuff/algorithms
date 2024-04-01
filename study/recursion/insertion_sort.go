package recursion

/*
插入排序
*/
func insertionSort(a []int) {
	insertionSort_recursion(a, 1)
}

func insertionSort_recursion(a []int, low int) {
	if low == len(a) {
		return
	}
	t := a[low]
	i := low - 1
	for ; i >= 0 && a[i] > t; i-- {
		a[i+1] = a[i]
	}
	a[i+1] = t
	insertionSort_recursion(a, low+1)
}
