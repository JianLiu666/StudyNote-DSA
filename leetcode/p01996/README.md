[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/the-number-of-weak-characters-in-the-game/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 1996. The Number of Weak Characters in the Game

You are playing a game that contains multiple characters, and each of the characters has **two** main properties: **attack** and **defense**. You are given a 2D integer array `properties` where `properties[i]` = [$\text{attack}_\text{i}$, $\text{defense}_\text{i}$] represents the properties of the $\text{i}^\text{th}$ character in the game.

A character is said to be **weak** if any other character has **both** attack and defense levels **strictly greater** than this character's attack and defense levels. More formally, a character `i` is said to be **weak** if there exists another character `j` where $\text{attack}_\text{j}$ > $\text{attack}_\text{i}$ and $\text{defense}_\text{j}$ > $\text{defense}_\text{i}$.

Return the number of **weak** characters.

### Example 1:

```
Input: properties = [[5,5],[6,3],[3,6]]
Output: 0

Explanation:
 - No character has strictly greater attack and defense than the other.
```

### Example 2:

```
Input: properties = [[2,2],[3,3]]
Output: 1

Explanation:
 - The first character is weak because the second character has a strictly greater attack and defense.
```

### Example 3:

```
Input: properties = [[1,5],[10,4],[4,3]]
Output: 1

Explanation:
 - The third character is weak because the second character has a strictly greater attack and defense.
```

### Constraints:

- 2 <= properties.length <= $10^5$
- `properties[i].length == 2`
- 1 <= $\text{attack}_\text{i}$, $\text{defense}_\text{i}$ <= $10^5$

### Related Topics

- Array
- Stack
- Greedy
- Sorting
- Monotonic Stack
  
---

# 解題方向

### Solved using Sorting concept

在這題的 LeetCodeCN [官方解法](https://leetcode.cn/problems/the-number-of-weak-characters-in-the-game/solution/you-xi-zhong-ruo-jiao-se-de-shu-liang-by-3d2g/) 中看到的解題思路，懶人包是對 `properties` 做排序

-  確保 `attack` 由大到小排序
-  `attack` 相同時，`defense` 由小到大排序

如此一來我們就可以確定接下來對 `properties` 遍歷的過程中，`attack` 始終都會是最大的，那麼在遍歷的過程中就只需要比較 `defense` 的大小即可:

 - `e.g. properties = [[5,1],[5,2],[5,3],[4,1],[4,2],[3,1]]`

如上所示，我們在遍歷的過程中只要記下當前能夠找到最大的 `defense`，就可以比對出 **Weak Character** 了

### Solved using Monotonic Stack concept 

一樣是來自 LeetCodeCN [官方解法](https://leetcode.cn/problems/the-number-of-weak-characters-in-the-game/solution/you-xi-zhong-ruo-jiao-se-de-shu-liang-by-3d2g/)

差別在於用一個由小到大排序的 Monotonic Stack 來維護 `defense`，也就是說當 `properties[j] > properties[i] ` 時就是我們要找到的 **Week Character**