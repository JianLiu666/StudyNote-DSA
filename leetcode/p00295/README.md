[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/find-median-from-data-stream/description/)
![image](https://img.shields.io/badge/Difficulty-Hard-red)

---

# 295. Find Median from Data Stream

自己看一下

### Related Topics

- Two Pointers
- Design
- Sorting
- Heap (Priority Queue)
- Data Stream
  
---

# 解題方向

### Solved using Heap (Priority Queue)

用兩個不同的 priority queue 解決這個問題, `min-heap` 跟 `max-heap`  
實作的想法是參考 LeetCode Solutions 這位大神的[解答](https://leetcode.com/problems/find-median-from-data-stream/description/)  

這個做法真的很巧妙，透過兩個不同排序方式的 heap，我們可以把需要的 median 固定在各自的第一個元素上  
詳細流程請直接看程式碼  