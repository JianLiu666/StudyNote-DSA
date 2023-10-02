[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/minimum-interval-to-include-each-query/description/)
![image](https://img.shields.io/badge/Difficulty-Hard-red)

---

# 1851. Minimum Interval to Include Each Query

自己看一下

### Related Topics

- Array
- Binary Search
- Line Sweep
- Sorting
- Heap (Priority Queue)
  
---

# 解題方向

先對 intervals 排序, 以 start 做 ascending order, 目的是讓最早的時刻放在最前面  

對 Queries 一樣做 ascending order, 這樣就可以讓 interval 對齊了, 陸續將元素在區間內的 intervals 放進 MinHeap 內  
MinHeap 的 comparision function 以 range size 為主, 越小的範圍越靠前  
只要把不符合的 intervals 從 MinHeap 去除掉後, 留下來的第一個 element in MinHeap 就是最小的區間  

剩下的直接看 code 吧