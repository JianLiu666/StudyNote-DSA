[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/minimum-height-trees/description/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 310. Minimum Height Trees

自己看一下

### Related Topics

- Depth-First Search
- Breadth-First Search
- Graph
- Topological Sort
  
---

# 解題方向

用 Topological Sort 逐步把 indegree == 1 的節點剝掉 (因為 tree 是 undirected graph)  
直到 indegree == 1 的節點數量剩下 2 個以下時就是答案了  

至於為什麼是剩下兩個, 直接看下圖  

```
當剩下三個以上時一定還能再剝掉至少一次

1 - 2 - 3 剝掉後剩下 2

1 - 3 - 4 剝掉後剩下 3
2 /

1 - 3 - 4 - 5 剝掉後剩下 3 - 4, 此時兩邊做為 root node 都符合 MHT
2 / 
```