[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/letter-tile-possibilities/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 1079. Letter Tile Possibilities

You have `n` `tiles`, where each tile has one letter `tiles[i]` printed on it.

Return the number of possible non-empty sequences of letters you can make using the letters printed on those `tiles`.

### Example 1:

```
Input: tiles = "AAB"
Output: 8

Explanation:
 - The possible sequences are "A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA".
```

### Example 2:

```
Input: tiles = "AAABBC"
Output: 188
```

### Example 3:

```
Input: tiles = "V"
Output: 1
```

### Constraints:

- `1 <= tiles.length <= 7`
- `tiles` consists of uppercase English letters.

### Related Topics

- String
- Backtracking
  
---

# 解題方向

### Solved using Backtracking concept

Enumeration 的想法是每次挑出一個字母後把剩餘的 `tiles` 繼續往下傳遞，直到 `tiles == ""` 時終止

有鑒於題目允許 `tiles` 出現重複字母，因此需要透過 `Hash Table` 將已經走過的路提早結束 (pruning)

### Solved using Depth-First Search concept

在 Golang Submission Details 看到一個用 DFS 解決的方法，詳細直接看程式碼吧：

```go
func dfs(letters [26]int) int {
    sum := 0
    
    // 長度為 ['A', 'Z'] = 26
    for i := 0; i < 26; i++ {
        // 尚未未處理的字母
        if letters[i] > 0 {
            // 找到一種 sequence
            sum++
            // 移除自己的字母後, 再繼續往下查找
            letters[i]--
            sum += dfs(letters)
            // 把自己的字母加回來，提供後續其他字母使用
            letters[i]++
        }
    }

    return sum
}

```