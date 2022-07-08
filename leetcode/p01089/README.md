[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/duplicate-zeros/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 1089. Duplicate Zeros

Given a fixed-length integer array `arr`, duplicate each occurrence of zero, shifting the remaining elements to the right.

**Note** that elements beyond the length of the original array are not written. Do the above modifications to the input array in place and do not return anything.

### Example 1:

```
Input: arr = [1,0,2,3,0,4,5,0]
Output: [1,0,0,2,3,0,0,4]

Explanation:
 - After calling your function, the input array is modified to: [1,0,0,2,3,0,0,4]
```

### Example 2:

```
Input: arr = [1,2,3]
Output: [1,2,3]

Explanation:
 - After calling your function, the input array is modified to: [1,2,3]
```

### Constraints:

- 1 <= arr.length <= $10^4$
- `0 <= arr[i] <= 9`

### Related Topics

- Array
- Two Pointers
  
---

# 解題方向

題目要求的輸出結果是 fixed-length integer array，也就是說處理完畢後的陣列會因為補零導致原始資料遺失

首先我們必須先決定出要從哪裡開始作為起點往後複製，也就是 `arr` 中出現的 0 的**有效數量**

e.g.

```
                    v: 這個 0 是會被複製到的
arr1 = [1,5,2,0,6,8,0,6,0]
res  = [1,5,2,0,0,6,8,0,0]

                      v: 這個 0 是不會被複製到的
arr2 = [1,5,2,0,6,8,1,0,0]
res  = [1,5,2,0,0,6,8,1,0]
```

#### Two Pointers (head and tail pointer)

將 head pointer 定位在正確位置後，開始由後往前將 head pointer 的資料複製到 tail pointer，

e.g.

```
step.1
                                          v : head
         +-----+-----+-----+-----+-----+-----+-----+-----+
  arr1 = |  1  |  0  |  2  |  3  |  0  |  4  |  5  |  4  |
         +-----+-----+-----+-----+-----+-----+-----+-----+
                                                      ^ : tail (copied from arr1[5])

step.2-1
                                    v : head
         +-----+-----+-----+-----+-----+-----+-----+-----+
  arr1 = |  1  |  0  |  2  |  3  |  0  |  4  |  0  |  4  |
         +-----+-----+-----+-----+-----+-----+-----+-----+
                                                ^ : tail (copied from arr1[4])

step.2-2 : 遇到 0 多複製一次
                                    v : head
         +-----+-----+-----+-----+-----+-----+-----+-----+
  arr1 = |  1  |  0  |  2  |  3  |  0  |  0  |  0  |  4  |
         +-----+-----+-----+-----+-----+-----+-----+-----+
                                          ^ : tail (copied from arr1[4])

step.3
                              v : head
         +-----+-----+-----+-----+-----+-----+-----+-----+
  arr1 = |  1  |  0  |  2  |  3  |  3  |  0  |  0  |  4  |
         +-----+-----+-----+-----+-----+-----+-----+-----+
                                    ^ : tail (copied from arr1[3])

step.4
                        v : head
         +-----+-----+-----+-----+-----+-----+-----+-----+
  arr1 = |  1  |  0  |  2  |  2  |  3  |  0  |  0  |  4  |
         +-----+-----+-----+-----+-----+-----+-----+-----+
                              ^ : tail (copied from arr1[2])

step.5-1
                  v : head
         +-----+-----+-----+-----+-----+-----+-----+-----+
  arr1 = |  1  |  0  |  0  |  2  |  3  |  0  |  0  |  4  |
         +-----+-----+-----+-----+-----+-----+-----+-----+
                        ^ : tail (copied from arr1[1])

step.5-2 : 追到 head, 可以不用再往下處理了
                  v : head
         +-----+-----+-----+-----+-----+-----+-----+-----+
  arr1 = |  1  |  0  |  0  |  2  |  3  |  0  |  0  |  4  |
         +-----+-----+-----+-----+-----+-----+-----+-----+
                  ^ : tail
```