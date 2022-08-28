[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/rotate-image/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 48. Rotate Image

You are given an `n x n` 2D `matrix` representing an image, rotate the image by **90** degrees (clockwise).

You have to rotate the image [in-place](https://en.wikipedia.org/wiki/In-place_algorithm), which means you have to modify the input 2D matrix directly. **DO NOT** allocate another 2D matrix and do the rotation.

### Example 1:

![image](./image/mat1.jpeg)

```
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [[7,4,1],[8,5,2],[9,6,3]]
```

### Example 2:

![image](./image/mat2.jpeg)

```
Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
```

### Constraints:

- `n == matrix.length == matrix[i].length`
- `1 <= n <= 20`
- `-1000 <= matrix[i][j] <= 1000`

### Related Topics

- Array
- Math
- Matrix
  
---

# 解題方向

### Solved using Linear Algebra concept

在線性代數中這個方法叫做 **Transpose and Reflect** *(老師抱歉...)*

懶人包請看這題的 [Solution](https://leetcode.com/problems/rotate-image/solution/)

### Solved using Brute Force concept

示意圖如下:

```
+---+---+---+---+
| A |   |   | B |
+---+---+---+---+
|   |   |   |   |
+---+---+---+---+
|   |   |   |   |
+---+---+---+---+
| D |   |   | C |
+---+---+---+---+
```

順序:
1. `B, tmp = A, B`
2. `C, tmp = tmp, C`
3. `D, tmp = tmp, D`
4. `A = tmp`

不斷交換直到整個 `matrix` 都尋訪過一遍