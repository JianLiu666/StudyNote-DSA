[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/word-break/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 139. Word Break

自己看一下

### Related Topics

- Array
- Hash Table
- String
- Dynamic Programming
- Trie
- Memorization
  
---

# 解題方向

稍微複雜一點的 Trie, 重點在 search 的方法需要加入遞迴檢查, 詳細直接看程式碼吧  
值得注意一提的是加入剪枝的策略, 確保在遇到極端測資時重複處理次數過多導致 TLE 的問題  