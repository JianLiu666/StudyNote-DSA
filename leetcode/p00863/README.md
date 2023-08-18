[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/description/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 863. All Nodes Distance K in Binary Tree

自己看一下

### Related Topics

- Tree
- Depth-First Search
- Breadth-First Search
- Binary Tree
  
---

# 解題方向

### Memorization

能想出這個解法的人真的是天才: [LeetCode Solution](https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/solutions/143798/1ms-beat-100-simple-java-dfs-with-without-hashmap-including-explanation/)

詳細的解法可以直接看程式碼，這裡只記錄一下我在看這個解法時的一些片段

step.1: 紀錄從 `root` 出發沿路到 `target` 這之間每個 `node` 與 `target` 之間的距離
 - 目的是可以讓 dfs 的過程中知道現在這個節點與 target 之間的距離會有多遠

但這件事情只是用在與 target 相反邊的 subtree 才會越來越遠, 若是 target 同方向的 subtree 只會越走越近
所以作者才會在每一次的 dfs 都檢查一次 map 不斷修正這個 node 與 target 之間的距離

i.e. 只要存在 map 內的 nodes 表示離 target 越來越近, 反之不在 map 內的 nodes 就是與 target 越來越遠

### Graph

官神開示，先把 tree 轉成 graph 後就可以用 BFS 找到嚕
 - [傳送門在此](https://www.youtube.com/watch?v=53yZy6BWVzc): 雖然他還是用 DFS 解的就是了...