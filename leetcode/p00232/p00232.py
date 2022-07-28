from collections import deque

class MyQueue:

    def __init__(self):
        self.input = deque()
        self.output = deque()

    def push(self, x: int) -> None:
        self.input.append(x)

    def pop(self) -> int:
        if self.empty():
            return -1
        
        if len(self.output) == 0:
            while self.input:
                self.output.append(self.input.pop())
        
        return self.output.pop()

    def peek(self) -> int:
        if self.empty():
            return -1
        
        if len(self.output) == 0:
            while self.input:
                self.output.append(self.input.pop())
            
        return self.output[-1]


    def empty(self) -> bool:
        return True if len(self.input) + len(self.output) == 0 else False

if __name__ == '__main__':
    obj = MyQueue()
    obj.push(1)
    obj.push(2)
    assert obj.peek() == 1
    assert obj.pop() == 1
    assert obj.empty() == False
