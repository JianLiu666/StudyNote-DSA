[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/partition-list/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)
---

# 86. Partition List

Given the `head` of a linked list and a value `x`, partition it such that all nodes **less than** `x` come before nodes **greater than or equal** to `x`.

You should **preserve** the original relative order of the nodes in each of the two partitions.

### Example 1:

![image](./image/partition.jpeg)

```
Input: head = [1,4,3,2,5,2], x = 3
Output: [1,2,2,4,3,5]
```

### Example 2:

```
Input: head = [2,1], x = 2
Output: [1,2]
```

### Constraints:

- The number of nodes in the list is in the range `[0, 200]`.
- `-100 <= Node.val <= 100`
- `-200 <= x <= 200`

### Related Topics

- Linked List
- Two Pointers
  
---

# 解題方向

題目的兩個重點:

- 以 `x` 為準，區分出左右兩個 `Linked Lists`
- 維持原本的相對位置

直接來吧，用兩個 `cursors` 分別記住切分之後的左半部與右半部，只要對原始的 `Linked List` 尋訪完一次之後，再將左右半部拼起來就好。

```
模擬一下 current 尋訪時的串接流程, 如下所示: (x=3)

current.val < x, 所以將 left_cursor.next 指到這個節點，反之就是將 right_cursor.next 指到這個節點

   head
    v
  +---+      +---+      +---+      +---+      +---+      +---+
  | 1 | ---> | 4 | ---> | 3 | ---> | 2 | ---> | 5 | ---> | 2 |
  +---+      +---+      +---+      +---+      +---+      +---+
    ^          ^
    |------┐  current'
  current  |  2. move current to next node
           |
 (dummy)   | 1. because current.val < x,
left_head  |    let left_cursor.next point to here
    v      |
  +---+    |
  | 0 | ---┘
  +---+
    ^
left_cursor
 
 (dummy)
right_head
    v
  +---+
  | 0 |
  +---+
    ^
right_cursor


如果把 current 移動的每一步以 s1, s2, ..., sn 表示，則最後結果如下: (x=3)
 
   head                                                  current
    v                                                      v
  +---+      +---+      +---+      +---+      +---+      +---+
  | 1 | ---> | 4 | ---> | 3 | ---> | 2 | ---> | 5 | ---> | 2 |
  +---+      +---+      +---+      +---+      +---+      +---+
   s1         s2         s3         s4         s5         s6

 (dummy)   
left_head                        left_cursor
    v                                v
  +---+      +---+      +---+      +---+
  | 0 | ---> | 1 | ---> | 2 | ---> | 2 | --┐
  +---+      +---+      +---+      +---+   |
              s1         s4         s6     | 1. let left_cursor.next point
                                           |    to right_head.next
               ┌---------------------------┘
 (dummy)       |
right_head     |                 right_cursor
    v          v                     v
  +---+      +---+      +---+      +---+          2. let right_cursor.next point to none,
  | 0 | ---> | 4 | ---> | 3 | ---> | 5 | --> None    otherwise it will become cyclic list.
  +---+      +---+      +---+      +---+            
              s2         s3         s5

```