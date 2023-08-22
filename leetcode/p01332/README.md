[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/remove-palindromic-subsequences/description/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 1332. Remove Palindromic Subsequences

自己看一下

### Related Topics

- Two Pointers
- String
  
---

# 解題方向

重點在於了解兩個名詞定義:
 - subsequence: 不需要不需要不需要連續
 - subarray: 需要連續

一旦分清楚之後這題秒殺
 - 空字串 = 0
 - 能夠回文 = 1
 - 不能夠回文 = 2次

```text
舉例:

  abbbababababaaaba
  __  ________  ___

  第一次可以拿掉所有 _ 位置的字母後剩下

  ..bb........aa...
  
  又是一個回文再拿掉就解決
```

屬於那種說破之後寫了一個寂寞的題目...