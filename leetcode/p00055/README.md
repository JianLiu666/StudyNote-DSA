[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/jump-game/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 55. Jump Game

You are given an integer array `nums`. You are initially positioned at the array's **first index**, and each element in the array represents your maximum jump length at that position.

Return `true` if you can reach the last index, or `false` otherwise.

### Example 1:

```
Input: nums = [2,3,1,1,4]
Output: true

Explanation:
 - Jump 1 step from index 0 to 1, then 3 steps to the last index.
```

### Example 2:

```
Input: nums = [3,2,1,0,4]
Output: false

Explanation:
 - You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.
```

### Constraints:

- 1 <= `nums.length` <= $10^4$
- 0 <= `nums[i]` <= $10^5$

### Related Topics

- Array
- Dynamic Programming
- Greedy
  
---

# 解題方向

**Bottom Up**

我們的重點只要放在能夠走到得最遠距離，一旦當前位置超出能夠走到的最遠距離時就能提早結束

反之，一旦最遠距離超過 `nums` 的長度時，也可以提早結束