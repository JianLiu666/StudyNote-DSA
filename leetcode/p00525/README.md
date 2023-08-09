[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/contiguous-array/description/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 525. Contiguous Array

Given a binary array `nums`, return the maximum length of a contiguous subarray with an equal number of `0` and `1`.

### Example 1:

```
Input: nums = [0,1]
Output: 2
Explanation: 
 - [0, 1] is the longest contiguous subarray with an equal number of 0 and 1.
```

### Example 2:

Input: nums = [0,1,0]
Output: 2
Explanation: 
 - [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.

### Constraints:

- 1 <= nums.length <= $10^5$
- `nums[i]` is either `0` or `1`.

### Related Topics

- Array
- Hash Table
- Prefix Sum
  
---

# 解題方向

Prefix Sum 的變化題，詳細可以直接看程式碼
