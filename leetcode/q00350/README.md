# `Easy` Intersection of Two Arrays II

Given two integer arrays `nums1` and `nums2`, return an array of their intersection. Each element in the result must appear as many times as it shows in both arrays and you may return the result in **any order**.

### Example 1:

```
Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2,2]
```

### Example 2:

```
Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [4,9]
Explanation:
 - [9,4] is also accepted.
``` 

### Constraints:

- `1 <= nums1.length, nums2.length <= 1000`
- `0 <= nums1[i], nums2[i] <= 1000`
 

#### Follow up:

- What if the given array is already sorted? How would you optimize your algorithm?
- What if `nums1`'s size is small compared to `nums2`'s size? Which algorithm is better?
- What if elements of `nums2` are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?

### Related Topics:

- Array
- Hash Table
- Two Pointers
- Binary Search
- Sorting

---

# 解題方向

### Considered Time Complexity

Table-Table Join

- 透過 hash map 將兩邊的陣列以及出現次數都記錄下來，再比對兩個 Map 重疊的部分。
  - e.g., dict1 : { `num` : `count(num)` }

Stream-Table Join

- 如果考慮到 follow up 情境中的記憶體限制條件，無法一次將 `nums2` 的資料一次保存在記憶體時，可以只將 `nums1` 的資料用 hash map 保存。

### Considered Space Complexity

先將兩個陣列排序過後，在用 two pointers 的概念依序比較，直到其中一個陣列先到盡頭為止。

e.g. 
- if `nums1[pointer1]` == `nums2[pointer2]`, 表示為交集。
- if `nums1[pointer1]` > `nums2[pointer2]`, moving `pointer2`
- if `nums1[pointer1]` < `nums2[pointer2]`, moving `pointer1`