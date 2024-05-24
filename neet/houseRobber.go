package neet

import "slices"

func Rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 3 {
		return max(nums[0]+nums[2], nums[1])
	}

	base := []int{0, 1}

	var answers []int
	for _, curr := range base {
		ans := nums[curr]
		for i := curr; i <= len(nums)-3; i++ {
			if curr == len(nums)-3 {
				ans += nums[len(nums)-1]
				break
			}
			if nums[curr+2] >= nums[curr+3] {
				curr += 2
				ans += nums[curr]
			} else {
				curr += 3
				ans += nums[curr]
			}
		}
		answers = append(answers, ans)
	}

	return slices.Max(answers)
}
