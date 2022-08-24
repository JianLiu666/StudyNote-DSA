[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/contains-duplicate-iii/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 220. Contains Duplicate III

Given an integer array `nums` and two integers `k` and `t`, return `true` if there are **two distinct indices** `i` and `j` in the array such that `abs(nums[i] - nums[j]) <= t` and `abs(i - j) <= k`.

### Example 1:

```
Input: nums = [1,2,3,1], k = 3, t = 0
Output: true
```

### Example 2:

```
Input: nums = [1,0,1,1], k = 1, t = 2
Output: true
```

### Example 3:

```
Input: nums = [1,5,9,1,5,9], k = 2, t = 3
Output: false
```

### Constraints:

- 1 <= nums.length <= 2 * $10^4$
- -$2^{31}$ <= nums[i] <= $2^{31}$ - 1
- 0 <= k <= $10^4$
- 0 <= t <= $2^{31}$ - 1

### Related Topics

- Array
- Sliding Window
- Sorting
- Bucket Sort
- Ordered Set
  
---

# 解題方向

### Solved using Sliding Window concept

一邊遍歷 `nums` 一邊維護一個長度為 `k+1` 的 sliding window，每當遇到新資料時：

- 如果 sliding window 少於 `k+1`，以 ascending order 的方式將新資料加進 window
- 當 sliding window 已經超過 `k+1` 筆時，要先將最舊的一筆資料移除(i.e., `nums[i-k-1]`) 後，才能加入新資料
- 每當有新資料加入時就跟左右鄰居相比，看看差異值是否落在 `t` 之內

### Solved using Bucket Sort concept

