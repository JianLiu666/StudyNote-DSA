[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/two-sum/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

# 1. Two Sum

Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `target`.

You may assume that each input would have **exactly one solution**, and you may not use the same element twice.

You can return the answer in any order.

### Example1:

```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
```

### Example 2:

```
Input: nums = [3,2,4], target = 6
Output: [1,2]
```

### Example 3:

```
Input: nums = [3,3], target = 6
Output: [0,1]
```

### Constraints:

- 2 <= nums.length <= $10^4$
- -$10^9$ <= nums[i] <= $10^9$
- -$10^9$ <= target <= $10^9$
- **Only one valid answer exists.**

### Related Topics

- Array
- Hash Table

---

# 解題方向

透過 Hash Map 紀錄出現過的數字，每次迭代時比對 Map 內是否有加總後符合答案的數字