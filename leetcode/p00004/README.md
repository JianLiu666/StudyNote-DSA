[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/median-of-two-sorted-arrays/)
![image](https://img.shields.io/badge/Difficulty-Hard-red)

---

# 4. Median of Two Sorted Arrays

Given two sorted arrays `nums1` and `nums2` of size `m` and `n` respectively, return **the median** of the two sorted arrays.

The overall run time complexity should be `O(log (m+n))`.

### Example 1:

```
Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000

Explanation:
 - merged array = [1,2,3] and median is 2.
```

### Example 2:

```
Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000

Explanation:
 - merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
```

### Constraints:

- `nums1.length == m`
- `nums2.length == n`
- `0 <= m <= 1000`
- `0 <= n <= 1000`
- `1 <= m + n <= 2000`
- -$10^6$ <= nums1[i], nums2[i] <= $10^6$

### Related Topics

- Array
- Binary Search
- Divide and Conquer
  
---

# 解題方向

LeetCodeCN 的這篇[精選解答](https://leetcode.cn/problems/median-of-two-sorted-arrays/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by-w-2/)寫得不錯可以參考

這題感覺只要過一陣子不寫就會忘記了QQ

**解法三:**

我們可以把中位數問題轉換成找到第 `k` 與 `k+1` 個位置的數字，在排除資料時用 Binary Search 的思路優化，加速收斂