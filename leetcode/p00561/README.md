[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/array-partition/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 561. Array Partition

Given an integer array `nums` of `2n` integers, group these integers into `n` pairs ($a_1$, $b_1$), ($a_2$, $b_2$), ..., ($a_n$, $b_n$) such that the sum of min($a_i$, $b_i$) for all `i` is maximized. Return the maximized sum.

### Example 1:

```
Input: nums = [1,4,3,2]
Output: 4

Explanation:
 - All possible pairings (ignoring the ordering of elements) are:
 - 1. (1, 4), (2, 3) -> min(1, 4) + min(2, 3) = 1 + 2 = 3
 - 2. (1, 3), (2, 4) -> min(1, 3) + min(2, 4) = 1 + 2 = 3
 - 3. (1, 2), (3, 4) -> min(1, 2) + min(3, 4) = 1 + 3 = 4
 - So the maximum possible sum is 4.
```

### Example 2:

```
Input: nums = [6,2,6,5,1,2]
Output: 9

Explanation:
 - The optimal pairing is (2, 1), (2, 5), (6, 6). min(2, 1) + min(2, 5) + min(6, 6) = 1 + 2 + 6 = 9.
```

### Constraints:

- 1 <= n <= $10^4$
- `nums.length == 2 * n`
- -$10^4$ <= nums[i] <= $10^4$

### Related Topics

- Array
- Greedy
- Sorting
- Couting Sort
  
---

# 解題方向

基本的 `Array` 暖身題，把 `Array` 由小到大排序後間隔累計即可。