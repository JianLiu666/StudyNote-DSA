[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/minimum-size-subarray-sum/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 209. Minimum Size Subarray Sum

iven an array of positive integers `nums` and a positive integer `target`, return the minimal length of a **contiguous subarray** [ $\text{nums}_l$, $\text{nums}_{l+1}$, ..., $\text{nums}_{r-1}$, $\text{nums}_r$ ] of which the sum is greater than or equal to `target`. If there is no such subarray, return `0` instead.

### Example 1:

```
Input: target = 7, nums = [2,3,1,2,4,3]
Output: 2

Explanation:
 - The subarray [4,3] has the minimal length under the problem constraint.
```

### Example 2:

```
Input: target = 4, nums = [1,4,4]
Output: 1
```

### Example 3:

```
Input: target = 11, nums = [1,1,1,1,1,1,1,1]
Output: 0
```

### Constraints:

- 1 <= target <= $10^9$
- 1 <= nums.length <= $10^5$
- 1 <= nums[i] <= $10^4$

#### Follow up:

If you have figured out the `O(n)` solution, try coding another solution of which the time complexity is `O(n log(n))`.

### Related Topics

- Array
- Binary Search
- Sliding Window
- Prefix Sum
  
---

# 解題方向

Two Pointers (head and tail pointers) 解決，我們將 `head pointer` 與 `tail pointer` 的距離控制在總和剛好超過 `target` 的長度，只要超過 `target` 就調整距離。

不斷移動直到 `nums` 盡頭或是找到最小的 Subarray Sum (i.e., 1)