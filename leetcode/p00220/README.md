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

### Solved using Binary Search concept

一邊遍歷 `nums` 一邊維護一個長度為 `k+1` 的 sliding window，每當遇到新資料時都用 binary search 的方式查找 sliding window 內的資料：

- 如果 `window` 少於 `k+1`，以 ascending order 的方式將新資料加進 `window`
- 當 `window` 已經超過 `k+1` 筆時，要先將最舊的一筆資料移除(i.e., `nums[i-k-1]`) 後，才能加入新資料
- 每當有新資料加入時就跟左右鄰居相比，看看差異值是否落在 `t` 之內

### Solved using Bucket Sort concept

同上，差別在從 sliding window 改由 buckets 的概念來存放資料，每個 bucket 的資料範圍為 `t+1`，範例如下：

```
nums = [1,5,9,2,4,7,9,5], k = 3, t = 2

1: 1/(2+1) = bucket 0
5: 5/(2+1) = bucket 1
9: 9/(2+1) = bucket 3
2: 2/(2+1) = bucket 0
4: 4/(2+1) = bucket 1
7: 7/(2+1) = bucket 2
9: 9/(2+1) = bucket 3
5: 5/(2+1) = bucket 1

- bucket 0 has [1,2]
- bucket 1 has [5,4,5]
- bucket 2 has [7]
- bucket 3 has [9,9]
```

換句話說，只要落在同一個 `bucket` 內的資料差值一定小於 `k`，另外鄰近的 buckets 也**可能存在**差值小於 `k` 的資料，因此也必須檢查，總結一下流程：

- 如果當 `buckets` 的資料少於 `k+1` 筆時，就直接將新資料加進 `bucket`
- 當 `buckets` 的資料超過 `k+1` 筆時，先把最舊的一筆移除後再加入新資料
- 每當有新資料加入時：
  - 如果同一個 `bucket` 內已經存在資料，表示答案存在
  - 如果 `bucket-1` 或 `bucket+1` 內存在資料，且兩筆資料差值小於 `k` 也表示答案存在
