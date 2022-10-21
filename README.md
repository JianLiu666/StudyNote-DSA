# Data Structure and Algorithm

- [Data Structure and Algorithm](#data-structure-and-algorithm)
  - [Introduction](#introduction)
  - [My DSA Road Map](#my-dsa-road-map)
    - [Data Structure Notes](#data-structure-notes)
    - [Algorithm Notes](#algorithm-notes)
  - [References](#references)

---
## Introduction

- 施工中 ...

## My DSA Road Map

```
Data Structures

  +-------+                     +-------------+               +------------+         +---------------+          +------+                        +-------+
  | Array | ------------------> | Linked List | ------------> | Hash Table | ------> | Queue & Stack | -------> | Tree | ---------------------> | Graph | ---> keep going
  +-------+                     +-------------+               +------------+         +---------------+          +------+                        +-------+
    |                             |                             |                      |                          |                               |
    ├─ Two-Pointer Techniques     ├─ Two-Pointer Techniques     ├─ Hash Functions      ├─ Queue                   ├─ Tree Traversals              ├─ Adjacency Matrix
    ├─ In-place Algorithms        └─ Cycle Detection            ├─ Rolling Hash        |   └─ Circular Queue      |   ├─ Pre-order Traversal      ├─ Adjacency List
    ├─ Prefix Sum                                               ├─ Chaining            └─ Stack                   |   ├─ In-order Traversal       ├─ Undirected Graphs
    └─ String                                                   ├─ Open Addressing         ├─ Monotonic Stack     |   ├─ Post-order Traversal     |   └─ Rooted Tree
        ├─ Knuth–Morris–Pratt                                   └─ Memoization             └─ Explicit Stack      |   └─ Morris Traversal         ├─ Directed Graphs
        └─ Manacher Algorithm                                                                                     ├─ Binary Tree                  |   ├─ Directed Acyclic Graph
                                                                                                                  |   ├─ Binary Search Tree       |   └─ Topological Sorting
                                                                                                                  |   └─ Binary Heap Tree         └─ Disjoint Set
                                                                                                                  ├─ N-ary Tree                       └─ Minimum Spanning Trees
                                                                                                                  ├─ Trie                     
                                                                                                                  └─ Segment Tree

Algorithms

  +---------+                   +-----------+               +-----------+               +---------------------+
  | Sorting | ----------------> | Searching | ------------> | Recursion | ------------> | Dynamic Programming | ---> keep going
  +---------+                   +-----------+               +-----------+               +---------------------+
    |                             |                           |                            |
    ├─ Comparison                 ├─ Binary Search            ├─ Tail Recursion            ├─ Top-down
    |   ├─ Stable                 ├─ Breadth-First Search     ├─ Enumeration               └─ Bottom-up
    |   |   ├─ Bubble Sort        ├─ Depth-First Search       |   ├─ Backtracking     
    |   |   ├─ Insertion Sort     └─ Graph Search             |   └─ Branch and Bound 
    |   |   └─ Merge Sort             ├─ Dijkstra             ├─ Divide and Conquer   
    |   └─ Unstable                   ├─ Bellman-Ford         └─ Master Theorem       
    |       ├─ Selection Sort         ├─ Floyd-Warshall   
    |       ├─ Quick Sort             └─ Union Find       
    |       └─ Heap Sort      
    └─ Non-comparison         
        └─ Stable             
            ├─ Counting Sort  
            └─ Bucket Sort    
``` 

### Data Structure Notes

- [Array](./notes/array.md)
- [Linked List](./notes/linked-list.md)
- [Hash Table](./notes/hash-table.md)

### Algorithm Notes

- Sorting
- Searching

## References

- [官神請受我一拜🧎🏻](https://github.com/wisdompeak/LeetCode)
- [CList-LeetCode](https://clist.by/resource/leetcode.com/)
- [LeetCode Problem Rating](https://zerotrac.github.io/leetcode_problem_rating/#/)
- [Grind 75](https://www.techinterviewhandbook.org/grind75)