class MyCircularQueue:

    def __init__(self, k: int):
        self.queue = [-1] * k
        self.head = 0
        self.tail = k-1
        self.size = 0
        self.capacity = k

    def enQueue(self, value: int) -> bool:
        if self.isFull() == True:
            return False
        
        self.tail += 1
        if self.tail == self.capacity:
            self.tail = 0
        
        self.queue[self.tail] = value
        self.size += 1
        
        return True

    def deQueue(self) -> bool:
        if self.isEmpty() == True:
            return False
        
        self.head += 1
        if self.head == self.capacity:
            self.head = 0
        
        self.size -= 1

        return True

    def Front(self) -> int:
        if self.size == 0:
            return -1
        
        return self.queue[self.head]

    def Rear(self) -> int:
        if self.size == 0:
            return -1
        
        return self.queue[self.tail]

    def isEmpty(self) -> bool:
        return self.size == 0

    def isFull(self) -> bool:
        return self.size == self.capacity

if __name__ == '__main__':
    obj = MyCircularQueue(3)
    assert obj.enQueue(1) == True
    assert obj.enQueue(2) == True
    assert obj.enQueue(3) == True
    assert obj.enQueue(4) == False
    assert obj.Rear() == 3
    assert obj.isFull() == True
    assert obj.deQueue() == True
    assert obj.enQueue(4) == True
    assert obj.Rear() == 4

    obj = MyCircularQueue(3)
    assert obj.enQueue(7) == True
    assert obj.deQueue() == True
    assert obj.Front() == -1
    assert obj.deQueue() == False
    assert obj.Front() == -1
    assert obj.Rear() == -1
    assert obj.enQueue(0) == True
    assert obj.isFull() == False
    assert obj.deQueue() == True
    assert obj.Rear() == -1
    assert obj.enQueue(3) == True