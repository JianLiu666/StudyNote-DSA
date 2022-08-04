[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/find-k-th-smallest-pair-distance/)
![image](https://img.shields.io/badge/Difficulty-Hard-red)

---

# 719. Find K-th Smallest Pair Distance

he **distance of a pair** of integers `a` and `b` is defined as the absolute difference between `a` and `b`.

Given an integer array `nums` and an integer `k`, return the $k^{th}$ smallest **distance among all the pairs** `nums[i]` and `nums[j]` where `0 <= i < j < nums.length`.

### Example 1:

```
Input: nums = [1,3,1], k = 1
Output: 0

Explanation: Here are all the pairs:
 - (1,3) -> 2
 - (1,1) -> 0
 - (3,1) -> 2
 - Then the 1st smallest distance pair is (1,1), and its distance is 0.
```

### Example 2:

```
Input: nums = [1,1,1], k = 2
Output: 0
```

### Example 3:

```
Input: nums = [1,6,1], k = 3
Output: 5
```

### Constraints:

- `n == nums.length`
- 2 <= n <= $10^4$
- 0 <= nums[i] <= $10^6$
- `1 <= k <= n * (n - 1) / 2`

### Related Topics

- Array
- Two Pointers
- Binary Search
- Sorting
  
---

# 解題方向

這位作者寫的這篇 [Binary Search 核心原理](https://leetcode.com/problems/find-k-th-smallest-pair-distance/discuss/769705/Python-Clear-explanation-Powerful-Ultimate-Binary-Search-Template.-Solved-many-problems.) 根本是教科書等級，強烈建議一定要看過一遍！

這題的核心概念就是 Binary Search + Two Pointers 的 sliding window 方法

為了加速 `pair distance` 的比對過程，第一步必須先對 `nums` 做排序，才能將複雜度從 O($n^2$) 降到 O(n)

接著只要再決定好 Binary Search 的 `[low, high]` 範圍，Done