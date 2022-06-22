'''
Problem Statement
    You have to implement the find_mid() function which will take a linked list as an input and return the value of the middle node. 
    If the length of the list is even, the middle value will occur at length/2. For a list of odd length, the middle value will be (length+1)/2

Input
    A singly linked list.

Sample Input
    LinkedList = 7->14->10->21

Output
    The integer value of the middle node.

Sample Output
    14
'''

class Node:
    def __init__(self, data):
        self.data = data
        self.next = None

# My Solution
# Time Complexity: O(2n) -> O(n)
class LinkedList1:
    def __init__(self):
        self.head = None

    def insert_at_head(self, data):
        new_node = Node(data)
        if self.head == None:
            self.head = new_node
            return
        
        new_node.next = self.head
        self.head = new_node

    def print_list(self):
        if self.head == None:
            print("None")
            return
        
        cur_node = self.head
        while cur_node.next != None:
            print(cur_node.data, end=" -> ")
            cur_node = cur_node.next
        
        print(cur_node.data, "-> None")

    def get_length(self):
        if self.head == None:
            return 0
        
        count = 1
        cur_node = self.head
        while cur_node.next != None:
            count += 1
            cur_node = cur_node.next
        
        return count

    def find_mid(self):
        if self.head == None:
            return None
        
        cur_node = self.head

        # e.g. even: (4-1)//2 = 1 -> 從 head 往下數1個 = 第2個
        # e.g. odd:  (5-1)//2 = 2 -> 從 head 往下數2個 = 第3個
        mid = (self.get_length()-1)//2
        for i in range (mid):
            cur_node = cur_node.next

        return cur_node.data

def find_mid_sol1():
    lst = LinkedList1()
    lst.insert_at_head(22)
    lst.insert_at_head(21)
    lst.insert_at_head(10)
    lst.insert_at_head(14)
    lst.insert_at_head(7)

    lst.print_list()
    print(lst.find_mid())

find_mid_sol1()

# Solution: Two Pointers
# Time Complexity: O(n)
class LinkedList2:
    def __init__(self):
        self.head = None
    
    def insert_at_head(self, data):
        new_node = Node(data)
        if self.head == None:
            self.head = new_node
            return
        
        new_node.next = self.head
        self.head = new_node

    def print_list(self):
        if self.head == None:
            print("None")
            return
        
        cur_node = self.head
        while cur_node.next != None:
            print(cur_node.data, end=" -> ")
            cur_node = cur_node.next
        
        print(cur_node.data, "-> None")
    
    def find_mid(self):
        if self.head == None:
            return -1
        
        cur_node = self.head
        if cur_node.next == None:
            # only one element exists in list, so return this value
            return cur_node.data
        
        # Move mid_node (Slower) one step at a time
        # Move cur_node (Faster) two steps at a time
        # When cur_node reaches at end, mid_node will be at the middle of List 
        mid_node = cur_node
        cur_node = cur_node.next.next

        while cur_node:
            mid_node = mid_node.next
            cur_node = cur_node.next
            if cur_node == None:
                return mid_node.data
            cur_node = cur_node.next
        
        return mid_node.data

def find_mid_sol2():
    lst = LinkedList2()
    lst.insert_at_head(22)
    lst.insert_at_head(21)
    lst.insert_at_head(10)
    lst.insert_at_head(14)
    lst.insert_at_head(7)

    lst.print_list()
    print(lst.find_mid())

find_mid_sol2()