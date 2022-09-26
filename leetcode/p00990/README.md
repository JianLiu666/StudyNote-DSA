[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/satisfiability-of-equality-equations/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 990. Satisfiability of Equality Equations

You are given an array of strings `equations` that represent relationships between variables where each string `equations[i]` is of length `4` and takes one of two different forms: " x~i~ == y~i~ " or " x~i~ != y~i~ ".Here, x~i~ and y~i~ are lowercase letters (not necessarily different) that represent one-letter variable names.

Return `true` if it is possible to assign integers to variable names so as to satisfy all the given equations, or `false` otherwise.

### Example 1:

```
Input: equations = ["a==b","b!=a"]
Output: false

Explanation:
 - If we assign say, a = 1 and b = 1, then the first equation is satisfied, but not the second.
 - There is no way to assign the variables to satisfy both equations.
```

### Example 2:

```
Input: equations = ["b==a","a==b"]
Output: true

Explanation:
 -  We could assign a = 1 and b = 1 to satisfy both equations.
```

### Constraints:

- `1 <= equations.length <= 500`
- `equations[i].length == 4`
- `equations[i][0]` is a lowercase letter.
- `equations[i][1]` is either `'='` or `'!'`.
- `equations[i][2]` is `'='`.
- `equations[i][3]` is a lowercase letter.

### Related Topics

- Array
- String
- Union Find
- Graph
  
---

# 解題方向

### Solved using Union Find concept

暖身題，第一次遍歷 `equations` 建立好併查集後，在對 `equations` 遍歷一次檢查群組關係即可