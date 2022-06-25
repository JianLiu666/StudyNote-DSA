'''
Implement a Queue Using Stacks

Problem Statement:
    You have to implement the enqueue() and dequeue() functions using the MyStack class we created earlier. 
    enqueue() will insert a value into the queue and dequeue() will remove a value from the queue.

Input:
    enqueue(): A value to insert into the queue
    dequeue(): Does not require any input

Sample Input:
    value = 5 # [1, 2, 3, 4]
    enqueue(value)
    dequeue()

Output:
    enqueue(): Does not return anything
    dequeue(): Pops out and returns the oldest value in the queue

Sample Output:
    True # [1, 2, 3, 4, 5]
    1 # [2, 3, 4, 5]
'''

class Stack:
    def __init__(self):
        self.list = []
        self.length = 0
    
    def is_empty(self):
        return self.length == 0

    def peek(self):
        return None if self.is_empty() else self.list[-1]

    def size(self):
        return self.length

    def push(self, data):
        self.list.append(data)
        self.length += 1
    
    def pop(self):
        if self.is_empty():
            return None
        
        data = self.list.pop()
        self.length -= 1
        return data

# Solution1: Two Stacks Working in enqueue()
# Time Complexity:
#  - enqueue: O(n)
#  - dequeue: O(1)
class Queue1:
    def __init__(self):
        self.main = Stack()
        self.second = Stack()
    
    def enqueue(self, data):
        if self.main.is_empty() and self.second.is_empty():
            self.main.push(data)
            return

        # 確保這次放進去的會在 stack 最下面
        while self.main.is_empty() == False:
            self.second.push(self.main.pop())
        
        self.main.push(data)

        while self.second.is_empty() == False:
            self.main.push(self.second.pop())
    
    def dequeue(self):
        return None if self.main.is_empty() else self.main.pop()

def solution1():
    queue = Queue1()
    count = 1
    for i in range(3):
        queue.enqueue(count)
        count += 1
    for i in range(2):
        print(queue.dequeue())
    for i in range(2):
        queue.enqueue(count)
        count += 1
    for i in range(3):
        print(queue.dequeue())

# Solution 2: Two Stacks Working in dequeue()
# Time Complexity:
#  - Enqueue: O(1)
#  - Dequeue: O(n)
class Queue2:
    def __init__(self):
        self.main = Stack()
        self.second = Stack()

    def enqueue(self, data):
        self.main.push(data)
    
    def dequeue(self):
        if self.main.is_empty() and self.second.is_empty():
            return None
        
        if self.second.is_empty() == False:
            return self.second.pop()
        
        while self.main.is_empty() == False:
            self.second.push(self.main.pop())
        
        return self.second.pop()

def solution2():
    queue = Queue2()
    count = 1
    for i in range(3):
        queue.enqueue(count)
        count += 1
    for i in range(2):
        print(queue.dequeue())
    for i in range(2):
        queue.enqueue(count)
        count += 1
    for i in range(3):
        print(queue.dequeue())

if __name__ == "__main__" :
  solution1()
  print("----------")
  solution2()