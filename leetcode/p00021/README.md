[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/merge-two-sorted-lists/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 21. Merge Two Sorted Lists

You are given the heads of two sorted linked lists `list1` and `list2`.

Merge the two lists in a one **sorted** list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

### Example 1:

![image](./image/merge_ex1.jpeg)

```
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
```

### Example 2:

```
Input: list1 = [], list2 = []
Output: []
```

### Example 3:

```
Input: list1 = [], list2 = [0]
Output: [0]
```

### Constraints:

- The number of nodes in both lists is in the range `[0, 50]`.
- `-100 <= Node.val <= 100`
- Both `list1` and `list2` are sorted in **non-decreasing** order.

### Related Topics

- Linked List
- Recursion
  
---

# 解題方向

基本的 Linked List 暖身題，Two Pointers (head and tail pointer) 解決。

直接展示吧：

```
step.0 初始化狀態，決定 head pointer 與 tail pointer 起點

         tail
         head       list1
          v          v
        +---+      +---+      +---+
list1 = | 1 | ---> | 3 | ---> | 5 |
        +---+      +---+      +---+

        +---+      +---+      +---+
list2 = | 2 | ---> | 4 | ---> | 6 |
        +---+      +---+      +---+
          ^
         list2

step.1 以 tail pointer 當作尋訪指針，指向小的位置 (i.e. list1.val > list2.val)

         tail
         head       list1
          v          v
        +---+      +---+      +---+
list1 = | 1 | ---X | 3 | ---> | 5 |
        +---+  |   +---+      +---+
               /
              /
             /
            /
           /
          |
          v
        +---+      +---+      +---+
list2 = | 2 | ---> | 4 | ---> | 6 |
        +---+      +---+      +---+
          ^
         list2

step.2 這次是 list1.val < list2.val，以此類推直到 tail pointer 尋訪到盡頭

         head       list1
          v          v
        +---+      +---+      +---+
list1 = | 1 |      | 3 | ---> | 5 |
        +---+      +---+      +---+
          |          ^
          |          |
          |         /
          |        /
          |       /
          |      /
          v     /
        +---+  |   +---+      +---+
list2 = | 2 | ---X | 4 | ---> | 6 |
        +---+      +---+      +---+
          ^          ^
         tail       list2
```