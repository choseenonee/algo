from typing import List

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        m = {}
        for i in range(len(nums)):
            sec = target-nums[i]
            if sec in m.keys():
                return [i, m[sec]]
            else:
                m[nums[i]] = i
