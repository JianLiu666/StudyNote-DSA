[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/happy-number/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 202. Happy Number

Write an algorithm to determine if a number `n` is happy.

A **happy number** is a number defined by the following process:

- Starting with any positive integer, replace the number by the sum of the squares of its digits.
- Repeat the process until the number equals 1 (where it will stay), or it **loops endlessly in a cycle** which does not include 1.
- Those numbers for which this process **ends in 1** are happy.

Return `true` if n is a happy number, and `false` if not.

### Example 1:

```
Input: n = 19
Output: true
```

Explanation:
 - $1^2$ + $9^2$ = 82
 - $8^2$ + $2^2$ = 68
 - $6^2$ + $8^2$ = 100
 - $1^2$ + $0^2$ + $0^2$ = 1

### Example 2:

```
Input: n = 2
Output: false
```

### Constraints:

- 1 <= n <= $2^{31}$ - 1

### Related Topics

- Hash Table
- Math
- Two Pointers
  
---

# 解題方向

這題意外的有點有趣，根據題目描述，我們只有三種情境需要思考
- 收斂到 1
- 陷入循環
- ~~無限遞增~~，不可能的原因請參考本題的 [solution](https://leetcode.com/problems/happy-number/solution/)


```
case 1: HAPPY NUMBER!

      +-----+    +-----+    +-----+    +-----+    +-----+    +-----+
  n = |   7 | -> |  49 | -> |  97 | -> | 130 | -> |  10 | -> |   1 |
      +-----+    +-----+    +-----+    +-----+    +-----+    +-----+

case 2: 陷入循環

      +-----+    +-----+    +-----+    +-----+    +-----+    +-----+    +-----+
  n = | 116 | -> |  38 | -> |  73 | -> |  58 | -> |  89 | -> | 145 | -> |  42 |
      +-----+    +-----+    +-----+    +-----+    +-----+    +-----+    +-----+
                                          ^                                |
                                          |                                v
                                       +-----+    +-----+    +-----+    +-----+
                                       |  37 | <- |  16 | <- |   4 | <- |  20 |
                                       +-----+    +-----+    +-----+    +-----+

```

### Solved using hash table concept

透過 `Hash Set` 將出現過的數字記錄下來，看是先收斂到 1 還是遇到重複的數字

- Time Complexity: `O(logn)`
- Space Complexity: `O(logn)`

### Solved using two pointer concept

slow and fast pointers，不需再額外紀錄出現過的數字，只要看 `fast pointer` 先收斂到 1 或是追到 `slow pointer` 即可

- Time Complexity: `O(logn)`
- Space Complexity: `O(1)`

### Solved using mathematically concept

從上圖我們可以發現循環的規律是 4 -> 16 -> 37 -> 58 -> 89 -> 145 -> 42 -> 20 -> 4 -> ...

所以一但數字進到這個環狀線上的話就可以結束了，可以用一個 `Hash Set` 將循環數字都先定義好，或是乾脆只挑其中一個數字當作終止條件也可以

- Time Complexity: `O(logn)`
- Space Complexity: `O(1)`