[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/task-scheduler/description/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 621. Task Scheduler

自己看一下

### Related Topics

- Array
- Hash Table
- Greedy
- Sorting
- Heap (Priority Queue)
- Counting
  
---

# 解題方向

用 hash table 把相同編號的 task 能夠執行的間隔時間區分出來後當作 min-heap 的排序條件  
接著不斷從 min-heap 取出能夠被執行的 task, 直到整個 min-heap 內沒有任何 task 就是完成時間  