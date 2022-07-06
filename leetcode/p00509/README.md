[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/fibonacci-number/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 509. Fibonacci Number

The **Fibonacci numbers**, commonly denoted `F(n)` form a sequence, called the **Fibonacci sequence**, such that each number is the sum of the two preceding ones, starting from `0` and `1`. That is,

```
F(0) = 0, F(1) = 1
F(n) = F(n - 1) + F(n - 2), for n > 1.
```

Given `n`, calculate `F(n)`.

### Example 1:

```
Input: n = 2
Output: 1

Explanation:
 - F(2) = F(1) + F(0) = 1 + 0 = 1.
```

### Example 2:

```
Input: n = 3
Output: 2

Explanation:
 - F(3) = F(2) + F(1) = 1 + 1 = 2.
```

### Example 3:

```
Input: n = 4
Output: 3

Explanation:
 - F(4) = F(3) + F(2) = 2 + 1 = 3.
```

### Constraints:

- `0 <= n <= 30`

### Related Topics

- Math
- Dynamic Programming
- Recursion
- Memorization

---

# 解題方向

本來覺得這題就只是一個一般的暖身題，意外發現一個怪怪的知識 ...

### Solved using basic recursion concept

```python
# Time Complexity: O(2^n)
# Space Complexity: O(n), function call stacks
def fib(n: int) -> int:
    if n < 2:
        return n
    
    return fib(n-1) + fib(n-2)
```

### Solved using dynamic programming concept

Bottom-up DP

```python
# Time Complexity: O(n)
# Space Complexity: O(1)
def fib(n: int) -> int:
    if n < 2:
        return n

    fn, fn_1, fn_2: 0, 1, 0
    for i in range(1, n):
        fn = fn_1 + fn_2
        fn_2 = fn_1
        fn_1 = fn

    return fn
```

### Solved using Binet's Nth-term Formula

跟鬼一樣的公式解來了 -> [here](https://r-knott.surrey.ac.uk/Fibonacci/fibFormula.html)

`F(n)` = round( $Phi^n$ / $\sqrt{5}$ ), where $Phi$ = ( $\sqrt{5}$ + 1 ) / 2