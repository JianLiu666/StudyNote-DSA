[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/evaluate-division/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 399. Evaluate Division

You are given an array of variable pairs `equations` and an array of real numbers `values`, where equations[i] = [A~i~, B~i~] and `values[i]` represent the equation A~i~ / B~i~ = values[i]. Each A~i~ or B~i~ is a string that represents a single variable.

You are also given some `queries`, where queries[j] = [C~j~, D~j~] represents the $j^{th}$ query where you must find the answer for C~j~ / D~j~ = ?.

Return the answers to all queries. If a single answer cannot be determined, return `-1.0`.

**Note:** The input is always valid. You may assume that evaluating the queries will not result in division by zero and that there is no contradiction.

### Example 1:

```
Input: equations = [["a","b"],["b","c"]], values = [2.0,3.0], queries = [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
Output: [6.00000,0.50000,-1.00000,1.00000,-1.00000]

Explanation: 
 - Given: a / b = 2.0, b / c = 3.0
 - queries are: a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ?
 - return: [6.0, 0.5, -1.0, 1.0, -1.0 ]
```

### Example 2:

```
Input: equations = [["a","b"],["b","c"],["bc","cd"]], values = [1.5,2.5,5.0], queries = [["a","c"],["c","b"],["bc","cd"],["cd","bc"]]
Output: [3.75000,0.40000,5.00000,0.20000]
```

### Example 3:

```
Input: equations = [["a","b"]], values = [0.5], queries = [["a","b"],["b","a"],["a","c"],["x","y"]]
Output: [0.50000,2.00000,-1.00000,-1.00000]
```

### Constraints:

- `1 <= equations.length <= 20`
- `equations[i].length == 2`
- 1 <= A~i~.length, B~i~.length <= 5
- `values.length == equations.length`
- `0.0 < values[i] <= 20.0`
- `1 <= queries.length <= 20`
- `queries[i].length == 2`
- 1 <= C~j~.length, D~j~.length <= 5
- A~i~, B~i~, C~j~, D~j~ consist of lower case English letters and digits.

### Related Topics

- Array
- Depth-First Search
- Breadth-First Search
- Union Find
- Graph
- Shortest Path
  
---

# 解題方向

### Solved using Union Find concept

`Union Find` 的進階題，在 `Find()` 與 `Union()` 的過程中要不斷維護 `edge` 的資料變化

詳細邏輯可以參考 LeetCodeCN 的[官方解法](https://leetcode.cn/problems/evaluate-division/solution/399-chu-fa-qiu-zhi-nan-du-zhong-deng-286-w45d/)

### Solved using Shortest Path Algorithm concept

施工中 ...

