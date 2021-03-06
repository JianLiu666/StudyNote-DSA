[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/add-two-numbers/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 2. Add Two Numbers

You are given two **non-empty** linked lists representing two non-negative integers. The digits are stored in **reverse order**, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

### Example 1:

![image](./image/addtwonumber1.jpeg)

```
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]

Explanation:
 - 342 + 465 = 807.
```

### Example 2:

```
Input: l1 = [0], l2 = [0]
Output: [0]
```

### Example 3:

```
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]
```

### Constraints:

- The number of nodes in each linked list is in the range `[1, 100]`.
- `0 <= Node.val <= 9`
- It is guaranteed that the list represents a number that does not have leading zeros.

### Related Topics

- Linked List
- Math
- Recursion
  
---

# 解題方向

基本的 Linked List 暖身題，透過兩個指針來組成的 Linked List 即可