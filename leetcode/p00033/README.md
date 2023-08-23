[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/search-in-rotated-sorted-array/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 33. Search in Rotated Sorted Array

There is an integer array `nums` sorted in ascending order (with **distinct** values).

Prior to being passed to your function, `nums` is **possibly rotated** at an unknown pivot index `k` (`1 <= k < nums.length`) such that the resulting array is `[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]` (**0-indexed**). For example, `[0,1,2,4,5,6,7]` might be rotated at pivot index `3` and become `[4,5,6,7,0,1,2]`.

Given the array `nums` **after** the possible rotation and an integer `target`, return the index of `target` if it is in `nums`, or `-1` if it is not in `nums`.

You must write an algorithm with `O(log n)` runtime complexity.
 
### Example 1:

```
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
```

### Example 2:

```
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
```

### Example 3:

```
Input: nums = [1], target = 0
Output: -1
```

### Constraints:

- `1 <= nums.length <= 5000`
- -$10^4$ <= nums[i] <= $10^4$
- All values of `nums` are **unique**.
- `nums` is an ascending array that is possibly rotated.
- -$10^4$ <= target <= $10^4$

### Related Topics

- Array
- Binary Search

---

# 解題方向

2023.08.23 更新
 - 詳細直接看這篇 [blog](https://www.cnblogs.com/grandyang/p/4325648.html) 吧, 講得非常清楚
 - 下面的以後有需要再回來整理

題目表示將一個有序的陣列隨機挑一個位置旋轉後得到的新陣列 `nums` 作為此題的輸入，也就是說我們可以把他視為兩個有序陣列來看，即：

```
+-------------------------------+      +-------------------------------+
|           nums                +      |   sub arr1    |   sub arr2    +
+---+---+---+---+---+---+---+---+  =>  +---+---+---+---+---+---+---+---+
| 4 | 5 | 6 | 7 | 0 | 1 | 2 | 3 |      | 4 | 5 | 6 | 7 | 0 | 1 | 2 | 3 |
+---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
```

這樣一來，只要在查找時多判斷現在是落在左右哪個陣列中，確認要使用的邊界條件後，就可以當作是基本的 Binary Search 問題了，核心概念如下：

```python

while left <= right:
    mid = left + (right-left) // 2
    
    if nums[mid] == target:
        # 直接命中，可以下班了
        return mid

    if nums[mid] >= nums[left]:
        # 表示 middle index 落在左邊的陣列
        if target < nums[left] or target > nums[mid]:
            # 示意圖
            #       +-------------------------------+        +-------------------------------+
            #       |   sub arr1    |   sub arr2    +        |   sub arr1    |   sub arr2    +
            #       +---+---+---+---+---+---+---+---+   =>   +---+---+---+---+---+---+---+---+
            #       | 4 | 5 | 6 | 7 | 0 | 1 | 2 | 3 |        | 4 | 5 | 6 | 7 | 0 | 1 | 2 | 3 |
            #       | l |   |   | m |   |   |   | r |        | x |   |   | m | l |   |   | r |
            #       +---+---+---+---+---+---+---+---+        +---+---+---+---+---+---+---+---+
            #   <---| out of range  |--->                              ^ new left index
            left = mid + 1
        else:
            right = mid - 1
    else:
        # 表示 middle index 落在右邊的陣列
        if target < nums[mid] or target > nums[right]:
            # 示意圖
            #       +-------------------------------+        +-------------------------------+
            #       |   sub arr1    |   sub arr2    +        |   sub arr1    |   sub arr2    +
            #       +---+---+---+---+---+---+---+---+   =>   +---+---+---+---+---+---+---+---+
            #       | 4 | 5 | 6 | 7 | 0 | 1 | 2 | 3 |        | 4 | 5 | 6 | 7 | 0 | 1 | 2 | 3 |
            #       |   | l |   |   | m |   |   | r |        |   | l |   | r | m |   |   | x |
            #       +---+---+---+---+---+---+---+---+        +---+---+---+---+---+---+---+---+
            #                   <---| out of range  |--->                  ^ new right index
            right = mid - 1
        else:
            left = mid + 1
```