[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/non-overlapping-intervals/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# No. Title

Given an array of intervals `intervals` where `intervals[i] = [starti, endi]`, return the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.

### Example 1:

```
Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
Output: 1

Explanation:
 - [1,3] can be removed and the rest of the intervals are non-overlapping.
```

### Example 2:

```
Input: intervals = [[1,2],[1,2],[1,2]]
Output: 2

Explanation:
 - You need to remove two [1,2] to make the rest of the intervals non-overlapping.
```

### Example 3:

```
Input: intervals = [[1,2],[2,3]]
Output: 0

Explanation:
 - You don't need to remove any of the intervals since they're already non-overlapping.
```

### Constraints:

- 1 <= intervals.length <= $10^5$
- `intervals[i].length == 2`
- -5 * $10^4$ <= starti < endi <= 5 * $10^4$

### Related Topics

- Array
- Dynamic Programming
- Greedy
- Sorting
  
---

# 解題方向

### Solved using Greedy concept

[這篇討論](https://leetcode.com/problems/non-overlapping-intervals/discuss/276056/Python-Greedy-Interval-Scheduling)解釋的很詳細，只要越早結束可以有越多的剩餘時間處理其他事物

所以只要把 `intervals` 重新以 `endTime` 由小到大排序後計數有多少個重複的 `intervals` 需要被剔除即可