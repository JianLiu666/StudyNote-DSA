class Node:
    def __init__(self, val):
        self.val = val
        self.next = None
        self.prev = None

class MyLinkedList:
    def __init__(self):
        self.head = None
        self.tail = None
        self.size = 0

    def get(self, index: int) -> int:
        if self.size == 0 or index < 0 or index >= self.size:
            return -1

        current = self.head
        for i in range(index):
            current = current.next
        
        return current.val

    def addAtHead(self, val: int) -> None:
        node = Node(val)
        if self.size == 0:
            self.head = node
            self.tail = node
            self.size += 1
            return
        
        self.head.prev = node
        node.next = self.head
        self.head = node
        self.size += 1

    def addAtTail(self, val: int) -> None:
        node = Node(val)
        if self.size == 0:
            self.head = node
            self.tail = node
            self.size += 1
            return
        
        self.tail.next = node
        node.prev = self.tail
        self.tail = node
        self.size += 1

    def addAtIndex(self, index: int, val: int) -> None:
        if index > self.size:
            return
        if index < 0:
            index = 0
        
        if index == 0:
            self.addAtHead(val)
            return
        
        if index == self.size:
            self.addAtTail(val)
            return
        
        node = Node(val)
        current = self.head
        for _ in range(0, index-1, 1):
            current = current.next
        
        node.next = current.next
        node.prev = current
        current.next.prev = node
        current.next = node
        self.size += 1

    def deleteAtIndex(self, index: int) -> None:
        if self.size == 0 or index < 0 or index >= self.size:
            return

        current = self.head
        for _ in range(index):
            current = current.next
        
        if current.prev != None:
            current.prev.next = current.next
        else:
            self.head = current.next
        
        if current.next != None:
            current.next.prev = current.prev
        else:
            self.tail = current.prev
        
        current.prev = None
        current.next = None
        
        self.size -= 1

if __name__ == '__main__':
    s = MyLinkedList()
    s.addAtHead(1)
    s.addAtTail(3)
    s.addAtIndex(1,2)
    assert s.get(1) == 2
    s.deleteAtIndex(1)
    assert s.get(1) == 3

    s = MyLinkedList()
    s.addAtHead(1)
    s.addAtTail(3)
    s.addAtIndex(1,2)
    assert s.get(1) == 2
    s.deleteAtIndex(0)
    assert s.get(0) == 2
