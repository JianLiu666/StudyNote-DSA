[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/removing-stars-from-a-string/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 2390. Removing Stars From a String

You are given a string `s`, which contains stars `*`.

In one operation, you can:

- Choose a star in `s`.
- Remove the closest **non-star** character to its **left**, as well as remove the star itself.

Return the string after **all** stars have been removed.

**Note:**

- The input will be generated such that the operation is always possible.
- It can be shown that the resulting string will always be unique.

### Example 1:

```
Input: s = "leet**cod*e"
Output: "lecoe"

Explanation: Performing the removals from left to right:
 - The closest character to the 1st star is 't' in "leet**cod*e". s becomes "lee*cod*e".
 - The closest character to the 2nd star is 'e' in "lee*cod*e". s becomes "lecod*e".
 - The closest character to the 3rd star is 'd' in "lecod*e". s becomes "lecoe".
 - There are no more stars, so we return "lecoe".
```

### Example 2:

```
Input: s = "erase*****"
Output: ""

Explanation:
 -  The entire string is removed, so we return an empty string.
```

### Constraints:

- 1 <= s.length <= $10^5$
- `s` consists of lowercase English letters and stars `*`.
- The operation above can be performed on `s`.

### Related Topics

- String
- Stack
- Simulation
  
---

# 解題方向

### Solved using Array concept

打 [Weekly Contest 308](https://leetcode.com/contest/weekly-contest-308/) 時的做法，當時的想法是每一個 `*` 都會一併消除前一個字元，所以我就從 `s` 的尾巴往前遍歷，只要遇到 `*` 時就跳過一個字元，直到遍歷完整個 `s`

### Solved using Stack concept

不知道是不是第一次打週賽太緊張，當下完全沒想到這就是 `Stack` 🤦‍♂️

只要不是 `*` 的字元就加進 `Stack`，遇到 `*` 時就從 `Stack` 中移除掉最後一個，最後用 **FIFO** 的順序輸出即可