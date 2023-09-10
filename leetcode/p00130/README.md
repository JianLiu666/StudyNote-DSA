[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/surrounded-regions/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 130. Surrounded Regions

自己看一下

### Related Topics

- Array
- Depth-First Search
- Breadth-First Search
- Union Find
- Matrix
  
---

# 解題方向

第一次遍歷 `board` 時從邊緣是 `'O'` 的節點開始用 DFS 或 BFS 的方式走訪，並且在走訪過程中將 `'O'` 標記成 `'V'`  
第二次遍歷 `board` 時將所有被標記成 `'V'` 的節點轉換成 `'O'`，其餘的節點轉換成 `'X'`
