[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/longest-repeating-character-replacement/description/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 424. Longest Repeating Character Replacement

You are given `a` string s and an integer `k`. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most `k` times.

Return the length of the longest substring containing the same letter you can get after performing the above operations.

### Example 1:

```
Input: s = "ABAB", k = 2
Output: 4
Explanation: 
 - Replace the two 'A's with two 'B's or vice versa.
```

### Example 2:

```
Input: s = "AABABBA", k = 1
Output: 4
Explanation: 
 - Replace the one 'A' in the middle with 'B' and form "AABBBBA".
 - The substring "BBBB" has the longest repeating letters, which is 4.
 - There may exists other ways to achive this answer too.
```

### Constraints:

 - 1 <= s.length <= $10^5$
 - `s` consists of only uppercase English letters.
 - `0 <= k <= s.length`

### Related Topics

- Hash Table
- String
- Sliding Window
  
---

# 解題方向

1. 控制好 sliding window 的範圍
2. 用 hash table 紀錄 sliding window 內的字母出現頻率, 每次迭代都檢查目前 window 內最多出現次數的字母數量是多少, 加上 replacement 後能否全部替換掉來找到最大值, 詳細看 code
