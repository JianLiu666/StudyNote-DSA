[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/spiral-matrix/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 54. Spiral Matrix

Given an `m x n` `matrix`, return all elements of the `matrix` in spiral order.

### Example 1:

![image](./image/spiral1.jpeg)

```
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]
```

### Example 2:

![image](./image/spiral.jpeg)

```
Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]
```

### Constraints:

- `m == matrix.length`
- `n == matrix[i].length`
- `1 <= m, n <= 10`
- `-100 <= matrix[i][j] <= 100`

### Related Topics

- Array
- Matrix
- Simulation
  
---

# 解題方向

基本的 `2D Array` 暖身題，只要在螺旋循環時確保上下邊界，不要走到已經訪問過的元素即可。