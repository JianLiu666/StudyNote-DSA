[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/rotate-list/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 61. Rotate List

Given the `head` of a linked list, rotate the list to the right by `k` places.

### Example 1:

![image](./image/rotate1.jpeg)

```
Input: head = [1,2,3,4,5], k = 2
Output: [4,5,1,2,3]
```

### Example 2:

![image](./image/rotate2.jpeg)

```
Input: head = [0,1,2], k = 4
Output: [2,0,1]
```

### Constraints:

- The number of nodes in the list is in the range `[0, 500]`.
- `-100 <= Node.val <= 100`
- 0 <= k <= 2 * $10^9$

### Related Topics

- Linked List
- Two Pointers
  
---

# 解題方向

根據題目可以發現只要旋轉次數跟 `Linked List` 長度相同時就跟沒轉一樣，剩下的就只是一般的 Two Pointers 做法

- i.e., `rotate_point = k % len(linked_list)`