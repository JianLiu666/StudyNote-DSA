[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/determine-if-two-strings-are-close/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 1657. Determine if Two Strings Are Close

自己看一下

### Related Topics

- Hash Table
- String
- Sorting
  
---

# 解題方向

operation1 告訴我們位置順序不重要，只要所有字母的出現頻率是相同的，就可以一直換到都一樣為止  
operation2 告訴我們只要所有字母的出現頻率是一樣多的，就可以交換到一樣  

因此我們只需要把所有字母的出現頻率記錄下來，比較：
1. 是否兩個字串擁有的字母是一樣的 (只比較字母)
2. 這些字母的出現頻率是否一樣 (只比較頻率)