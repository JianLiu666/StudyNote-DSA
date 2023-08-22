[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/)
![image](https://img.shields.io/badge/Difficulty-Hard-red)

---

# 815. Bus Routes

自己看一下

### Related Topics

- Array
- Hash Table
- Breadth-First Search
  
---

# 解題方向

題目的問題是我們需要搭乘幾次公車後才能從 `source` 站牌抵達到 `target` 站牌  
題目所提供的只有公車的行駛路徑 `routes`, 我們需要自己把站牌對應的公車班次建立出來 `stop2bus`  

有了這個 graph 結構後我們就可以用 bfs 的方式找到最短路徑了  
詳細直接看程式碼吧  