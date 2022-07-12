[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/odd-even-linked-list/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 328. Odd Even Linked List

Given the `head` of a singly linked list, group all the nodes with odd indices together followed by the nodes with even indices, and return the reordered list.

The **first** node is considered **odd**, and the **second** node is **even**, and so on.

Note that the relative order inside both the even and odd groups should remain as it was in the input.

You must solve the problem in `O(1)` extra space complexity and `O(n)` time complexity.

### Example 1:

![image](./image/oddeven-linked-list.jpeg)

```
Input: head = [1,2,3,4,5]
Output: [1,3,5,2,4]
```

### Example 2:

```
Input: head = [2,1,3,5,6,4,7]
Output: [2,3,6,7,1,5,4]
```

### Constraints:

- The number of nodes in the linked list is in the range [0, $10^4$].
- -$10^6$ <= Node.val <= $10^6$

### Related Topics

- Linked List
  
---

# 解題方向

Two Pointers (previous and current pointer)

簡單來說就是不斷地將 `current pointer` 指到的下一個節點搬到 `previous pointer` 的下一個節點，然後兩個指針往下移動直到結束

用講的可能有點抽象，直接圖解：

```
step.0 原始資料

      v head     v cur
    +---+      +---+      +---+      +---+      +---+      +---+            
    | 1 | ---> | 2 | ---> | 3 | ---> | 4 | ---> | 5 | ---> | 6 | ---> none  
    +---+      +---+      +---+      +---+      +---+      +---+            
      ^ pre      ^ tmp

step.1 將 previous pointer 的下一個位置指向 current pointer 的下一個位置

      v pre     v cur
    +---+      +---+      +---+      +---+      +---+      +---+            
    | 1 | ---X | 2 | ---> | 3 | ---> | 4 | ---> | 5 | ---> | 6 | ---> none  
    +---+  |   +---+      +---+      +---+      +---+      +---+            
           |     ^ tmp      ^
           |                |
           └----------------┘

step.2 將 current pointer 的下一個位置指向 current pointer 的往後兩個位置

                      ┌----------------┐
                      |                |
      v pre     v cur |                v
    +---+      +---+  |   +---+      +---+      +---+      +---+            
    | 1 | ---X | 2 | ---X | 3 | ---> | 4 | ---> | 5 | ---> | 6 | ---> none  
    +---+  |   +---+      +---+      +---+      +---+      +---+            
           |     ^ tmp      ^
           |                |
           └----------------┘

step.3 把 previous pointer 新指向的下一個位置的下一個位置指回 tmp (好饒舌QQ)

                      ┌----------------┐
                      |                |
      v pre     v cur |                v
    +---+      +---+  |   +---+      +---+      +---+      +---+            
    | 1 | ---X | 2 | ---X | 3 | ---X | 4 | ---> | 5 | ---> | 6 | ---> none  
    +---+  |   +---+      +---+  |   +---+      +---+      +---+            
           |     ^ tmp      ^    |
           |     |          |    |
           └-----+----------┘    |
                 └---------------┘

step.4 把 previous pointer 跟 current pointer 都移動到各自新的下一個位置

                      ┌----------------┐
                      |                |
                      |     v pre      v
    +---+      +---+  |   +---+      +---+      +---+      +---+            
    | 1 | ---X | 2 | ---X | 3 | ---X | 4 | ---> | 5 | ---> | 6 | ---> none  
    +---+  |   +---+      +---+  |   +---+      +---+      +---+            
           |     ^ tmp      ^    |     ^ cur
           |     |          |    |
           └-----+----------┘    |
                 └---------------┘

把示意圖按新的排序整理一下

      v head     v pre                 v cur
    +---+      +---+      +---+      +---+      +---+      +---+            
    | 1 | ---> | 3 | ---> | 2 | ---> | 4 | ---> | 5 | ---> | 6 | ---> none  
    +---+      +---+      +---+      +---+      +---+      +---+            
```