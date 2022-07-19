[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/reverse-words-in-a-string/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 151. Reverse Words in a String

Given an input string `s`, reverse the order of the **words**.

A **word** is defined as a sequence of non-space characters. The **words** in `s` will be separated by at least one space.

Return a string of the words in reverse order concatenated by a single space.

**Note** that `s` may contain leading or trailing spaces or multiple spaces between two words. The returned string should only have a single space separating the words. Do not include any extra spaces.

### Example 1:

```
Input: s = "the sky is blue"
Output: "blue is sky the"
```

### Example 2:

```
Input: s = "  hello world  "
Output: "world hello"

Explanation:
 - Your reversed string should not contain leading or trailing spaces.
```

### Example 3:

```
Input: s = "a good   example"
Output: "example good a"

Explanation:
 - You need to reduce multiple spaces between two words to a single space in the reversed string.
```

### Constraints:

- 1 <= s.length <= $10^4$
- `s` contains English letters (upper-case and lower-case), digits, and spaces `' '`.
- There is **at least one** word in `s`.

#### Follow-up:

If the string data type is mutable in your language, can you solve it **in-place** with `O(1)` extra space?

### Related Topics

- Two Pointers
- String
  
---

# 解題方向

不管是 Golang 或 Python 的 `string` 都是 **immutable** 的，所以必須用一個 `string` 取代。

因為是要反轉這個 `string`，所以只要從後往前遍歷整個 `string`，過程中注意每個單字的邊界判斷即可。