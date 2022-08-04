[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/k-th-symbol-in-grammar/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 779. K-th Symbol in Grammar

We build a table of `n` rows (**1-indexed**). We start by writing `0` in the $1^{st}$ row. Now in every subsequent row, we look at the previous row and replace each occurrence of `0` with `01`, and each occurrence of `1` with `10`.

- For example, for `n = 3`, the $1^{st}$ row is `0`, the $2^{nd}$ row is `01`, and the $3^{rd}$ row is `0110`.

Given two integer `n` and `k`, return the $k^{th}$ (**1-indexed**) symbol in the $n^{th}$ row of a table of `n` rows.

### Example 1:

```
Input: n = 1, k = 1
Output: 0

Explanation:
 - row 1: 0
```

### Example 2:

```
Input: n = 2, k = 1
Output: 0

Explanation: 
 - row 1: 0
 - row 2: 01
```

### Example 3:

```
Input: n = 2, k = 2
Output: 1

Explanation: 
 - row 1: 0
 - row 2: 01
```

### Constraints:

- `1 <= n <= 30`
- 1 <= k <= $2^n$ - 1


### Related Topics

- Math
- Bit Manipulation
- Recursion
  
---

# 解題方向

[神之想法](https://leetcode.com/problems/k-th-symbol-in-grammar/discuss/1151025/Java-Easy-to-understand-solution)，用 `Tree` 的角度來看這道題目，如下：

```
               0
       ┌---------------┐
       0               1
   ┌-------┐       ┌-------┐
   0       1       1       0
 ┌---┐   ┌---┐   ┌---┐   ┌---┐
 0   1   1   0   1   0   0   1
┌-┐ ┌-┐ ┌-┐ ┌-┐ ┌-┐ ┌-┐ ┌-┐ ┌-┐
0 1 1 0 1 0 0 1 1 0 0 1 0 1 1 0
```

