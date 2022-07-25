[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 34. Find First and Last Position of Element in Sorted Array

Given an array of integers `nums` sorted in non-decreasing order, find the starting and ending position of a given `target` value.

If `target` is not found in the array, return `[-1, -1]`.

You must write an algorithm with `O(log n)` runtime complexity.

### Example 1:

```
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
```

### Example 2:

```
Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
```

### Example 3:

```
Input: nums = [], target = 0
Output: [-1,-1]
```

### Constraints:

- 0 <= nums.length <= $10^5$
- -$10^9$ <= nums[i] <= $10^9$
- `nums` is a non-decreasing array.
- -$10^9$ <= target <= $10^9$

### Related Topics

- Array
- Binary Search
  
---

# 解題方向

Binary Search 的變化題型，`target` 不是唯一值(unique) 時的問題，核心思想不變，只要把已確定的資料排除後不斷重複，直到找到左右邊界為止

- Left edge condition : `cursor == 0 or nums[cursor-1] != target`
- Right edge condition : `cursor == len(nums)-1 or nums[cursor+1] != target`

