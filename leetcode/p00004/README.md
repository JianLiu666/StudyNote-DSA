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

#### 2026.05.10 重刷 NeetCode 150 時卡關看解答的做法

```
分析問題:
- 已知 nums1, nums2 是兩個不一定等長的 ascending order array, 找到 p50 返回
- 如果元素總數量是偶數, 就要返回平均值

間單想法:
- median = (nums1.length + nums2.length + 1) // 2
- 可能落在 nums1 or nums2

視覺化來看:

nums1: [ ... mid1 ... ]
nums2: [ ... mid2 ... ], mid1 & mid2 左邊元素加總的數量要等於 median
- 現在的問題是不知道 nums1, nums2 應該要排除多少左邊數字？

用這個方式來看:
nums1: [ ... l1 | r1 ... ]
nums2: [ ... l2 | r2 ... ]
        左半部 ^   ^ 右半部

如果 l1 <= r2 && l2 <= r1, 找到完全符合的中位數
如果 l1 > r2, 表示 nums1 佔據太多元素, 需要收窄 mid1
反之如果 l2 > r1, 表示 nums2 佔據太多元素, 需要收窄 mid2

只要可以定位一邊的 nums, 就能同時定位兩邊
用長度較小的 array 來定位上下界 (low, high), 即:
- mid1 = (low + high) / 2
- mid2 = median - mid1
```

#### 2023.09.13 在刷一次後試著用自己的想法分析一遍  

```
分析問題:
題目限定要用 log(m+n) 的複雜度, 所以基本上的思路就只剩 binary search 了

考慮兩個 array 的可能情況
 - case 1: [x x x x x x x] [y y y y y y] 兩個 array 彼此之間沒有重疊
 - case 2: [x x x x x x x]     兩個 array 的 element 彼此重疊
               [y y y y y y y]

所以我們在不能保證 nums range 的情況下, 沒辦法直接只對一個 array 做 binary search

換個角度思考, 這個題目的問題等價於找到兩個 array 的所有元素中第 k 小的元素 (i.e. k = median)
要找到兩個 soreted array 中第 k 小的 element, 可以用 binary search 的概念處理

每一次都比較 nums1[k/2], nums2[k/2] 位置中的元素誰大誰小, 把小的那邊刪掉, 如下

if k = 7, k/2 = 3(無條件捨去)

            v
nums1: [1 2 3 x x x x]

            v
nums2: [2 3 4 y y y y]

如上圖所示, nums2 > nums1 (i.e. 4 > 3), 我們可以直接把 nums1 的 [1 2 3] 捨去, 變成:

nums1:       [x x x x]
nums2: [2 3 4 y y y y]

先證明這個做法有效且保證安全, 每次比較兩個 array 之間的第 k/2 小的元素時, 每一次只會刪去 k/2-1 或 k/2 個元素
例如: k=7, 這時比較完後會刪去3個元素

原理在於, 兩邊都是 soreted array, 所以當我們刪去較小的那邊時, 排除掉的就是 1, 2, 3, ..., k 的前三個數字
我們就可以接著更新 k 到 4, 在下一次迭代時的目標是找到兩個 array 中排序第 4 小的位置

迭代的終止條件:
 1. k = 1 時, 只要比較兩個 array 的第一個起始位置誰比較小, 誰就是答案
 2. 可能會發生其中有一邊已經被完全排空(i.e. len(nums) == 0)了, 那就只要返回另一個 array 的第 k-1 index 就是答案

另外要注意的事情是, 當兩個 sorted array 的總個數是偶數時, 我們需要把 median 的左右兩個鄰近元素取平均
所以可以做一個簡單的操作是, 先將 left, right 都取出來, 如果是奇數時, 那 left == right, 如果是偶數時也符合題目要求

left  = (len(nums1) + len(nums2) + 1) / 2 -> +1 的原因是因為要從 1 開始計數
right = (len(nums1) + len(nums2) + 2) / 2
```

#### previous  

LeetCodeCN 的這篇[精選解答](https://leetcode.cn/problems/median-of-two-sorted-arrays/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by-w-2/)寫得不錯可以參考

這題感覺只要過一陣子不寫就會忘記了QQ

**解法三:**

我們可以把中位數問題轉換成找到第 `k` 與 `k+1` 個位置的數字，在排除資料時用 Binary Search 的思路優化，加速收斂