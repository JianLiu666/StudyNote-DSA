[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/remove-linked-list-elements/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 203. Remove Linked List Elements

Given the `head` of a linked list and an integer `val`, remove all the nodes of the linked list that has `Node.val == val`, and return the new head.

![image](./image/removelinked-list.jpeg)

### Example 1:

```
Input: head = [1,2,6,3,4,5,6], val = 6
Output: [1,2,3,4,5]
```

### Example 2:

```
Input: head = [], val = 1
Output: []
```

### Example 3:

```
Input: head = [7,7,7,7], val = 7
Output: []
```

### Constraints:

- The number of nodes in the list is in the range [0, $10^4$].
- `1 <= Node.val <= 50`
- `0 <= val <= 50`

### Related Topics

- Linked List
- Recursion
  
---

# 解題方向

基本 Linked List 練習題