[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/pascals-triangle-ii/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 119. Pascal's Triangle II

Given an integer `rowIndex`, return the $\text{rowIndex}^{th}$ (**0-indexed**) row of the **Pascal's triangle**.

In **Pascal's triangle**, each number is the sum of the two numbers directly above it as shown:

![image](./image/PascalTriangleAnimated.gif)

### Example 1:

```
Input: rowIndex = 3
Output: [1,3,3,1]
```

### Example 2:

```
Input: rowIndex = 0
Output: [1]
```

### Example 3:

```
Input: rowIndex = 1
Output: [1,1]
```

### Constraints:

- `0 <= rowIndex <= 33`
 

#### Follow up:

Could you optimize your algorithm to use only `O(rowIndex)` extra space?

### Related Topics:

- Array
- Dynamic Programming

---

# 解題方向

題目限制能用的 extra space 只有 `O(rowIndex)`，想要用 `Hash Table` 是不可能的了，那遞迴估計也可以放棄了

這題的思路是想辦法用 `in-place` 的角度切入，反覆在同一個陣列裡面重複疊加數字上去