package searchbinary

/*
二分查找基础版(插入点)
数据a是一个有序的数组
找到则返回索引
找不到返回-(插入点+1)
问题1：为什么是i<=j 意味着区间内有未比较的元素，而不是i<j？
	i == j 意味着 i,j 它们指向的元素也会参与比较
	i < j 只意味着m指向的元素参与比较
时间复杂度 fn = (floor(log2(n) + 1)) * 5 + 4
	- 最坏情况：O(log(n))
	- 最好情况: 如果待查找元素恰好在数组中央，只需要循环一次O(1)
空间复杂度
	- 需要常数个指针i,j,m，因此额外占用的空间是O(1)
*/
func BinarySearch(a []int, target int) int {
	i, j := 0, len(a)-1
	for i <= j {
		m := (i + j) >> 1
		if target < a[m] {
			j = m - 1
		} else if a[m] < target {
			i = m + 1
		} else {
			return m
		}
	}
	return -(i + 1)
}

/*
二分查找基础版
数据a是一个有序的数组
找到则返回索引
找不到返回-1
问题1：为什么是i<=j 意味着区间内有未比较的元素，而不是i<j？
	i == j 意味着 i,j 它们指向的元素也会参与比较
	i < j 只意味着m指向的元素参与比较
时间复杂度 fn = (floor(log2(n) + 1)) * 5 + 4
	- 最坏情况：O(log(n))
	- 最好情况: 如果待查找元素恰好在数组中央，只需要循环一次O(1)
空间复杂度
	- 需要常数个指针i,j,m，因此额外占用的空间是O(1)
*/
func BinarySearchBasic(a []int, target int) int {
	i, j := 0, len(a)-1
	for i <= j {
		m := (i + j) >> 1
		if target < a[m] {
			j = m - 1
		} else if a[m] < target {
			i = m + 1
		} else {
			return m
		}
	}
	return -1
}

/*
二分查找改动版
数据a是一个有序的数组
找到则返回索引
找不到返回-1
*/
func BinarySearchAlternative(a []int, target int) int {
	i, j := 0, len(a)
	for i < j {
		m := (i + j) >> 1
		if target < a[m] {
			j = m
		} else if a[m] < target {
			i = m + 1
		} else {
			return m
		}
	}
	return -1
}

/*
二分查找平衡版
数据a是一个有序的数组
找到则返回索引
找不到返回-1
*/
func BinarySearchBalance(a []int, target int) int {
	i, j := 0, len(a)
	for 1 < j-i {
		m := (i + j) >> 1
		if target < a[m] {
			j = m
		} else {
			i = m
		}
	}
	if a[i] == target {
		return i
	}
	return -1
}

/*
二分查找LeftMost，如果有重复，返回第一个索引，如果没找到，返回插入的索引
a 查找升序数组
target 查找目标

*/
func BinarySearchLeftMost(a []int, target int) int {
	i, j := 0, len(a)-1
	for i <= j {
		m := (i + j) >> 1
		if target <= a[m] {
			j = m - 1
		} else {
			i = m + 1
		}
	}
	return i
}

/*
二分查找Right，如果有重复，返回最后一个索引，如果没找到，返回插入的索引
a 查找升序数组
target 查找目标
*/
func BinarySearchRightMost(a []int, target int) int {
	i, j := 0, len(a)-1
	for i <= j {
		m := (i + j) >> 1
		if target < a[m] {
			j = m - 1
		} else {
			i = m + 1
		}
	}
	return i - 1
}
