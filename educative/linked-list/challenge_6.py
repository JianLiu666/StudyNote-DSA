'''
Detect Loop in a Linked List

Problem Statement:
    By definition, a loop is formed when a node in your linked list points to a previously traversed node.
    You must implement the detect_loop() function which will take a linked list as input and deduce whether or not a loop is present.

Input:
    A singly linked list.

Sample Input:
    LinkedList = 7->14->21->7 # Both '7's are the same node. Not duplicates.

Output:
    Returns True if the given linked list contains a loop. Otherwise, it returns False

Sample Output:
    True
'''

# Solution: Floyd’s Cycle-Finding Algorithm
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

    def detect_loop(self):
        # head->none 在這邊就擋掉了
        if self.head == None:
            return False

        onestep = self.head
        twostep = self.head

        # head->1->none 在這邊就先擋掉了
        # head->1->2->... 開始檢查
        while onestep and twostep and twostep.next:
            onestep = onestep.next
            twostep = twostep.next.next

            if onestep.data == twostep.data:
                return True
        
        return False

    def get_head(self):
        return self.head

    def print_list(self):
        if self.head == None:
            print("None")
            return
        
        cur_node = self.head
        while cur_node.next != None:
            print(cur_node.data, end=" -> ")
            cur_node = cur_node.next
        
        print(cur_node.data, "-> None")

def detect_loop(lst:LinkedList):
    return lst.detect_loop()

lst = LinkedList()

lst.insert_at_head(21)
lst.insert_at_head(14)
lst.insert_at_head(7)

# Adding a loop
head = lst.get_head()
node = lst.get_head()

for i in range(4):
    if node.next is None:
        node.next = head.next
        break
    node = node.next

print(detect_loop(lst))

