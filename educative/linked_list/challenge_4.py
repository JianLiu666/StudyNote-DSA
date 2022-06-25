'''
Find the Length of a Linked List

Problem Statement:
    In this problem, you have to implement the length() function which will find the length of a given linked list.

Input:
    A linked list.

Sample Input:
    linkedlist = 0->1->2->3->4

Output:
    The number of nodes in the list.

Sample Output:
    5 
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

    def insert_at_head(self, data):
        new_node = Node(data)
        if self.head == None:
            self.head = new_node
            return
        
        new_node.next = self.head
        self.head = new_node

    def length(self):
        if self.head == None:
            return 0
        
        cur_node = self.head
        count = 0
        while cur_node != None:
            count += 1
            cur_node = cur_node.next
        
        return count

def length(lst:LinkedList):
    return lst.length()

lst = LinkedList()
lst.insert_at_head(4)
lst.insert_at_head(3)
lst.insert_at_head(2)
lst.insert_at_head(1)
lst.insert_at_head(0)
print(length(lst))