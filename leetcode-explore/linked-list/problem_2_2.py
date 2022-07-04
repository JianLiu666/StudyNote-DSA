from typing import Optional

class Node:
    def __init__(self, val: int):
        self.val = val
        self.next = None

class Solution:
    def hasCycle(self, head: Optional[Node]) -> bool:
        if head == None:
            return False

        slow_node = head
        fast_node = head
        while fast_node.next != None:
            slow_node = slow_node.next
            fast_node = fast_node.next
            if fast_node.next == None:
                return False
            
            fast_node = fast_node.next
            if slow_node == fast_node:
                return True
        
        return False

if __name__ == '__main__':
    s = Solution()

    head = Node(1)
    cur_node = head
    for i in range(3):
        new_node = Node(1)
        cur_node.next = new_node
        cur_node = cur_node.next
    cur_node.next = head

    assert s.hasCycle(head) == True


