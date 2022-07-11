[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/remove-nth-node-from-end-of-list/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 19. Remove Nth Node From End of List

Given the `head` of a linked list, remove the $n^{th}$ node from the end of the list and return its head.

### Example 1:

![image](./image/remove_ex1.jpeg)

```
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
```

### Example 2:

```
Input: head = [1], n = 1
Output: []
```

### Example 3:

```
Input: head = [1,2], n = 1
Output: [1]
```

### Constraints:

- The number of nodes in the list is `sz`.
- `1 <= sz <= 30`
- `0 <= Node.val <= 100`
- `1 <= n <= sz`

#### Follow up:

Could you do this in one pass?

### Related Topics

- Linked List
- Two Pointers
  
---

# 解題方向

Two Pointers : (first and second pointer)

先讓 `first pointer` 往前移動 `n` 步後，`second pointer` 才開始往下移動

等到 `first pointer` 走到尾巴時， `second pointer` 目前停留位置的下一個節點就是我們要移除的節點了

圖解：

```
n = 2

1. first pointer 先出發，往前移動兩步

              v first pointer
    1 -> 2 -> 3 -> 4 -> 5
    ^ second pointer

2. first pointer 與 second pointer 開始同時等速移動，直到 first pointer 移動到末端

                        v first pointer
    1 -> 2 -> 3 -> 4 -> 5
              ^ second pointer
```