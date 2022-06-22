'''
Problem Statement
    In this function, we can pass a particular value that we want to delete from the list. 
    The node containing this value could be anywhere in the list. It is also possible that such a node may not exist at all.

Input
    A linked list and an integer to be deleted.

Sample Input
    LinkedList = 3->2->1->0
    Integer = 2

Output
    True if the value is deleted. Otherwise, False.

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

    def insert_at_head(self, value):
        new_node = Node(value)
        if self.head == None:
            self.head = new_node
            return
        
        new_node.next = self.head
        self.head = new_node
    
    def delete_at_head(self):
        if self.head == None:
            return
        
        self.head = self.head.next

    def delete_by_value(self, value):
        if self.head == None:
            return False
        
        cur_node = self.head
        if cur_node.data == value:
            self.delete_at_head()
            return True
        
        while cur_node.next != None:
            pre_node = cur_node
            cur_node = cur_node.next
            if cur_node.data == value:
                pre_node.next = cur_node.next
                return True
        
        return False


    def print_list(self):
        if self.head == None:
            print("None")
            return
        
        cur_node = self.head
        while cur_node.next != None:
            print(cur_node.data, end=" -> ")
            cur_node = cur_node.next
        
        print(cur_node.data, "-> None")


def delete(lst:LinkedList, value):
    return lst.delete_by_value(value)

lst = LinkedList()
lst.insert_at_head(1)
lst.insert_at_head(4)
lst.insert_at_head(3)
lst.insert_at_head(2)
lst.print_list()
delete(lst, 4)
lst.print_list()