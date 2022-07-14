from typing import Optional, List

class Node:
    def __init__(self, val=0, prev=None, next=None, child=None):
        self.val = val
        self.prev = prev
        self.next = next
        self.child = child

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def flatten(self, head: 'Optional[Node]') -> 'Optional[Node]':
        def recursive(head: 'Optional[Node]') -> 'Optional[Node]':
            dummy = Node(0)
            dummy.next = head
            current = head

            while current:
                if current.child:
                    tmp = current.next
                    child_tail = recursive(current.child)
                    child_tail.next = current.next
                    if current.next:
                        current.next.prev = child_tail
                    current.next = current.child
                    current.child.prev = current
                    current.child = None
                    current = tmp
                    dummy = child_tail
                else:
                    current = current.next
                    dummy = dummy.next
            
            return dummy
        
        recursive(head)
        return head

def checkDoublyLinkedList(head: 'Optional[Node]') -> List[int]:
    dummy = Node(0)
    dummy.next = head
    current = head
    res = []
    while current:
        if current.child != None:
            return []
        if current != head and current.prev != dummy:
            return []
        
        res.append(current.val)
        current = current.next
        dummy = dummy.next

    return res

if __name__ == '__main__':
    s = Solution()

    # Example 1
    list1 = [Node(1), Node(2), Node(3), Node(4), Node(5), Node(6)]
    list1[0].next = list1[1]
    list1[len(list1)-1].prev = list1[len(list1)-2]
    for i in range(1, len(list1)-1):
        list1[i].next = list1[i+1]
        list1[i].prev = list1[i-1]
    
    list2 = [Node(7), Node(8), Node(9), Node(10)]
    list2[0].next = list2[1]
    list2[len(list2)-1].prev = list2[len(list2)-2]
    for i in range(1, len(list2)-1):
        list2[i].next = list2[i+1]
        list2[i].prev = list2[i-1]
    
    list1[2].child = list2[0]

    list3 = [Node(11), Node(12)]
    list3[0].next = list3[1]
    list3[1].prev = list3[0]

    list2[1].child = list3[0]

    assert checkDoublyLinkedList(s.flatten(list1[0])) == [1,2,3,7,8,11,12,9,10,4,5,6]
