[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/uncrossed-lines/description/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 1035. Uncrossed Lines

You are given two integer arrays `nums1` and `nums2`. We write the integers of `nums1` and `nums2` (in the order they are given) on two separate horizontal lines.

We may draw connecting lines: a straight line connecting two numbers `nums1[i]` and `nums2[j]` such that:

- `nums1[i] == nums2[j]`, and
- the line we draw does not intersect any other connecting (non-horizontal) line.
No
te that a connecting line cannot intersect even at the endpoints (i.e., each number can only belong to one connecting line).

Return *the maximum number of connecting lines we can draw in this way*.

### Example 1:

![image](./image/142.png)

```
Input: nums1 = [1,4,2], nums2 = [1,2,4]
Output: 2

Explanation: 
 - We can draw 2 uncrossed lines as in the diagram.
 - We cannot draw 3 uncrossed lines, because the line from nums1[1] = 4 to nums2[2] = 4 will intersect the line from nums1[2]=2 to nums2[1]=2.
```

### Example 2:

```
Input: nums1 = [2,5,1,2,5], nums2 = [10,5,2,1,5,2]
Output: 3
```

### Example 3:

```
Input: nums1 = [1,3,7,1,7,5], nums2 = [1,9,2,5,1]
Output: 2
```

### Constraints:

- `1 <= nums1.length, nums2.length <= 500`
- `1 <= nums1[i], nums2[j] <= 2000`

### Related Topics

- Array
- Dynamic Programming
  
---

# 解題方向

我們的目標是定義出 dynamic programing 的轉移方程，以本題為例，可以將 nums1, nums2 視為一個 matrix
以 example 1 為例，如下圖所示：

```
   1 4 2
  +- - -
1 |
2 |
3 |
```

當 `nums1[m]` 與 `nums2[n]` 等價時，我們可以累計一條連線數量
當 `nums1[m]` 與 `nums2[n]` 不相等時，我們就需要選擇保留之前較大的累計結果

所以換句話說，我們可以用下圖表示我們的目的：

```
最終目標是由左上至右下實際走訪一遍，並紀錄過程

   1 4 2
  +- - -
1 |S       S = start
2 |  ➘     E = end
3 |    E
```