'''
Return the Nth node from End

Problem Statement:
    In the find_nth function, a certain N is specified as an argument. 
    You simply need to return the node which is N spaces away from the None end of the linked list.

Input:
    A linked list and a position N.

Sample Input:
    LinkedList = 22->18->60->78->47->39->99 and n = 3

Output:
    The value of the node n positions from the end. Returns -1 if n is out of bounds.

Sample Output:
    47
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
        while cur_node.next:
            print(cur_node.data, end=" -> ")
            cur_node = cur_node.next
        print(cur_node.data, " -> None")

    def find_nth(self, n):
        if self.head == None:
            return -1
        
        count = 0
        cur_node = self.head
        while cur_node:
            count += 1
            cur_node = cur_node.next
        
        if n > count:
            return -1
        
        cur_node = self.head
        for i in range (count-n):
            cur_node = cur_node.next
        
        return cur_node.data

def find_nth_sol1():
    lst = LinkedList1()
    lst.insert_at_head(21)
    lst.insert_at_head(14)
    lst.insert_at_head(7)
    lst.insert_at_head(8)
    lst.insert_at_head(22)
    lst.insert_at_head(15)
    lst.print_list()

    print(lst.find_nth(5))
    print(lst.find_nth(1))
    print(lst.find_nth(10))

find_nth_sol1()

# Solution 2: Two Pointers
# Time Complexity: O(n)
class LinkedList2:
    def __init__(self):
        self.head = None

    def insert_at_head(self, data):
        new_node = Node(data)
        if self.head:
            new_node.next = self.head

        self.head = new_node
    
    def print_list(self):
        if self.head == None:
            print("None")
            return
        
        cur_node = self.head
        while cur_node.next:
            print(cur_node.data, end=" -> ")
            cur_node = cur_node.next
        
        print(cur_node.data, " -> None")

    def find_nth(self, n):
        if self.head == None:
            return -1
        
        nth_node = self.head
        end_node = self.head

        count = 0
        while count < n:
            end_node = end_node.next
            if end_node == None:
                return -1
            
            count += 1
        
        while end_node:
            end_node = end_node.next
            nth_node = nth_node.next
        
        return nth_node.data

def find_nth_sol2():
    lst = LinkedList2()
    lst.insert_at_head(21)
    lst.insert_at_head(14)
    lst.insert_at_head(7)
    lst.insert_at_head(8)
    lst.insert_at_head(22)
    lst.insert_at_head(15)
    lst.print_list()

    print(lst.find_nth(5))
    print(lst.find_nth(1))
    print(lst.find_nth(10))

find_nth_sol2()