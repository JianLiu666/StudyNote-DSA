[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/target-sum/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 494. Target Sum

You are given an integer array `nums` and an integer `target`.

You want to build an **expression** out of nums by adding one of the symbols `'+'` and `'-'` before each integer in nums and then concatenate all the integers.

- For example, if `nums = [2, 1]`, you can add a `'+'` before `2` and a `'-'` before `1` and concatenate them to build the expression `"+2-1"`.

Return the number of different **expressions** that you can build, which evaluates to `target`.

### Example 1:

```
Input: nums = [1,1,1,1,1], target = 3
Output: 5

Explanation: 
 - There are 5 ways to assign symbols to make the sum of nums be target 3.
   -1 + 1 + 1 + 1 + 1 = 3
   +1 - 1 + 1 + 1 + 1 = 3
   +1 + 1 - 1 + 1 + 1 = 3
   +1 + 1 + 1 - 1 + 1 = 3
   +1 + 1 + 1 + 1 - 1 = 3
```

### Example 2:

```
Input: nums = [1], target = 1
Output: 1
```

### Constraints:

- `1 <= nums.length <= 20`
- `0 <= nums[i] <= 1000`
- `0 <= sum(nums[i]) <= 1000`
- `-1000 <= target <= 1000`

### Related Topics

- Array
- Dynamic Programming
- Backtracking
  
---

# 解題方向

### Solved using Depth-First Search Concept

DFS 練習題，先一次將相同符號(e.g. `+`) 的算式求和之後，再從最末端開始變向依序遞迴即可。

單純的 DFS 需要花的時間是 O($2^n$)，我們可以用 memorization 的概念將算過的結果記錄下來，利用空間來換取時間，最終可以把時間複雜度從原本的 O($2^n$) 降到 O(n)