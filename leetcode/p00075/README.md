[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/sort-colors/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 75. Sort Colors

Given an array `nums` with `n` objects colored red, white, or blue, sort them [in-place](https://en.wikipedia.org/wiki/In-place_algorithm) so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

We will use the integers `0`, `1`, and `2` to represent the color red, white, and blue, respectively.

You must solve this problem without using the library's sort function.

### Example 1:

```
Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
```

### Example 2:

```
Input: nums = [2,0,1]
Output: [0,1,2]
```

### Constraints:

- `n == nums.length`
- `1 <= n <= 300`
- `nums[i]` is either `0`, `1`, or `2`.

#### Follow up: 

Could you come up with a one-pass algorithm using only constant extra space?

### Related Topics

- Array
- Two Pointers
- Sorting
  
---

# 解題方向

### Solved using Counting Sort concept

題目說明 `nums` 內的元素只有 `0`, `1`, `2` 三種可能，但 follow up 限制我們只能用有限的資料空間來處理，直接用 **counting sort** 解決即可

### Solved using Two Pointers concept

根據題目描述，我們可以建立三種指標 `low`, `mid`, `high` 來分別對應三個數字，核心概念是不斷的把 `0` 往前丟，`2` 往後丟，直到所有數字都整理完畢

詳細流程直接看程式碼吧