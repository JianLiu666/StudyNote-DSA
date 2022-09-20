[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/course-schedule-ii/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 210. Course Schedule II

There are a total of `numCourses` courses you have to take, labeled from `0` to `numCourses - 1`. You are given an array `prerequisites` where `prerequisites[i]` = [ $a_i$, $b_i$ ] indicates that you **must** take course $b_i$  first if you want to take course $a_i$.

- For example, the pair `[0, 1]`, indicates that to take course `0` you have to first take course `1`.

Return the ordering of courses you should take to finish all courses. If there are many valid answers, return **any** of them. If it is impossible to finish all courses, return **an empty array**.

### Example 1:

```
Input: numCourses = 2, prerequisites = [[1,0]]
Output: [0,1]

Explanation:
 - There are a total of 2 courses to take. To take course 1 you should have finished course 0. So the correct course order is [0,1].
```

### Example 2:

```
Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
Output: [0,2,1,3]

Explanation:
 - There are a total of 4 courses to take. To take course 3 you should have finished both courses 1 and 2. Both courses 1 and 2 should be taken after you finished course 0.
 - So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3].
```

### Example 3:

```
Input: numCourses = 1, prerequisites = []
Output: [0]
```

### Constraints:

- `1 <= numCourses <= 2000`
- `0 <= prerequisites.length <= numCourses * (numCourses - 1)`
- `prerequisites[i].length == 2`
- 0 <= $a_i$, $b_i$ < numCourses
- $a_i$ != $b_i$
- All the pairs [ $a_i$, $b_i$ ] are **distinct**.

### Related Topics

- Depth-First Search
- Breadth-First Search
- Graph
- Topological Sort
  
---

# 解題方向

### Solved using Depth-First Search concept

這題的核心問題在於檢查 `Graph` 本身是否為合法的 DAG (i.e., 不存在循環)

首先先將 `Graph` 按照題目給定的 `prerequisites` 建構完畢後，用 DFS 搭配 `Hash Table` 的觀念切入來檢查是否出現循環, 一旦檢查到循環路徑提早結束

反之就將走過的路徑紀錄下來即可