[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/palindrome-linked-list/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 234. Palindrome Linked List

Given the `head` of a singly linked list, return `true` if it is a palindrome.

### Example 1:

![image](./image/pal1linked-list.jpeg)

```
Input: head = [1,2,2,1]
Output: true
```

### Example 2:

![image](./image/pal2linked-list.jpeg)

```
Input: head = [1,2]
Output: false
```

### Constraints:

- The number of nodes in the list is in the range [1, $10^5$].
- `0 <= Node.val <= 9`

#### Follow up:

Could you do it in `O(n)` time and `O(1)` space?

### Related Topics

- Linked List
- Two Pointers
- Stack
- Recursion
  
---

# 解題方向

Two Pointers (slow and fast pointer)
- `fast pointer` 的移動速度定為 `slow pointer` 的兩倍
- 當 `fast pointer` 走到終點時，`slow pointer` 的位置會恰好在整個 linked list 的中間位置

如果題目沒有對空間有額外要求的話，`slow pointer` 在移動的過程中可以將資料寫進 `Stack` 內，等到 `fast pointer` 走到終點後，接下來的路程在逐一從 `Stack` 中把資料拿出來跟現在指針位置上的資料做比對。

但礙於題目限制所以不能這樣處理，但題目沒有說不能破壞原始的鏈路結構，所以我們可以從 `slow pointer` 的停止位置開始依序反轉第二段的 `linked list`

如此一來，只要等 `slow pointer` 也走到終點後，就會獲得一個已經反轉過的 `linked list`，就可以將兩個 `linked lists` 做比較確認是否相同。