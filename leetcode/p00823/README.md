[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/binary-trees-with-factors/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 823. Binary Trees With Factors

Given an array of unique integers, `arr`, where each integer `arr[i]` is strictly greater than `1`.

We make a binary tree using these integers, and each number may be used for any number of times. Each non-leaf node's value should be equal to the product of the values of its children.

Return the number of binary trees we can make. The answer may be too large so return the answer **modulo** $10^9$ + 7.

### Example 1:

```
Input: arr = [2,4]
Output: 3

Explanation:
 - We can make these trees: [2], [4], [4, 2, 2]
```

### Example 2:

```
Input: arr = [2,4,5,10]
Output: 7

Explanation:
 - We can make these trees: [2], [4], [5], [10], [4, 2, 2], [10, 2, 5], [10, 5, 2].
```

### Constraints:

- `1 <= arr.length <= 1000`
- 2 <= arr[i] <= $10^9$
- All the values of `arr` are **unique**.

### Related Topics

- Array
- Hash Table
- Dynamic Programming
  
---

# 解題方向

### Solved using Dynamic Programming concept

一樣是經典的 DP 問題，給定一個跟 `arr` 一樣大的 `counter` 來記錄已經算過的 `tree`，特別注意每個 `tree` 都有可能成為更大數字的 `subtree`，所以數字會變得很大，記得 **mod**

另外在找 child node 時為了加速查找，可以用 `Hash Table` 把 `arr` 每一個元素的 `index` 跟 `value` 記錄下來加速

剩下的就直接看程式碼吧