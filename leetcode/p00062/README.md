[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/unique-paths/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 62. Unique Paths

There is a robot on an `m x n` grid. The robot is initially located at the **top-left corner** (i.e., `grid[0][0]`). The robot tries to move to the **bottom-right corner** (i.e., `grid[m - 1][n - 1]`). The robot can only move either down or right at any point in time.

Given the two integers `m` and `n`, return the number of possible unique paths that the robot can take to reach the bottom-right corner.

The test cases are generated so that the answer will be less than or equal to 2 * $10^9$.

### Example 1:

![image](./image/robot_maze.png)

```
Input: m = 3, n = 7
Output: 28
```

### Example 2:

```
Input: m = 3, n = 2
Output: 3

Explanation: 
 - From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
   1. Right -> Down -> Down
   2. Down -> Down -> Right
   3. Down -> Right -> Down
```

### Constraints:

- `1 <= m, n <= 100`

### Related Topics

- Math
- Dynamic Programming
- Combinatorics
  
---

# 解題方向

`Dynamic Programming` 暖身題，按照題目描述從左上角出發，模擬機器人行走方式不斷更新 `matrix[row][col]` 的元素即可

