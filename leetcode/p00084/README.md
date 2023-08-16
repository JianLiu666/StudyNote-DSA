[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/largest-rectangle-in-histogram/)
![image](https://img.shields.io/badge/Difficulty-Hard-red)

---

# 84. Largest Rectangle in Histogram

Given an array of integers `heights` representing the histogram's bar height where the width of each bar is `1`, return the area of the largest rectangle in the histogram.

### Example 1:

![image](./image/histogram.jpeg)

```
Input: heights = [2,1,5,6,2,3]
Output: 10

Explanation:
 - The above is a histogram where width of each bar is 1.
 - The largest rectangle is shown in the red area, which has an area = 10 units.
```

### Example 2:

![image](./image/histogram-1.jpeg)

```
Input: heights = [2,4]
Output: 4
```

### Constraints:

- 1 <= heights.length <= $10^5$
- 0 <= heights[i] <= $10^4$

### Related Topics

- Array
- Stack
- Monotonic Stack
  
---

# 解題方向

### Solved using Stack concept

這題的解題思路有點刁鑽，自己想半天想不知道怎麼切入這問題，後來還是只能先找大神的作法來學習QQ

- LeetCodeCN [精選解答](https://leetcode.cn/problems/largest-rectangle-in-histogram/solution/84-by-ikaruga/)

看圖理解後的推導過程就不難了，剩下的直接在程式碼裡面說明吧

#### 2023.08.17

[multi-pass solution](https://www.youtube.com/watch?v=mesaogfSjD4) 請參考官神前半段講解