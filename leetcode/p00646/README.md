[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/maximum-length-of-pair-chain/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 646. Maximum Length of Pair Chain

You are given an array of `n` pairs `pairs` where pairs[i] = [$\text{left}_i$, $\text{right}_i$] and $\text{left}_i$ < $\text{right}_i$.

A pair `p2 = [c, d]` **follows** a pair `p1 = [a, b]` if `b < c`. A **chain** of pairs can be formed in this fashion.

Return the length longest chain which can be formed.

You do not need to use up all the given intervals. You can select pairs in any order.

### Example 1:

```
Input: pairs = [[1,2],[2,3],[3,4]]
Output: 2

Explanation:
 - The longest chain is [1,2] -> [3,4].
```

### Example 2:

```
Input: pairs = [[1,2],[7,8],[4,5]]
Output: 3

Explanation:
 - The longest chain is [1,2] -> [4,5] -> [7,8].
```

### Constraints:

- `n == pairs.length`
- `1 <= n <= 1000`
- -1000 <= $\text{left}_i$ < $\text{right}_i$ <= 1000

### Related Topics

- Array
- Dynamic Programming
- Greedy
- Sorting
  
---

# 解題方向

### Solved using Sorting concept

題目目的是找出最長的 Chain，因此我們必須從最小的 `pairs[i][1]` 開始挑起，才能串出最長的長度

因此只要將 `pairs` 以 `pairs[i][1] < pairs[j][1]` 為主將最小的 $\text{right}_i$ 往前靠，在對 `pairs` 遍歷一次計算長度即可