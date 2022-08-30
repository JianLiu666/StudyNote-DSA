[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/maximum-xor-of-two-numbers-in-an-array/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 421. Maximum XOR of Two Numbers in an Array

Given an integer array `nums`, return the maximum result of `nums[i] XOR nums[j]`, where `0 <= i <= j < n`.

### Example 1:

```
Input: nums = [3,10,5,25,2,8]
Output: 28

Explanation:
 - The maximum result is 5 XOR 25 = 28.
```

### Example 2:

```
Input: nums = [14,70,53,83,49,91,36,80,92,51,66,70]
Output: 127
```

### Constraints:

- 1 <= nums.length <= 2 * $10^5$
- 0 <= nums[i] <= $2^{31}$ - 1

### Related Topics

- Array
- Hash Table
- Bit Manipulation
- Trie
  
---

# 解題方向

### Solved using Trie concept

LeetCodeCN 有兩篇講得很不錯的討論，[這篇](https://leetcode.cn/problems/maximum-xor-of-two-numbers-in-an-array/solution/gong-shui-san-xie-noxiang-xin-ke-xue-xi-bmjdg/)跟[這篇](https://leetcode.cn/problems/maximum-xor-of-two-numbers-in-an-array/solution/python3-qiao-miao-li-yong-qian-zhui-shu-0alcy/)

核心概念是如果要讓 `nums[i] ^ nums[j]` 後的結果越大，就要找出重複性越低且 **XOR** 後 `1` 最多的兩數，這個搜尋過程就可以透過 `Trie` 的概念來處理：

- 第一步先將 `nums` 內的所有元素新增至 `Trie` 內 (Binary 表示)
- 除非子節點只有相同位元，否則查找時盡量往相異位元的分支尋訪，讓當前位元最後的 **XOR** 結果能夠始終保持為 `1`
  - e.g. `1 -> 0` 或 `0 -> 1`

剩下的就直接看程式碼吧