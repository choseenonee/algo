package leetcode

func LongestSubarray(nums []int) int {
	subarray := make([]int, 0, len(nums))

	var previousOne bool
	var hasZero bool
	var counter int
	var maxSoloSubArray int
	for _, val := range nums {
		if val == 1 { // && !previousOne
			previousOne = true
			counter += 1
			//} else if val == 1 && previousOne {
			//	counter += 1
		} else {
			if previousOne {
				subarray = append(subarray, counter)
			}
			hasZero = true
			maxSoloSubArray = max(maxSoloSubArray, counter)
			counter = 0
			subarray = append(subarray, 0)
			previousOne = false
		}
	}
	if previousOne {
		subarray = append(subarray, counter)
		maxSoloSubArray = max(maxSoloSubArray, counter)
	}

	//fmt.Println(subarray)

	var maxVal int
	for firstIdx := 0; firstIdx < len(subarray)-2; firstIdx++ {
		maxVal = max(maxVal, subarray[firstIdx]+subarray[firstIdx+2], maxSoloSubArray)
	}

	if len(subarray) < 3 {
		if !hasZero {
			return maxSoloSubArray - 1
		}
		return maxSoloSubArray
	}

	if !hasZero {
		return maxVal - 1
	}

	return maxVal
}
