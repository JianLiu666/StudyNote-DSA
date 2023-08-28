[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/longest-zigzag-path-in-a-binary-tree/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 1372. Longest ZigZag Path in a Binary Tree

自己看一下

### Related Topics

- Dynamic Programming
- Tree
- Depth-First Search
- Binary Tree
  
---

# 解題方向

`DFS` + `Memorization` 解決  
這題如果是單純的用 DFS 走訪，並且每個 node 都作為 root 找出 zigzag 的話馬上就會遇到 TLE 的問題  
所以我們的目標是讓走訪只進行一次，可以用後序遍歷從 leaf node 開始找起  
只要走訪過的 node 都加到 hash table 內，這樣就不需要每次都要重複走訪了  