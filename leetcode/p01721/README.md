[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/swapping-nodes-in-a-linked-list/description/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 1721. Swapping Nodes in a Linked List

You are given the `head` of a linked list, and an integer `k`.

Return the head of the linked list after swapping the values of the $k^{th}$ node from the beginning and the $k^{th}$ node from the end (the list is 1-indexed).

### Example 1:

![image](./image/linked1.jpg) 

```
Input: head = [1,2,3,4,5], k = 2
Output: [1,4,3,2,5]
```

### Example 2:

```
Input: head = [7,9,6,6,7,8,3,0,9,5], k = 5
Output: [7,9,6,6,8,7,3,0,9,5]
```

### Constraints:

- The number of nodes in the list is `n`.
- 1 <= k <= n <= $10^5$
- `0 <= Node.val <= 100`

### Related Topics

- Linked List
- Two Pointers
  
---

# 解題方向

Linked List 暖身題，透過 Two Pointer 方式找到 right node 就可以跟 left node 進行交換

```
以 Example 1 為例:

Step.1 以 for loop 定位好 left node, cursor node (用來定位 right node)

        left, cursor
           |
           v
+---+    +---+    +---+    +---+    +---+
| 1 | -> | 2 | -> | 3 | -> | 4 | -> | 5 |
+---+    +---+    +---+    +---+    +---+

Step.2 將 right node 放在 head 的位置 (i.e. right 與 cursor 的距離就表示 right 到 linked list end 的距離)

          left, cursor
           |
           v
+---+    +---+    +---+    +---+    +---+
| 1 | -> | 2 | -> | 3 | -> | 4 | -> | 5 |
+---+    +---+    +---+    +---+    +---+
  ^
  |
 right

Step.3 開始 while loop 不斷移動 right 與 curosr 直到 cursor 走到 linked list end

          left                       cursor
           |                          |
           v                          v
+---+    +---+    +---+    +---+    +---+
| 1 | -> | 2 | -> | 3 | -> | 4 | -> | 5 |
+---+    +---+    +---+    +---+    +---+
                             ^
                             |
                            right

Step.4 交換 left 與 right 所在位置的 value

          left                  
           |                    
           v                    
+---+    +---+    +---+    +---+    +---+
| 1 | -> | 4 | -> | 3 | -> | 2 | -> | 5 |
+---+    +---+    +---+    +---+    +---+
                             ^
                             |
                            right
```