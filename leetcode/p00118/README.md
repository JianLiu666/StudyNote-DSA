[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/pascals-triangle/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 118. Pascal's Triangle

Given an integer `numRows`, return the first numRows of **Pascal's triangle**.

In **Pascal's triangle**, each number is the sum of the two numbers directly above it as shown:

![image](./image/PascalTriangleAnimated.gif)

### Example 1:

```
Input: numRows = 5
Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
```

### Example 2:

```
Input: numRows = 1
Output: [[1]]
```

### Constraints:

- `1 <= numRows <= 30`

### Related Topics:

- Array
- Dynamic Programming

---

# 解題方向

Brute-force, 就是單純將每層結果當成下一層的輸入傳遞進去計算