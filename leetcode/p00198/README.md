[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/house-robber/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 198. House Robber

You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and **it will automatically contact the police if two adjacent houses were broken into on the same night.**

Given an integer array `nums` representing the amount of money of each house, return the maximum amount of money you can rob tonight **without alerting the police.**

### Example 1:

```
Input: nums = [1,2,3,1]
Output: 4
Explanation:
 - Rob house 1 (money = 1) and then rob house 3 (money = 3).
 - Total amount you can rob = 1 + 3 = 4.
```

### Example 2:

```
Input: nums = [2,7,9,3,1]
Output: 12
Explanation: 
 - Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
 - Total amount you can rob = 2 + 9 + 1 = 12.
```

### Constraints:

- `1 <= nums.length <= 100`
- `0 <= nums[i] <= 400`

### Related Topics

- Array
- Dynamic Programming
  
---

# 解題方向

經典的 Dynamic Programming 問題，第一次解還有點不熟悉，記錄一下思考邏輯：

我們的目標是搶完整條街之後拿到最多的錢，因此對於最後一間 `n` 只有兩種可能：搶與不搶

在把這個問題往下延伸就會變成：
- 搶了最後一間，表示結束時他身上的錢會有 `nums[n]+rob[n-2]` 這麼多 (`rob` 為當下的累計總額)
- 不搶最後一間，表示結束時他身上的錢會有 `rob[n-1]` 這麼多

因此，我們可以將最終的結果以這個等式描述：`rob_maximum = max(nums[n]+rob[n-2], rob[n-1])`

確定完結果，我們接著繼續來看 `rob`：
- `rob[0]` 只有一種可能，搶就對了，即 `rob[0] == nums[0]`
- `rob[1]` 就有兩種可能了，不是搶第一間，就是搶第二間，錢多的算，即 `rob[1] = max(nums[1], nums[0])`

將上述邏輯已遞迴的形式表示，如下：

```python
# Time Complexity: O(2^n)
# Space Complexity: O(n)
def house_robber(nums: List[int]) -> int:
    def rob(n: int) -> int:
        if n == 0:
            return nums[0]
        if n == 1:
            return max(nums[1], nums[0])
    
        return max(nums[n]+rob[n-2], rob[n-1])
    
    return rob(len(nums)-1)
```

將已經算過的結果紀錄在 Hash Map，利用空間優化一下時間複雜度

```python
# Time Complexity: O(n)
# Space Complexity: O(n)
def house_robber(nums: List[int]) -> int:
    memo = {}
    def rob(n: int) -> int:
        if n == 0:
            return nums[0]
        if n == 1:
            return max(nums[1], nums[0])
        if n in memo:
            return memo[n]

        memo[n] = max(nums[n]+rob[n-2], rob[n-1])
        return memo[n]
    
    return rob(len(nums)-1)
```

也許我們可以換個形式表示

```python
#        +-----+-----+-----+-----+-----+
# nums = |  2  |  7  |  9  |  3  |  1  |
#        +-----+-----+-----+-----+-----+
#
#                max(2,7)    max(3+7,11)
#        +-----+-----+-----+-----+-----+
# robs = |  2  |  7  |  11 |  11 |  12 |
#        +-----+-----+-----+-----+-----+
#                      max(9+2,7)  max(1+11, 11)

# Time Complexity: O(n)
# Space Complexity: O(n)
def house_robber(nums: List[int]) -> int:
    robs = [0] * len(nums)
    robs[0] = nums[0]
    robs[1] = max(nums[1], nums[0])

    for i in range(2, len(nums)):
        robs[i] = max(nums[i]+robs[i-2], robs[i-1])

    return robs[-1]
```
