[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/sort-array-by-parity/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 905. Sort Array By Parity

Given an integer array `nums`, move all the even integers at the beginning of the array followed by all the odd integers.

Return **any array** that satisfies this condition.

### Example 1:

```
Input: nums = [3,1,2,4]
Output: [2,4,3,1]

Explanation:
 - The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.
```

### Example 2:

```
Input: nums = [0]
Output: [0]
```

### Constraints:

 - `1 <= nums.length <= 5000`
 - `0 <= nums[i] <= 5000`

### Related Topics

- Array
- Two Pointers
- Sorting
  
---

# 解題方向

Two Pointers (head and tail pointer)

從頭尾兩端開始搜尋，只要兩邊都滿足條件就交換，直到頭尾指針相遇