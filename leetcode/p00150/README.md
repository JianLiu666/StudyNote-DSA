[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/evaluate-reverse-polish-notation/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 150. Evaluate Reverse Polish Notation

Evaluate the value of an arithmetic expression in [Reverse Polish Notation](https://en.wikipedia.org/wiki/Reverse_Polish_notation).

Valid operators are `+`, `-`, `*`, and `/`. Each operand may be an integer or another expression.

**Note** that division between two integers should truncate toward zero.

It is guaranteed that the given RPN expression is always valid. That means the expression would always evaluate to a result, and there will not be any division by zero operation.

### Example 1:

```
Input: tokens = ["2","1","+","3","*"]
Output: 9

Explanation: ((2 + 1) * 3) = 9
```

### Example 2:

```
Input: tokens = ["4","13","5","/","+"]
Output: 6

Explanation: (4 + (13 / 5)) = 6
```

### Example 3:

```
Input: tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
Output: 22

Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
  = ((10 * (6 / (12 * -11))) + 17) + 5
  = ((10 * (6 / -132)) + 17) + 5
  = ((10 * 0) + 17) + 5
  = (0 + 17) + 5
  = 17 + 5
  = 22
```

### Constraints:

- 1 <= tokens.length <= $10^4$
- `tokens[i]` is either an operator: `"+"`, `"-"`, `"*"`, or `"/"`, or an integer in the range `[-200, 200]`.

### Related Topics

- Array
- Math
- Stack
  
---

# 解題方向

`Stack` 暖身題，遇到 `"+"`, `"-"`, `"*"`, `"/"` 時就從 `Stack` 取出兩個數字運算後放回去 `Stack` ，重複直到算式結束
