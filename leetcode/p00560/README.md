[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/subarray-sum-equals-k/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 560. Subarray Sum Equals K

Given an array of integers `nums` and an integer `k`, return the total number of subarrays whose sum equals to `k`.

A subarray is a contiguous **non-empty** sequence of elements within an array.

### Example 1:

```
Input: nums = [1,1,1], k = 2
Output: 2
```

### Example 2:

```
Input: nums = [1,2,3], k = 3
Output: 2
```

### Constraints:

- 1 <= nums.length <= 2 * $10^4$
- `-1000 <= nums[i] <= 1000`
- -$10^7$ <= k <= $10^7$

### Related Topics

- Array
- Hash Table
- Prefix Sum
  
---

# 解題方向

### Solved using Prefix Sum concept

**LeetCodeCN** 上有相當精闢的解釋，可以參考[這篇](https://leetcode.cn/problems/subarray-sum-equals-k/solution/dai-ni-da-tong-qian-zhui-he-cong-zui-ben-fang-fa-y/)，或是[這篇](https://leetcode.cn/problems/subarray-sum-equals-k/solution/bao-li-jie-fa-qian-zhui-he-qian-zhui-he-you-hua-ja/)

舉個例子記錄一下，讓自己以後方便回憶:

```
nums = [1,1,1,1,1,3,-3,1,4,0], k = 5

 - i = 0, nums[0] =  1, presum =  1, count = 0, memo = {0:1, 1:1}
 - i = 1, nums[1] =  1, presum =  2, count = 0, memo = {0:1, 1:1, 2:1}
 - i = 2, nums[2] =  1, presum =  3, count = 0, memo = {0:1, 1:1, 2:1, 3:1}
 - i = 3, nums[3] =  1, presum =  4, count = 0, memo = {0:1, 1:1, 2:1, 3:1, 4:1}
 - i = 4, nums[4] =  1, presum =  5, count = 1, memo = {0:1, 1:1, 2:1, 3:1, 4:1, 5:1}
 - i = 5, nums[5] =  3, presum =  8, count = 2, memo = {0:1, 1:1, 2:1, 3:1, 4:1, 5:1, 8:1}
 - i = 6, nums[6] = -3, presum =  5, count = 3, memo = {0:1, 1:1, 2:1, 3:1, 4:1, 5:2, 8:1}
 - i = 7, nums[7] =  1, presum =  6, count = 4, memo = {0:1, 1:1, 2:1, 3:1, 4:1, 5:2, 6:1, 8:1}
 - i = 8, nums[8] =  4, presum = 10, count = 6, memo = {0:1, 1:1, 2:1, 3:1, 4:1, 5:2, 6:1, 8:1, 10:1}
 - i = 9, nums[9] =  0, presum = 10, count = 8, memo = {0:1, 1:1, 2:1, 3:1, 4:1, 5:2, 6:1, 8:1, 10:2}
```

- 如果 `nums` 是非負陣列的話，`presum` 只會單調遞增，即 `memo[presum]` 最多只會出現一次，且一定是連續的
- 即使 `nums[i]` 出現負數造成 `memo[presum]` 抵銷也沒關係，`memo` 的 `key(presum)` **一定是連續的**！
- 只要加入一個新的元素後又找到相同的 `presum` 時也可以放心累計，因為之前的 `subarray` 組合都可以再多 concatenate 一個新的元素變成新的 `subarrays`！