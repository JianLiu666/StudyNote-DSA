[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/diagonal-traverse/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 498. Diagonal Traverse

Given an `m x n` matrix mat, return an array of all the elements of the array in a diagonal order.

### Example 1:

![image](./image/diag1-grid.jpeg)

```
Input: mat = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,4,7,5,3,6,8,9]
```

### Example 2:

```
Input: mat = [[1,2],[3,4]]
Output: [1,2,3,4]
```

### Constraints:

- `m == mat.length`
- `n == mat[i].length`
- 1 <= m, n <= $10^4$
- 1 <= m * n <= $10^4$
- -$10^5$ <= mat[i][j] <= $10^5$

### Related Topics

- Array
- Matrix
- Simulation
  
---

# 解題方向

應該可以算是基本的 `2D Array` 暖身題，只要注意好邊界條件的判斷，確實轉換方向即可