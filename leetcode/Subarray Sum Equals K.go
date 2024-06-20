package leetcode

func SubarraySum(nums []int, k int) int {
	var curSum int
	var ans int

	for left := 0; left < len(nums); {
		curSum += nums[left]
		if curSum == k {
			ans += 1
		}

		var wentRight = false
		for right := left + 1; right < len(nums); right += 1 {
			wentRight = true
			curSum += nums[right]
			if curSum == k {
				ans += 1
			}
		}

		curSum = 0
		left += 1

		if !wentRight {
			break
		}

	}

	return ans
}

func SubarraySumSmart(nums []int, k int) int {
	sum, counter := 0, 0
	ps := map[int]int{}
	ps[0] = 1

	for _, n := range nums {
		sum += n

		rem := sum - k
		if _, ok := ps[rem]; ok {
			counter += ps[rem]
		}
		ps[sum]++
	}
	return counter
}
