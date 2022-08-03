[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/find-the-duplicate-number/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 287. Find the Duplicate Number

Given an array of integers `nums` containing `n + 1` integers where each integer is in the range `[1, n]` inclusive.

There is only **one repeated number** in `nums`, return this repeated number.

You must solve the problem **without** modifying the array `nums` and uses only constant extra space.

### Example 1:

```
Input: nums = [1,3,4,2,2]
Output: 2
```

### Example 2:

```
Input: nums = [3,1,3,4,2]
Output: 3
```

### Constraints:

- 1 <= n <= $10^5$
- `nums.length == n + 1`
- `1 <= nums[i] <= n`
- All the integers in `nums` appear only **once** except for **precisely one integer** which appears **two or more** times.

#### Follow up:

- How can we prove that at least one duplicate number must exist in `nums`?
- Can you solve the problem in linear runtime complexity?


### Related Topics

- Array
- Two Pointers
- Binary Search
- Bit Manipulation
  
---

# 解題方向

### Solved using Binary Search concept

`Binary Search` 的變化題，在不能排序的前提下，Search space 不在是常見的 `index` 為主，而是要考慮整個 `array(range)`，所以時間複雜度會是 `O(nlogn)`

我們已知資料範圍在 `[1, n]` 之間，每一次決定好 `mid` 後，都對整個 `array` 遍歷一次計數有多少筆資料比 `mid` 還要小

- 如果 `count(nums, mid) <= mid` 的話表示 `[1, mid]` 這段資料範圍不存在重複的數字可以排除，即 `low = mid+1`
- 如果 `count(nums, mid) > mid` 的話表示 `[1, mid]` 這段資料範圍包含重複的數字，則必須收斂範圍後繼續檢查，即 `high = mid`

隨著資料空間的逐步收斂，當 `low = high` 時我們可以確定這個數字就是重複數字！

類似的題目像是這題 [378. Kth Smallest Element in a Sorted Matrix](./../p00378/README.md)

### Solved using Cycle Detection concept

鬼神等級的想法，用 `Linked List` 的 `Cycle Detection` 來解決這個問題，因為資料範圍落在 `[1, n]` 之間，因此 `array` 內的資料可以等價成 **node pointer reference**，重複數字表示陷入循環

因此我們可以先用 Two Pointers 的想法切入找到環之後，再透過一個 `pointer` 從原點出發，直到跟 `slow pointer` 相撞時的那個數字即為重複數字