[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/push-dominoes/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 838. Push Dominoes

There are `n` dominoes in a line, and we place each domino vertically upright. In the beginning, we simultaneously push some of the dominoes either to the left or to the right.

After each second, each domino that is falling to the left pushes the adjacent domino on the left. Similarly, the dominoes falling to the right push their adjacent dominoes standing on the right.

When a vertical domino has dominoes falling on it from both sides, it stays still due to the balance of the forces.

For the purposes of this question, we will consider that a falling domino expends no additional force to a falling or already fallen domino.

You are given a string `dominoes` representing the initial state where:

- `dominoes[i] = 'L'`, if the $i^{th}$ domino has been pushed to the left,
- `dominoes[i] = 'R'`, if the $i^{th}$ domino has been pushed to the right, and
- `dominoes[i] = '.'`, if the $i^{th}$ domino has not been pushed.

Return a string representing the final state.

### Example 1:

```
Input: dominoes = "RR.L"
Output: "RR.L"

Explanation:
 - The first domino expends no additional force on the second domino.
```

### Example 2:

![image](./image/domino.png)

```
Input: dominoes = ".L.R...LR..L.."
Output: "LL.RR.LLRRLL.."
```

### Constraints:

- `n == dominoes.length`
- 1 <= `n` <= $10^5$
- `dominoes[i]` is either `'L'`, `'R'`, or `'.'`.

### Related Topics

- Two Pointers
- String
- Dynamic Programming
  
---

# 解題方向

解題思路我覺得 LeetCodeCN [這篇](https://leetcode.cn/problems/push-dominoes/solution/fu-xue-ming-zhu-miao-dong-xi-lie-xiang-x-xkts/)講的蠻不錯的，有圖解推推

只要理解題意，按照骨牌倒塌的原理依序將四種可能情境都實作即可：

- `[L ... L]` -> 這中間的骨牌會全都往左倒
- `[L ... R]` -> 這中間的骨牌會維持原樣
- `[R ... L]` -> 這中間的骨牌會往中心點倒
- `[R ... R]` -> 這中間的骨牌會全都往右倒