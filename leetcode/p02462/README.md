[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/total-cost-to-hire-k-workers/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 2462. Total Cost to Hire K Workers

自己看一下

### Related Topics

- Array
- Two Pointers
- Heap (Priority Queue)
- Simulation
  
---

# 解題方向

這題麻煩的地方在於每此挑選 worker 時會因為 `candidates` 而受限，所以我們必須維護兩個 pointers `left` & `right` 來幫助我們找到目前可參考的 workers 有哪些  
剩下的就是基本的 Min-Heap 了  