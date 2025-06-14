from typing import Optional 
import os

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        list = ListNode()
        head = list

        if list1 is None and list2 is None:
            return None
        elif list1 is not None and list2 is not None:
            if list1.val <= list2.val:
                list.val = list1.val
                list1 = list1.next
            else:
                list.val = list2.val
                list2 = list2.next
        elif list1 is not None:
            list.val = list1.val
            list1 = list1.next 
        elif list2 is not None:
            list.val = list2.val
            list2 = list2.next 
        else:
            return "WTF???"       

        while list1 is not None or list2 is not None:
            if list1 is None:
                while list2 is not None:
                    list.next = list2
                    list = list.next
                    list2 = list2.next
                return head
            if list2 is None:
                while list1 is not None:
                    list.next = list1
                    list = list.next
                    list1 = list1.next
                return head    

            if list1.val <= list2.val:
                list.next = list1
                list = list.next
                list1 = list1.next
            else:
                list.next = list2
                list = list.next
                list2 = list2.next

        return head


# constructListFromSlice returns head
def constructListFromSlice(arr):
    if len(arr) == 0:
        return None
    if len(arr) == 1:
        head = ListNode()
        head.val = arr[0]
        return head
    
    head = ListNode()
    head.val = arr[0]

    tail = ListNode()
    head.next = tail

    for i in range(1, len(arr)):
        tail.val = arr[i]
        new_tail = ListNode()
        if i < len(arr) - 1:
            tail.next = new_tail
            tail = new_tail

    return head


def linked_list_to_array(head):
    result = []
    current = head
    while current:
        result.append(current.val)
        current = current.next
    return result

        
test_cases = [
    [
        [1,2,4],
        [1,3,4],
        [1,1,2,3,4,4],
    ],
    [
        [],
        [0],
        [0],
    ],
    [
        [0],
        [],
        [0],
    ]
]

a = Solution()

for test_case in test_cases:
    res = linked_list_to_array(a.mergeTwoLists(constructListFromSlice(test_case[0]), constructListFromSlice(test_case[1])))
    if res != test_case[2]:
        print(f"data: '{test_case[0]}' ; '{test_case[1]}', got: '{res}' expected: '{test_case[2]}'")
        os.exit(1)

print("PASS")
