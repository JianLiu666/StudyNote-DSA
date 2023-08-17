[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/longest-valid-parentheses/)
![image](https://img.shields.io/badge/Difficulty-Hard-red)

---

# Longest Valid Parentheses

自己看一下

### Related Topics

- String
- Dynamic Programming
- Stack
  
---

# 解題方向

括號問題就是 stack 無誤！

必須考慮的像是 `((())` 跟 `)(())` 這種無法完全匹配的
我們可以先把能夠匹配的消去之後看看剩下的會是什麼

```
例如：
(((())) 當我現在走到第 i 步時
      i 

中間的括號應該早就被匹配完了如下所示

(((()))
((****i
```

因此我就可以用 `stack` 來維護還沒有被匹配成功的位置, 從那個位置往後移一格就是最後一次能夠正確匹配的位置,詳細看程式碼吧  
或者是直接去看官神的[影片](https://www.youtube.com/watch?v=677VaZhd4dg), 官神真的神
