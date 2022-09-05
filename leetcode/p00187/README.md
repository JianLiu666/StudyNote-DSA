[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/repeated-dna-sequences/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 187. Repeated DNA Sequences

The **DNA sequence** is composed of a series of nucleotides abbreviated as `'A'`, `'C'`, `'G'`, and `'T'`.

- For example, `"ACGAATTCCG"` is a **DNA sequence**.

When studying **DNA**, it is useful to identify repeated sequences within the DNA.

Given a string `s` that represents a **DNA sequence**, return all the `10`**-letter-long** sequences (substrings) that occur more than once in a DNA molecule. You may return the answer in **any order**.

### Example 1:

```
Input: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
Output: ["AAAAACCCCC","CCCCCAAAAA"]
```

### Example 2:

```
Input: s = "AAAAAAAAAAAAA"
Output: ["AAAAAAAAAA"]
```

### Constraints:

- 1 <= s.length <= $10^5$
- `s[i]` is either `'A'`, `'C'`, `'G'`, or `'T'`.

### Related Topics

- Hash Table
- String
- Bit Manipulation
- Sliding Window
- Rolling Hash
- Hash Function
  
---

# 解題方向

### Solved using Brute Force concept

以長度為 10 的 substring 對 `s` 遍歷一遍，只要出現重複的 substring 就記錄下來

