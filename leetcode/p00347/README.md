[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/top-k-frequent-elements/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 347. Top K Frequent Elements

Given an integer array `nums` and an integer `k`, return the `k` most frequent elements. You may return the answer in **any order**.

### Example 1:

```
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
```

### Example 2:

```
Input: nums = [1], k = 1
Output: [1]
```

### Constraints:

- 1 <= nums.length <= $10^5$
- `k` is in the range `[1, the number of unique elements in the array]`.
- It is **guaranteed** that the answer is **unique**.

#### Follow up

Your algorithm's time complexity must be better than `O(n log n)`, where n is the array's size.

### Related Topics

- Array
- Hash Table
- Divide and Conquer
- Sorting
- Heap (Priority Queue)
- Bucket Sort
- Counting
- Quickselect
  
---

# 解題方向

### Solved using bucket sort concept

1. 將 `nums` 中所有數字的出現頻率用 `Hash Table` 記錄
2. 用一個 `array` 當作 `buckets`，按照出現頻率將數字放進相對的 `bucket`
3. 從 `buckets` 的尾巴往前依序拿取 `bucket` 內的數字，直到取回的數字滿足題目指定的 `k`
