[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/reduce-array-size-to-the-half/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 1338. Reduce Array Size to The Half

You are given an integer array `arr`. You can choose a set of integers and remove all the occurrences of these integers in the array.

Return the minimum size of the set so that **at least** half of the integers of the array are removed.

### Example 1:

```
Input: arr = [3,3,3,3,5,5,5,2,2,7]
Output: 2
Explanation:
 - Choosing {3,7} will make the new array [5,5,5,2,2] which has size 5 (i.e equal to half of the size of the old array).
 - Possible sets of size 2 are {3,5},{3,2},{5,2}.
 - Choosing set {2,7} is not possible as it will make the new array [3,3,3,3,5,5,5] which has a size greater than half of the size of the old array.
```

### Example 2:

```
Input: arr = [7,7,7,7,7,7]
Output: 1

Explanation:
 - The only possible set you can choose is {7}. This will make the new array empty.
```

### Constraints:

- 2 <= arr.length <= $10^5$
- `arr.length` is even.
- 1 <= arr[i] <= $10^5$


### Related Topics

- Array
- Hash Table
- Greedy
- Sorting
- Heap (Priority Queue)
  
---

# 解題方向

### Solved using Sorting concept

先對 `arr` 遍歷一遍並用 `Hash Table` 將每筆資料的出現頻率統計完畢後，再以出現頻率為主重新排序一遍 `Hash Table` 裡面的 unique keys

最後逐步以排序後的 `key` 遞減直到資料筆數低於 `arr` 的一半即可

### Solved using Hash Table

1. 先對 `arr` 遍歷一遍並用 `Hash Table` 統計每筆資料的出現頻率，遍歷的過程中要順便擦掉 `arr` 的原始資料
2. 以 `Hash Table` 內的出現頻率為 **index**，填回 `arr`，把 `arr` 視為出現頻率的統計表，如下：
    - e.g. 
      - `arr = [1,2,3,3,3,3]`
      - `hash = {1:1, 2:1, 3:4}`
      - `new_arr = [2,0,0,1,0,0]`，即有兩筆資料出現1次`(index=0)`，有一筆資料出現4次`(index=3)`
3. 接著只要從 `arr` 尾巴遍歷回來，逐步減去出現次數直到資料筆數低於 `arr` 的一半即可
