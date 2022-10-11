[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/triangle/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 120. Triangle

Given a `triangle` array, return the minimum path sum from top to bottom.

For each step, you may move to an adjacent number of the row below. More formally, if you are on index `i` on the current row, you may move to either index `i` or index `i + 1` on the next row.

### Example 1:

```
Input: triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
Output: 11

Explanation: The triangle looks like:
    2
   3 4
  6 5 7
 4 1 8 3
 The minimum path sum from top to bottom is 2 + 3 + 5 + 1 = 11 (underlined above).
```

### Example 2:

```
Input: triangle = [[-10]]
Output: -10
```

### Constraints:

- `1 <= triangle.length <= 200`
- `triangle[0].length == 1`
- `triangle[i].length == triangle[i - 1].length + 1`
- $-10^4$ <= `triangle[i][j]` <= $10^4$

#### Follow up: 

Could you do this using only `O(n)` extra space, where `n` is the total number of rows in the triangle?

### Related Topics

- Array
- Dynamic Programming
  
---

# 解題方向

用 Bottom-up 的想法從起點一路往後推到終點，每次的目標都找出當前加總的最小值