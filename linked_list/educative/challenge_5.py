'''
Reverse a Linked List

Problem Statement:
    You have to define the reverse function, which takes a singly linked list and produces the exact opposite list.
    i.e., the links of the output linked list should be reversed.

Input:
    A singly linked list.

Sample Input:
    The input linked list object:
    LinkedList = 0->1->2->3-4

Output:
    The reversed linked list.

Sample Output:
    The reversed linked list:
    LinkedList = 4->3->2->1->0
'''

class Node:
    def __init__(self, data):
        self.data = data
        self.next = None

# My Solution
# Time Complexity: O(n)
class LinkedList1:
    def __init__(self):
        self.head = None

    def insert_at_head(self, value):
        new_node = Node(value)
        
        if self.head == None:
            self.head = new_node
            return
        
        new_node.next = self.head
        self.head = new_node

    def reverse(self):
        if self.head == None:
            return
        
        old_head = self.head
        cur_node = self.head
        nxt_node = cur_node.next
        while nxt_node != None:
            tmp_node = nxt_node.next
            nxt_node.next = cur_node
            cur_node = nxt_node
            nxt_node = tmp_node
        
        self.head = cur_node
        old_head.next = None

    def print_list(self):
        if self.head == None:
            print("Linked list is empty")
            return
        
        cur_node = self.head
        while cur_node.next != None:
            print(cur_node.data, end=" -> ")
            cur_node = cur_node.next
        print(cur_node.data, "-> None")

def reverse_sol1():
    lst = LinkedList1()
    lst.insert_at_head(6)
    lst.insert_at_head(4)
    lst.insert_at_head(9)
    lst.insert_at_head(10)
    lst.print_list()
    lst.reverse()
    lst.print_list()

reverse_sol1()

# Solution: Iterative Pointer Manipulation#
# Time Complexity: O(n)
class LinkedList2:
    def __init__(self):
        self.head = None

    def insert_at_head(self, value):
        new_node = Node(value)
        
        if self.head == None:
            self.head = new_node
            return
        
        new_node.next = self.head
        self.head = new_node

    def reverse(self):
        if self.head == None:
            return

        previous, current, next = None, self.head, None
        while current:
            next = current.next
            current.next = previous
            previous = current
            current = next
            self.head = previous

    def print_list(self):
        if self.head == None:
            print("Linked list is empty")
            return
        
        cur_node = self.head
        while cur_node.next != None:
            print(cur_node.data, end=" -> ")
            cur_node = cur_node.next
        print(cur_node.data, "-> None")

def reverse_sol2():
    lst = LinkedList2()
    lst.insert_at_head(6)
    lst.insert_at_head(4)
    lst.insert_at_head(9)
    lst.insert_at_head(10)
    lst.print_list()
    lst.reverse()
    lst.print_list()

reverse_sol2()