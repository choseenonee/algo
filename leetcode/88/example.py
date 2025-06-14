from typing import List
import os

class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        i = m - 1
        j = n - 1
        k = m + n - 1

        while i >= 0 and j >= 0: 
            if nums1[i] > nums2[j]: 
                nums1[k] = nums1[i]
                i -= 1
            else: 
                nums1[k] = nums2[j]
                j -= 1
            k -= 1
        
        while j >= 0: 
            nums1[k] = nums2[j]
            j -= 1
            k -= 1


test_cases = [
    {
    "input": {
      "nums1": [1, 2, 3, 4, 0, 0],
      "m": 4,
      "nums2": [2, 5],
      "n": 2
    },
    "expected": [1, 2, 2, 3, 4, 5]
  },
#   {
#     "input": {
#       "nums1": [1, 2, 3, 0, 0, 0],
#       "m": 3,
#       "nums2": [4, 5, 6],
#       "n": 3
#     },
#     "expected": [1, 2, 3, 4, 5, 6]
#   },
#   {
#     "input": {
#       "nums1": [1],
#       "m": 1,
#       "nums2": [],
#       "n": 0
#     },
#     "expected": [1]
#   },
#   {
#     "input": {
#       "nums1": [0],
#       "m": 0,
#       "nums2": [1],
#       "n": 1
#     },
#     "expected": [1]
#   },
#     {
#         "input":
#         {
#             "nums1": [],
#             "m": 0,
#             "nums2": [],
#             "n": 0,
#         },
#         "expected": [],
#     },
#     {
#         "input":
#         {
#             "nums1": [1,2,3,0,0,0,0],
#             "m": 3,
#             "nums2": [2,3,4,5],
#             "n": 4,
#         },
#         "expected": [1,2,2,3,3,4,5],
#     },
#     {
#         "input": 
#         {
#             "nums1": [1,2,3,4,0,0,0],
#             "m": 4,
#             "nums2": [2,3,4],
#             "n": 3,
#         },
#         "expected": [1,2,2,3,3,4,4],
#     },
#     {
#         "input": 
#         {
#             "nums1": [2,0],
#             "m": 1,
#             "nums2": [1],
#             "n": 1,
#         },
#         "expected": [1,2],
#     },
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
