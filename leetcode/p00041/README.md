[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/first-missing-positive/)
![image](https://img.shields.io/badge/Difficulty-Hard-red)

---

# 41. First Missing Positive

Given an unsorted integer array `nums`, return the smallest missing positive integer.

You must implement an algorithm that runs in `O(n)` time and uses constant extra space.

### Example 1:

```
Input: nums = [1,2,0]
Output: 3
```

### Example 2:

```
Input: nums = [3,4,-1,1]
Output: 2
```

### Example 3:

```
Input: nums = [7,8,9,11,12]
Output: 1
```

### Constraints:

- 1 <= nums.length <= 5 * $10^5$
- -$2^{31}$ <= nums[i] <= $2^{31}$ - 1

### Related Topics

- Array
- Hash Table

### Food for thought

- Think about how you would solve the problem in non-consistent space. Can you apply that logic to the existing space?
- We don't care about duplicates or non-positive integers
- Remember that `O(2n)` = `O(n)`

---

# 解題方向

因為題目同時規定空間複雜度要符合 constant extra space 的限制，所以用 Hash Map 查找的解法可以放棄了 ...

只能依靠陣列本身的空間來做判斷，也算是很常出現的解題方向 (in-place algorithm)

以下有兩種切入方式：

### Solved using Swap concept

我們的目標是將所有落在 `0 < nums[i] <= len(nums)` 的資料都移動到正確的位置上，即 `nums[i] == nums[nums[i]-1]`
- e.g., `if nums = [1,2,3,4], where i = 0 then nums[0] == nums[nums[0]-1]`

所以我們可以依序對 `nums` 內的所有位置都檢查一遍，直到符合規範的數字都已經落到正確的位置上。

```python
for i in range(len(nums)):
    # 乍看之下這邊看起來會很像是 O(n^2) 的操作，但我們的目標是將不再正確位置上的數字移動到正確的位置
    # 也就是說，只要數字經過一次移動後就不需要再交換了，因此實際情況下兩個迴圈的操作次數會落在 N <= operations <= 2N-1
    while nums[i] > 0 and nums[i] <= len(nums) and nums[i] != nums[nums[i]-1]:
        nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
```

### Solved using Hash concept

首先對 `nums` 進行第一次整理，把超出範圍的數字歸零，確保後續的 hashing 都只專注在條件範圍內的數字。
- e.g., `if nums[i] < 0 or nums[i] > len(nums) then nums[i] = 0`

接下來就是這個做法的精彩之處了，將 `nums` 內非零數字的出現次數計數在其**正確的位置上**，即：
- `nums[nums[i] % (len(nums)+1)] += len(nums)+1`

沒事，我當初也看了很久，讓我們一步一步來：

```python
nums = [1, 1, 1, 2, 2]
size = len(nums) # 5

for num in nums:
    # where num = 1 (num[0]), nums = [1,  6,  1, 2, 2]
    # where num = 1 (num[1]), nums = [1, 11,  1, 2, 2]
    # where num = 1 (num[2]), nums = [1, 16,  1, 2, 2]
    # where num = 2 (num[3]), nums = [1, 16,  6, 2, 2]
    # where num = 2 (num[4]), nums = [1, 16, 11, 2, 2]
    nums[num%size] += size

print(nums) # hashed: [1, 16, 11, 2, 2]
```

最後，我們只要從 `nums[1]` 開始往後依序尋找，找到第一個出現頻率為零 (i.e., `nums[i] / len(nums)`) 的位置，就是我們要找的答案。

這個方法最精彩的地方在於不只同時紀錄每個數字的出現頻率，也變相地把原始資料 (i.e., `nums[i] % len(nums)`)也保存下來。