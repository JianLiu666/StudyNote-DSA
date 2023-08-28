[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 1466. Reorder Routes to Make All Paths Lead to the City Zero

自己看一下

### Related Topics

- Depth-First Search
- Breadth-First Search
- Graph
  
---

# 解題方向

### Brute-force

一開始的想法很單純，就是先把 tree 轉成 DAG 後，就可以跟 connections 比較差異，把差異記下來  
但實際寫法是先把 connections 建成一個 graph 後，在用 graph 轉成 DAG...


### Using Adjacency List

跟我一開始想到的邏輯一樣，但只要直接轉成 adjacency list 處理就好，直接看程式碼吧
