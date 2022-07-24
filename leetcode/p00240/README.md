[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/search-a-2d-matrix-ii/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 240. Search a 2D Matrix II

Write an efficient algorithm that searches for a value `target` in an `m x n` integer matrix `matrix`. This matrix has the following properties:

- Integers in each row are sorted in ascending from left to right.
- Integers in each column are sorted in ascending from top to bottom.

### Example 1:

![image](./image/searchgrid2.jpeg)

```
Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
Output: true
```

### Example 2:

![image](./image/searchgrid.jpeg)

```
Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
Output: false
```

### Constraints:

- `m == matrix.length`
- `n == matrix[i].length`
- `1 <= n, m <= 300`
- -$10^9$ <= matrix[i][j] <= $10^9$
- All the integers in each row are **sorted** in ascending order.
- All the integers in each column are **sorted** in ascending order.
- -$10^9$ <= target <= $10^9$

### Related Topics

- Array
- Binary Search
- Divide and Conquer
- Matrix
  
---

# 解題方向

### Solved using binary search algorithm

最直覺的想法，只要對每一個 `row array` 都做一遍 Binary Search，就可以找到答案

依照題意，在查找之前可以先檢查 `array` 的頭尾確認 `target` 是否落在這個區間，但本質上來說的時間複雜度還是 `O(nlogn)`

### Solved using divide and conquer concept

依照題意可以確定的是，這是一個已經排序過的 `matrix`，因此從正確的角度切入，就可以確保每一次的迭代都能往正確的方向收斂，即：

- 從 `bottom-left` 切入，只要 target 小於 `matrix[col][row]` 就只能往上移動，反之則往右移動
- 從 `top-right` 切入，只要 target 小於 `matrix[col][row]` 時就只能往左移動，反之則往下移動

最壞情況為 `target` 落在切入點的對角線最遠距離，時間複雜度為 `O(m+n)`