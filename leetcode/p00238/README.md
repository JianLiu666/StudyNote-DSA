[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/product-of-array-except-self/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 238. Product of Array Except Self

Given an integer array `nums`, return an array `answer` such that `answer[i]` is equal to the product of all the elements of `nums` except `nums[i]`.

The product of any prefix or suffix of `nums` is **guaranteed** to fit in a **32-bit** integer.

You must write an algorithm that runs in `O(n)` time and without using the division operation.

### Example 1:

```
Input: nums = [1,2,3,4]
Output: [24,12,8,6]
```

### Example 2:

```
Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]
```

### Constraints:

- 2 <= nums.length <= $10^5$
- `-30 <= nums[i] <= 30`
- The product of any prefix or suffix of `nums` is **guaranteed** to fit in a **32-bit** integer.

#### Follow up: 

Can you solve the problem in `O(1)` extra space complexity? (The output array **does not** count as extra space for space complexity analysis.)

### Related Topics

- Array
- Prefix Sum
  
---

# 解題方向

### Solved using Brute-Force concept

- 第一遍對 `nums` 遍歷時計算所有非零元素的乘積，與 `0` 出現的次數
- 如果 `nums` 內的 `0` 有兩個以上時表示 `nums` 的結果為 `0`，可以直接返回長度為 `nums` 的空陣列
- 否則就直接將 `nums[i]` 的元素更新為 `product / nums[i]` 即可 (若遇到 `0` 時則直接替換成 `product`) 

### Solved using Prefix Sum concept 

另一個比較優雅的寫法：

- 第一次遍歷時從頭開始做一次 **prefix sum**
- 第二次遍歷時從尾巴再做一次 **prefix sum**

因此 `nums[i]` 的資料就包含了除了自己位置以外的乘積結果
