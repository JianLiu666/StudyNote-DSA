[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/maximum-length-of-repeated-subarray/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 718. Maximum Length of Repeated Subarray

Given two integer arrays `nums1` and `nums2`, return the maximum length of a subarray that appears in **both** arrays.

### Example 1:

```
Input: nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
Output: 3

Explanation:
 - The repeated subarray with maximum length is [3,2,1].
```

### Example 2:

```
Input: nums1 = [0,0,0,0,0], nums2 = [0,0,0,0,0]
Output: 5
```

### Constraints:

- `1 <= nums1.length, nums2.length <= 1000`
- `0 <= nums1[i], nums2[i] <= 100`

### Related Topics

- Array
- Binary Search
- Dynamic Programming
- Sliding Window
- Rolling Hash
- Hash Function
  
---

# 解題方向

### Solved using Dynamic Programming concept

用一個 2D array 紀錄 subarray 的長度累計 + 兩層的 foreach 不斷移動 `nums1[i]` 遍歷 `nums2` 一邊紀錄最大長度