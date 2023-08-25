[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/maximal-square/description/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 221. Maximal Square

自己看一下

### Related Topics

- Array
- Dynamic Programming
- Matrix
  
---

# 解題方向

DP 經典題，這裡要探討的是如何確定最優子結構怎麼保存

對一個正方形而言, 只要我們確定在 `matrix[row][col]` 這個位置上相對應的 左/上/左上 位置都能夠成正方形, 那就保證這個位置也能夠成正方形
 - i.e. `matrix[row-1][col] > 0` and `matrix[row][col-1] > 0` and `matrix[row-1][col-1] > 0`

而且我們要取的是最小值, 避免發生以下狀況

```
1 0 1 1 此時最大的正方形面積只有 9, 如果我們取鄰近邊的最大值, 那就會變成 16
1 1 1 1 
1 1 1 1
1 1 1 1
```

剩下的直接看程式碼吧