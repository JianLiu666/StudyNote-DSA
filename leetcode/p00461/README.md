[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/hamming-distance/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 461. Hamming Distance

The [Hamming distance](https://en.wikipedia.org/wiki/Hamming_distance) between two integers is the number of positions at which the corresponding bits are different.

Given two integers `x` and `y`, return the **Hamming distance** between them.

### Example 1:

```
Input: x = 1, y = 4
Output: 2

Explanation:
 1   (0 0 0 1)
 4   (0 1 0 0)
        ↑   ↑
 The above arrows point to positions where the corresponding bits are different.
```

### Example 2:

```
Input: x = 3, y = 1
Output: 1
```

### Constraints:

- 0 <= x, y <= $2^{31}$ - 1

### Related Topics

- Bit Manipulation
  
---

# 解題方向

將兩數經過 XOR 運算後可以得到所有位置的差異值，再把所有差異值加總起來即可