package searchbinary

func searchBasic(a []int, target int) int {
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

func search2(a []int, target int) int {
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

func search3(a []int, target int) int {
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

func searchInsert(a []int, target int) int {
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

func searchRange(a []int, target int) []int {
	left := searchLeftMost(a, target)
	if left == -1 {
		return []int{-1, -1}
	}
	return []int{left, searchRightMost(a, target)}
}

func searchLeftMost(a []int, target int) int {
	i, j := 0, len(a)-1
	condidate := -1
	for i <= j {
		m := (i + j) >> 1
		if target < a[m] {
			j = m - 1
		} else if a[m] < target {
			i = m + 1
		} else {
			condidate = m
			j = m - 1
		}
	}
	return condidate
}

func searchRightMost(a []int, target int) int {
	i, j := 0, len(a)-1
	condidate := -1
	for i <= j {
		m := (i + j) >> 1
		if target < a[m] {
			j = m - 1
		} else if a[m] < target {
			i = m + 1
		} else {
			condidate = m
			i = m + 1
		}
	}
	return condidate
}
