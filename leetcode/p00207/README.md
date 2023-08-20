[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/course-schedule/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 207. Course Schedule

There are a total of `numCourses` courses you have to take, labeled from `0` to `numCourses - 1`. You are given an array `prerequisites` where `prerequisites[i]` = [ $a_i$, $b_i$ ] indicates that you must take course $b_i$  first if you want to take course $a_i$.

- For example, the pair `[0, 1]`, indicates that to take course `0` you have to first take course `1`.

Return `true` if you can finish all courses. Otherwise, return `false`.

### Example 1:

```
Input: numCourses = 2, prerequisites = [[1,0]]
Output: true

Explanation:
 - There are a total of 2 courses to take. 
 - To take course 1 you should have finished course 0. 
 - So it is possible.
```

### Example 2:

```
Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
Output: false

Explanation:
 - There are a total of 2 courses to take. 
 - To take course 1 you should have finished course 0, and to take course 0 you should also have finished course 1.
 - So it is impossible.
```

### Constraints:

- `1 <= numCourses <= 2000`
- `0 <= prerequisites.length <= 5000`
- `prerequisites[i].length == 2`
- 0 <= $a_i$, $b_i$ < numCourses
- All the pairs prerequisites[i] are **unique**.

### Related Topics

- Depth-First Search
- Breadth-First Search
- Graph
- Topological Sort
  
---

# 解題方向

### Solved using Depth-First Search concept

這題的核心問題在於檢查 `Graph` 本身是否為合法的 DAG (i.e., 不存在循環)

首先先將 `Graph` 按照題目給定的 `prerequisites` 建構完畢後，用 DFS 搭配 `Hash Table` 的觀念切入來檢查是否出現循環

更新了 BFS 與 DFS 的做法