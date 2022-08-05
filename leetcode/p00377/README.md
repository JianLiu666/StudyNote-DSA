[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/combination-sum-iv/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 377. Combination Sum IV

Given an array of **distinct** integers `nums` and a target integer `target`, return the number of possible combinations that add up to `target`.

The test cases are generated so that the answer can fit in a **32-bit** integer.

### Example 1:

```
Input: nums = [1,2,3], target = 4
Output: 7

Explanation:
 - The possible combination ways are:
    - (1, 1, 1, 1)
    - (1, 1, 2)
    - (1, 2, 1)
    - (1, 3)
    - (2, 1, 1)
    - (2, 2)
    - (3, 1)
 - Note that different sequences are counted as different combinations.
```

### Example 2:

```
Input: nums = [9], target = 3
Output: 0
```

### Constraints:

- `1 <= nums.length <= 200`
- `1 <= nums[i] <= 1000`
- All the elements of `nums` are **unique**.
- `1 <= target <= 1000`

### Related Topics

- Array
- Dynamic Programming
  
---

# 解題方向

### Solved using Recursion concept

題目的懶人包就是找出有幾種方法可以湊出加總起來等於 `target` 的運算式，用 Recursion 的表達邏輯就會像這樣：

- 依序將每個數字都拿出來試試看，不斷減去直到 `target == 0`，表示又找到一個運算式了！

```python
def combinationSum4(nums: List[int], target: int) -> int:
   if target == 0:
      return 1
   
   res = 0
   for n in nums:
      if target >= n:
         res += combinationSum4(nums, target-n)
   
   return res
```

但這種簡單暴力的做法基本上下一秒就會遇到 stack overflow 的問題了，節哀

### Solved using Dynamic Programming concept

**Top-down solution**

其實就是 recursion + memoization

Top-down 片段程式碼如下：

```python
# @param memo is the memoization in range from 0 to target, this array init value is -1
# @param nums is the problem given parameter
# @param target is the problem given parameter
#
# @return int is the accumulation of the combinations of the target
def helper(memo: List[int], nums: List[int], target: int) -> int:
   if memo[target] != -1:
      return memo[target]
   
   res = 0
   for n in nums:
      if target >= n:
         helper(memo, nums, target-n)

   memo[target] = res
   return res
```

**Bottom-up solution**

從 `target = 0` 時開始出發，一路計算當 `target = 1, 2, 3, 4, ... ..., n` 時的結果

因為每次計算時都有可能參考到過去任何一刻的計算結果，所以必須將每一次的 `target` 都保存起來

程式碼如下：

```python
def combinationSum4(nums: List[int], target: int) -> int:
   memo = [0 for _ in range(target+1)]
   memo[0] = 1

   for i in range(1, target+1, 1):
      for n in nums:
            if i-n >= 0:
               memo[i] += memo[i-n]
   
   return memo[target]
```