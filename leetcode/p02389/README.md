[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/longest-subsequence-with-limited-sum/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 2389. Longest Subsequence With Limited Sum

You are given an integer array `nums` of length `n`, and an integer array `queries` of length `m`.

Return an array `answer` of length `m` where `answer[i]` is the **maximum** size of a **subsequence** that you can take from `nums` such that the **sum** of its elements is less than or equal to `queries[i]`.

A **subsequence** is an array that can be derived from another array by deleting some or no elements without changing the order of the remaining elements.

### Example 1:

```
Input: nums = [4,5,2,1], queries = [3,10,21]
Output: [2,3,4]

Explanation: We answer the queries as follows:
 - The subsequence [2,1] has a sum less than or equal to 3. It can be proven that 2 is the maximum size of such a subsequence, so answer[0] = 2.
 - The subsequence [4,5,1] has a sum less than or equal to 10. It can be proven that 3 is the maximum size of such a subsequence, so answer[1] = 3.
 - The subsequence [4,5,2,1] has a sum less than or equal to 21. It can be proven that 4 is the maximum size of such a subsequence, so answer[2] = 4.
```

### Example 2:

```
Input: nums = [2,3,4,5], queries = [1]
Output: [0]

Explanation:
 - The empty subsequence is the only subsequence that has a sum less than or equal to 1, so answer[0] = 0.
```

### Constraints:

- `n == nums.length`
- `m == queries.length`
- `1 <= n, m <= 1000`
- 1 <= nums[i], queries[i] <= $10^6$

### Related Topics

- Array
- Binary Search
- Greedy
- Sorting
- Prefix Sum
  
---

# 解題方向

打 [Weekly Contest 308](https://leetcode.com/contest/weekly-contest-308/) 時太緊張被 **subsequence** 的解釋誤導，其實說穿了就是 `Sorting` 的暖身題

只要把 `nums` 按 **Ascending Order** 排序後與 `queries` 互相遍歷一次即可，邏輯是這樣：

- 題目要我們返回 longest subsequence with limited sum，所以我們就從小的先塞，如果小的都塞不進去那大的就更不用說了