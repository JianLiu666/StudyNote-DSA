[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/move-zeroes/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# Move Zeroes

Given an integer array `nums`, move all `0`'s to the end of it while maintaining the relative order of the non-zero elements.

**Note** that you must do this in-place without making a copy of the array.

### Example 1:

```
Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]
```

### Example 2:

```
Input: nums = [0]
Output: [0]
```

### Constraints:

- 1 <= nums.length <= $10^4$
- -$2^{31}$ <= nums[i] <= $2^{31}$ - 1

#### Follow up:

Could you minimize the total number of operations done?

### Related Topics:

- Array
- Two Pointers

---

# 解題方向

Two Pointers: slow and fast pointer

 - Slow : 只處理非零數字
 - Fast : index in foreach loop