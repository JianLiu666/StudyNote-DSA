[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/count-vowels-permutation/)
![image](https://img.shields.io/badge/Difficulty-Hard-red)

---

# 1220. Count Vowels Permutation

Given an integer `n`, your task is to count how many strings of length `n` can be formed under the following rules:

 - Each character is a lower case vowel (`'a'`, `'e'`, `'i'`, `'o'`, `'u'`)
 - Each vowel `'a'` may only be followed by an `'e'`.
 - Each vowel `'e'` may only be followed by an `'a'` or an `'i'`.
 - Each vowel `'i'` may not be followed by another `'i'`.
 - Each vowel `'o'` may only be followed by an `'i'` or a `'u'`.
 - Each vowel `'u'` may only be followed by an `'a'`.

Since the answer may be too large, return it modulo $10^9$ + 7.

### Example 1:

```
Input: n = 1
Output: 5

Explanation:
 - All possible strings are: "a", "e", "i" , "o" and "u".
```

### Example 2:

```
Input: n = 2
Output: 10

Explanation:
 - All possible strings are: "ae", "ea", "ei", "ia", "ie", "io", "iu", "oi", "ou" and "ua".
```

### Example 3: 

```
Input: n = 5
Output: 68
```

### Constraints:

- 1 <= n <= 2 * $10^4$

### Related Topics

- Dynamic Programming
  
---

# 解題方向

### Solved using Dynamic Programming concept

跟經典的上樓梯問題其實很像，每一次的子問題都只跟上一次的結果有關，因此我們只需要給定一個 `fixed array` 來儲存 `a,e,i,o,u` 每一次的變化持續累加即可

### Soved using Matrix concept

透過矩陣快速冪方式實現

施工中 ...