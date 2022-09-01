[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/multiply-strings/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 43. Multiply Strings

Given two non-negative integers `num1` and `num2` represented as strings, return the product of `num1` and `num2`, also represented as a string.

**Note:** You must not use any built-in BigInteger library or convert the inputs to integer directly.

### Example 1:

```
Input: num1 = "2", num2 = "3"
Output: "6"
```

### Example 2:

```
Input: num1 = "123", num2 = "456"
Output: "56088"
```

### Constraints:

- `1 <= num1.length, num2.length <= 200`
- `num1` and `num2` consist of digits only.
- Both `num1` and `num2` do not contain any leading zero, except the number `0` itself.


### Related Topics

- Math
- String
- Simulation
  
---

# 解題方向

四則運算的乘法運算，由於輸入的數字 `num1` 跟 `num2` 是以 `string` 表示，所以在處理乘法時我們需要反過來從尾巴開始做才會是最低位元

再乘法過程中會一直去改變中途的數字，所以我選擇先用 `[]byte` 來暫存數字，`[]byte` 的最大長度會等於 `len(num1) + len(num2)` (e.g. `999 x 999 = 998001`)

接著只要依照乘法規則不斷的遍歷 `num1` 與 `num2`，注意進位問題後即可，剩下就直接看程式碼吧