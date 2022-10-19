[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/top-k-frequent-words/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 692. Top K Frequent Words

Given an array of strings `words` and an integer `k`, return the `k` most frequent strings.

Return the answer **sorted** by **the frequency** from highest to lowest. Sort the words with the same frequency by their **lexicographical order**.

### Example 1:

```
Input: words = ["i","love","leetcode","i","love","coding"], k = 2
Output: ["i","love"]

Explanation:
 - "i" and "love" are the two most frequent words.
 - Note that "i" comes before "love" due to a lower alphabetical order.
```

### Example 2:

```
Input: words = ["the","day","is","sunny","the","the","the","sunny","is","is"], k = 4
Output: ["the","is","sunny","day"]

Explanation:
 - "the", "is", "sunny" and "day" are the four most frequent words, with the number of occurrence being 4, 3, 2 and 1 respectively.
```

### Constraints:

- `1 <= words.length <= 500`
- `1 <= words[i].length <= 10`
- `words[i]` consists of lowercase English letters.
- `k` is in the range `[1, The number of unique words[i]]`

#### Follow-up: 

Could you solve it in `O(n log(k))` time and `O(n)` extra space?

### Related Topics

- Hash Table
- String
- Trie
- Sorting
- Heap (Priority Queue)
- Bucket Sort
- Counting
  
---

# 解題方向

### Solved using Priority Queue concept

`Max Heap` 暖身題, 注意排序規則即可

- 按出現頻率高低排序
- 頻率相同時，按字母順序排序
