'''
min( ) Function Using a Stack

Problem Statement:
    You have to implement the MinStack class which will have a min() function. 
    Whenever min() is called, the minimum value of the stack is returned in O(1) time. 
    The element is not popped from the stack. Its value is simply returned.

Input:
    min() operates on an object of MinStack and doesn’t take any input

Output:
    Returns minimum number in O(1) time

Sample Output:
    min_stack = [9, 3, 1, 4, 2, 5]
    min_stack.min()
    Result = 1
'''

# Solution: Two Stacks
# Time Complexity: O(1)
class Stack:
    def __init__(self):
        self.main = []
        self.length = 0

    def is_empty(self):
        return self.length == 0

    def size(self):
        return self.length

    def push(self, data):
        self.main.append(data)
        self.length += 1

    def pop(self):
        if self.is_empty():
            return None
        
        self.length -= 1
        return self.main.pop()

    def peek(self):
        if self.is_empty():
            return None
        
        return self.main[-1]

class MinStack:
    def __init__(self):
        self.main = Stack()
        self.mins = Stack()

    def push(self, value):
        self.main.push(value)
        if self.mins.is_empty() or self.mins.peek() > value:
            self.mins.push(value)
        else:
            self.mins.push(self.mins.peek())

    def pop(self):
        self.mins.pop()
        return self.main.pop()

    def min(self):
        return self.mins.peek()

if __name__ == "__main__" :
    stack = MinStack()
    stack.push(5)
    stack.push(0)
    stack.push(2)
    stack.push(4)
    stack.push(1)
    stack.push(3)
    stack.push(0)
    print("Main stack:", stack.main.main)
    print("Min stack:", stack.mins.main)
    print("Minimum value: " + str(stack.min()))
    stack.pop()
    stack.push(-2)
    print("Main stack:", stack.main.main)
    print("Min stack:", stack.mins.main)
    print("Minimum value: " + str(stack.min()))
    stack.pop()
    print("Main stack:", stack.main.main)
    print("Min stack:", stack.mins.main)
    print("Minimum value: " + str(stack.min()))