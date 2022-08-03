[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/find-smallest-letter-greater-than-target/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 744. Find Smallest Letter Greater Than Target

Given a characters array `letters` that is sorted in **non-decreasing** order and a character `target`, return the smallest character in the array that is larger than `target`.

**Note** that the letters wrap around.

- For example, if `target == 'z'` and `letters == ['a', 'b']`, the answer is `'a'`.

### Example 1:

```
Input: letters = ["c","f","j"], target = "a"
Output: "c"
```

### Example 2:

```
Input: letters = ["c","f","j"], target = "c"
Output: "f"
```

### Example 3:

```
Input: letters = ["c","f","j"], target = "d"
Output: "f"
```

### Constraints:

- 2 <= letters.length <= $10^4$
- `letters[i]` is a lowercase English letter.
- `letters` is sorted in **non-decreasing** order.
- `letters` contains at least two different characters.
- `target` is a lowercase English letter.

### Related Topics

- Array
- Binary Search
  
---

# 解題方向

Binary Search 暖身題，只要注意收斂時的判斷即可

- `if letters[mid] <= target: low = mid + 1` : 如同題目所述，我們要找到的是比 `target` 還要大的字元
- `if letters[mid] > target: high = mid` : 保留這個字元，繼續收斂

最終當 `left` 收斂到跟 `right` 相同時，就是我們要找的答案
