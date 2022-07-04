[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/3sum/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 15. 3Sum

Given an integer array nums, return all the triplets `[nums[i], nums[j], nums[k]]` such that `i != j`, `i != k`, and `j != k`, and `nums[i] + nums[j] + nums[k] == 0`.

Notice that the solution set must not contain duplicate triplets.

### Example 1:

```
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
```

### Example 2:

```
Input: nums = []
Output: []
```

### Example 3:

```
Input: nums = [0]
Output: []
```

### Constraints:

- `0 <= nums.length <= 3000`
- -$10^5$ <= nums[i] <= $10^5$

### Related Topics

- Array
- Two Pointers
- Sorting

---

# 解題方向

因為題目給定的 `nums` 是無序的，因此第一件事情就是先對 `nums` 排序

排序好了之後就可以用 Two Pointers 的想法切入，把複雜度從 O($n^3$) 降到 O($n^2$)，想法如下：

- 第一個 foreach loop 固定第一位數字 -> O(n)
- 第二跟第三位數字靠 Two Pointers (`head`, `tail`) 一起尋找 -> O(n)

i.e.

 - `3Sum` = `nums[idx]` + `nums[head]` + `nums[tail]`

e.g.

```
+-------+-------+-------+-------+-------+-------+
| fixed | head  |       |       |       |  tail |
+-------+-------+-------+-------+-------+-------+
|  -4   |  -1   |  -1   |   0   |   1   |   2   |
+-------+-------+-------+-------+-------+-------+

3Sum = (-4) + (-1) + 2 != 0 (not found, moving head pointer to next distinct value)

+-------+-------+-------+-------+-------+-------+
| fixed |       |       | head  |       |  tail |
+-------+-------+-------+-------+-------+-------+
|  -4   |  -1   |  -1   |   0   |   1   |   2   |
+-------+-------+-------+-------+-------+-------+

3Sum = (-4) + 0 + 2 != 0 (not found, keep moving ...)

+-------+-------+-------+-------+-------+-------+
| fixed |       |       |       | head  |  tail |
+-------+-------+-------+-------+-------+-------+
|  -4   |  -1   |  -1   |   0   |   1   |   2   |
+-------+-------+-------+-------+-------+-------+

3Sum = (-4) + 1 + 2 != 0 (not found, moving head pointer and fixed pointer)

+-------+-------+-------+-------+-------+-------+
|       | fixed | head  |       |       |  tail |
+-------+-------+-------+-------+-------+-------+
|  -4   |  -1   |  -1   |   0   |   1   |   2   |
+-------+-------+-------+-------+-------+-------+

3Sum = (-1) + (-1) + 2 == 0 (congrats!)
```