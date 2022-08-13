[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/valid-palindrome/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 125. Valid Palindrome

A phrase is a **palindrome** if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string `s`, return `true` if it is a **palindrome**, or `false` otherwise.

### Example 1:

```
Input: s = "A man, a plan, a canal: Panama"
Output: true

Explanation:
 - "amanaplanacanalpanama" is a palindrome.
```

### Example 2:

Input: s = "race a car"
Output: false
Explanation:
 - "raceacar" is not a palindrome.

### Example 3:

```
Input: s = " "
Output: true
Explanation:
 - s is an empty string "" after removing non-alphanumeric characters.
 - Since an empty string reads the same forward and backward, it is a palindrome.
```

### Constraints:

- 1 <= s.length <= 2 * $10^5$
- `s` consists only of printable ASCII characters.

### Related Topics

- Two Pointers
- String
  
---

# 解題方向

### Solved using Two Pointers concept 

head and tail pointers, 只要注意當前字元是否合法即可:

- `[0, 9]`, `[A, Z]`, `[a, z]`