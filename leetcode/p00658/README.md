[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/find-k-closest-elements/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 658. Find K Closest Elements

Given a **sorted** integer array `arr`, two integers `k` and `x`, return the `k` closest integers to `x` in the array. The result should also be sorted in ascending order.

An integer `a` is closer to `x` than an integer `b` if:

- `|a - x| < |b - x|`, or
- `|a - x| == |b - x|` and `a < b`

### Example 1:

```
Input: arr = [1,2,3,4,5], k = 4, x = 3
Output: [1,2,3,4]
```

### Example 2:

```
Input: arr = [1,2,3,4,5], k = 4, x = -1
Output: [1,2,3,4]
```

### Constraints:

- `1 <= k <= arr.length`
- 1 <= arr.length <= $10^4$
- `arr` is sorted in **ascending** order.
- -$10^4$ <= arr[i], x <= $10^4$

### Related Topics

- Array
- Two Pointers
- Binary Search
- Sorting
- Heap (Priority Queue)
  
---

# 解題方向

### Solved using Binary Search concept

`Binary Search` 變化題型，這次不只有找到目標值 `x`，而是要找到符合的範圍區間(i.e. 距離 `x` 最近的 `k` 個數字)

- 距離 `x` 越遠的數字其差距(i.e. `x - arr[i]`) 越大
  - e.g. `x - arr[i] > arr[i+k] - x`，表示左邊離中心較遠，因此必須向右移動