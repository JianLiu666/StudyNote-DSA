[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/next-permutation/description/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 31. Next Permutation

自己看一下

### Related Topics

- Array
- Two Pointers
  
---

# 解題方向

這題的考點在如何找到下一個按照 lexicographically 排序後大一點的數字

如下所示：

```
4 6 5 4 3 2
  < < < < <

如果數字一直是遞增的，那我們連換的資格都沒有

4 6 5 4 3 2
i   j
一旦發現不是遞增之後, 我們就可以嘗試從尾巴開始在找一次, 找到比現在這個定位點(i)更大的數字後(j)跟他交換

5 6 4 4 3 2
i
接著只要再把 i+1 以後的數字以 ascending order 排序過一遍後就是剛好比原本的 nums 還要再大一點的順序了
```