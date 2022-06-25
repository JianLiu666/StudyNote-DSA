'''
Union & Intersection of Linked Lists

Problem Statement:
    Union and intersection are two of the most popular operations which can be performed on data sets. 

    Now, you will be implementing them for linked lists! Let’s take a look at their definitions:
     - Union: Given two lists, A and B, the union is the list that contains elements or objects that belong to either A, B, or to both.
     - Intersection: Given two lists, A and B, the intersection is the largest list which contains all the elements that are common to both the sets.

    The union function will take two linked lists and return their union.
    The intersection function will return all the elements that are common between two linked lists.

Input:
    Two linked lists.

Sample Input:
    list1 = 10->20->80->60
    list2 = 15->20->30->60->45

Output:
    A list containing the union of the two lists.
    A list containing the intersection of the two lists.

Sample Output:
    union = 10->20->80->60->15->30->45
    intersection = 20->60
'''

# My Solution
# Time Complexity: 
#  - Union: O(m)+O(l(l-1)/2) -> O(l^2), where m=len(lst1), n=len(lst2), l=m+n
#  - Intersection: max(O(mn), O(min(m,n)^2)), where m=len(lst1), n=len(lst2)
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
            return
        
        if self.head.next == None:
            return

        outer_node = self.head
        while outer_node:
            inner_node = outer_node
            while inner_node:
                if inner_node.next:
                    if outer_node.data == inner_node.next.data:
                        next_node = inner_node.next.next
                        inner_node.next = next_node
                    else:
                        inner_node = inner_node.next
                else: 
                    inner_node = inner_node.next
            
            outer_node = outer_node.next
    
    def search(self, data):
        if self.head == None:
            return False

        cur_node = self.head
        while cur_node:
            if cur_node.data == data:
                return True
            cur_node = cur_node.next
        
        return False

def union(lst1:LinkedList, lst2:LinkedList):
    if lst1.head == None:
        return lst2
    if lst2.head == None:
        return lst1
    
    tail_node = lst1.head
    while tail_node.next:
        tail_node = tail_node.next
    
    tail_node.next = lst2.head
    lst1.remove_duplicates()

    return lst1

ulist1 = LinkedList()
ulist1.insert_at_head(8)
ulist1.insert_at_head(22)
ulist1.insert_at_head(15)
print("ulist1:")
ulist1.print_list()

ulist2 = LinkedList()
ulist2.insert_at_head(21)
ulist2.insert_at_head(14)
ulist2.insert_at_head(7)
print("ulist2:")
ulist2.print_list()

union_list = union(ulist1,ulist2)
print("Union of ulist1 and ulist2:")
union_list.print_list()

def intersection(lst1:LinkedList, lst2:LinkedList):
    if lst1.head == None:
        return lst2
    if lst2.head == None:
        return lst1
    
    result = LinkedList()
    result_node = result.head

    lst1_node = lst1.head
    while lst1_node:
        if lst2.search(lst1_node.data):
            new_node = Node(lst1_node.data)
            if result.head == None:
                result.head = new_node
                result_node = result.head
            else:
                result_node.next = new_node
                result_node = new_node

        
        lst1_node = lst1_node.next

    result.remove_duplicates()

    return result

ilist1 = LinkedList()
ilist1.insert_at_head(14)
ilist1.insert_at_head(22)
ilist1.insert_at_head(15)
print("ilist1:")
ilist1.print_list()

ilist2 = LinkedList()
ilist2.insert_at_head(21)
ilist2.insert_at_head(14)
ilist2.insert_at_head(15)
print("ilist2:")
ilist2.print_list()

intersection_list = intersection(ilist1, ilist2)
print("Intersection of ilist1 and ilist2:")
intersection_list.print_list()