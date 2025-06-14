package leetcode

func Rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 3 {
		return max(nums[0]+nums[2], nums[1])
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	} else if len(nums) == 4 {
		return max(nums[0]+max(nums[2], nums[3]), nums[1]+nums[3])
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = nums[1]
	for idx, val := range nums {
		if idx == 2 {
			dp[idx] = val + dp[0]
			continue
		}
		if idx == 3 {
			dp[idx] = val + max(dp[1], dp[0])
			continue
		}
		if idx < 3 {
			continue
		}
		dp[idx] = val + max(dp[idx-2], dp[idx-3])
	}

	// slices.Max(dp)
	return max(dp[len(dp)-1], dp[len(dp)-2])
}
