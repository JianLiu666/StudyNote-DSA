[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/sudoku-solver/)
![image](https://img.shields.io/badge/Difficulty-Hard-red)

---

# 37. Sudoku Solver

Write a program to solve a Sudoku puzzle by filling the empty cells.

A sudoku solution must satisfy **all of the following rules**:

1. Each of the digits `1-9` must occur exactly once in each row.
2. Each of the digits `1-9` must occur exactly once in each column.
3. Each of the digits `1-9` must occur exactly once in each of the 9 `3x3` sub-boxes of the grid.

The `'.'` character indicates empty cells.

### Example 1:

![image](./image/Sudoku-by-L2G-20050714.svg.webp)

```
Input: board = [
    ["5","3",".",".","7",".",".",".","."],
    ["6",".",".","1","9","5",".",".","."],
    [".","9","8",".",".",".",".","6","."],
    ["8",".",".",".","6",".",".",".","3"],
    ["4",".",".","8",".","3",".",".","1"],
    ["7",".",".",".","2",".",".",".","6"],
    [".","6",".",".",".",".","2","8","."],
    [".",".",".","4","1","9",".",".","5"],
    [".",".",".",".","8",".",".","7","9"]]

Output: [
    ["5","3","4","6","7","8","9","1","2"],
    ["6","7","2","1","9","5","3","4","8"],
    ["1","9","8","3","4","2","5","6","7"],
    ["8","5","9","7","6","1","4","2","3"],
    ["4","2","6","8","5","3","7","9","1"],
    ["7","1","3","9","2","4","8","5","6"],
    ["9","6","1","5","3","7","2","8","4"],
    ["2","8","7","4","1","9","6","3","5"],
    ["3","4","5","2","8","6","1","7","9"]]

Explanation: 
 - The input board is shown above and the only valid solution is shown below:
```

![image](./image/250px-Sudoku-by-L2G-20050714_solution.svg.png)

### Constraints:

 - `board.length == 9`
 - `board[i].length == 9`
 - `board[i][j]` is a digit or `'.'`.
 - It is **guaranteed** that the input board has only one solution.

### Related Topics

- Array
- Backtracking
- Matrix
  
---

# 解題方向

經典的 `Backtracking` 練習題，只要注意數獨特有規則並維護好數字更新的 pattern 即可

