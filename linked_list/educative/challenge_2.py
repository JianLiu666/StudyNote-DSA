'''
Problem Statement
    The search algorithm in a linked list can be generalized to the following steps:
     - Start from the head node.
     - Traverse the list till you either find a node with the given value or you reach the end node which will indicate that the given node doesn’t exist in the list.

Input
    A linked list and an integer to be searched.

Sample Input
    Linked List = 5->90->10->4  
    Integer = 4

Output
    True if the integer is found. False otherwise.

Sample Output
    True
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

    def is_exists(self, data):
        if self.head is None:
            return False
        
        cur_node = self.head
        exists = False
        while exists != True:
            if cur_node.data == data:
                exists = True
                break
            
            if cur_node.next != None:
                cur_node = cur_node.next
            else:
                break
        
        return exists
    
    def print_list(self):
        if self.head is None:
            print("This list is empty.")
            return
        
        cur_node = self.head
        while cur_node.next != None:
            print(cur_node.data, end=" -> ")
            cur_node = cur_node.next
        print(cur_node.data, "-> None")
        

    def insert_at_head(self, data):
        new_node = Node(data)
        if self.head is None:
            self.head = new_node
        else:
            new_node.next = self.head
            self.head = new_node

def search(lst: LinkedList, value):
    return lst.is_exists(value)

lst = LinkedList()
lst.insert_at_head(4)
lst.insert_at_head(10)
lst.insert_at_head(40)
lst.insert_at_head(5)
lst.print_list()
print(search(lst, 4))