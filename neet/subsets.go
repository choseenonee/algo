package neet

//func subsets(nums []int) [][]int {
//	n := int(math.Pow(2.0, float64(len(nums))))
//
//	ans := make([][]int, 0, n)
//
//	for i := 1; i <= n; i++ {
//		subset := make([]int, 0, bits.OnesCount(uint(i)))
//		for idx, val := range nums {
//			indexInBitMask := int(math.Pow(2.0, float64(idx)))
//			if indexInBitMask&i != 0 {
//				subset = append(subset, val)
//			}
//		}
//		ans = append(ans, subset)
//	}
//
//	return ans
//}

// prettier solution, backtracking algo
func subsets(nums []int) [][]int {
	var result [][]int
	var curr []int

	var explore func(int)
	explore = func(idx int) {
		if idx == len(nums) {
			subset := make([]int, 0, len(curr))
			copy(subset, curr)
			result = append(result, subset)
			return
		}

		curr = append(curr, nums[idx])
		explore(idx + 1)
		curr = curr[:len(curr)-1]

		explore(idx + 1)
	}

	explore(0)
	return result
}
