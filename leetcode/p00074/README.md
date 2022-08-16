[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/search-a-2d-matrix/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)
---

# 74. Search a 2D Matrix

Write an efficient algorithm that searches for a value `target` in an `m x n` integer matrix `matrix`. This matrix has the following properties:

- Integers in each row are sorted from left to right.
- The first integer of each row is greater than the last integer of the previous row.

### Example 1:

![image](./image/mat.jpeg)

```
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true
```

### Example 2:

![image](./image/mat2.jpeg)

```
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
Output: false
```

### Constraints:

- `m == matrix.length`
- `n == matrix[i].length`
- `1 <= m, n <= 100`
- -$10^4$ <= matrix[i][j], target <= $10^4$

### Related Topics

- Array
- Binary Search
- Matrix
  
---

# 解題方向

依照題意可以確定的是，這是一個已經排序過的 `matrix`，因此從正確的角度切入，就可以確保每一次的迭代都能往正確的方向收斂，即：

- 從 `bottom-left` 切入，只要 `target` 小於 `matrix[col][row]` 就只能往上移動，反之則往右移動
- 從 `top-right` 切入，只要 `target` 小於 `matrix[col][row]` 時就只能往左移動，反之則往下移動

