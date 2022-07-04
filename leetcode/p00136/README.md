[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/single-number/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 136. Single Number

Given a **non-empty** array of integers `nums`, every element appears twice except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.

### Example 1:

```
Input: nums = [2,2,1]
Output: 1
```

### Example 2:

```
Input: nums = [4,1,2,1,2]
Output: 4
```

### Example 3:

```
Input: nums = [1]
Output: 1
```

### Constraints:

- 1 <= nums.length <= 3 * $10^4$
- -3 * $10^4$ <= nums[i] <= 3 * $10^4$
- Each element in the array appears twice except for one element which appears only once.

### Related Topics:

- Array
- Bit Manipulation

---

# 解題方向

### Solved using data structure

- 遍歷一遍陣列, 將所有數字計數到 HashMap。
- 遍歷一遍 HashMap, 找到只出現過一次的數字。

但這樣做不符合題目要求，需要額外的空間複雜度 `O(n)` 來紀錄出現過的數字。

### Solved using bit manipulation

互斥或 (Exclusive or, XOR)
- 具備下列性質
   - 交換律 (Commutativity) : A $\oplus{}$ B = B $\oplus{}$ A
   - 結合律 (Associativity) : A $\oplus{}$ ( B $\oplus{}$ C ) = ( A $\oplus{}$ B ) $\oplus{}$ C

真值表 (Truth table)
  
| A     | B     | A $\oplus{}$ B |
| :---: | :---: | :---: |
| False | False | False |
| False | `True`  | `True`  |
| `True`  | False | `True`  |
| `True`  | `True`  | False |

換句話說，按照 XOR 的特性，只要遍歷一遍陣列就可以找到唯一一個只出現過一次的數字
- 根據真值表可得知，當兩數相同時則為零
- 根據交換律與結合律的特性可得知，即使不是連續出現也沒關係