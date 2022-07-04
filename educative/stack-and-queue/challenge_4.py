'''
Next Greater Element Using a Stack

Problem Statement:
    You are required to implement the next_greater_element() function. 
    For each element i in a list, the function finds the first element to its right which is greater than element i. 
    If for any element such a value does not exist, the answer is -1.

    Note: 
        The next greater element is the first element towards the right which is greater than the given element. 
        For example, in the list [1, 3, 8, 4, 10, 5], the next greater element of 3 is 8 and the next greater element for 8 is 10.

Input:
    An integer list.

Sample Input:
    list = [4, 6, 3, 2, 8, 1]

Output:
    A list containing the next greater element of each element from the input list.

Sample Output:
    result = [6, 8, 8, 8, -1, -1]
'''

# Solution: Stack Iteration
# Time Complexity: O(n^2)
from typing import List


class Stack:
    def __init__(self):
        self.main = []
        self.length = 0
    
    def is_empty(self):
        return self.length == 0
    
    def size(self):
        return self.length
    
    def peek(self):
        if self.is_empty():
            return None
        
        return self.main[-1]
    
    def push(self, data):
        self.main.append(data)
        self.length += 1
    
    def pop(self):
        if self.is_empty():
            return None
        
        data = self.main.pop()
        self.length -= 1
        return data

def next_greater_element(lst: List):
    if len(lst) < 1:
        return lst
    
    stack = Stack()

    result = [-1] * len(lst)
    for i in range(len(lst)-1, -1, -1):
        if stack.is_empty() == False:
            while stack.is_empty() == False and stack.peek() <= lst[i]:
                stack.pop()
        if stack.is_empty() == False:
            result[i] = stack.peek()
        
        stack.push(lst[i])

    return result

list1 = [13, 3, 12, 16, 15, 11, 1, 2]
print(list1)
print(next_greater_element(list1))