[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/substring-with-concatenation-of-all-words/)
![image](https://img.shields.io/badge/Difficulty-Hard-red)

---

# 30. Substring with Concatenation of All Words

You are given a string `s` and an array of strings `words` of **the same length**. Return all starting indices of substring(s) in `s` that is a concatenation of each word in `words` **exactly once**, **in any order**, and **without any intervening characters**.

You can return the answer in **any order**.

### Example 1:

```
Input: s = "barfoothefoobarman", words = ["foo","bar"]
Output: [0,9]

Explanation:
 - Substrings starting at index 0 and 9 are "barfoo" and "foobar" respectively.
 - The output order does not matter, returning [9,0] is fine too.
```

### Example 2:

```
Input: s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
Output: []
```

### Example 3:

```
Input: s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
Output: [6,9,12]
```

### Constraints:

- 1 <= s.length <= $10^4$
- `1 <= words.length <= 5000`
- `1 <= words[i].length <= 30`
- `s` and `words[i]` consist of lowercase English letters.

### Related Topics

- Hash Table
- String
- Sliding Window
  
---

# 解題方向

這題最麻煩的地方除了 `words` 的組成順序可以任意置換之外，也可能出現在 `s` 重疊的部分，例如：

 - `s = "aaaaaaaaaa", words = ["aa", "aa"]`

我們可以分成兩個階段來處理：

**step.1 :**

實作 `substring` 的判斷邏輯，可以用 `Hash Table` 將 `words` 的出現頻率記錄下來加速查找

**step.2 :**

已知 `words` 內的每個單字長度皆相同，因此我們每次遍歷時可以用相同的長度來查找 `s` 內的單字，我們只要不斷以不同起點重複查找 `s`，即可解決掉 `s` 內單字重疊的部分。

詳細可以看程式碼註解