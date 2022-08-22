[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/koko-eating-bananas/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 875. Koko Eating Bananas

Koko loves to eat bananas. There are `n` piles of bananas, the $i^{th}$ pile has `piles[i]` bananas. The guards have gone and will come back in `h` hours.

Koko can decide her bananas-per-hour eating speed of `k`. Each hour, she chooses some pile of bananas and eats `k` bananas from that pile. If the pile has less than `k` bananas, she eats all of them instead and will not eat any more bananas during this hour.

Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.

Return the minimum integer `k` such that she can eat all the bananas within `h` hours.

### Example 1:

```
Input: piles = [3,6,7,11], h = 8
Output: 4
```

### Example 2:

```
Input: piles = [30,11,23,4,20], h = 5
Output: 30
```

### Example 3:

```
Input: piles = [30,11,23,4,20], h = 6
Output: 23
```

### Constraints:

- 1 <= piles.length <= $10^4$
- piles.length <= h <= $10^9$
- 1 <= piles[i] <= $10^9$

### Related Topics

- Array
- Binary Search
  
---

# 解題方向

這題算是 Binary Search 的變形題，我們的目的在找出既可以讓猴子悠哉的吃，又能夠趕在守衛回來之前全吃光的速度

直接破題，最慢的吃法就是一根一根的吃，最快的吃法就是一次吃完整串，即 :
- `min` : 1
- `max` : maximum of value in array

目的明確，再來看終止條件，我們不能只單單找到適合的速度就跳出，否則就會發生這種慘案

```
Input: piles = [1,1,1,999999999], h = 10
Output: 156250000
Expected: 142857143
```

原因在於一樣在規定的10小時內吃完，但 `Expected` 給出的速度在吃最後一串香蕉時可以幾乎剛好在7小時時限內吃完

因此即使找到適合的答案時，還是必須先暫時保留起來繼續縮小範圍，直到 `low == high` 時才可以結束查找