[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/power-of-four/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 342. Power of Four

Given an integer `n`, return `true` if it is a power of four. Otherwise, return `false`.

An integer `n` is a power of four, if there exists an integer `x` such that n == $4^x$.

### Example 1:

```
Input: n = 16
Output: true
```

### Example 2:

```
Input: n = 5
Output: false
```

### Example 3:

```
Input: n = 1
Output: true
```

### Constraints:

- -$2^{31}$ <= n <= $2^{31}$ - 1

#### Follow up: 

Could you solve it without loops/recursion?

### Related Topics

- Math
- Bit Manipulation
- Recursion
  
---

# 解題方向

### Solved using Recursion concept

只要不能被 `4` 整除就提早中斷

### Solved using Bit Manipulation concept

第一步先確認 `n` 符合 `2` 的次方，可以用這個方式來確認： `n & (n-1) == 0`

- e.g. `100 & 011 = 0`
- e.g. `1000 & 0111 = 0`

接著只要確認 `n` 是不是只符合 `4` 的次方，轉為二進制來看的話可以檢查所有 `1` 能出現的位置，即：

- `01010101010101010101010101010101(2)`
- `55555555(16)`

詳細解釋請參考[這篇](https://leetcode.com/problems/power-of-four/discuss/80456/O(1)-one-line-solution-without-loops)