[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/sliding-window-maximum/description/)
![image](https://img.shields.io/badge/Difficulty-Hard-red)

---

# 239. Sliding Window Maximum

You are given an array of integers `nums`, there is a sliding window of size `k` which is moving from the very left of the array to the very right. You can only see the `k` numbers in the window. Each time the sliding window moves right by one position.

Return the max sliding window.

### Example 1:

```
Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
Output: [3,3,5,5,6,7]
Explanation: 
  Window position                Max
  ---------------               -----
  [1  3  -1] -3  5  3  6  7       3
   1 [3  -1  -3] 5  3  6  7       3
   1  3 [-1  -3  5] 3  6  7       5
   1  3  -1 [-3  5  3] 6  7       5
   1  3  -1  -3 [5  3  6] 7       6
   1  3  -1  -3  5 [3  6  7]      7
```

### Example 2:

```
Input: nums = [1], k = 1
Output: [1]
```

### Constraints:

- 1 <= nums.length <= $10^5$
- -$10^4$ <= nums[i] <= $10^4$
- `1 <= k <= nums.length`

### Related Topics

- Array
- Queue
- Sliding Window
- Heap (Priority Queue)
- Monotonic Queue
  
---

# 解題方向

兩種解法:

一種是對 sliding window 建立 BST 維護, 所以每次查找時的時間複雜度都是 `O(logk)`，但 golang 不提供這種資料結構  
最接近的是 goDS 提供的 `treeset`，但是 `treeset` 不能處理 duplicate values (請移駕到 C++ 的 `multiset`)  

另一種做法就是用 dequeue 維護 sliding window，確保 sliding window 最大的數放在 head，且不斷確認 index 是否超出 window 範圍  
詳細流程請直接看程式碼