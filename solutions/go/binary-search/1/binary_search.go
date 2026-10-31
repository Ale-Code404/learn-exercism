package binarysearch

func SearchInts(list []int, key int) int {
	length := len(list)
	if length == 0 {
		return -1
	}

	left := 0
	right := length - 1

	for left <= right {
		pivot := left + (right-left)/2

		middle := list[pivot]
		if middle == key {
			return pivot
		}

		if middle > key {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}

	return -1
}
