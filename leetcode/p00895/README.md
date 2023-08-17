[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/maximum-frequency-stack/description/)
![image](https://img.shields.io/badge/Difficulty-Hard-red)

---

# 895. Maximum Frequency Stack

自己看一下黑

### Related Topics

- Hash Table
- Stack
- Design
- Ordered Set
  
---

# 解題方向

這題的重點在怎麼同時維護 `key` 的出現頻率跟 `stack` 的思路  

我們可以用一個 `hash table` 紀錄 `key` 的出現頻率
在用一個 `hash table` 紀錄每個頻率的 `stack`，這樣就可以保證當同一個 `key` 一直放進去的同時，他都會存在在每一個頻率的 `stack` 內