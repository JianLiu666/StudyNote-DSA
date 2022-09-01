[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/set-matrix-zeroes/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 73. Set Matrix Zeroes

Given an `m x n` integer matrix `matrix`, if an element is `0`, set its entire row and column to `0`'s.

You must do it [in place](https://en.wikipedia.org/wiki/In-place_algorithm).

### Example 1:

![image](./image/mat1.jpeg)

```
Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
Output: [[1,0,1],[0,0,0],[1,0,1]]
```

### Example 2:

![image](./image/mat2.jpeg)

```
Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
```

### Constraints:

- `m == matrix.length`
- `n == matrix[0].length`
- `1 <= m, n <= 200`
- -$2^{31}$ <= matrix[i][j] <= $2^{31}$ - 1

#### Follow up:

- A straightforward solution using `O(mn)` space is probably a bad idea.
- A simple improvement uses `O(m + n)` space, but still not the best solution.
- Could you devise a constant space solution?

### Related Topics

- Array
- Hash Table
- Matrix
  
---

# 解題方向

先在 LeetCodeCN 上看到的題目，詳細思路可以參考官方提供的[解法](https://leetcode.cn/problems/zero-matrix-lcci/solution/ling-ju-zhen-by-leetcode-solution-7ogg/)

`Space Complexity = O(1)` 的做法有點有趣，詳細可以直接看程式碼，核心概念就是把暫存資料放在同一個 `matrix` 的最前面來處理