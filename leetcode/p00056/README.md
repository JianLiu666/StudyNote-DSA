[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/merge-intervals/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 56. Merge Intervals

Given an array of `intervals` where `intervals[i]` = [$\text{start}_i$, $\text{end}_i$], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

### Example 1:

```
Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]

Explanation:
 - Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
```

### Example 2:

```
Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]

Explanation:
 - Intervals [1,4] and [4,5] are considered overlapping.
```

### Constraints:

- 1 <= intervals.length <= $10^4$
- `intervals[i].length == 2`
- 0 <= starti <= endi <= $10^4$

### Related Topics

- Array
- Sorting
  
---

# 解題方向

題目給定的 `intervals` 本身是無序的，為了避免下面這種情況發生，因此第一件要做的事情就是將 `intervals` 先以 ascending order 排序過一遍
 - `intervals = [[1,2],[3,4],[5,6],[7,8],[1,10]]`

當 `intervals` 已經是一個有序陣列後事情就簡單多了，`intervals[i]` 的左邊已經確定，我們只需要檢查 `result[j][1]` 是否介於 `result[j][0]` 跟 `intervals[i-1][1]` 之間即可，是的話就將其合併，否則就將這個 `interval` 插入 `result` 內