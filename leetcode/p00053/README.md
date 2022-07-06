[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/maximum-subarray/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 53. Maximum Subarray

Given an integer array `nums`, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

A **subarray** is a **contiguous** part of an array.

### Example 1:

```
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6

Explanation: 
 - [4,-1,2,1] has the largest sum = 6.
```
### Example 2:

```
Input: nums = [1]
Output: 1
```

### Example 3:

```
Input: nums = [5,4,-1,7,8]
Output: 23
```

### Constraints:

- 1 <= nums.length <= $10^5$
- -$10^4$ <= nums[i] <= $10^4$

#### Follow up: 

If you have figured out the `O(n)` solution, try coding another solution using the **divide and conquer** approach, which is more subtle.

### Related Topics

- Array
- Divide and Conquer
- Dynamic Programming
  
---

# 解題方向

我們的目標是找到 `nums` 中加總數字最大的 `sub array`，我的理解是只要確保我們持續累加的數字不會讓總和為負，就可以持續連續下去

- i.e., 只要本次加總後總和為負數，代表我需要靠下一個數字幫我把這次產生的洞給填滿，那不如歸零重新計算加總