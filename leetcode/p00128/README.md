[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/longest-consecutive-sequence/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 128. Longest Consecutive Sequence

Given an unsorted array of integers `nums`, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in `O(n)` time.

### Example 1:

```
Input: nums = [100,4,200,1,3,2]
Output: 4

Explanation:
 - The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
```

### Example 2:

```
Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9
```

### Constraints:

- 0 <= nums.length <= $10^5$
- $-10^9$ <= nums[i] <= $10^9$

### Related Topics

- Array
- Hash Table
- Union Find
  
---

# 解題方向

### Solved using Hash Table concept

對 `nums` 遍歷一次將資料存到 `Hash Table` 內，接著對 `Hash Table` 遍歷一次，過程中只要找到能夠代表**起點**的資料，就開始向上計數，直到碰到連續數列的末端為止