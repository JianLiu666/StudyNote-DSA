[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 219. Contains Duplicate II

Given an integer array `nums` and an integer `k`, return `true` if there are two **distinct indices** `i` and `j` in the array such that `nums[i] == nums[j]` and `abs(i - j) <= k`.

### Example 1:

```
Input: nums = [1,2,3,1], k = 3
Output: true
```

### Example 2:

```
Input: nums = [1,0,1,1], k = 1
Output: true
```

### Example 3:

```
Input: nums = [1,2,3,1,2,3], k = 2
Output: false
```

### Constraints:

- 1 <= nums.length <= $10^5$
- -$10^9$ <= nums[i] <= $10^9$
- 0 <= k <= $10^5$

### Related Topics

- Array
- Hash Table
- Sliding Window
  
---

# 解題方向

`Hash Table` 暖身題，把出現過的數字紀錄在 `Hash Table` 拿來後續比對使用即可。