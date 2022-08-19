[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/remove-duplicates-from-sorted-list/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 83. Remove Duplicates from Sorted List

Given the `head` of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list **sorted** as well.

### Example 1:

![image](./image/list1.jpeg)

```
Input: head = [1,1,2]
Output: [1,2]
```

### Example 2:

![image](./image/list2.jpeg)

```
Input: head = [1,1,2,3,3]
Output: [1,2,3]
```

### Constraints:

- The number of nodes in the list is in the range `[0, 300]`.
- `-100 <= Node.val <= 100`
- The list is guaranteed to be **sorted** in ascending order.

### Related Topics

- Linked List
  
---

# 解題方向

`Singly Linked List` 暖身題，只要控制好 `current` 跟 `next` 兩個 `pointers`，如果 `current.val == next.val` 相等時 `next` 就繼續往前尋訪

直到 `current` 與 `next` 不同就能一次移除中間這段重複節點
