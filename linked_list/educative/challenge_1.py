'''
Problem Statement
    We need to insert a new object at the end of the linked list. 
    You can naturally guess, that this newly added node will point to None as it is at the tail.

Input
    A linked list and an integer value.

Sample Input
    Linked List = 0->1->2
    integer = 3

Output
    The updated linked list with the value inserted.

Sample Output
    Linked List = 0->1->2->3
'''

# My Solution
# Time Complexity: O(n)
class Node:
    def __init__(self, data):
        self.data = data
        self.next = None

class LinkedList:
    def __init__(self):
        self.head = None

    def is_empty(self):
        return True if self.head is None else False

    def print_list(self):
        if self.head is None:
            print("List is empty.")
            return
        
        cur_node = self.head
        while cur_node.next is not None:
            print(cur_node.data, end=" -> ")
            cur_node = cur_node.next
        print(cur_node.data, "-> None")

    def get_head(self):
        return self.head

    def insert_at_head(self, data):
        new_node = Node(data)
        new_node.next = self.head
        self.head = new_node
    
    def insert_at_tail(self, data):
        new_node = Node(data)

        if self.head is None:
            self.head = new_node
            return

        cur_node = self.head
        while cur_node.next is not None:
            cur_node = cur_node.next
        
        cur_node.next = new_node

def insert_at_tail(lst: LinkedList, value):
    if lst.is_empty():
        lst.insert_at_head(value)
    else:
        lst.insert_at_tail(value)

lst = LinkedList()
lst.print_list()
insert_at_tail(lst, 0)
lst.print_list()
insert_at_tail(lst, 1)
lst.print_list()
insert_at_tail(lst, 2)
lst.print_list()
insert_at_tail(lst, 3)
lst.print_list()