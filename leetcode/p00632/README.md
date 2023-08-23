[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/)
![image](https://img.shields.io/badge/Difficulty-Hard-red)

---

# 632. Smallest Range Covering Elements from K Lists

自己看一下

### Related Topics

- Array
- Hash Table
- Greedy
- Sliding Window
- Sorting
- Heap (Priority Queue)

---

# 解題方向

這題也是 min heap 的應用，詳細思路官神說得很清楚，可以直接參考他的[影片](https://www.youtube.com/watch?v=ejVD92bJe34)  

這裡我只記錄一些重點，如下：

```
A: [4,10,15,24,26]
B: [0,9,12,20]
C: [5,18,22,30]

我們要做的事情是確保所有的 array (A, B, C) 的元素都至少有一個在 heap 內，即:

[0, 4, 5]

接著每次都 pop 出一個值，看彈出去的是哪個陣列的元素，就要再從那個陣列中補上下一個元素進來，直到有某個陣列已經不能再補值了
此時 min-heap 中維護的元素頭跟尾就會是我們需要的範圍端點
```

golang 沒有辦法像 c++ 一樣用一個 set 輕巧的解決, 所以 tricky 的解法請直接參考程式碼