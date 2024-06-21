package leetcode

func MoveZeroes(nums []int) {
	for left := 0; left < len(nums)-1; {
		if nums[left] == 0 {
			for right := left + 1; right < len(nums); right++ {
				if nums[right] != 0 {
					nums[left] = nums[right]
					nums[right] = 0
					//left++
					break
				}
			}
			left++
		} else {
			left++
		}
	}
}
