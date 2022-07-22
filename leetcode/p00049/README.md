[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/group-anagrams/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 49. Group Anagrams

Given an array of strings `strs`, group **the anagrams** together. You can return the answer in **any order**.

An **Anagram** is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

### Example 1:

```
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
```

### Example 2:

```
Input: strs = [""]
Output: [[""]]
```

### Example 3:

```
Input: strs = ["a"]
Output: [["a"]]
```

### Constraints:

- 1 <= strs.length <= $10^4$
- `0 <= strs[i].length <= 100`
- `strs[i]` consists of lowercase English letters.

### Related Topics

- Array
- Hash Table
- String
- Sorting
  
---

# 解題方向

這題的重點在如何設計一個適合的 `key` 將所有符合 anagrams 的單字關聯在一起，思路如下：

- 將單字重新排序，因為範圍已知且不須在意順序，所以可以直接用 `couting sort` 解決

決定好怎麼設計 anagrams key 之後，就只需把整個陣列尋訪一次做好分類即可
