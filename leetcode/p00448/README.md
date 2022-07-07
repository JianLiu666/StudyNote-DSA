[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 448. Find All Numbers Disappeared in an Array

Given an array `nums` of `n` integers where `nums[i]` is in the range `[1, n]`, return an array of all the integers in the range `[1, n]` that do not appear in `nums`.

### Example 1:

```
Input: nums = [4,3,2,7,8,2,3,1]
Output: [5,6]
```

### Example 2:

```
Input: nums = [1,1]
Output: [2]
```

### Constraints:

- `n == nums.length`
- 1 <= n <= $10^5$
- `1 <= nums[i] <= n`

#### Follow up: 

Could you do it without extra space and in `O(n)` runtime? You may assume the returned list does not count as extra space.

### Related Topics

- Array
- Hash Table

---

# 解題方向

善用 in-place method 處理 Array 的技巧 : Hash

e.g.

如果我們想知道 `arr1` 中的每個數字出現了幾次，我們可以用另一個同樣長度的 `arr2` 來表示每個數字的出現次數，即：

```python
arr1 = [4, 3, 2, 7, 8, 2, 3, 1]
arr2 = [0] * len(arr1)

print(arr1) # [4, 3, 2, 7, 8, 2, 3, 1]
print(arr2) # [0, 0, 0, 0, 0, 0, 0, 0]

for num in nums:
    arr2[num-1] += 1

print(arr1) # [4, 3, 2, 7, 8, 2, 3, 1]
print(arr2) # [1, 2, 2, 1, 0, 0, 1, 1] -> 可以發現 arr1 中沒有出現過 5, 6 這兩個數字
```

現在來看看怎麼只用一個陣列就做到這件事情，讓 `arr1` 不只保存原始資料同時也能保存出現次數

```python
arr = [4, 3, 2, 7, 8, 2, 3, 1]
size = len(arr) + 1

for i in range(len(nums)):
    arr[(arr[i]%size)-1] += size

print(arr) # [13, 21, 20, 16, 8, 2, 12, 10], 同時包含原始資料與出現次數

counter = []
for num in arr:
    counter.append(num//size)

print(counter) # [1, 2, 2, 1, 0, 0, 1, 1], 每個次數出現次數

raw = []
for num in nums:
    raw.append(num%size)

print(raw) # [4, 3, 2, 7, 8, 2, 3, 1], 原始資料
```