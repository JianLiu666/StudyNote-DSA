[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/rotate-array/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 189. Rotate Array

Given an array, rotate the array to the right by `k` steps, where `k` is non-negative.

### Example 1:

```
Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]

Explanation:
 - rotate 1 steps to the right: [7,1,2,3,4,5,6]
 - rotate 2 steps to the right: [6,7,1,2,3,4,5]
 - rotate 3 steps to the right: [5,6,7,1,2,3,4]
```

### Example 2:

```
Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]

Explanation: 
 - rotate 1 steps to the right: [99,-1,-100,3]
 - rotate 2 steps to the right: [3,99,-1,-100]
```

### Constraints:

 - 1 <= nums.length <= $10^5$
 - -$2^{31}$ <= nums[i] <= $2^{31}$ - 1
 - 0 <= k <= $10^5$

#### Follow up:

- Try to come up with as many solutions as you can. There are at least **three** different ways to solve this problem.
- Could you do it in-place with `O(1)` extra space?

### Related Topics

- Array
- Math
- Two Pointers
  
---

# 解題方向

### Solved using reversion concept

Two Pointers (head and tail pointers)

- Step.1 將整個 `nums` 反轉一遍
- Step.2 將前半部的 `nums[:k-1]` 反轉回來
- Step.3 將後半部的 `nums[k:]` 反轉回來

### Solved using cyclic replacement concept

從題目的規律來看，`k` 表示數字往後移動的位置，換句話說只要不斷地將數字往後移動 `k` 格，直到全部移動過一遍即可

直接上圖:

```
if k = 2

   step.1 
  ┌-------┐
  |       v                               tmp=3
+---+---+---+---+---+---+---+    +---+---+---+---+---+---+---+
| 1 | 2 | 3 | 4 | 5 | 6 | 7 | -> | 1 | 2 | 1 | 4 | 5 | 6 | 7 |
+---+---+---+---+---+---+---+    +---+---+---+---+---+---+---+

           step.2 
          ┌-------┐
          |       v                               tmp=5
+---+---+---+---+---+---+---+    +---+---+---+---+---+---+---+
| 1 | 2 | 1 | 4 | 5 | 6 | 7 | -> | 1 | 2 | 1 | 4 | 3 | 6 | 7 |
+---+---+---+---+---+---+---+    +---+---+---+---+---+---+---+

                   step.3 
                  ┌-------┐
                  |       v                               tmp=7
+---+---+---+---+---+---+---+    +---+---+---+---+---+---+---+
| 1 | 2 | 1 | 4 | 3 | 6 | 7 | -> | 1 | 2 | 1 | 4 | 3 | 6 | 5 |
+---+---+---+---+---+---+---+    +---+---+---+---+---+---+---+

             step.4 
      ┌-------------------┐
      v                   |           tmp=2
+---+---+---+---+---+---+---+    +---+---+---+---+---+---+---+
| 1 | 2 | 1 | 4 | 3 | 6 | 5 | -> | 1 | 7 | 1 | 4 | 3 | 6 | 5 |
+---+---+---+---+---+---+---+    +---+---+---+---+---+---+---+

       step.5 
      ┌-------┐
      |       v                               tmp=4
+---+---+---+---+---+---+---+    +---+---+---+---+---+---+---+
| 1 | 7 | 1 | 4 | 3 | 6 | 5 | -> | 1 | 7 | 1 | 2 | 3 | 6 | 5 |
+---+---+---+---+---+---+---+    +---+---+---+---+---+---+---+

               step.6 
              ┌-------┐
              |       v                               tmp=6
+---+---+---+---+---+---+---+    +---+---+---+---+---+---+---+
| 1 | 7 | 1 | 2 | 3 | 6 | 5 | -> | 1 | 7 | 1 | 2 | 3 | 4 | 5 |
+---+---+---+---+---+---+---+    +---+---+---+---+---+---+---+

         step.7
  ┌-------------------┐
  v                   |           tmp=1
+---+---+---+---+---+---+---+    +---+---+---+---+---+---+---+
| 1 | 7 | 1 | 2 | 3 | 4 | 5 | -> | 6 | 7 | 1 | 2 | 3 | 4 | 5 |
+---+---+---+---+---+---+---+    +---+---+---+---+---+---+---+
```