'''
Reversing First k Elements of Queue

Problem Statement:
    Implement the function reverseK(queue, k) which takes a queue and a number “k” as input and reverses the first “k” elements of the queue.

Sample Input:
    Queue = [1,2,3,4,5,6,7,8,9,10], k = 5

Output:
    The queue with first “k” elements reversed. Remember to return the queue itself!
    In case the value of “k” is larger than the size of the queue, is smaller than 0, or if the queue is empty, simply return None instead.

Sample Output:
    Queue = [5,4,3,2,1,6,7,8,9,10]
'''

# SolutionL Using a Stack
# Time Complexity: O(n)
class Node:
    def __init__(self, data):
        self.data = data
        self.next = None
        self.previous = None

class DoublyLinkedList:
    def __init__(self):
        self.head = None
        self.tail = None
        self.length = 0
    
    def is_empty(self):
        return self.length == 0

    def get_head(self):
        return None if self.is_empty() else self.head.data
    
    def get_tail(self):
        return None if self.is_empty() else self.tail.data

    def insert_at_tail(self, value):
        new_node = Node(value)
        if self.is_empty():
            self.head = new_node
            self.tail = new_node
        else:
            self.tail.next = new_node
            new_node.previous = self.tail
            self.tail = new_node

        self.length += 1

    def remove_at_head(self):
        if self.is_empty():
            return

        if self.length == 0:
            self.head = None
            self.tail = None
        else:
            self.head = self.head.next
            if self.head:
                self.head.previous = None
        
        self.length -= 1

    def __str__(self):
        text = ""
        if self.is_empty():
            return text
        
        current_node = self.head
        text += "[" + str(current_node.data) + ", "

        current_node = current_node.next
        while current_node.next:
            text += str(current_node.data) + ", "
            current_node = current_node.next
        
        text += str(current_node.data) + "]"

        return text

class Stack:
    def __init__(self):
        self.list = []
        self.size = 0

    def is_empty(self):
        return self.size == 0

    def peek(self):
        if self.is_empty():
            return None
        
        return self.list[-1]

    def size(self):
        return self.size
    
    def push(self, data):
        self.list.append(data)
        self.size += 1

    def pop(self):
        if self.is_empty():
            return None
        
        self.size -= 1
        return self.list.pop()

class Queue:
    def __init__(self):
        self.items = DoublyLinkedList()

    def is_empty(self):
        return self.items.is_empty()

    def front(self):
        return None if self.is_empty() else self.items.get_head()

    def rear(self):
        return None if self.is_empty() else self.items.get_tail()

    def size(self):
        return self.items.length

    def enqueue(self, data):
        self.items.insert_at_tail(data)

    def dequeue(self):
        data = self.items.get_head()
        self.items.remove_at_head()

        return data
    
    def display(self):
        return self.items.__str__()

def reverseK(queue:Queue, k):
    if queue.is_empty() or k < 0 or k > queue.size():
        return None

    stack = Stack()
    for i in range(k):
        stack.push(queue.dequeue())
    
    while stack.is_empty() == False:
        queue.enqueue(stack.pop())
    
    size = queue.size()
    for i in range(size-k):
        queue.enqueue(queue.dequeue())

    return queue

if __name__ == "__main__":
    queue = Queue()
    queue.enqueue(1)
    queue.enqueue(2)
    queue.enqueue(3)
    queue.enqueue(4)
    queue.enqueue(5)
    queue.enqueue(6)
    queue.enqueue(7)
    queue.enqueue(8)
    queue.enqueue(9)
    queue.enqueue(10)
    print("the queue before reversing:")
    print(queue.items)
    reverseK(queue, 5)
    print("the queue after reversing:")
    print(queue.items)