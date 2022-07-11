from typing import List, Optional

class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def getIntersectionNode_alignment(self, headA: ListNode, headB: ListNode) -> Optional[ListNode]:
        if headA == None or headB == None:
            return None
        
        lenA, lenB = 0, 0
        currentA, currentB = headA, headB

        while currentA:
            lenA += 1
            currentA = currentA.next

        while currentB:
            lenB += 1
            currentB = currentB.next
        
        currentA, currentB = headA, headB
        
        for _ in range(0, lenA - lenB, 1):
            currentA = currentA.next

        for _ in range(0, lenB - lenA, 1):
            currentB = currentB.next
        
        while currentA:
            if currentA == currentB:
                return currentA
            
            currentA = currentA.next
            currentB = currentB.next
        
        return None

    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def getIntersectionNode_cycle(self, headA: ListNode, headB: ListNode) -> Optional[ListNode]:
        currentA = headA
        currentB = headB
        
        while currentA != currentB:
            currentA = currentA.next if currentA != None else headB
            currentB = currentB.next if currentB != None else headA
        
        return currentA

if __name__ == '__main__':
    s = Solution()

    # Example 1:
    listA = [ListNode(4), ListNode(1)]
    listB = [ListNode(5), ListNode(6), ListNode(1)]
    intersects = [ListNode(8), ListNode(4), ListNode(5)]

    for i in range(len(listA)):
        if i != len(listA)-1:
            listA[i].next = listA[i+1]
        else:
            listA[i].next = intersects[0]
    
    for i in range(len(listB)):
        if i != len(listB)-1:
            listB[i].next = listB[i+1]
        else:
            listB[i].next = intersects[0]
    
    assert s.getIntersectionNode_alignment(listA[0], listB[0]) == intersects[0]
    assert s.getIntersectionNode_cycle(listA[0], listB[0]) == intersects[0]

    # Example 2:
    listA = [ListNode(1), ListNode(9), ListNode(1)]
    listB = [ListNode(3)]
    intersects = [ListNode(2), ListNode(4)]

    for i in range(len(listA)):
        if i != len(listA)-1:
            listA[i].next = listA[i+1]
        else:
            listA[i].next = intersects[0]
    
    for i in range(len(listB)):
        if i != len(listB)-1:
            listB[i].next = listB[i+1]
        else:
            listB[i].next = intersects[0]
    
    assert s.getIntersectionNode_alignment(listA[0], listB[0]) == intersects[0]
    assert s.getIntersectionNode_cycle(listA[0], listB[0]) == intersects[0]

    # Example 3:
    listA = [ListNode(2), ListNode(6), ListNode(4)]
    listB = [ListNode(1), ListNode(5)]

    for i in range(len(listA)-1):
        listA[i].next = listA[i+1]
    
    for i in range(len(listB)-1):
        listB[i].next = listB[i+1]
    
    assert s.getIntersectionNode_alignment(listA[0], listB[0]) == None
    assert s.getIntersectionNode_cycle(listA[0], listB[0]) == None
