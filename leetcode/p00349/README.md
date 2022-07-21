[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/intersection-of-two-arrays/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 349. Intersection of Two Arrays

Given two integer arrays `nums1` and `nums2`, return an array of their intersection. Each element in the result must be **unique** and you may return the result in **any order**.

### Example 1:

```
Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]
```

### Example 2:

```
Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]

Explanation:
 - [4,9] is also accepted.
```

### Constraints:

- `1 <= nums1.length, nums2.length <= 1000`
- `0 <= nums1[i], nums2[i] <= 1000`

### Related Topics

- Array
- Hash Table
- Two Pointers
- Binary Search
- Sorting
  
---

# 解題方向

**Table-Table Join**

直覺解法，用兩個 `Hash Set` 將資料保存後依序比對找出重疊部分

- Time Complexity: `O(m+2n)`, suppose `n` <= `m`
- Space Complexity: `O(m+n)`

**Stream-Table Join**

只選擇將較小的 `Array` 保存在 `Hash Set`，然後對另一個 `Array` 尋訪一次逐一比對

- Time Complexity: `O(m+n)`
- Space Complexity: `O(n)`, suppose `n` <= `m`
