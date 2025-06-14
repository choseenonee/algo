import os

class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        if len(needle) > len(haystack):
            return -1
        if len(needle) == len(haystack):
            return 0 if needle == haystack else -1
        
        for i in range(len(haystack)-len(needle)+1):
            if haystack[i:i+len(needle)] == needle:
                return i

        return -1

test_cases = [
    [
        "sadbutsad",
        "sad",
        0
    ],
    [
        "leetcode",
        "leeto",
        -1,
    ],
    [
        "abc",
        "c",
        2
    ],
    [
        "aaaa",
        "baaa",
        -1,
    ]
]

a = Solution()

for test_case in test_cases:
    res = a.strStr(test_case[0], test_case[1])
    if res != test_case[2]:
        print(f"data: '{test_case[0]}' '{test_case[1]}', got: '{res}' expected: '{test_case[2]}'")
        os.exit(1)

print("PASS")
