[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/middle-of-the-linked-list/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 876. Middle of the Linked List

Given the `head` of a singly linked list, return the middle node of the linked list.

If there are two middle nodes, return **the second middle** node.

### Example 1:

![image](./image/lc-midlist1.jpeg)

```
Input: head = [1,2,3,4,5]
Output: [3,4,5]

Explanation:
 - The middle node of the list is node 3.
```

### Example 2:

![image](./image/lc-midlist2.jpeg)

```
Input: head = [1,2,3,4,5,6]
Output: [4,5,6]

Explanation:
 - Since the list has two middle nodes with values 3 and 4, we return the second one.
```

### Constraints:

- The number of nodes in the list is in the range `[1, 100]`.
- `1 <= Node.val <= 100`

### Related Topics

- Linked List
- Two Pointers
  
---

# 解題方向

基本 `Linked List` 暖身題，透過 Two Pointers (slow and fast pointer) 解決

- 給定 `fast pointer` 的移動速度是 `slow pointer` 的兩倍，當 `fast pointer` 走到終點時 `slow pointer` 的位置就恰好會在 `Linked List` 的中間