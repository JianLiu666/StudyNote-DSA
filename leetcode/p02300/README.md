[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)
![image](https://img.shields.io/badge/Difficulty-Hard-red)

---

# 2300. Successful Pairs of Spells and Potions

自己看一下

### Related Topics

- Array
- Two Pointers
- Binary Search
- Sorting
  
---

# 解題方向

Binary Search 變形題，考點在 sorted array with duplicated values 中找出符合 target 的第一個數字  
剩下的直接看程式碼吧

自己也順便複習一下:  
- 如果要在 sorted array 中找到最接近的 index 該怎麼做?
- 如果可以完全命中的話那就直接回傳 `mid` 的位置就好
- 不然就是一樣移動 `mid`, 直到最後會發現最接近的 `index` 會落在 `left` 的位置
- 但如果是 sorted array with duplicated value 就要注意很有可能現在找到的這個 target 不是第一個, 所以還是要繼續搜尋
  - 直到 left, right 互相撞到之後看 left 停留的位置才會是答案
- 繼續延續, 所以最後 left-1 的位置就代表第一個小於 target 的數字