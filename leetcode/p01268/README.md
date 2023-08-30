[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/search-suggestions-system/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 1268. Search Suggestions System

自己看一下

### Related Topics

- Array
- String
- Binary Search
- Trie
- Sorting
- Heap (Priority Queue)
  
---

# 解題方向

目前的做法是 trie + 每層記下所有經過的 words 有哪些, 但這樣可能會造成儲存空間的浪費, 可以用 binary search 優化, 在 suggest 的時候查找最小的前三個  
之後有需要再回頭優化這段程式碼  