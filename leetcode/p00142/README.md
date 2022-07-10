[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/linked-list-cycle-ii/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 142. Linked List Cycle II

Given the `head` of a linked list, return the node where the cycle begins. If there is no cycle, return `null`.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the `next` pointer. Internally, `pos` is used to denote the index of the node that tail's `next` pointer is connected to (**0-indexed**). It is `-1` if there is no cycle. **Note that** `pos` **is not passed as a parameter**.

**Do not modify** the linked list.

### Example 1:

![image](./image/circularlinkedlist_test1.png)

```
Input: head = [3,2,0,-4], pos = 1
Output: tail connects to node index 1

Explanation:
 - There is a cycle in the linked list, where tail connects to the second node.
```

### Example 2:

![image](./image/circularlinkedlist_test2.png)

```
Input: head = [1,2], pos = 0
Output: tail connects to node index 0

Explanation:
 - There is a cycle in the linked list, where tail connects to the first node.
```

### Example 3:

```
Input: head = [1], pos = -1
Output: no cycle

Explanation:
 - There is no cycle in the linked list.
```

### Constraints:

- The number of the nodes in the list is in the range [0, $10^4$].
- -$10^5$ <= Node.val <= $10^5$
- `pos` is `-1` or a **valid index** in the linked-list.
 
#### Follow up:

Can you solve it using `O(1)` (i.e. constant) memory?

### Related Topics

- Hash Table
- Linked List
- Two Pointers
  
---

# 解題方向

### Solved using data structure

將遍歷過的 `ListNode` 保存在 `Hash map`，每次移動時檢查是否曾出現過。

### Solved using two pointer concept

老規矩，直接圖解：

```
         +---+    +---+    +---+    +---+    +---+    +---+
 head -> | 1 | -> | 2 | -> | 3 | -> | 4 | -> | 5 | -> | 6 |
         +---+    +---+    +---+    +---+    +---+    +---+
                                      ^                 |
                                      |                 v
                                    +---+    +---+    +---+
                                    | 9 | <- | 8 | <- | 7 |
                                    +---+    +---+    +---+
```

現在我們看看使用 Two Pointers 策略之後的結果：

- 給定 `fast pointer` 的移動速度是 `slow pointer` 的兩倍

```
假設 s0, f0 分別表示 slow pointer 與 fast pointer 的初始位置，即：

          f0
          s0
         +---+    +---+    +---+    +---+    +---+    +---+
 head -> | 1 | -> | 2 | -> | 3 | -> | 4 | -> | 5 | -> | 6 |
         +---+    +---+    +---+    +---+    +---+    +---+
                                      ^                 |
                                      |                 v
                                    +---+    +---+    +---+
                                    | 9 | <- | 8 | <- | 7 |
                                    +---+    +---+    +---+

經過第一次移動之後，slow pointer 與 fast pointer 目前的所在位置，即：

          f0
          s0       s1       f1
         +---+    +---+    +---+    +---+    +---+    +---+
 head -> | 1 | -> | 2 | -> | 3 | -> | 4 | -> | 5 | -> | 6 |
         +---+    +---+    +---+    +---+    +---+    +---+
                                      ^                 |
                                      |                 v
                                    +---+    +---+    +---+
                                    | 9 | <- | 8 | <- | 7 |
                                    +---+    +---+    +---+

現在我們將時間快轉到 slow pointer 與 fast pointer 相遇時的結果，如下：

                                              f5
          f0                s2                s4
          s0       s1       f1       s3       f2       s5
         +---+    +---+    +---+    +---+    +---+    +---+
 head -> | 1 | -> | 2 | -> | 3 | -> | 4 | -> | 5 | -> | 6 |
         +---+    +---+    +---+    +---+    +---+    +---+
                                      ^                 |
                                      |                 v
                                    +---+    +---+    +---+
                                    | 9 | <- | 8 | <- | 7 |
                                    +---+    +---+    +---+
                                     f4                f3
                                                       s6
                                                       f6
```

如上圖所示，`node7` 即 `slow ponter` 與 `fast pointer` 的相遇點

現在回到題目，我們不只要確認 Linked List 是否循環，還要找出循環的起點，要做到這件事情，我們必須先定義幾個變數以利我們後續推導：

- `L1` : 表示從 `head` 到循環點的路徑，即 `1 -> 2 -> 3 -> 4`
- `L2` : 表示從循環點到 `slow pointer` 與 `fast pointer` 相遇點的路徑，即 `4 -> 5 -> 6 -> 7`
- `C` : 表示形成循環的路徑，即 `4 -> 5 -> 6 -> 7 -> 8 -> 9`

有了這幾個變數，我們可以將 `slow pointer` 的整個移動軌跡表示成 `L1 + L2`，即 `1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7`

還記得我們在一開始就先定義 `fast pointer` 的移動速度是 `slow pointer` 的兩倍吧？因此我們可以把 `fast pointer` 的移動軌跡表示成 `2 * (L1 + L2)`

重點來了，`fast pointer` 有一半的移動距離(i.e., `L1 + L2`) 是為了追上 `slow pointer` 而多走的距離，換句話說就是在閉環內打轉的距離，即 `L1 + L2 = n * C`
- 這裡的 `n` 表示在閉環內重複循環的次數
- 把等式移動一下後變成 `L1 = nC - L2`

有趣的事情來了，對照一下上圖後發現 `slow pointer` 目前停留的位置恰恰好就正在 `nC - L2` 的距離

換句話說，如果我們現在在給定一個新的 `slow pointer'` 從 `head` 開始出發，當 `slow pointer` 與 `slow pointer'` 兩個指針相遇時的位置恰好就會在循環點上，如下圖所示：

```
                                     s'3
          s'0      s'1      s'2      s3
         +---+    +---+    +---+    +---+    +---+    +---+
 head -> | 1 | -> | 2 | -> | 3 | -> | 4 | -> | 5 | -> | 6 |
         +---+    +---+    +---+    +---+    +---+    +---+
                                      ^                 |
                                      |                 v
                                    +---+    +---+    +---+
                                    | 9 | <- | 8 | <- | 7 |
                                    +---+    +---+    +---+
                                     s2       s1       s0
```