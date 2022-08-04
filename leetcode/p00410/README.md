[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/split-array-largest-sum/)
![image](https://img.shields.io/badge/Difficulty-Hard-red)

---

# 410. Split Array Largest Sum

Given an array `nums` which consists of non-negative integers and an integer `m`, you can split the array into `m` non-empty continuous subarrays.

Write an algorithm to minimize the largest sum among these `m` subarrays.

### Example 1:

```
Input: nums = [7,2,5,10,8], m = 2
Output: 18

Explanation:
 - There are four ways to split nums into two subarrays.
 - The best way is to split it into [7,2,5] and [10,8],
 - where the largest sum among the two subarrays is only 18.
```

### Example 2:

```
Input: nums = [1,2,3,4,5], m = 2
Output: 9
```

### Example 3:

```
Input: nums = [1,4,4], m = 3
Output: 4
```

### Constraints:

- `1 <= nums.length <= 1000`
- 0 <= nums[i] <= $10^6$
- `1 <= m <= min(50, nums.length)`

### Related Topics

- Array
- Binary Search
- Dynamic Programming
- Greedy
  
---

# 解題方向

建議直接拜讀這位大神的[文章](https://leetcode.com/problems/split-array-largest-sum/discuss/89819/C%2B%2B-Fast-Very-clear-explanation-Clean-Code-Solution-with-Greedy-Algorithm-and-Binary-Search)，我就只簡單記錄一下懶人包了

這題考的是 `Binary Search` 的延伸應用，我們的目標不是 `target index` 而是一段 `target range` (the smallest sum among the `m` subarrays)

按照大神的思路先定定出 `[low, high]` 的區間，舉例 `array = [1,2,3,4,5]`：

- `low` : maximum of the single element, i.e., `[[1],[2],[3],[4],[5]]`, the maximum = `5`
- `high` : sum of whole array elements, i.e., `[[1,2,3,4,5]]` = `15`

換句話說，不管 `m` 等於多少，結果一定會落在 `[5, 15]` 的區間內，接著就只要考慮如何實作 `cut function` 即可

- 已知目標值是 `mid`，只要 `acc+nums[i]` 還沒超過 `mid` 就持續累計，反之就當作切出一個新的 subarray 重新累計

再說一次，真的太神了...