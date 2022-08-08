[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/longest-increasing-subsequence/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 300. Longest Increasing Subsequence

Given an integer array `nums`, return the length of the longest strictly increasing subsequence.

A **subsequence** is a sequence that can be derived from an array by deleting some or no elements without changing the order of the remaining elements. For example, `[3,6,2,7]` is a subsequence of the array `[0,3,1,6,2,2,7]`.

### Example 1:

```
Input: nums = [10,9,2,5,3,7,101,18]
Output: 4

Explanation:
 - The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
```

### Example 2:

```
Input: nums = [0,1,0,3,2,3]
Output: 4
```

### Example 3:

```
Input: nums = [7,7,7,7,7,7,7]
Output: 1
```

### Constraints:

- `1 <= nums.length <= 2500`
- -$10^4$ <= nums[i] <= $10^4$

#### Follow up:

Can you come up with an algorithm that runs in `O(n log(n))` time complexity?


### Related Topics

- Array
- Binary Search
- Dynamic Programming
  
---

# 解題方向

### Solved using Dynamic Programming concept

我自己認為這題的 DP 切入點不是這麼直觀好理解，假設 `dp[i]` 表示對應回 `nums[i]` 的 longest increasing subsequence，範圍在 `nums[0...i]` 之間。

接續這個想法，我們可以在每移動一次 `nums[i+1]` 時就比對一次從 `nums[0...i]` 的範圍，如果符合 longes increasing subsequence 的條件時就把之前的結果 `+1` 表示為新的 subsequence，pseudo code 如下：

```python
for i in range(len(nums)):
    for j in range(i):
        if nums[i] > nums[j] and dp[i] < dp[j] + 1:
            dp[i] = dp[j] + 1
```

### Solved using Binary Search with Greedy concept

我也覺得 Binary Search 不是這麼直觀可以想到的 *(總感覺是我自己太廢...)*，這位大神的 [Solution 2](https://leetcode.com/problems/longest-increasing-subsequence/discuss/1326308/C%2B%2BPython-DP-Binary-Search-BIT-Solutions-Picture-explain-O(NlogN)) 解釋得很清楚，我就不獻醜了

必須注意的是用這種方式做出來的 subsequence 不等於 `nums` 中真正的 longest increasing subsequence 順序，只是長度相等