[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/add-strings/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 415. Add Strings

Given two non-negative integers, `num1` and `num2` represented as string, return the sum of `num1` and `num2` as a string.

You must solve the problem without using any built-in library for handling large integers (such as `BigInteger`). You must also not convert the inputs to integers directly.

### Example 1:

```
Input: num1 = "11", num2 = "123"
Output: "134"
```

### Example 2:

```
Input: num1 = "456", num2 = "77"
Output: "533"
```

### Example 3:

```
Input: num1 = "0", num2 = "0"
Output: "0"
```

### Constraints:

- 1 <= num1.length, num2.length <= $10^4$
- `num1` and `num2` consist of only digits.
- `num1` and `num2` don't have any leading zeros except for the zero itself.

### Related Topics

- Math
- String
- Simulation
  
---

# 解題方向

`String` 暖身題，因為題目給定的 `num1` 與 `num2` 都是數字表示，所以在相加時需要從字串的尾巴往前逐一相加，只要多留意進位問題即可