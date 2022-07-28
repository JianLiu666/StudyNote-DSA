from collections import deque

class MyStack:

    def __init__(self):
        self.queue = deque()
        self.size = 0

    def push(self, x: int) -> None:
        size = len(self.queue)
        self.queue.append(x)
        for _ in range(size):
            self.queue.append(self.queue.popleft())
        
        self.size += 1

    def pop(self) -> int:
        if self.empty():
            return -1
        
        self.size -= 1
        return self.queue.popleft()

    def top(self) -> int:
        if self.empty():
            return -1

        return self.queue[0]

    def empty(self) -> bool:
        return len(self.queue) == 0

if __name__ == '__main__':
    obj = MyStack()

    obj.push(1)
    obj.push(2)
    assert obj.top() == 2
    assert obj.pop() == 2
    assert obj.empty() == False
