[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 378. Kth Smallest Element in a Sorted Matrix

Given an `n x n` `matrix` where each of the rows and columns is sorted in ascending order, return the $k^{th}$ smallest element in the matrix.

Note that it is the $k^{th}$ smallest element in the sorted order, not the $k^{th}$ distinct element.

You must find a solution with a memory complexity better than O($n^2$).

### Example 1:

```
Input: matrix = [[1,5,9],[10,11,13],[12,13,15]], k = 8
Output: 13
Explanation:
 - The elements in the matrix are [1,5,9,10,11,12,13,13,15], and the 8th smallest number is 13
```

### Example 2:

```
Input: matrix = [[-5]], k = 1
Output: -5
```

### Constraints:

- `n == matrix.length == matrix[i].length`
- `1 <= n <= 300`
- -$10^9$ <= matrix[i][j] <= $10^9$
- All the rows and columns of `matrix` are **guaranteed** to be sorted in **non-decreasing order**.
- 1 <= k <= $n^2$

#### Follow up:

- Could you solve the problem with a constant memory (i.e., `O(1)` memory complexity)?
- Could you solve the problem in `O(n)` time complexity? The solution may be too advanced for an interview but you may find reading [this paper](http://www.cse.yorku.ca/~andy/pubs/X+Y.pdf) fun.


### Related Topics

- Array
- Binary Search
- Sorting
- Heap (Priority Queue)
- Matrix
  
---

# 解題方向

### Solved using Binary Search concept

`Binary Search` 的變化題，只能說這個解法真的跟鬼一樣，首先我們要了解 Binary Search 的 search space 有兩種面向：

- `index`: 基本用法，找到指定的 target(index)，例如 [153. Find Minimum in Rotated Sorted Array](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/)
- `range`: 進階用法，找到連續的 targets，例如 [287. Find the Duplicate Number
](https://leetcode.com/problems/find-the-duplicate-number/)

這題的技巧就是運用第二種想法來找到第 $k^{th}$ 小的數字：

- 每次決定好 `mid` 後都遍歷一次 `matrix` 計數有多少個比 `mid` 還小的數字，如果低於 `k`，則表示 `mid` 太小，要放寬區間 (i.e. `low = mid + 1`)
  - 題目說明 `matrix` 為 **non-decreasing order**，換句話說我們在遍歷時可以從每一個 `row` 的尾巴開始比較，加速統計過程

現在我們來看最後一個問題，為什麼能確保最終 **value of low** 一定存在於 `matrix`？

- 起初 `[matrix[0][0], mid]` 範圍內的數量小於 `k`，因此我們將範圍擴大到 `[matrix[0][0], mid+1]`，這時範圍內的數量滿足了，我們可以表示 `mid+1` 一定存在於 `matrix` 內，否則這個情況不可能發生
- 因此我們要做的事情就是不斷的調整 `low`, `high` 縮小 `mid` 的可能範圍，直到 `low = high` 時表示 `mid` 必定為我們要的值