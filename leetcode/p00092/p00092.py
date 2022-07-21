from typing import Optional, List

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def reverseBetween(self, head: Optional[ListNode], left: int, right: int) -> Optional[ListNode]:
        if left == right:
            return head
        
        reverse_head = None
        reverse_head_prev = None
        reverse_tail = None
        
        previous = None
        current = head
        
        step = 1
        while current:
            if step == left:
                reverse_head_prev = previous
                reverse_tail = current
                
                tmp = current.next
                current.next = None
                previous = current
                current = tmp  
            
            elif step == right:
                reverse_tail.next = current.next
                current.next = previous

                if reverse_head_prev:
                    reverse_head_prev.next = current
                    return head
                else:
                    reverse_head = current
                    return reverse_head
            
            elif step > left and step < right:
                tmp = current.next
                current.next = previous
                previous = current
                current = tmp
            
            else:
                previous = current
                current = current.next
            
            step += 1

        return None

def getVals(head: Optional[ListNode]) -> List[int]:
    res = []

    while head:
        res.append(head.val)
        head = head.next

    return res

if __name__ == '__main__':
    s = Solution()

    list = [ListNode(1), ListNode(2), ListNode(3), ListNode(4), ListNode(5)]
    for i in range(0, len(list)-1):
        list[i].next = list[i+1]
    
    assert getVals(s.reverseBetween(list[0], 2, 4)) == [1,4,3,2,5]