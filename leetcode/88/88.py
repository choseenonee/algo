from typing import List
import os

class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        p1 = 0
        p2 = 0
        cd = []
        while p1 < m+n:
            if len(cd) > 0:
                if p2 < n and cd[0] <= nums2[p2]:
                    if p1 < m:
                        cd.append(nums1[p1])
                    nums1[p1] = cd.pop(0)
                    p1 += 1
                elif p2 >= n:
                    if p1 < m:
                        cd.append(nums1[p1])
                    nums1[p1] = cd.pop(0)
                    p1 += 1
                elif p1 < m and p2 < n and nums2[p2] <= nums1[p1]:
                    if p1 < m:
                        cd.append(nums1[p1])
                    nums1[p1] = nums2[p2]
                    p1 += 1
                    p2 += 1
                elif p1 >= m and p2 < n:
                    nums1[p1] = nums2[p2]
                    p1 += 1
                    p2 += 1    
                else:
                    p1 += 1
            else:
                if p1 < m and p2 < n and nums2[p2] <= nums1[p1]:
                    if p1 < m:
                        cd.append(nums1[p1])
                    nums1[p1] = nums2[p2]
                    p1 += 1
                    p2 += 1
                elif p1 >= m and p2 < n:
                    nums1[p1] = nums2[p2]
                    p1 += 1
                    p2 += 1
                else:
                    p1 += 1


test_cases = [
    {
    "input": {
      "nums1": [1, 2, 3, 0, 0, 0],
      "m": 3,
      "nums2": [2, 5, 6],
      "n": 3
    },
    "expected": [1, 2, 2, 3, 5, 6]
  },
  {
    "input": {
      "nums1": [1, 2, 3, 0, 0, 0],
      "m": 3,
      "nums2": [4, 5, 6],
      "n": 3
    },
    "expected": [1, 2, 3, 4, 5, 6]
  },
  {
    "input": {
      "nums1": [1],
      "m": 1,
      "nums2": [],
      "n": 0
    },
    "expected": [1]
  },
  {
    "input": {
      "nums1": [0],
      "m": 0,
      "nums2": [1],
      "n": 1
    },
    "expected": [1]
  },
    {
        "input":
        {
            "nums1": [],
            "m": 0,
            "nums2": [],
            "n": 0,
        },
        "expected": [],
    },
    {
        "input":
        {
            "nums1": [1,2,3,0,0,0,0],
            "m": 3,
            "nums2": [2,3,4,5],
            "n": 4,
        },
        "expected": [1,2,2,3,3,4,5],
    },
    {
        "input": 
        {
            "nums1": [1,2,3,4,0,0,0],
            "m": 4,
            "nums2": [2,3,4],
            "n": 3,
        },
        "expected": [1,2,2,3,3,4,4],
    },
    {
        "input": 
        {
            "nums1": [2,0],
            "m": 1,
            "nums2": [1],
            "n": 1,
        },
        "expected": [1,2],
    },
]

a = Solution()

for test_case in test_cases:
    inpt = test_case["input"]
    expected = test_case["expected"]
    a.merge(inpt["nums1"], inpt["m"], inpt["nums2"], inpt["n"])
    if inpt["nums1"] != expected:
        print(f"data: '{inpt}', got: '{inpt["nums1"]}' expected: '{expected}'")
        os.exit(1)

print("PASS")
