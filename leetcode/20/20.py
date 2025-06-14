import os

class Solution:
    def isValid(self, s: str) -> bool:
        count = [0,0,0]
        last_open = []
        for i in range(len(s)):
            if s[i] == '(':
                last_open.append(s[i])

                count[0] += 1
            elif s[i] == ')':
                if len(last_open) > 0:
                    if last_open[-1] != '(':
                        return False
                    last_open.pop()
                
                if count[0] <= 0:
                    return False
                count[0] -= 1
            elif s[i] == '[':
                last_open.append(s[i])

                count[1] += 1
            elif s[i] == ']':
                if len(last_open) > 0:
                    if last_open[-1] != '[':
                        return False
                    last_open.pop()
                
                if count[1] <= 0:
                    return False
                
                count[1] -= 1
            elif s[i] == '{':
                last_open.append(s[i])

                count[2] += 1
            elif s[i] == '}':
                if len(last_open) > 0:
                    if last_open[-1] != '{':
                        return False
                    last_open.pop()

                if count[2] <= 0:
                    return False
                count[2] -= 1    

        if sum(count) == 0:
            return True
        
        return False 

test_cases = [
    [
        "(((})))",
        False,
    ],
    [
        "[",
        False,
    ],
    [
        "}",
        False,
    ],
    [
        "()",
        True,
    ],
    [
        "({[]})",
        True,
    ],
    [
        "([)]",
        False,
    ]
]

a = Solution()

for test_case in test_cases:
    res = a.isValid(test_case[0])
    if res != test_case[1]:
        print(f"data: '{test_case[0]}', got: '{res}' expected: '{test_case[1]}'")
        os.exit(1)

print("PASS")
