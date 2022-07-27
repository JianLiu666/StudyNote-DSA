class MinStack:

    def __init__(self):
        self.arr = []
        self.min = []

    def push(self, val: int) -> None:
        self.arr.append(val)

        if len(self.min) == 0 or val <= self.min[-1]:
            self.min.append(val)

    def pop(self) -> None:
        if self.min[-1] == self.arr[-1]:
            self.min.pop()
        
        self.arr.pop()

    def top(self) -> int:
        return self.arr[-1]

    def getMin(self) -> int:
        return self.min[-1]

if __name__ == '__main__':
    obj = MinStack()
    obj.push(-2)
    obj.push(0)
    obj.push(-3)
    assert obj.getMin() == -3
    obj.pop()
    assert obj.top() == 0
    assert obj.getMin() == -2
