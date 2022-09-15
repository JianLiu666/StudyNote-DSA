[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/find-original-array-from-doubled-array/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 2007. Find Original Array From Doubled Array

An integer array `original` is transformed into a **doubled** array `changed` by appending **twice the value** of every element in `original`, and then randomly **shuffling** the resulting array.

Given an array `changed`, return `original` if `changed` is a **doubled** array. If `changed` is not a **doubled** array, return an empty array. The elements in `original` may be returned in **any** order.

### Example 1:

```
Input: changed = [1,3,4,2,6,8]
Output: [1,3,4]

Explanation: One possible original array could be [1,3,4]:
 - Twice the value of 1 is 1 * 2 = 2.
 - Twice the value of 3 is 3 * 2 = 6.
 - Twice the value of 4 is 4 * 2 = 8.
 Other original arrays could be [4,3,1] or [3,1,4].
```

### Example 2:

```
Input: changed = [6,3,0,1]
Output: []

Explanation:
 - changed is not a doubled array.
```

### Example 3:

```
Input: changed = [1]
Output: []

Explanation:
 - changed is not a doubled array.
```

### Constraints:

- 1 <= changed.length <= $10^5$
- 0 <= changed[i] <= $10^5$

### Related Topics

- Array
- Hash Table
- Greedy
- Sorting
  
---

# 解題方向

只要對 `changed` 從小到大排序後可以保證，如果存在 `original` 時，那 `changed` 的第一筆資料一定是 `original` 的第一筆資料

接著在用 `Hash Table` 紀錄 `changed` 內所有元素的出現次數作為第二次遍歷時的參考

遍歷過程中只要不斷確認 `changed[i]` 與 `changed[i*2]` 是否同時存在於 `Hash Table` 即可確認是否存在 `original`