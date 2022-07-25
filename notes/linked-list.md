# Linked List

> *from [LeetCode Explore: Linked List](https://leetcode.com/explore/learn/card/linked-list/209/singly-linked-list//)*

---

- [Linked List](#linked-list)
  - [Introduction](#introduction)
    - [Singly Linked List](#singly-linked-list)
      - [Basic Operations](#basic-operations)
        - [Search : find a node in the linked list](#search--find-a-node-in-the-linked-list)
        - [Insertion : adds a new element to the linked list](#insertion--adds-a-new-element-to-the-linked-list)
        - [Deletion : removes the existing elements](#deletion--removes-the-existing-elements)
    - [Doubly Linked List](#doubly-linked-list)
  - [Two-Pointer in Linked List](#two-pointer-in-linked-list)

---

## Introduction

### Singly Linked List

Each node in a singly-linked list contains not only value but also **a reference field** to link to the next node.

By this way, the singly-linked list organizes all the nodes in a sequence.

**For example:**

```
 (head)
  node1          node2          node3
+---+---+      +---+---+      +---+---+
| 1 |   | ---> | 2 |   | ---> | 3 |   | ---> none
+---+---+      +---+---+      +---+---+
  ^   ^          ^   ^          ^   ^ 
 val ptr        val ptr        val ptr
```

#### Basic Operations

##### Search : find a node in the linked list

Unlike the array, we are not able to access a random element in a singly-linked list in constant time.

It take us `O(N)` time on average to **visit an element by index**, where *N* is the length of the linked list.

```
If we want to find 3 in the linked list:

Step.1 : let a current pointer start from the head node.

     NO!
  +---+---+      +---+---+      +---+---+
  | 1 |   | ---> | 2 |   | ---> | 3 |   | ---> none
  +---+---+      +---+---+      +---+---+
      ^
   current

Step.2 : move pointer to next node.

                    NO!
  +---+---+      +---+---+      +---+---+
  | 1 |   | ---> | 2 |   | ---> | 3 |   | ---> none
  +---+---+      +---+---+      +---+---+
                     ^
               -> current

Step.3 : until we found the node with value of 3

                                 GOTCHA!
  +---+---+      +---+---+      +---+---+
  | 1 |   | ---> | 2 |   | ---> | 3 |   | ---> none
  +---+---+      +---+---+      +---+---+
                                    ^
                              -> current

```

##### Insertion : adds a new element to the linked list

If we want to add a new element to the linked list, we have a few thing need to do :

```
Here is a given linked list, we want to add a node D between node B and node C

   node A                      node B         node C                      node N
  +---+---+                   +---+---+      +---+---+                   +---+---+
  | 1 |   | ---> ... ... ---> | 2 |   | ---X | 3 |   | ---> ... ... ---> | n |   | ---> none
  +---+---+                   +---+---+      +---+---+                   +---+---+
                                    |          ^
     2. let node B's "next" field   v          | 1. let node D's "next" filed
        point to node D           +---+---+    |    point to node C
                                  | 4 |   | ---┘
                                  +---+---+
                                   node D
```

Unlike an array, we don't need to move all elements past the inserted element, just in `O(1)` time complexity. But we need find the specific node first, it will cost `O(N)` time.

What about adding a new node at the head of the list or end of the list? Think about it.

##### Deletion : removes the existing elements

Let's talk about delete opearation : 

```
If we want to remove node B from the list.

   node A         node B         node C                      node N
  +---+---+      +---+---+      +---+---+                   +---+---+
  | 1 |   | ---X | 2 |   | ---> | 3 |   | ---> ... ... ---> | n |   | ---> none
  +---+---+      +---+---+      +---+---+                   +---+---+
        |                         ^
        └-------------------------┘
         1. let node A's "next" field
            point to node C
```

Like linked list insertion shown above, deletion is `O(1)` time complexity, but we also need to find the specific node which we want to delete.

Again, what about deleting the last node or the first node?

### Doubly Linked List

施工中 ...

## Two-Pointer in Linked List

施工中 ...