[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/number-of-arithmetic-triplets/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 2367. Number of Arithmetic Triplets

You are given a **0-indexed**, **strictly increasing** integer array `nums` and a positive integer `diff`. A triplet `(i, j, k)` is an **arithmetic triplet** if the following conditions are met:

- `i < j < k`,
- `nums[j] - nums[i] == diff`, and
- `nums[k] - nums[j] == diff`.

Return the number of unique **arithmetic triplets**.

### Example 1:

```
Input: nums = [0,1,4,6,7,10], diff = 3
Output: 2

Explanation:
 - (1, 2, 4) is an arithmetic triplet because both 7 - 4 == 3 and 4 - 1 == 3.
 - (2, 4, 5) is an arithmetic triplet because both 10 - 7 == 3 and 7 - 4 == 3. 
```

### Example 2:

```
Input: nums = [4,5,6,7,8,9], diff = 2
Output: 2

Explanation:
 - (0, 2, 4) is an arithmetic triplet because both 8 - 6 == 2 and 6 - 4 == 2.
 - (1, 3, 5) is an arithmetic triplet because both 9 - 7 == 2 and 7 - 5 == 2.
```

### Constraints:

- `3 <= nums.length <= 200`
- `0 <= nums[i] <= 200`
- `1 <= diff <= 50`
- `nums` is **strictly** increasing.

### Related Topics

- Array
- Hash Table
- Two Pointers
- Enumeration
  
---

# 解題方向

### Solved using Hash Table

因為題目明確說明 `nums` 為 **strictly increasing**，因此我們可以先將 `nums` 用 `Hash Table` 記錄起來後，在對 `nums` 遍歷一次：

- 以 `(i, j, k)` 的 `k` 為主對 `Hash Table` 查找 `memo[k-diff] && memo[k+diff]` 同時存在，則代表符合條件的 `(i, j, k)` 存在