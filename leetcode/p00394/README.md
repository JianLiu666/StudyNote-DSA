[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/decode-string/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 394. Decode String

Given an encoded string, return its decoded string.

The encoding rule is: `k[encoded_string]`, where the `encoded_string` inside the square brackets is being repeated exactly `k` times. Note that k is guaranteed to be a positive integer.

You may assume that the input string is always valid; there are no extra white spaces, square brackets are well-formed, etc. Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, `k`. For example, there will not be input like `3a` or `2[4]`.

The test cases are generated so that the length of the output will never exceed $10^5$.

### Example 1:

```
Input: s = "3[a]2[bc]"
Output: "aaabcbc"
```

### Example 2:

```
Input: s = "3[a2[c]]"
Output: "accaccacc"
```

### Example 3:

```
Input: s = "2[abc]3[cd]ef"
Output: "abcabccdcdcdef"
```

### Constraints:

- `1 <= s.length <= 30`
- `s` consists of lowercase English letters, digits, and square brackets `'[]'`.
- `s` is guaranteed to be **a valid** input.
- All the integers in `s` are in the range `[1, 300]`.

### Related Topics

- String
- Stack
- Recursion
  
---

# 解題方向

用 `stack` 將我們遇到的每個數字或字串存起來，遇到 `]` 後對 `stack` 的資料進行運算。

在歷遍字串 `s` 時，我們總共有三件事情必須處理：
- 數字
  - 每當遇到數字時就存進 `stack`
- 連續字母
  - 如果 `stack` 是空的，表示這段字串不需要經過運算，可以直接輸出
  - 否則就必須把這段字串放進 `stack`，直到遇到 `]` 後拿出來處理
- `]`
  - 處理 `stack` 內的資料，有兩種情況要注意:
    - e.g. `3[a]`, 直接處理
    - e.g. `3[a4[e]]`, `3[4[e]a]`, 此時必須先將這些字串組合完後才能開始複製

