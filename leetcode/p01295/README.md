[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/find-numbers-with-even-number-of-digits/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 1295. Find Numbers with Even Number of Digits

Given an array `nums` of integers, return how many of them contain an **even number** of digits.

### Example 1:

```
Input: nums = [12,345,2,6,7896]
Output: 2

Explanation: 
 - 12 contains 2 digits (even number of digits). 
 - 345 contains 3 digits (odd number of digits). 
 - 2 contains 1 digit (odd number of digits). 
 - 6 contains 1 digit (odd number of digits). 
 - 7896 contains 4 digits (even number of digits). 
 - Therefore only 12 and 7896 contain an even number of digits.
```

### Example 2:

```
Input: nums = [555,901,482,1771]
Output: 1 

Explanation: 
 - Only 1771 contains an even number of digits.
```

### Constraints:

- `1 <= nums.length <= 500`
- 1 <= nums[i] <= $10^5$

### Related Topics

- Array
  
---

# 解題方向

#### Solved using string type

轉成 `string` 後檢查字元長度

e.g.
```python
if len(str(n))%2 == 0:
    count += 1
```

#### Solved using problem constaints

根據題目範圍 `1 - 100000`，設計條件判斷

e.g.
```python
if n >= 10 and n <= 99:
    count += 1
elif n >= 1000 and n <= 9999:
    count += 1
elif n == 100000:
    count += 1
```

#### Solved using division

連續整除 10 直到數字歸零後，判斷除法次數

e.g.
```python
division = 0
for n > 0:
    n //= 10
    division += 1

if division % 2 == 0:
    count += 1
```

#### Solved using logarithm and bit manipulation

$log{_{10}}{n}$ 後就可以知道數字長度，接著在對結果做做 AND 1 運算

e.g.
```python
count += int(math.log10(n)) & 1
```