[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/reverse-linked-list-ii/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 92. Reverse Linked List II

Given the `head` of a singly linked list and two integers `left` and `right` where `left <= right`, reverse the nodes of the list from position `left` to position `right`, and return the reversed list. 

### Example 1:

![image](./image/rev2ex2.jpeg)

```
Input: head = [1,2,3,4,5], left = 2, right = 4
Output: [1,4,3,2,5]
```

### Example 2:

```
Input: head = [5], left = 1, right = 1
Output: [5]
```

### Constraints:

- The number of nodes in the list is `n`.
- `1 <= n <= 500`
- `-500 <= Node.val <= 500`
- `1 <= left <= right <= n`
 
#### Follow up: 

Could you do it in one pass?

### Related Topics

- Linked List
  
---

# 解題方向

本質還是跟一般的反轉 `Linked List` 問題一樣，只要透過幾個 `cursors` 定位好再反轉時需要操作的 `nodes` 即可

如圖所示：

```
step.0:

    head       left                 right
     v          v                     v
   +---+      +---+      +---+      +---+      +---+
   | 1 | ---> | 2 | ---> | 3 | ---> | 4 | ---> | 5 |
   +---+      +---+      +---+      +---+      +---+
     ^
  current

step.1: 當 current 尋訪到 left 時，將需要的 cursors 都定位好，以及開始準備反轉 nodes

  current 尋訪                                              標記 cursors

    head       left                 right            reverse_head_prev   reverse_tail
     v          v                     v                       v          v
   +---+      +---+      +---+      +---+      +---+        +---+      +---+      +---+      +---+      +---+
   | 1 | ---> | 2 | ---> | 3 | ---> | 4 | ---> | 5 |   =>   | 1 | ---> | 2 | ---> | 3 | ---> | 4 | ---> | 5 |
   +---+      +---+      +---+      +---+      +---+        +---+      +---+      +---+      +---+      +---+
     ^          ^                                             ^          ^
  previous   current                                       previous   current

  反轉流程
                      1. set tmp
                           v
   +---+      +---+      +---+      +---+      +---+
   | 1 | ---> | 2 | ---X | 3 | ---> | 4 | ---> | 5 |
   +---+      +---+  |   +---+      +---+      +---+
                ^    |     ^
    3. move previous |   4. move current
       to here       |      to here
                     v
               2. set to none

step.2: current 介於 left 與 right 之間時，只要考慮怎麼反轉指針即可

                2. reverse pointer
                ┌---------------┐  1. set tmp
                v               |     v
   +---+      +---+      +---+  |   +---+      +---+
   | 1 | ---> | 2 |      | 3 | ---X | 4 | ---> | 5 |
   +---+      +---+      +---+      +---+      +---+
                |          ^          ^     
                v   3. move previous  4. move current
               none    to here           to here

step.3: current 尋訪到 right 時，進行收尾

                           2. reverse pointer
        3.redirect pointer ┌---------------┐
          to here          |               |
     ┌---------------------+----------┐    |     
     |                     v          v    |     
   +---+      +---+      +---+      +---+  |   +---+        +---+      +---+      +---+      +---+      +---+
   | 1 | ---> | 2 | <--- | 3 | <--- | 4 | ---X | 5 |   =>   | 1 | ---> | 4 | ---> | 3 | ---> | 2 | ---> | 5 |
   +---+      +---+      +---+      +---+      +---+        +---+      +---+      +---+      +---+      +---+
                |                                ^
                ├--------------------------------┘
                |  1. redirect pointer to here
                X
               none
```