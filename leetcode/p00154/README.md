[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/)
![image](https://img.shields.io/badge/Difficulty-Hard-red)

---

# 154. Find Minimum in Rotated Sorted Array II

Suppose an array of length `n` sorted in ascending order is **rotated** between `1` and `n` times. For example, the array `nums = [0,1,4,4,5,6,7]` might become:

- `[4,5,6,7,0,1,4]` if it was rotated `4` times.
- `[0,1,4,4,5,6,7]` if it was rotated `7` times.

Notice that **rotating** an array `[a[0], a[1], a[2], ..., a[n-1]]` 1 time results in the array `[a[n-1], a[0], a[1], a[2], ..., a[n-2]]`.

Given the sorted rotated array `nums` that may contain **duplicates**, return the minimum element of this array.

You must decrease the overall operation steps as much as possible.

### Example 1:

```
Input: nums = [1,3,5]
Output: 1
```

### Example 2:

```
Input: nums = [2,2,2,0,1]
Output: 0
```

### Constraints:

- `n == nums.length`
- `1 <= n <= 5000`
- `-5000 <= nums[i] <= 5000`
- `nums` is sorted and rotated between `1` and `n` times.

#### Follow up:

This problem is similar to [Find Minimum in Rotated Sorted Array](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/), but `nums` may contain **duplicates**. Would this affect the runtime complexity? How and why?

### Related Topics

- Array
- Binary Search
  
---

# 解題方向

[153. Find Minimum in Rotated Sorted Array](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/) 的延伸問題，如果 `Array` 內的元素有重複的話該如何解決？

Example :

```
給定一個 array 如下：

   head                mid                     tail
    v                   v                       v
  +---+---+---+---+---+---+---+---+---+---+---+---+
  | 3 | 3 | 3 | 3 | 3 | 3 | 3 | 3 | 1 | 3 | 3 | 3 |
  +---+---+---+---+---+---+---+---+---+---+---+---+
```

- `if nums[mid] > nums[tail] then tail = mid` : 這沒問題
- `if nums[mid] < nums[tail] then head = mid + 1` : 這也沒有問題
- `when nums[mid] == nums[tail]` : 如果直接讓 `tail = mid`，就會錯過正確的元素了，此時只能逐步的將 tail 往前位移直到不是重複的元素，或與 `mid` 重疊為止

因此，在最壞的情況下會從 Binaray Search 的 `O(logn)` 退化成 Linear Search 的 `O(n)`