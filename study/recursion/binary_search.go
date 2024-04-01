package recursion

/*
二分查找法（递归版）
*/
func binarySearch(a []int, target int) int {
	return binarySearch_recursion(a, target, 0, len(a)-1)
}

func binarySearch_recursion(a []int, target, i, j int) int {
	if i > j {
		return -1
	}
	m := (i + j) >> 1
	if target < a[m] {
		return binarySearch_recursion(a, target, i, m-1)
	} else if a[m] < target {
		return binarySearch_recursion(a, target, m+1, j)
	} else {
		return m
	}
}
