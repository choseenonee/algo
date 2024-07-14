package other

func BinarySearch(arr []int, val int) (int, bool) {
	left := 0
	right := len(arr) - 1
	for left != right {
		curIdx := (right + left) / 2
		if arr[curIdx] == val {
			return curIdx, true
		} else {
			if arr[curIdx] > val {
				right = curIdx
			} else {
				left = curIdx + 1
			}
		}
	}

	return -1, false
}
