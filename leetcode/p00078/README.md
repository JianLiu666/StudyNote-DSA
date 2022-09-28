[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/subsets/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 78. Subsets

Given an integer array `nums` of **unique** elements, return all possible subsets (the power set).

The solution set **must not** contain duplicate subsets. Return the solution in **any order**.

### Example 1:

```
Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
```

### Example 2:

```
Input: nums = [0]
Output: [[],[0]]
```

### Constraints:

- `1 <= nums.length <= 10`
- `-10 <= nums[i] <= 10`
- All the numbers of `nums` are **unique**.

### Related Topics

- Array
- Backtracking
- Bit Manipulation
  
---

# 解題方向

`Backtracking` 暖身題，每個元素分別都處理拿與不拿的情境即可
