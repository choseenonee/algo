package leetcode

func SubarraySum(nums []int, k int) int {
	var curSum int
	var ans int
	for left := 0; left < len(nums); {
		curSum += nums[left]
		if curSum >= k {
			if curSum == k {
				ans += 1
			}
			curSum = 0
			left += 1
		} else {
			var wentRight = false
			for right := left + 1; right < len(nums); right += 1 {
				wentRight = true
				curSum += nums[right]
				if curSum >= k {
					if curSum == k {
						ans += 1
					}
					curSum = 0
					left += 1
					break
				}
			}
			if !wentRight {
				break
			}
		}
	}

	return ans
}
