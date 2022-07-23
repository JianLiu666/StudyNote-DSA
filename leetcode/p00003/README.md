[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/longest-substring-without-repeating-characters/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 3. Longest Substring Without Repeating Characters

Given a string `s`, find the length of the **longest substring** without repeating characters.

### Example 1:

```
Input: s = "abcabcbb"
Output: 3

Explanation:
 - The answer is "abc", with the length of 3.
```

### Example 2:

```
Input: s = "bbbbb"
Output: 1

Explanation:
 - The answer is "b", with the length of 1.
```

### Example 3:

```
Input: s = "pwwkew"
Output: 3

Explanation:
 - The answer is "wke", with the length of 3.
 - Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
```

### Constraints:

- 0 <= s.length <= 5 * $10^4$
- `s` consists of English letters, digits, symbols and spaces.

### Related Topics

- Hash Table
- String
- Sliding Window
  
---

# 解題方向

題目本身不難，只要注意好邊界條件的判斷即可：

- 不重複的起始點可能在 `cursor` 的前面，例如 `abcadefg`
- 重複的字母可能在起始點之前，與現在判定的 `substring` 無關，例如 `tmmzuxt`
