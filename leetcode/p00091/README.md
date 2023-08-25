[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/decode-ways/description/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 91. Decode Ways

自己看一下

### Related Topics

- String
- Dynamic Programming
  
---

# 解題方向

這題是比較麻煩的 DP 題，原因在於需要先判斷字串內的數字是不是合法的 (i.e. "1" - "26")

核心思想跟登高問題一樣，要到達 `i` 的位置，可能的走法會是 `i-1` 跟 `i-2` 的走法加在一起  

因此最優子結構的問題變成:  
 - 判斷 `s[i:i+1]` 這個位置只要合法，代表可以沿用 `dp[i-1]` 的結果
 - 判斷 `s[i-1:i+1]` 這個位置只要合法, 代表可以沿用 `dp[i-2]` 的結果

為了避免 `s[0]` 這個位置出現不合理的答案，所以我們可以強迫給定一個初始值在 `s[0]` 之前，即
 - `s = "#" + s` 且定義 `dp[0] = 1` (這邊給定初始值的方式有點 tricky)
 - 接著就可以判斷 `s[1]`(原本的`s[0]`) 這個位置到底合不合法了

後面就是 `dp` 的狀態轉移方程
 - `dp[i] = dp[i-1](只要合法) + dp[i-2](只要合法)`

剩下詳細的直接看程式碼吧
