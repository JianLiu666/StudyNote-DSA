'''
Remove Duplicates from Linked List

Problem Statement:
    You will now be implementing the remove_duplicates() function. 
    When a linked list is passed to this function, it removes any node which is a duplicate of another existing node.

Input:
    A linked list.

Sample Input:
    LinkedList = 1->2->2->2->3->4->4->5->6

Output:
    A list with all the duplicates removed.

Sample Output:
    LinkedList = 1->2->3->4->5->6
'''

# My Solution
# Time Complexity: O(n(n-1)/2) -> O(n^2)
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
    
    def print_list(self):
        if self.head == None:
            print("None")
            return
        
        cur_node = self.head
        while cur_node.next != None:
            print(cur_node.data, end=" -> ")
            cur_node = cur_node.next
        
        print(cur_node.data, "-> None")

    def remove_duplicates(self):
        if self.head == None:
            return None
        
        # If list only has one node, leave it unchanged
        if self.head.next == None:
            return lst

        outer_node = self.head
        while outer_node:
            # Iterator for the inner loop
            inner_node = outer_node
            while inner_node:
                if inner_node.next:
                    if outer_node.data == inner_node.next.data:
                        # Duplicate found, so now removing it
                        next_node = inner_node.next.next
                        inner_node.next = next_node
                    else:
                        # Otherwise simply iterate ahead
                        inner_node = inner_node.next
                else:
                    # Otherwise simply iterate ahead
                    inner_node = inner_node.next
            outer_node = outer_node.next

def remove_duplicates(lst:LinkedList):
    return lst.remove_duplicates()

lst = LinkedList()
lst.insert_at_head(7)
lst.insert_at_head(7)
lst.insert_at_head(7)
lst.insert_at_head(22)
lst.insert_at_head(14)
lst.insert_at_head(21)
lst.insert_at_head(14)
lst.insert_at_head(7)

lst.print_list()
remove_duplicates(lst)
lst.print_list()