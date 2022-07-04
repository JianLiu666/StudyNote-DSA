class Node:

    def __init__(self, val: int):
        self.val = val
        self.next = None

class MyLinkedList:

    def __init__(self):
        self.head = None
        self.size = 0

    def get(self, index: int) -> int:
        if index < 0 or index >= self.size:
            return -1
        
        cur_node = self.head
        for i in range(index):
            cur_node = cur_node.next
        
        return cur_node.val

    def addAtHead(self, val: int) -> None:
        self.addAtIndex(0, val)

    def addAtTail(self, val: int) -> None:
        self.addAtIndex(self.size, val)

    def addAtIndex(self, index: int, val: int) -> None:
        if index > self.size:
            return
        if index < 0:
            index == 0
        
        new_node = Node(val)
        if index == 0:
            new_node.next = self.head
            self.head = new_node
            self.size += 1
            return
        
        cur_node = self.head
        for i in range(1, index):
            cur_node = cur_node.next
        
        new_node.next = cur_node.next
        cur_node.next = new_node
        self.size += 1

    def deleteAtIndex(self, index: int) -> None:
        if index < 0 or index >= self.size:
            return

        if index == 0:
            self.head = self.head.next
            self.size -= 1
            return
        
        cur_node = self.head
        for i in range(1, index):
            cur_node = cur_node.next
        
        cur_node.next = cur_node.next.next
        self.size -= 1

if __name__ == '__main__':
    obj = MyLinkedList()

    obj.addAtHead(1)
    obj.addAtTail(3)
    obj.addAtIndex(1, 2)
    assert obj.get(1) == 2
    
    obj.deleteAtIndex(1)
    assert obj.get(1) == 3