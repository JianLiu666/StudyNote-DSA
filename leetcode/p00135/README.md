[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/candy/)
![image](https://img.shields.io/badge/Difficulty-Hard-red)

---

# 135. Candy

自己看一下

### Related Topics

- Array
- Greedy
  
---

# 解題方向

為了確定每個小孩拿到的糖果數量都要符合限制，所以我們必須從 rating 最小的孩子開始發糖果，依序發到 rating 最大的孩子身上  
這樣一來就可以確保當我們在發糖果給第 `ratings[i]` 的孩子時，如果左右邊的孩子 rating 比他小時，我們知道要多發幾個糖果給現在這個孩子  

因此初步的思想就是用 min-heap 來取出現在要發的孩子是 `ratings[i]`  

接著有以下幾種可能性，身邊孩子的 rating 可能比自己小、跟自己一樣大、比自己大，依序列出來之後會像這樣：
- `[ v . v ]`: 當左右的 rating 都比自己小時，為了符合條件我們需要找出最大的 `max(left, right)` 後在多發一個糖果給 `ratings[i]` 的孩子
- `[ v . = ] [ = . v ] [ v . ^ ] [ ^ . v ]`: 這些情況下，我們要知道比自己 rating 還小的孩子拿到多少糖果，在多發一個給 `ratings[i]`
- `[ = . = ] [ = . ^ ] [ ^ . = ] [ ^ . ^ ]`: 這些情況下，我們都是最小的 rating，所以只需要發一顆糖果給 `ratings[i]` 就好


列舉完之後，剩下直接看程式碼吧  